// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/setting_pb2/setting.proto

package setting_pb2

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

// Setting Container for on-chain configuration key/value pairs
type Setting struct {
	// List of setting entries - more than one implies a state key collision
	Entries              []*Setting_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Setting) Reset()         { *m = Setting{} }
func (m *Setting) String() string { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()    {}
func (*Setting) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9dfce539bcfa55f, []int{0}
}

func (m *Setting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Setting.Unmarshal(m, b)
}
func (m *Setting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Setting.Marshal(b, m, deterministic)
}
func (m *Setting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Setting.Merge(m, src)
}
func (m *Setting) XXX_Size() int {
	return xxx_messageInfo_Setting.Size(m)
}
func (m *Setting) XXX_DiscardUnknown() {
	xxx_messageInfo_Setting.DiscardUnknown(m)
}

var xxx_messageInfo_Setting proto.InternalMessageInfo

func (m *Setting) GetEntries() []*Setting_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Contains a setting entry (or entries, in the case of collisions).
type Setting_Entry struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Setting_Entry) Reset()         { *m = Setting_Entry{} }
func (m *Setting_Entry) String() string { return proto.CompactTextString(m) }
func (*Setting_Entry) ProtoMessage()    {}
func (*Setting_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9dfce539bcfa55f, []int{0, 0}
}

func (m *Setting_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Setting_Entry.Unmarshal(m, b)
}
func (m *Setting_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Setting_Entry.Marshal(b, m, deterministic)
}
func (m *Setting_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Setting_Entry.Merge(m, src)
}
func (m *Setting_Entry) XXX_Size() int {
	return xxx_messageInfo_Setting_Entry.Size(m)
}
func (m *Setting_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Setting_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Setting_Entry proto.InternalMessageInfo

func (m *Setting_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Setting_Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Setting)(nil), "Setting")
	proto.RegisterType((*Setting_Entry)(nil), "Setting.Entry")
}

func init() { proto.RegisterFile("protobuf/setting_pb2/setting.proto", fileDescriptor_f9dfce539bcfa55f) }

var fileDescriptor_f9dfce539bcfa55f = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x4e, 0x2d, 0x29, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0x48, 0x32,
	0x82, 0xb1, 0xf5, 0xc0, 0x92, 0x4a, 0x29, 0x5c, 0xec, 0xc1, 0x10, 0x01, 0x21, 0x0d, 0x2e, 0xf6,
	0xd4, 0xbc, 0x92, 0xa2, 0xcc, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x3e, 0x3d,
	0xa8, 0x94, 0x9e, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x4c, 0x5a, 0x4a, 0x9f, 0x8b, 0x15, 0x2c,
	0x22, 0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62,
	0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x81, 0xc5, 0x20, 0x1c, 0x27,
	0x35, 0x2e, 0xd1, 0xe2, 0xc4, 0xf2, 0x92, 0xfc, 0xfc, 0x92, 0x0c, 0xbd, 0xe2, 0x94, 0x6c, 0x3d,
	0x98, 0xc3, 0x02, 0x18, 0xa3, 0xb8, 0x91, 0xdc, 0x96, 0xc4, 0x06, 0x96, 0x30, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0xd3, 0x70, 0x8c, 0xba, 0x00, 0x00, 0x00,
}
