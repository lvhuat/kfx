// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	code.proto
	error.proto
	grpc.proto
	types.proto

It has these top-level messages:
	Request
	Response
	SymbolEventMessage
	GetCommonSettingRequest
	GetCommonSettingResponse
	GetPositionsRequest
	GetPricesRequest
	ConfigServiceRequest
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Codes int32

const (
	Codes_Ping Codes = 0
	// 鉴权,配置
	Codes_ConfigService Codes = 10
	Codes_GetCoreStatus Codes = 11
	// 基本设置
	Codes_GetCommonSetting Codes = 1001
	// 管理员权限
	Codes_GetManagerSetting Codes = 1002
	// 节假日设置
	Codes_GetHolidaySetting Codes = 1003
	// 时间设置
	Codes_GetTimeSetting Codes = 1004
	// 申请账户
	Codes_Apply Codes = 1101
	// 获取账户信息
	Codes_GetAccountInfo Codes = 1102
	// 获取账户入金
	Codes_GetAccountMargin Codes = 1103
	// 检验密码
	Codes_CheckPassword Codes = 1104
	// 更新密码
	Codes_UpdatePassword Codes = 1105
	// 获取账户组
	Codes_GetGroups Codes = 1106
	Codes_GetGroup  Codes = 1107
	// 获取品种配置
	Codes_GetSymbols Codes = 1201
	// 获取K线
	Codes_GetCharts Codes = 1205
	// 获取当前报价
	Codes_GetPrices Codes = 1210
	// 交易
	Codes_GetOrders Codes = 1315
	// 获取当前持仓
	Codes_GetPositions Codes = 1320
	// 获取历史平仓
	Codes_GetClosed Codes = 1325
	// 开市价单
	Codes_OpenMarket Codes = 1330
	// 开限价单
	Codes_OpenLimit Codes = 1331
	// 开止损单
	Codes_OpenStop Codes = 1332
	// 更新订单
	Codes_UpdateOrder Codes = 1333
	// 更新持仓单
	Codes_UpdatePosition Codes = 1334
	// 平仓
	Codes_ClosePosition Codes = 1335
	// 入金
	Codes_Deposit Codes = 1350
	// 出金
	Codes_Withdraw Codes = 1351
	// 入信用金
	Codes_CreditIn Codes = 1355
	// 出信用金
	Codes_CreditOut Codes = 1356
	// 获取历史出入金记录
	Codes_GetHistoryCashflows Codes = 1370
	// 获取历史信用金操作记录
	Codes_GetHistoryCredits Codes = 1371
	// 工具类
	Codes_CalPositionProfit Codes = 1401
	// 交易推送
	Codes_TradeEvent Codes = 4001
	// 出入金推送
	Codes_CashflowEvent Codes = 4010
	// 报价推送
	Codes_PriceEvent Codes = 4009
	// 组配置推送
	Codes_GroupEvent Codes = 4005
	// 账户推送
	Codes_AccountEvent Codes = 4006
	// 品种推送
	Codes_SymbolEvent Codes = 4008
)

