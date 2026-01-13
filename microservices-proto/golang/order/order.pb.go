

package order

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf


const _ = proto.ProtoPackageIsVersion3 

type CreateOrderRequest struct {
	CostumerId           int32        `protobuf:"varint,1,opt,name=costumer_id,json=costumerId,proto3" json:"costumer_id,omitempty"`
	OrderItems           []*OrderItem `protobuf:"bytes,2,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
	TotalPrice           float32      `protobuf:"fixed32,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
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

func (m *CreateOrderRequest) GetCostumerId() int32 {
	if m != nil {
		return m.CostumerId
	}
	return 0
}

func (m *CreateOrderRequest) GetOrderItems() []*OrderItem {
	if m != nil {
		return m.OrderItems
	}
	return nil
}

func (m *CreateOrderRequest) GetTotalPrice() float32 {
	if m != nil {
		return m.TotalPrice
	}
	return 0
}

type OrderItem struct {
	ProductCode          string   `protobuf:"bytes,1,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	UnitPrice            float32  `protobuf:"fixed32,2,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	Quantity             int32    `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderItem) Reset()         { *m = OrderItem{} }
func (m *OrderItem) String() string { return proto.CompactTextString(m) }
func (*OrderItem) ProtoMessage()    {}
func (*OrderItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *OrderItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderItem.Unmarshal(m, b)
}
func (m *OrderItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderItem.Marshal(b, m, deterministic)
}
func (m *OrderItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderItem.Merge(m, src)
}
func (m *OrderItem) XXX_Size() int {
	return xxx_messageInfo_OrderItem.Size(m)
}
func (m *OrderItem) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderItem.DiscardUnknown(m)
}

var xxx_messageInfo_OrderItem proto.InternalMessageInfo

func (m *OrderItem) GetProductCode() string {
	if m != nil {
		return m.ProductCode
	}
	return ""
}

func (m *OrderItem) GetUnitPrice() float32 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *OrderItem) GetQuantity() int32 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type CreateOrderResponse struct {
	OrderId              int32    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
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

func (m *CreateOrderResponse) GetOrderId() int32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateOrderRequest)(nil), "CreateOrderRequest")
	proto.RegisterType((*OrderItem)(nil), "OrderItem")
	proto.RegisterType((*CreateOrderResponse)(nil), "CreateOrderResponse")
}

func init() {
	proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077)
}

var fileDescriptor_cd01338c35d87077 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6b, 0xc2, 0x30,
	0x14, 0xc6, 0xa9, 0xa2, 0xd3, 0xd7, 0x9d, 0xe2, 0x0e, 0x4e, 0x18, 0x73, 0x9e, 0x84, 0x61, 0x3b,
	0xf4, 0x2a, 0x3b, 0xcc, 0x93, 0x07, 0xd9, 0xe8, 0x71, 0x97, 0x12, 0x93, 0x87, 0x0b, 0xb4, 0x7d,
	0x35, 0x79, 0x19, 0xec, 0xbc, 0x7f, 0x7c, 0x34, 0x56, 0x61, 0x6c, 0xb7, 0xf6, 0x47, 0xbe, 0xdf,
	0xc7, 0x97, 0x40, 0x4c, 0x56, 0xa3, 0x4d, 0x6a, 0x4b, 0x4c, 0xb3, 0xef, 0x08, 0xc4, 0xc6, 0xa2,
	0x64, 0x7c, 0x6d, 0x68, 0x86, 0x47, 0x8f, 0x8e, 0xc5, 0x3d, 0xc4, 0x8a, 0x1c, 0xfb, 0x12, 0x6d,
	0x6e, 0xf4, 0x38, 0x9a, 0x46, 0xf3, 0x5e, 0x06, 0x67, 0xb4, 0xd5, 0xe2, 0xb1, 0xd5, 0xe4, 0x86,
	0xb1, 0x74, 0xe3, 0xce, 0xb4, 0x3b, 0x8f, 0x97, 0x90, 0x04, 0xc9, 0x96, 0xb1, 0xcc, 0x80, 0xce,
	0x9f, 0xae, 0xb1, 0x31, 0xb1, 0x2c, 0xf2, 0xda, 0x1a, 0x85, 0xe3, 0xee, 0x34, 0x9a, 0x77, 0x32,
	0x08, 0xe8, 0xad, 0x21, 0x33, 0x03, 0xc3, 0x4b, 0x52, 0x3c, 0xc0, 0x75, 0x6d, 0x49, 0x7b, 0xc5,
	0xb9, 0x22, 0x8d, 0xa1, 0x7c, 0x98, 0xc5, 0x2d, 0xdb, 0x90, 0x46, 0x71, 0x07, 0xe0, 0x2b, 0xc3,
	0xad, 0xaf, 0x13, 0x7c, 0xc3, 0x86, 0x04, 0x9d, 0x98, 0xc0, 0xe0, 0xe8, 0x65, 0xc5, 0x86, 0xbf,
	0x42, 0x59, 0x2f, 0xbb, 0xfc, 0xcf, 0x9e, 0x60, 0xf4, 0x6b, 0xaf, 0xab, 0xa9, 0x72, 0x28, 0x6e,
	0x61, 0xd0, 0xee, 0x39, 0xaf, 0xbd, 0x3a, 0x0d, 0xd0, 0xcb, 0x35, 0xf4, 0xc2, 0x59, 0xb1, 0x82,
	0xfe, 0x29, 0x2a, 0x46, 0xc9, 0xdf, 0x3b, 0x9b, 0xdc, 0x24, 0xff, 0x88, 0x5f, 0x9e, 0xdf, 0xd7,
	0x07, 0xc3, 0x1f, 0x7e, 0x9f, 0x28, 0x2a, 0xd3, 0x1d, 0x16, 0x85, 0xac, 0x0c, 0x2e, 0x76, 0xd2,
	0x36, 0xa9, 0xb4, 0x34, 0xca, 0x92, 0x43, 0xfb, 0x69, 0x14, 0xba, 0x45, 0x78, 0x93, 0xf4, 0x40,
	0x85, 0xac, 0x0e, 0x69, 0xe8, 0xdf, 0xf7, 0x03, 0x5b, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x95,
	0xf0, 0xbc, 0x8e, 0xb6, 0x01, 0x00, 0x00,
}
