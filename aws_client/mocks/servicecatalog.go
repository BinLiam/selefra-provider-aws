package mocks

import (
	context "context"
	"github.com/selefra/selefra-provider-aws/constants"
	reflect "reflect"

	servicecatalog "github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	gomock "github.com/golang/mock/gomock"
)

type MockServicecatalogClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicecatalogClientMockRecorder
}

type MockServicecatalogClientMockRecorder struct {
	mock *MockServicecatalogClient
}

func NewMockServicecatalogClient(ctrl *gomock.Controller) *MockServicecatalogClient {
	mock := &MockServicecatalogClient{ctrl: ctrl}
	mock.recorder = &MockServicecatalogClientMockRecorder{mock}
	return mock
}

func (m *MockServicecatalogClient) EXPECT() *MockServicecatalogClientMockRecorder {
	return m.recorder
}

func (m *MockServicecatalogClient) DescribeConstraint(arg0 context.Context, arg1 *servicecatalog.DescribeConstraintInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeConstraintOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeConstraint, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeConstraintOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeConstraint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeConstraint, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeConstraint), varargs...)
}

func (m *MockServicecatalogClient) DescribeCopyProductStatus(arg0 context.Context, arg1 *servicecatalog.DescribeCopyProductStatusInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeCopyProductStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeCopyProductStatus, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeCopyProductStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeCopyProductStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeCopyProductStatus, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeCopyProductStatus), varargs...)
}

func (m *MockServicecatalogClient) DescribePortfolio(arg0 context.Context, arg1 *servicecatalog.DescribePortfolioInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribePortfolioOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePortfolio, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribePortfolioOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribePortfolio(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePortfolio, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribePortfolio), varargs...)
}

func (m *MockServicecatalogClient) DescribePortfolioShareStatus(arg0 context.Context, arg1 *servicecatalog.DescribePortfolioShareStatusInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribePortfolioShareStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePortfolioShareStatus, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribePortfolioShareStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribePortfolioShareStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePortfolioShareStatus, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribePortfolioShareStatus), varargs...)
}

func (m *MockServicecatalogClient) DescribePortfolioShares(arg0 context.Context, arg1 *servicecatalog.DescribePortfolioSharesInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribePortfolioSharesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribePortfolioShares, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribePortfolioSharesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribePortfolioShares(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribePortfolioShares, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribePortfolioShares), varargs...)
}

func (m *MockServicecatalogClient) DescribeProduct(arg0 context.Context, arg1 *servicecatalog.DescribeProductInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProductOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProduct, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeProductOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeProduct(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProduct, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeProduct), varargs...)
}

func (m *MockServicecatalogClient) DescribeProductAsAdmin(arg0 context.Context, arg1 *servicecatalog.DescribeProductAsAdminInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProductAsAdminOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProductAsAdmin, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeProductAsAdminOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeProductAsAdmin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProductAsAdmin, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeProductAsAdmin), varargs...)
}

func (m *MockServicecatalogClient) DescribeProductView(arg0 context.Context, arg1 *servicecatalog.DescribeProductViewInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProductViewOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProductView, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeProductViewOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeProductView(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProductView, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeProductView), varargs...)
}

func (m *MockServicecatalogClient) DescribeProvisionedProduct(arg0 context.Context, arg1 *servicecatalog.DescribeProvisionedProductInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProvisionedProductOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProvisionedProduct, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeProvisionedProductOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeProvisionedProduct(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProvisionedProduct, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeProvisionedProduct), varargs...)
}

func (m *MockServicecatalogClient) DescribeProvisionedProductPlan(arg0 context.Context, arg1 *servicecatalog.DescribeProvisionedProductPlanInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProvisionedProductPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProvisionedProductPlan, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeProvisionedProductPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeProvisionedProductPlan(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProvisionedProductPlan, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeProvisionedProductPlan), varargs...)
}

func (m *MockServicecatalogClient) DescribeProvisioningArtifact(arg0 context.Context, arg1 *servicecatalog.DescribeProvisioningArtifactInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProvisioningArtifactOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProvisioningArtifact, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeProvisioningArtifactOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeProvisioningArtifact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProvisioningArtifact, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeProvisioningArtifact), varargs...)
}

