// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/gsagula/go/pkg/mod/github.com/aws/aws-sdk-go@v1.29.31/service/acm/acmiface/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	acm "github.com/aws/aws-sdk-go/service/acm"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockACMAPI is a mock of ACMAPI interface
type MockACMAPI struct {
	ctrl     *gomock.Controller
	recorder *MockACMAPIMockRecorder
}

// MockACMAPIMockRecorder is the mock recorder for MockACMAPI
type MockACMAPIMockRecorder struct {
	mock *MockACMAPI
}

// NewMockACMAPI creates a new mock instance
func NewMockACMAPI(ctrl *gomock.Controller) *MockACMAPI {
	mock := &MockACMAPI{ctrl: ctrl}
	mock.recorder = &MockACMAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockACMAPI) EXPECT() *MockACMAPIMockRecorder {
	return m.recorder
}

// AddTagsToCertificate mocks base method
func (m *MockACMAPI) AddTagsToCertificate(arg0 *acm.AddTagsToCertificateInput) (*acm.AddTagsToCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTagsToCertificate", arg0)
	ret0, _ := ret[0].(*acm.AddTagsToCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTagsToCertificate indicates an expected call of AddTagsToCertificate
func (mr *MockACMAPIMockRecorder) AddTagsToCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTagsToCertificate", reflect.TypeOf((*MockACMAPI)(nil).AddTagsToCertificate), arg0)
}

// AddTagsToCertificateWithContext mocks base method
func (m *MockACMAPI) AddTagsToCertificateWithContext(arg0 aws.Context, arg1 *acm.AddTagsToCertificateInput, arg2 ...request.Option) (*acm.AddTagsToCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddTagsToCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.AddTagsToCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTagsToCertificateWithContext indicates an expected call of AddTagsToCertificateWithContext
func (mr *MockACMAPIMockRecorder) AddTagsToCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTagsToCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).AddTagsToCertificateWithContext), varargs...)
}

// AddTagsToCertificateRequest mocks base method
func (m *MockACMAPI) AddTagsToCertificateRequest(arg0 *acm.AddTagsToCertificateInput) (*request.Request, *acm.AddTagsToCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTagsToCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.AddTagsToCertificateOutput)
	return ret0, ret1
}

// AddTagsToCertificateRequest indicates an expected call of AddTagsToCertificateRequest
func (mr *MockACMAPIMockRecorder) AddTagsToCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTagsToCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).AddTagsToCertificateRequest), arg0)
}

// DeleteCertificate mocks base method
func (m *MockACMAPI) DeleteCertificate(arg0 *acm.DeleteCertificateInput) (*acm.DeleteCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCertificate", arg0)
	ret0, _ := ret[0].(*acm.DeleteCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCertificate indicates an expected call of DeleteCertificate
func (mr *MockACMAPIMockRecorder) DeleteCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCertificate", reflect.TypeOf((*MockACMAPI)(nil).DeleteCertificate), arg0)
}

// DeleteCertificateWithContext mocks base method
func (m *MockACMAPI) DeleteCertificateWithContext(arg0 aws.Context, arg1 *acm.DeleteCertificateInput, arg2 ...request.Option) (*acm.DeleteCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.DeleteCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCertificateWithContext indicates an expected call of DeleteCertificateWithContext
func (mr *MockACMAPIMockRecorder) DeleteCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).DeleteCertificateWithContext), varargs...)
}

// DeleteCertificateRequest mocks base method
func (m *MockACMAPI) DeleteCertificateRequest(arg0 *acm.DeleteCertificateInput) (*request.Request, *acm.DeleteCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.DeleteCertificateOutput)
	return ret0, ret1
}

// DeleteCertificateRequest indicates an expected call of DeleteCertificateRequest
func (mr *MockACMAPIMockRecorder) DeleteCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).DeleteCertificateRequest), arg0)
}

// DescribeCertificate mocks base method
func (m *MockACMAPI) DescribeCertificate(arg0 *acm.DescribeCertificateInput) (*acm.DescribeCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCertificate", arg0)
	ret0, _ := ret[0].(*acm.DescribeCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCertificate indicates an expected call of DescribeCertificate
func (mr *MockACMAPIMockRecorder) DescribeCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCertificate", reflect.TypeOf((*MockACMAPI)(nil).DescribeCertificate), arg0)
}

// DescribeCertificateWithContext mocks base method
func (m *MockACMAPI) DescribeCertificateWithContext(arg0 aws.Context, arg1 *acm.DescribeCertificateInput, arg2 ...request.Option) (*acm.DescribeCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.DescribeCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCertificateWithContext indicates an expected call of DescribeCertificateWithContext
func (mr *MockACMAPIMockRecorder) DescribeCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).DescribeCertificateWithContext), varargs...)
}

