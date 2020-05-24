// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/azureclient/mgmt/network (interfaces: InterfacesClient,LoadBalancersClient,PrivateEndpointsClient,PrivateLinkServicesClient,PublicIPAddressesClient,SubnetsClient,VirtualNetworksClient)

// Package mock_network is a generated GoMock package.
package mock_network

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-07-01/network"
	gomock "github.com/golang/mock/gomock"
)

// MockInterfacesClient is a mock of InterfacesClient interface
type MockInterfacesClient struct {
	ctrl     *gomock.Controller
	recorder *MockInterfacesClientMockRecorder
}

// MockInterfacesClientMockRecorder is the mock recorder for MockInterfacesClient
type MockInterfacesClientMockRecorder struct {
	mock *MockInterfacesClient
}

// NewMockInterfacesClient creates a new mock instance
func NewMockInterfacesClient(ctrl *gomock.Controller) *MockInterfacesClient {
	mock := &MockInterfacesClient{ctrl: ctrl}
	mock.recorder = &MockInterfacesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfacesClient) EXPECT() *MockInterfacesClientMockRecorder {
	return m.recorder
}

// CreateOrUpdateAndWait mocks base method
func (m *MockInterfacesClient) CreateOrUpdateAndWait(arg0 context.Context, arg1, arg2 string, arg3 network.Interface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateAndWait", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateAndWait indicates an expected call of CreateOrUpdateAndWait
func (mr *MockInterfacesClientMockRecorder) CreateOrUpdateAndWait(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateAndWait", reflect.TypeOf((*MockInterfacesClient)(nil).CreateOrUpdateAndWait), arg0, arg1, arg2, arg3)
}

// DeleteAndWait mocks base method
func (m *MockInterfacesClient) DeleteAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait
func (mr *MockInterfacesClientMockRecorder) DeleteAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockInterfacesClient)(nil).DeleteAndWait), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockInterfacesClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockInterfacesClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterfacesClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// MockLoadBalancersClient is a mock of LoadBalancersClient interface
type MockLoadBalancersClient struct {
	ctrl     *gomock.Controller
	recorder *MockLoadBalancersClientMockRecorder
}

// MockLoadBalancersClientMockRecorder is the mock recorder for MockLoadBalancersClient
type MockLoadBalancersClientMockRecorder struct {
	mock *MockLoadBalancersClient
}

// NewMockLoadBalancersClient creates a new mock instance
func NewMockLoadBalancersClient(ctrl *gomock.Controller) *MockLoadBalancersClient {
	mock := &MockLoadBalancersClient{ctrl: ctrl}
	mock.recorder = &MockLoadBalancersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoadBalancersClient) EXPECT() *MockLoadBalancersClientMockRecorder {
	return m.recorder
}

// CreateOrUpdateAndWait mocks base method
func (m *MockLoadBalancersClient) CreateOrUpdateAndWait(arg0 context.Context, arg1, arg2 string, arg3 network.LoadBalancer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateAndWait", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateAndWait indicates an expected call of CreateOrUpdateAndWait
func (mr *MockLoadBalancersClientMockRecorder) CreateOrUpdateAndWait(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateAndWait", reflect.TypeOf((*MockLoadBalancersClient)(nil).CreateOrUpdateAndWait), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockLoadBalancersClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockLoadBalancersClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLoadBalancersClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// MockPrivateEndpointsClient is a mock of PrivateEndpointsClient interface
type MockPrivateEndpointsClient struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointsClientMockRecorder
}

// MockPrivateEndpointsClientMockRecorder is the mock recorder for MockPrivateEndpointsClient
type MockPrivateEndpointsClientMockRecorder struct {
	mock *MockPrivateEndpointsClient
}

// NewMockPrivateEndpointsClient creates a new mock instance
func NewMockPrivateEndpointsClient(ctrl *gomock.Controller) *MockPrivateEndpointsClient {
	mock := &MockPrivateEndpointsClient{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrivateEndpointsClient) EXPECT() *MockPrivateEndpointsClientMockRecorder {
	return m.recorder
}

// CreateOrUpdateAndWait mocks base method
func (m *MockPrivateEndpointsClient) CreateOrUpdateAndWait(arg0 context.Context, arg1, arg2 string, arg3 network.PrivateEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateAndWait", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateAndWait indicates an expected call of CreateOrUpdateAndWait
func (mr *MockPrivateEndpointsClientMockRecorder) CreateOrUpdateAndWait(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateAndWait", reflect.TypeOf((*MockPrivateEndpointsClient)(nil).CreateOrUpdateAndWait), arg0, arg1, arg2, arg3)
}

// DeleteAndWait mocks base method
func (m *MockPrivateEndpointsClient) DeleteAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait
func (mr *MockPrivateEndpointsClientMockRecorder) DeleteAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockPrivateEndpointsClient)(nil).DeleteAndWait), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockPrivateEndpointsClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.PrivateEndpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.PrivateEndpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPrivateEndpointsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPrivateEndpointsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// MockPrivateLinkServicesClient is a mock of PrivateLinkServicesClient interface
type MockPrivateLinkServicesClient struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateLinkServicesClientMockRecorder
}

// MockPrivateLinkServicesClientMockRecorder is the mock recorder for MockPrivateLinkServicesClient
type MockPrivateLinkServicesClientMockRecorder struct {
	mock *MockPrivateLinkServicesClient
}

