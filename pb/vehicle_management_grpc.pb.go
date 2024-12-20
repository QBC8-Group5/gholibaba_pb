// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: vehicle_management.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	VehicleManagementService_CreateVehicle_FullMethodName         = "/VehicleManagementService/CreateVehicle"
	VehicleManagementService_RequestVehicleForTrip_FullMethodName = "/VehicleManagementService/RequestVehicleForTrip"
	VehicleManagementService_MatchVehicle_FullMethodName          = "/VehicleManagementService/MatchVehicle"
)

// VehicleManagementServiceClient is the client API for VehicleManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleManagementServiceClient interface {
	CreateVehicle(ctx context.Context, in *CreateVehicleRequest, opts ...grpc.CallOption) (*CreateVehicleResponse, error)
	RequestVehicleForTrip(ctx context.Context, in *RequestVehicleForTripRequest, opts ...grpc.CallOption) (*RequestVehicleForTripResponse, error)
	MatchVehicle(ctx context.Context, in *MatchVehicleRequest, opts ...grpc.CallOption) (*MatchVehicleResponse, error)
}

type vehicleManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleManagementServiceClient(cc grpc.ClientConnInterface) VehicleManagementServiceClient {
	return &vehicleManagementServiceClient{cc}
}

func (c *vehicleManagementServiceClient) CreateVehicle(ctx context.Context, in *CreateVehicleRequest, opts ...grpc.CallOption) (*CreateVehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateVehicleResponse)
	err := c.cc.Invoke(ctx, VehicleManagementService_CreateVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleManagementServiceClient) RequestVehicleForTrip(ctx context.Context, in *RequestVehicleForTripRequest, opts ...grpc.CallOption) (*RequestVehicleForTripResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestVehicleForTripResponse)
	err := c.cc.Invoke(ctx, VehicleManagementService_RequestVehicleForTrip_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleManagementServiceClient) MatchVehicle(ctx context.Context, in *MatchVehicleRequest, opts ...grpc.CallOption) (*MatchVehicleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MatchVehicleResponse)
	err := c.cc.Invoke(ctx, VehicleManagementService_MatchVehicle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleManagementServiceServer is the server API for VehicleManagementService service.
// All implementations must embed UnimplementedVehicleManagementServiceServer
// for forward compatibility.
type VehicleManagementServiceServer interface {
	CreateVehicle(context.Context, *CreateVehicleRequest) (*CreateVehicleResponse, error)
	RequestVehicleForTrip(context.Context, *RequestVehicleForTripRequest) (*RequestVehicleForTripResponse, error)
	MatchVehicle(context.Context, *MatchVehicleRequest) (*MatchVehicleResponse, error)
	mustEmbedUnimplementedVehicleManagementServiceServer()
}

// UnimplementedVehicleManagementServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVehicleManagementServiceServer struct{}

func (UnimplementedVehicleManagementServiceServer) CreateVehicle(context.Context, *CreateVehicleRequest) (*CreateVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVehicle not implemented")
}
func (UnimplementedVehicleManagementServiceServer) RequestVehicleForTrip(context.Context, *RequestVehicleForTripRequest) (*RequestVehicleForTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVehicleForTrip not implemented")
}
func (UnimplementedVehicleManagementServiceServer) MatchVehicle(context.Context, *MatchVehicleRequest) (*MatchVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MatchVehicle not implemented")
}
func (UnimplementedVehicleManagementServiceServer) mustEmbedUnimplementedVehicleManagementServiceServer() {
}
func (UnimplementedVehicleManagementServiceServer) testEmbeddedByValue() {}

// UnsafeVehicleManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleManagementServiceServer will
// result in compilation errors.
type UnsafeVehicleManagementServiceServer interface {
	mustEmbedUnimplementedVehicleManagementServiceServer()
}

func RegisterVehicleManagementServiceServer(s grpc.ServiceRegistrar, srv VehicleManagementServiceServer) {
	// If the following call pancis, it indicates UnimplementedVehicleManagementServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VehicleManagementService_ServiceDesc, srv)
}

func _VehicleManagementService_CreateVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleManagementServiceServer).CreateVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleManagementService_CreateVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleManagementServiceServer).CreateVehicle(ctx, req.(*CreateVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleManagementService_RequestVehicleForTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVehicleForTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleManagementServiceServer).RequestVehicleForTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleManagementService_RequestVehicleForTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleManagementServiceServer).RequestVehicleForTrip(ctx, req.(*RequestVehicleForTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleManagementService_MatchVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleManagementServiceServer).MatchVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VehicleManagementService_MatchVehicle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleManagementServiceServer).MatchVehicle(ctx, req.(*MatchVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VehicleManagementService_ServiceDesc is the grpc.ServiceDesc for VehicleManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VehicleManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VehicleManagementService",
	HandlerType: (*VehicleManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVehicle",
			Handler:    _VehicleManagementService_CreateVehicle_Handler,
		},
		{
			MethodName: "RequestVehicleForTrip",
			Handler:    _VehicleManagementService_RequestVehicleForTrip_Handler,
		},
		{
			MethodName: "MatchVehicle",
			Handler:    _VehicleManagementService_MatchVehicle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vehicle_management.proto",
}
