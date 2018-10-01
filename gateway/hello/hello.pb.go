// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type HelloResponse struct {
	Data                 string   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "hello.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "hello.HelloResponse")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce) }

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x64, 0xd2, 0xf3, 0xf3,
	0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0x8a, 0x21, 0x8a, 0x94, 0x94, 0xb8, 0x78, 0x3c, 0x40, 0xca, 0x82, 0x52, 0x0b, 0x4b,
	0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x5c, 0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xc0, 0x6c, 0x25, 0x65, 0x2e, 0x5e, 0xa8, 0x9a, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x6c, 0x8a, 0x8c, 0xe6, 0xb2, 0x70, 0xb1, 0x82, 0x55, 0x09, 0x05, 0xc0, 0x18, 0xc2, 0x7a, 0x10,
	0xe7, 0x20, 0x5b, 0x20, 0x25, 0x82, 0x2a, 0x08, 0x31, 0x51, 0x49, 0xa6, 0xe9, 0xf2, 0x93, 0xc9,
	0x4c, 0x62, 0x4a, 0x82, 0xfa, 0x65, 0x86, 0xfa, 0xe9, 0x89, 0x25, 0xa9, 0xe5, 0x89, 0x95, 0xfa,
	0x60, 0x85, 0x56, 0x8c, 0x5a, 0x42, 0xf1, 0x50, 0x47, 0x86, 0x64, 0xe6, 0xa6, 0xe6, 0x97, 0x96,
	0x90, 0x62, 0xb0, 0x0a, 0xd8, 0x60, 0x39, 0x25, 0x49, 0x0c, 0x83, 0xe3, 0x4b, 0x20, 0xa6, 0x81,
	0x2c, 0xc8, 0xe5, 0x12, 0x42, 0x36, 0x2b, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0x97, 0x14, 0x6b, 0xb4,
	0xc0, 0xd6, 0xa8, 0x28, 0xc9, 0x23, 0x5b, 0x83, 0xac, 0x2f, 0xbe, 0x18, 0x6c, 0xa6, 0x15, 0xa3,
	0x96, 0x06, 0xa3, 0x50, 0x3e, 0x97, 0x30, 0x8a, 0x76, 0xd2, 0xed, 0xd3, 0x06, 0xdb, 0xa7, 0xaa,
	0xa4, 0x80, 0xc5, 0x3e, 0x88, 0x12, 0x84, 0x85, 0x06, 0x8c, 0x42, 0x0d, 0x8c, 0x5c, 0x52, 0xc8,
	0xa6, 0x92, 0x6f, 0xb1, 0x31, 0xd8, 0x62, 0x5d, 0x25, 0x0d, 0x9c, 0x1e, 0x2d, 0xc2, 0x70, 0x80,
	0x06, 0xa3, 0x01, 0xa3, 0x13, 0x7b, 0x14, 0x24, 0x3d, 0x26, 0xb1, 0x81, 0x13, 0x9e, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x06, 0xd5, 0xbd, 0x11, 0xac, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloTimeout(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloRequestStream(ctx context.Context, opts ...grpc.CallOption) (Hello_HelloRequestStreamClient, error)
	HelloResponseStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Hello_HelloResponseStreamClient, error)
	HelloRequestResponseStream(ctx context.Context, opts ...grpc.CallOption) (Hello_HelloRequestResponseStreamClient, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello.Hello/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) HelloTimeout(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello.Hello/HelloTimeout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) HelloRequestStream(ctx context.Context, opts ...grpc.CallOption) (Hello_HelloRequestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[0], "/hello.Hello/HelloRequestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloHelloRequestStreamClient{stream}
	return x, nil
}

type Hello_HelloRequestStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloHelloRequestStreamClient struct {
	grpc.ClientStream
}

func (x *helloHelloRequestStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloHelloRequestStreamClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) HelloResponseStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Hello_HelloResponseStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[1], "/hello.Hello/HelloResponseStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloHelloResponseStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hello_HelloResponseStreamClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloHelloResponseStreamClient struct {
	grpc.ClientStream
}

func (x *helloHelloResponseStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) HelloRequestResponseStream(ctx context.Context, opts ...grpc.CallOption) (Hello_HelloRequestResponseStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Hello_serviceDesc.Streams[2], "/hello.Hello/HelloRequestResponseStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloHelloRequestResponseStreamClient{stream}
	return x, nil
}

type Hello_HelloRequestResponseStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloHelloRequestResponseStreamClient struct {
	grpc.ClientStream
}

func (x *helloHelloRequestResponseStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloHelloRequestResponseStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloTimeout(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloRequestStream(Hello_HelloRequestStreamServer) error
	HelloResponseStream(*HelloRequest, Hello_HelloResponseStreamServer) error
	HelloRequestResponseStream(Hello_HelloRequestResponseStreamServer) error
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_HelloTimeout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).HelloTimeout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/HelloTimeout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).HelloTimeout(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_HelloRequestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).HelloRequestStream(&helloHelloRequestStreamServer{stream})
}

type Hello_HelloRequestStreamServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloHelloRequestStreamServer struct {
	grpc.ServerStream
}

func (x *helloHelloRequestStreamServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloHelloRequestStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Hello_HelloResponseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServer).HelloResponseStream(m, &helloHelloResponseStreamServer{stream})
}

type Hello_HelloResponseStreamServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type helloHelloResponseStreamServer struct {
	grpc.ServerStream
}

func (x *helloHelloResponseStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Hello_HelloRequestResponseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).HelloRequestResponseStream(&helloHelloRequestResponseStreamServer{stream})
}

type Hello_HelloRequestResponseStreamServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloHelloRequestResponseStreamServer struct {
	grpc.ServerStream
}

func (x *helloHelloRequestResponseStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloHelloRequestResponseStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Hello_Hello_Handler,
		},
		{
			MethodName: "HelloTimeout",
			Handler:    _Hello_HelloTimeout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloRequestStream",
			Handler:       _Hello_HelloRequestStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloResponseStream",
			Handler:       _Hello_HelloResponseStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HelloRequestResponseStream",
			Handler:       _Hello_HelloRequestResponseStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}