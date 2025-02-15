package mocks

import (
	context "context"
	"github.com/selefra/selefra-provider-aws/constants"
	reflect "reflect"

	cloudhsmv2 "github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudhsmv2Client struct {
	ctrl     *gomock.Controller
	recorder *MockCloudhsmv2ClientMockRecorder
}

type MockCloudhsmv2ClientMockRecorder struct {
	mock *MockCloudhsmv2Client
}

func NewMockCloudhsmv2Client(ctrl *gomock.Controller) *MockCloudhsmv2Client {
	mock := &MockCloudhsmv2Client{ctrl: ctrl}
	mock.recorder = &MockCloudhsmv2ClientMockRecorder{mock}
	return mock
}

func (m *MockCloudhsmv2Client) EXPECT() *MockCloudhsmv2ClientMockRecorder {
	return m.recorder
}

func (m *MockCloudhsmv2Client) DescribeBackups(arg0 context.Context, arg1 *cloudhsmv2.DescribeBackupsInput, arg2 ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeBackups, varargs...)
	ret0, _ := ret[0].(*cloudhsmv2.DescribeBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudhsmv2ClientMockRecorder) DescribeBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeBackups, reflect.TypeOf((*MockCloudhsmv2Client)(nil).DescribeBackups), varargs...)
}

func (m *MockCloudhsmv2Client) DescribeClusters(arg0 context.Context, arg1 *cloudhsmv2.DescribeClustersInput, arg2 ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeClusters, varargs...)
	ret0, _ := ret[0].(*cloudhsmv2.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudhsmv2ClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeClusters, reflect.TypeOf((*MockCloudhsmv2Client)(nil).DescribeClusters), varargs...)
}

func (m *MockCloudhsmv2Client) ListTags(arg0 context.Context, arg1 *cloudhsmv2.ListTagsInput, arg2 ...func(*cloudhsmv2.Options)) (*cloudhsmv2.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTags, varargs...)
	ret0, _ := ret[0].(*cloudhsmv2.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudhsmv2ClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTags, reflect.TypeOf((*MockCloudhsmv2Client)(nil).ListTags), varargs...)
}
