// Code generated by protoc-gen-go. DO NOT EDIT.
// source: promotion.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	PromotionCode        string   `protobuf:"bytes,1,opt,name=promotionCode,proto3" json:"promotionCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_93258b77978a8f53, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPromotionCode() string {
	if m != nil {
		return m.PromotionCode
	}
	return ""
}

type Response struct {
	Charge               uint64   `protobuf:"varint,1,opt,name=charge,proto3" json:"charge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_93258b77978a8f53, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCharge() uint64 {
	if m != nil {
		return m.Charge
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() { proto.RegisterFile("promotion.proto", fileDescriptor_93258b77978a8f53) }

var fileDescriptor_93258b77978a8f53 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0xcf,
	0xcd, 0x2f, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0xfa, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x2a, 0x5c, 0xbc, 0x70, 0x45,
	0xce, 0xf9, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xa8, 0x82, 0x4a, 0x4a, 0x5c,
	0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c, 0xc9, 0x19, 0x89,
	0x45, 0xe9, 0x10, 0xa5, 0x2c, 0x41, 0x50, 0x9e, 0x91, 0x19, 0x17, 0x67, 0x00, 0x4c, 0x93, 0x90,
	0x26, 0x17, 0x5b, 0x58, 0x6a, 0x51, 0x66, 0x5a, 0xa5, 0x10, 0x1f, 0xc4, 0x6a, 0x3d, 0xa8, 0x85,
	0x52, 0xfc, 0x70, 0x3e, 0xc4, 0xbc, 0x24, 0x36, 0x30, 0xdf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x81, 0xd9, 0xd2, 0x0d, 0xad, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PromotionClient is the client API for Promotion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PromotionClient interface {
	Verify(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type promotionClient struct {
	cc grpc.ClientConnInterface
}

func NewPromotionClient(cc grpc.ClientConnInterface) PromotionClient {
	return &promotionClient{cc}
}

func (c *promotionClient) Verify(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Promotion/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromotionServer is the server API for Promotion service.
type PromotionServer interface {
	Verify(context.Context, *Request) (*Response, error)
}

// UnimplementedPromotionServer can be embedded to have forward compatible implementations.
type UnimplementedPromotionServer struct {
}

func (*UnimplementedPromotionServer) Verify(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}

func RegisterPromotionServer(s *grpc.Server, srv PromotionServer) {
	s.RegisterService(&_Promotion_serviceDesc, srv)
}

func _Promotion_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Promotion/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServer).Verify(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Promotion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Promotion",
	HandlerType: (*PromotionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _Promotion_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promotion.proto",
}