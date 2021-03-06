// Code generated by protoc-gen-go.
// source: transport.proto
// DO NOT EDIT!

/*
Package transport is a generated protocol buffer package.

It is generated from these files:
	transport.proto

It has these top-level messages:
	Message
*/
package transport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Message is the envelope/wrapper for all messages.
type Message struct {
	// ID is a globally unique message.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Timestamp is the timestamp in nanoseconds with the message was published.
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// Payload is the actual payload being published that will be consumed.
	// This is expected to use protobuf encoding.
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// Deprecated. Use status instead.
	Error string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	// Cause is the ID of a message that resulted in this message being published.
	Cause string `protobuf:"bytes,5,opt,name=cause" json:"cause,omitempty"`
	// Subject is the subject the message is published.
	Subject string `protobuf:"bytes,6,opt,name=subject" json:"subject,omitempty"`
	// Queue is the queue this message was received by.
	Queue string `protobuf:"bytes,7,opt,name=queue" json:"queue,omitempty"`
	// Reply is the inbox ID of the publisher of this message.
	Reply string `protobuf:"bytes,8,opt,name=reply" json:"reply,omitempty"`
	// Error that occurs in during transport or in the application.
	Status *google_rpc.Status `protobuf:"bytes,9,opt,name=status" json:"status,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Message) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

func (m *Message) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Message) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *Message) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func (m *Message) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "transport.Message")
}

func init() { proto.RegisterFile("transport.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0xdd, 0x6d, 0x6d, 0x14, 0x85, 0x20, 0x38, 0x88, 0x87, 0xe2, 0xa9, 0x78, 0xe8,
	0x82, 0x3e, 0x87, 0x97, 0xfa, 0x04, 0xd9, 0x74, 0x28, 0x95, 0xae, 0x89, 0x33, 0x93, 0xc3, 0x3e,
	0xb5, 0xaf, 0x20, 0x49, 0xac, 0x1e, 0xbf, 0x6f, 0x3e, 0x06, 0x7e, 0x7d, 0x2b, 0x64, 0x3f, 0x39,
	0x78, 0x92, 0x21, 0x90, 0x17, 0x6f, 0xda, 0x3f, 0xf1, 0x70, 0x3f, 0x7b, 0x3f, 0xaf, 0x78, 0xa0,
	0xe0, 0x0e, 0x2c, 0x56, 0x22, 0x97, 0xe6, 0xe9, 0x5b, 0xe9, 0xe6, 0x0d, 0x99, 0xed, 0x8c, 0xe6,
	0x46, 0x57, 0xcb, 0x04, 0xaa, 0x53, 0x7d, 0x3b, 0x56, 0xcb, 0x64, 0x1e, 0x75, 0x2b, 0xcb, 0x09,
	0x59, 0xec, 0x29, 0x40, 0xd5, 0xa9, 0x7e, 0x37, 0xfe, 0x0b, 0x03, 0xba, 0x09, 0xf6, 0xbc, 0x7a,
	0x3b, 0xc1, 0x45, 0xa7, 0xfa, 0xeb, 0x71, 0x43, 0x73, 0xa7, 0xf7, 0x48, 0xe4, 0x09, 0x76, 0xf9,
	0x55, 0x81, 0x64, 0x9d, 0x8d, 0x8c, 0xb0, 0x2f, 0x36, 0x43, 0xfa, 0xc2, 0xf1, 0xf8, 0x81, 0x4e,
	0xa0, 0xce, 0x7e, 0xc3, 0xd4, 0x7f, 0x45, 0x8c, 0x08, 0x4d, 0xe9, 0x33, 0x24, 0x4b, 0x18, 0xd6,
	0x33, 0x5c, 0x16, 0x9b, 0xc1, 0x3c, 0xeb, 0xba, 0xac, 0x82, 0xb6, 0x53, 0xfd, 0xd5, 0x8b, 0x19,
	0xca, 0xde, 0x81, 0x82, 0x1b, 0xde, 0xf3, 0x65, 0xfc, 0x2d, 0x8e, 0x75, 0x1e, 0xfe, 0xfa, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0x4d, 0x61, 0xd5, 0x2f, 0x01, 0x00, 0x00,
}
