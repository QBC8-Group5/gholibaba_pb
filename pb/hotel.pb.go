// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: hotel.proto

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

type HotelStatus int32

const (
	HotelStatus_HotelStatus_Active  HotelStatus = 0
	HotelStatus_HotelStatus_Blocked HotelStatus = 1
)

// Enum value maps for HotelStatus.
var (
	HotelStatus_name = map[int32]string{
		0: "HotelStatus_Active",
		1: "HotelStatus_Blocked",
	}
	HotelStatus_value = map[string]int32{
		"HotelStatus_Active":  0,
		"HotelStatus_Blocked": 1,
	}
)

func (x HotelStatus) Enum() *HotelStatus {
	p := new(HotelStatus)
	*p = x
	return p
}

func (x HotelStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HotelStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_hotel_proto_enumTypes[0].Descriptor()
}

func (HotelStatus) Type() protoreflect.EnumType {
	return &file_hotel_proto_enumTypes[0]
}

func (x HotelStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HotelStatus.Descriptor instead.
func (HotelStatus) EnumDescriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{0}
}

type PricingType int32

const (
	PricingType_PricingType_Twelve_Hour     PricingType = 0
	PricingType_PricingType_TwentyFour_Hour PricingType = 1
)

// Enum value maps for PricingType.
var (
	PricingType_name = map[int32]string{
		0: "PricingType_Twelve_Hour",
		1: "PricingType_TwentyFour_Hour",
	}
	PricingType_value = map[string]int32{
		"PricingType_Twelve_Hour":     0,
		"PricingType_TwentyFour_Hour": 1,
	}
)

func (x PricingType) Enum() *PricingType {
	p := new(PricingType)
	*p = x
	return p
}

func (x PricingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PricingType) Descriptor() protoreflect.EnumDescriptor {
	return file_hotel_proto_enumTypes[1].Descriptor()
}

func (PricingType) Type() protoreflect.EnumType {
	return &file_hotel_proto_enumTypes[1]
}

func (x PricingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PricingType.Descriptor instead.
func (PricingType) EnumDescriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{1}
}

type RoomReservationStatus int32

const (
	RoomReservationStatus_RoomReservationStatus_Active     RoomReservationStatus = 0
	RoomReservationStatus_RoomReservationStatus_CheckedOut RoomReservationStatus = 1
	RoomReservationStatus_RoomReservationStatus_Cancelled  RoomReservationStatus = 2
)

// Enum value maps for RoomReservationStatus.
var (
	RoomReservationStatus_name = map[int32]string{
		0: "RoomReservationStatus_Active",
		1: "RoomReservationStatus_CheckedOut",
		2: "RoomReservationStatus_Cancelled",
	}
	RoomReservationStatus_value = map[string]int32{
		"RoomReservationStatus_Active":     0,
		"RoomReservationStatus_CheckedOut": 1,
		"RoomReservationStatus_Cancelled":  2,
	}
)

func (x RoomReservationStatus) Enum() *RoomReservationStatus {
	p := new(RoomReservationStatus)
	*p = x
	return p
}

func (x RoomReservationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoomReservationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_hotel_proto_enumTypes[2].Descriptor()
}

func (RoomReservationStatus) Type() protoreflect.EnumType {
	return &file_hotel_proto_enumTypes[2]
}

func (x RoomReservationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoomReservationStatus.Descriptor instead.
func (RoomReservationStatus) EnumDescriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{2}
}

type Hotel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HotelId     string      `protobuf:"bytes,1,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	OwnerUserId string      `protobuf:"bytes,2,opt,name=owner_user_id,json=ownerUserId,proto3" json:"owner_user_id,omitempty"`
	HotelName   string      `protobuf:"bytes,3,opt,name=hotel_name,json=hotelName,proto3" json:"hotel_name,omitempty"`
	Location    string      `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Status      HotelStatus `protobuf:"varint,5,opt,name=status,proto3,enum=HotelStatus" json:"status,omitempty"`
}

func (x *Hotel) Reset() {
	*x = Hotel{}
	mi := &file_hotel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hotel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hotel) ProtoMessage() {}

func (x *Hotel) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hotel.ProtoReflect.Descriptor instead.
func (*Hotel) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{0}
}

func (x *Hotel) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *Hotel) GetOwnerUserId() string {
	if x != nil {
		return x.OwnerUserId
	}
	return ""
}

func (x *Hotel) GetHotelName() string {
	if x != nil {
		return x.HotelName
	}
	return ""
}

