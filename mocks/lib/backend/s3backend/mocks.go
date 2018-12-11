// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/kraken/lib/backend/s3backend (interfaces: S3)

// Package mocks3backend is a generated GoMock package.
package mocks3backend

import (
	s3 "github.com/aws/aws-sdk-go/service/s3"
	s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockS3 is a mock of S3 interface
type MockS3 struct {
	ctrl     *gomock.Controller
	recorder *MockS3MockRecorder
}

// MockS3MockRecorder is the mock recorder for MockS3
type MockS3MockRecorder struct {
	mock *MockS3
}

// NewMockS3 creates a new mock instance
func NewMockS3(ctrl *gomock.Controller) *MockS3 {
	mock := &MockS3{ctrl: ctrl}
	mock.recorder = &MockS3MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockS3) EXPECT() *MockS3MockRecorder {
	return m.recorder
}

// Download mocks base method
func (m *MockS3) Download(arg0 io.WriterAt, arg1 *s3.GetObjectInput, arg2 ...func(*s3manager.Downloader)) (int64, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Download", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download
func (mr *MockS3MockRecorder) Download(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockS3)(nil).Download), varargs...)
}

// HeadObject mocks base method
func (m *MockS3) HeadObject(arg0 *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	ret := m.ctrl.Call(m, "HeadObject", arg0)
	ret0, _ := ret[0].(*s3.HeadObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadObject indicates an expected call of HeadObject
func (mr *MockS3MockRecorder) HeadObject(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadObject", reflect.TypeOf((*MockS3)(nil).HeadObject), arg0)
}

// ListObjectsPages mocks base method
func (m *MockS3) ListObjectsPages(arg0 *s3.ListObjectsInput, arg1 func(*s3.ListObjectsOutput, bool) bool) error {
	ret := m.ctrl.Call(m, "ListObjectsPages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListObjectsPages indicates an expected call of ListObjectsPages
func (mr *MockS3MockRecorder) ListObjectsPages(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsPages", reflect.TypeOf((*MockS3)(nil).ListObjectsPages), arg0, arg1)
}

// Upload mocks base method
func (m *MockS3) Upload(arg0 *s3manager.UploadInput, arg1 ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].(*s3manager.UploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload
func (mr *MockS3MockRecorder) Upload(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockS3)(nil).Upload), varargs...)
}
