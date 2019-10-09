// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/validator_pb2/validator.proto

package validator_pb2

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

type Message_MessageType int32

const (
	Message_DEFAULT Message_MessageType = 0
	// Registration request from the transaction processor to the validator
	Message_TP_REGISTER_REQUEST Message_MessageType = 1
	// Registration response from the validator to the
	// transaction processor
	Message_TP_REGISTER_RESPONSE Message_MessageType = 2
	// Tell the validator that the transaction processor
	// won't take any more transactions
	Message_TP_UNREGISTER_REQUEST Message_MessageType = 3
	// Response from the validator to the tp that it won't
	// send any more transactions
	Message_TP_UNREGISTER_RESPONSE Message_MessageType = 4
	// Process Request from the validator/executor to the
	// transaction processor
	Message_TP_PROCESS_REQUEST Message_MessageType = 5
	// Process response from the transaction processor to the validator/executor
	Message_TP_PROCESS_RESPONSE Message_MessageType = 6
	// State get request from the transaction processor to validator/context_manager
	Message_TP_STATE_GET_REQUEST Message_MessageType = 7
	// State get response from the validator/context_manager to the transaction processor
	Message_TP_STATE_GET_RESPONSE Message_MessageType = 8
	// State set request from the transaction processor to the validator/context_manager
	Message_TP_STATE_SET_REQUEST Message_MessageType = 9
	// State set response from the validator/context_manager to the transaction processor
	Message_TP_STATE_SET_RESPONSE Message_MessageType = 10
	// State delete request from the transaction processor to the validator/context_manager
	Message_TP_STATE_DELETE_REQUEST Message_MessageType = 11
	// State delete response from the validator/context_manager to the transaction processor
	Message_TP_STATE_DELETE_RESPONSE Message_MessageType = 12
	// Message to append data to a transaction receipt
	Message_TP_RECEIPT_ADD_DATA_REQUEST Message_MessageType = 13
	// Response from validator to tell transaction processor that data has been appended
	Message_TP_RECEIPT_ADD_DATA_RESPONSE Message_MessageType = 14
	// Message to add event
	Message_TP_EVENT_ADD_REQUEST Message_MessageType = 15
	// Response from validator to tell transaction processor that event has been created
	Message_TP_EVENT_ADD_RESPONSE Message_MessageType = 16
	// Submission of a batchlist from the web api or another client to the validator
	Message_CLIENT_BATCH_SUBMIT_REQUEST Message_MessageType = 100
	// Response from the validator to the web api/client that the submission was accepted
	Message_CLIENT_BATCH_SUBMIT_RESPONSE Message_MessageType = 101
	// A request to list blocks from the web api/client to the validator
	Message_CLIENT_BLOCK_LIST_REQUEST        Message_MessageType = 102
	Message_CLIENT_BLOCK_LIST_RESPONSE       Message_MessageType = 103
	Message_CLIENT_BLOCK_GET_BY_ID_REQUEST   Message_MessageType = 104
	Message_CLIENT_BLOCK_GET_RESPONSE        Message_MessageType = 105
	Message_CLIENT_BATCH_LIST_REQUEST        Message_MessageType = 106
	Message_CLIENT_BATCH_LIST_RESPONSE       Message_MessageType = 107
	Message_CLIENT_BATCH_GET_REQUEST         Message_MessageType = 108
	Message_CLIENT_BATCH_GET_RESPONSE        Message_MessageType = 109
	Message_CLIENT_TRANSACTION_LIST_REQUEST  Message_MessageType = 110
	Message_CLIENT_TRANSACTION_LIST_RESPONSE Message_MessageType = 111
	Message_CLIENT_TRANSACTION_GET_REQUEST   Message_MessageType = 112
	Message_CLIENT_TRANSACTION_GET_RESPONSE  Message_MessageType = 113
	// Client state request of the current state hash to be retrieved from the journal
	Message_CLIENT_STATE_CURRENT_REQUEST Message_MessageType = 114
	// Response with the current state hash
	Message_CLIENT_STATE_CURRENT_RESPONSE Message_MessageType = 115
	// A request of all the addresses under a particular prefix, for a state hash.
	Message_CLIENT_STATE_LIST_REQUEST Message_MessageType = 116
	// The response of the addresses
	Message_CLIENT_STATE_LIST_RESPONSE Message_MessageType = 117
	// Get the address:data entry at a particular address
	Message_CLIENT_STATE_GET_REQUEST Message_MessageType = 118
	// The response with the entry
	Message_CLIENT_STATE_GET_RESPONSE Message_MessageType = 119
	// A request for the status of a batch or batches
	Message_CLIENT_BATCH_STATUS_REQUEST Message_MessageType = 120
	// A response with the batch statuses
	Message_CLIENT_BATCH_STATUS_RESPONSE Message_MessageType = 121
	// A request for one or more transaction receipts
	Message_CLIENT_RECEIPT_GET_REQUEST Message_MessageType = 122
	// A response with the receipts
	Message_CLIENT_RECEIPT_GET_RESPONSE     Message_MessageType = 123
	Message_CLIENT_BLOCK_GET_BY_NUM_REQUEST Message_MessageType = 124
	// A request for a validator's peers
	Message_CLIENT_PEERS_GET_REQUEST Message_MessageType = 125
	// A response with the validator's peers
	Message_CLIENT_PEERS_GET_RESPONSE                  Message_MessageType = 126
	Message_CLIENT_BLOCK_GET_BY_TRANSACTION_ID_REQUEST Message_MessageType = 127
	Message_CLIENT_BLOCK_GET_BY_BATCH_ID_REQUEST       Message_MessageType = 128
	// A request for a validator's status
	Message_CLIENT_STATUS_GET_REQUEST Message_MessageType = 129
	// A response with the validator's status
	Message_CLIENT_STATUS_GET_RESPONSE Message_MessageType = 130
	// Message types for events
	Message_CLIENT_EVENTS_SUBSCRIBE_REQUEST    Message_MessageType = 500
	Message_CLIENT_EVENTS_SUBSCRIBE_RESPONSE   Message_MessageType = 501
	Message_CLIENT_EVENTS_UNSUBSCRIBE_REQUEST  Message_MessageType = 502
	Message_CLIENT_EVENTS_UNSUBSCRIBE_RESPONSE Message_MessageType = 503
	Message_CLIENT_EVENTS                      Message_MessageType = 504
	Message_CLIENT_EVENTS_GET_REQUEST          Message_MessageType = 505
	Message_CLIENT_EVENTS_GET_RESPONSE         Message_MessageType = 506
	// Temp message types until a discussion can be had about gossip msg
	Message_GOSSIP_MESSAGE                         Message_MessageType = 200
	Message_GOSSIP_REGISTER                        Message_MessageType = 201
	Message_GOSSIP_UNREGISTER                      Message_MessageType = 202
	Message_GOSSIP_BLOCK_REQUEST                   Message_MessageType = 205
	Message_GOSSIP_BLOCK_RESPONSE                  Message_MessageType = 206
	Message_GOSSIP_BATCH_BY_BATCH_ID_REQUEST       Message_MessageType = 207
	Message_GOSSIP_BATCH_BY_TRANSACTION_ID_REQUEST Message_MessageType = 208
	Message_GOSSIP_BATCH_RESPONSE                  Message_MessageType = 209
	Message_GOSSIP_GET_PEERS_REQUEST               Message_MessageType = 210
	Message_GOSSIP_GET_PEERS_RESPONSE              Message_MessageType = 211
	Message_GOSSIP_CONSENSUS_MESSAGE               Message_MessageType = 212
	Message_NETWORK_ACK                            Message_MessageType = 300
	Message_NETWORK_CONNECT                        Message_MessageType = 301
	Message_NETWORK_DISCONNECT                     Message_MessageType = 302
	// Message types for Authorization Types
	Message_AUTHORIZATION_CONNECTION_RESPONSE Message_MessageType = 600
	Message_AUTHORIZATION_VIOLATION           Message_MessageType = 601
	Message_AUTHORIZATION_TRUST_REQUEST       Message_MessageType = 602
	Message_AUTHORIZATION_TRUST_RESPONSE      Message_MessageType = 603
	Message_AUTHORIZATION_CHALLENGE_REQUEST   Message_MessageType = 604
	Message_AUTHORIZATION_CHALLENGE_RESPONSE  Message_MessageType = 605
	Message_AUTHORIZATION_CHALLENGE_SUBMIT    Message_MessageType = 606
	Message_AUTHORIZATION_CHALLENGE_RESULT    Message_MessageType = 607
	Message_PING_REQUEST                      Message_MessageType = 700
	Message_PING_RESPONSE                     Message_MessageType = 701
	// Consensus service messages
	Message_CONSENSUS_REGISTER_REQUEST          Message_MessageType = 800
	Message_CONSENSUS_REGISTER_RESPONSE         Message_MessageType = 801
	Message_CONSENSUS_SEND_TO_REQUEST           Message_MessageType = 802
	Message_CONSENSUS_SEND_TO_RESPONSE          Message_MessageType = 803
	Message_CONSENSUS_BROADCAST_REQUEST         Message_MessageType = 804
	Message_CONSENSUS_BROADCAST_RESPONSE        Message_MessageType = 805
	Message_CONSENSUS_INITIALIZE_BLOCK_REQUEST  Message_MessageType = 806
	Message_CONSENSUS_INITIALIZE_BLOCK_RESPONSE Message_MessageType = 807
	Message_CONSENSUS_FINALIZE_BLOCK_REQUEST    Message_MessageType = 808
	Message_CONSENSUS_FINALIZE_BLOCK_RESPONSE   Message_MessageType = 809
	Message_CONSENSUS_SUMMARIZE_BLOCK_REQUEST   Message_MessageType = 828
	Message_CONSENSUS_SUMMARIZE_BLOCK_RESPONSE  Message_MessageType = 829
	Message_CONSENSUS_CANCEL_BLOCK_REQUEST      Message_MessageType = 810
	Message_CONSENSUS_CANCEL_BLOCK_RESPONSE     Message_MessageType = 811
	Message_CONSENSUS_CHECK_BLOCKS_REQUEST      Message_MessageType = 812
	Message_CONSENSUS_CHECK_BLOCKS_RESPONSE     Message_MessageType = 813
	Message_CONSENSUS_COMMIT_BLOCK_REQUEST      Message_MessageType = 814
	Message_CONSENSUS_COMMIT_BLOCK_RESPONSE     Message_MessageType = 815
	Message_CONSENSUS_IGNORE_BLOCK_REQUEST      Message_MessageType = 816
	Message_CONSENSUS_IGNORE_BLOCK_RESPONSE     Message_MessageType = 817
	Message_CONSENSUS_FAIL_BLOCK_REQUEST        Message_MessageType = 818
	Message_CONSENSUS_FAIL_BLOCK_RESPONSE       Message_MessageType = 819
	Message_CONSENSUS_SETTINGS_GET_REQUEST      Message_MessageType = 820
	Message_CONSENSUS_SETTINGS_GET_RESPONSE     Message_MessageType = 821
	Message_CONSENSUS_STATE_GET_REQUEST         Message_MessageType = 822
	Message_CONSENSUS_STATE_GET_RESPONSE        Message_MessageType = 823
	Message_CONSENSUS_BLOCKS_GET_REQUEST        Message_MessageType = 824
	Message_CONSENSUS_BLOCKS_GET_RESPONSE       Message_MessageType = 825
	Message_CONSENSUS_CHAIN_HEAD_GET_REQUEST    Message_MessageType = 826
	Message_CONSENSUS_CHAIN_HEAD_GET_RESPONSE   Message_MessageType = 827
	// Consensus notification messages
	Message_CONSENSUS_NOTIFY_PEER_CONNECTED    Message_MessageType = 900
	Message_CONSENSUS_NOTIFY_PEER_DISCONNECTED Message_MessageType = 901
	Message_CONSENSUS_NOTIFY_PEER_MESSAGE      Message_MessageType = 902
	Message_CONSENSUS_NOTIFY_BLOCK_NEW         Message_MessageType = 903
	Message_CONSENSUS_NOTIFY_BLOCK_VALID       Message_MessageType = 904
	Message_CONSENSUS_NOTIFY_BLOCK_INVALID     Message_MessageType = 905
	Message_CONSENSUS_NOTIFY_BLOCK_COMMIT      Message_MessageType = 906
	Message_CONSENSUS_NOTIFY_ACK               Message_MessageType = 999
)