var Codes_name = map[int32]string{
	0:    "Ping",
	10:   "ConfigService",
	11:   "GetCoreStatus",
	1001: "GetCommonSetting",
	1002: "GetManagerSetting",
	1003: "GetHolidaySetting",
	1004: "GetTimeSetting",
	1101: "Apply",
	1102: "GetAccountInfo",
	1103: "GetAccountMargin",
	1104: "CheckPassword",
	1105: "UpdatePassword",
	1106: "GetGroups",
	1107: "GetGroup",
	1201: "GetSymbols",
	1205: "GetCharts",
	1210: "GetPrices",
	1315: "GetOrders",
	1320: "GetPositions",
	1325: "GetClosed",
	1330: "OpenMarket",
	1331: "OpenLimit",
	1332: "OpenStop",
	1333: "UpdateOrder",
	1334: "UpdatePosition",
	1335: "ClosePosition",
	1350: "Deposit",
	1351: "Withdraw",
	1355: "CreditIn",
	1356: "CreditOut",
	1370: "GetHistoryCashflows",
	1371: "GetHistoryCredits",
	1401: "CalPositionProfit",
	4001: "TradeEvent",
	4010: "CashflowEvent",
	4009: "PriceEvent",
	4005: "GroupEvent",
	4006: "AccountEvent",
	4008: "SymbolEvent",
}
var Codes_value = map[string]int32{
	"Ping":                0,
	"ConfigService":       10,
	"GetCoreStatus":       11,
	"GetCommonSetting":    1001,
	"GetManagerSetting":   1002,
	"GetHolidaySetting":   1003,
	"GetTimeSetting":      1004,
	"Apply":               1101,
	"GetAccountInfo":      1102,
	"GetAccountMargin":    1103,
	"CheckPassword":       1104,
	"UpdatePassword":      1105,
	"GetGroups":           1106,
	"GetGroup":            1107,
	"GetSymbols":          1201,
	"GetCharts":           1205,
	"GetPrices":           1210,
	"GetOrders":           1315,
	"GetPositions":        1320,
	"GetClosed":           1325,
	"OpenMarket":          1330,
	"OpenLimit":           1331,
	"OpenStop":            1332,
	"UpdateOrder":         1333,
	"UpdatePosition":      1334,
	"ClosePosition":       1335,
	"Deposit":             1350,
	"Withdraw":            1351,
	"CreditIn":            1355,
	"CreditOut":           1356,
	"GetHistoryCashflows": 1370,
	"GetHistoryCredits":   1371,
	"CalPositionProfit":   1401,
	"TradeEvent":          4001,
	"CashflowEvent":       4010,
	"PriceEvent":          4009,
	"GroupEvent":          4005,
	"AccountEvent":        4006,
	"SymbolEvent":         4008,
}

func (x Codes) String() string {
	return proto.EnumName(Codes_name, int32(x))
}
func (Codes) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("pb.Codes", Codes_name, Codes_value)
}

