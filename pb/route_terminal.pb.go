// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.3
// source: route_terminal.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TerminalType int32

const (
	TerminalType_TerminalType_Airport     TerminalType = 0
	TerminalType_TerminalType_RailStation TerminalType = 1
	TerminalType_TerminalType_BusTerminal TerminalType = 2
	TerminalType_TerminalType_Port        TerminalType = 3
)

// Enum value maps for TerminalType.
var (
	TerminalType_name = map[int32]string{
		0: "TerminalType_Airport",
		1: "TerminalType_RailStation",
		2: "TerminalType_BusTerminal",
		3: "TerminalType_Port",
	}
	TerminalType_value = map[string]int32{
		"TerminalType_Airport":     0,
		"TerminalType_RailStation": 1,
		"TerminalType_BusTerminal": 2,
		"TerminalType_Port":        3,
	}
)

func (x TerminalType) Enum() *TerminalType {
	p := new(TerminalType)
	*p = x
	return p
}

func (x TerminalType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TerminalType) Descriptor() protoreflect.EnumDescriptor {
	return file_route_terminal_proto_enumTypes[0].Descriptor()
}

func (TerminalType) Type() protoreflect.EnumType {
	return &file_route_terminal_proto_enumTypes[0]
}

func (x TerminalType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TerminalType.Descriptor instead.
func (TerminalType) EnumDescriptor() ([]byte, []int) {
	return file_route_terminal_proto_rawDescGZIP(), []int{0}
}

type Terminal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TerminalId   string       `protobuf:"bytes,1,opt,name=terminal_id,json=terminalId,proto3" json:"terminal_id,omitempty"`
	TerminalName string       `protobuf:"bytes,2,opt,name=terminal_name,json=terminalName,proto3" json:"terminal_name,omitempty"`
	TerminalType TerminalType `protobuf:"varint,3,opt,name=terminal_type,json=terminalType,proto3,enum=TerminalType" json:"terminal_type,omitempty"`
	Location     string       `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Terminal) Reset() {
	*x = Terminal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_terminal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Terminal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Terminal) ProtoMessage() {}

func (x *Terminal) ProtoReflect() protoreflect.Message {
	mi := &file_route_terminal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Terminal.ProtoReflect.Descriptor instead.
func (*Terminal) Descriptor() ([]byte, []int) {
	return file_route_terminal_proto_rawDescGZIP(), []int{0}
}

func (x *Terminal) GetTerminalId() string {
	if x != nil {
		return x.TerminalId
	}
	return ""
}

func (x *Terminal) GetTerminalName() string {
	if x != nil {
		return x.TerminalName
	}
	return ""
}

func (x *Terminal) GetTerminalType() TerminalType {
	if x != nil {
		return x.TerminalType
	}
	return TerminalType_TerminalType_Airport
}

func (x *Terminal) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId        string  `protobuf:"bytes,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	FromTerminalId string  `protobuf:"bytes,2,opt,name=from_terminal_id,json=fromTerminalId,proto3" json:"from_terminal_id,omitempty"`
	ToTerminalId   string  `protobuf:"bytes,3,opt,name=to_terminal_id,json=toTerminalId,proto3" json:"to_terminal_id,omitempty"`
	Distance       float32 `protobuf:"fixed32,4,opt,name=distance,proto3" json:"distance,omitempty"`
	Code           string  `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_terminal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_route_terminal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_route_terminal_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

func (x *Route) GetFromTerminalId() string {
	if x != nil {
		return x.FromTerminalId
	}
	return ""
}

func (x *Route) GetToTerminalId() string {
	if x != nil {
		return x.ToTerminalId
	}
	return ""
}

func (x *Route) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Route) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type CreateTerminalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Terminal *Terminal `protobuf:"bytes,1,opt,name=terminal,proto3" json:"terminal,omitempty"`
}

func (x *CreateTerminalRequest) Reset() {
	*x = CreateTerminalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_terminal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTerminalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTerminalRequest) ProtoMessage() {}

func (x *CreateTerminalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_terminal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTerminalRequest.ProtoReflect.Descriptor instead.
func (*CreateTerminalRequest) Descriptor() ([]byte, []int) {
	return file_route_terminal_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTerminalRequest) GetTerminal() *Terminal {
	if x != nil {
		return x.Terminal
	}
	return nil
}

type CreateTerminalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	TerminalId string `protobuf:"bytes,2,opt,name=terminal_id,json=terminalId,proto3" json:"terminal_id,omitempty"`
}

func (x *CreateTerminalResponse) Reset() {
	*x = CreateTerminalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_terminal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTerminalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTerminalResponse) ProtoMessage() {}

func (x *CreateTerminalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_terminal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTerminalResponse.ProtoReflect.Descriptor instead.
func (*CreateTerminalResponse) Descriptor() ([]byte, []int) {
	return file_route_terminal_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTerminalResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateTerminalResponse) GetTerminalId() string {
	if x != nil {
		return x.TerminalId
	}
	return ""
}

type CreateRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route *Route `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *CreateRouteRequest) Reset() {
	*x = CreateRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_terminal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRouteRequest) ProtoMessage() {}