var Message_MessageType_name = map[int32]string{
	0:   "DEFAULT",
	1:   "TP_REGISTER_REQUEST",
	2:   "TP_REGISTER_RESPONSE",
	3:   "TP_UNREGISTER_REQUEST",
	4:   "TP_UNREGISTER_RESPONSE",
	5:   "TP_PROCESS_REQUEST",
	6:   "TP_PROCESS_RESPONSE",
	7:   "TP_STATE_GET_REQUEST",
	8:   "TP_STATE_GET_RESPONSE",
	9:   "TP_STATE_SET_REQUEST",
	10:  "TP_STATE_SET_RESPONSE",
	11:  "TP_STATE_DELETE_REQUEST",
	12:  "TP_STATE_DELETE_RESPONSE",
	13:  "TP_RECEIPT_ADD_DATA_REQUEST",
	14:  "TP_RECEIPT_ADD_DATA_RESPONSE",
	15:  "TP_EVENT_ADD_REQUEST",
	16:  "TP_EVENT_ADD_RESPONSE",
	100: "CLIENT_BATCH_SUBMIT_REQUEST",
	101: "CLIENT_BATCH_SUBMIT_RESPONSE",
	102: "CLIENT_BLOCK_LIST_REQUEST",
	103: "CLIENT_BLOCK_LIST_RESPONSE",
	104: "CLIENT_BLOCK_GET_BY_ID_REQUEST",
	105: "CLIENT_BLOCK_GET_RESPONSE",
	106: "CLIENT_BATCH_LIST_REQUEST",
	107: "CLIENT_BATCH_LIST_RESPONSE",
	108: "CLIENT_BATCH_GET_REQUEST",
	109: "CLIENT_BATCH_GET_RESPONSE",
	110: "CLIENT_TRANSACTION_LIST_REQUEST",
	111: "CLIENT_TRANSACTION_LIST_RESPONSE",
	112: "CLIENT_TRANSACTION_GET_REQUEST",
	113: "CLIENT_TRANSACTION_GET_RESPONSE",
	114: "CLIENT_STATE_CURRENT_REQUEST",
	115: "CLIENT_STATE_CURRENT_RESPONSE",
	116: "CLIENT_STATE_LIST_REQUEST",
	117: "CLIENT_STATE_LIST_RESPONSE",
	118: "CLIENT_STATE_GET_REQUEST",
	119: "CLIENT_STATE_GET_RESPONSE",
	120: "CLIENT_BATCH_STATUS_REQUEST",
	121: "CLIENT_BATCH_STATUS_RESPONSE",
	122: "CLIENT_RECEIPT_GET_REQUEST",
	123: "CLIENT_RECEIPT_GET_RESPONSE",
	124: "CLIENT_BLOCK_GET_BY_NUM_REQUEST",
	125: "CLIENT_PEERS_GET_REQUEST",
	126: "CLIENT_PEERS_GET_RESPONSE",
	127: "CLIENT_BLOCK_GET_BY_TRANSACTION_ID_REQUEST",
	128: "CLIENT_BLOCK_GET_BY_BATCH_ID_REQUEST",
	129: "CLIENT_STATUS_GET_REQUEST",
	130: "CLIENT_STATUS_GET_RESPONSE",
	500: "CLIENT_EVENTS_SUBSCRIBE_REQUEST",
	501: "CLIENT_EVENTS_SUBSCRIBE_RESPONSE",
	502: "CLIENT_EVENTS_UNSUBSCRIBE_REQUEST",
	503: "CLIENT_EVENTS_UNSUBSCRIBE_RESPONSE",
	504: "CLIENT_EVENTS",
	505: "CLIENT_EVENTS_GET_REQUEST",
	506: "CLIENT_EVENTS_GET_RESPONSE",
	200: "GOSSIP_MESSAGE",
	201: "GOSSIP_REGISTER",
	202: "GOSSIP_UNREGISTER",
	205: "GOSSIP_BLOCK_REQUEST",
	206: "GOSSIP_BLOCK_RESPONSE",
	207: "GOSSIP_BATCH_BY_BATCH_ID_REQUEST",
	208: "GOSSIP_BATCH_BY_TRANSACTION_ID_REQUEST",
	209: "GOSSIP_BATCH_RESPONSE",
	210: "GOSSIP_GET_PEERS_REQUEST",
	211: "GOSSIP_GET_PEERS_RESPONSE",
	212: "GOSSIP_CONSENSUS_MESSAGE",
	300: "NETWORK_ACK",
	301: "NETWORK_CONNECT",
	302: "NETWORK_DISCONNECT",
	600: "AUTHORIZATION_CONNECTION_RESPONSE",
	601: "AUTHORIZATION_VIOLATION",
	602: "AUTHORIZATION_TRUST_REQUEST",
	603: "AUTHORIZATION_TRUST_RESPONSE",
	604: "AUTHORIZATION_CHALLENGE_REQUEST",
	605: "AUTHORIZATION_CHALLENGE_RESPONSE",
	606: "AUTHORIZATION_CHALLENGE_SUBMIT",
	607: "AUTHORIZATION_CHALLENGE_RESULT",
	700: "PING_REQUEST",
	701: "PING_RESPONSE",
	800: "CONSENSUS_REGISTER_REQUEST",
	801: "CONSENSUS_REGISTER_RESPONSE",
	802: "CONSENSUS_SEND_TO_REQUEST",
	803: "CONSENSUS_SEND_TO_RESPONSE",
	804: "CONSENSUS_BROADCAST_REQUEST",
	805: "CONSENSUS_BROADCAST_RESPONSE",
	806: "CONSENSUS_INITIALIZE_BLOCK_REQUEST",
	807: "CONSENSUS_INITIALIZE_BLOCK_RESPONSE",
	808: "CONSENSUS_FINALIZE_BLOCK_REQUEST",
	809: "CONSENSUS_FINALIZE_BLOCK_RESPONSE",
	828: "CONSENSUS_SUMMARIZE_BLOCK_REQUEST",
	829: "CONSENSUS_SUMMARIZE_BLOCK_RESPONSE",
	810: "CONSENSUS_CANCEL_BLOCK_REQUEST",
	811: "CONSENSUS_CANCEL_BLOCK_RESPONSE",
	812: "CONSENSUS_CHECK_BLOCKS_REQUEST",
	813: "CONSENSUS_CHECK_BLOCKS_RESPONSE",
	814: "CONSENSUS_COMMIT_BLOCK_REQUEST",
	815: "CONSENSUS_COMMIT_BLOCK_RESPONSE",
	816: "CONSENSUS_IGNORE_BLOCK_REQUEST",
	817: "CONSENSUS_IGNORE_BLOCK_RESPONSE",
	818: "CONSENSUS_FAIL_BLOCK_REQUEST",
	819: "CONSENSUS_FAIL_BLOCK_RESPONSE",
	820: "CONSENSUS_SETTINGS_GET_REQUEST",
	821: "CONSENSUS_SETTINGS_GET_RESPONSE",
	822: "CONSENSUS_STATE_GET_REQUEST",
	823: "CONSENSUS_STATE_GET_RESPONSE",
	824: "CONSENSUS_BLOCKS_GET_REQUEST",
	825: "CONSENSUS_BLOCKS_GET_RESPONSE",
	826: "CONSENSUS_CHAIN_HEAD_GET_REQUEST",
	827: "CONSENSUS_CHAIN_HEAD_GET_RESPONSE",
	900: "CONSENSUS_NOTIFY_PEER_CONNECTED",
	901: "CONSENSUS_NOTIFY_PEER_DISCONNECTED",
	902: "CONSENSUS_NOTIFY_PEER_MESSAGE",
	903: "CONSENSUS_NOTIFY_BLOCK_NEW",
	904: "CONSENSUS_NOTIFY_BLOCK_VALID",
	905: "CONSENSUS_NOTIFY_BLOCK_INVALID",
	906: "CONSENSUS_NOTIFY_BLOCK_COMMIT",
	999: "CONSENSUS_NOTIFY_ACK",
}

