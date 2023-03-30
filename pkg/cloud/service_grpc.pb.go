// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cloud

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestKubeCloudAPIClient is the client API for TestKubeCloudAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestKubeCloudAPIClient interface {
	// Deprecated, use ExecuteAsync instead,
	// Will remove this after we fully migrate to ExecuteAsync.
	Execute(ctx context.Context, opts ...grpc.CallOption) (TestKubeCloudAPI_ExecuteClient, error)
	Send(ctx context.Context, opts ...grpc.CallOption) (TestKubeCloudAPI_SendClient, error)
	Call(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
	ExecuteAsync(ctx context.Context, opts ...grpc.CallOption) (TestKubeCloudAPI_ExecuteAsyncClient, error)
	GetLogsStream(ctx context.Context, opts ...grpc.CallOption) (TestKubeCloudAPI_GetLogsStreamClient, error)
}

type testKubeCloudAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTestKubeCloudAPIClient(cc grpc.ClientConnInterface) TestKubeCloudAPIClient {
	return &testKubeCloudAPIClient{cc}
}

func (c *testKubeCloudAPIClient) Execute(ctx context.Context, opts ...grpc.CallOption) (TestKubeCloudAPI_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestKubeCloudAPI_ServiceDesc.Streams[0], "/cloud.TestKubeCloudAPI/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &testKubeCloudAPIExecuteClient{stream}
	return x, nil
}

type TestKubeCloudAPI_ExecuteClient interface {
	Send(*ExecuteResponse) error
	Recv() (*ExecuteRequest, error)
	grpc.ClientStream
}

type testKubeCloudAPIExecuteClient struct {
	grpc.ClientStream
}

func (x *testKubeCloudAPIExecuteClient) Send(m *ExecuteResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testKubeCloudAPIExecuteClient) Recv() (*ExecuteRequest, error) {
	m := new(ExecuteRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testKubeCloudAPIClient) Send(ctx context.Context, opts ...grpc.CallOption) (TestKubeCloudAPI_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestKubeCloudAPI_ServiceDesc.Streams[1], "/cloud.TestKubeCloudAPI/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &testKubeCloudAPISendClient{stream}
	return x, nil
}

type TestKubeCloudAPI_SendClient interface {
	Send(*WebsocketData) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type testKubeCloudAPISendClient struct {
	grpc.ClientStream
}

func (x *testKubeCloudAPISendClient) Send(m *WebsocketData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testKubeCloudAPISendClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testKubeCloudAPIClient) Call(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/cloud.TestKubeCloudAPI/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testKubeCloudAPIClient) ExecuteAsync(ctx context.Context, opts ...grpc.CallOption) (TestKubeCloudAPI_ExecuteAsyncClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestKubeCloudAPI_ServiceDesc.Streams[2], "/cloud.TestKubeCloudAPI/ExecuteAsync", opts...)
	if err != nil {
		return nil, err
	}
	x := &testKubeCloudAPIExecuteAsyncClient{stream}
	return x, nil
}

type TestKubeCloudAPI_ExecuteAsyncClient interface {
	Send(*ExecuteResponse) error
	Recv() (*ExecuteRequest, error)
	grpc.ClientStream
}

type testKubeCloudAPIExecuteAsyncClient struct {
	grpc.ClientStream
}

func (x *testKubeCloudAPIExecuteAsyncClient) Send(m *ExecuteResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testKubeCloudAPIExecuteAsyncClient) Recv() (*ExecuteRequest, error) {
	m := new(ExecuteRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testKubeCloudAPIClient) GetLogsStream(ctx context.Context, opts ...grpc.CallOption) (TestKubeCloudAPI_GetLogsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestKubeCloudAPI_ServiceDesc.Streams[3], "/cloud.TestKubeCloudAPI/GetLogsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testKubeCloudAPIGetLogsStreamClient{stream}
	return x, nil
}

type TestKubeCloudAPI_GetLogsStreamClient interface {
	Send(*LogsStreamResponse) error
	Recv() (*LogsStreamRequest, error)
	grpc.ClientStream
}

type testKubeCloudAPIGetLogsStreamClient struct {
	grpc.ClientStream
}

func (x *testKubeCloudAPIGetLogsStreamClient) Send(m *LogsStreamResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testKubeCloudAPIGetLogsStreamClient) Recv() (*LogsStreamRequest, error) {
	m := new(LogsStreamRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestKubeCloudAPIServer is the server API for TestKubeCloudAPI service.
// All implementations must embed UnimplementedTestKubeCloudAPIServer
// for forward compatibility
type TestKubeCloudAPIServer interface {
	// Deprecated, use ExecuteAsync instead,
	// Will remove this after we fully migrate to ExecuteAsync.
	Execute(TestKubeCloudAPI_ExecuteServer) error
	Send(TestKubeCloudAPI_SendServer) error
	Call(context.Context, *CommandRequest) (*CommandResponse, error)
	ExecuteAsync(TestKubeCloudAPI_ExecuteAsyncServer) error
	GetLogsStream(TestKubeCloudAPI_GetLogsStreamServer) error
	mustEmbedUnimplementedTestKubeCloudAPIServer()
}

// UnimplementedTestKubeCloudAPIServer must be embedded to have forward compatible implementations.
type UnimplementedTestKubeCloudAPIServer struct {
}

func (UnimplementedTestKubeCloudAPIServer) Execute(TestKubeCloudAPI_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedTestKubeCloudAPIServer) Send(TestKubeCloudAPI_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedTestKubeCloudAPIServer) Call(context.Context, *CommandRequest) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedTestKubeCloudAPIServer) ExecuteAsync(TestKubeCloudAPI_ExecuteAsyncServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteAsync not implemented")
}
func (UnimplementedTestKubeCloudAPIServer) GetLogsStream(TestKubeCloudAPI_GetLogsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLogsStream not implemented")
}
func (UnimplementedTestKubeCloudAPIServer) mustEmbedUnimplementedTestKubeCloudAPIServer() {}

// UnsafeTestKubeCloudAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestKubeCloudAPIServer will
// result in compilation errors.
type UnsafeTestKubeCloudAPIServer interface {
	mustEmbedUnimplementedTestKubeCloudAPIServer()
}

func RegisterTestKubeCloudAPIServer(s grpc.ServiceRegistrar, srv TestKubeCloudAPIServer) {
	s.RegisterService(&TestKubeCloudAPI_ServiceDesc, srv)
}

func _TestKubeCloudAPI_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestKubeCloudAPIServer).Execute(&testKubeCloudAPIExecuteServer{stream})
}

type TestKubeCloudAPI_ExecuteServer interface {
	Send(*ExecuteRequest) error
	Recv() (*ExecuteResponse, error)
	grpc.ServerStream
}

type testKubeCloudAPIExecuteServer struct {
	grpc.ServerStream
}

func (x *testKubeCloudAPIExecuteServer) Send(m *ExecuteRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testKubeCloudAPIExecuteServer) Recv() (*ExecuteResponse, error) {
	m := new(ExecuteResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestKubeCloudAPI_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestKubeCloudAPIServer).Send(&testKubeCloudAPISendServer{stream})
}

type TestKubeCloudAPI_SendServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*WebsocketData, error)
	grpc.ServerStream
}