func (m *MockServicecatalogClient) DescribeProvisioningParameters(arg0 context.Context, arg1 *servicecatalog.DescribeProvisioningParametersInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeProvisioningParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeProvisioningParameters, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeProvisioningParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeProvisioningParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeProvisioningParameters, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeProvisioningParameters), varargs...)
}

func (m *MockServicecatalogClient) DescribeRecord(arg0 context.Context, arg1 *servicecatalog.DescribeRecordInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeRecordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeRecord, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeRecord(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeRecord, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeRecord), varargs...)
}

func (m *MockServicecatalogClient) DescribeServiceAction(arg0 context.Context, arg1 *servicecatalog.DescribeServiceActionInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeServiceActionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeServiceAction, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeServiceActionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeServiceAction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeServiceAction, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeServiceAction), varargs...)
}

func (m *MockServicecatalogClient) DescribeServiceActionExecutionParameters(arg0 context.Context, arg1 *servicecatalog.DescribeServiceActionExecutionParametersInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeServiceActionExecutionParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeServiceActionExecutionParameters, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeServiceActionExecutionParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeServiceActionExecutionParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeServiceActionExecutionParameters, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeServiceActionExecutionParameters), varargs...)
}

func (m *MockServicecatalogClient) DescribeTagOption(arg0 context.Context, arg1 *servicecatalog.DescribeTagOptionInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.DescribeTagOptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.DescribeTagOption, varargs...)
	ret0, _ := ret[0].(*servicecatalog.DescribeTagOptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) DescribeTagOption(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.DescribeTagOption, reflect.TypeOf((*MockServicecatalogClient)(nil).DescribeTagOption), varargs...)
}

func (m *MockServicecatalogClient) GetAWSOrganizationsAccessStatus(arg0 context.Context, arg1 *servicecatalog.GetAWSOrganizationsAccessStatusInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.GetAWSOrganizationsAccessStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetAWSOrganizationsAccessStatus, varargs...)
	ret0, _ := ret[0].(*servicecatalog.GetAWSOrganizationsAccessStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) GetAWSOrganizationsAccessStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetAWSOrganizationsAccessStatus, reflect.TypeOf((*MockServicecatalogClient)(nil).GetAWSOrganizationsAccessStatus), varargs...)
}

func (m *MockServicecatalogClient) GetProvisionedProductOutputs(arg0 context.Context, arg1 *servicecatalog.GetProvisionedProductOutputsInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.GetProvisionedProductOutputsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetProvisionedProductOutputs, varargs...)
	ret0, _ := ret[0].(*servicecatalog.GetProvisionedProductOutputsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) GetProvisionedProductOutputs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetProvisionedProductOutputs, reflect.TypeOf((*MockServicecatalogClient)(nil).GetProvisionedProductOutputs), varargs...)
}

func (m *MockServicecatalogClient) ListAcceptedPortfolioShares(arg0 context.Context, arg1 *servicecatalog.ListAcceptedPortfolioSharesInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListAcceptedPortfolioSharesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListAcceptedPortfolioShares, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListAcceptedPortfolioSharesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListAcceptedPortfolioShares(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAcceptedPortfolioShares, reflect.TypeOf((*MockServicecatalogClient)(nil).ListAcceptedPortfolioShares), varargs...)
}

func (m *MockServicecatalogClient) ListBudgetsForResource(arg0 context.Context, arg1 *servicecatalog.ListBudgetsForResourceInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListBudgetsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBudgetsForResource, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListBudgetsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListBudgetsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBudgetsForResource, reflect.TypeOf((*MockServicecatalogClient)(nil).ListBudgetsForResource), varargs...)
}

func (m *MockServicecatalogClient) ListConstraintsForPortfolio(arg0 context.Context, arg1 *servicecatalog.ListConstraintsForPortfolioInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListConstraintsForPortfolioOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListConstraintsForPortfolio, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListConstraintsForPortfolioOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListConstraintsForPortfolio(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListConstraintsForPortfolio, reflect.TypeOf((*MockServicecatalogClient)(nil).ListConstraintsForPortfolio), varargs...)
}