// DescribeCertificateRequest mocks base method
func (m *MockACMAPI) DescribeCertificateRequest(arg0 *acm.DescribeCertificateInput) (*request.Request, *acm.DescribeCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.DescribeCertificateOutput)
	return ret0, ret1
}

// DescribeCertificateRequest indicates an expected call of DescribeCertificateRequest
func (mr *MockACMAPIMockRecorder) DescribeCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).DescribeCertificateRequest), arg0)
}

// ExportCertificate mocks base method
func (m *MockACMAPI) ExportCertificate(arg0 *acm.ExportCertificateInput) (*acm.ExportCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportCertificate", arg0)
	ret0, _ := ret[0].(*acm.ExportCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportCertificate indicates an expected call of ExportCertificate
func (mr *MockACMAPIMockRecorder) ExportCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportCertificate", reflect.TypeOf((*MockACMAPI)(nil).ExportCertificate), arg0)
}

// ExportCertificateWithContext mocks base method
func (m *MockACMAPI) ExportCertificateWithContext(arg0 aws.Context, arg1 *acm.ExportCertificateInput, arg2 ...request.Option) (*acm.ExportCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExportCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.ExportCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExportCertificateWithContext indicates an expected call of ExportCertificateWithContext
func (mr *MockACMAPIMockRecorder) ExportCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).ExportCertificateWithContext), varargs...)
}

// ExportCertificateRequest mocks base method
func (m *MockACMAPI) ExportCertificateRequest(arg0 *acm.ExportCertificateInput) (*request.Request, *acm.ExportCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.ExportCertificateOutput)
	return ret0, ret1
}

// ExportCertificateRequest indicates an expected call of ExportCertificateRequest
func (mr *MockACMAPIMockRecorder) ExportCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).ExportCertificateRequest), arg0)
}

// GetCertificate mocks base method
func (m *MockACMAPI) GetCertificate(arg0 *acm.GetCertificateInput) (*acm.GetCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificate", arg0)
	ret0, _ := ret[0].(*acm.GetCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificate indicates an expected call of GetCertificate
func (mr *MockACMAPIMockRecorder) GetCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificate", reflect.TypeOf((*MockACMAPI)(nil).GetCertificate), arg0)
}

// GetCertificateWithContext mocks base method
func (m *MockACMAPI) GetCertificateWithContext(arg0 aws.Context, arg1 *acm.GetCertificateInput, arg2 ...request.Option) (*acm.GetCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.GetCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificateWithContext indicates an expected call of GetCertificateWithContext
func (mr *MockACMAPIMockRecorder) GetCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).GetCertificateWithContext), varargs...)
}

// GetCertificateRequest mocks base method
func (m *MockACMAPI) GetCertificateRequest(arg0 *acm.GetCertificateInput) (*request.Request, *acm.GetCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.GetCertificateOutput)
	return ret0, ret1
}

// GetCertificateRequest indicates an expected call of GetCertificateRequest
func (mr *MockACMAPIMockRecorder) GetCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).GetCertificateRequest), arg0)
}

// ImportCertificate mocks base method
func (m *MockACMAPI) ImportCertificate(arg0 *acm.ImportCertificateInput) (*acm.ImportCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportCertificate", arg0)
	ret0, _ := ret[0].(*acm.ImportCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportCertificate indicates an expected call of ImportCertificate
func (mr *MockACMAPIMockRecorder) ImportCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportCertificate", reflect.TypeOf((*MockACMAPI)(nil).ImportCertificate), arg0)
}

// ImportCertificateWithContext mocks base method
func (m *MockACMAPI) ImportCertificateWithContext(arg0 aws.Context, arg1 *acm.ImportCertificateInput, arg2 ...request.Option) (*acm.ImportCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ImportCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.ImportCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportCertificateWithContext indicates an expected call of ImportCertificateWithContext
func (mr *MockACMAPIMockRecorder) ImportCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).ImportCertificateWithContext), varargs...)
}

// ImportCertificateRequest mocks base method
func (m *MockACMAPI) ImportCertificateRequest(arg0 *acm.ImportCertificateInput) (*request.Request, *acm.ImportCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.ImportCertificateOutput)
	return ret0, ret1
}

// ImportCertificateRequest indicates an expected call of ImportCertificateRequest
func (mr *MockACMAPIMockRecorder) ImportCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).ImportCertificateRequest), arg0)
}