var Message_MessageType_value = map[string]int32{
	"DEFAULT":                                    0,
	"TP_REGISTER_REQUEST":                        1,
	"TP_REGISTER_RESPONSE":                       2,
	"TP_UNREGISTER_REQUEST":                      3,
	"TP_UNREGISTER_RESPONSE":                     4,
	"TP_PROCESS_REQUEST":                         5,
	"TP_PROCESS_RESPONSE":                        6,
	"TP_STATE_GET_REQUEST":                       7,
	"TP_STATE_GET_RESPONSE":                      8,
	"TP_STATE_SET_REQUEST":                       9,
	"TP_STATE_SET_RESPONSE":                      10,
	"TP_STATE_DELETE_REQUEST":                    11,
	"TP_STATE_DELETE_RESPONSE":                   12,
	"TP_RECEIPT_ADD_DATA_REQUEST":                13,
	"TP_RECEIPT_ADD_DATA_RESPONSE":               14,
	"TP_EVENT_ADD_REQUEST":                       15,
	"TP_EVENT_ADD_RESPONSE":                      16,
	"CLIENT_BATCH_SUBMIT_REQUEST":                100,
	"CLIENT_BATCH_SUBMIT_RESPONSE":               101,
	"CLIENT_BLOCK_LIST_REQUEST":                  102,
	"CLIENT_BLOCK_LIST_RESPONSE":                 103,
	"CLIENT_BLOCK_GET_BY_ID_REQUEST":             104,
	"CLIENT_BLOCK_GET_RESPONSE":                  105,
	"CLIENT_BATCH_LIST_REQUEST":                  106,
	"CLIENT_BATCH_LIST_RESPONSE":                 107,
	"CLIENT_BATCH_GET_REQUEST":                   108,
	"CLIENT_BATCH_GET_RESPONSE":                  109,
	"CLIENT_TRANSACTION_LIST_REQUEST":            110,
	"CLIENT_TRANSACTION_LIST_RESPONSE":           111,
	"CLIENT_TRANSACTION_GET_REQUEST":             112,
	"CLIENT_TRANSACTION_GET_RESPONSE":            113,
	"CLIENT_STATE_CURRENT_REQUEST":               114,
	"CLIENT_STATE_CURRENT_RESPONSE":              115,
	"CLIENT_STATE_LIST_REQUEST":                  116,
	"CLIENT_STATE_LIST_RESPONSE":                 117,
	"CLIENT_STATE_GET_REQUEST":                   118,
	"CLIENT_STATE_GET_RESPONSE":                  119,
	"CLIENT_BATCH_STATUS_REQUEST":                120,
	"CLIENT_BATCH_STATUS_RESPONSE":               121,
	"CLIENT_RECEIPT_GET_REQUEST":                 122,
	"CLIENT_RECEIPT_GET_RESPONSE":                123,
	"CLIENT_BLOCK_GET_BY_NUM_REQUEST":            124,
	"CLIENT_PEERS_GET_REQUEST":                   125,
	"CLIENT_PEERS_GET_RESPONSE":                  126,
	"CLIENT_BLOCK_GET_BY_TRANSACTION_ID_REQUEST": 127,
	"CLIENT_BLOCK_GET_BY_BATCH_ID_REQUEST":       128,
	"CLIENT_STATUS_GET_REQUEST":                  129,
	"CLIENT_STATUS_GET_RESPONSE":                 130,
	"CLIENT_EVENTS_SUBSCRIBE_REQUEST":            500,
	"CLIENT_EVENTS_SUBSCRIBE_RESPONSE":           501,
	"CLIENT_EVENTS_UNSUBSCRIBE_REQUEST":          502,
	"CLIENT_EVENTS_UNSUBSCRIBE_RESPONSE":         503,
	"CLIENT_EVENTS":                              504,
	"CLIENT_EVENTS_GET_REQUEST":                  505,
	"CLIENT_EVENTS_GET_RESPONSE":                 506,
	"GOSSIP_MESSAGE":                             200,
	"GOSSIP_REGISTER":                            201,
	"GOSSIP_UNREGISTER":                          202,
	"GOSSIP_BLOCK_REQUEST":                       205,
	"GOSSIP_BLOCK_RESPONSE":                      206,
	"GOSSIP_BATCH_BY_BATCH_ID_REQUEST":           207,
	"GOSSIP_BATCH_BY_TRANSACTION_ID_REQUEST":     208,
	"GOSSIP_BATCH_RESPONSE":                      209,
	"GOSSIP_GET_PEERS_REQUEST":                   210,
	"GOSSIP_GET_PEERS_RESPONSE":                  211,
	"GOSSIP_CONSENSUS_MESSAGE":                   212,
	"NETWORK_ACK":                                300,
	"NETWORK_CONNECT":                            301,
	"NETWORK_DISCONNECT":                         302,
	"AUTHORIZATION_CONNECTION_RESPONSE":          600,
	"AUTHORIZATION_VIOLATION":                    601,
	"AUTHORIZATION_TRUST_REQUEST":                602,
	"AUTHORIZATION_TRUST_RESPONSE":               603,
	"AUTHORIZATION_CHALLENGE_REQUEST":            604,
	"AUTHORIZATION_CHALLENGE_RESPONSE":           605,
	"AUTHORIZATION_CHALLENGE_SUBMIT":             606,
	"AUTHORIZATION_CHALLENGE_RESULT":             607,
	"PING_REQUEST":                               700,
	"PING_RESPONSE":                              701,
	"CONSENSUS_REGISTER_REQUEST":                 800,
	"CONSENSUS_REGISTER_RESPONSE":                801,
	"CONSENSUS_SEND_TO_REQUEST":                  802,
	"CONSENSUS_SEND_TO_RESPONSE":                 803,
	"CONSENSUS_BROADCAST_REQUEST":                804,
	"CONSENSUS_BROADCAST_RESPONSE":               805,
	"CONSENSUS_INITIALIZE_BLOCK_REQUEST":         806,
	"CONSENSUS_INITIALIZE_BLOCK_RESPONSE":        807,
	"CONSENSUS_FINALIZE_BLOCK_REQUEST":           808,
	"CONSENSUS_FINALIZE_BLOCK_RESPONSE":          809,
	"CONSENSUS_SUMMARIZE_BLOCK_REQUEST":          828,
	"CONSENSUS_SUMMARIZE_BLOCK_RESPONSE":         829,
	"CONSENSUS_CANCEL_BLOCK_REQUEST":             810,
	"CONSENSUS_CANCEL_BLOCK_RESPONSE":            811,
	"CONSENSUS_CHECK_BLOCKS_REQUEST":             812,
	"CONSENSUS_CHECK_BLOCKS_RESPONSE":            813,
	"CONSENSUS_COMMIT_BLOCK_REQUEST":             814,
	"CONSENSUS_COMMIT_BLOCK_RESPONSE":            815,
	"CONSENSUS_IGNORE_BLOCK_REQUEST":             816,
	"CONSENSUS_IGNORE_BLOCK_RESPONSE":            817,
	"CONSENSUS_FAIL_BLOCK_REQUEST":               818,
	"CONSENSUS_FAIL_BLOCK_RESPONSE":              819,
	"CONSENSUS_SETTINGS_GET_REQUEST":             820,
	"CONSENSUS_SETTINGS_GET_RESPONSE":            821,
	"CONSENSUS_STATE_GET_REQUEST":                822,
	"CONSENSUS_STATE_GET_RESPONSE":               823,
	"CONSENSUS_BLOCKS_GET_REQUEST":               824,
	"CONSENSUS_BLOCKS_GET_RESPONSE":              825,
	"CONSENSUS_CHAIN_HEAD_GET_REQUEST":           826,
	"CONSENSUS_CHAIN_HEAD_GET_RESPONSE":          827,
	"CONSENSUS_NOTIFY_PEER_CONNECTED":            900,
	"CONSENSUS_NOTIFY_PEER_DISCONNECTED":         901,
	"CONSENSUS_NOTIFY_PEER_MESSAGE":              902,
	"CONSENSUS_NOTIFY_BLOCK_NEW":                 903,
	"CONSENSUS_NOTIFY_BLOCK_VALID":               904,
	"CONSENSUS_NOTIFY_BLOCK_INVALID":             905,
	"CONSENSUS_NOTIFY_BLOCK_COMMIT":              906,
	"CONSENSUS_NOTIFY_ACK":                       999,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}

