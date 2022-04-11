// Automatically generated by MockGen. DO NOT EDIT!
// Source: cgroup.go

package cgroup

import (
	gomock "github.com/golang/mock/gomock"
	configs "github.com/opencontainers/runc/libcontainer/configs"
)

// Mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *_MockManagerRecorder
}

// Recorder for MockManager (not exported)
type _MockManagerRecorder struct {
	mock *MockManager
}

func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &_MockManagerRecorder{mock}
	return mock
}

func (_m *MockManager) EXPECT() *_MockManagerRecorder {
	return _m.recorder
}

func (_m *MockManager) Set(r *configs.Resources) error {
	ret := _m.ctrl.Call(_m, "Set", r)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockManagerRecorder) Set(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Set", arg0)
}

func (_m *MockManager) GetBasePathToHostSubsystem(subsystem string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetBasePathToHostSubsystem", subsystem)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) GetBasePathToHostSubsystem(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetBasePathToHostSubsystem", arg0)
}

func (_m *MockManager) GetCgroupVersion() CgroupVersion {
	ret := _m.ctrl.Call(_m, "GetCgroupVersion")
	ret0, _ := ret[0].(CgroupVersion)
	return ret0
}

func (_mr *_MockManagerRecorder) GetCgroupVersion() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCgroupVersion")
}

func (_m *MockManager) GetCpuSet() (string, error) {
	ret := _m.ctrl.Call(_m, "GetCpuSet")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockManagerRecorder) GetCpuSet() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCpuSet")
}