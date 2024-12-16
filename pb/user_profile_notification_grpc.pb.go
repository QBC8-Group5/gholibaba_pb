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

// UserProfileServiceClient is the client API for UserProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProfileServiceClient interface {
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error)
}

type userProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProfileServiceClient(cc grpc.ClientConnInterface) UserProfileServiceClient {
	return &userProfileServiceClient{cc}
}

func (c *userProfileServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, "/UserProfileService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*UpdateProfileResponse, error) {
	out := new(UpdateProfileResponse)
	err := c.cc.Invoke(ctx, "/UserProfileService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfileServiceServer is the server API for UserProfileService service.
// All implementations must embed UnimplementedUserProfileServiceServer
// for forward compatibility
type UserProfileServiceServer interface {
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
	mustEmbedUnimplementedUserProfileServiceServer()
}

// UnimplementedUserProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserProfileServiceServer struct {
}

func (UnimplementedUserProfileServiceServer) GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserProfileServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserProfileServiceServer) mustEmbedUnimplementedUserProfileServiceServer() {}

// UnsafeUserProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProfileServiceServer will
// result in compilation errors.
type UnsafeUserProfileServiceServer interface {
	mustEmbedUnimplementedUserProfileServiceServer()
}

func RegisterUserProfileServiceServer(s *grpc.Server, srv UserProfileServiceServer) {
	s.RegisterService(&_UserProfileService_serviceDesc, srv)
}

func _UserProfileService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProfileService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServiceServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfileService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserProfileService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserProfileService",
	HandlerType: (*UserProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfile",
			Handler:    _UserProfileService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UserProfileService_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_profile_notification.proto",
}

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error)
	MarkNotificationRead(ctx context.Context, in *MarkNotificationReadRequest, opts ...grpc.CallOption) (*MarkNotificationReadResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) ListNotifications(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error) {
	out := new(ListNotificationsResponse)
	err := c.cc.Invoke(ctx, "/NotificationService/ListNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) MarkNotificationRead(ctx context.Context, in *MarkNotificationReadRequest, opts ...grpc.CallOption) (*MarkNotificationReadResponse, error) {
	out := new(MarkNotificationReadResponse)
	err := c.cc.Invoke(ctx, "/NotificationService/MarkNotificationRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	ListNotifications(context.Context, *ListNotificationsRequest) (*ListNotificationsResponse, error)
	MarkNotificationRead(context.Context, *MarkNotificationReadRequest) (*MarkNotificationReadResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) ListNotifications(context.Context, *ListNotificationsRequest) (*ListNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotifications not implemented")
}
func (UnimplementedNotificationServiceServer) MarkNotificationRead(context.Context, *MarkNotificationReadRequest) (*MarkNotificationReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkNotificationRead not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_ListNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ListNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationService/ListNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ListNotifications(ctx, req.(*ListNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_MarkNotificationRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkNotificationReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).MarkNotificationRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationService/MarkNotificationRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).MarkNotificationRead(ctx, req.(*MarkNotificationReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNotifications",
			Handler:    _NotificationService_ListNotifications_Handler,
		},
		{
			MethodName: "MarkNotificationRead",
			Handler:    _NotificationService_MarkNotificationRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_profile_notification.proto",
}