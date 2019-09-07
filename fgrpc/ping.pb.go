// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ping.proto

package fgrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PingMessage struct {
	Seq                  int64    `protobuf:"varint,1,opt,name=seq,proto3" json:"seq,omitempty"`
	Ts                   int64    `protobuf:"varint,2,opt,name=ts,proto3" json:"ts,omitempty"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Delay                string   `protobuf:"bytes,4,opt,name=delay,proto3" json:"delay,omitempty"`
	ErrorPercent         int32    `protobuf:"varint,5,opt,name=errorPercent,proto3" json:"errorPercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ping_47e7f6f1b61f822a, []int{0}
}
func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (dst *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(dst, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetSeq() int64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *PingMessage) GetTs() int64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *PingMessage) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PingMessage) GetDelay() string {
	if m != nil {
		return m.Delay
	}
	return ""
}

func (m *PingMessage) GetErrorPercent() int32 {
	if m != nil {
		return m.ErrorPercent
	}
	return 0
}

func init() {
	proto.RegisterType((*PingMessage)(nil), "fgrpc.PingMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingServerClient is the client API for PingServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingServerClient interface {
	Ping(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessage, error)
}

type pingServerClient struct {
	cc *grpc.ClientConn
}

func NewPingServerClient(cc *grpc.ClientConn) PingServerClient {
	return &pingServerClient{cc}
}

func (c *pingServerClient) Ping(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessage, error) {
	out := new(PingMessage)
	err := c.cc.Invoke(ctx, "/fgrpc.PingServer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServerServer is the server API for PingServer service.
type PingServerServer interface {
	Ping(context.Context, *PingMessage) (*PingMessage, error)
}

func RegisterPingServerServer(s *grpc.Server, srv PingServerServer) {
	s.RegisterService(&_PingServer_serviceDesc, srv)
}

func _PingServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fgrpc.PingServer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServerServer).Ping(ctx, req.(*PingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fgrpc.PingServer",
	HandlerType: (*PingServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingServer_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}

func init() { proto.RegisterFile("ping.proto", fileDescriptor_ping_47e7f6f1b61f822a) }

var fileDescriptor_ping_47e7f6f1b61f822a = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x8a, 0xc2, 0x30,
	0x14, 0x45, 0x27, 0x6d, 0x33, 0xc3, 0xbc, 0x19, 0x44, 0x1e, 0x2e, 0x82, 0xab, 0xd2, 0x55, 0x57,
	0x45, 0x74, 0xef, 0x1f, 0x08, 0xa5, 0x7e, 0x41, 0x6c, 0x9f, 0xa1, 0x50, 0x9a, 0xf8, 0x12, 0x84,
	0xae, 0xfd, 0x71, 0x69, 0x44, 0x50, 0x70, 0x77, 0xcf, 0x59, 0x9d, 0x0b, 0xe0, 0xfa, 0xd1, 0x54,
	0x8e, 0x6d, 0xb0, 0x28, 0xcf, 0x86, 0x5d, 0x5b, 0xdc, 0x04, 0xfc, 0xd5, 0xfd, 0x68, 0x0e, 0xe4,
	0xbd, 0x36, 0x84, 0x4b, 0x48, 0x3d, 0x5d, 0x94, 0xc8, 0x45, 0x99, 0x36, 0xf3, 0xc4, 0x05, 0x24,
	0xc1, 0xab, 0x24, 0x8a, 0x24, 0x78, 0x54, 0xf0, 0xe3, 0xf4, 0x34, 0x58, 0xdd, 0xa9, 0x34, 0x17,
	0xe5, 0x6f, 0xf3, 0x44, 0x5c, 0x81, 0xec, 0x68, 0xd0, 0x93, 0xca, 0xa2, 0x7f, 0x00, 0x16, 0xf0,
	0x4f, 0xcc, 0x96, 0x6b, 0xe2, 0x96, 0xc6, 0xa0, 0x64, 0x2e, 0x4a, 0xd9, 0xbc, 0xb9, 0xed, 0x1e,
	0x60, 0x8e, 0x38, 0x12, 0x5f, 0x89, 0x71, 0x03, 0xd9, 0x4c, 0x88, 0x55, 0x6c, 0xac, 0x5e, 0xfa,
	0xd6, 0x1f, 0x5c, 0xf1, 0x75, 0xfa, 0x8e, 0x9f, 0x76, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69,
	0xd4, 0x9d, 0xb0, 0xe1, 0x00, 0x00, 0x00,
}
