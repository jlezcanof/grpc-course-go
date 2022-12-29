// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: currency/exchangeRate/currency.proto

package exchangeRate

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

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	ExhangeRate(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (CurrencyService_ExhangeRateClient, error)
}

type currencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyServiceClient(cc grpc.ClientConnInterface) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) ExhangeRate(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (CurrencyService_ExhangeRateClient, error) {
	stream, err := c.cc.NewStream(ctx, &CurrencyService_ServiceDesc.Streams[0], "/currency.CurrencyService/ExhangeRate", opts...)
	if err != nil {
		return nil, err
	}
	x := &currencyServiceExhangeRateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CurrencyService_ExhangeRateClient interface {
	Recv() (*CurrencyResponse, error)
	grpc.ClientStream
}

type currencyServiceExhangeRateClient struct {
	grpc.ClientStream
}

func (x *currencyServiceExhangeRateClient) Recv() (*CurrencyResponse, error) {
	m := new(CurrencyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
// All implementations must embed UnimplementedCurrencyServiceServer
// for forward compatibility
type CurrencyServiceServer interface {
	ExhangeRate(*CurrencyRequest, CurrencyService_ExhangeRateServer) error
	mustEmbedUnimplementedCurrencyServiceServer()
}

// UnimplementedCurrencyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyServiceServer struct {
}

func (UnimplementedCurrencyServiceServer) ExhangeRate(*CurrencyRequest, CurrencyService_ExhangeRateServer) error {
	return status.Errorf(codes.Unimplemented, "method ExhangeRate not implemented")
}
func (UnimplementedCurrencyServiceServer) mustEmbedUnimplementedCurrencyServiceServer() {}

// UnsafeCurrencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServiceServer will
// result in compilation errors.
type UnsafeCurrencyServiceServer interface {
	mustEmbedUnimplementedCurrencyServiceServer()
}

func RegisterCurrencyServiceServer(s grpc.ServiceRegistrar, srv CurrencyServiceServer) {
	s.RegisterService(&CurrencyService_ServiceDesc, srv)
}

func _CurrencyService_ExhangeRate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CurrencyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CurrencyServiceServer).ExhangeRate(m, &currencyServiceExhangeRateServer{stream})
}

type CurrencyService_ExhangeRateServer interface {
	Send(*CurrencyResponse) error
	grpc.ServerStream
}

type currencyServiceExhangeRateServer struct {
	grpc.ServerStream
}

func (x *currencyServiceExhangeRateServer) Send(m *CurrencyResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CurrencyService_ServiceDesc is the grpc.ServiceDesc for CurrencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "currency.CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExhangeRate",
			Handler:       _CurrencyService_ExhangeRate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "currency/exchangeRate/currency.proto",
}