type testKubeCloudAPISendServer struct {
	grpc.ServerStream
}

func (x *testKubeCloudAPISendServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testKubeCloudAPISendServer) Recv() (*WebsocketData, error) {
	m := new(WebsocketData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestKubeCloudAPI_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestKubeCloudAPIServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloud.TestKubeCloudAPI/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestKubeCloudAPIServer).Call(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestKubeCloudAPI_ExecuteAsync_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestKubeCloudAPIServer).ExecuteAsync(&testKubeCloudAPIExecuteAsyncServer{stream})
}

type TestKubeCloudAPI_ExecuteAsyncServer interface {
	Send(*ExecuteRequest) error
	Recv() (*ExecuteResponse, error)
	grpc.ServerStream
}

type testKubeCloudAPIExecuteAsyncServer struct {
	grpc.ServerStream
}

func (x *testKubeCloudAPIExecuteAsyncServer) Send(m *ExecuteRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testKubeCloudAPIExecuteAsyncServer) Recv() (*ExecuteResponse, error) {
	m := new(ExecuteResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestKubeCloudAPI_GetLogsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestKubeCloudAPIServer).GetLogsStream(&testKubeCloudAPIGetLogsStreamServer{stream})
}

type TestKubeCloudAPI_GetLogsStreamServer interface {
	Send(*LogsStreamRequest) error
	Recv() (*LogsStreamResponse, error)
	grpc.ServerStream
}

type testKubeCloudAPIGetLogsStreamServer struct {
	grpc.ServerStream
}

func (x *testKubeCloudAPIGetLogsStreamServer) Send(m *LogsStreamRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testKubeCloudAPIGetLogsStreamServer) Recv() (*LogsStreamResponse, error) {
	m := new(LogsStreamResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestKubeCloudAPI_ServiceDesc is the grpc.ServiceDesc for TestKubeCloudAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestKubeCloudAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloud.TestKubeCloudAPI",
	HandlerType: (*TestKubeCloudAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _TestKubeCloudAPI_Call_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _TestKubeCloudAPI_Execute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Send",
			Handler:       _TestKubeCloudAPI_Send_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ExecuteAsync",
			Handler:       _TestKubeCloudAPI_ExecuteAsync_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetLogsStream",
			Handler:       _TestKubeCloudAPI_GetLogsStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
