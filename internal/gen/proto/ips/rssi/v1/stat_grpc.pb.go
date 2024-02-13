// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: ips/rssi/v1/stat.proto

package rssiv1

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

// StatCollectionServiceClient is the client API for StatCollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatCollectionServiceClient interface {
	CollectData(ctx context.Context, in *CollectDataRequest, opts ...grpc.CallOption) (*CollectDataResponse, error)
	GetStatData(ctx context.Context, in *GetStatDataRequest, opts ...grpc.CallOption) (*GetStatDataResponse, error)
	ReDoDataProcessing(ctx context.Context, in *ReDoDataProcessingRequest, opts ...grpc.CallOption) (*ReDoDataProcessingResponse, error)
}

type statCollectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatCollectionServiceClient(cc grpc.ClientConnInterface) StatCollectionServiceClient {
	return &statCollectionServiceClient{cc}
}

func (c *statCollectionServiceClient) CollectData(ctx context.Context, in *CollectDataRequest, opts ...grpc.CallOption) (*CollectDataResponse, error) {
	out := new(CollectDataResponse)
	err := c.cc.Invoke(ctx, "/ips.rssi.v1.StatCollectionService/CollectData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statCollectionServiceClient) GetStatData(ctx context.Context, in *GetStatDataRequest, opts ...grpc.CallOption) (*GetStatDataResponse, error) {
	out := new(GetStatDataResponse)
	err := c.cc.Invoke(ctx, "/ips.rssi.v1.StatCollectionService/GetStatData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statCollectionServiceClient) ReDoDataProcessing(ctx context.Context, in *ReDoDataProcessingRequest, opts ...grpc.CallOption) (*ReDoDataProcessingResponse, error) {
	out := new(ReDoDataProcessingResponse)
	err := c.cc.Invoke(ctx, "/ips.rssi.v1.StatCollectionService/ReDoDataProcessing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatCollectionServiceServer is the server API for StatCollectionService service.
// All implementations must embed UnimplementedStatCollectionServiceServer
// for forward compatibility
type StatCollectionServiceServer interface {
	CollectData(context.Context, *CollectDataRequest) (*CollectDataResponse, error)
	GetStatData(context.Context, *GetStatDataRequest) (*GetStatDataResponse, error)
	ReDoDataProcessing(context.Context, *ReDoDataProcessingRequest) (*ReDoDataProcessingResponse, error)
	mustEmbedUnimplementedStatCollectionServiceServer()
}

// UnimplementedStatCollectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatCollectionServiceServer struct {
}

func (UnimplementedStatCollectionServiceServer) CollectData(context.Context, *CollectDataRequest) (*CollectDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectData not implemented")
}
func (UnimplementedStatCollectionServiceServer) GetStatData(context.Context, *GetStatDataRequest) (*GetStatDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatData not implemented")
}
func (UnimplementedStatCollectionServiceServer) ReDoDataProcessing(context.Context, *ReDoDataProcessingRequest) (*ReDoDataProcessingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReDoDataProcessing not implemented")
}
func (UnimplementedStatCollectionServiceServer) mustEmbedUnimplementedStatCollectionServiceServer() {}

// UnsafeStatCollectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatCollectionServiceServer will
// result in compilation errors.
type UnsafeStatCollectionServiceServer interface {
	mustEmbedUnimplementedStatCollectionServiceServer()
}

func RegisterStatCollectionServiceServer(s grpc.ServiceRegistrar, srv StatCollectionServiceServer) {
	s.RegisterService(&StatCollectionService_ServiceDesc, srv)
}

func _StatCollectionService_CollectData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatCollectionServiceServer).CollectData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ips.rssi.v1.StatCollectionService/CollectData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatCollectionServiceServer).CollectData(ctx, req.(*CollectDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatCollectionService_GetStatData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatCollectionServiceServer).GetStatData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ips.rssi.v1.StatCollectionService/GetStatData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatCollectionServiceServer).GetStatData(ctx, req.(*GetStatDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatCollectionService_ReDoDataProcessing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReDoDataProcessingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatCollectionServiceServer).ReDoDataProcessing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ips.rssi.v1.StatCollectionService/ReDoDataProcessing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatCollectionServiceServer).ReDoDataProcessing(ctx, req.(*ReDoDataProcessingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StatCollectionService_ServiceDesc is the grpc.ServiceDesc for StatCollectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatCollectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ips.rssi.v1.StatCollectionService",
	HandlerType: (*StatCollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectData",
			Handler:    _StatCollectionService_CollectData_Handler,
		},
		{
			MethodName: "GetStatData",
			Handler:    _StatCollectionService_GetStatData_Handler,
		},
		{
			MethodName: "ReDoDataProcessing",
			Handler:    _StatCollectionService_ReDoDataProcessing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ips/rssi/v1/stat.proto",
}