// NewMockPrivateLinkServicesClient creates a new mock instance
func NewMockPrivateLinkServicesClient(ctrl *gomock.Controller) *MockPrivateLinkServicesClient {
	mock := &MockPrivateLinkServicesClient{ctrl: ctrl}
	mock.recorder = &MockPrivateLinkServicesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrivateLinkServicesClient) EXPECT() *MockPrivateLinkServicesClientMockRecorder {
	return m.recorder
}

// DeletePrivateEndpointConnection mocks base method
func (m *MockPrivateLinkServicesClient) DeletePrivateEndpointConnection(arg0 context.Context, arg1, arg2, arg3 string) (network.PrivateLinkServicesDeletePrivateEndpointConnectionFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePrivateEndpointConnection", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.PrivateLinkServicesDeletePrivateEndpointConnectionFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePrivateEndpointConnection indicates an expected call of DeletePrivateEndpointConnection
func (mr *MockPrivateLinkServicesClientMockRecorder) DeletePrivateEndpointConnection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePrivateEndpointConnection", reflect.TypeOf((*MockPrivateLinkServicesClient)(nil).DeletePrivateEndpointConnection), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockPrivateLinkServicesClient) List(arg0 context.Context, arg1 string) ([]network.PrivateLinkService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]network.PrivateLinkService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockPrivateLinkServicesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPrivateLinkServicesClient)(nil).List), arg0, arg1)
}

// MockPublicIPAddressesClient is a mock of PublicIPAddressesClient interface
type MockPublicIPAddressesClient struct {
	ctrl     *gomock.Controller
	recorder *MockPublicIPAddressesClientMockRecorder
}

// MockPublicIPAddressesClientMockRecorder is the mock recorder for MockPublicIPAddressesClient
type MockPublicIPAddressesClientMockRecorder struct {
	mock *MockPublicIPAddressesClient
}

// NewMockPublicIPAddressesClient creates a new mock instance
func NewMockPublicIPAddressesClient(ctrl *gomock.Controller) *MockPublicIPAddressesClient {
	mock := &MockPublicIPAddressesClient{ctrl: ctrl}
	mock.recorder = &MockPublicIPAddressesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublicIPAddressesClient) EXPECT() *MockPublicIPAddressesClientMockRecorder {
	return m.recorder
}

// DeleteAndWait mocks base method
func (m *MockPublicIPAddressesClient) DeleteAndWait(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAndWait", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAndWait indicates an expected call of DeleteAndWait
func (mr *MockPublicIPAddressesClientMockRecorder) DeleteAndWait(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAndWait", reflect.TypeOf((*MockPublicIPAddressesClient)(nil).DeleteAndWait), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockPublicIPAddressesClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.PublicIPAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.PublicIPAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPublicIPAddressesClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPublicIPAddressesClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockPublicIPAddressesClient) List(arg0 context.Context, arg1 string) ([]network.PublicIPAddress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]network.PublicIPAddress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockPublicIPAddressesClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPublicIPAddressesClient)(nil).List), arg0, arg1)
}

// MockSubnetsClient is a mock of SubnetsClient interface
type MockSubnetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubnetsClientMockRecorder
}

// MockSubnetsClientMockRecorder is the mock recorder for MockSubnetsClient
type MockSubnetsClientMockRecorder struct {
	mock *MockSubnetsClient
}

// NewMockSubnetsClient creates a new mock instance
func NewMockSubnetsClient(ctrl *gomock.Controller) *MockSubnetsClient {
	mock := &MockSubnetsClient{ctrl: ctrl}
	mock.recorder = &MockSubnetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSubnetsClient) EXPECT() *MockSubnetsClientMockRecorder {
	return m.recorder
}

// CreateOrUpdateAndWait mocks base method
func (m *MockSubnetsClient) CreateOrUpdateAndWait(arg0 context.Context, arg1, arg2, arg3 string, arg4 network.Subnet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateAndWait", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateAndWait indicates an expected call of CreateOrUpdateAndWait
func (mr *MockSubnetsClientMockRecorder) CreateOrUpdateAndWait(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateAndWait", reflect.TypeOf((*MockSubnetsClient)(nil).CreateOrUpdateAndWait), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method
func (m *MockSubnetsClient) Get(arg0 context.Context, arg1, arg2, arg3, arg4 string) (network.Subnet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(network.Subnet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSubnetsClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubnetsClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// MockVirtualNetworksClient is a mock of VirtualNetworksClient interface
type MockVirtualNetworksClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualNetworksClientMockRecorder
}

// MockVirtualNetworksClientMockRecorder is the mock recorder for MockVirtualNetworksClient
type MockVirtualNetworksClientMockRecorder struct {
	mock *MockVirtualNetworksClient
}

// NewMockVirtualNetworksClient creates a new mock instance
func NewMockVirtualNetworksClient(ctrl *gomock.Controller) *MockVirtualNetworksClient {
	mock := &MockVirtualNetworksClient{ctrl: ctrl}
	mock.recorder = &MockVirtualNetworksClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVirtualNetworksClient) EXPECT() *MockVirtualNetworksClientMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockVirtualNetworksClient) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 network.VirtualNetwork) (network.VirtualNetworksCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.VirtualNetworksCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockVirtualNetworksClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockVirtualNetworksClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockVirtualNetworksClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.VirtualNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.VirtualNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockVirtualNetworksClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualNetworksClient)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockVirtualNetworksClient) List(arg0 context.Context, arg1 string) ([]network.VirtualNetwork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]network.VirtualNetwork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockVirtualNetworksClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualNetworksClient)(nil).List), arg0, arg1)
}
