// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/google_ads_field_category.proto

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

// The category of the artifact.
type GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory int32

const (
	// Unspecified
	GoogleAdsFieldCategoryEnum_UNSPECIFIED GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 0
	// Unknown
	GoogleAdsFieldCategoryEnum_UNKNOWN GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 1
	// The described artifact is a resource.
	GoogleAdsFieldCategoryEnum_RESOURCE GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 2
	// The described artifact is a field and is an attribute of a resource.
	// Including a resource attribute field in a query may segment the query if
	// the resource to which it is attributed segments the resource found in
	// the FROM clause.
	GoogleAdsFieldCategoryEnum_ATTRIBUTE GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 3
	// The described artifact is a field and always segments search queries.
	GoogleAdsFieldCategoryEnum_SEGMENT GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 5
	// The described artifact is a field and is a metric. It never segments
	// search queries.
	GoogleAdsFieldCategoryEnum_METRIC GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory = 6
)

var GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "RESOURCE",
	3: "ATTRIBUTE",
	5: "SEGMENT",
	6: "METRIC",
}
var GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"RESOURCE":    2,
	"ATTRIBUTE":   3,
	"SEGMENT":     5,
	"METRIC":      6,
}

func (x GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) String() string {
	return proto.EnumName(GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_name, int32(x))
}
func (GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_google_ads_field_category_97cde39ce3aca2c0, []int{0, 0}
}

// Container for enum that determines if the described artifact is a resource
// or a field, and if it is a field, when it segments search queries.
type GoogleAdsFieldCategoryEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleAdsFieldCategoryEnum) Reset()         { *m = GoogleAdsFieldCategoryEnum{} }
func (m *GoogleAdsFieldCategoryEnum) String() string { return proto.CompactTextString(m) }
func (*GoogleAdsFieldCategoryEnum) ProtoMessage()    {}
func (*GoogleAdsFieldCategoryEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_google_ads_field_category_97cde39ce3aca2c0, []int{0}
}
func (m *GoogleAdsFieldCategoryEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleAdsFieldCategoryEnum.Unmarshal(m, b)
}
func (m *GoogleAdsFieldCategoryEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleAdsFieldCategoryEnum.Marshal(b, m, deterministic)
}
func (dst *GoogleAdsFieldCategoryEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleAdsFieldCategoryEnum.Merge(dst, src)
}
func (m *GoogleAdsFieldCategoryEnum) XXX_Size() int {
	return xxx_messageInfo_GoogleAdsFieldCategoryEnum.Size(m)
}
func (m *GoogleAdsFieldCategoryEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleAdsFieldCategoryEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleAdsFieldCategoryEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GoogleAdsFieldCategoryEnum)(nil), "google.ads.googleads.v0.enums.GoogleAdsFieldCategoryEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory", GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_name, GoogleAdsFieldCategoryEnum_GoogleAdsFieldCategory_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/google_ads_field_category.proto", fileDescriptor_google_ads_field_category_97cde39ce3aca2c0)
}

var fileDescriptor_google_ads_field_category_97cde39ce3aca2c0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0x5c, 0x98, 0x50, 0x7c, 0x62, 0x4a, 0x71, 0x7c, 0x5a, 0x66, 0x6a, 0x4e, 0x4a, 0x7c, 0x72,
	0x62, 0x49, 0x6a, 0x7a, 0x7e, 0x51, 0xa5, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x2c, 0x44,
	0x81, 0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xbb, 0x5e, 0x99, 0x81, 0x1e, 0x58, 0xbb, 0x52, 0x17,
	0x23, 0x97, 0x94, 0x3b, 0x58, 0xd8, 0x31, 0xa5, 0xd8, 0x0d, 0x64, 0x80, 0x33, 0x54, 0xbf, 0x6b,
	0x5e, 0x69, 0xae, 0x52, 0x0e, 0x97, 0x18, 0x76, 0x59, 0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0,
	0x00, 0x57, 0x67, 0x4f, 0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f,
	0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x46, 0x21, 0x1e, 0x2e, 0x8e, 0x20, 0xd7, 0x60, 0xff, 0xd0,
	0x20, 0x67, 0x57, 0x01, 0x26, 0x21, 0x5e, 0x2e, 0x4e, 0xc7, 0x90, 0x90, 0x20, 0x4f, 0xa7, 0xd0,
	0x10, 0x57, 0x01, 0x66, 0x90, 0xca, 0x60, 0x57, 0x77, 0x5f, 0x57, 0xbf, 0x10, 0x01, 0x56, 0x21,
	0x2e, 0x2e, 0x36, 0x5f, 0xd7, 0x90, 0x20, 0x4f, 0x67, 0x01, 0x36, 0xa7, 0x33, 0x8c, 0x5c, 0x8a,
	0xc9, 0xf9, 0xb9, 0x7a, 0x78, 0x9d, 0xec, 0x24, 0x8d, 0xdd, 0x45, 0x01, 0x20, 0xef, 0x06, 0x30,
	0x46, 0x39, 0x41, 0x75, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0xe5, 0x17, 0xa5, 0xeb, 0xa7,
	0xa7, 0xe6, 0x81, 0x03, 0x03, 0x16, 0x7e, 0x05, 0x99, 0xc5, 0x38, 0x82, 0xd3, 0x1a, 0x4c, 0x2e,
	0x62, 0x62, 0x76, 0x77, 0x74, 0x5c, 0xc5, 0x24, 0x0b, 0xb1, 0x49, 0xcf, 0x31, 0xa5, 0x58, 0x0f,
	0x6e, 0xa9, 0x5e, 0x98, 0x81, 0x1e, 0x28, 0x6c, 0x8a, 0x4f, 0xc1, 0xe4, 0x63, 0x1c, 0x53, 0x8a,
	0x63, 0xe0, 0xf2, 0x31, 0x61, 0x06, 0x31, 0x60, 0xf9, 0x24, 0x36, 0xb0, 0xa5, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0x56, 0xff, 0x0d, 0xc2, 0x01, 0x00, 0x00,
}
