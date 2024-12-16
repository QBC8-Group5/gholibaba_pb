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

// RouteTerminalServiceClient is the client API for RouteTerminalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteTerminalServiceClient interface {
	CreateTerminal(ctx context.Context, in *CreateTerminalRequest, opts ...grpc.CallOption) (*CreateTerminalResponse, error)
	CreateRoute(ctx context.Context, in *CreateRouteRequest, opts ...grpc.CallOption) (*CreateRouteResponse, error)
	GetRouteInfo(ctx context.Context, in *GetRouteInfoRequest, opts ...grpc.CallOption) (*GetRouteInfoResponse, error)
}

type routeTerminalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteTerminalServiceClient(cc grpc.ClientConnInterface) RouteTerminalServiceClient {
	return &routeTerminalServiceClient{cc}
}

func (c *routeTerminalServiceClient) CreateTerminal(ctx context.Context, in *CreateTerminalRequest, opts ...grpc.CallOption) (*CreateTerminalResponse, error) {
	out := new(CreateTerminalResponse)
	err := c.cc.Invoke(ctx, "/RouteTerminalService/CreateTerminal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTerminalServiceClient) CreateRoute(ctx context.Context, in *CreateRouteRequest, opts ...grpc.CallOption) (*CreateRouteResponse, error) {
	out := new(CreateRouteResponse)
	err := c.cc.Invoke(ctx, "/RouteTerminalService/CreateRoute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeTerminalServiceClient) GetRouteInfo(ctx context.Context, in *GetRouteInfoRequest, opts ...grpc.CallOption) (*GetRouteInfoResponse, error) {
	out := new(GetRouteInfoResponse)
	err := c.cc.Invoke(ctx, "/RouteTerminalService/GetRouteInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteTerminalServiceServer is the server API for RouteTerminalService service.
// All implementations must embed UnimplementedRouteTerminalServiceServer
// for forward compatibility
type RouteTerminalServiceServer interface {
	CreateTerminal(context.Context, *CreateTerminalRequest) (*CreateTerminalResponse, error)
	CreateRoute(context.Context, *CreateRouteRequest) (*CreateRouteResponse, error)
	GetRouteInfo(context.Context, *GetRouteInfoRequest) (*GetRouteInfoResponse, error)
	mustEmbedUnimplementedRouteTerminalServiceServer()
}

// UnimplementedRouteTerminalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRouteTerminalServiceServer struct {
}

func (UnimplementedRouteTerminalServiceServer) CreateTerminal(context.Context, *CreateTerminalRequest) (*CreateTerminalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTerminal not implemented")
}
func (UnimplementedRouteTerminalServiceServer) CreateRoute(context.Context, *CreateRouteRequest) (*CreateRouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoute not implemented")
}
func (UnimplementedRouteTerminalServiceServer) GetRouteInfo(context.Context, *GetRouteInfoRequest) (*GetRouteInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouteInfo not implemented")
}
func (UnimplementedRouteTerminalServiceServer) mustEmbedUnimplementedRouteTerminalServiceServer() {}

// UnsafeRouteTerminalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteTerminalServiceServer will
// result in compilation errors.
type UnsafeRouteTerminalServiceServer interface {
	mustEmbedUnimplementedRouteTerminalServiceServer()
}

func RegisterRouteTerminalServiceServer(s *grpc.Server, srv RouteTerminalServiceServer) {
	s.RegisterService(&_RouteTerminalService_serviceDesc, srv)
}

func _RouteTerminalService_CreateTerminal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTerminalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTerminalServiceServer).CreateTerminal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteTerminalService/CreateTerminal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTerminalServiceServer).CreateTerminal(ctx, req.(*CreateTerminalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTerminalService_CreateRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTerminalServiceServer).CreateRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteTerminalService/CreateRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTerminalServiceServer).CreateRoute(ctx, req.(*CreateRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteTerminalService_GetRouteInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRouteInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteTerminalServiceServer).GetRouteInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RouteTerminalService/GetRouteInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteTerminalServiceServer).GetRouteInfo(ctx, req.(*GetRouteInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteTerminalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RouteTerminalService",
	HandlerType: (*RouteTerminalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTerminal",
			Handler:    _RouteTerminalService_CreateTerminal_Handler,
		},
		{
			MethodName: "CreateRoute",
			Handler:    _RouteTerminalService_CreateRoute_Handler,
		},
		{
			MethodName: "GetRouteInfo",
			Handler:    _RouteTerminalService_GetRouteInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route_terminal.proto",
}
