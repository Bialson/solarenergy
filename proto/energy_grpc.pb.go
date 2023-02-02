// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/energy.proto

package proto

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

// SolarServiceClient is the client API for SolarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SolarServiceClient interface {
	GetSolarEnergyFromHomesByParams(ctx context.Context, in *PowerConsumptionRequest, opts ...grpc.CallOption) (SolarService_GetSolarEnergyFromHomesByParamsClient, error)
}

type solarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSolarServiceClient(cc grpc.ClientConnInterface) SolarServiceClient {
	return &solarServiceClient{cc}
}

func (c *solarServiceClient) GetSolarEnergyFromHomesByParams(ctx context.Context, in *PowerConsumptionRequest, opts ...grpc.CallOption) (SolarService_GetSolarEnergyFromHomesByParamsClient, error) {
	stream, err := c.cc.NewStream(ctx, &SolarService_ServiceDesc.Streams[0], "/solarservice.SolarService/GetSolarEnergyFromHomesByParams", opts...)
	if err != nil {
		return nil, err
	}
	x := &solarServiceGetSolarEnergyFromHomesByParamsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SolarService_GetSolarEnergyFromHomesByParamsClient interface {
	Recv() (*PowerFromHomes, error)
	grpc.ClientStream
}

type solarServiceGetSolarEnergyFromHomesByParamsClient struct {
	grpc.ClientStream
}

func (x *solarServiceGetSolarEnergyFromHomesByParamsClient) Recv() (*PowerFromHomes, error) {
	m := new(PowerFromHomes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SolarServiceServer is the server API for SolarService service.
// All implementations must embed UnimplementedSolarServiceServer
// for forward compatibility
type SolarServiceServer interface {
	GetSolarEnergyFromHomesByParams(*PowerConsumptionRequest, SolarService_GetSolarEnergyFromHomesByParamsServer) error
	mustEmbedUnimplementedSolarServiceServer()
}

// UnimplementedSolarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSolarServiceServer struct {
}

func (UnimplementedSolarServiceServer) GetSolarEnergyFromHomesByParams(*PowerConsumptionRequest, SolarService_GetSolarEnergyFromHomesByParamsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSolarEnergyFromHomesByParams not implemented")
}
func (UnimplementedSolarServiceServer) mustEmbedUnimplementedSolarServiceServer() {}

// UnsafeSolarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SolarServiceServer will
// result in compilation errors.
type UnsafeSolarServiceServer interface {
	mustEmbedUnimplementedSolarServiceServer()
}

func RegisterSolarServiceServer(s grpc.ServiceRegistrar, srv SolarServiceServer) {
	s.RegisterService(&SolarService_ServiceDesc, srv)
}

func _SolarService_GetSolarEnergyFromHomesByParams_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PowerConsumptionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SolarServiceServer).GetSolarEnergyFromHomesByParams(m, &solarServiceGetSolarEnergyFromHomesByParamsServer{stream})
}

type SolarService_GetSolarEnergyFromHomesByParamsServer interface {
	Send(*PowerFromHomes) error
	grpc.ServerStream
}

type solarServiceGetSolarEnergyFromHomesByParamsServer struct {
	grpc.ServerStream
}

func (x *solarServiceGetSolarEnergyFromHomesByParamsServer) Send(m *PowerFromHomes) error {
	return x.ServerStream.SendMsg(m)
}

// SolarService_ServiceDesc is the grpc.ServiceDesc for SolarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SolarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "solarservice.SolarService",
	HandlerType: (*SolarServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSolarEnergyFromHomesByParams",
			Handler:       _SolarService_GetSolarEnergyFromHomesByParams_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/energy.proto",
}
