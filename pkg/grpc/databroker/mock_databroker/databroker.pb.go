// Code generated by MockGen. DO NOT EDIT.
// Source: databroker.pb.go
//
// Generated by this command:
//
//	mockgen -source=databroker.pb.go -destination ./mock_databroker/databroker.pb.go DataBrokerServiceClient
//

// Package mock_databroker is a generated GoMock package.
package mock_databroker

import (
	context "context"
	reflect "reflect"

	databroker "github.com/pomerium/pomerium/pkg/grpc/databroker"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockisSyncLatestResponse_Response is a mock of isSyncLatestResponse_Response interface.
type MockisSyncLatestResponse_Response struct {
	ctrl     *gomock.Controller
	recorder *MockisSyncLatestResponse_ResponseMockRecorder
}

// MockisSyncLatestResponse_ResponseMockRecorder is the mock recorder for MockisSyncLatestResponse_Response.
type MockisSyncLatestResponse_ResponseMockRecorder struct {
	mock *MockisSyncLatestResponse_Response
}

// NewMockisSyncLatestResponse_Response creates a new mock instance.
func NewMockisSyncLatestResponse_Response(ctrl *gomock.Controller) *MockisSyncLatestResponse_Response {
	mock := &MockisSyncLatestResponse_Response{ctrl: ctrl}
	mock.recorder = &MockisSyncLatestResponse_ResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisSyncLatestResponse_Response) EXPECT() *MockisSyncLatestResponse_ResponseMockRecorder {
	return m.recorder
}

// isSyncLatestResponse_Response mocks base method.
func (m *MockisSyncLatestResponse_Response) isSyncLatestResponse_Response() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isSyncLatestResponse_Response")
}

// isSyncLatestResponse_Response indicates an expected call of isSyncLatestResponse_Response.
func (mr *MockisSyncLatestResponse_ResponseMockRecorder) isSyncLatestResponse_Response() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isSyncLatestResponse_Response", reflect.TypeOf((*MockisSyncLatestResponse_Response)(nil).isSyncLatestResponse_Response))
}

// MockDataBrokerServiceClient is a mock of DataBrokerServiceClient interface.
type MockDataBrokerServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDataBrokerServiceClientMockRecorder
}

// MockDataBrokerServiceClientMockRecorder is the mock recorder for MockDataBrokerServiceClient.
type MockDataBrokerServiceClientMockRecorder struct {
	mock *MockDataBrokerServiceClient
}

// NewMockDataBrokerServiceClient creates a new mock instance.
func NewMockDataBrokerServiceClient(ctrl *gomock.Controller) *MockDataBrokerServiceClient {
	mock := &MockDataBrokerServiceClient{ctrl: ctrl}
	mock.recorder = &MockDataBrokerServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataBrokerServiceClient) EXPECT() *MockDataBrokerServiceClientMockRecorder {
	return m.recorder
}

// AcquireLease mocks base method.
func (m *MockDataBrokerServiceClient) AcquireLease(ctx context.Context, in *databroker.AcquireLeaseRequest, opts ...grpc.CallOption) (*databroker.AcquireLeaseResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AcquireLease", varargs...)
	ret0, _ := ret[0].(*databroker.AcquireLeaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireLease indicates an expected call of AcquireLease.
func (mr *MockDataBrokerServiceClientMockRecorder) AcquireLease(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireLease", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).AcquireLease), varargs...)
}

// Get mocks base method.
func (m *MockDataBrokerServiceClient) Get(ctx context.Context, in *databroker.GetRequest, opts ...grpc.CallOption) (*databroker.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*databroker.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDataBrokerServiceClientMockRecorder) Get(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).Get), varargs...)
}

// ListTypes mocks base method.
func (m *MockDataBrokerServiceClient) ListTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*databroker.ListTypesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTypes", varargs...)
	ret0, _ := ret[0].(*databroker.ListTypesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTypes indicates an expected call of ListTypes.
func (mr *MockDataBrokerServiceClientMockRecorder) ListTypes(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTypes", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).ListTypes), varargs...)
}

// Patch mocks base method.
func (m *MockDataBrokerServiceClient) Patch(ctx context.Context, in *databroker.PatchRequest, opts ...grpc.CallOption) (*databroker.PatchResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*databroker.PatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockDataBrokerServiceClientMockRecorder) Patch(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).Patch), varargs...)
}

