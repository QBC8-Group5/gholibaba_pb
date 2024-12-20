// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: travel_agency.proto

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
	TravelAgencyService_CreateTour_FullMethodName        = "/TravelAgencyService/CreateTour"
	TravelAgencyService_CreateTourPackage_FullMethodName = "/TravelAgencyService/CreateTourPackage"
	TravelAgencyService_ListTours_FullMethodName         = "/TravelAgencyService/ListTours"
)

// TravelAgencyServiceClient is the client API for TravelAgencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TravelAgencyServiceClient interface {
	CreateTour(ctx context.Context, in *CreateTourRequest, opts ...grpc.CallOption) (*CreateTourResponse, error)
	CreateTourPackage(ctx context.Context, in *CreateTourPackageRequest, opts ...grpc.CallOption) (*CreateTourPackageResponse, error)
	ListTours(ctx context.Context, in *ListToursRequest, opts ...grpc.CallOption) (*ListToursResponse, error)
}

type travelAgencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTravelAgencyServiceClient(cc grpc.ClientConnInterface) TravelAgencyServiceClient {
	return &travelAgencyServiceClient{cc}
}

func (c *travelAgencyServiceClient) CreateTour(ctx context.Context, in *CreateTourRequest, opts ...grpc.CallOption) (*CreateTourResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTourResponse)
	err := c.cc.Invoke(ctx, TravelAgencyService_CreateTour_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelAgencyServiceClient) CreateTourPackage(ctx context.Context, in *CreateTourPackageRequest, opts ...grpc.CallOption) (*CreateTourPackageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTourPackageResponse)
	err := c.cc.Invoke(ctx, TravelAgencyService_CreateTourPackage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *travelAgencyServiceClient) ListTours(ctx context.Context, in *ListToursRequest, opts ...grpc.CallOption) (*ListToursResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListToursResponse)
	err := c.cc.Invoke(ctx, TravelAgencyService_ListTours_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TravelAgencyServiceServer is the server API for TravelAgencyService service.
// All implementations must embed UnimplementedTravelAgencyServiceServer
// for forward compatibility.
type TravelAgencyServiceServer interface {
	CreateTour(context.Context, *CreateTourRequest) (*CreateTourResponse, error)
	CreateTourPackage(context.Context, *CreateTourPackageRequest) (*CreateTourPackageResponse, error)
	ListTours(context.Context, *ListToursRequest) (*ListToursResponse, error)
	mustEmbedUnimplementedTravelAgencyServiceServer()
}

// UnimplementedTravelAgencyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTravelAgencyServiceServer struct{}

func (UnimplementedTravelAgencyServiceServer) CreateTour(context.Context, *CreateTourRequest) (*CreateTourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTour not implemented")
}
func (UnimplementedTravelAgencyServiceServer) CreateTourPackage(context.Context, *CreateTourPackageRequest) (*CreateTourPackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTourPackage not implemented")
}
func (UnimplementedTravelAgencyServiceServer) ListTours(context.Context, *ListToursRequest) (*ListToursResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTours not implemented")
}
func (UnimplementedTravelAgencyServiceServer) mustEmbedUnimplementedTravelAgencyServiceServer() {}
func (UnimplementedTravelAgencyServiceServer) testEmbeddedByValue()                             {}

// UnsafeTravelAgencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TravelAgencyServiceServer will
// result in compilation errors.
type UnsafeTravelAgencyServiceServer interface {
	mustEmbedUnimplementedTravelAgencyServiceServer()
}

func RegisterTravelAgencyServiceServer(s grpc.ServiceRegistrar, srv TravelAgencyServiceServer) {
	// If the following call pancis, it indicates UnimplementedTravelAgencyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TravelAgencyService_ServiceDesc, srv)
}

func _TravelAgencyService_CreateTour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAgencyServiceServer).CreateTour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TravelAgencyService_CreateTour_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAgencyServiceServer).CreateTour(ctx, req.(*CreateTourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelAgencyService_CreateTourPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTourPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAgencyServiceServer).CreateTourPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TravelAgencyService_CreateTourPackage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAgencyServiceServer).CreateTourPackage(ctx, req.(*CreateTourPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TravelAgencyService_ListTours_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListToursRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelAgencyServiceServer).ListTours(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TravelAgencyService_ListTours_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelAgencyServiceServer).ListTours(ctx, req.(*ListToursRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TravelAgencyService_ServiceDesc is the grpc.ServiceDesc for TravelAgencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TravelAgencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TravelAgencyService",
	HandlerType: (*TravelAgencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTour",
			Handler:    _TravelAgencyService_CreateTour_Handler,
		},
		{
			MethodName: "CreateTourPackage",
			Handler:    _TravelAgencyService_CreateTourPackage_Handler,
		},
		{
			MethodName: "ListTours",
			Handler:    _TravelAgencyService_ListTours_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "travel_agency.proto",
}
