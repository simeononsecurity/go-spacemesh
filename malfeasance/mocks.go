// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface.go
//
// Generated by this command:
//
//	mockgen -typed -package=malfeasance -destination=./mocks.go -source=./interface.go
//
// Package malfeasance is a generated GoMock package.
package malfeasance

import (
	reflect "reflect"

	types "github.com/spacemeshos/go-spacemesh/common/types"
	signing "github.com/spacemeshos/go-spacemesh/signing"
	gomock "go.uber.org/mock/gomock"
)

// MockSigVerifier is a mock of SigVerifier interface.
type MockSigVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockSigVerifierMockRecorder
}

// MockSigVerifierMockRecorder is the mock recorder for MockSigVerifier.
type MockSigVerifierMockRecorder struct {
	mock *MockSigVerifier
}

// NewMockSigVerifier creates a new mock instance.
func NewMockSigVerifier(ctrl *gomock.Controller) *MockSigVerifier {
	mock := &MockSigVerifier{ctrl: ctrl}
	mock.recorder = &MockSigVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSigVerifier) EXPECT() *MockSigVerifierMockRecorder {
	return m.recorder
}

// Verify mocks base method.
func (m *MockSigVerifier) Verify(arg0 signing.Domain, arg1 types.NodeID, arg2 []byte, arg3 types.EdSignature) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockSigVerifierMockRecorder) Verify(arg0, arg1, arg2, arg3 any) *SigVerifierVerifyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockSigVerifier)(nil).Verify), arg0, arg1, arg2, arg3)
	return &SigVerifierVerifyCall{Call: call}
}

// SigVerifierVerifyCall wrap *gomock.Call
type SigVerifierVerifyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SigVerifierVerifyCall) Return(arg0 bool) *SigVerifierVerifyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SigVerifierVerifyCall) Do(f func(signing.Domain, types.NodeID, []byte, types.EdSignature) bool) *SigVerifierVerifyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SigVerifierVerifyCall) DoAndReturn(f func(signing.Domain, types.NodeID, []byte, types.EdSignature) bool) *SigVerifierVerifyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Mocktortoise is a mock of tortoise interface.
type Mocktortoise struct {
	ctrl     *gomock.Controller
	recorder *MocktortoiseMockRecorder
}

// MocktortoiseMockRecorder is the mock recorder for Mocktortoise.
type MocktortoiseMockRecorder struct {
	mock *Mocktortoise
}

// NewMocktortoise creates a new mock instance.
func NewMocktortoise(ctrl *gomock.Controller) *Mocktortoise {
	mock := &Mocktortoise{ctrl: ctrl}
	mock.recorder = &MocktortoiseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocktortoise) EXPECT() *MocktortoiseMockRecorder {
	return m.recorder
}

// OnMalfeasance mocks base method.
func (m *Mocktortoise) OnMalfeasance(arg0 types.NodeID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnMalfeasance", arg0)
}

// OnMalfeasance indicates an expected call of OnMalfeasance.
func (mr *MocktortoiseMockRecorder) OnMalfeasance(arg0 any) *tortoiseOnMalfeasanceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnMalfeasance", reflect.TypeOf((*Mocktortoise)(nil).OnMalfeasance), arg0)
	return &tortoiseOnMalfeasanceCall{Call: call}
}

// tortoiseOnMalfeasanceCall wrap *gomock.Call
type tortoiseOnMalfeasanceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *tortoiseOnMalfeasanceCall) Return() *tortoiseOnMalfeasanceCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *tortoiseOnMalfeasanceCall) Do(f func(types.NodeID)) *tortoiseOnMalfeasanceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *tortoiseOnMalfeasanceCall) DoAndReturn(f func(types.NodeID)) *tortoiseOnMalfeasanceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
