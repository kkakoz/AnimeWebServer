// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/user/user.proto

package userpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type LoginReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bc452c1a8bd0aa8, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bc452c1a8bd0aa8, []int{1}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LoginRes struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	CreateAt             string   `protobuf:"bytes,5,opt,name=createAt,proto3" json:"createAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}
func (*LoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bc452c1a8bd0aa8, []int{2}
}

func (m *LoginRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRes.Unmarshal(m, b)
}
func (m *LoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRes.Marshal(b, m, deterministic)
}
func (m *LoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRes.Merge(m, src)
}
func (m *LoginRes) XXX_Size() int {
	return xxx_messageInfo_LoginRes.Size(m)
}
func (m *LoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRes proto.InternalMessageInfo

func (m *LoginRes) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LoginRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginRes) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginRes) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

type UserInfoRes struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoRes) Reset()         { *m = UserInfoRes{} }
func (m *UserInfoRes) String() string { return proto.CompactTextString(m) }
func (*UserInfoRes) ProtoMessage()    {}
func (*UserInfoRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bc452c1a8bd0aa8, []int{3}
}

func (m *UserInfoRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRes.Unmarshal(m, b)
}
func (m *UserInfoRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRes.Marshal(b, m, deterministic)
}
func (m *UserInfoRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRes.Merge(m, src)
}
func (m *UserInfoRes) XXX_Size() int {
	return xxx_messageInfo_UserInfoRes.Size(m)
}
func (m *UserInfoRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRes.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRes proto.InternalMessageInfo

func (m *UserInfoRes) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfoRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfoRes) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserInfoList struct {
	UserinfoList         []*UserInfoRes `protobuf:"bytes,1,rep,name=userinfoList,proto3" json:"userinfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UserInfoList) Reset()         { *m = UserInfoList{} }
func (m *UserInfoList) String() string { return proto.CompactTextString(m) }
func (*UserInfoList) ProtoMessage()    {}
func (*UserInfoList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bc452c1a8bd0aa8, []int{4}
}

func (m *UserInfoList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoList.Unmarshal(m, b)
}
func (m *UserInfoList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoList.Marshal(b, m, deterministic)
}
func (m *UserInfoList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoList.Merge(m, src)
}
func (m *UserInfoList) XXX_Size() int {
	return xxx_messageInfo_UserInfoList.Size(m)
}
func (m *UserInfoList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoList.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoList proto.InternalMessageInfo

func (m *UserInfoList) GetUserinfoList() []*UserInfoRes {
	if m != nil {
		return m.UserinfoList
	}
	return nil
}

type Id struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_3bc452c1a8bd0aa8, []int{5}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "LoginReq")
	proto.RegisterType((*RegisterReq)(nil), "RegisterReq")
	proto.RegisterType((*LoginRes)(nil), "LoginRes")
	proto.RegisterType((*UserInfoRes)(nil), "UserInfoRes")
	proto.RegisterType((*UserInfoList)(nil), "UserInfoList")
	proto.RegisterType((*Id)(nil), "Id")
}

func init() { proto.RegisterFile("api/user/user.proto", fileDescriptor_3bc452c1a8bd0aa8) }

