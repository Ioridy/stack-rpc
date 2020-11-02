// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-stack/examples/greeter/greeter.proto

package greeter

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcac6c3c12bb4962, []int{0}
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

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcac6c3c12bb4962, []int{1}
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

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() {
	proto.RegisterFile("cmd/protoc-gen-stack/examples/greeter/greeter.proto", fileDescriptor_dcac6c3c12bb4962)
}

var fileDescriptor_dcac6c3c12bb4962 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0x3d, 0x4e, 0xc4, 0x30,
	0x10, 0x85, 0x65, 0x7e, 0x92, 0xcd, 0x50, 0x00, 0xae, 0xd0, 0x6a, 0x91, 0x90, 0xab, 0x15, 0xd2,
	0x66, 0xd0, 0x6e, 0xc7, 0x05, 0xa0, 0x0e, 0x27, 0x30, 0x61, 0x64, 0x22, 0xd9, 0x1e, 0x63, 0x1b,
	0x89, 0x9a, 0x2b, 0x70, 0x34, 0xae, 0xc0, 0x41, 0x50, 0x9c, 0xa4, 0xdb, 0x6a, 0x9e, 0xf4, 0xbe,
	0x6f, 0x66, 0xe0, 0xd0, 0xbb, 0x37, 0x0c, 0x91, 0x33, 0xf7, 0x3b, 0x43, 0x7e, 0xe7, 0x86, 0x3e,
	0x32, 0xd2, 0x97, 0x76, 0xc1, 0x52, 0x42, 0x13, 0x89, 0x32, 0xc5, 0x65, 0xb6, 0x85, 0x5c, 0x6f,
	0x0c, 0xb3, 0xb1, 0x84, 0x3a, 0x0c, 0xa8, 0xbd, 0xe7, 0xac, 0xf3, 0xc0, 0x3e, 0x4d, 0xad, 0xba,
	0x85, 0xba, 0xa3, 0x8f, 0x4f, 0x4a, 0x59, 0x4a, 0x38, 0xf3, 0xda, 0xd1, 0x8d, 0xb8, 0x13, 0xdb,
	0xa6, 0x2b, 0x59, 0x6d, 0x60, 0xd5, 0x51, 0x0a, 0xec, 0x13, 0xc9, 0x2b, 0x38, 0x75, 0xc9, 0xcc,
	0xf5, 0x18, 0xf7, 0x1e, 0xea, 0xa7, 0xe9, 0x96, 0x44, 0x38, 0x7f, 0x26, 0x6b, 0x59, 0xae, 0xda,
	0x79, 0xdf, 0xba, 0x69, 0x17, 0x55, 0x5d, 0x7f, 0xff, 0xfe, 0xfd, 0x9c, 0x5c, 0xa8, 0x0a, 0xdf,
	0x47, 0xe8, 0x51, 0xdc, 0xcb, 0x3d, 0x54, 0x2f, 0x39, 0x92, 0x76, 0xc7, 0x8d, 0xcb, 0x62, 0x34,
	0xb2, 0xc6, 0x54, 0xa8, 0xad, 0x78, 0x10, 0xaf, 0x55, 0xf9, 0xf9, 0xf0, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0x4c, 0x4a, 0xe5, 0x08, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Greeter_StreamClient, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Greeter/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Greeter_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greeter_serviceDesc.Streams[0], "/Greeter/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterStreamClient{stream}
	return x, nil
}

type Greeter_StreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type greeterStreamClient struct {
	grpc.ClientStream
}

func (x *greeterStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	Hello(context.Context, *Request) (*Response, error)
	Stream(Greeter_StreamServer) error
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) Hello(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedGreeterServer) Stream(srv Greeter_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).Stream(&greeterStreamServer{stream})
}

type Greeter_StreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type greeterStreamServer struct {
	grpc.ServerStream
}

func (x *greeterStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Greeter_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Greeter_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cmd/protoc-gen-stack/examples/greeter/greeter.proto",
}
