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
	DelayNanos           int64    `protobuf:"varint,4,opt,name=delayNanos,proto3" json:"delayNanos,omitempty"`
	ErrorPercent         int32    `protobuf:"varint,5,opt,name=errorPercent,proto3" json:"errorPercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ping_c98dcc1e6f531dd9, []int{0}
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

func (m *PingMessage) GetDelayNanos() int64 {
	if m != nil {
		return m.DelayNanos
	}
	return 0
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

func init() { proto.RegisterFile("ping.proto", fileDescriptor_ping_c98dcc1e6f531dd9) }

var fileDescriptor_ping_c98dcc1e6f531dd9 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xcf, 0x8a, 0xc2, 0x30,
	0x10, 0x87, 0x37, 0xfd, 0xb3, 0xcb, 0xce, 0x2e, 0xcb, 0x32, 0xa7, 0xe0, 0x41, 0x4a, 0x4f, 0x3d,
	0x15, 0xd1, 0xbb, 0x6f, 0xa0, 0x94, 0xfa, 0x04, 0xb1, 0x1d, 0x43, 0xa1, 0x24, 0x71, 0x12, 0x84,
	0x3e, 0x84, 0xef, 0x2c, 0x0d, 0x08, 0x15, 0xbc, 0xcd, 0xf7, 0xc1, 0xc0, 0xf7, 0x03, 0x70, 0x83,
	0xd1, 0xb5, 0x63, 0x1b, 0x2c, 0xe6, 0x17, 0xcd, 0xae, 0x2b, 0xef, 0x02, 0x7e, 0x9a, 0xc1, 0xe8,
	0x03, 0x79, 0xaf, 0x34, 0xe1, 0x3f, 0xa4, 0x9e, 0xae, 0x52, 0x14, 0xa2, 0x4a, 0xdb, 0xf9, 0xc4,
	0x3f, 0x48, 0x82, 0x97, 0x49, 0x14, 0x49, 0xf0, 0x28, 0xe1, 0xcb, 0xa9, 0x69, 0xb4, 0xaa, 0x97,
	0x69, 0x21, 0xaa, 0xef, 0xf6, 0x89, 0xb8, 0x06, 0xe8, 0x69, 0x54, 0xd3, 0x51, 0x19, 0xeb, 0x65,
	0x16, 0x3f, 0x16, 0x06, 0x4b, 0xf8, 0x25, 0x66, 0xcb, 0x0d, 0x71, 0x47, 0x26, 0xc8, 0xbc, 0x10,
	0x55, 0xde, 0xbe, 0xb8, 0xed, 0x1e, 0x60, 0xce, 0x39, 0x11, 0xdf, 0x88, 0x71, 0x03, 0xd9, 0x4c,
	0x88, 0x75, 0xac, 0xad, 0x17, 0xa5, 0xab, 0x37, 0xae, 0xfc, 0x38, 0x7f, 0xc6, 0x75, 0xbb, 0x47,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x5d, 0x69, 0x08, 0xeb, 0x00, 0x00, 0x00,
}