// Put mocks base method.
func (m *MockDataBrokerServiceClient) Put(ctx context.Context, in *databroker.PutRequest, opts ...grpc.CallOption) (*databroker.PutResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Put", varargs...)
	ret0, _ := ret[0].(*databroker.PutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockDataBrokerServiceClientMockRecorder) Put(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).Put), varargs...)
}

// Query mocks base method.
func (m *MockDataBrokerServiceClient) Query(ctx context.Context, in *databroker.QueryRequest, opts ...grpc.CallOption) (*databroker.QueryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*databroker.QueryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDataBrokerServiceClientMockRecorder) Query(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).Query), varargs...)
}

// ReleaseLease mocks base method.
func (m *MockDataBrokerServiceClient) ReleaseLease(ctx context.Context, in *databroker.ReleaseLeaseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReleaseLease", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseLease indicates an expected call of ReleaseLease.
func (mr *MockDataBrokerServiceClientMockRecorder) ReleaseLease(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseLease", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).ReleaseLease), varargs...)
}

// RenewLease mocks base method.
func (m *MockDataBrokerServiceClient) RenewLease(ctx context.Context, in *databroker.RenewLeaseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenewLease", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewLease indicates an expected call of RenewLease.
func (mr *MockDataBrokerServiceClientMockRecorder) RenewLease(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewLease", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).RenewLease), varargs...)
}

// SetOptions mocks base method.
func (m *MockDataBrokerServiceClient) SetOptions(ctx context.Context, in *databroker.SetOptionsRequest, opts ...grpc.CallOption) (*databroker.SetOptionsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetOptions", varargs...)
	ret0, _ := ret[0].(*databroker.SetOptionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetOptions indicates an expected call of SetOptions.
func (mr *MockDataBrokerServiceClientMockRecorder) SetOptions(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOptions", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).SetOptions), varargs...)
}

// Sync mocks base method.
func (m *MockDataBrokerServiceClient) Sync(ctx context.Context, in *databroker.SyncRequest, opts ...grpc.CallOption) (databroker.DataBrokerService_SyncClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sync", varargs...)
	ret0, _ := ret[0].(databroker.DataBrokerService_SyncClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sync indicates an expected call of Sync.
func (mr *MockDataBrokerServiceClientMockRecorder) Sync(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).Sync), varargs...)
}

// SyncLatest mocks base method.
func (m *MockDataBrokerServiceClient) SyncLatest(ctx context.Context, in *databroker.SyncLatestRequest, opts ...grpc.CallOption) (databroker.DataBrokerService_SyncLatestClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncLatest", varargs...)
	ret0, _ := ret[0].(databroker.DataBrokerService_SyncLatestClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncLatest indicates an expected call of SyncLatest.
func (mr *MockDataBrokerServiceClientMockRecorder) SyncLatest(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncLatest", reflect.TypeOf((*MockDataBrokerServiceClient)(nil).SyncLatest), varargs...)
}

// MockDataBrokerService_SyncClient is a mock of DataBrokerService_SyncClient interface.
type MockDataBrokerService_SyncClient struct {
	ctrl     *gomock.Controller
	recorder *MockDataBrokerService_SyncClientMockRecorder
}

// MockDataBrokerService_SyncClientMockRecorder is the mock recorder for MockDataBrokerService_SyncClient.
type MockDataBrokerService_SyncClientMockRecorder struct {
	mock *MockDataBrokerService_SyncClient
}

// NewMockDataBrokerService_SyncClient creates a new mock instance.
func NewMockDataBrokerService_SyncClient(ctrl *gomock.Controller) *MockDataBrokerService_SyncClient {
	mock := &MockDataBrokerService_SyncClient{ctrl: ctrl}
	mock.recorder = &MockDataBrokerService_SyncClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataBrokerService_SyncClient) EXPECT() *MockDataBrokerService_SyncClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockDataBrokerService_SyncClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockDataBrokerService_SyncClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockDataBrokerService_SyncClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockDataBrokerService_SyncClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockDataBrokerService_SyncClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDataBrokerService_SyncClient)(nil).Context))
}

// Header mocks base method.
func (m *MockDataBrokerService_SyncClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockDataBrokerService_SyncClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockDataBrokerService_SyncClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockDataBrokerService_SyncClient) Recv() (*databroker.SyncResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*databroker.SyncResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockDataBrokerService_SyncClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockDataBrokerService_SyncClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockDataBrokerService_SyncClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockDataBrokerService_SyncClientMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDataBrokerService_SyncClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockDataBrokerService_SyncClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockDataBrokerService_SyncClientMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDataBrokerService_SyncClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockDataBrokerService_SyncClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockDataBrokerService_SyncClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockDataBrokerService_SyncClient)(nil).Trailer))
}

