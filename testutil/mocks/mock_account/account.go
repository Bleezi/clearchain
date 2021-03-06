// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/tendermint/clearchain/types (interfaces: AccountGetter,AccountSetter,AccountGetterSetter)

package mock_account

import (
	gomock "github.com/golang/mock/gomock"
	types "github.com/tendermint/clearchain/types"
)

// Mock of AccountGetter interface
type MockAccountGetter struct {
	ctrl     *gomock.Controller
	recorder *_MockAccountGetterRecorder
}

// Recorder for MockAccountGetter (not exported)
type _MockAccountGetterRecorder struct {
	mock *MockAccountGetter
}

func NewMockAccountGetter(ctrl *gomock.Controller) *MockAccountGetter {
	mock := &MockAccountGetter{ctrl: ctrl}
	mock.recorder = &_MockAccountGetterRecorder{mock}
	return mock
}

func (_m *MockAccountGetter) EXPECT() *_MockAccountGetterRecorder {
	return _m.recorder
}

func (_m *MockAccountGetter) GetAccount(_param0 string) *types.Account {
	ret := _m.ctrl.Call(_m, "GetAccount", _param0)
	ret0, _ := ret[0].(*types.Account)
	return ret0
}

func (_mr *_MockAccountGetterRecorder) GetAccount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccount", arg0)
}

// Mock of AccountSetter interface
type MockAccountSetter struct {
	ctrl     *gomock.Controller
	recorder *_MockAccountSetterRecorder
}

// Recorder for MockAccountSetter (not exported)
type _MockAccountSetterRecorder struct {
	mock *MockAccountSetter
}

func NewMockAccountSetter(ctrl *gomock.Controller) *MockAccountSetter {
	mock := &MockAccountSetter{ctrl: ctrl}
	mock.recorder = &_MockAccountSetterRecorder{mock}
	return mock
}

func (_m *MockAccountSetter) EXPECT() *_MockAccountSetterRecorder {
	return _m.recorder
}

func (_m *MockAccountSetter) SetAccount(_param0 string, _param1 *types.Account) {
	_m.ctrl.Call(_m, "SetAccount", _param0, _param1)
}

func (_mr *_MockAccountSetterRecorder) SetAccount(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetAccount", arg0, arg1)
}

// Mock of AccountGetterSetter interface
type MockAccountGetterSetter struct {
	ctrl     *gomock.Controller
	recorder *_MockAccountGetterSetterRecorder
}

// Recorder for MockAccountGetterSetter (not exported)
type _MockAccountGetterSetterRecorder struct {
	mock *MockAccountGetterSetter
}

func NewMockAccountGetterSetter(ctrl *gomock.Controller) *MockAccountGetterSetter {
	mock := &MockAccountGetterSetter{ctrl: ctrl}
	mock.recorder = &_MockAccountGetterSetterRecorder{mock}
	return mock
}

func (_m *MockAccountGetterSetter) EXPECT() *_MockAccountGetterSetterRecorder {
	return _m.recorder
}

func (_m *MockAccountGetterSetter) GetAccount(_param0 string) *types.Account {
	ret := _m.ctrl.Call(_m, "GetAccount", _param0)
	ret0, _ := ret[0].(*types.Account)
	return ret0
}

func (_mr *_MockAccountGetterSetterRecorder) GetAccount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccount", arg0)
}

func (_m *MockAccountGetterSetter) SetAccount(_param0 string, _param1 *types.Account) {
	_m.ctrl.Call(_m, "SetAccount", _param0, _param1)
}

func (_mr *_MockAccountGetterSetterRecorder) SetAccount(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetAccount", arg0, arg1)
}
