// Automatically generated by MockGen. DO NOT EDIT!
// Source: manager.go

package virtwrap

import (
	gomock "github.com/golang/mock/gomock"
	v1 "kubevirt.io/api/core/v1"

	v10 "kubevirt.io/kubevirt/pkg/handler-launcher-com/cmd/v1"
	cmd_client "kubevirt.io/kubevirt/pkg/virt-handler/cmd-client"
	api "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
	stats "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/stats"
)

// Mock of DomainManager interface
type MockDomainManager struct {
	ctrl     *gomock.Controller
	recorder *_MockDomainManagerRecorder
}

// Recorder for MockDomainManager (not exported)
type _MockDomainManagerRecorder struct {
	mock *MockDomainManager
}

func NewMockDomainManager(ctrl *gomock.Controller) *MockDomainManager {
	mock := &MockDomainManager{ctrl: ctrl}
	mock.recorder = &_MockDomainManagerRecorder{mock}
	return mock
}

func (_m *MockDomainManager) EXPECT() *_MockDomainManagerRecorder {
	return _m.recorder
}

func (_m *MockDomainManager) SyncVMI(_param0 *v1.VirtualMachineInstance, _param1 bool, _param2 *v10.VirtualMachineOptions) (*api.DomainSpec, error) {
	ret := _m.ctrl.Call(_m, "SyncVMI", _param0, _param1, _param2)
	ret0, _ := ret[0].(*api.DomainSpec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDomainManagerRecorder) SyncVMI(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SyncVMI", arg0, arg1, arg2)
}

func (_m *MockDomainManager) PauseVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "PauseVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) PauseVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PauseVMI", arg0)
}

func (_m *MockDomainManager) UnpauseVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "UnpauseVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) UnpauseVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnpauseVMI", arg0)
}

func (_m *MockDomainManager) FreezeVMI(_param0 *v1.VirtualMachineInstance, _param1 int32) error {
	ret := _m.ctrl.Call(_m, "FreezeVMI", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) FreezeVMI(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FreezeVMI", arg0, arg1)
}

func (_m *MockDomainManager) UnfreezeVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "UnfreezeVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) UnfreezeVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnfreezeVMI", arg0)
}

func (_m *MockDomainManager) SoftRebootVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "SoftRebootVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) SoftRebootVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SoftRebootVMI", arg0)
}

func (_m *MockDomainManager) KillVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "KillVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) KillVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KillVMI", arg0)
}

func (_m *MockDomainManager) DeleteVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "DeleteVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) DeleteVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteVMI", arg0)
}

func (_m *MockDomainManager) SignalShutdownVMI(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "SignalShutdownVMI", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) SignalShutdownVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SignalShutdownVMI", arg0)
}

func (_m *MockDomainManager) MarkGracefulShutdownVMI() {
	_m.ctrl.Call(_m, "MarkGracefulShutdownVMI")
}

func (_mr *_MockDomainManagerRecorder) MarkGracefulShutdownVMI() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MarkGracefulShutdownVMI")
}

func (_m *MockDomainManager) ListAllDomains() ([]*api.Domain, error) {
	ret := _m.ctrl.Call(_m, "ListAllDomains")
	ret0, _ := ret[0].([]*api.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDomainManagerRecorder) ListAllDomains() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAllDomains")
}

func (_m *MockDomainManager) MigrateVMI(_param0 *v1.VirtualMachineInstance, _param1 *cmd_client.MigrationOptions) error {
	ret := _m.ctrl.Call(_m, "MigrateVMI", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) MigrateVMI(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MigrateVMI", arg0, arg1)
}

func (_m *MockDomainManager) PrepareMigrationTarget(_param0 *v1.VirtualMachineInstance, _param1 bool, _param2 *v10.VirtualMachineOptions) error {
	ret := _m.ctrl.Call(_m, "PrepareMigrationTarget", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) PrepareMigrationTarget(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PrepareMigrationTarget", arg0, arg1, arg2)
}

func (_m *MockDomainManager) GetDomainStats() ([]*stats.DomainStats, error) {
	ret := _m.ctrl.Call(_m, "GetDomainStats")
	ret0, _ := ret[0].([]*stats.DomainStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDomainManagerRecorder) GetDomainStats() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDomainStats")
}

func (_m *MockDomainManager) CancelVMIMigration(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "CancelVMIMigration", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) CancelVMIMigration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CancelVMIMigration", arg0)
}

func (_m *MockDomainManager) GetGuestInfo() v1.VirtualMachineInstanceGuestAgentInfo {
	ret := _m.ctrl.Call(_m, "GetGuestInfo")
	ret0, _ := ret[0].(v1.VirtualMachineInstanceGuestAgentInfo)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) GetGuestInfo() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGuestInfo")
}

func (_m *MockDomainManager) GetUsers() []v1.VirtualMachineInstanceGuestOSUser {
	ret := _m.ctrl.Call(_m, "GetUsers")
	ret0, _ := ret[0].([]v1.VirtualMachineInstanceGuestOSUser)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) GetUsers() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUsers")
}

func (_m *MockDomainManager) GetFilesystems() []v1.VirtualMachineInstanceFileSystem {
	ret := _m.ctrl.Call(_m, "GetFilesystems")
	ret0, _ := ret[0].([]v1.VirtualMachineInstanceFileSystem)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) GetFilesystems() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFilesystems")
}

func (_m *MockDomainManager) FinalizeVirtualMachineMigration(_param0 *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "FinalizeVirtualMachineMigration", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) FinalizeVirtualMachineMigration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FinalizeVirtualMachineMigration", arg0)
}

func (_m *MockDomainManager) HotplugHostDevices(vmi *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "HotplugHostDevices", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) HotplugHostDevices(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HotplugHostDevices", arg0)
}

func (_m *MockDomainManager) InterfacesStatus() []api.InterfaceStatus {
	ret := _m.ctrl.Call(_m, "InterfacesStatus")
	ret0, _ := ret[0].([]api.InterfaceStatus)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) InterfacesStatus() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InterfacesStatus")
}

func (_m *MockDomainManager) GetGuestOSInfo() *api.GuestOSInfo {
	ret := _m.ctrl.Call(_m, "GetGuestOSInfo")
	ret0, _ := ret[0].(*api.GuestOSInfo)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) GetGuestOSInfo() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetGuestOSInfo")
}

func (_m *MockDomainManager) Exec(_param0 string, _param1 string, _param2 []string, _param3 int32) (string, error) {
	ret := _m.ctrl.Call(_m, "Exec", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDomainManagerRecorder) Exec(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exec", arg0, arg1, arg2, arg3)
}

func (_m *MockDomainManager) GuestPing(_param0 string) error {
	ret := _m.ctrl.Call(_m, "GuestPing", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) GuestPing(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GuestPing", arg0)
}

func (_m *MockDomainManager) MemoryDump(vmi *v1.VirtualMachineInstance, dumpPath string) error {
	ret := _m.ctrl.Call(_m, "MemoryDump", vmi, dumpPath)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) MemoryDump(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MemoryDump", arg0, arg1)
}

func (_m *MockDomainManager) CreateSnapshotVMI(vmi *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "CreateSnapshotVMI", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDomainManagerRecorder) CreateSnapshotVMI(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateSnapshotVMI", arg0)
}
