// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/post-publish.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/post-publish.proto

It has these top-level messages:
	PublishPostMessage
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type PublishPostMessage struct {
	Title   string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *PublishPostMessage) Reset()                    { *m = PublishPostMessage{} }
func (m *PublishPostMessage) String() string            { return proto1.CompactTextString(m) }
func (*PublishPostMessage) ProtoMessage()               {}
func (*PublishPostMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PublishPostMessage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PublishPostMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto1.RegisterType((*PublishPostMessage)(nil), "mycodesmells.golangexamples.nats.proto.PublishPostMessage")
}

func init() { proto1.RegisterFile("proto/post-publish.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcd, 0x31, 0xeb, 0xc2, 0x30,
	0x10, 0x05, 0x70, 0xfa, 0x87, 0xbf, 0x62, 0xc6, 0xe0, 0x90, 0x51, 0x1c, 0xc4, 0xa5, 0x09, 0xa8,
	0x9f, 0x40, 0x5c, 0x85, 0xe2, 0xe8, 0x96, 0xd6, 0x23, 0x0d, 0x24, 0xb9, 0xe0, 0x5d, 0x41, 0xbf,
	0xbd, 0x98, 0x50, 0x70, 0x3a, 0xde, 0xe3, 0x1e, 0x3f, 0xa1, 0xf2, 0x13, 0x19, 0x4d, 0x46, 0xe2,
	0x36, 0x4f, 0x7d, 0xf0, 0x34, 0xea, 0x52, 0xc9, 0x5d, 0x7c, 0x0f, 0xf8, 0x00, 0x8a, 0x10, 0x02,
	0x69, 0x87, 0xc1, 0x26, 0x07, 0x2f, 0x1b, 0x73, 0x00, 0xd2, 0xc9, 0x32, 0xd5, 0xbf, 0xed, 0x45,
	0xc8, 0xae, 0x0e, 0x3b, 0x24, 0xbe, 0x02, 0x91, 0x75, 0x20, 0xd7, 0xe2, 0x9f, 0x3d, 0x07, 0x50,
	0xcd, 0xa6, 0xd9, 0xaf, 0x6e, 0x35, 0x48, 0x25, 0x96, 0x03, 0x26, 0x86, 0xc4, 0xea, 0xaf, 0xf4,
	0x73, 0x3c, 0x9f, 0xee, 0x07, 0xe7, 0x79, 0x9c, 0x7a, 0x3d, 0x60, 0x34, 0xbf, 0xb4, 0xa9, 0x74,
	0x3b, 0xdb, 0xe6, 0x6b, 0x9b, 0x62, 0xf7, 0x8b, 0x72, 0x8e, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x39, 0x0a, 0x1a, 0xc4, 0xc6, 0x00, 0x00, 0x00,
}