// MockDataBrokerService_SyncLatestClient is a mock of DataBrokerService_SyncLatestClient interface.
type MockDataBrokerService_SyncLatestClient struct {
	ctrl     *gomock.Controller
	recorder *MockDataBrokerService_SyncLatestClientMockRecorder
}

// MockDataBrokerService_SyncLatestClientMockRecorder is the mock recorder for MockDataBrokerService_SyncLatestClient.
type MockDataBrokerService_SyncLatestClientMockRecorder struct {
	mock *MockDataBrokerService_SyncLatestClient
}

// NewMockDataBrokerService_SyncLatestClient creates a new mock instance.
func NewMockDataBrokerService_SyncLatestClient(ctrl *gomock.Controller) *MockDataBrokerService_SyncLatestClient {
	mock := &MockDataBrokerService_SyncLatestClient{ctrl: ctrl}
	mock.recorder = &MockDataBrokerService_SyncLatestClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataBrokerService_SyncLatestClient) EXPECT() *MockDataBrokerService_SyncLatestClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockDataBrokerService_SyncLatestClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockDataBrokerService_SyncLatestClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockDataBrokerService_SyncLatestClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockDataBrokerService_SyncLatestClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockDataBrokerService_SyncLatestClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDataBrokerService_SyncLatestClient)(nil).Context))
}

// Header mocks base method.
func (m *MockDataBrokerService_SyncLatestClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockDataBrokerService_SyncLatestClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockDataBrokerService_SyncLatestClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockDataBrokerService_SyncLatestClient) Recv() (*databroker.SyncLatestResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*databroker.SyncLatestResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockDataBrokerService_SyncLatestClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockDataBrokerService_SyncLatestClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockDataBrokerService_SyncLatestClient) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockDataBrokerService_SyncLatestClientMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDataBrokerService_SyncLatestClient)(nil).RecvMsg), m)
}

// SendMsg mocks base method.
func (m_2 *MockDataBrokerService_SyncLatestClient) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockDataBrokerService_SyncLatestClientMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDataBrokerService_SyncLatestClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockDataBrokerService_SyncLatestClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockDataBrokerService_SyncLatestClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockDataBrokerService_SyncLatestClient)(nil).Trailer))
}

// MockDataBrokerServiceServer is a mock of DataBrokerServiceServer interface.
type MockDataBrokerServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockDataBrokerServiceServerMockRecorder
}

// MockDataBrokerServiceServerMockRecorder is the mock recorder for MockDataBrokerServiceServer.
type MockDataBrokerServiceServerMockRecorder struct {
	mock *MockDataBrokerServiceServer
}

// NewMockDataBrokerServiceServer creates a new mock instance.
func NewMockDataBrokerServiceServer(ctrl *gomock.Controller) *MockDataBrokerServiceServer {
	mock := &MockDataBrokerServiceServer{ctrl: ctrl}
	mock.recorder = &MockDataBrokerServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataBrokerServiceServer) EXPECT() *MockDataBrokerServiceServerMockRecorder {
	return m.recorder
}

