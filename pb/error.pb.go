// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Errors int32

const (
	Errors_OK Errors = 0
	// 组件错误
	Errors_RedisNotWorking Errors = 10001
	Errors_MysqlNotWorking Errors = 10002
	Errors_CoreNotWorking  Errors = 10003
)

var Errors_name = map[int32]string{
	0:     "OK",
	10001: "RedisNotWorking",
	10002: "MysqlNotWorking",
	10003: "CoreNotWorking",
}
var Errors_value = map[string]int32{
	"OK":              0,
	"RedisNotWorking": 10001,
	"MysqlNotWorking": 10002,
	"CoreNotWorking":  10003,
}

func (x Errors) String() string {
	return proto.EnumName(Errors_name, int32(x))
}
func (Errors) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterEnum("pb.Errors", Errors_name, Errors_value)
}

func init() { proto.RegisterFile("error.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2d, 0x2a, 0xca,
	0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0xd2, 0x0a, 0xe4, 0x62, 0x73,
	0x05, 0x09, 0x15, 0x0b, 0xb1, 0x71, 0x31, 0xf9, 0x7b, 0x0b, 0x30, 0x08, 0x89, 0x70, 0xf1, 0x07,
	0xa5, 0xa6, 0x64, 0x16, 0xfb, 0xe5, 0x97, 0x84, 0xe7, 0x17, 0x65, 0x67, 0xe6, 0xa5, 0x0b, 0x4c,
	0xf4, 0x03, 0x89, 0xfa, 0x56, 0x16, 0x17, 0xe6, 0x20, 0x89, 0x4e, 0xf2, 0x13, 0x12, 0xe6, 0xe2,
	0x73, 0xce, 0x2f, 0x4a, 0x45, 0x12, 0x9c, 0xec, 0x97, 0xc4, 0x06, 0x36, 0xdd, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x84, 0x1d, 0x75, 0xb4, 0x6c, 0x00, 0x00, 0x00,
}