func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6829a7b19f8a01ed, []int{1, 0}
}

// A list of messages to be transmitted together.
type MessageList struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MessageList) Reset()         { *m = MessageList{} }
func (m *MessageList) String() string { return proto.CompactTextString(m) }
func (*MessageList) ProtoMessage()    {}
func (*MessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6829a7b19f8a01ed, []int{0}
}

func (m *MessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageList.Unmarshal(m, b)
}
func (m *MessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageList.Marshal(b, m, deterministic)
}
func (m *MessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageList.Merge(m, src)
}
func (m *MessageList) XXX_Size() int {
	return xxx_messageInfo_MessageList.Size(m)
}
func (m *MessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageList.DiscardUnknown(m)
}

var xxx_messageInfo_MessageList proto.InternalMessageInfo

func (m *MessageList) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

// The message passed between the validator and client, containing the
// header fields and content.
type Message struct {
	// The type of message, used to determine how to 'route' the message
	// to the appropriate handler as well as how to deserialize the
	// content.
	MessageType Message_MessageType `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=Message_MessageType" json:"message_type,omitempty"`
	// The identifier used to correlate response messages to their related
	// request messages.  correlation_id should be set to a random string
	// for messages which are not responses to previously sent messages.  For
	// response messages, correlation_id should be set to the same string as
	// contained in the request message.
	CorrelationId string `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// The content of the message, defined by message_type.  In many
	// cases, this data has been serialized with Protocol Buffers or
	// CBOR.
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_6829a7b19f8a01ed, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil {
		return m.MessageType
	}
	return Message_DEFAULT
}

func (m *Message) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *Message) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterType((*MessageList)(nil), "MessageList")
	proto.RegisterType((*Message)(nil), "Message")
}

func init() {
	proto.RegisterFile("protobuf/validator_pb2/validator.proto", fileDescriptor_6829a7b19f8a01ed)
}

var fileDescriptor_6829a7b19f8a01ed = []byte{
	// 1387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x58, 0x49, 0x73, 0xd4, 0x46,
	0x1b, 0xfe, 0xe4, 0x19, 0x34, 0xd0, 0x5e, 0x10, 0x8d, 0xc1, 0x03, 0x78, 0x19, 0x8c, 0xe1, 0x73,
	0x92, 0x2a, 0xa7, 0x0a, 0x0e, 0x39, 0xcb, 0x9a, 0xf6, 0x4c, 0x97, 0x67, 0xa4, 0x89, 0xba, 0x05,
	0x05, 0x17, 0x95, 0xc1, 0x03, 0x38, 0x80, 0xc7, 0xf1, 0x0c, 0x10, 0xb2, 0x27, 0x95, 0xfd, 0x4f,
	0x64, 0x25, 0x2b, 0x90, 0x7d, 0x27, 0xf7, 0x24, 0x55, 0xd9, 0x73, 0xc8, 0x9e, 0x63, 0xfe, 0x40,
	0xf6, 0x53, 0x4a, 0xd2, 0x2b, 0xa9, 0xbb, 0xa5, 0xf1, 0xc9, 0xa3, 0x7e, 0x9e, 0xf7, 0x79, 0xd7,
	0xee, 0x96, 0x85, 0x0e, 0xac, 0xad, 0x77, 0x7a, 0x9d, 0xe3, 0xe7, 0x4f, 0xde, 0x7c, 0x61, 0xe9,
	0xec, 0xca, 0xf2, 0x52, 0xaf, 0xb3, 0xee, 0xaf, 0x1d, 0x3f, 0x98, 0x3e, 0xcd, 0x85, 0x84, 0xe9,
	0x43, 0x68, 0xb0, 0xd9, 0xee, 0x76, 0x97, 0x4e, 0xb5, 0x1b, 0x2b, 0xdd, 0x1e, 0x9e, 0x41, 0x9b,
	0xcf, 0x45, 0x8f, 0xdd, 0xb2, 0x56, 0x29, 0xcc, 0x0e, 0x1e, 0xdc, 0x3c, 0x07, 0xb8, 0x9b, 0x20,
	0xd3, 0x97, 0xa7, 0x50, 0x09, 0x56, 0xf1, 0x2d, 0x68, 0x08, 0xd6, 0xfd, 0xde, 0xa5, 0xb5, 0x76,
	0x59, 0xab, 0x68, 0xb3, 0x23, 0x07, 0x47, 0x63, 0xab, 0xf8, 0x2f, 0xbf, 0xb4, 0xd6, 0x76, 0x07,
	0xcf, 0xa5, 0x0f, 0x78, 0x3f, 0x1a, 0x39, 0xd1, 0x59, 0x5f, 0x6f, 0x9f, 0x5d, 0xea, 0xad, 0x74,
	0x56, 0xfd, 0x95, 0xe5, 0xf2, 0x40, 0x45, 0x9b, 0xdd, 0xe2, 0x0e, 0x0b, 0xab, 0x74, 0x19, 0x97,
	0x51, 0xe9, 0x44, 0x67, 0xb5, 0xd7, 0x5e, 0xed, 0x95, 0x0b, 0x15, 0x6d, 0x76, 0xc8, 0x8d, 0x1f,
	0xa7, 0x9f, 0x9c, 0x4c, 0x62, 0x0f, 0x05, 0x07, 0x51, 0xa9, 0x4a, 0x16, 0x4c, 0xaf, 0xc1, 0x8d,
	0xff, 0xe1, 0x31, 0xb4, 0x9d, 0xb7, 0x7c, 0x97, 0xd4, 0x28, 0xe3, 0xc4, 0xf5, 0x5d, 0x72, 0xab,
	0x47, 0x18, 0x37, 0x34, 0x5c, 0x46, 0xa3, 0x32, 0xc0, 0x5a, 0x8e, 0xcd, 0x88, 0x31, 0x80, 0x77,
	0xa1, 0x1d, 0xbc, 0xe5, 0x7b, 0x76, 0xc6, 0xa8, 0x80, 0x77, 0xa3, 0x9d, 0x2a, 0x04, 0x66, 0x45,
	0xbc, 0x13, 0x61, 0xde, 0xf2, 0x5b, 0xae, 0x63, 0x11, 0xc6, 0x12, 0x9b, 0x4d, 0x10, 0x41, 0xba,
	0x0e, 0x06, 0x3a, 0x44, 0xc0, 0xb8, 0xc9, 0x89, 0x5f, 0x23, 0x3c, 0x31, 0x29, 0x41, 0x04, 0x22,
	0x02, 0x46, 0x9b, 0x25, 0x23, 0x26, 0x18, 0x6d, 0x91, 0x8c, 0x98, 0x68, 0x84, 0xf0, 0x1e, 0x34,
	0x96, 0x40, 0x55, 0xd2, 0x20, 0x9c, 0x24, 0x76, 0x83, 0x78, 0x1c, 0x95, 0xb3, 0x20, 0x98, 0x0e,
	0xe1, 0x29, 0xb4, 0x27, 0x2c, 0x93, 0x45, 0x68, 0x8b, 0xfb, 0x66, 0xb5, 0xea, 0x57, 0x4d, 0x6e,
	0x26, 0xe6, 0xc3, 0xb8, 0x82, 0xc6, 0xf3, 0x09, 0x20, 0x31, 0x02, 0x21, 0x93, 0xc3, 0xc4, 0x8e,
	0xf0, 0xd8, 0x76, 0x2b, 0x84, 0x2c, 0x22, 0x60, 0x64, 0x04, 0x7e, 0xad, 0x06, 0x0d, 0x80, 0x79,
	0x93, 0x5b, 0x75, 0x9f, 0x79, 0xf3, 0x4d, 0x9a, 0xa6, 0xbb, 0x1c, 0xf8, 0xcd, 0x27, 0x80, 0x44,
	0x1b, 0x4f, 0xa0, 0x5d, 0x31, 0xa3, 0xe1, 0x58, 0x8b, 0x7e, 0x83, 0xb2, 0x54, 0xe0, 0x24, 0x9e,
	0x44, 0xbb, 0xf3, 0x60, 0x30, 0x3f, 0x85, 0xa7, 0xd1, 0xa4, 0x84, 0x07, 0x8d, 0x98, 0x3f, 0xea,
	0xd3, 0x34, 0x81, 0xd3, 0x19, 0x17, 0x52, 0xb3, 0x56, 0x44, 0x38, 0x8c, 0x51, 0x8a, 0xe0, 0x36,
	0x31, 0x02, 0x11, 0x06, 0xf3, 0x33, 0x41, 0x67, 0x24, 0x5c, 0x1c, 0x92, 0xb3, 0x19, 0x71, 0xc9,
	0xf7, 0x39, 0xbc, 0x0f, 0x4d, 0x01, 0xcc, 0x5d, 0xd3, 0x66, 0xa6, 0xc5, 0xa9, 0x63, 0xcb, 0x11,
	0xac, 0xe2, 0x19, 0x54, 0xe9, 0x4f, 0x02, 0xa9, 0x8e, 0x50, 0x09, 0x91, 0x25, 0x46, 0xb3, 0xd6,
	0xc7, 0x9d, 0x14, 0xd3, 0xed, 0x42, 0xcf, 0xa2, 0x71, 0xb3, 0x3c, 0xd7, 0x0d, 0x9e, 0x62, 0x99,
	0x75, 0xbc, 0x17, 0x4d, 0xf4, 0x61, 0x80, 0x48, 0x57, 0xc8, 0x3b, 0xa2, 0x48, 0x29, 0xf5, 0x84,
	0xa2, 0x4a, 0x30, 0x98, 0x9f, 0x17, 0x8a, 0x9a, 0xdd, 0x79, 0x17, 0x32, 0xe2, 0x52, 0x02, 0x17,
	0xb3, 0x53, 0xc9, 0x4d, 0xee, 0xa5, 0x9b, 0xfd, 0x8e, 0xec, 0x54, 0xc6, 0x04, 0x90, 0xb8, 0x24,
	0xc4, 0x17, 0xef, 0x19, 0x31, 0x82, 0x3b, 0x05, 0x17, 0x32, 0x0e, 0x02, 0x77, 0x09, 0x95, 0x96,
	0xe6, 0xd2, 0xf6, 0x9a, 0x89, 0xca, 0xdd, 0x42, 0x96, 0x2d, 0x42, 0x5c, 0x26, 0xf9, 0xb8, 0x47,
	0xc8, 0x52, 0x44, 0xc1, 0xc3, 0xbd, 0x78, 0x0e, 0xdd, 0x98, 0xe7, 0x41, 0x6c, 0xac, 0xb0, 0x0b,
	0xee, 0xc3, 0x37, 0xa0, 0x99, 0x3c, 0x7e, 0x54, 0x01, 0x81, 0x79, 0xbf, 0x86, 0x27, 0xa5, 0xfa,
	0x7a, 0x72, 0x60, 0x0f, 0x68, 0x78, 0x4a, 0xea, 0x9e, 0xa7, 0x84, 0xf6, 0xa0, 0x86, 0x67, 0x92,
	0xec, 0xc3, 0x63, 0x83, 0x05, 0xfb, 0x9e, 0x59, 0x2e, 0x9d, 0x4f, 0x8f, 0xb4, 0xdf, 0x0b, 0x78,
	0x7f, 0x32, 0xd7, 0x39, 0x2c, 0x10, 0xfb, 0xa3, 0x80, 0x0f, 0xa0, 0xbd, 0x32, 0xcd, 0xb3, 0xb3,
	0x72, 0x7f, 0x16, 0xf0, 0xff, 0xd1, 0xf4, 0x46, 0x3c, 0x10, 0xfc, 0xab, 0x80, 0x31, 0x1a, 0x96,
	0x88, 0xc6, 0xdf, 0x05, 0x21, 0x65, 0x30, 0x16, 0x53, 0xfe, 0xa7, 0x20, 0xa4, 0x2c, 0xe1, 0x20,
	0xfa, 0x6f, 0x01, 0x6f, 0x47, 0x23, 0x35, 0x87, 0x31, 0xda, 0xf2, 0x9b, 0x84, 0x31, 0xb3, 0x46,
	0x8c, 0x8f, 0x35, 0x3c, 0x8a, 0xb6, 0xc2, 0x62, 0x7c, 0x17, 0x19, 0x9f, 0x68, 0x78, 0x27, 0xda,
	0x06, 0xab, 0xe9, 0x1d, 0x65, 0x7c, 0xaa, 0xe1, 0x5d, 0x68, 0x14, 0xd6, 0xa3, 0x0e, 0xc5, 0xee,
	0x3f, 0xd3, 0xf0, 0x6e, 0xb4, 0x43, 0x81, 0xc0, 0xf3, 0xe7, 0x5a, 0x50, 0xc6, 0x18, 0x0b, 0x7b,
	0x99, 0xd7, 0xd4, 0x2f, 0x34, 0x7c, 0x13, 0x3a, 0xa0, 0xd2, 0xfa, 0xcc, 0xca, 0x97, 0x92, 0xbf,
	0x90, 0x9c, 0xf8, 0xfb, 0x4a, 0xc3, 0x13, 0xa8, 0x0c, 0x58, 0x50, 0x83, 0x68, 0x36, 0x63, 0xd3,
	0xaf, 0xc3, 0xe1, 0xc9, 0x81, 0xc1, 0xfc, 0x1b, 0xd1, 0xdc, 0x0a, 0x96, 0x6c, 0xe6, 0xb1, 0xa4,
	0x64, 0xdf, 0x6a, 0xd8, 0x40, 0x83, 0x36, 0xe1, 0x47, 0x1c, 0x77, 0xd1, 0x37, 0xad, 0x45, 0xe3,
	0xca, 0x40, 0x50, 0xc4, 0x78, 0xc5, 0x72, 0x6c, 0x9b, 0x58, 0xdc, 0xb8, 0x3a, 0x80, 0xc7, 0x10,
	0x8e, 0x57, 0xab, 0x94, 0xc5, 0xc0, 0xb5, 0x81, 0x60, 0x5c, 0x4c, 0x8f, 0xd7, 0x1d, 0x97, 0x1e,
	0x33, 0xc3, 0xcc, 0x00, 0x0b, 0x7e, 0x26, 0x71, 0x7c, 0x57, 0xc4, 0xe3, 0x68, 0x4c, 0xe6, 0x1d,
	0xa6, 0x4e, 0x23, 0xfc, 0x65, 0x7c, 0x5f, 0xc4, 0x15, 0xb4, 0x47, 0x46, 0xb9, 0xeb, 0x09, 0x27,
	0xd8, 0x0f, 0x45, 0xbc, 0x17, 0x8d, 0xe7, 0x33, 0xc0, 0xc5, 0x8f, 0xc5, 0x60, 0x1b, 0x28, 0xa1,
	0xd4, 0xcd, 0x46, 0x83, 0xd8, 0xb5, 0x74, 0x6e, 0x7f, 0x2a, 0x06, 0xfd, 0xeb, 0xcf, 0x02, 0xb1,
	0x9f, 0x8b, 0x78, 0x1f, 0x9a, 0xec, 0x47, 0x8b, 0x6e, 0x55, 0xe3, 0x97, 0x0d, 0x49, 0x2e, 0x61,
	0xc1, 0xcb, 0xd6, 0xaf, 0x45, 0xbc, 0x0d, 0x0d, 0xb5, 0xa8, 0x5d, 0x4b, 0x62, 0xb8, 0xbe, 0x29,
	0xd8, 0x12, 0xb0, 0x04, 0x0e, 0x3f, 0xda, 0x14, 0x8e, 0x7c, 0xd2, 0xa1, 0xcc, 0x6b, 0xd6, 0x53,
	0x7a, 0x50, 0xa3, 0x5c, 0x02, 0x48, 0x3c, 0xad, 0x87, 0xbb, 0x2a, 0x61, 0x30, 0x62, 0x57, 0x7d,
	0xee, 0x24, 0x0a, 0xcf, 0xe8, 0xb2, 0x8b, 0x14, 0x07, 0x81, 0x67, 0x15, 0x17, 0xf3, 0xae, 0x63,
	0x56, 0x2d, 0x53, 0x68, 0xc3, 0x73, 0x7a, 0xd0, 0x86, 0x7c, 0x06, 0x88, 0x5c, 0xd6, 0xc3, 0x83,
	0x21, 0xa1, 0x50, 0x9b, 0x72, 0x6a, 0x36, 0xe8, 0x31, 0xa2, 0xec, 0xb2, 0xe7, 0x75, 0x3c, 0x8b,
	0xf6, 0x6d, 0x48, 0x04, 0xc9, 0x17, 0xf4, 0xf0, 0xe8, 0x4a, 0x98, 0x0b, 0xd4, 0xce, 0x13, 0x7c,
	0x51, 0x0f, 0x8f, 0xae, 0xfe, 0x34, 0x90, 0x7b, 0x49, 0xe1, 0x31, 0xaf, 0xd9, 0x34, 0xdd, 0xac,
	0xde, 0x75, 0x25, 0x93, 0x2c, 0x2f, 0xee, 0x9d, 0x1e, 0xcc, 0x41, 0x4a, 0xb4, 0x4c, 0xdb, 0x22,
	0x0d, 0x45, 0xed, 0x65, 0x3d, 0x3c, 0xa5, 0xfb, 0x91, 0x40, 0xea, 0x15, 0x55, 0xaa, 0x4e, 0xac,
	0xc5, 0x88, 0x94, 0x6e, 0xfa, 0x2b, 0xaa, 0x94, 0x4c, 0x02, 0xa9, 0xab, 0xaa, 0x94, 0xd3, 0x0c,
	0x5e, 0x05, 0xe5, 0xa8, 0xae, 0xa9, 0x52, 0x32, 0x09, 0xa4, 0x5e, 0x55, 0xa4, 0x68, 0xcd, 0x76,
	0x5c, 0xb5, 0x5c, 0xaf, 0x29, 0x52, 0x0a, 0x09, 0xa4, 0x5e, 0x57, 0x26, 0x68, 0xc1, 0xa4, 0x6a,
	0xa5, 0xde, 0xd0, 0xf1, 0x34, 0x9a, 0xe8, 0x43, 0x01, 0x99, 0x37, 0x95, 0x88, 0x18, 0xe1, 0x9c,
	0xda, 0x35, 0xf9, 0x1a, 0x79, 0x4b, 0x89, 0x48, 0x21, 0x81, 0xd4, 0xdb, 0xca, 0xd4, 0x67, 0x5f,
	0x80, 0xde, 0x51, 0x62, 0xce, 0x79, 0x09, 0x7a, 0x57, 0xdd, 0x18, 0x51, 0x33, 0x44, 0x95, 0xf7,
	0x94, 0xb4, 0x24, 0x0a, 0xc8, 0xbc, 0xaf, 0x4c, 0xba, 0x55, 0x37, 0xa9, 0xed, 0xd7, 0x89, 0x59,
	0x95, 0xa4, 0x3e, 0x50, 0x26, 0x38, 0x43, 0x03, 0xb9, 0x0f, 0x95, 0x02, 0xd8, 0x0e, 0xa7, 0x0b,
	0x47, 0xc3, 0x3b, 0x22, 0x3e, 0xa5, 0x49, 0xd5, 0x78, 0xa8, 0x24, 0xcf, 0xb9, 0xc8, 0x4a, 0x8f,
	0x7a, 0x52, 0x35, 0x1e, 0x2e, 0xc9, 0x19, 0x88, 0xc4, 0xf8, 0x46, 0x79, 0xa4, 0x24, 0x1f, 0x32,
	0xc0, 0x89, 0xda, 0x67, 0x93, 0x23, 0xc6, 0xa3, 0x25, 0xb9, 0x52, 0x12, 0xe1, 0xb0, 0xd9, 0xa0,
	0x55, 0xe3, 0xb1, 0x92, 0xdc, 0x5c, 0x89, 0x42, 0xed, 0x88, 0xf4, 0x78, 0x7e, 0x30, 0x11, 0x29,
	0x1a, 0x63, 0xe3, 0x89, 0xe0, 0x9f, 0xc6, 0xd1, 0x0c, 0x27, 0xb8, 0xe7, 0x7e, 0x2b, 0xcd, 0xcf,
	0xa2, 0x1d, 0xdd, 0xa5, 0x8b, 0xbd, 0x4e, 0xa7, 0x77, 0x7a, 0xae, 0xbb, 0x7c, 0x66, 0x2e, 0xfe,
	0x26, 0xd0, 0xd2, 0x8e, 0x0d, 0x4b, 0x9f, 0x05, 0x8e, 0xeb, 0x21, 0x74, 0xe8, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd8, 0x39, 0x9a, 0xf3, 0x37, 0x10, 0x00, 0x00,
}
