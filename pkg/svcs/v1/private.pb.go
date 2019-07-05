// Code generated by protoc-gen-go. DO NOT EDIT.
// source: svcs/v1/private.proto

package svcsv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/octofoxio/sparkle/pkg/common/v1"
	v11 "github.com/octofoxio/sparkle/pkg/entities/v1"
	grpc "google.golang.org/grpc"
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

type GetUserByAccessTokenInput struct {
	ID                   *v1.String `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetUserByAccessTokenInput) Reset()         { *m = GetUserByAccessTokenInput{} }
func (m *GetUserByAccessTokenInput) String() string { return proto.CompactTextString(m) }
func (*GetUserByAccessTokenInput) ProtoMessage()    {}
func (*GetUserByAccessTokenInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb17c849bb4215b6, []int{0}
}

func (m *GetUserByAccessTokenInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByAccessTokenInput.Unmarshal(m, b)
}
func (m *GetUserByAccessTokenInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByAccessTokenInput.Marshal(b, m, deterministic)
}
func (m *GetUserByAccessTokenInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByAccessTokenInput.Merge(m, src)
}
func (m *GetUserByAccessTokenInput) XXX_Size() int {
	return xxx_messageInfo_GetUserByAccessTokenInput.Size(m)
}
func (m *GetUserByAccessTokenInput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByAccessTokenInput.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByAccessTokenInput proto.InternalMessageInfo

func (m *GetUserByAccessTokenInput) GetID() *v1.String {
	if m != nil {
		return m.ID
	}
	return nil
}

type GetUserByAccessTokenOutput struct {
	Result               *v11.User `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetUserByAccessTokenOutput) Reset()         { *m = GetUserByAccessTokenOutput{} }
func (m *GetUserByAccessTokenOutput) String() string { return proto.CompactTextString(m) }
func (*GetUserByAccessTokenOutput) ProtoMessage()    {}
func (*GetUserByAccessTokenOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb17c849bb4215b6, []int{1}
}

func (m *GetUserByAccessTokenOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByAccessTokenOutput.Unmarshal(m, b)
}
func (m *GetUserByAccessTokenOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByAccessTokenOutput.Marshal(b, m, deterministic)
}
func (m *GetUserByAccessTokenOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByAccessTokenOutput.Merge(m, src)
}
func (m *GetUserByAccessTokenOutput) XXX_Size() int {
	return xxx_messageInfo_GetUserByAccessTokenOutput.Size(m)
}
func (m *GetUserByAccessTokenOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByAccessTokenOutput.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByAccessTokenOutput proto.InternalMessageInfo

func (m *GetUserByAccessTokenOutput) GetResult() *v11.User {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserByAccessTokenInput)(nil), "sparkle.svcs.v1.GetUserByAccessTokenInput")
	proto.RegisterType((*GetUserByAccessTokenOutput)(nil), "sparkle.svcs.v1.GetUserByAccessTokenOutput")
}

func init() { proto.RegisterFile("svcs/v1/private.proto", fileDescriptor_fb17c849bb4215b6) }

