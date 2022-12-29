// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: hello/hellopb/hellopb.proto

package hellopb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Unary
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Server streaming
	// The service return hello/greeting in different languages
	HelloManyLanguages(ctx context.Context, in *HelloManyLanguagesRequest, opts ...grpc.CallOption) (HelloService_HelloManyLanguagesClient, error)
	// Client streaming
	// Send many hellos and response with one goodbye for all people
	HellosGoodbye(ctx context.Context, opts ...grpc.CallOption) (HelloService_HellosGoodbyeClient, error)
	// bidirectional streaming
	// It will send many hellos and the the server will response a goodbye by each one of them
	Goodbye(ctx context.Context, opts ...grpc.CallOption) (HelloService_GoodbyeClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) HelloManyLanguages(ctx context.Context, in *HelloManyLanguagesRequest, opts ...grpc.CallOption) (HelloService_HelloManyLanguagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], "/hello.HelloService/HelloManyLanguages", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloManyLanguagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_HelloManyLanguagesClient interface {
	Recv() (*HelloManyLanguagesResponse, error)
	grpc.ClientStream
}

type helloServiceHelloManyLanguagesClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloManyLanguagesClient) Recv() (*HelloManyLanguagesResponse, error) {
	m := new(HelloManyLanguagesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) HellosGoodbye(ctx context.Context, opts ...grpc.CallOption) (HelloService_HellosGoodbyeClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], "/hello.HelloService/HellosGoodbye", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHellosGoodbyeClient{stream}
	return x, nil
}

type HelloService_HellosGoodbyeClient interface {
	Send(*HellosGoodbyeRequest) error
	CloseAndRecv() (*HellosGoodbyeResponse, error)
	grpc.ClientStream
}

type helloServiceHellosGoodbyeClient struct {
	grpc.ClientStream
}

func (x *helloServiceHellosGoodbyeClient) Send(m *HellosGoodbyeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceHellosGoodbyeClient) CloseAndRecv() (*HellosGoodbyeResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HellosGoodbyeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) Goodbye(ctx context.Context, opts ...grpc.CallOption) (HelloService_GoodbyeClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[2], "/hello.HelloService/Goodbye", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceGoodbyeClient{stream}
	return x, nil
}

type HelloService_GoodbyeClient interface {
	Send(*GoodbyeRequest) error
	Recv() (*GoodbyeResponse, error)
	grpc.ClientStream
}

type helloServiceGoodbyeClient struct {
	grpc.ClientStream
}

func (x *helloServiceGoodbyeClient) Send(m *GoodbyeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceGoodbyeClient) Recv() (*GoodbyeResponse, error) {
	m := new(GoodbyeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	// Unary
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Server streaming
	// The service return hello/greeting in different languages
	HelloManyLanguages(*HelloManyLanguagesRequest, HelloService_HelloManyLanguagesServer) error
	// Client streaming
	// Send many hellos and response with one goodbye for all people
	HellosGoodbye(HelloService_HellosGoodbyeServer) error
	// bidirectional streaming
	// It will send many hellos and the the server will response a goodbye by each one of them
	Goodbye(HelloService_GoodbyeServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedHelloServiceServer) HelloManyLanguages(*HelloManyLanguagesRequest, HelloService_HelloManyLanguagesServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloManyLanguages not implemented")
}
func (UnimplementedHelloServiceServer) HellosGoodbye(HelloService_HellosGoodbyeServer) error {
	return status.Errorf(codes.Unimplemented, "method HellosGoodbye not implemented")
}
func (UnimplementedHelloServiceServer) Goodbye(HelloService_GoodbyeServer) error {
	return status.Errorf(codes.Unimplemented, "method Goodbye not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_HelloManyLanguages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloManyLanguagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).HelloManyLanguages(m, &helloServiceHelloManyLanguagesServer{stream})
}

type HelloService_HelloManyLanguagesServer interface {
	Send(*HelloManyLanguagesResponse) error
	grpc.ServerStream
}

type helloServiceHelloManyLanguagesServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloManyLanguagesServer) Send(m *HelloManyLanguagesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_HellosGoodbye_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).HellosGoodbye(&helloServiceHellosGoodbyeServer{stream})
}

type HelloService_HellosGoodbyeServer interface {
	SendAndClose(*HellosGoodbyeResponse) error
	Recv() (*HellosGoodbyeRequest, error)
	grpc.ServerStream
}

type helloServiceHellosGoodbyeServer struct {
	grpc.ServerStream
}

func (x *helloServiceHellosGoodbyeServer) SendAndClose(m *HellosGoodbyeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceHellosGoodbyeServer) Recv() (*HellosGoodbyeRequest, error) {
	m := new(HellosGoodbyeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_Goodbye_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).Goodbye(&helloServiceGoodbyeServer{stream})
}

type HelloService_GoodbyeServer interface {
	Send(*GoodbyeResponse) error
	Recv() (*GoodbyeRequest, error)
	grpc.ServerStream
}

type helloServiceGoodbyeServer struct {
	grpc.ServerStream
}

func (x *helloServiceGoodbyeServer) Send(m *GoodbyeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceGoodbyeServer) Recv() (*GoodbyeRequest, error) {
	m := new(GoodbyeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloManyLanguages",
			Handler:       _HelloService_HelloManyLanguages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HellosGoodbye",
			Handler:       _HelloService_HellosGoodbye_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Goodbye",
			Handler:       _HelloService_Goodbye_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello/hellopb/hellopb.proto",
}