// AcquireLease mocks base method.
func (m *MockDataBrokerServiceServer) AcquireLease(arg0 context.Context, arg1 *databroker.AcquireLeaseRequest) (*databroker.AcquireLeaseResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcquireLease", arg0, arg1)
	ret0, _ := ret[0].(*databroker.AcquireLeaseResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireLease indicates an expected call of AcquireLease.
func (mr *MockDataBrokerServiceServerMockRecorder) AcquireLease(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireLease", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).AcquireLease), arg0, arg1)
}

// Get mocks base method.
func (m *MockDataBrokerServiceServer) Get(arg0 context.Context, arg1 *databroker.GetRequest) (*databroker.GetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*databroker.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDataBrokerServiceServerMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).Get), arg0, arg1)
}

// ListTypes mocks base method.
func (m *MockDataBrokerServiceServer) ListTypes(arg0 context.Context, arg1 *emptypb.Empty) (*databroker.ListTypesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTypes", arg0, arg1)
	ret0, _ := ret[0].(*databroker.ListTypesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTypes indicates an expected call of ListTypes.
func (mr *MockDataBrokerServiceServerMockRecorder) ListTypes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTypes", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).ListTypes), arg0, arg1)
}

// Patch mocks base method.
func (m *MockDataBrokerServiceServer) Patch(arg0 context.Context, arg1 *databroker.PatchRequest) (*databroker.PatchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", arg0, arg1)
	ret0, _ := ret[0].(*databroker.PatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockDataBrokerServiceServerMockRecorder) Patch(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).Patch), arg0, arg1)
}

// Put mocks base method.
func (m *MockDataBrokerServiceServer) Put(arg0 context.Context, arg1 *databroker.PutRequest) (*databroker.PutResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(*databroker.PutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockDataBrokerServiceServerMockRecorder) Put(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).Put), arg0, arg1)
}

// Query mocks base method.
func (m *MockDataBrokerServiceServer) Query(arg0 context.Context, arg1 *databroker.QueryRequest) (*databroker.QueryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0, arg1)
	ret0, _ := ret[0].(*databroker.QueryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockDataBrokerServiceServerMockRecorder) Query(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).Query), arg0, arg1)
}

// ReleaseLease mocks base method.
func (m *MockDataBrokerServiceServer) ReleaseLease(arg0 context.Context, arg1 *databroker.ReleaseLeaseRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseLease", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReleaseLease indicates an expected call of ReleaseLease.
func (mr *MockDataBrokerServiceServerMockRecorder) ReleaseLease(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseLease", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).ReleaseLease), arg0, arg1)
}

// RenewLease mocks base method.
func (m *MockDataBrokerServiceServer) RenewLease(arg0 context.Context, arg1 *databroker.RenewLeaseRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewLease", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewLease indicates an expected call of RenewLease.
func (mr *MockDataBrokerServiceServerMockRecorder) RenewLease(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewLease", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).RenewLease), arg0, arg1)
}

// SetOptions mocks base method.
func (m *MockDataBrokerServiceServer) SetOptions(arg0 context.Context, arg1 *databroker.SetOptionsRequest) (*databroker.SetOptionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOptions", arg0, arg1)
	ret0, _ := ret[0].(*databroker.SetOptionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetOptions indicates an expected call of SetOptions.
func (mr *MockDataBrokerServiceServerMockRecorder) SetOptions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOptions", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).SetOptions), arg0, arg1)
}

// Sync mocks base method.
func (m *MockDataBrokerServiceServer) Sync(arg0 *databroker.SyncRequest, arg1 databroker.DataBrokerService_SyncServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockDataBrokerServiceServerMockRecorder) Sync(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).Sync), arg0, arg1)
}