func (x *CreateRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_terminal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRouteRequest.ProtoReflect.Descriptor instead.
func (*CreateRouteRequest) Descriptor() ([]byte, []int) {
	return file_route_terminal_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRouteRequest) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

type CreateRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	RouteId string `protobuf:"bytes,2,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
}

func (x *CreateRouteResponse) Reset() {
	*x = CreateRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_terminal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRouteResponse) ProtoMessage() {}

func (x *CreateRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_terminal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRouteResponse.ProtoReflect.Descriptor instead.
func (*CreateRouteResponse) Descriptor() ([]byte, []int) {
	return file_route_terminal_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRouteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateRouteResponse) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

type GetRouteInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId string `protobuf:"bytes,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
}

func (x *GetRouteInfoRequest) Reset() {
	*x = GetRouteInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_terminal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRouteInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteInfoRequest) ProtoMessage() {}

func (x *GetRouteInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_terminal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteInfoRequest.ProtoReflect.Descriptor instead.
func (*GetRouteInfoRequest) Descriptor() ([]byte, []int) {
	return file_route_terminal_proto_rawDescGZIP(), []int{6}
}

func (x *GetRouteInfoRequest) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

type GetRouteInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route *Route `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *GetRouteInfoResponse) Reset() {
	*x = GetRouteInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_terminal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRouteInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteInfoResponse) ProtoMessage() {}

func (x *GetRouteInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_terminal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteInfoResponse.ProtoReflect.Descriptor instead.
func (*GetRouteInfoResponse) Descriptor() ([]byte, []int) {
	return file_route_terminal_proto_rawDescGZIP(), []int{7}
}

func (x *GetRouteInfoResponse) GetRoute() *Route {
	if x != nil {
		return x.Route
	}
	return nil
}

var File_route_terminal_proto protoreflect.FileDescriptor

var file_route_terminal_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0d, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa2, 0x01, 0x0a, 0x05, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x10, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x6f, 0x5f,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x6f, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x08, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x22,
	0x53, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x4a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2a, 0x7b, 0x0a, 0x0c,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x41, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x52, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x42, 0x75, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x50, 0x6f, 0x72, 0x74, 0x10, 0x03, 0x32, 0xd0, 0x01, 0x0a, 0x14, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x14, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x51, 0x42, 0x43, 0x38, 0x2d,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x35, 0x2f, 0x67, 0x68, 0x6f, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61,
	0x5f, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_route_terminal_proto_rawDescOnce sync.Once
	file_route_terminal_proto_rawDescData = file_route_terminal_proto_rawDesc
)

func file_route_terminal_proto_rawDescGZIP() []byte {
	file_route_terminal_proto_rawDescOnce.Do(func() {
		file_route_terminal_proto_rawDescData = protoimpl.X.CompressGZIP(file_route_terminal_proto_rawDescData)
	})
	return file_route_terminal_proto_rawDescData
}

var file_route_terminal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_route_terminal_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_route_terminal_proto_goTypes = []interface{}{
	(TerminalType)(0),              // 0: TerminalType
	(*Terminal)(nil),               // 1: Terminal
	(*Route)(nil),                  // 2: Route
	(*CreateTerminalRequest)(nil),  // 3: CreateTerminalRequest
	(*CreateTerminalResponse)(nil), // 4: CreateTerminalResponse
	(*CreateRouteRequest)(nil),     // 5: CreateRouteRequest
	(*CreateRouteResponse)(nil),    // 6: CreateRouteResponse
	(*GetRouteInfoRequest)(nil),    // 7: GetRouteInfoRequest
	(*GetRouteInfoResponse)(nil),   // 8: GetRouteInfoResponse
}
var file_route_terminal_proto_depIdxs = []int32{
	0, // 0: Terminal.terminal_type:type_name -> TerminalType
	1, // 1: CreateTerminalRequest.terminal:type_name -> Terminal
	2, // 2: CreateRouteRequest.route:type_name -> Route
	2, // 3: GetRouteInfoResponse.route:type_name -> Route
	3, // 4: RouteTerminalService.CreateTerminal:input_type -> CreateTerminalRequest
	5, // 5: RouteTerminalService.CreateRoute:input_type -> CreateRouteRequest
	7, // 6: RouteTerminalService.GetRouteInfo:input_type -> GetRouteInfoRequest
	4, // 7: RouteTerminalService.CreateTerminal:output_type -> CreateTerminalResponse
	6, // 8: RouteTerminalService.CreateRoute:output_type -> CreateRouteResponse
	8, // 9: RouteTerminalService.GetRouteInfo:output_type -> GetRouteInfoResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_route_terminal_proto_init() }
func file_route_terminal_proto_init() {
	if File_route_terminal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_route_terminal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Terminal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_terminal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_terminal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTerminalRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_terminal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTerminalResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_terminal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_terminal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRouteResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_terminal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRouteInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_terminal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRouteInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_route_terminal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_route_terminal_proto_goTypes,
		DependencyIndexes: file_route_terminal_proto_depIdxs,
		EnumInfos:         file_route_terminal_proto_enumTypes,
		MessageInfos:      file_route_terminal_proto_msgTypes,
	}.Build()
	File_route_terminal_proto = out.File
	file_route_terminal_proto_rawDesc = nil
	file_route_terminal_proto_goTypes = nil
	file_route_terminal_proto_depIdxs = nil
}
