//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2022 Red Hat, Inc.
 *
 */

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.CertRotateConfigCA":                   schema_pkg_apis_hco_v1beta1_CertRotateConfigCA(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.CertRotateConfigServer":               schema_pkg_apis_hco_v1beta1_CertRotateConfigServer(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConverged":                       schema_pkg_apis_hco_v1beta1_HyperConverged(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedCertConfig":             schema_pkg_apis_hco_v1beta1_HyperConvergedCertConfig(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedFeatureGates":           schema_pkg_apis_hco_v1beta1_HyperConvergedFeatureGates(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedObsoleteCPUs":           schema_pkg_apis_hco_v1beta1_HyperConvergedObsoleteCPUs(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedSpec":                   schema_pkg_apis_hco_v1beta1_HyperConvergedSpec(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedStatus":                 schema_pkg_apis_hco_v1beta1_HyperConvergedStatus(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedWorkloadUpdateStrategy": schema_pkg_apis_hco_v1beta1_HyperConvergedWorkloadUpdateStrategy(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.LiveMigrationConfigurations":          schema_pkg_apis_hco_v1beta1_LiveMigrationConfigurations(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.MediatedDevicesConfiguration":         schema_pkg_apis_hco_v1beta1_MediatedDevicesConfiguration(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.MediatedHostDevice":                   schema_pkg_apis_hco_v1beta1_MediatedHostDevice(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.NodeMediatedDeviceTypesConfig":        schema_pkg_apis_hco_v1beta1_NodeMediatedDeviceTypesConfig(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.OperandResourceRequirements":          schema_pkg_apis_hco_v1beta1_OperandResourceRequirements(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.PciHostDevice":                        schema_pkg_apis_hco_v1beta1_PciHostDevice(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.PermittedHostDevices":                 schema_pkg_apis_hco_v1beta1_PermittedHostDevices(ref),
		"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.StorageImportConfig":                  schema_pkg_apis_hco_v1beta1_StorageImportConfig(ref),
	}
}

func schema_pkg_apis_hco_v1beta1_CertRotateConfigCA(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CertRotateConfigCA contains the tunables for TLS certificates.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"duration": {
						SchemaProps: spec.SchemaProps{
							Description: "The requested 'duration' (i.e. lifetime) of the Certificate. This should comply with golang's ParseDuration format (https://golang.org/pkg/time/#ParseDuration)",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
						},
					},
					"renewBefore": {
						SchemaProps: spec.SchemaProps{
							Description: "The amount of time before the currently issued certificate's `notAfter` time that we will begin to attempt to renew the certificate. This should comply with golang's ParseDuration format (https://golang.org/pkg/time/#ParseDuration)",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Duration"},
	}
}

func schema_pkg_apis_hco_v1beta1_CertRotateConfigServer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CertRotateConfigServer contains the tunables for TLS certificates.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"duration": {
						SchemaProps: spec.SchemaProps{
							Description: "The requested 'duration' (i.e. lifetime) of the Certificate. This should comply with golang's ParseDuration format (https://golang.org/pkg/time/#ParseDuration)",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
						},
					},
					"renewBefore": {
						SchemaProps: spec.SchemaProps{
							Description: "The amount of time before the currently issued certificate's `notAfter` time that we will begin to attempt to renew the certificate. This should comply with golang's ParseDuration format (https://golang.org/pkg/time/#ParseDuration)",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Duration"},
	}
}

func schema_pkg_apis_hco_v1beta1_HyperConverged(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HyperConverged is the Schema for the hyperconvergeds API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedSpec", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_hco_v1beta1_HyperConvergedCertConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HyperConvergedCertConfig holds the CertConfig entries for the HCO operands",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"ca": {
						SchemaProps: spec.SchemaProps{
							Description: "CA configuration - CA certs are kept in the CA bundle as long as they are valid",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.CertRotateConfigCA"),
						},
					},
					"server": {
						SchemaProps: spec.SchemaProps{
							Description: "Server configuration - Certs are rotated and discarded",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.CertRotateConfigServer"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.CertRotateConfigCA", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.CertRotateConfigServer"},
	}
}

func schema_pkg_apis_hco_v1beta1_HyperConvergedFeatureGates(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HyperConvergedFeatureGates is a set of optional feature gates to enable or disable new features that are not enabled by default yet.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"withHostPassthroughCPU": {
						SchemaProps: spec.SchemaProps{
							Description: "Allow migrating a virtual machine with CPU host-passthrough mode. This should be enabled only when the Cluster is homogeneous from CPU HW perspective doc here",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"sriovLiveMigration": {
						SchemaProps: spec.SchemaProps{
							Description: "Allow migrating a virtual machine with SRIOV interfaces.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"enableCommonBootImageImport": {
						SchemaProps: spec.SchemaProps{
							Description: "Opt-in to automatic delivery/updates of the common data import cron templates. There are two sources for the data import cron templates: hard coded list of common templates, and custom templates that can be added to the dataImportCronTemplates field. This feature gates only control the common templates. It is possible to use custom templates by adding them to the dataImportCronTemplates field.",
							Default:     false,
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_hco_v1beta1_HyperConvergedObsoleteCPUs(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HyperConvergedObsoleteCPUs allows avoiding scheduling of VMs for obsolete CPU models",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"minCPUModel": {
						SchemaProps: spec.SchemaProps{
							Description: "MinCPUModel is the Minimum CPU model that is used for basic CPU features; e.g. Penryn or Haswell. The default value for this field is nil, but in KubeVirt, the default value is \"Penryn\", if nothing else is set. Use this field to override KubeVirt default value.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"cpuModels": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "CPUModels is a list of obsolete CPU models. When the node-labeller obtains the list of obsolete CPU models, it eliminates those CPU models and creates labels for valid CPU models. The default values for this field is nil, however, HCO uses opinionated values, and adding values to this list will add them to the opinionated values.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_hco_v1beta1_HyperConvergedSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HyperConvergedSpec defines the desired state of HyperConverged",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"localStorageClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "LocalStorageClassName the name of the local storage class.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"infra": {
						SchemaProps: spec.SchemaProps{
							Description: "infra HyperConvergedConfig influences the pod configuration (currently only placement) for all the infra components needed on the virtualization enabled cluster but not necessarely directly on each node running VMs/VMIs.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedConfig"),
						},
					},
					"workloads": {
						SchemaProps: spec.SchemaProps{
							Description: "workloads HyperConvergedConfig influences the pod configuration (currently only placement) of components which need to be running on a node where virtualization workloads should be able to run. Changes to Workloads HyperConvergedConfig can be applied only without existing workload.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedConfig"),
						},
					},
					"featureGates": {
						SchemaProps: spec.SchemaProps{
							Description: "featureGates is a map of feature gate flags. Setting a flag to `true` will enable the feature. Setting `false` or removing the feature gate, disables the feature.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedFeatureGates"),
						},
					},
					"liveMigrationConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "Live migration limits and timeouts are applied so that migration processes do not overwhelm the cluster.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.LiveMigrationConfigurations"),
						},
					},
					"permittedHostDevices": {
						SchemaProps: spec.SchemaProps{
							Description: "PermittedHostDevices holds information about devices allowed for passthrough",
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.PermittedHostDevices"),
						},
					},
					"mediatedDevicesConfiguration": {
						SchemaProps: spec.SchemaProps{
							Description: "MediatedDevicesConfiguration holds information about MDEV types to be defined on nodes, if available",
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.MediatedDevicesConfiguration"),
						},
					},
					"certConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "certConfig holds the rotation policy for internal, self-signed certificates",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedCertConfig"),
						},
					},
					"resourceRequirements": {
						SchemaProps: spec.SchemaProps{
							Description: "ResourceRequirements describes the resource requirements for the operand workloads.",
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.OperandResourceRequirements"),
						},
					},
					"scratchSpaceStorageClass": {
						SchemaProps: spec.SchemaProps{
							Description: "Override the storage class used for scratch space during transfer operations. The scratch space storage class is determined in the following order: value of scratchSpaceStorageClass, if that doesn't exist, use the default storage class, if there is no default storage class, use the storage class of the DataVolume, if no storage class specified, use no storage class for scratch space",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"vddkInitImage": {
						SchemaProps: spec.SchemaProps{
							Description: "VDDK Init Image eventually used to import VMs from external providers",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"defaultCPUModel": {
						SchemaProps: spec.SchemaProps{
							Description: "DefaultCPUModel defines a cluster default for CPU model: default CPU model is set when VMI doesn't have any CPU model. When VMI has CPU model set, then VMI's CPU model is preferred. When default CPU model is not set and VMI's CPU model is not set too, host-model will be set. Default CPU model can be changed when kubevirt is running.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"obsoleteCPUs": {
						SchemaProps: spec.SchemaProps{
							Description: "ObsoleteCPUs allows avoiding scheduling of VMs for obsolete CPU models",
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedObsoleteCPUs"),
						},
					},
					"commonTemplatesNamespace": {
						SchemaProps: spec.SchemaProps{
							Description: "CommonTemplatesNamespace defines namespace in which common templates will be deployed. It overrides the default openshift namespace.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"storageImport": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageImport contains configuration for importing containerized data",
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.StorageImportConfig"),
						},
					},
					"workloadUpdateStrategy": {
						SchemaProps: spec.SchemaProps{
							Description: "WorkloadUpdateStrategy defines at the cluster level how to handle automated workload updates",
							Ref:         ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedWorkloadUpdateStrategy"),
						},
					},
					"dataImportCronTemplates": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "DataImportCronTemplates holds list of data import cron templates (golden images)",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("kubevirt.io/ssp-operator/api/v1beta1.DataImportCronTemplate"),
									},
								},
							},
						},
					},
					"uninstallStrategy": {
						SchemaProps: spec.SchemaProps{
							Description: "UninstallStrategy defines how to proceed on uninstall when workloads (VirtualMachines, DataVolumes) still exist. BlockUninstallIfWorkloadsExist will prevent the CR from being removed when workloads still exist. BlockUninstallIfWorkloadsExist is the safest choice to protect your workloads from accidental data loss, so it's strongly advised. RemoveWorkloads will cause all the workloads to be cascading deleted on uninstall. WARNING: please notice that RemoveWorkloads will cause your workloads to be deleted as soon as this CR will be, even accidentally, deleted. Please correctly consider the implications of this option before setting it. BlockUninstallIfWorkloadsExist is the default behaviour.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedCertConfig", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedConfig", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedFeatureGates", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedObsoleteCPUs", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.HyperConvergedWorkloadUpdateStrategy", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.LiveMigrationConfigurations", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.MediatedDevicesConfiguration", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.OperandResourceRequirements", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.PermittedHostDevices", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.StorageImportConfig", "kubevirt.io/ssp-operator/api/v1beta1.DataImportCronTemplate"},
	}
}

func schema_pkg_apis_hco_v1beta1_HyperConvergedStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HyperConvergedStatus defines the observed state of HyperConverged",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions describes the state of the HyperConverged resource.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.Condition"),
									},
								},
							},
						},
					},
					"relatedObjects": {
						SchemaProps: spec.SchemaProps{
							Description: "RelatedObjects is a list of objects created and maintained by this operator. Object references will be added to this list after they have been created AND found in the cluster.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
					"versions": {
						SchemaProps: spec.SchemaProps{
							Description: "Versions is a list of HCO component versions, as name/version pairs. The version with a name of \"operator\" is the HCO version itself, as described here: https://github.com/openshift/cluster-version-operator/blob/master/docs/dev/clusteroperator.md#version",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.Version"),
									},
								},
							},
						},
					},
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "ObservedGeneration reflects the HyperConverged resource generation. If the ObservedGeneration is less than the resource generation in metadata, the status is out of date",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"dataImportSchedule": {
						SchemaProps: spec.SchemaProps{
							Description: "DataImportSchedule is the cron expression that is used in for the hard-coded data import cron templates. HCO generates the value of this field once and stored in the status field, so will survive restart.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.Version", "k8s.io/api/core/v1.ObjectReference", "k8s.io/apimachinery/pkg/apis/meta/v1.Condition"},
	}
}

