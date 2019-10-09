// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/transaction_pb2/transaction.proto

package transaction_pb2

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

type TransactionHeader struct {
	// Public key for the client who added this transaction to a batch
	BatcherPublicKey string `protobuf:"bytes,1,opt,name=batcher_public_key,json=batcherPublicKey,proto3" json:"batcher_public_key,omitempty"`
	// A list of transaction signatures that describe the transactions that
	// must be processed before this transaction can be valid
	Dependencies []string `protobuf:"bytes,2,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	// The family name correlates to the transaction processor's family name
	// that this transaction can be processed on, for example 'intkey'
	FamilyName string `protobuf:"bytes,3,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	// The family version correlates to the transaction processor's family
	// version that this transaction can be processed on, for example "1.0"
	FamilyVersion string `protobuf:"bytes,4,opt,name=family_version,json=familyVersion,proto3" json:"family_version,omitempty"`
	// A list of addresses that are given to the context manager and control
	// what addresses the transaction processor is allowed to read from.
	Inputs []string `protobuf:"bytes,5,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// A random string that provides uniqueness for transactions with
	// otherwise identical fields.
	Nonce string `protobuf:"bytes,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// A list of addresses that are given to the context manager and control
	// what addresses the transaction processor is allowed to write to.
	Outputs []string `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty"`
	//The sha512 hash of the encoded payload
	PayloadSha512 string `protobuf:"bytes,9,opt,name=payload_sha512,json=payloadSha512,proto3" json:"payload_sha512,omitempty"`
	// Public key for the client that signed the TransactionHeader
	SignerPublicKey      string   `protobuf:"bytes,10,opt,name=signer_public_key,json=signerPublicKey,proto3" json:"signer_public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionHeader) Reset()         { *m = TransactionHeader{} }
func (m *TransactionHeader) String() string { return proto.CompactTextString(m) }
func (*TransactionHeader) ProtoMessage()    {}
func (*TransactionHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_332502cd0484d4bc, []int{0}
}

func (m *TransactionHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionHeader.Unmarshal(m, b)
}
func (m *TransactionHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionHeader.Marshal(b, m, deterministic)
}
func (m *TransactionHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionHeader.Merge(m, src)
}
func (m *TransactionHeader) XXX_Size() int {
	return xxx_messageInfo_TransactionHeader.Size(m)
}
func (m *TransactionHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionHeader.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionHeader proto.InternalMessageInfo

func (m *TransactionHeader) GetBatcherPublicKey() string {
	if m != nil {
		return m.BatcherPublicKey
	}
	return ""
}

func (m *TransactionHeader) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *TransactionHeader) GetFamilyName() string {
	if m != nil {
		return m.FamilyName
	}
	return ""
}

func (m *TransactionHeader) GetFamilyVersion() string {
	if m != nil {
		return m.FamilyVersion
	}
	return ""
}

func (m *TransactionHeader) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *TransactionHeader) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *TransactionHeader) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *TransactionHeader) GetPayloadSha512() string {
	if m != nil {
		return m.PayloadSha512
	}
	return ""
}

func (m *TransactionHeader) GetSignerPublicKey() string {
	if m != nil {
		return m.SignerPublicKey
	}
	return ""
}

type Transaction struct {
	// The serialized version of the TransactionHeader
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The signature derived from signing the header
	HeaderSignature string `protobuf:"bytes,2,opt,name=header_signature,json=headerSignature,proto3" json:"header_signature,omitempty"`
	// The payload is the encoded family specific information of the
	// transaction. Example cbor({'Verb': verb, 'Name': name,'Value': value})
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_332502cd0484d4bc, []int{1}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Transaction) GetHeaderSignature() string {
	if m != nil {
		return m.HeaderSignature
	}
	return ""
}

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// A simple list of transactions that needs to be serialized before
// it can be transmitted to a batcher.
type TransactionList struct {
	Transactions         []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransactionList) Reset()         { *m = TransactionList{} }
func (m *TransactionList) String() string { return proto.CompactTextString(m) }
func (*TransactionList) ProtoMessage()    {}
func (*TransactionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_332502cd0484d4bc, []int{2}
}

func (m *TransactionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionList.Unmarshal(m, b)
}
func (m *TransactionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionList.Marshal(b, m, deterministic)
}
func (m *TransactionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionList.Merge(m, src)
}
func (m *TransactionList) XXX_Size() int {
	return xxx_messageInfo_TransactionList.Size(m)
}
func (m *TransactionList) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionList.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionList proto.InternalMessageInfo

func (m *TransactionList) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionHeader)(nil), "TransactionHeader")
	proto.RegisterType((*Transaction)(nil), "Transaction")
	proto.RegisterType((*TransactionList)(nil), "TransactionList")
}

func init() {
	proto.RegisterFile("protobuf/transaction_pb2/transaction.proto", fileDescriptor_332502cd0484d4bc)
}

var fileDescriptor_332502cd0484d4bc = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xd1, 0x4a, 0xeb, 0x40,
	0x10, 0x86, 0x49, 0x7b, 0xda, 0xd2, 0x69, 0xce, 0x69, 0xbb, 0x9c, 0x73, 0xd8, 0x3b, 0x4b, 0x40,
	0xa8, 0x45, 0xa2, 0x56, 0x7c, 0x01, 0xbd, 0x11, 0x14, 0x29, 0xa9, 0x78, 0xe1, 0x4d, 0xd8, 0x24,
	0x5b, 0xb3, 0xb6, 0xdd, 0x0d, 0xd9, 0x8d, 0x92, 0x37, 0xf4, 0xb1, 0x24, 0xb3, 0x89, 0xa6, 0xde,
	0xe5, 0xff, 0xe6, 0xcf, 0xcc, 0xf0, 0xef, 0xc0, 0x22, 0xcb, 0x95, 0x51, 0x51, 0xb1, 0x39, 0x33,
	0x39, 0x93, 0x9a, 0xc5, 0x46, 0x28, 0x19, 0x66, 0xd1, 0xb2, 0xad, 0x7d, 0x34, 0x79, 0x1f, 0x1d,
	0x98, 0x3e, 0x7e, 0xd3, 0x5b, 0xce, 0x12, 0x9e, 0x93, 0x53, 0x20, 0x11, 0x33, 0x71, 0xca, 0xf3,
	0x30, 0x2b, 0xa2, 0x9d, 0x88, 0xc3, 0x2d, 0x2f, 0xa9, 0x33, 0x73, 0xe6, 0xc3, 0x60, 0x52, 0x57,
	0x56, 0x58, 0xb8, 0xe3, 0x25, 0xf1, 0xc0, 0x4d, 0x78, 0xc6, 0x65, 0xc2, 0x65, 0x2c, 0xb8, 0xa6,
	0x9d, 0x59, 0x77, 0x3e, 0x0c, 0x0e, 0x18, 0x39, 0x82, 0xd1, 0x86, 0xed, 0xc5, 0xae, 0x0c, 0x25,
	0xdb, 0x73, 0xda, 0xc5, 0x56, 0x60, 0xd1, 0x03, 0xdb, 0x73, 0x72, 0x0c, 0x7f, 0x6a, 0xc3, 0x1b,
	0xcf, 0xb5, 0x50, 0x92, 0xfe, 0x42, 0xcf, 0x6f, 0x4b, 0x9f, 0x2c, 0x24, 0xff, 0xa1, 0x2f, 0x64,
	0x56, 0x18, 0x4d, 0x7b, 0x38, 0xa5, 0x56, 0xe4, 0x2f, 0xf4, 0xa4, 0x92, 0x31, 0xa7, 0x7d, 0xfc,
	0xcb, 0x0a, 0x42, 0x61, 0xa0, 0x0a, 0x83, 0xf6, 0x01, 0xda, 0x1b, 0x59, 0x8d, 0xcb, 0x58, 0xb9,
	0x53, 0x2c, 0x09, 0x75, 0xca, 0xae, 0x2e, 0x96, 0x74, 0x68, 0xc7, 0xd5, 0x74, 0x8d, 0x90, 0x2c,
	0x60, 0xaa, 0xc5, 0x8b, 0x3c, 0xcc, 0x01, 0xd0, 0x39, 0xb6, 0x85, 0xaf, 0x18, 0xbc, 0x57, 0x18,
	0xb5, 0x92, 0xac, 0x36, 0x4d, 0x31, 0x4d, 0xcc, 0xcd, 0x0d, 0x6a, 0x45, 0x4e, 0x60, 0x62, 0xbf,
	0xc2, 0xaa, 0x01, 0x33, 0x45, 0xce, 0x69, 0xc7, 0x76, 0xb4, 0x7c, 0xdd, 0xe0, 0x6a, 0xfd, 0x7a,
	0x1d, 0x0c, 0xcc, 0x0d, 0x1a, 0xe9, 0xdd, 0xc0, 0xb8, 0x35, 0xeb, 0x5e, 0x68, 0x43, 0xce, 0xc1,
	0x6d, 0x3d, 0xaf, 0xa6, 0xce, 0xac, 0x3b, 0x1f, 0x2d, 0x5d, 0xbf, 0xe5, 0x0b, 0x0e, 0x1c, 0xd7,
	0x0b, 0xf8, 0xa7, 0xd9, 0xbb, 0x51, 0xca, 0xa4, 0xbe, 0x4e, 0xb6, 0x7e, 0x73, 0x36, 0x2b, 0xe7,
	0x79, 0xfc, 0xe3, 0x72, 0xa2, 0x3e, 0x16, 0x2f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xda, 0xaa,
	0x5a, 0xe9, 0x5c, 0x02, 0x00, 0x00,
}