// ListCertificates mocks base method
func (m *MockACMAPI) ListCertificates(arg0 *acm.ListCertificatesInput) (*acm.ListCertificatesOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCertificates", arg0)
	ret0, _ := ret[0].(*acm.ListCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCertificates indicates an expected call of ListCertificates
func (mr *MockACMAPIMockRecorder) ListCertificates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificates", reflect.TypeOf((*MockACMAPI)(nil).ListCertificates), arg0)
}

// ListCertificatesWithContext mocks base method
func (m *MockACMAPI) ListCertificatesWithContext(arg0 aws.Context, arg1 *acm.ListCertificatesInput, arg2 ...request.Option) (*acm.ListCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCertificatesWithContext", varargs...)
	ret0, _ := ret[0].(*acm.ListCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCertificatesWithContext indicates an expected call of ListCertificatesWithContext
func (mr *MockACMAPIMockRecorder) ListCertificatesWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificatesWithContext", reflect.TypeOf((*MockACMAPI)(nil).ListCertificatesWithContext), varargs...)
}

// ListCertificatesRequest mocks base method
func (m *MockACMAPI) ListCertificatesRequest(arg0 *acm.ListCertificatesInput) (*request.Request, *acm.ListCertificatesOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCertificatesRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.ListCertificatesOutput)
	return ret0, ret1
}

// ListCertificatesRequest indicates an expected call of ListCertificatesRequest
func (mr *MockACMAPIMockRecorder) ListCertificatesRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificatesRequest", reflect.TypeOf((*MockACMAPI)(nil).ListCertificatesRequest), arg0)
}

// ListCertificatesPages mocks base method
func (m *MockACMAPI) ListCertificatesPages(arg0 *acm.ListCertificatesInput, arg1 func(*acm.ListCertificatesOutput, bool) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCertificatesPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListCertificatesPages indicates an expected call of ListCertificatesPages
func (mr *MockACMAPIMockRecorder) ListCertificatesPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificatesPages", reflect.TypeOf((*MockACMAPI)(nil).ListCertificatesPages), arg0, arg1)
}

// ListCertificatesPagesWithContext mocks base method
func (m *MockACMAPI) ListCertificatesPagesWithContext(arg0 aws.Context, arg1 *acm.ListCertificatesInput, arg2 func(*acm.ListCertificatesOutput, bool) bool, arg3 ...request.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCertificatesPagesWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListCertificatesPagesWithContext indicates an expected call of ListCertificatesPagesWithContext
func (mr *MockACMAPIMockRecorder) ListCertificatesPagesWithContext(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificatesPagesWithContext", reflect.TypeOf((*MockACMAPI)(nil).ListCertificatesPagesWithContext), varargs...)
}

// ListTagsForCertificate mocks base method
func (m *MockACMAPI) ListTagsForCertificate(arg0 *acm.ListTagsForCertificateInput) (*acm.ListTagsForCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForCertificate", arg0)
	ret0, _ := ret[0].(*acm.ListTagsForCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForCertificate indicates an expected call of ListTagsForCertificate
func (mr *MockACMAPIMockRecorder) ListTagsForCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForCertificate", reflect.TypeOf((*MockACMAPI)(nil).ListTagsForCertificate), arg0)
}

// ListTagsForCertificateWithContext mocks base method
func (m *MockACMAPI) ListTagsForCertificateWithContext(arg0 aws.Context, arg1 *acm.ListTagsForCertificateInput, arg2 ...request.Option) (*acm.ListTagsForCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.ListTagsForCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForCertificateWithContext indicates an expected call of ListTagsForCertificateWithContext
func (mr *MockACMAPIMockRecorder) ListTagsForCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).ListTagsForCertificateWithContext), varargs...)
}

// ListTagsForCertificateRequest mocks base method
func (m *MockACMAPI) ListTagsForCertificateRequest(arg0 *acm.ListTagsForCertificateInput) (*request.Request, *acm.ListTagsForCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.ListTagsForCertificateOutput)
	return ret0, ret1
}

// ListTagsForCertificateRequest indicates an expected call of ListTagsForCertificateRequest
func (mr *MockACMAPIMockRecorder) ListTagsForCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).ListTagsForCertificateRequest), arg0)
}

