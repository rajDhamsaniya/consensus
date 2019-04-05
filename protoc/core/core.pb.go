// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package helloworld

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

type Executetx struct {
	Tx                   string   `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	Args                 []byte   `protobuf:"bytes,2,opt,name=Args,proto3" json:"Args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Executetx) Reset()         { *m = Executetx{} }
func (m *Executetx) String() string { return proto.CompactTextString(m) }
func (*Executetx) ProtoMessage()    {}
func (*Executetx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0}
}

func (m *Executetx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Executetx.Unmarshal(m, b)
}
func (m *Executetx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Executetx.Marshal(b, m, deterministic)
}
func (m *Executetx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Executetx.Merge(m, src)
}
func (m *Executetx) XXX_Size() int {
	return xxx_messageInfo_Executetx.Size(m)
}
func (m *Executetx) XXX_DiscardUnknown() {
	xxx_messageInfo_Executetx.DiscardUnknown(m)
}

var xxx_messageInfo_Executetx proto.InternalMessageInfo

func (m *Executetx) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

func (m *Executetx) GetArgs() []byte {
	if m != nil {
		return m.Args
	}
	return nil
}

type ExecResponse struct {
	Sign                 string   `protobuf:"bytes,1,opt,name=Sign,proto3" json:"Sign,omitempty"`
	Result               []byte   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecResponse) Reset()         { *m = ExecResponse{} }
func (m *ExecResponse) String() string { return proto.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()    {}
func (*ExecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1}
}

func (m *ExecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecResponse.Unmarshal(m, b)
}
func (m *ExecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecResponse.Marshal(b, m, deterministic)
}
func (m *ExecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecResponse.Merge(m, src)
}
func (m *ExecResponse) XXX_Size() int {
	return xxx_messageInfo_ExecResponse.Size(m)
}
func (m *ExecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecResponse proto.InternalMessageInfo

func (m *ExecResponse) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *ExecResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{2}
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

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{3}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Executetx)(nil), "helloworld.Executetx")
	proto.RegisterType((*ExecResponse)(nil), "helloworld.ExecResponse")
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_f7e43720d1edc0fe) }

var fileDescriptor_f7e43720d1edc0fe = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xfc, 0x30,
	0x10, 0xc5, 0xb7, 0xe5, 0xcf, 0xee, 0x7f, 0x87, 0xaa, 0x30, 0xe0, 0x52, 0xd6, 0xcb, 0x92, 0x83,
	0xec, 0xa9, 0x8a, 0xde, 0x3c, 0x08, 0xbb, 0x22, 0xeb, 0x71, 0xe9, 0x0a, 0x9e, 0x63, 0x1d, 0x6a,
	0x21, 0x4d, 0x62, 0x92, 0x62, 0xfb, 0x49, 0xfd, 0x3a, 0x92, 0xd2, 0x68, 0x11, 0x4f, 0xde, 0xde,
	0x24, 0xef, 0x97, 0x47, 0xde, 0x00, 0x14, 0xca, 0x50, 0xa6, 0x8d, 0x72, 0x0a, 0xe1, 0x95, 0x84,
	0x50, 0xef, 0xca, 0x88, 0x17, 0x76, 0x01, 0xf3, 0xfb, 0x96, 0x8a, 0xc6, 0x91, 0x6b, 0xf1, 0x18,
	0x62, 0xd7, 0xa6, 0xd1, 0x2a, 0x5a, 0xcf, 0xf3, 0xd8, 0xb5, 0x88, 0xf0, 0x6f, 0x63, 0x4a, 0x9b,
	0xc6, 0xab, 0x68, 0x9d, 0xe4, 0xbd, 0x66, 0x37, 0x90, 0x78, 0x20, 0x27, 0xab, 0x95, 0xb4, 0xe4,
	0x3d, 0x87, 0xaa, 0x94, 0x03, 0xd5, 0x6b, 0x5c, 0xc0, 0xd4, 0x90, 0x6d, 0x84, 0x1b, 0xc8, 0x61,
	0x62, 0x0c, 0x92, 0x07, 0x1f, 0x9d, 0xd3, 0x5b, 0x43, 0xd6, 0x79, 0x56, 0xf2, 0x9a, 0x02, 0xeb,
	0x35, 0x3b, 0x07, 0x18, 0x3c, 0x5a, 0x74, 0x98, 0xc2, 0xac, 0x26, 0x6b, 0x79, 0x19, 0x4c, 0x61,
	0xbc, 0xfa, 0x88, 0x60, 0xb6, 0x33, 0x44, 0x8e, 0x0c, 0xde, 0xc2, 0xff, 0x03, 0xef, 0x7a, 0x0c,
	0xd3, 0xec, 0xfb, 0x77, 0xd9, 0x38, 0x6d, 0xb9, 0xf8, 0xe5, 0x46, 0x8b, 0x8e, 0x4d, 0xf0, 0x0e,
	0x8e, 0x02, 0xbf, 0x29, 0x79, 0x25, 0xff, 0xf4, 0xc8, 0x0e, 0x70, 0x68, 0xf2, 0xd1, 0x70, 0x69,
	0x79, 0xe1, 0x2a, 0x25, 0xf1, 0x74, 0xec, 0xff, 0x6a, 0x7a, 0x99, 0xfe, 0x3c, 0x0e, 0x7d, 0xb2,
	0xc9, 0xf6, 0x12, 0xce, 0x2a, 0x95, 0x95, 0x46, 0x17, 0x19, 0xb5, 0xbc, 0xd6, 0x82, 0xec, 0xc8,
	0xbd, 0x3d, 0xe9, 0x53, 0x9f, 0xbc, 0xde, 0xfb, 0x75, 0xee, 0xa3, 0xe7, 0x69, 0xbf, 0xd7, 0xeb,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x05, 0x69, 0x40, 0xe5, 0x01, 0x00, 0x00,
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
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// Sends another greeting
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	ExecuteTransaction(ctx context.Context, in *Executetx, opts ...grpc.CallOption) (*ExecResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHelloAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ExecuteTransaction(ctx context.Context, in *Executetx, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/ExecuteTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// Sends another greeting
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
	ExecuteTransaction(context.Context, *Executetx) (*ExecResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ExecuteTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Executetx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ExecuteTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/ExecuteTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ExecuteTransaction(ctx, req.(*Executetx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Greeter_SayHelloAgain_Handler,
		},
		{
			MethodName: "ExecuteTransaction",
			Handler:    _Greeter_ExecuteTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}