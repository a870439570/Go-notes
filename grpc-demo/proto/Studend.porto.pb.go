// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/Studend.porto

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

type MyRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyRequest) Reset()         { *m = MyRequest{} }
func (m *MyRequest) String() string { return proto.CompactTextString(m) }
func (*MyRequest) ProtoMessage()    {}
func (*MyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a18f3d1a0276fb7, []int{0}
}

func (m *MyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyRequest.Unmarshal(m, b)
}
func (m *MyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyRequest.Marshal(b, m, deterministic)
}
func (m *MyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyRequest.Merge(m, src)
}
func (m *MyRequest) XXX_Size() int {
	return xxx_messageInfo_MyRequest.Size(m)
}
func (m *MyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MyRequest proto.InternalMessageInfo

func (m *MyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MyResponse struct {
	Realname             string   `protobuf:"bytes,2,opt,name=realname,proto3" json:"realname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyResponse) Reset()         { *m = MyResponse{} }
func (m *MyResponse) String() string { return proto.CompactTextString(m) }
func (*MyResponse) ProtoMessage()    {}
func (*MyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a18f3d1a0276fb7, []int{1}
}

func (m *MyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyResponse.Unmarshal(m, b)
}
func (m *MyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyResponse.Marshal(b, m, deterministic)
}
func (m *MyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyResponse.Merge(m, src)
}
func (m *MyResponse) XXX_Size() int {
	return xxx_messageInfo_MyResponse.Size(m)
}
func (m *MyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MyResponse proto.InternalMessageInfo

func (m *MyResponse) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func init() {
	proto.RegisterType((*MyRequest)(nil), "proto.MyRequest")
	proto.RegisterType((*MyResponse)(nil), "proto.MyResponse")
}

func init() { proto.RegisterFile("proto/Studend.porto", fileDescriptor_1a18f3d1a0276fb7) }

var fileDescriptor_1a18f3d1a0276fb7 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x0f, 0x2e, 0x29, 0x4d, 0x49, 0xcd, 0x4b, 0xd1, 0x2b, 0xc8, 0x2f, 0x2a, 0xc9, 0x17,
	0x62, 0x05, 0x0b, 0x2a, 0xc9, 0x73, 0x71, 0xfa, 0x56, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x4a, 0x1a, 0x5c, 0x5c, 0x20, 0x05, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x52, 0x5c,
	0x1c, 0x45, 0xa9, 0x89, 0x39, 0x60, 0x55, 0x4c, 0x60, 0x55, 0x70, 0xbe, 0x91, 0x3b, 0x97, 0x30,
	0xc4, 0x8a, 0x92, 0x00, 0x90, 0xd1, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x06, 0x5c,
	0xec, 0xee, 0xa9, 0x25, 0x7e, 0x89, 0xb9, 0xa9, 0x42, 0x02, 0x7a, 0x60, 0x4b, 0xf5, 0xe0, 0x36,
	0x4a, 0x09, 0x22, 0x89, 0x40, 0xac, 0x50, 0x62, 0x48, 0x62, 0x03, 0x8b, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x61, 0x9f, 0x2c, 0x26, 0xb8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StudentProtoServiceClient is the client API for StudentProtoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentProtoServiceClient interface {
	GetName(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error)
}

type studentProtoServiceClient struct {
	cc *grpc.ClientConn
}

func NewStudentProtoServiceClient(cc *grpc.ClientConn) StudentProtoServiceClient {
	return &studentProtoServiceClient{cc}
}

func (c *studentProtoServiceClient) GetName(ctx context.Context, in *MyRequest, opts ...grpc.CallOption) (*MyResponse, error) {
	out := new(MyResponse)
	err := c.cc.Invoke(ctx, "/proto.StudentProtoService/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentProtoServiceServer is the server API for StudentProtoService service.
type StudentProtoServiceServer interface {
	GetName(context.Context, *MyRequest) (*MyResponse, error)
}

// UnimplementedStudentProtoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStudentProtoServiceServer struct {
}

func (*UnimplementedStudentProtoServiceServer) GetName(ctx context.Context, req *MyRequest) (*MyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}

func RegisterStudentProtoServiceServer(s *grpc.Server, srv StudentProtoServiceServer) {
	s.RegisterService(&_StudentProtoService_serviceDesc, srv)
}

func _StudentProtoService_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentProtoServiceServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StudentProtoService/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentProtoServiceServer).GetName(ctx, req.(*MyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudentProtoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StudentProtoService",
	HandlerType: (*StudentProtoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetName",
			Handler:    _StudentProtoService_GetName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Studend.porto",
}