func (m *MockServicecatalogClient) ListLaunchPaths(arg0 context.Context, arg1 *servicecatalog.ListLaunchPathsInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListLaunchPathsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListLaunchPaths, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListLaunchPathsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListLaunchPaths(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListLaunchPaths, reflect.TypeOf((*MockServicecatalogClient)(nil).ListLaunchPaths), varargs...)
}

func (m *MockServicecatalogClient) ListOrganizationPortfolioAccess(arg0 context.Context, arg1 *servicecatalog.ListOrganizationPortfolioAccessInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListOrganizationPortfolioAccessOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListOrganizationPortfolioAccess, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListOrganizationPortfolioAccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListOrganizationPortfolioAccess(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListOrganizationPortfolioAccess, reflect.TypeOf((*MockServicecatalogClient)(nil).ListOrganizationPortfolioAccess), varargs...)
}

func (m *MockServicecatalogClient) ListPortfolioAccess(arg0 context.Context, arg1 *servicecatalog.ListPortfolioAccessInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListPortfolioAccessOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPortfolioAccess, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListPortfolioAccessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListPortfolioAccess(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPortfolioAccess, reflect.TypeOf((*MockServicecatalogClient)(nil).ListPortfolioAccess), varargs...)
}

func (m *MockServicecatalogClient) ListPortfolios(arg0 context.Context, arg1 *servicecatalog.ListPortfoliosInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListPortfoliosOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPortfolios, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListPortfoliosOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListPortfolios(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPortfolios, reflect.TypeOf((*MockServicecatalogClient)(nil).ListPortfolios), varargs...)
}

func (m *MockServicecatalogClient) ListPortfoliosForProduct(arg0 context.Context, arg1 *servicecatalog.ListPortfoliosForProductInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListPortfoliosForProductOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPortfoliosForProduct, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListPortfoliosForProductOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListPortfoliosForProduct(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPortfoliosForProduct, reflect.TypeOf((*MockServicecatalogClient)(nil).ListPortfoliosForProduct), varargs...)
}

func (m *MockServicecatalogClient) ListPrincipalsForPortfolio(arg0 context.Context, arg1 *servicecatalog.ListPrincipalsForPortfolioInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListPrincipalsForPortfolioOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListPrincipalsForPortfolio, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListPrincipalsForPortfolioOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListPrincipalsForPortfolio(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListPrincipalsForPortfolio, reflect.TypeOf((*MockServicecatalogClient)(nil).ListPrincipalsForPortfolio), varargs...)
}

func (m *MockServicecatalogClient) ListProvisionedProductPlans(arg0 context.Context, arg1 *servicecatalog.ListProvisionedProductPlansInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListProvisionedProductPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProvisionedProductPlans, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListProvisionedProductPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListProvisionedProductPlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProvisionedProductPlans, reflect.TypeOf((*MockServicecatalogClient)(nil).ListProvisionedProductPlans), varargs...)
}

func (m *MockServicecatalogClient) ListProvisioningArtifacts(arg0 context.Context, arg1 *servicecatalog.ListProvisioningArtifactsInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListProvisioningArtifactsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProvisioningArtifacts, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListProvisioningArtifactsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListProvisioningArtifacts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProvisioningArtifacts, reflect.TypeOf((*MockServicecatalogClient)(nil).ListProvisioningArtifacts), varargs...)
}

func (m *MockServicecatalogClient) ListProvisioningArtifactsForServiceAction(arg0 context.Context, arg1 *servicecatalog.ListProvisioningArtifactsForServiceActionInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListProvisioningArtifactsForServiceActionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListProvisioningArtifactsForServiceAction, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListProvisioningArtifactsForServiceActionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListProvisioningArtifactsForServiceAction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListProvisioningArtifactsForServiceAction, reflect.TypeOf((*MockServicecatalogClient)(nil).ListProvisioningArtifactsForServiceAction), varargs...)
}