var fileDescriptor_fb17c849bb4215b6 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4b, 0x3b, 0x31,
	0x10, 0xc6, 0xff, 0xbb, 0x7f, 0xec, 0x21, 0x0a, 0x85, 0xf5, 0x05, 0xbb, 0x27, 0xe9, 0xc9, 0x17,
	0x48, 0x48, 0x3d, 0x7a, 0xb2, 0x14, 0xa5, 0xa7, 0x96, 0x8d, 0x16, 0x91, 0x5e, 0xda, 0x10, 0xd7,
	0xb0, 0xdd, 0x4d, 0x4c, 0x66, 0x83, 0x7a, 0xf2, 0xb3, 0x78, 0xf4, 0xa3, 0xf8, 0xa9, 0x24, 0xdd,
	0xd4, 0x43, 0xad, 0xe0, 0x29, 0x30, 0x79, 0x9e, 0xdf, 0xcc, 0x3c, 0x83, 0xf6, 0xad, 0xe3, 0x96,
	0x38, 0x4a, 0xb4, 0x91, 0x6e, 0x06, 0x02, 0x6b, 0xa3, 0x40, 0x25, 0x6d, 0xab, 0x67, 0xa6, 0x58,
	0x08, 0xec, 0xbf, 0xb1, 0xa3, 0x69, 0x87, 0xab, 0xb2, 0x54, 0x55, 0x50, 0x96, 0x12, 0xa4, 0x0b,
	0xda, 0xf4, 0x40, 0x54, 0x20, 0x41, 0x8a, 0x25, 0xa6, 0xb6, 0xc2, 0x34, 0xf5, 0xee, 0x15, 0xea,
	0x5c, 0x0b, 0xb8, 0xb5, 0xc2, 0xf4, 0x5f, 0x2e, 0x39, 0x17, 0xd6, 0xde, 0xa8, 0x42, 0x54, 0xc3,
	0x4a, 0xd7, 0x90, 0x9c, 0xa0, 0x78, 0x38, 0x38, 0x8c, 0x8e, 0xa2, 0xe3, 0xed, 0x5e, 0x07, 0xaf,
	0xba, 0x35, 0x4d, 0xb0, 0xa3, 0x98, 0x81, 0x91, 0x55, 0x9e, 0xc5, 0xc3, 0x41, 0x77, 0x84, 0xd2,
	0x4d, 0x9c, 0x51, 0x0d, 0x1e, 0x44, 0x51, 0x2b, 0x13, 0xb6, 0x5e, 0xc0, 0x0f, 0xd8, 0x6a, 0x2c,
	0x8f, 0xf3, 0xee, 0x2c, 0x08, 0x7b, 0xaf, 0x68, 0x8b, 0x69, 0x59, 0x88, 0xe4, 0x09, 0xed, 0x6d,
	0x22, 0x27, 0xa7, 0x78, 0x6d, 0x7d, 0xfc, 0xeb, 0x22, 0xe9, 0xd9, 0x9f, 0xb4, 0xcd, 0xb0, 0xdd,
	0x7f, 0xfd, 0xb7, 0x08, 0xed, 0x72, 0x55, 0xae, 0x9b, 0xfa, 0x3b, 0xe3, 0x26, 0xff, 0xb1, 0x8f,
	0x6e, 0x1c, 0xdd, 0x93, 0x5c, 0xc2, 0x63, 0x3d, 0xf7, 0x89, 0x10, 0xc5, 0x41, 0x3d, 0xa8, 0x67,
	0xa9, 0x48, 0x70, 0x11, 0x5d, 0xe4, 0x24, 0x1c, 0xee, 0xc2, 0xbf, 0x8e, 0xbe, 0xc7, 0xff, 0x19,
	0xbb, 0xfb, 0x88, 0xdb, 0x2c, 0x80, 0x99, 0x07, 0x4f, 0xe8, 0xe7, 0x77, 0x65, 0xea, 0x2b, 0xd3,
	0x09, 0x9d, 0xb7, 0x96, 0xe7, 0x39, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x01, 0x6c, 0x12,
	0xfb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SpikeClient is the client API for Spike service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpikeClient interface {
	GetUserByAccessToken(ctx context.Context, in *GetUserByAccessTokenInput, opts ...grpc.CallOption) (*GetUserByAccessTokenOutput, error)
}

type spikeClient struct {
	cc *grpc.ClientConn
}

func NewSpikeClient(cc *grpc.ClientConn) SpikeClient {
	return &spikeClient{cc}
}

func (c *spikeClient) GetUserByAccessToken(ctx context.Context, in *GetUserByAccessTokenInput, opts ...grpc.CallOption) (*GetUserByAccessTokenOutput, error) {
	out := new(GetUserByAccessTokenOutput)
	err := c.cc.Invoke(ctx, "/sparkle.svcs.v1.Spike/GetUserByAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpikeServer is the server API for Spike service.
type SpikeServer interface {
	GetUserByAccessToken(context.Context, *GetUserByAccessTokenInput) (*GetUserByAccessTokenOutput, error)
}

func RegisterSpikeServer(s *grpc.Server, srv SpikeServer) {
	s.RegisterService(&_Spike_serviceDesc, srv)
}

func _Spike_GetUserByAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByAccessTokenInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpikeServer).GetUserByAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sparkle.svcs.v1.Spike/GetUserByAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpikeServer).GetUserByAccessToken(ctx, req.(*GetUserByAccessTokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _Spike_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sparkle.svcs.v1.Spike",
	HandlerType: (*SpikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByAccessToken",
			Handler:    _Spike_GetUserByAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svcs/v1/private.proto",
}
