// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: vehicle_management.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VehicleType int32

const (
	VehicleType_VehicleType_Plane VehicleType = 0
	VehicleType_VehicleType_Bus   VehicleType = 1
	VehicleType_VehicleType_Train VehicleType = 2
	VehicleType_VehicleType_Ship  VehicleType = 3
)

// Enum value maps for VehicleType.
var (
	VehicleType_name = map[int32]string{
		0: "VehicleType_Plane",
		1: "VehicleType_Bus",
		2: "VehicleType_Train",
		3: "VehicleType_Ship",
	}
	VehicleType_value = map[string]int32{
		"VehicleType_Plane": 0,
		"VehicleType_Bus":   1,
		"VehicleType_Train": 2,
		"VehicleType_Ship":  3,
	}
)

func (x VehicleType) Enum() *VehicleType {
	p := new(VehicleType)
	*p = x
	return p
}

func (x VehicleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VehicleType) Descriptor() protoreflect.EnumDescriptor {
	return file_vehicle_management_proto_enumTypes[0].Descriptor()
}

func (VehicleType) Type() protoreflect.EnumType {
	return &file_vehicle_management_proto_enumTypes[0]
}

func (x VehicleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VehicleType.Descriptor instead.
func (VehicleType) EnumDescriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{0}
}

type VehicleStatus int32

const (
	VehicleStatus_VehicleStatus_Active   VehicleStatus = 0
	VehicleStatus_VehicleStatus_Inactive VehicleStatus = 1
)

// Enum value maps for VehicleStatus.
var (
	VehicleStatus_name = map[int32]string{
		0: "VehicleStatus_Active",
		1: "VehicleStatus_Inactive",
	}
	VehicleStatus_value = map[string]int32{
		"VehicleStatus_Active":   0,
		"VehicleStatus_Inactive": 1,
	}
)

func (x VehicleStatus) Enum() *VehicleStatus {
	p := new(VehicleStatus)
	*p = x
	return p
}

func (x VehicleStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VehicleStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_vehicle_management_proto_enumTypes[1].Descriptor()
}

func (VehicleStatus) Type() protoreflect.EnumType {
	return &file_vehicle_management_proto_enumTypes[1]
}

func (x VehicleStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VehicleStatus.Descriptor instead.
func (VehicleStatus) EnumDescriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{1}
}

type VehicleRequestStatus int32

const (
	VehicleRequestStatus_VehicleRequestStatus_Pending  VehicleRequestStatus = 0
	VehicleRequestStatus_VehicleRequestStatus_Matched  VehicleRequestStatus = 1
	VehicleRequestStatus_VehicleRequestStatus_Rejected VehicleRequestStatus = 2
)

// Enum value maps for VehicleRequestStatus.
var (
	VehicleRequestStatus_name = map[int32]string{
		0: "VehicleRequestStatus_Pending",
		1: "VehicleRequestStatus_Matched",
		2: "VehicleRequestStatus_Rejected",
	}
	VehicleRequestStatus_value = map[string]int32{
		"VehicleRequestStatus_Pending":  0,
		"VehicleRequestStatus_Matched":  1,
		"VehicleRequestStatus_Rejected": 2,
	}
)

func (x VehicleRequestStatus) Enum() *VehicleRequestStatus {
	p := new(VehicleRequestStatus)
	*p = x
	return p
}

func (x VehicleRequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VehicleRequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_vehicle_management_proto_enumTypes[2].Descriptor()
}

func (VehicleRequestStatus) Type() protoreflect.EnumType {
	return &file_vehicle_management_proto_enumTypes[2]
}

