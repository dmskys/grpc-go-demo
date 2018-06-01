// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package pbproto is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserRequest
	UserResponse
*/
package user

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

// message 对应生成代码的 struct
// 定义客户端请求的数据格式
type UserRequest struct {
	// [修饰符] 类型 字段名 = 标识符;
	Uid int32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *UserRequest) Reset()                    { *m = UserRequest{} }
func (m *UserRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()               {}
func (*UserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

// 定义服务端响应的数据格式
type UserResponse struct {
	Id    int32    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age   int32    `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	Title []string `protobuf:"bytes,4,rep,name=title" json:"title,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserResponse) GetTitle() []string {
	if m != nil {
		return m.Title
	}
	return nil
}

func init() {
	proto.RegisterType((*UserRequest)(nil), "pbproto.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "pbproto.UserResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserInfoService service

type UserInfoServiceClient interface {
	// rpc 定义服务内的 GetUserInfo 远程调用
	// 微服务中获取用户信息的 RPC 函数
	GetUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userInfoServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserInfoServiceClient(cc *grpc.ClientConn) UserInfoServiceClient {
	return &userInfoServiceClient{cc}
}

func (c *userInfoServiceClient) GetUserInfo(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/pbproto.UserInfoService/GetUserInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserInfoService service

type UserInfoServiceServer interface {
	// rpc 定义服务内的 GetUserInfo 远程调用
	// 微服务中获取用户信息的 RPC 函数
	GetUserInfo(context.Context, *UserRequest) (*UserResponse, error)
}

func RegisterUserInfoServiceServer(s *grpc.Server, srv UserInfoServiceServer) {
	s.RegisterService(&_UserInfoService_serviceDesc, srv)
}

func _UserInfoService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbproto.UserInfoService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetUserInfo(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbproto.UserInfoService",
	HandlerType: (*UserInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _UserInfoService_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0x3f, 0x0b, 0xc2, 0x30,
	0x10, 0x47, 0xe9, 0x3f, 0xa5, 0x57, 0x51, 0x39, 0x2a, 0x04, 0x17, 0x4b, 0xa7, 0x4e, 0x1d, 0x74,
	0xf3, 0x0b, 0x88, 0x83, 0x4b, 0xc4, 0xc5, 0xad, 0xb5, 0x67, 0x09, 0x68, 0x53, 0x93, 0xd4, 0xcf,
	0x2f, 0x6d, 0x54, 0xc4, 0xed, 0x77, 0x8f, 0xe3, 0xf1, 0x00, 0x3a, 0x4d, 0x2a, 0x6f, 0x95, 0x34,
	0x12, 0xc7, 0x6d, 0x39, 0x8c, 0x74, 0x05, 0xd1, 0x49, 0x93, 0xe2, 0xf4, 0xe8, 0x48, 0x1b, 0x9c,
	0x83, 0xd7, 0x89, 0x8a, 0x39, 0x89, 0x93, 0x05, 0xbc, 0x9f, 0xe9, 0x19, 0x26, 0xf6, 0x41, 0xb7,
	0xb2, 0xd1, 0x84, 0x53, 0x70, 0xbf, 0x0f, 0xae, 0xa8, 0x10, 0xc1, 0x6f, 0x8a, 0x3b, 0x31, 0x37,
	0x71, 0xb2, 0x90, 0x0f, 0xbb, 0xb7, 0x14, 0x35, 0x31, 0xcf, 0x5a, 0x8a, 0x9a, 0x30, 0x86, 0xc0,
	0x08, 0x73, 0x23, 0xe6, 0x27, 0x5e, 0x16, 0x72, 0x7b, 0xac, 0x0f, 0x30, 0xeb, 0xdd, 0xfb, 0xe6,
	0x2a, 0x8f, 0xa4, 0x9e, 0xe2, 0x42, 0xb8, 0x85, 0x68, 0x47, 0xe6, 0x43, 0x31, 0xce, 0xdf, 0xa1,
	0xf9, 0x4f, 0xe5, 0x72, 0xf1, 0x47, 0x6d, 0x5a, 0x39, 0x1a, 0xd8, 0xe6, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x39, 0xda, 0x18, 0x05, 0xe9, 0x00, 0x00, 0x00,
}