// RemoveTagsFromCertificate mocks base method
func (m *MockACMAPI) RemoveTagsFromCertificate(arg0 *acm.RemoveTagsFromCertificateInput) (*acm.RemoveTagsFromCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTagsFromCertificate", arg0)
	ret0, _ := ret[0].(*acm.RemoveTagsFromCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveTagsFromCertificate indicates an expected call of RemoveTagsFromCertificate
func (mr *MockACMAPIMockRecorder) RemoveTagsFromCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTagsFromCertificate", reflect.TypeOf((*MockACMAPI)(nil).RemoveTagsFromCertificate), arg0)
}

// RemoveTagsFromCertificateWithContext mocks base method
func (m *MockACMAPI) RemoveTagsFromCertificateWithContext(arg0 aws.Context, arg1 *acm.RemoveTagsFromCertificateInput, arg2 ...request.Option) (*acm.RemoveTagsFromCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveTagsFromCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.RemoveTagsFromCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveTagsFromCertificateWithContext indicates an expected call of RemoveTagsFromCertificateWithContext
func (mr *MockACMAPIMockRecorder) RemoveTagsFromCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTagsFromCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).RemoveTagsFromCertificateWithContext), varargs...)
}

// RemoveTagsFromCertificateRequest mocks base method
func (m *MockACMAPI) RemoveTagsFromCertificateRequest(arg0 *acm.RemoveTagsFromCertificateInput) (*request.Request, *acm.RemoveTagsFromCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTagsFromCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.RemoveTagsFromCertificateOutput)
	return ret0, ret1
}

// RemoveTagsFromCertificateRequest indicates an expected call of RemoveTagsFromCertificateRequest
func (mr *MockACMAPIMockRecorder) RemoveTagsFromCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTagsFromCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).RemoveTagsFromCertificateRequest), arg0)
}

// RenewCertificate mocks base method
func (m *MockACMAPI) RenewCertificate(arg0 *acm.RenewCertificateInput) (*acm.RenewCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewCertificate", arg0)
	ret0, _ := ret[0].(*acm.RenewCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewCertificate indicates an expected call of RenewCertificate
func (mr *MockACMAPIMockRecorder) RenewCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewCertificate", reflect.TypeOf((*MockACMAPI)(nil).RenewCertificate), arg0)
}

// RenewCertificateWithContext mocks base method
func (m *MockACMAPI) RenewCertificateWithContext(arg0 aws.Context, arg1 *acm.RenewCertificateInput, arg2 ...request.Option) (*acm.RenewCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenewCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.RenewCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewCertificateWithContext indicates an expected call of RenewCertificateWithContext
func (mr *MockACMAPIMockRecorder) RenewCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).RenewCertificateWithContext), varargs...)
}

// RenewCertificateRequest mocks base method
func (m *MockACMAPI) RenewCertificateRequest(arg0 *acm.RenewCertificateInput) (*request.Request, *acm.RenewCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.RenewCertificateOutput)
	return ret0, ret1
}

// RenewCertificateRequest indicates an expected call of RenewCertificateRequest
func (mr *MockACMAPIMockRecorder) RenewCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).RenewCertificateRequest), arg0)
}

// RequestCertificate mocks base method
func (m *MockACMAPI) RequestCertificate(arg0 *acm.RequestCertificateInput) (*acm.RequestCertificateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestCertificate", arg0)
	ret0, _ := ret[0].(*acm.RequestCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestCertificate indicates an expected call of RequestCertificate
func (mr *MockACMAPIMockRecorder) RequestCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCertificate", reflect.TypeOf((*MockACMAPI)(nil).RequestCertificate), arg0)
}

// RequestCertificateWithContext mocks base method
func (m *MockACMAPI) RequestCertificateWithContext(arg0 aws.Context, arg1 *acm.RequestCertificateInput, arg2 ...request.Option) (*acm.RequestCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestCertificateWithContext", varargs...)
	ret0, _ := ret[0].(*acm.RequestCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestCertificateWithContext indicates an expected call of RequestCertificateWithContext
func (mr *MockACMAPIMockRecorder) RequestCertificateWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCertificateWithContext", reflect.TypeOf((*MockACMAPI)(nil).RequestCertificateWithContext), varargs...)
}

// RequestCertificateRequest mocks base method
func (m *MockACMAPI) RequestCertificateRequest(arg0 *acm.RequestCertificateInput) (*request.Request, *acm.RequestCertificateOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestCertificateRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.RequestCertificateOutput)
	return ret0, ret1
}

// RequestCertificateRequest indicates an expected call of RequestCertificateRequest
func (mr *MockACMAPIMockRecorder) RequestCertificateRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCertificateRequest", reflect.TypeOf((*MockACMAPI)(nil).RequestCertificateRequest), arg0)
}

