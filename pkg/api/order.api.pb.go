// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: api/order.api.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OrderItemStatus int32

const (
	OrderItemStatus_StatusUnknown OrderItemStatus = 0
	OrderItemStatus_StatusCreated OrderItemStatus = 1
	OrderItemStatus_StatusDeleted OrderItemStatus = 2
)

// Enum value maps for OrderItemStatus.
var (
	OrderItemStatus_name = map[int32]string{
		0: "StatusUnknown",
		1: "StatusCreated",
		2: "StatusDeleted",
	}
	OrderItemStatus_value = map[string]int32{
		"StatusUnknown": 0,
		"StatusCreated": 1,
		"StatusDeleted": 2,
	}
)

func (x OrderItemStatus) Enum() *OrderItemStatus {
	p := new(OrderItemStatus)
	*p = x
	return p
}

func (x OrderItemStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderItemStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_order_api_proto_enumTypes[0].Descriptor()
}

func (OrderItemStatus) Type() protoreflect.EnumType {
	return &file_api_order_api_proto_enumTypes[0]
}

func (x OrderItemStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderItemStatus.Descriptor instead.
func (OrderItemStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{0}
}

type Direction int32

const (
	Direction_ASC  Direction = 0
	Direction_DESC Direction = 1
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "ASC",
		1: "DESC",
	}
	Direction_value = map[string]int32{
		"ASC":  0,
		"DESC": 1,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_api_order_api_proto_enumTypes[1].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_api_order_api_proto_enumTypes[1]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{1}
}

type OrderItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter *OrderItemFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *OrderItemRequest) Reset() {
	*x = OrderItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemRequest) ProtoMessage() {}

func (x *OrderItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemRequest.ProtoReflect.Descriptor instead.
func (*OrderItemRequest) Descriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{0}
}

func (x *OrderItemRequest) GetFilter() *OrderItemFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type OrderItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *OrderItem `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OrderItemResponse) Reset() {
	*x = OrderItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemResponse) ProtoMessage() {}

func (x *OrderItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemResponse.ProtoReflect.Descriptor instead.
func (*OrderItemResponse) Descriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItemResponse) GetValue() *OrderItem {
	if x != nil {
		return x.Value
	}
	return nil
}

type OrderItemListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter     *OrderItemFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Orders     []*Order         `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders,omitempty"`
	Pagination *Pagination      `protobuf:"bytes,3,opt,name=pagination,proto3,oneof" json:"pagination,omitempty"`
}

func (x *OrderItemListRequest) Reset() {
	*x = OrderItemListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemListRequest) ProtoMessage() {}

func (x *OrderItemListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemListRequest.ProtoReflect.Descriptor instead.
func (*OrderItemListRequest) Descriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItemListRequest) GetFilter() *OrderItemFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *OrderItemListRequest) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *OrderItemListRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type OrderItemListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*OrderItem `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *OrderItemListResponse) Reset() {
	*x = OrderItemListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemListResponse) ProtoMessage() {}

func (x *OrderItemListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemListResponse.ProtoReflect.Descriptor instead.
func (*OrderItemListResponse) Descriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{3}
}

func (x *OrderItemListResponse) GetValue() []*OrderItem {
	if x != nil {
		return x.Value
	}
	return nil
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Enum
	Status OrderItemStatus `protobuf:"varint,2,opt,name=status,proto3,enum=order.api.OrderItemStatus" json:"status,omitempty"`
	// Record creation date
	TsCreate *timestamp.Timestamp `protobuf:"bytes,3,opt,name=ts_create,json=tsCreate,proto3" json:"ts_create,omitempty"`
	// Record modification date
	TsModify *timestamp.Timestamp `protobuf:"bytes,4,opt,name=ts_modify,json=tsModify,proto3" json:"ts_modify,omitempty"`
	// UUID
	UserId string `protobuf:"bytes,11,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Record name
	Name string `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	// Record description
	Description string `protobuf:"bytes,13,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{4}
}

func (x *OrderItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderItem) GetStatus() OrderItemStatus {
	if x != nil {
		return x.Status
	}
	return OrderItemStatus_StatusUnknown
}

func (x *OrderItem) GetTsCreate() *timestamp.Timestamp {
	if x != nil {
		return x.TsCreate
	}
	return nil
}

func (x *OrderItem) GetTsModify() *timestamp.Timestamp {
	if x != nil {
		return x.TsModify
	}
	return nil
}

func (x *OrderItem) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OrderItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type OrderItemFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids    []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	UserId string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *OrderItemFilter) Reset() {
	*x = OrderItemFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItemFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemFilter) ProtoMessage() {}

func (x *OrderItemFilter) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemFilter.ProtoReflect.Descriptor instead.
func (*OrderItemFilter) Descriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{5}
}