var fileDescriptor_3bc452c1a8bd0aa8 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xed, 0x24, 0x38, 0x93, 0x08, 0xa1, 0x6d, 0xa0, 0x8b, 0xdb, 0x8a, 0x6a, 0x25, 0xa4,
	0x8a, 0xc3, 0x1a, 0x95, 0x5b, 0x91, 0x10, 0x44, 0x42, 0x28, 0x52, 0x0f, 0xc8, 0xc0, 0x85, 0x13,
	0x9b, 0x7a, 0x1a, 0xad, 0x88, 0x77, 0xcd, 0xee, 0x36, 0x88, 0x2b, 0xbf, 0xc0, 0xa7, 0xf1, 0x01,
	0x5c, 0xf8, 0x8a, 0x9e, 0x90, 0xd7, 0x5e, 0x2b, 0xc9, 0x95, 0x5e, 0xac, 0x7d, 0xf3, 0xde, 0xbc,
	0x99, 0xf1, 0x0c, 0x1c, 0x88, 0x5a, 0xe6, 0x37, 0x16, 0x8d, 0xff, 0xf0, 0xda, 0x68, 0xa7, 0xb3,
	0xa3, 0x95, 0xd6, 0xab, 0x35, 0xe6, 0x1e, 0x2d, 0x6f, 0xae, 0x73, 0xac, 0x6a, 0xf7, 0xa3, 0x23,
	0x8f, 0x3b, 0xb2, 0x49, 0x14, 0x4a, 0x69, 0x27, 0x9c, 0xd4, 0xca, 0x76, 0xec, 0xe1, 0x46, 0xac,
	0x65, 0x29, 0x1c, 0xe6, 0xe1, 0xd1, 0x12, 0xec, 0x3d, 0xa4, 0x97, 0x7a, 0x25, 0x55, 0x81, 0xdf,
	0xc8, 0x09, 0x0c, 0xb1, 0x12, 0x72, 0x4d, 0xa3, 0xd3, 0xe8, 0x6c, 0x3c, 0xbf, 0x77, 0x3b, 0x1f,
	0x98, 0xf8, 0x4b, 0x54, 0xb4, 0x51, 0xf2, 0x14, 0xd2, 0x5a, 0x58, 0xfb, 0x5d, 0x9b, 0x92, 0xc6,
	0x5e, 0x31, 0xbe, 0x9d, 0x8f, 0xcc, 0xe0, 0xc1, 0x88, 0x42, 0xd1, 0x53, 0xcc, 0xc0, 0xa4, 0xc0,
	0x95, 0xb4, 0x0e, 0xcd, 0x9d, 0x99, 0x92, 0x13, 0x18, 0x28, 0x51, 0x21, 0x4d, 0xb6, 0x25, 0x09,
	0x85, 0xc2, 0x87, 0xd9, 0xa6, 0x9f, 0xc2, 0x92, 0xfb, 0x10, 0xcb, 0xd2, 0x57, 0x4b, 0x8a, 0x58,
	0x96, 0x84, 0x74, 0xa9, 0xde, 0xbd, 0xd5, 0x93, 0x59, 0x68, 0xca, 0xfb, 0x85, 0x5e, 0x66, 0x30,
	0x74, 0xfa, 0x2b, 0x2a, 0x3a, 0x68, 0xa3, 0x1e, 0x90, 0x0c, 0xd2, 0x2b, 0x83, 0xc2, 0xe1, 0x1b,
	0x47, 0x87, 0x9e, 0xe8, 0x31, 0x7b, 0x07, 0x93, 0x4f, 0x16, 0xcd, 0x42, 0x5d, 0xeb, 0xff, 0x2a,
	0xcd, 0x5e, 0xc3, 0x34, 0x18, 0x5d, 0x4a, 0xeb, 0xc8, 0x73, 0x98, 0x36, 0x8b, 0x97, 0x1d, 0xa6,
	0xd1, 0x69, 0x72, 0x36, 0x39, 0x9f, 0xf2, 0xad, 0x6a, 0xc5, 0x8e, 0x82, 0xcd, 0x20, 0x5e, 0x94,
	0xfb, 0x1d, 0x9c, 0xff, 0x89, 0xda, 0x0e, 0x3f, 0xa0, 0xd9, 0xc8, 0x2b, 0x24, 0xaf, 0x60, 0xe8,
	0x7f, 0x14, 0x19, 0xf3, 0xb0, 0xf6, 0xac, 0x7f, 0x5a, 0xf6, 0xe4, 0xe7, 0xef, 0xbf, 0xbf, 0xe2,
	0xc7, 0x6c, 0x96, 0x0b, 0x25, 0x2b, 0x14, 0x75, 0xdd, 0x1e, 0xe1, 0xba, 0x11, 0x5c, 0x44, 0xcf,
	0xc8, 0x47, 0x48, 0xc3, 0x72, 0xc9, 0x94, 0x6f, 0xed, 0x39, 0x7b, 0xc4, 0xdb, 0x03, 0xe4, 0xe1,
	0x3a, 0xf9, 0xdb, 0xe6, 0x3a, 0x19, 0xf3, 0x96, 0xc7, 0xec, 0x70, 0xcf, 0xd2, 0x74, 0xb9, 0x8d,
	0xeb, 0x05, 0xa4, 0x61, 0x30, 0x92, 0xf0, 0x45, 0x99, 0xed, 0x0c, 0xca, 0x8e, 0xbc, 0xc5, 0x43,
	0x72, 0xb0, 0x67, 0xd1, 0xcc, 0x3e, 0x9f, 0x7c, 0x1e, 0xf3, 0xfc, 0x65, 0x83, 0xeb, 0xe5, 0x72,
	0xe4, 0x8b, 0xbf, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x80, 0x55, 0xfc, 0xf4, 0x3f, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserInfoRes, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*UserInfoRes, error) {
	out := new(UserInfoRes)
	err := c.cc.Invoke(ctx, "/UserService/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
	Register(context.Context, *RegisterReq) (*emptypb.Empty, error)
	UserInfo(context.Context, *Id) (*UserInfoRes, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserInfo(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _UserService_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/user.proto",
}