// ResendValidationEmail mocks base method
func (m *MockACMAPI) ResendValidationEmail(arg0 *acm.ResendValidationEmailInput) (*acm.ResendValidationEmailOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResendValidationEmail", arg0)
	ret0, _ := ret[0].(*acm.ResendValidationEmailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResendValidationEmail indicates an expected call of ResendValidationEmail
func (mr *MockACMAPIMockRecorder) ResendValidationEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResendValidationEmail", reflect.TypeOf((*MockACMAPI)(nil).ResendValidationEmail), arg0)
}

// ResendValidationEmailWithContext mocks base method
func (m *MockACMAPI) ResendValidationEmailWithContext(arg0 aws.Context, arg1 *acm.ResendValidationEmailInput, arg2 ...request.Option) (*acm.ResendValidationEmailOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResendValidationEmailWithContext", varargs...)
	ret0, _ := ret[0].(*acm.ResendValidationEmailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResendValidationEmailWithContext indicates an expected call of ResendValidationEmailWithContext
func (mr *MockACMAPIMockRecorder) ResendValidationEmailWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResendValidationEmailWithContext", reflect.TypeOf((*MockACMAPI)(nil).ResendValidationEmailWithContext), varargs...)
}

// ResendValidationEmailRequest mocks base method
func (m *MockACMAPI) ResendValidationEmailRequest(arg0 *acm.ResendValidationEmailInput) (*request.Request, *acm.ResendValidationEmailOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResendValidationEmailRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.ResendValidationEmailOutput)
	return ret0, ret1
}

// ResendValidationEmailRequest indicates an expected call of ResendValidationEmailRequest
func (mr *MockACMAPIMockRecorder) ResendValidationEmailRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResendValidationEmailRequest", reflect.TypeOf((*MockACMAPI)(nil).ResendValidationEmailRequest), arg0)
}

// UpdateCertificateOptions mocks base method
func (m *MockACMAPI) UpdateCertificateOptions(arg0 *acm.UpdateCertificateOptionsInput) (*acm.UpdateCertificateOptionsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCertificateOptions", arg0)
	ret0, _ := ret[0].(*acm.UpdateCertificateOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCertificateOptions indicates an expected call of UpdateCertificateOptions
func (mr *MockACMAPIMockRecorder) UpdateCertificateOptions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificateOptions", reflect.TypeOf((*MockACMAPI)(nil).UpdateCertificateOptions), arg0)
}

// UpdateCertificateOptionsWithContext mocks base method
func (m *MockACMAPI) UpdateCertificateOptionsWithContext(arg0 aws.Context, arg1 *acm.UpdateCertificateOptionsInput, arg2 ...request.Option) (*acm.UpdateCertificateOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCertificateOptionsWithContext", varargs...)
	ret0, _ := ret[0].(*acm.UpdateCertificateOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCertificateOptionsWithContext indicates an expected call of UpdateCertificateOptionsWithContext
func (mr *MockACMAPIMockRecorder) UpdateCertificateOptionsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificateOptionsWithContext", reflect.TypeOf((*MockACMAPI)(nil).UpdateCertificateOptionsWithContext), varargs...)
}

// UpdateCertificateOptionsRequest mocks base method
func (m *MockACMAPI) UpdateCertificateOptionsRequest(arg0 *acm.UpdateCertificateOptionsInput) (*request.Request, *acm.UpdateCertificateOptionsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCertificateOptionsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*acm.UpdateCertificateOptionsOutput)
	return ret0, ret1
}

// UpdateCertificateOptionsRequest indicates an expected call of UpdateCertificateOptionsRequest
func (mr *MockACMAPIMockRecorder) UpdateCertificateOptionsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCertificateOptionsRequest", reflect.TypeOf((*MockACMAPI)(nil).UpdateCertificateOptionsRequest), arg0)
}

// WaitUntilCertificateValidated mocks base method
func (m *MockACMAPI) WaitUntilCertificateValidated(arg0 *acm.DescribeCertificateInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitUntilCertificateValidated", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilCertificateValidated indicates an expected call of WaitUntilCertificateValidated
func (mr *MockACMAPIMockRecorder) WaitUntilCertificateValidated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilCertificateValidated", reflect.TypeOf((*MockACMAPI)(nil).WaitUntilCertificateValidated), arg0)
}

// WaitUntilCertificateValidatedWithContext mocks base method
func (m *MockACMAPI) WaitUntilCertificateValidatedWithContext(arg0 aws.Context, arg1 *acm.DescribeCertificateInput, arg2 ...request.WaiterOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WaitUntilCertificateValidatedWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitUntilCertificateValidatedWithContext indicates an expected call of WaitUntilCertificateValidatedWithContext
func (mr *MockACMAPIMockRecorder) WaitUntilCertificateValidatedWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitUntilCertificateValidatedWithContext", reflect.TypeOf((*MockACMAPI)(nil).WaitUntilCertificateValidatedWithContext), varargs...)
}
