// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/base/error.proto

package basepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ErrorCode int32

const (
	ErrorCode_EC_UNUSE          ErrorCode = 0
	ErrorCode_EC_SUCCESS        ErrorCode = 200
	ErrorCode_EC_BUSINESSERR    ErrorCode = 400
	ErrorCode_EC_UNAUTHORIZED   ErrorCode = 402
	ErrorCode_EC_SERVER_UNKNOWN ErrorCode = 500
)

var ErrorCode_name = map[int32]string{
	0:   "EC_UNUSE",
	200: "EC_SUCCESS",
	400: "EC_BUSINESSERR",
	402: "EC_UNAUTHORIZED",
	500: "EC_SERVER_UNKNOWN",
}

var ErrorCode_value = map[string]int32{
	"EC_UNUSE":          0,
	"EC_SUCCESS":        200,
	"EC_BUSINESSERR":    400,
	"EC_UNAUTHORIZED":   402,
	"EC_SERVER_UNKNOWN": 500,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_585e389d4e591de1, []int{0}
}

func init() {
	proto.RegisterEnum("ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("api/base/error.proto", fileDescriptor_585e389d4e591de1) }

var fileDescriptor_585e389d4e591de1 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2c, 0xc8, 0xd4,
	0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0xd7, 0xca, 0xe5, 0xe2, 0x74, 0x05, 0x71, 0x9d, 0xf3, 0x53, 0x52, 0x85, 0x78, 0xb8, 0x38, 0x5c,
	0x9d, 0xe3, 0x43, 0xfd, 0x42, 0x83, 0x5d, 0x05, 0x18, 0x84, 0xf8, 0xb9, 0xb8, 0x5c, 0x9d, 0xe3,
	0x83, 0x43, 0x9d, 0x9d, 0x5d, 0x83, 0x83, 0x05, 0x4e, 0x30, 0x0a, 0x09, 0x73, 0xf1, 0xb9, 0x3a,
	0xc7, 0x3b, 0x85, 0x06, 0x7b, 0xfa, 0xb9, 0x06, 0x07, 0xbb, 0x06, 0x05, 0x09, 0x4c, 0x60, 0x16,
	0x12, 0xe1, 0xe2, 0x07, 0xeb, 0x71, 0x0c, 0x0d, 0xf1, 0xf0, 0x0f, 0xf2, 0x8c, 0x72, 0x75, 0x11,
	0x98, 0xc4, 0x2c, 0x24, 0xc6, 0x25, 0x08, 0xd2, 0xeb, 0x1a, 0x14, 0xe6, 0x1a, 0x14, 0x1f, 0xea,
	0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xf0, 0x85, 0xd9, 0x89, 0x3b, 0x8a, 0x53, 0x4f, 0xdf, 0x1a,
	0xe4, 0x8a, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0x13, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9c,
	0xc0, 0xde, 0x0f, 0x9a, 0x00, 0x00, 0x00,
}