func (x *Hotel) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Hotel) GetStatus() HotelStatus {
	if x != nil {
		return x.Status
	}
	return HotelStatus_HotelStatus_Active
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId      string      `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	HotelId     string      `protobuf:"bytes,2,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
	RoomType    string      `protobuf:"bytes,3,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	Capacity    int32       `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	BasePrice   float32     `protobuf:"fixed32,5,opt,name=base_price,json=basePrice,proto3" json:"base_price,omitempty"`
	PricingType PricingType `protobuf:"varint,6,opt,name=pricing_type,json=pricingType,proto3,enum=PricingType" json:"pricing_type,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	mi := &file_hotel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{1}
}

func (x *Room) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Room) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

func (x *Room) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *Room) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Room) GetBasePrice() float32 {
	if x != nil {
		return x.BasePrice
	}
	return 0
}

func (x *Room) GetPricingType() PricingType {
	if x != nil {
		return x.PricingType
	}
	return PricingType_PricingType_Twelve_Hour
}

type RoomReservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationId string                 `protobuf:"bytes,1,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
	RoomId        string                 `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartTime     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Price         float32                `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	Status        RoomReservationStatus  `protobuf:"varint,7,opt,name=status,proto3,enum=RoomReservationStatus" json:"status,omitempty"`
}

func (x *RoomReservation) Reset() {
	*x = RoomReservation{}
	mi := &file_hotel_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoomReservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomReservation) ProtoMessage() {}

func (x *RoomReservation) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomReservation.ProtoReflect.Descriptor instead.
func (*RoomReservation) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{2}
}

func (x *RoomReservation) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

func (x *RoomReservation) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *RoomReservation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RoomReservation) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *RoomReservation) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *RoomReservation) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *RoomReservation) GetStatus() RoomReservationStatus {
	if x != nil {
		return x.Status
	}
	return RoomReservationStatus_RoomReservationStatus_Active
}

type CreateHotelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hotel *Hotel `protobuf:"bytes,1,opt,name=hotel,proto3" json:"hotel,omitempty"`
}

func (x *CreateHotelRequest) Reset() {
	*x = CreateHotelRequest{}
	mi := &file_hotel_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateHotelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHotelRequest) ProtoMessage() {}

func (x *CreateHotelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHotelRequest.ProtoReflect.Descriptor instead.
func (*CreateHotelRequest) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{3}
}

func (x *CreateHotelRequest) GetHotel() *Hotel {
	if x != nil {
		return x.Hotel
	}
	return nil
}

type CreateHotelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	HotelId string `protobuf:"bytes,2,opt,name=hotel_id,json=hotelId,proto3" json:"hotel_id,omitempty"`
}

func (x *CreateHotelResponse) Reset() {
	*x = CreateHotelResponse{}
	mi := &file_hotel_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateHotelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHotelResponse) ProtoMessage() {}

func (x *CreateHotelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHotelResponse.ProtoReflect.Descriptor instead.
func (*CreateHotelResponse) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{4}
}

func (x *CreateHotelResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateHotelResponse) GetHotelId() string {
	if x != nil {
		return x.HotelId
	}
	return ""
}

type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	mi := &file_hotel_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRoomRequest) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

type CreateRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	RoomId  string `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *CreateRoomResponse) Reset() {
	*x = CreateRoomResponse{}
	mi := &file_hotel_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResponse) ProtoMessage() {}

func (x *CreateRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResponse.ProtoReflect.Descriptor instead.
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{6}
}

func (x *CreateRoomResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateRoomResponse) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type ReserveRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId    string                 `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	UserId    string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *ReserveRoomRequest) Reset() {
	*x = ReserveRoomRequest{}
	mi := &file_hotel_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReserveRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveRoomRequest) ProtoMessage() {}

func (x *ReserveRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveRoomRequest.ProtoReflect.Descriptor instead.
func (*ReserveRoomRequest) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{7}
}

