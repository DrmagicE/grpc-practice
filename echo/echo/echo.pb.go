// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package echo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Message struct {
	MessageId            int32    `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
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

func (m *Message) GetMessageId() int32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *Message) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "echo.Message")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x1c, 0xb9, 0xd8, 0x7d, 0x53,
	0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0x64, 0xb8, 0x38, 0x73, 0x21, 0x4c, 0xcf, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0x84, 0x80, 0x90, 0x04, 0x17, 0x7b, 0x41, 0x62, 0x65, 0x4e, 0x7e,
	0x62, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x6b, 0xf4, 0x8a, 0x91, 0x8b, 0xdb,
	0x35, 0x39, 0x23, 0x3f, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x9b, 0x8b, 0x33, 0x34,
	0x2f, 0xb1, 0xa8, 0x12, 0x24, 0x26, 0xc4, 0xab, 0x07, 0xb6, 0x12, 0x6a, 0x87, 0x14, 0x2a, 0x57,
	0x89, 0x41, 0xc8, 0x9c, 0x4b, 0xd8, 0x39, 0x27, 0x33, 0x35, 0xaf, 0x24, 0xb8, 0xa4, 0x28, 0x35,
	0x31, 0x37, 0x33, 0x2f, 0x9d, 0x18, 0x6d, 0x1a, 0x8c, 0x20, 0x8d, 0x20, 0x0b, 0x53, 0x8b, 0x48,
	0xd2, 0x68, 0xc0, 0x28, 0x64, 0xc3, 0x25, 0xe6, 0x94, 0x99, 0x92, 0x59, 0x94, 0x9a, 0x5c, 0x92,
	0x99, 0x9f, 0x97, 0x98, 0x03, 0xd7, 0x4f, 0xd8, 0x52, 0x03, 0xc6, 0x24, 0x36, 0x70, 0xe0, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x23, 0xaa, 0x08, 0xc0, 0x4a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	UnaryEcho(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	ClientStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (EchoService_ClientStreamingEchoClient, error)
	ServerStreamingEcho(ctx context.Context, in *Message, opts ...grpc.CallOption) (EchoService_ServerStreamingEchoClient, error)
	BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (EchoService_BidirectionalStreamingClient, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) UnaryEcho(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/echo.EchoService/UnaryEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) ClientStreamingEcho(ctx context.Context, opts ...grpc.CallOption) (EchoService_ClientStreamingEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoService_serviceDesc.Streams[0], "/echo.EchoService/ClientStreamingEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceClientStreamingEchoClient{stream}
	return x, nil
}

type EchoService_ClientStreamingEchoClient interface {
	Send(*Message) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type echoServiceClientStreamingEchoClient struct {
	grpc.ClientStream
}

func (x *echoServiceClientStreamingEchoClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceClientStreamingEchoClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) ServerStreamingEcho(ctx context.Context, in *Message, opts ...grpc.CallOption) (EchoService_ServerStreamingEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoService_serviceDesc.Streams[1], "/echo.EchoService/ServerStreamingEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceServerStreamingEchoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EchoService_ServerStreamingEchoClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type echoServiceServerStreamingEchoClient struct {
	grpc.ClientStream
}

func (x *echoServiceServerStreamingEchoClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoServiceClient) BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (EchoService_BidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoService_serviceDesc.Streams[2], "/echo.EchoService/BidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceBidirectionalStreamingClient{stream}
	return x, nil
}

type EchoService_BidirectionalStreamingClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type echoServiceBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *echoServiceBidirectionalStreamingClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceBidirectionalStreamingClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	UnaryEcho(context.Context, *Message) (*Message, error)
	ClientStreamingEcho(EchoService_ClientStreamingEchoServer) error
	ServerStreamingEcho(*Message, EchoService_ServerStreamingEchoServer) error
	BidirectionalStreaming(EchoService_BidirectionalStreamingServer) error
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_UnaryEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).UnaryEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.EchoService/UnaryEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).UnaryEcho(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_ClientStreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).ClientStreamingEcho(&echoServiceClientStreamingEchoServer{stream})
}

type EchoService_ClientStreamingEchoServer interface {
	SendAndClose(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type echoServiceClientStreamingEchoServer struct {
	grpc.ServerStream
}

func (x *echoServiceClientStreamingEchoServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceClientStreamingEchoServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EchoService_ServerStreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServiceServer).ServerStreamingEcho(m, &echoServiceServerStreamingEchoServer{stream})
}

type EchoService_ServerStreamingEchoServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type echoServiceServerStreamingEchoServer struct {
	grpc.ServerStream
}

func (x *echoServiceServerStreamingEchoServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _EchoService_BidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).BidirectionalStreaming(&echoServiceBidirectionalStreamingServer{stream})
}

type EchoService_BidirectionalStreamingServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type echoServiceBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *echoServiceBidirectionalStreamingServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceBidirectionalStreamingServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryEcho",
			Handler:    _EchoService_UnaryEcho_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStreamingEcho",
			Handler:       _EchoService_ClientStreamingEcho_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStreamingEcho",
			Handler:       _EchoService_ServerStreamingEcho_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidirectionalStreaming",
			Handler:       _EchoService_BidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "echo.proto",
}
