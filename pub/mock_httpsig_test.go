// Code generated by MockGen. DO NOT EDIT.
// Source: ../../httpsig/httpsig.go

// Package pub is a generated GoMock package.
package pub

import (
	crypto "crypto"
	httpsig "github.com/go-fed/httpsig"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockSigner is a mock of Signer interface
type MockSigner struct {
	ctrl     *gomock.Controller
	recorder *MockSignerMockRecorder
}

// MockSignerMockRecorder is the mock recorder for MockSigner
type MockSignerMockRecorder struct {
	mock *MockSigner
}

// NewMockSigner creates a new mock instance
func NewMockSigner(ctrl *gomock.Controller) *MockSigner {
	mock := &MockSigner{ctrl: ctrl}
	mock.recorder = &MockSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSigner) EXPECT() *MockSignerMockRecorder {
	return m.recorder
}

// SignRequest mocks base method
func (m *MockSigner) SignRequest(pKey crypto.PrivateKey, pubKeyId string, r *http.Request, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignRequest", pKey, pubKeyId, r, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignRequest indicates an expected call of SignRequest
func (mr *MockSignerMockRecorder) SignRequest(pKey, pubKeyId, r, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignRequest", reflect.TypeOf((*MockSigner)(nil).SignRequest), pKey, pubKeyId, r, body)
}

// SignResponse mocks base method
func (m *MockSigner) SignResponse(pKey crypto.PrivateKey, pubKeyId string, r http.ResponseWriter, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignResponse", pKey, pubKeyId, r, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignResponse indicates an expected call of SignResponse
func (mr *MockSignerMockRecorder) SignResponse(pKey, pubKeyId, r, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignResponse", reflect.TypeOf((*MockSigner)(nil).SignResponse), pKey, pubKeyId, r, body)
}

// MockSSHSigner is a mock of SSHSigner interface
type MockSSHSigner struct {
	ctrl     *gomock.Controller
	recorder *MockSSHSignerMockRecorder
}

// MockSSHSignerMockRecorder is the mock recorder for MockSSHSigner
type MockSSHSignerMockRecorder struct {
	mock *MockSSHSigner
}

// NewMockSSHSigner creates a new mock instance
func NewMockSSHSigner(ctrl *gomock.Controller) *MockSSHSigner {
	mock := &MockSSHSigner{ctrl: ctrl}
	mock.recorder = &MockSSHSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSSHSigner) EXPECT() *MockSSHSignerMockRecorder {
	return m.recorder
}

// SignRequest mocks base method
func (m *MockSSHSigner) SignRequest(pubKeyId string, r *http.Request, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignRequest", pubKeyId, r, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignRequest indicates an expected call of SignRequest
func (mr *MockSSHSignerMockRecorder) SignRequest(pubKeyId, r, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignRequest", reflect.TypeOf((*MockSSHSigner)(nil).SignRequest), pubKeyId, r, body)
}

// SignResponse mocks base method
func (m *MockSSHSigner) SignResponse(pubKeyId string, r http.ResponseWriter, body []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignResponse", pubKeyId, r, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignResponse indicates an expected call of SignResponse
func (mr *MockSSHSignerMockRecorder) SignResponse(pubKeyId, r, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignResponse", reflect.TypeOf((*MockSSHSigner)(nil).SignResponse), pubKeyId, r, body)
}

// MockVerifier is a mock of Verifier interface
type MockVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierMockRecorder
}

// MockVerifierMockRecorder is the mock recorder for MockVerifier
type MockVerifierMockRecorder struct {
	mock *MockVerifier
}

// NewMockVerifier creates a new mock instance
func NewMockVerifier(ctrl *gomock.Controller) *MockVerifier {
	mock := &MockVerifier{ctrl: ctrl}
	mock.recorder = &MockVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVerifier) EXPECT() *MockVerifierMockRecorder {
	return m.recorder
}

// KeyId mocks base method
func (m *MockVerifier) KeyId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyId")
	ret0, _ := ret[0].(string)
	return ret0
}

// KeyId indicates an expected call of KeyId
func (mr *MockVerifierMockRecorder) KeyId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyId", reflect.TypeOf((*MockVerifier)(nil).KeyId))
}

// Verify mocks base method
func (m *MockVerifier) Verify(pKey crypto.PublicKey, algo httpsig.Algorithm) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", pKey, algo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify
func (mr *MockVerifierMockRecorder) Verify(pKey, algo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVerifier)(nil).Verify), pKey, algo)
}
