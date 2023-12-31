// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/infra/grpc/protofiles/order.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type XBlank struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XBlank) Reset()         { *m = XBlank{} }
func (m *XBlank) String() string { return proto.CompactTextString(m) }
func (*XBlank) ProtoMessage()    {}
func (*XBlank) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{0}
}

func (m *XBlank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XBlank.Unmarshal(m, b)
}
func (m *XBlank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XBlank.Marshal(b, m, deterministic)
}
func (m *XBlank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XBlank.Merge(m, src)
}
func (m *XBlank) XXX_Size() int {
	return xxx_messageInfo_XBlank.Size(m)
}
func (m *XBlank) XXX_DiscardUnknown() {
	xxx_messageInfo_XBlank.DiscardUnknown(m)
}

var xxx_messageInfo_XBlank proto.InternalMessageInfo

type Order struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice           float32  `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{1}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Order) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

func (m *Order) GetFinalPrice() float32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

type CreateOrderRequest struct {
	Price                float32  `protobuf:"fixed32,1,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,2,opt,name=tax,proto3" json:"tax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{2}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateOrderRequest) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

type CreateOrderResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Tax                  float32  `protobuf:"fixed32,3,opt,name=tax,proto3" json:"tax,omitempty"`
	FinalPrice           float32  `protobuf:"fixed32,4,opt,name=final_price,json=finalPrice,proto3" json:"final_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{3}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrderResponse) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *CreateOrderResponse) GetTax() float32 {
	if m != nil {
		return m.Tax
	}
	return 0
}

func (m *CreateOrderResponse) GetFinalPrice() float32 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

type ListOrdersResponse struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOrdersResponse) Reset()         { *m = ListOrdersResponse{} }
func (m *ListOrdersResponse) String() string { return proto.CompactTextString(m) }
func (*ListOrdersResponse) ProtoMessage()    {}
func (*ListOrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a81083cb550ac9a5, []int{4}
}

func (m *ListOrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOrdersResponse.Unmarshal(m, b)
}
func (m *ListOrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOrdersResponse.Marshal(b, m, deterministic)
}
func (m *ListOrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOrdersResponse.Merge(m, src)
}
func (m *ListOrdersResponse) XXX_Size() int {
	return xxx_messageInfo_ListOrdersResponse.Size(m)
}
func (m *ListOrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOrdersResponse proto.InternalMessageInfo

func (m *ListOrdersResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*XBlank)(nil), "pb._blank")
	proto.RegisterType((*Order)(nil), "pb.Order")
	proto.RegisterType((*CreateOrderRequest)(nil), "pb.CreateOrderRequest")
	proto.RegisterType((*CreateOrderResponse)(nil), "pb.CreateOrderResponse")
	proto.RegisterType((*ListOrdersResponse)(nil), "pb.ListOrdersResponse")
}

func init() {
	proto.RegisterFile("internal/infra/grpc/protofiles/order.proto", fileDescriptor_a81083cb550ac9a5)
}

var fileDescriptor_a81083cb550ac9a5 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x65, 0xf7, 0xdf, 0xea, 0xdf, 0x0b, 0x42, 0x95, 0x41, 0x25, 0xea, 0x42, 0xc8, 0x14,
	0x81, 0x94, 0x54, 0x65, 0x60, 0x41, 0x0c, 0xb0, 0x22, 0x81, 0xd2, 0x8d, 0xa5, 0x38, 0xcd, 0x05,
	0x59, 0x04, 0xc7, 0xd8, 0x06, 0xf1, 0xa5, 0xf8, 0x8e, 0xc8, 0x4e, 0xd5, 0x50, 0x35, 0x2b, 0x5b,
	0xf2, 0xbb, 0xbb, 0xf7, 0x9e, 0xce, 0x07, 0xe7, 0x42, 0x5a, 0xd4, 0x92, 0xd7, 0x99, 0x90, 0x95,
	0xe6, 0xd9, 0x8b, 0x56, 0xeb, 0x4c, 0xe9, 0xc6, 0x36, 0x95, 0xa8, 0xd1, 0x64, 0x8d, 0x2e, 0x51,
	0xa7, 0x1e, 0x30, 0xaa, 0x8a, 0xf8, 0x3f, 0x8c, 0x56, 0x45, 0xcd, 0xe5, 0x6b, 0xfc, 0x0c, 0xc3,
	0x07, 0x57, 0x64, 0x87, 0x40, 0x45, 0x19, 0x92, 0x88, 0x24, 0xe3, 0x9c, 0x8a, 0x92, 0x1d, 0xc3,
	0x50, 0x69, 0xb1, 0xc6, 0x90, 0x46, 0x24, 0xa1, 0x79, 0xfb, 0xc3, 0x26, 0x30, 0xb0, 0xfc, 0x2b,
	0x1c, 0x78, 0xe6, 0x3e, 0xd9, 0x29, 0x04, 0x95, 0x90, 0xbc, 0x5e, 0xb5, 0xdd, 0xff, 0x7c, 0x05,
	0x3c, 0x7a, 0x74, 0x24, 0xbe, 0x06, 0x76, 0xa7, 0x91, 0x5b, 0xf4, 0x3e, 0x39, 0xbe, 0x7f, 0xa0,
	0xb1, 0x9d, 0x3c, 0xe9, 0x91, 0xa7, 0x5b, 0xf9, 0x58, 0xc2, 0xd1, 0xce, 0xb4, 0x51, 0x8d, 0x34,
	0xf8, 0x77, 0x69, 0xaf, 0x80, 0xdd, 0x0b, 0x63, 0xbd, 0x9b, 0xd9, 0xda, 0x9d, 0xc1, 0xc8, 0xaf,
	0xd0, 0x84, 0x24, 0x1a, 0x24, 0xc1, 0x62, 0x9c, 0xaa, 0x22, 0x6d, 0x13, 0x6d, 0x0a, 0x8b, 0x6f,
	0x02, 0x07, 0x9e, 0x2c, 0x51, 0x7f, 0x3a, 0xf3, 0x1b, 0x08, 0x7e, 0x25, 0x67, 0x53, 0x37, 0xb2,
	0xbf, 0x88, 0xd9, 0xc9, 0x1e, 0xdf, 0x78, 0xce, 0x01, 0xba, 0x24, 0x0c, 0x5c, 0x5b, 0xfb, 0x66,
	0x33, 0x2f, 0xd5, 0x93, 0xf2, 0x02, 0x26, 0x1d, 0x5d, 0x5a, 0x8d, 0xfc, 0x6d, 0x67, 0xae, 0x4b,
	0x3d, 0x27, 0xb7, 0xe1, 0xd3, 0xb4, 0xf7, 0x68, 0x8a, 0x62, 0xe4, 0xef, 0xe4, 0xf2, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x38, 0x52, 0x95, 0xfb, 0x55, 0x02, 0x00, 0x00,
}