func (x *OrderItemFilter) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *OrderItemFilter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column    string    `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Direction Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=order.api.Direction" json:"direction,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{6}
}

func (x *Order) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *Order) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_ASC
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_order_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_api_order_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_api_order_api_proto_rawDescGZIP(), []int{7}
}

func (x *Pagination) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_api_order_api_proto protoreflect.FileDescriptor

var file_api_order_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x46, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x11, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x14, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x43, 0x0a, 0x15,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x90, 0x02, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x74, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x08, 0x74, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x09,
	0x74, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x74, 0x73, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x53, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x12, 0x32, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x2a, 0x4a, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x02, 0x2a,
	0x1e, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x53, 0x43, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x32,
	0xb4, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x6b, 0x6f, 0x76, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_order_api_proto_rawDescOnce sync.Once
	file_api_order_api_proto_rawDescData = file_api_order_api_proto_rawDesc
)

func file_api_order_api_proto_rawDescGZIP() []byte {
	file_api_order_api_proto_rawDescOnce.Do(func() {
		file_api_order_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_order_api_proto_rawDescData)
	})
	return file_api_order_api_proto_rawDescData
}

var file_api_order_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_order_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_order_api_proto_goTypes = []interface{}{
	(OrderItemStatus)(0),          // 0: order.api.OrderItemStatus
	(Direction)(0),                // 1: order.api.Direction
	(*OrderItemRequest)(nil),      // 2: order.api.OrderItemRequest
	(*OrderItemResponse)(nil),     // 3: order.api.OrderItemResponse
	(*OrderItemListRequest)(nil),  // 4: order.api.OrderItemListRequest
	(*OrderItemListResponse)(nil), // 5: order.api.OrderItemListResponse
	(*OrderItem)(nil),             // 6: order.api.OrderItem
	(*OrderItemFilter)(nil),       // 7: order.api.OrderItemFilter
	(*Order)(nil),                 // 8: order.api.Order
	(*Pagination)(nil),            // 9: order.api.Pagination
	(*timestamp.Timestamp)(nil),   // 10: google.protobuf.Timestamp
}
var file_api_order_api_proto_depIdxs = []int32{
	7,  // 0: order.api.OrderItemRequest.filter:type_name -> order.api.OrderItemFilter
	6,  // 1: order.api.OrderItemResponse.value:type_name -> order.api.OrderItem
	7,  // 2: order.api.OrderItemListRequest.filter:type_name -> order.api.OrderItemFilter
	8,  // 3: order.api.OrderItemListRequest.orders:type_name -> order.api.Order
	9,  // 4: order.api.OrderItemListRequest.pagination:type_name -> order.api.Pagination
	6,  // 5: order.api.OrderItemListResponse.value:type_name -> order.api.OrderItem
	0,  // 6: order.api.OrderItem.status:type_name -> order.api.OrderItemStatus
	10, // 7: order.api.OrderItem.ts_create:type_name -> google.protobuf.Timestamp
	10, // 8: order.api.OrderItem.ts_modify:type_name -> google.protobuf.Timestamp
	1,  // 9: order.api.Order.direction:type_name -> order.api.Direction
	2,  // 10: order.api.OrderService.GetOrderItem:input_type -> order.api.OrderItemRequest
	4,  // 11: order.api.OrderService.GetOrderItemList:input_type -> order.api.OrderItemListRequest
	3,  // 12: order.api.OrderService.GetOrderItem:output_type -> order.api.OrderItemResponse
	5,  // 13: order.api.OrderService.GetOrderItemList:output_type -> order.api.OrderItemListResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_api_order_api_proto_init() }
func file_api_order_api_proto_init() {
	if File_api_order_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_order_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemRequest); i {
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
		file_api_order_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemResponse); i {
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
		file_api_order_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemListRequest); i {
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
		file_api_order_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemListResponse); i {
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
		file_api_order_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_api_order_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItemFilter); i {
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
		file_api_order_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_api_order_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
	file_api_order_api_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_order_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_order_api_proto_goTypes,
		DependencyIndexes: file_api_order_api_proto_depIdxs,
		EnumInfos:         file_api_order_api_proto_enumTypes,
		MessageInfos:      file_api_order_api_proto_msgTypes,
	}.Build()
	File_api_order_api_proto = out.File
	file_api_order_api_proto_rawDesc = nil
	file_api_order_api_proto_goTypes = nil
	file_api_order_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	GetOrderItem(ctx context.Context, in *OrderItemRequest, opts ...grpc.CallOption) (*OrderItemResponse, error)
	GetOrderItemList(ctx context.Context, in *OrderItemListRequest, opts ...grpc.CallOption) (*OrderItemListResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) GetOrderItem(ctx context.Context, in *OrderItemRequest, opts ...grpc.CallOption) (*OrderItemResponse, error) {
	out := new(OrderItemResponse)
	err := c.cc.Invoke(ctx, "/order.api.OrderService/GetOrderItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderItemList(ctx context.Context, in *OrderItemListRequest, opts ...grpc.CallOption) (*OrderItemListResponse, error) {
	out := new(OrderItemListResponse)
	err := c.cc.Invoke(ctx, "/order.api.OrderService/GetOrderItemList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	GetOrderItem(context.Context, *OrderItemRequest) (*OrderItemResponse, error)
	GetOrderItemList(context.Context, *OrderItemListRequest) (*OrderItemListResponse, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) GetOrderItem(context.Context, *OrderItemRequest) (*OrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderItem not implemented")
}
func (*UnimplementedOrderServiceServer) GetOrderItemList(context.Context, *OrderItemListRequest) (*OrderItemListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderItemList not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_GetOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.api.OrderService/GetOrderItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderItem(ctx, req.(*OrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderItemList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderItemListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderItemList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.api.OrderService/GetOrderItemList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderItemList(ctx, req.(*OrderItemListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.api.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrderItem",
			Handler:    _OrderService_GetOrderItem_Handler,
		},
		{
			MethodName: "GetOrderItemList",
			Handler:    _OrderService_GetOrderItemList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/order.api.proto",
}