func (x VehicleRequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VehicleRequestStatus.Descriptor instead.
func (VehicleRequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{2}
}

type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId       string                 `protobuf:"bytes,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	OwnerUserId     string                 `protobuf:"bytes,2,opt,name=owner_user_id,json=ownerUserId,proto3" json:"owner_user_id,omitempty"`
	VehicleType     VehicleType            `protobuf:"varint,3,opt,name=vehicle_type,json=vehicleType,proto3,enum=VehicleType" json:"vehicle_type,omitempty"`
	Capacity        int32                  `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Speed           float32                `protobuf:"fixed32,5,opt,name=speed,proto3" json:"speed,omitempty"`
	ManufactureDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=manufacture_date,json=manufactureDate,proto3" json:"manufacture_date,omitempty"`
	HourlyCost      float32                `protobuf:"fixed32,7,opt,name=hourly_cost,json=hourlyCost,proto3" json:"hourly_cost,omitempty"`
	Status          VehicleStatus          `protobuf:"varint,8,opt,name=status,proto3,enum=VehicleStatus" json:"status,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	mi := &file_vehicle_management_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_management_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{0}
}

func (x *Vehicle) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

func (x *Vehicle) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *Vehicle) GetVehicleType() VehicleType {
	if x != nil {
		return x.VehicleType
	}
	return VehicleType_VehicleType_Plane
}

func (x *Vehicle) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Vehicle) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *Vehicle) GetManufactureDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ManufactureDate
	}
	return nil
}

func (x *Vehicle) GetHourlyCost() float32 {
	if x != nil {
		return x.HourlyCost
	}
	return 0
}

func (x *Vehicle) GetStatus() VehicleStatus {
	if x != nil {
		return x.Status
	}
	return VehicleStatus_VehicleStatus_Active
}

type VehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId         string                 `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	CompanyId         string                 `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	VehicleType       string                 `protobuf:"bytes,3,opt,name=vehicle_type,json=vehicleType,proto3" json:"vehicle_type,omitempty"`
	StartTerminalId   string                 `protobuf:"bytes,4,opt,name=start_terminal_id,json=startTerminalId,proto3" json:"start_terminal_id,omitempty"`
	EndTerminalId     string                 `protobuf:"bytes,5,opt,name=end_terminal_id,json=endTerminalId,proto3" json:"end_terminal_id,omitempty"`
	StartTime         *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	MinPassengerCount int32                  `protobuf:"varint,7,opt,name=min_passenger_count,json=minPassengerCount,proto3" json:"min_passenger_count,omitempty"`
	Status            string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VehicleRequest) Reset() {
	*x = VehicleRequest{}
	mi := &file_vehicle_management_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleRequest) ProtoMessage() {}

func (x *VehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_management_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleRequest.ProtoReflect.Descriptor instead.
func (*VehicleRequest) Descriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{1}
}

func (x *VehicleRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *VehicleRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *VehicleRequest) GetVehicleType() string {
	if x != nil {
		return x.VehicleType
	}
	return ""
}

func (x *VehicleRequest) GetStartTerminalId() string {
	if x != nil {
		return x.StartTerminalId
	}
	return ""
}

func (x *VehicleRequest) GetEndTerminalId() string {
	if x != nil {
		return x.EndTerminalId
	}
	return ""
}

func (x *VehicleRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *VehicleRequest) GetMinPassengerCount() int32 {
	if x != nil {
		return x.MinPassengerCount
	}
	return 0
}

func (x *VehicleRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type VehicleMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchId   string `protobuf:"bytes,1,opt,name=match_id,json=matchId,proto3" json:"match_id,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	VehicleId string `protobuf:"bytes,3,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
}

func (x *VehicleMatch) Reset() {
	*x = VehicleMatch{}
	mi := &file_vehicle_management_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VehicleMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleMatch) ProtoMessage() {}

func (x *VehicleMatch) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_management_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleMatch.ProtoReflect.Descriptor instead.
func (*VehicleMatch) Descriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{2}
}

func (x *VehicleMatch) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *VehicleMatch) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *VehicleMatch) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

type CreateVehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vehicle *Vehicle `protobuf:"bytes,1,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
}

func (x *CreateVehicleRequest) Reset() {
	*x = CreateVehicleRequest{}
	mi := &file_vehicle_management_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVehicleRequest) ProtoMessage() {}

func (x *CreateVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_management_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVehicleRequest.ProtoReflect.Descriptor instead.
func (*CreateVehicleRequest) Descriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVehicleRequest) GetVehicle() *Vehicle {
	if x != nil {
		return x.Vehicle
	}
	return nil
}

type CreateVehicleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	VehicleId string `protobuf:"bytes,2,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
}

func (x *CreateVehicleResponse) Reset() {
	*x = CreateVehicleResponse{}
	mi := &file_vehicle_management_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateVehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVehicleResponse) ProtoMessage() {}

func (x *CreateVehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_management_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVehicleResponse.ProtoReflect.Descriptor instead.
func (*CreateVehicleResponse) Descriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{4}
}

func (x *CreateVehicleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateVehicleResponse) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

type RequestVehicleForTripRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleRequest *VehicleRequest `protobuf:"bytes,1,opt,name=vehicle_request,json=vehicleRequest,proto3" json:"vehicle_request,omitempty"`
}

func (x *RequestVehicleForTripRequest) Reset() {
	*x = RequestVehicleForTripRequest{}
	mi := &file_vehicle_management_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVehicleForTripRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVehicleForTripRequest) ProtoMessage() {}

func (x *RequestVehicleForTripRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_management_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVehicleForTripRequest.ProtoReflect.Descriptor instead.
func (*RequestVehicleForTripRequest) Descriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{5}
}

func (x *RequestVehicleForTripRequest) GetVehicleRequest() *VehicleRequest {
	if x != nil {
		return x.VehicleRequest
	}
	return nil
}

type RequestVehicleForTripResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	RequestId string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *RequestVehicleForTripResponse) Reset() {
	*x = RequestVehicleForTripResponse{}
	mi := &file_vehicle_management_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestVehicleForTripResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVehicleForTripResponse) ProtoMessage() {}

func (x *RequestVehicleForTripResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_management_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVehicleForTripResponse.ProtoReflect.Descriptor instead.
func (*RequestVehicleForTripResponse) Descriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{6}
}