func init() { proto.RegisterFile("code.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x4b, 0x6e, 0x13, 0x41,
	0x10, 0x86, 0x11, 0xc2, 0x78, 0x52, 0x89, 0x43, 0xa5, 0x23, 0x10, 0x3b, 0xef, 0x59, 0xb0, 0xe1,
	0x04, 0xd1, 0x80, 0x86, 0x48, 0x58, 0xb6, 0xe4, 0x20, 0xd6, 0xed, 0x99, 0xb2, 0xdd, 0xca, 0xb8,
	0x6b, 0xd4, 0x5d, 0x8e, 0xe5, 0xe3, 0x20, 0x04, 0x0a, 0x48, 0x2c, 0x78, 0x04, 0x24, 0x0e, 0x00,
	0x0b, 0xde, 0x8f, 0x15, 0x9c, 0x80, 0xc7, 0x05, 0xd8, 0xa1, 0xee, 0x99, 0x89, 0x58, 0xfe, 0x5f,
	0x57, 0xd5, 0x5f, 0x8f, 0x06, 0xc8, 0xb9, 0xa0, 0xab, 0x95, 0x63, 0x61, 0x75, 0xb6, 0x9a, 0x5c,
	0xf9, 0xd6, 0x81, 0x4e, 0xca, 0x05, 0x79, 0x95, 0xc0, 0xb9, 0x91, 0xb1, 0x33, 0x3c, 0xa3, 0x76,
	0xa0, 0x97, 0xb2, 0x9d, 0x9a, 0xd9, 0x98, 0xdc, 0x91, 0xc9, 0x09, 0x21, 0xa0, 0x8c, 0x24, 0x65,
	0x47, 0x63, 0xd1, 0xb2, 0xf4, 0xb8, 0xa9, 0x2e, 0x02, 0x46, 0xb4, 0x58, 0xb0, 0x1d, 0x93, 0x48,
	0xc8, 0xfd, 0xd9, 0x55, 0x97, 0x60, 0x27, 0x23, 0x19, 0x68, 0xab, 0x67, 0xe4, 0x5a, 0xfe, 0xab,
	0xe5, 0x37, 0xb9, 0x34, 0x85, 0x5e, 0xb7, 0xfc, 0x77, 0x57, 0xed, 0xc2, 0x76, 0x46, 0x72, 0x60,
	0x16, 0xd4, 0xc2, 0x3f, 0x5d, 0x05, 0xd0, 0xd9, 0xab, 0xaa, 0x72, 0x8d, 0xef, 0x93, 0x26, 0x60,
	0x2f, 0xcf, 0x79, 0x69, 0x65, 0xdf, 0x4e, 0x19, 0x3f, 0x24, 0x8d, 0x79, 0x03, 0x07, 0xda, 0xcd,
	0x8c, 0xc5, 0x8f, 0x89, 0x52, 0xd0, 0x4b, 0xe7, 0x94, 0x1f, 0x8e, 0xb4, 0xf7, 0x2b, 0x76, 0x05,
	0x7e, 0x8a, 0xf9, 0xb7, 0xab, 0x42, 0x0b, 0x9d, 0xc2, 0xcf, 0x89, 0xda, 0x86, 0x8d, 0x8c, 0x24,
	0x73, 0xbc, 0xac, 0x3c, 0x7e, 0x49, 0x54, 0x0f, 0x92, 0x56, 0xe3, 0xd7, 0x44, 0x5d, 0x00, 0xc8,
	0x48, 0xc6, 0xeb, 0xc5, 0x84, 0x4b, 0x8f, 0x4f, 0x36, 0x9a, 0xf8, 0x74, 0xae, 0x9d, 0x78, 0x3c,
	0x69, 0xf5, 0xc8, 0x99, 0x9c, 0x3c, 0xbe, 0x6a, 0xf5, 0xd0, 0x15, 0xe4, 0x3c, 0xde, 0x0b, 0xfb,
	0xda, 0x0a, 0xef, 0xec, 0x8d, 0x18, 0xb6, 0x1e, 0x8f, 0xa1, 0x2d, 0x51, 0xb2, 0xa7, 0x02, 0x1f,
	0x43, 0xf0, 0x18, 0x56, 0x64, 0x07, 0xda, 0x1d, 0x92, 0xe0, 0xd3, 0x18, 0x10, 0xc0, 0x2d, 0xb3,
	0x30, 0x82, 0xcf, 0x20, 0xf4, 0x14, 0xf4, 0x58, 0xb8, 0xc2, 0xe7, 0xa0, 0x10, 0x36, 0xeb, 0x39,
	0xa2, 0x0b, 0x9e, 0xc0, 0x7f, 0x93, 0x35, 0x3e, 0xf8, 0x02, 0xe2, 0x0a, 0x82, 0xc7, 0x29, 0x7b,
	0x09, 0x6a, 0x0b, 0xba, 0xd7, 0xa9, 0x0a, 0x00, 0x5f, 0xc7, 0xba, 0x77, 0x8c, 0xcc, 0x0b, 0xa7,
	0x57, 0xf8, 0x26, 0xca, 0xd4, 0x51, 0x61, 0x64, 0xdf, 0xe2, 0xdb, 0xd8, 0x45, 0x2d, 0x87, 0x4b,
	0xc1, 0x77, 0xa0, 0x2e, 0xc3, 0x6e, 0xb8, 0x9b, 0xf1, 0xc2, 0x6e, 0x9d, 0x6a, 0x3f, 0x9f, 0x96,
	0xbc, 0xf2, 0xf8, 0x1d, 0xda, 0x8b, 0x36, 0x2f, 0x31, 0xc7, 0xe3, 0x8f, 0xc8, 0x53, 0x5d, 0xb6,
	0xfe, 0x23, 0xc7, 0x53, 0x23, 0xf8, 0x37, 0x0e, 0x7c, 0xe0, 0x74, 0x41, 0x37, 0x8e, 0xc8, 0x0a,
	0xde, 0xed, 0xc7, 0x56, 0x9b, 0x82, 0x35, 0x7b, 0xd4, 0x0f, 0x41, 0x71, 0xab, 0x35, 0x78, 0x18,
	0x41, 0x3c, 0x4b, 0x0d, 0xee, 0xf7, 0xc3, 0x6a, 0x9b, 0xbb, 0xd7, 0xe8, 0x41, 0x3f, 0xac, 0xa6,
	0xbe, 0x55, 0x4d, 0x8e, 0xfb, 0x93, 0xf3, 0xf1, 0x87, 0x5f, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xb6, 0x8e, 0xff, 0x3c, 0xef, 0x02, 0x00, 0x00,
}
