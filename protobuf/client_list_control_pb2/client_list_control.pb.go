// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/client_list_control_pb2/client_list_control.proto

package client_list_control_pb2

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

// Paging controls to be sent with List requests.
// Attributes:
//     start: The id of a resource to start the page with
//     limit: The number of results per page, defaults to 100 and maxes out at 1000
type ClientPagingControls struct {
	Start                string   `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientPagingControls) Reset()         { *m = ClientPagingControls{} }
func (m *ClientPagingControls) String() string { return proto.CompactTextString(m) }
func (*ClientPagingControls) ProtoMessage()    {}
func (*ClientPagingControls) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e5ee926f9a97d8, []int{0}
}

func (m *ClientPagingControls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientPagingControls.Unmarshal(m, b)
}
func (m *ClientPagingControls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientPagingControls.Marshal(b, m, deterministic)
}
func (m *ClientPagingControls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientPagingControls.Merge(m, src)
}
func (m *ClientPagingControls) XXX_Size() int {
	return xxx_messageInfo_ClientPagingControls.Size(m)
}
func (m *ClientPagingControls) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientPagingControls.DiscardUnknown(m)
}

var xxx_messageInfo_ClientPagingControls proto.InternalMessageInfo

func (m *ClientPagingControls) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *ClientPagingControls) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// Information about the pagination used, sent back with List responses.
// Attributes:
//     next: The id of the first resource in the next page
//     start: The id of the first resource in the returned page
//     limit: The number of results per page, defaults to 100 and maxes out at 1000
type ClientPagingResponse struct {
	Next                 string   `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Start                string   `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientPagingResponse) Reset()         { *m = ClientPagingResponse{} }
func (m *ClientPagingResponse) String() string { return proto.CompactTextString(m) }
func (*ClientPagingResponse) ProtoMessage()    {}
func (*ClientPagingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e5ee926f9a97d8, []int{1}
}

func (m *ClientPagingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientPagingResponse.Unmarshal(m, b)
}
func (m *ClientPagingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientPagingResponse.Marshal(b, m, deterministic)
}
func (m *ClientPagingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientPagingResponse.Merge(m, src)
}
func (m *ClientPagingResponse) XXX_Size() int {
	return xxx_messageInfo_ClientPagingResponse.Size(m)
}
func (m *ClientPagingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientPagingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientPagingResponse proto.InternalMessageInfo

func (m *ClientPagingResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func (m *ClientPagingResponse) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *ClientPagingResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// Sorting controls to be sent with List requests. More than one can be sent.
// If so, the first is used, and additional controls are tie-breakers.
// Attributes:
//     keys: Nested set of keys to sort by (i.e. ['default, block_num'])
//     reverse: Whether or not to reverse the sort (i.e. descending order)
type ClientSortControls struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Reverse              bool     `protobuf:"varint,2,opt,name=reverse,proto3" json:"reverse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientSortControls) Reset()         { *m = ClientSortControls{} }
func (m *ClientSortControls) String() string { return proto.CompactTextString(m) }
func (*ClientSortControls) ProtoMessage()    {}
func (*ClientSortControls) Descriptor() ([]byte, []int) {
	return fileDescriptor_67e5ee926f9a97d8, []int{2}
}

func (m *ClientSortControls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientSortControls.Unmarshal(m, b)
}
func (m *ClientSortControls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientSortControls.Marshal(b, m, deterministic)
}
func (m *ClientSortControls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientSortControls.Merge(m, src)
}
func (m *ClientSortControls) XXX_Size() int {
	return xxx_messageInfo_ClientSortControls.Size(m)
}
func (m *ClientSortControls) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientSortControls.DiscardUnknown(m)
}

var xxx_messageInfo_ClientSortControls proto.InternalMessageInfo

func (m *ClientSortControls) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ClientSortControls) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

func init() {
	proto.RegisterType((*ClientPagingControls)(nil), "ClientPagingControls")
	proto.RegisterType((*ClientPagingResponse)(nil), "ClientPagingResponse")
	proto.RegisterType((*ClientSortControls)(nil), "ClientSortControls")
}

func init() {
	proto.RegisterFile("protobuf/client_list_control_pb2/client_list_control.proto", fileDescriptor_67e5ee926f9a97d8)
}

var fileDescriptor_67e5ee926f9a97d8 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0x4b, 0x6b, 0xc3, 0x30,
	0x0c, 0x07, 0x70, 0xdc, 0xc7, 0x1e, 0x3a, 0x9a, 0x8e, 0xf9, 0x18, 0x7a, 0xea, 0x29, 0x83, 0xee,
	0xb6, 0x63, 0xfa, 0x05, 0x8a, 0x07, 0x3b, 0xec, 0x12, 0x92, 0x4e, 0xeb, 0x4c, 0x3c, 0x2b, 0x58,
	0xda, 0xeb, 0xdb, 0x8f, 0xd8, 0x64, 0x0f, 0x48, 0x6f, 0xfa, 0x4b, 0xf6, 0x0f, 0x21, 0xb8, 0xeb,
	0x23, 0x09, 0xb5, 0x6f, 0xcf, 0x37, 0x07, 0xef, 0x30, 0x48, 0xed, 0x1d, 0x4b, 0x7d, 0xa0, 0x20,
	0x91, 0x7c, 0xdd, 0xb7, 0xdb, 0xa9, 0x7e, 0x99, 0x3e, 0xad, 0x2b, 0x58, 0xed, 0xd2, 0x70, 0xdf,
	0x1c, 0x5d, 0x38, 0xee, 0xf2, 0x90, 0xf5, 0x0a, 0x96, 0x2c, 0x4d, 0x14, 0xa3, 0x0a, 0xb5, 0xb9,
	0xb4, 0x39, 0x0c, 0x5d, 0xef, 0x5e, 0x9d, 0x98, 0x59, 0xa1, 0x36, 0x4b, 0x9b, 0xc3, 0xfa, 0xe1,
	0xbf, 0x61, 0x91, 0x7b, 0x0a, 0x8c, 0x5a, 0xc3, 0x22, 0xe0, 0xe7, 0x48, 0xa4, 0xfa, 0xd7, 0x9d,
	0x4d, 0xba, 0xf3, 0xbf, 0x6e, 0x05, 0x3a, 0xbb, 0xf7, 0x14, 0xe5, 0x67, 0x33, 0x0d, 0x8b, 0x0e,
	0xbf, 0xd8, 0xa8, 0x62, 0x3e, 0xa8, 0x43, 0xad, 0x0d, 0x9c, 0x47, 0x7c, 0xc7, 0xc8, 0x98, 0xdc,
	0x0b, 0x3b, 0xc6, 0x6a, 0x0b, 0x57, 0xdc, 0x7c, 0x08, 0x91, 0xbc, 0x94, 0xfc, 0xd4, 0x95, 0xe3,
	0xa9, 0xf6, 0xea, 0xf1, 0xfa, 0xc4, 0xb5, 0xda, 0xb3, 0xf4, 0xe8, 0xf6, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x96, 0xd5, 0x7f, 0xd9, 0x58, 0x01, 0x00, 0x00,
}