func (x *RequestVehicleForTripResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RequestVehicleForTripResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type MatchVehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (x *MatchVehicleRequest) Reset() {
	*x = MatchVehicleRequest{}
	mi := &file_vehicle_management_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchVehicleRequest) ProtoMessage() {}

func (x *MatchVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_management_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchVehicleRequest.ProtoReflect.Descriptor instead.
func (*MatchVehicleRequest) Descriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{7}
}

func (x *MatchVehicleRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type MatchVehicleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	VehicleId string `protobuf:"bytes,2,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
}

func (x *MatchVehicleResponse) Reset() {
	*x = MatchVehicleResponse{}
	mi := &file_vehicle_management_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchVehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchVehicleResponse) ProtoMessage() {}

func (x *MatchVehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_management_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchVehicleResponse.ProtoReflect.Descriptor instead.
func (*MatchVehicleResponse) Descriptor() ([]byte, []int) {
	return file_vehicle_management_proto_rawDescGZIP(), []int{8}
}

func (x *MatchVehicleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *MatchVehicleResponse) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

var File_vehicle_management_proto protoreflect.FileDescriptor

var file_vehicle_management_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x02, 0x0a, 0x07,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0c, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x45, 0x0a,
	0x10, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0f, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x68, 0x6f, 0x75, 0x72, 0x6c,
	0x79, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc8, 0x02,
	0x0a, 0x0e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d,
	0x69, 0x6e, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x67, 0x0a, 0x0c, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49,
	0x64, 0x22, 0x3a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x50, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x58, 0x0a, 0x1c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x38, 0x0a, 0x0f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0e, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x1d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x54, 0x72,
	0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x13, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x14, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x2a, 0x66, 0x0a, 0x0b, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f,
	0x42, 0x75, 0x73, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x5f, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x53, 0x68, 0x69, 0x70,
	0x10, 0x03, 0x2a, 0x45, 0x0a, 0x0d, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x00, 0x12, 0x1a, 0x0a,
	0x16, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x49,
	0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x01, 0x2a, 0x7d, 0x0a, 0x14, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x20, 0x0a, 0x1c, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x64, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02, 0x32, 0xef, 0x01, 0x0a, 0x18, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x69, 0x70, 0x12, 0x1d,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x46,
	0x6f, 0x72, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x46, 0x6f,
	0x72, 0x54, 0x72, 0x69, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x14, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x51, 0x42, 0x43, 0x38, 0x2d, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x35, 0x2f, 0x67, 0x68, 0x6f, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61, 0x5f, 0x70,
	0x62, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vehicle_management_proto_rawDescOnce sync.Once
	file_vehicle_management_proto_rawDescData = file_vehicle_management_proto_rawDesc
)

func file_vehicle_management_proto_rawDescGZIP() []byte {
	file_vehicle_management_proto_rawDescOnce.Do(func() {
		file_vehicle_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_vehicle_management_proto_rawDescData)
	})
	return file_vehicle_management_proto_rawDescData
}

var file_vehicle_management_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_vehicle_management_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_vehicle_management_proto_goTypes = []any{
	(VehicleType)(0),                      // 0: VehicleType
	(VehicleStatus)(0),                    // 1: VehicleStatus
	(VehicleRequestStatus)(0),             // 2: VehicleRequestStatus
	(*Vehicle)(nil),                       // 3: Vehicle
	(*VehicleRequest)(nil),                // 4: VehicleRequest
	(*VehicleMatch)(nil),                  // 5: VehicleMatch
	(*CreateVehicleRequest)(nil),          // 6: CreateVehicleRequest
	(*CreateVehicleResponse)(nil),         // 7: CreateVehicleResponse
	(*RequestVehicleForTripRequest)(nil),  // 8: RequestVehicleForTripRequest
	(*RequestVehicleForTripResponse)(nil), // 9: RequestVehicleForTripResponse
	(*MatchVehicleRequest)(nil),           // 10: MatchVehicleRequest
	(*MatchVehicleResponse)(nil),          // 11: MatchVehicleResponse
	(*timestamppb.Timestamp)(nil),         // 12: google.protobuf.Timestamp
}
var file_vehicle_management_proto_depIdxs = []int32{
	0,  // 0: Vehicle.vehicle_type:type_name -> VehicleType
	12, // 1: Vehicle.manufacture_date:type_name -> google.protobuf.Timestamp
	1,  // 2: Vehicle.status:type_name -> VehicleStatus
	12, // 3: VehicleRequest.start_time:type_name -> google.protobuf.Timestamp
	3,  // 4: CreateVehicleRequest.vehicle:type_name -> Vehicle
	4,  // 5: RequestVehicleForTripRequest.vehicle_request:type_name -> VehicleRequest
	6,  // 6: VehicleManagementService.CreateVehicle:input_type -> CreateVehicleRequest
	8,  // 7: VehicleManagementService.RequestVehicleForTrip:input_type -> RequestVehicleForTripRequest
	10, // 8: VehicleManagementService.MatchVehicle:input_type -> MatchVehicleRequest
	7,  // 9: VehicleManagementService.CreateVehicle:output_type -> CreateVehicleResponse
	9,  // 10: VehicleManagementService.RequestVehicleForTrip:output_type -> RequestVehicleForTripResponse
	11, // 11: VehicleManagementService.MatchVehicle:output_type -> MatchVehicleResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_vehicle_management_proto_init() }
func file_vehicle_management_proto_init() {
	if File_vehicle_management_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vehicle_management_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vehicle_management_proto_goTypes,
		DependencyIndexes: file_vehicle_management_proto_depIdxs,
		EnumInfos:         file_vehicle_management_proto_enumTypes,
		MessageInfos:      file_vehicle_management_proto_msgTypes,
	}.Build()
	File_vehicle_management_proto = out.File
	file_vehicle_management_proto_rawDesc = nil
	file_vehicle_management_proto_goTypes = nil
	file_vehicle_management_proto_depIdxs = nil
}