// SyncLatest mocks base method.
func (m *MockDataBrokerServiceServer) SyncLatest(arg0 *databroker.SyncLatestRequest, arg1 databroker.DataBrokerService_SyncLatestServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncLatest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncLatest indicates an expected call of SyncLatest.
func (mr *MockDataBrokerServiceServerMockRecorder) SyncLatest(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncLatest", reflect.TypeOf((*MockDataBrokerServiceServer)(nil).SyncLatest), arg0, arg1)
}

// MockDataBrokerService_SyncServer is a mock of DataBrokerService_SyncServer interface.
type MockDataBrokerService_SyncServer struct {
	ctrl     *gomock.Controller
	recorder *MockDataBrokerService_SyncServerMockRecorder
}

// MockDataBrokerService_SyncServerMockRecorder is the mock recorder for MockDataBrokerService_SyncServer.
type MockDataBrokerService_SyncServerMockRecorder struct {
	mock *MockDataBrokerService_SyncServer
}

// NewMockDataBrokerService_SyncServer creates a new mock instance.
func NewMockDataBrokerService_SyncServer(ctrl *gomock.Controller) *MockDataBrokerService_SyncServer {
	mock := &MockDataBrokerService_SyncServer{ctrl: ctrl}
	mock.recorder = &MockDataBrokerService_SyncServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataBrokerService_SyncServer) EXPECT() *MockDataBrokerService_SyncServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockDataBrokerService_SyncServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockDataBrokerService_SyncServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDataBrokerService_SyncServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockDataBrokerService_SyncServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockDataBrokerService_SyncServerMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDataBrokerService_SyncServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockDataBrokerService_SyncServer) Send(arg0 *databroker.SyncResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockDataBrokerService_SyncServerMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockDataBrokerService_SyncServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockDataBrokerService_SyncServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockDataBrokerService_SyncServerMockRecorder) SendHeader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockDataBrokerService_SyncServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockDataBrokerService_SyncServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockDataBrokerService_SyncServerMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDataBrokerService_SyncServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockDataBrokerService_SyncServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockDataBrokerService_SyncServerMockRecorder) SetHeader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockDataBrokerService_SyncServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockDataBrokerService_SyncServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockDataBrokerService_SyncServerMockRecorder) SetTrailer(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockDataBrokerService_SyncServer)(nil).SetTrailer), arg0)
}

// MockDataBrokerService_SyncLatestServer is a mock of DataBrokerService_SyncLatestServer interface.
type MockDataBrokerService_SyncLatestServer struct {
	ctrl     *gomock.Controller
	recorder *MockDataBrokerService_SyncLatestServerMockRecorder
}

// MockDataBrokerService_SyncLatestServerMockRecorder is the mock recorder for MockDataBrokerService_SyncLatestServer.
type MockDataBrokerService_SyncLatestServerMockRecorder struct {
	mock *MockDataBrokerService_SyncLatestServer
}

// NewMockDataBrokerService_SyncLatestServer creates a new mock instance.
func NewMockDataBrokerService_SyncLatestServer(ctrl *gomock.Controller) *MockDataBrokerService_SyncLatestServer {
	mock := &MockDataBrokerService_SyncLatestServer{ctrl: ctrl}
	mock.recorder = &MockDataBrokerService_SyncLatestServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataBrokerService_SyncLatestServer) EXPECT() *MockDataBrokerService_SyncLatestServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockDataBrokerService_SyncLatestServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockDataBrokerService_SyncLatestServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDataBrokerService_SyncLatestServer)(nil).Context))
}

// RecvMsg mocks base method.
func (m_2 *MockDataBrokerService_SyncLatestServer) RecvMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockDataBrokerService_SyncLatestServerMockRecorder) RecvMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDataBrokerService_SyncLatestServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockDataBrokerService_SyncLatestServer) Send(arg0 *databroker.SyncLatestResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockDataBrokerService_SyncLatestServerMockRecorder) Send(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockDataBrokerService_SyncLatestServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockDataBrokerService_SyncLatestServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockDataBrokerService_SyncLatestServerMockRecorder) SendHeader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockDataBrokerService_SyncLatestServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockDataBrokerService_SyncLatestServer) SendMsg(m any) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockDataBrokerService_SyncLatestServerMockRecorder) SendMsg(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDataBrokerService_SyncLatestServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockDataBrokerService_SyncLatestServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockDataBrokerService_SyncLatestServerMockRecorder) SetHeader(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockDataBrokerService_SyncLatestServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockDataBrokerService_SyncLatestServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockDataBrokerService_SyncLatestServerMockRecorder) SetTrailer(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockDataBrokerService_SyncLatestServer)(nil).SetTrailer), arg0)
}
