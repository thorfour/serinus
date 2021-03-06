// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/product_channel.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing the locality of a product offer.
type ProductChannelEnum_ProductChannel int32

const (
	// Not specified.
	ProductChannelEnum_UNSPECIFIED ProductChannelEnum_ProductChannel = 0
	// Used for return value only. Represents value unknown in this version.
	ProductChannelEnum_UNKNOWN ProductChannelEnum_ProductChannel = 1
	// The item is sold online.
	ProductChannelEnum_ONLINE ProductChannelEnum_ProductChannel = 2
	// The item is sold in local stores.
	ProductChannelEnum_LOCAL ProductChannelEnum_ProductChannel = 3
)

var ProductChannelEnum_ProductChannel_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ONLINE",
	3: "LOCAL",
}
var ProductChannelEnum_ProductChannel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ONLINE":      2,
	"LOCAL":       3,
}

func (x ProductChannelEnum_ProductChannel) String() string {
	return proto.EnumName(ProductChannelEnum_ProductChannel_name, int32(x))
}
func (ProductChannelEnum_ProductChannel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_product_channel_81186d77dbd458b5, []int{0, 0}
}

// Locality of a product offer.
type ProductChannelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductChannelEnum) Reset()         { *m = ProductChannelEnum{} }
func (m *ProductChannelEnum) String() string { return proto.CompactTextString(m) }
func (*ProductChannelEnum) ProtoMessage()    {}
func (*ProductChannelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_channel_81186d77dbd458b5, []int{0}
}
func (m *ProductChannelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductChannelEnum.Unmarshal(m, b)
}
func (m *ProductChannelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductChannelEnum.Marshal(b, m, deterministic)
}
func (dst *ProductChannelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductChannelEnum.Merge(dst, src)
}
func (m *ProductChannelEnum) XXX_Size() int {
	return xxx_messageInfo_ProductChannelEnum.Size(m)
}
func (m *ProductChannelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductChannelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProductChannelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProductChannelEnum)(nil), "google.ads.googleads.v0.enums.ProductChannelEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.ProductChannelEnum_ProductChannel", ProductChannelEnum_ProductChannel_name, ProductChannelEnum_ProductChannel_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/product_channel.proto", fileDescriptor_product_channel_81186d77dbd458b5)
}

var fileDescriptor_product_channel_81186d77dbd458b5 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0x82, 0xa2, 0xfc, 0x94, 0xd2, 0xe4, 0x92, 0xf8, 0xe4, 0x8c, 0xc4, 0xbc,
	0xbc, 0xd4, 0x1c, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x4a, 0xbd, 0xc4, 0x94,
	0x62, 0x3d, 0xb8, 0x26, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x26, 0xa5, 0x68, 0x2e, 0xa1, 0x00, 0x88,
	0x3e, 0x67, 0x88, 0x36, 0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x57, 0x2e, 0x3e, 0x54, 0x51, 0x21, 0x7e,
	0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21,
	0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x46, 0x21, 0x2e, 0x2e, 0x36,
	0x7f, 0x3f, 0x1f, 0x4f, 0x3f, 0x57, 0x01, 0x26, 0x21, 0x4e, 0x2e, 0x56, 0x1f, 0x7f, 0x67, 0x47,
	0x1f, 0x01, 0x66, 0xa7, 0x23, 0x8c, 0x5c, 0x8a, 0xc9, 0xf9, 0xb9, 0x7a, 0x78, 0x9d, 0xe0, 0x24,
	0x8c, 0x6a, 0x55, 0x00, 0xc8, 0xd9, 0x01, 0x8c, 0x51, 0x4e, 0x50, 0x5d, 0xe9, 0xf9, 0x39, 0x89,
	0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0xfa, 0xe9, 0xa9, 0x79, 0x60, 0x4f, 0xc1, 0x7c, 0x5f, 0x90,
	0x59, 0x8c, 0x23, 0x30, 0xac, 0xc1, 0xe4, 0x22, 0x26, 0x66, 0x77, 0x47, 0xc7, 0x55, 0x4c, 0xb2,
	0xee, 0x10, 0xa3, 0x1c, 0x53, 0x8a, 0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x40, 0x0f, 0xe4, 0xd9,
	0xe2, 0x53, 0x30, 0xf9, 0x18, 0xc7, 0x94, 0xe2, 0x18, 0xb8, 0x7c, 0x4c, 0x98, 0x41, 0x0c, 0x58,
	0x3e, 0x89, 0x0d, 0x6c, 0xa9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x47, 0xf7, 0x35, 0x74, 0x80,
	0x01, 0x00, 0x00,
}
