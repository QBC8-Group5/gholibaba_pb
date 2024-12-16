// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// TransportationCompanyServiceClient is the client API for TransportationCompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportationCompanyServiceClient interface {
	CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*CreateTripResponse, error)
	UpdateTripStatus(ctx context.Context, in *UpdateTripStatusRequest, opts ...grpc.CallOption) (*UpdateTripStatusResponse, error)
	AssignTechnicalTeam(ctx context.Context, in *AssignTechnicalTeamRequest, opts ...grpc.CallOption) (*AssignTechnicalTeamResponse, error)
	ListCompanyTrips(ctx context.Context, in *ListCompanyTripsRequest, opts ...grpc.CallOption) (*ListCompanyTripsResponse, error)
}

type transportationCompanyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportationCompanyServiceClient(cc grpc.ClientConnInterface) TransportationCompanyServiceClient {
	return &transportationCompanyServiceClient{cc}
}

func (c *transportationCompanyServiceClient) CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*CreateTripResponse, error) {
	out := new(CreateTripResponse)
	err := c.cc.Invoke(ctx, "/TransportationCompanyService/CreateTrip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportationCompanyServiceClient) UpdateTripStatus(ctx context.Context, in *UpdateTripStatusRequest, opts ...grpc.CallOption) (*UpdateTripStatusResponse, error) {
	out := new(UpdateTripStatusResponse)
	err := c.cc.Invoke(ctx, "/TransportationCompanyService/UpdateTripStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportationCompanyServiceClient) AssignTechnicalTeam(ctx context.Context, in *AssignTechnicalTeamRequest, opts ...grpc.CallOption) (*AssignTechnicalTeamResponse, error) {
	out := new(AssignTechnicalTeamResponse)
	err := c.cc.Invoke(ctx, "/TransportationCompanyService/AssignTechnicalTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transportationCompanyServiceClient) ListCompanyTrips(ctx context.Context, in *ListCompanyTripsRequest, opts ...grpc.CallOption) (*ListCompanyTripsResponse, error) {
	out := new(ListCompanyTripsResponse)
	err := c.cc.Invoke(ctx, "/TransportationCompanyService/ListCompanyTrips", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransportationCompanyServiceServer is the server API for TransportationCompanyService service.
// All implementations must embed UnimplementedTransportationCompanyServiceServer
// for forward compatibility
type TransportationCompanyServiceServer interface {
	CreateTrip(context.Context, *CreateTripRequest) (*CreateTripResponse, error)
	UpdateTripStatus(context.Context, *UpdateTripStatusRequest) (*UpdateTripStatusResponse, error)
	AssignTechnicalTeam(context.Context, *AssignTechnicalTeamRequest) (*AssignTechnicalTeamResponse, error)
	ListCompanyTrips(context.Context, *ListCompanyTripsRequest) (*ListCompanyTripsResponse, error)
	mustEmbedUnimplementedTransportationCompanyServiceServer()
}

// UnimplementedTransportationCompanyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransportationCompanyServiceServer struct {
}

func (UnimplementedTransportationCompanyServiceServer) CreateTrip(context.Context, *CreateTripRequest) (*CreateTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrip not implemented")
}
func (UnimplementedTransportationCompanyServiceServer) UpdateTripStatus(context.Context, *UpdateTripStatusRequest) (*UpdateTripStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTripStatus not implemented")
}
func (UnimplementedTransportationCompanyServiceServer) AssignTechnicalTeam(context.Context, *AssignTechnicalTeamRequest) (*AssignTechnicalTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignTechnicalTeam not implemented")
}
func (UnimplementedTransportationCompanyServiceServer) ListCompanyTrips(context.Context, *ListCompanyTripsRequest) (*ListCompanyTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompanyTrips not implemented")
}
func (UnimplementedTransportationCompanyServiceServer) mustEmbedUnimplementedTransportationCompanyServiceServer() {
}

// UnsafeTransportationCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportationCompanyServiceServer will
// result in compilation errors.
type UnsafeTransportationCompanyServiceServer interface {
	mustEmbedUnimplementedTransportationCompanyServiceServer()
}

func RegisterTransportationCompanyServiceServer(s *grpc.Server, srv TransportationCompanyServiceServer) {
	s.RegisterService(&_TransportationCompanyService_serviceDesc, srv)
}

func _TransportationCompanyService_CreateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportationCompanyServiceServer).CreateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransportationCompanyService/CreateTrip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportationCompanyServiceServer).CreateTrip(ctx, req.(*CreateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportationCompanyService_UpdateTripStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTripStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportationCompanyServiceServer).UpdateTripStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransportationCompanyService/UpdateTripStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportationCompanyServiceServer).UpdateTripStatus(ctx, req.(*UpdateTripStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportationCompanyService_AssignTechnicalTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignTechnicalTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportationCompanyServiceServer).AssignTechnicalTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransportationCompanyService/AssignTechnicalTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportationCompanyServiceServer).AssignTechnicalTeam(ctx, req.(*AssignTechnicalTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransportationCompanyService_ListCompanyTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCompanyTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransportationCompanyServiceServer).ListCompanyTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransportationCompanyService/ListCompanyTrips",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransportationCompanyServiceServer).ListCompanyTrips(ctx, req.(*ListCompanyTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransportationCompanyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TransportationCompanyService",
	HandlerType: (*TransportationCompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrip",
			Handler:    _TransportationCompanyService_CreateTrip_Handler,
		},
		{
			MethodName: "UpdateTripStatus",
			Handler:    _TransportationCompanyService_UpdateTripStatus_Handler,
		},
		{
			MethodName: "AssignTechnicalTeam",
			Handler:    _TransportationCompanyService_AssignTechnicalTeam_Handler,
		},
		{
			MethodName: "ListCompanyTrips",
			Handler:    _TransportationCompanyService_ListCompanyTrips_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transportation_company.proto",
}