func (m *MockServicecatalogClient) ListRecordHistory(arg0 context.Context, arg1 *servicecatalog.ListRecordHistoryInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListRecordHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListRecordHistory, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListRecordHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListRecordHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRecordHistory, reflect.TypeOf((*MockServicecatalogClient)(nil).ListRecordHistory), varargs...)
}

func (m *MockServicecatalogClient) ListResourcesForTagOption(arg0 context.Context, arg1 *servicecatalog.ListResourcesForTagOptionInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListResourcesForTagOptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListResourcesForTagOption, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListResourcesForTagOptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListResourcesForTagOption(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResourcesForTagOption, reflect.TypeOf((*MockServicecatalogClient)(nil).ListResourcesForTagOption), varargs...)
}

func (m *MockServicecatalogClient) ListServiceActions(arg0 context.Context, arg1 *servicecatalog.ListServiceActionsInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListServiceActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServiceActions, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListServiceActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListServiceActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServiceActions, reflect.TypeOf((*MockServicecatalogClient)(nil).ListServiceActions), varargs...)
}

func (m *MockServicecatalogClient) ListServiceActionsForProvisioningArtifact(arg0 context.Context, arg1 *servicecatalog.ListServiceActionsForProvisioningArtifactInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListServiceActionsForProvisioningArtifactOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListServiceActionsForProvisioningArtifact, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListServiceActionsForProvisioningArtifactOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListServiceActionsForProvisioningArtifact(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListServiceActionsForProvisioningArtifact, reflect.TypeOf((*MockServicecatalogClient)(nil).ListServiceActionsForProvisioningArtifact), varargs...)
}

func (m *MockServicecatalogClient) ListStackInstancesForProvisionedProduct(arg0 context.Context, arg1 *servicecatalog.ListStackInstancesForProvisionedProductInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListStackInstancesForProvisionedProductOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListStackInstancesForProvisionedProduct, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListStackInstancesForProvisionedProductOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListStackInstancesForProvisionedProduct(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListStackInstancesForProvisionedProduct, reflect.TypeOf((*MockServicecatalogClient)(nil).ListStackInstancesForProvisionedProduct), varargs...)
}

func (m *MockServicecatalogClient) ListTagOptions(arg0 context.Context, arg1 *servicecatalog.ListTagOptionsInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.ListTagOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListTagOptions, varargs...)
	ret0, _ := ret[0].(*servicecatalog.ListTagOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) ListTagOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListTagOptions, reflect.TypeOf((*MockServicecatalogClient)(nil).ListTagOptions), varargs...)
}

func (m *MockServicecatalogClient) SearchProducts(arg0 context.Context, arg1 *servicecatalog.SearchProductsInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.SearchProductsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchProducts, varargs...)
	ret0, _ := ret[0].(*servicecatalog.SearchProductsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) SearchProducts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchProducts, reflect.TypeOf((*MockServicecatalogClient)(nil).SearchProducts), varargs...)
}

func (m *MockServicecatalogClient) SearchProductsAsAdmin(arg0 context.Context, arg1 *servicecatalog.SearchProductsAsAdminInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.SearchProductsAsAdminOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchProductsAsAdmin, varargs...)
	ret0, _ := ret[0].(*servicecatalog.SearchProductsAsAdminOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) SearchProductsAsAdmin(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchProductsAsAdmin, reflect.TypeOf((*MockServicecatalogClient)(nil).SearchProductsAsAdmin), varargs...)
}

func (m *MockServicecatalogClient) SearchProvisionedProducts(arg0 context.Context, arg1 *servicecatalog.SearchProvisionedProductsInput, arg2 ...func(*servicecatalog.Options)) (*servicecatalog.SearchProvisionedProductsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.SearchProvisionedProducts, varargs...)
	ret0, _ := ret[0].(*servicecatalog.SearchProvisionedProductsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicecatalogClientMockRecorder) SearchProvisionedProducts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.SearchProvisionedProducts, reflect.TypeOf((*MockServicecatalogClient)(nil).SearchProvisionedProducts), varargs...)
}