func schema_pkg_apis_hco_v1beta1_HyperConvergedWorkloadUpdateStrategy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "HyperConvergedWorkloadUpdateStrategy defines options related to updating a KubeVirt install",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"workloadUpdateMethods": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "WorkloadUpdateMethods defines the methods that can be used to disrupt workloads during automated workload updates. When multiple methods are present, the least disruptive method takes precedence over more disruptive methods. For example if both LiveMigrate and Evict methods are listed, only VMs which are not live migratable will be restarted/shutdown. An empty list defaults to no automated workload updating.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"batchEvictionSize": {
						SchemaProps: spec.SchemaProps{
							Description: "BatchEvictionSize Represents the number of VMIs that can be forced updated per the BatchShutdownInteral interval",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"batchEvictionInterval": {
						SchemaProps: spec.SchemaProps{
							Description: "BatchEvictionInterval Represents the interval to wait before issuing the next batch of shutdowns",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Duration"},
	}
}

func schema_pkg_apis_hco_v1beta1_LiveMigrationConfigurations(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LiveMigrationConfigurations - Live migration limits and timeouts are applied so that migration processes do not overwhelm the cluster.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"parallelMigrationsPerCluster": {
						SchemaProps: spec.SchemaProps{
							Description: "Number of migrations running in parallel in the cluster.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"parallelOutboundMigrationsPerNode": {
						SchemaProps: spec.SchemaProps{
							Description: "Maximum number of outbound migrations per node.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"bandwidthPerMigration": {
						SchemaProps: spec.SchemaProps{
							Description: "Bandwidth limit of each migration, in MiB/s.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"completionTimeoutPerGiB": {
						SchemaProps: spec.SchemaProps{
							Description: "The migration will be canceled if it has not completed in this time, in seconds per GiB of memory. For example, a virtual machine instance with 6GiB memory will timeout if it has not completed migration in 4800 seconds. If the Migration Method is BlockMigration, the size of the migrating disks is included in the calculation.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"progressTimeout": {
						SchemaProps: spec.SchemaProps{
							Description: "The migration will be canceled if memory copy fails to make progress in this time, in seconds.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"network": {
						SchemaProps: spec.SchemaProps{
							Description: "The migrations will be performed over a dedicated multus network to minimize disruption to tenant workloads due to network saturation when VM live migrations are triggered.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_hco_v1beta1_MediatedDevicesConfiguration(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MediatedDevicesConfiguration holds inforamtion about MDEV types to be defined, if available",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"mediatedDevicesTypes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"nodeMediatedDeviceTypes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.NodeMediatedDeviceTypesConfig"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.NodeMediatedDeviceTypesConfig"},
	}
}

func schema_pkg_apis_hco_v1beta1_MediatedHostDevice(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MediatedHostDevice represents a host mediated device allowed for passthrough",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"mdevNameSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "name of a mediated device type required to identify a mediated device on a host",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceName": {
						SchemaProps: spec.SchemaProps{
							Description: "name by which a device is advertised and being requested",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"externalResourceProvider": {
						SchemaProps: spec.SchemaProps{
							Description: "indicates that this resource is being provided by an external device plugin",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disabled": {
						SchemaProps: spec.SchemaProps{
							Description: "HCO enforces the existence of several MediatedHostDevice objects. Set disabled field to true instead of remove these objects.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"mdevNameSelector", "resourceName"},
			},
		},
	}
}

func schema_pkg_apis_hco_v1beta1_NodeMediatedDeviceTypesConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NodeMediatedDeviceTypesConfig holds inforamtion about MDEV types to be defined in a specifc node that matches the NodeSelector field.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelector is a selector which must be true for the vmi to fit on a node. Selector which must match a node's labels for the vmi to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"mediatedDevicesTypes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
				Required: []string{"nodeSelector", "mediatedDevicesTypes"},
			},
		},
	}
}

func schema_pkg_apis_hco_v1beta1_OperandResourceRequirements(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OperandResourceRequirements is a list of resource requirements for the operand workloads pods",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"storageWorkloads": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageWorkloads defines the resources requirements for storage workloads. It will propagate to the CDI custom resource",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_hco_v1beta1_PciHostDevice(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PciHostDevice represents a host PCI device allowed for passthrough",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"pciDeviceSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "a combination of a vendor_id:product_id required to identify a PCI device on a host.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceName": {
						SchemaProps: spec.SchemaProps{
							Description: "name by which a device is advertised and being requested",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"externalResourceProvider": {
						SchemaProps: spec.SchemaProps{
							Description: "indicates that this resource is being provided by an external device plugin",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disabled": {
						SchemaProps: spec.SchemaProps{
							Description: "HCO enforces the existence of several PciHostDevice objects. Set disabled field to true instead of remove these objects.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"pciDeviceSelector", "resourceName"},
			},
		},
	}
}

func schema_pkg_apis_hco_v1beta1_PermittedHostDevices(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PermittedHostDevices holds information about devices allowed for passthrough",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"pciHostDevices": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"pciDeviceSelector",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.PciHostDevice"),
									},
								},
							},
						},
					},
					"mediatedDevices": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"mdevNameSelector",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.MediatedHostDevice"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.MediatedHostDevice", "github.com/kubevirt/hyperconverged-cluster-operator/pkg/apis/hco/v1beta1.PciHostDevice"},
	}
}

func schema_pkg_apis_hco_v1beta1_StorageImportConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageImportConfig contains configuration for importing containerized data",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"insecureRegistries": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "InsecureRegistries is a list of image registries URLs that are not secured. Setting an insecure registry URL in this list allows pulling images from this registry.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
