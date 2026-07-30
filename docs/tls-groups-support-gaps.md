# TLS Groups (Supported Curves) — Support Gaps

Tracking document for TLS groups propagation across all HCO-managed TLS endpoints.

The `openshift/api` PR [#2583](https://github.com/openshift/api/pull/2583) added a
`Groups []TLSGroup` field to `TLSProfileSpec`, enabling configuration of TLS
supported groups (elliptic curves) including post-quantum hybrids like `X25519MLKEM768`.

## Supported (this PR)

| Component | How TLS is configured | Groups support |
|---|---|---|
| HCO operator server | `MutateTLSConfig` sets `tls.Config` | `CurvePreferences` set from profile groups |
| HCO webhook server | `MutateTLSConfig` sets `tls.Config` | `CurvePreferences` set from profile groups |
| Nginx console plugin | Generated `nginx.conf` | `ssl_ecdh_curve` populated from profile groups |

## Automatically handled (no HCO changes needed)

| Component | How TLS is configured | Notes |
|---|---|---|
| SSP | Raw `openshiftconfigv1.TLSSecurityProfile` passed through | Groups flow through once SSP bumps its `openshift/api` dependency |
| CNAO | Raw `openshiftconfigv1.TLSSecurityProfile` passed through | Groups flow through once CNAO bumps its `openshift/api` dependency |

## Blocked on upstream operand API changes

These operands define their own `TLSProfileSpec` types that only have `Ciphers` and
`MinTLSVersion` — no `Groups` field. HCO's conversion functions cannot forward groups
until the upstream projects add the field to their APIs.

| Component | Operand API type | Upstream tracker |
|---|---|---|
| CDI | `cdiv1beta1.TLSProfileSpec` | TODO: file issue on containerized-data-importer |
| AAQ | `aaqv1alpha1.TLSProfileSpec` | TODO: file issue on application-aware-quota |
| Migration (Forklift) | `migrationv1alpha1.TLSProfileSpec` | TODO: file issue on kubevirt-migration-operator |
| KubeVirt | `kubevirtcorev1.TLSConfiguration` | TODO: file issue on kubevirt/kubevirt |

## Blocked on binary support (CLI-args-based components)

These components receive TLS configuration as container CLI arguments
(`--tls-cipher-suites`, `--tls-min-version`). Supporting groups would require each
binary to accept a new flag (e.g. `--tls-groups`).

| Component | Handler file | CLI args used |
|---|---|---|
| AIE webhook | `controllers/handlers/aie/deployment.go` | `--tls-min-version`, `--tls-cipher-suites` |
| Network Resources Injector | `controllers/handlers/netresinjector/deployment.go` | `-tls-min-version`, `-tls-cipher-suites` |
| KV UI Proxy | `controllers/handlers/kubevirtConsolePlugin.go` | `--tls-cipher-suites`, `--tls-min-version` |
| Observability Controller | `controllers/handlers/observabilitycontroller/deployment.go` | `--tls-security-profile`, `--tls-min-version`, `--tls-cipher-suites` |

## Testing gaps

The current TLS profile integration test (`hack/check_tlsprofile.sh`) uses nmap
`ssl-enum-ciphers` which only reports the negotiated group per cipher — it cannot
enumerate the full set of supported TLS groups. This means groups configuration
cannot be verified by the existing test.

A `custom-pqc` test was added to `hack/check_tlsprofile.sh` that sets a Custom
profile with TLS 1.3 + `[X25519MLKEM768, X25519]` groups. Since nmap's
`ssl-enum-ciphers` does not support X25519MLKEM768 in its `supported_groups`
extension, it falls back to negotiating X25519. The test verifies TLS 1.3 is
enforced and the groups configuration is active, but cannot confirm that a
PQC-capable client would actually negotiate X25519MLKEM768.

Options to achieve full PQC group verification:
- Replace nmap with [`openshift/tls-scanner`](https://github.com/openshift/tls-scanner)
  which wraps `testssl.sh` and supports full group enumeration plus a `-pqc-check`
  flag for post-quantum readiness. This would require rewriting `check_tlsprofile.sh`
  and all expected baseline files in `hack/tlsprofiles/`.
- A Go-based TLS client test that dials endpoints and inspects
  `tls.ConnectionState().CurveID` can verify which group was **negotiated** (e.g.
  confirming X25519MLKEM768 was selected), but cannot enumerate the full list of
  groups the server **supports** — Go's `tls.ConnectionState` only reports the
  single group that was actually used for the handshake. Enumerating all supported
  groups would require multiple connections with different `CurvePreferences` on the
  client side, which is fragile and incomplete. `openshift/tls-scanner` (via
  `testssl.sh`) is the more reliable option for full server-side group enumeration.

## Pending library-go integration

The TLSGroup→CurveID mapping functions (`CurveIDForTLSGroup`, `CurveIDsForTLSGroups`,
`ValidTLSGroups`) are implemented locally in `pkg/internal/tlssecprofile/tls_groups.go`,
API-compatible with [openshift/library-go PR #2347](https://github.com/openshift/library-go/pull/2347).

Once that PR merges and HCO bumps library-go, the local implementation should be
replaced with `github.com/openshift/library-go/pkg/crypto`.