func (x *ReserveRoomRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *ReserveRoomRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReserveRoomRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *ReserveRoomRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type ReserveRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success       bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ReservationId string `protobuf:"bytes,2,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
}

func (x *ReserveRoomResponse) Reset() {
	*x = ReserveRoomResponse{}
	mi := &file_hotel_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReserveRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReserveRoomResponse) ProtoMessage() {}

func (x *ReserveRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReserveRoomResponse.ProtoReflect.Descriptor instead.
func (*ReserveRoomResponse) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{8}
}

func (x *ReserveRoomResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ReserveRoomResponse) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

type ConfirmCheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservationId string `protobuf:"bytes,1,opt,name=reservation_id,json=reservationId,proto3" json:"reservation_id,omitempty"`
}

func (x *ConfirmCheckoutRequest) Reset() {
	*x = ConfirmCheckoutRequest{}
	mi := &file_hotel_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmCheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmCheckoutRequest) ProtoMessage() {}

func (x *ConfirmCheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmCheckoutRequest.ProtoReflect.Descriptor instead.
func (*ConfirmCheckoutRequest) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{9}
}

func (x *ConfirmCheckoutRequest) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

type ConfirmCheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ConfirmCheckoutResponse) Reset() {
	*x = ConfirmCheckoutResponse{}
	mi := &file_hotel_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfirmCheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmCheckoutResponse) ProtoMessage() {}

func (x *ConfirmCheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmCheckoutResponse.ProtoReflect.Descriptor instead.
func (*ConfirmCheckoutResponse) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{10}
}

func (x *ConfirmCheckoutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_hotel_proto protoreflect.FileDescriptor

var file_hotel_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7,
	0x01, 0x0a, 0x05, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x74, 0x65, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f,
	0x74, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x0c, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa2,
	0x02, 0x0a, 0x0f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x68, 0x6f, 0x74,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x48, 0x6f, 0x74, 0x65, 0x6c,
	0x52, 0x05, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x22, 0x4a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72,
	0x6f, 0x6f, 0x6d, 0x22, 0x47, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a,
	0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x3f, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x22, 0x33, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x3e, 0x0a, 0x0b, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x48, 0x6f, 0x74, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x10, 0x01, 0x2a, 0x4b, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x54, 0x77, 0x65, 0x6c, 0x76, 0x65, 0x5f, 0x48, 0x6f, 0x75, 0x72, 0x10,
	0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x54, 0x77, 0x65, 0x6e, 0x74, 0x79, 0x46, 0x6f, 0x75, 0x72, 0x5f, 0x48, 0x6f, 0x75, 0x72,
	0x10, 0x01, 0x2a, 0x84, 0x01, 0x0a, 0x15, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x1c,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x00, 0x12, 0x24,
	0x0a, 0x20, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x4f,
	0x75, 0x74, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x32, 0xff, 0x01, 0x0a, 0x0c, 0x48, 0x6f,
	0x74, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x74, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x17, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x51, 0x42, 0x43, 0x38, 0x2d, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x35, 0x2f, 0x67, 0x68, 0x6f, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61, 0x5f,
	0x70, 0x62, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hotel_proto_rawDescOnce sync.Once
	file_hotel_proto_rawDescData = file_hotel_proto_rawDesc
)

func file_hotel_proto_rawDescGZIP() []byte {
	file_hotel_proto_rawDescOnce.Do(func() {
		file_hotel_proto_rawDescData = protoimpl.X.CompressGZIP(file_hotel_proto_rawDescData)
	})
	return file_hotel_proto_rawDescData
}

var file_hotel_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_hotel_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_hotel_proto_goTypes = []any{
	(HotelStatus)(0),                // 0: HotelStatus
	(PricingType)(0),                // 1: PricingType
	(RoomReservationStatus)(0),      // 2: RoomReservationStatus
	(*Hotel)(nil),                   // 3: Hotel
	(*Room)(nil),                    // 4: Room
	(*RoomReservation)(nil),         // 5: RoomReservation
	(*CreateHotelRequest)(nil),      // 6: CreateHotelRequest
	(*CreateHotelResponse)(nil),     // 7: CreateHotelResponse
	(*CreateRoomRequest)(nil),       // 8: CreateRoomRequest
	(*CreateRoomResponse)(nil),      // 9: CreateRoomResponse
	(*ReserveRoomRequest)(nil),      // 10: ReserveRoomRequest
	(*ReserveRoomResponse)(nil),     // 11: ReserveRoomResponse
	(*ConfirmCheckoutRequest)(nil),  // 12: ConfirmCheckoutRequest
	(*ConfirmCheckoutResponse)(nil), // 13: ConfirmCheckoutResponse
	(*timestamppb.Timestamp)(nil),   // 14: google.protobuf.Timestamp
}
var file_hotel_proto_depIdxs = []int32{
	0,  // 0: Hotel.status:type_name -> HotelStatus
	1,  // 1: Room.pricing_type:type_name -> PricingType
	14, // 2: RoomReservation.start_time:type_name -> google.protobuf.Timestamp
	14, // 3: RoomReservation.end_time:type_name -> google.protobuf.Timestamp
	2,  // 4: RoomReservation.status:type_name -> RoomReservationStatus
	3,  // 5: CreateHotelRequest.hotel:type_name -> Hotel
	4,  // 6: CreateRoomRequest.room:type_name -> Room
	14, // 7: ReserveRoomRequest.start_time:type_name -> google.protobuf.Timestamp
	14, // 8: ReserveRoomRequest.end_time:type_name -> google.protobuf.Timestamp
	6,  // 9: HotelService.CreateHotel:input_type -> CreateHotelRequest
	8,  // 10: HotelService.CreateRoom:input_type -> CreateRoomRequest
	10, // 11: HotelService.ReserveRoom:input_type -> ReserveRoomRequest
	12, // 12: HotelService.ConfirmCheckout:input_type -> ConfirmCheckoutRequest
	7,  // 13: HotelService.CreateHotel:output_type -> CreateHotelResponse
	9,  // 14: HotelService.CreateRoom:output_type -> CreateRoomResponse
	11, // 15: HotelService.ReserveRoom:output_type -> ReserveRoomResponse
	13, // 16: HotelService.ConfirmCheckout:output_type -> ConfirmCheckoutResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_hotel_proto_init() }
func file_hotel_proto_init() {
	if File_hotel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hotel_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hotel_proto_goTypes,
		DependencyIndexes: file_hotel_proto_depIdxs,
		EnumInfos:         file_hotel_proto_enumTypes,
		MessageInfos:      file_hotel_proto_msgTypes,
	}.Build()
	File_hotel_proto = out.File
	file_hotel_proto_rawDesc = nil
	file_hotel_proto_goTypes = nil
	file_hotel_proto_depIdxs = nil
}
