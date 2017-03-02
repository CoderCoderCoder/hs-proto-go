// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/pegasus/bobnet/bobnet.proto
// DO NOT EDIT!

/*
Package bobnet is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/pegasus/bobnet/bobnet.proto

It has these top-level messages:
	PACKETTYPES
	DebugConsoleGetCmdList
	DebugConsoleCommand
	DebugConsoleUpdateFromPane
	DebugConsoleGetZones
	DebugConsoleCmdList
	DebugConsoleResponse
	DebugPaneNewItems
	DebugPaneDelItems
	DebugConsoleZones
	Deadend
	DeadendUtil
*/
package bobnet

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

type PACKETTYPES_BobNetCount int32

const (
	PACKETTYPES_COUNT PACKETTYPES_BobNetCount = 500
)

var PACKETTYPES_BobNetCount_name = map[int32]string{
	500: "COUNT",
}
var PACKETTYPES_BobNetCount_value = map[string]int32{
	"COUNT": 500,
}

func (x PACKETTYPES_BobNetCount) Enum() *PACKETTYPES_BobNetCount {
	p := new(PACKETTYPES_BobNetCount)
	*p = x
	return p
}
func (x PACKETTYPES_BobNetCount) String() string {
	return proto.EnumName(PACKETTYPES_BobNetCount_name, int32(x))
}
func (x *PACKETTYPES_BobNetCount) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PACKETTYPES_BobNetCount_value, data, "PACKETTYPES_BobNetCount")
	if err != nil {
		return err
	}
	*x = PACKETTYPES_BobNetCount(value)
	return nil
}
func (PACKETTYPES_BobNetCount) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type DebugConsoleGetCmdList_PacketID int32

const (
	DebugConsoleGetCmdList_ID DebugConsoleGetCmdList_PacketID = 125
)

var DebugConsoleGetCmdList_PacketID_name = map[int32]string{
	125: "ID",
}
var DebugConsoleGetCmdList_PacketID_value = map[string]int32{
	"ID": 125,
}

func (x DebugConsoleGetCmdList_PacketID) Enum() *DebugConsoleGetCmdList_PacketID {
	p := new(DebugConsoleGetCmdList_PacketID)
	*p = x
	return p
}
func (x DebugConsoleGetCmdList_PacketID) String() string {
	return proto.EnumName(DebugConsoleGetCmdList_PacketID_name, int32(x))
}
func (x *DebugConsoleGetCmdList_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleGetCmdList_PacketID_value, data, "DebugConsoleGetCmdList_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleGetCmdList_PacketID(value)
	return nil
}
func (DebugConsoleGetCmdList_PacketID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type DebugConsoleCommand_PacketID int32

const (
	DebugConsoleCommand_ID DebugConsoleCommand_PacketID = 123
)

var DebugConsoleCommand_PacketID_name = map[int32]string{
	123: "ID",
}
var DebugConsoleCommand_PacketID_value = map[string]int32{
	"ID": 123,
}

func (x DebugConsoleCommand_PacketID) Enum() *DebugConsoleCommand_PacketID {
	p := new(DebugConsoleCommand_PacketID)
	*p = x
	return p
}
func (x DebugConsoleCommand_PacketID) String() string {
	return proto.EnumName(DebugConsoleCommand_PacketID_name, int32(x))
}
func (x *DebugConsoleCommand_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleCommand_PacketID_value, data, "DebugConsoleCommand_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleCommand_PacketID(value)
	return nil
}
func (DebugConsoleCommand_PacketID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

type DebugConsoleUpdateFromPane_PacketID int32

const (
	DebugConsoleUpdateFromPane_ID DebugConsoleUpdateFromPane_PacketID = 145
)

var DebugConsoleUpdateFromPane_PacketID_name = map[int32]string{
	145: "ID",
}
var DebugConsoleUpdateFromPane_PacketID_value = map[string]int32{
	"ID": 145,
}

func (x DebugConsoleUpdateFromPane_PacketID) Enum() *DebugConsoleUpdateFromPane_PacketID {
	p := new(DebugConsoleUpdateFromPane_PacketID)
	*p = x
	return p
}
func (x DebugConsoleUpdateFromPane_PacketID) String() string {
	return proto.EnumName(DebugConsoleUpdateFromPane_PacketID_name, int32(x))
}
func (x *DebugConsoleUpdateFromPane_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleUpdateFromPane_PacketID_value, data, "DebugConsoleUpdateFromPane_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleUpdateFromPane_PacketID(value)
	return nil
}
func (DebugConsoleUpdateFromPane_PacketID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type DebugConsoleGetZones_PacketID int32

const (
	DebugConsoleGetZones_ID DebugConsoleGetZones_PacketID = 147
)

var DebugConsoleGetZones_PacketID_name = map[int32]string{
	147: "ID",
}
var DebugConsoleGetZones_PacketID_value = map[string]int32{
	"ID": 147,
}

func (x DebugConsoleGetZones_PacketID) Enum() *DebugConsoleGetZones_PacketID {
	p := new(DebugConsoleGetZones_PacketID)
	*p = x
	return p
}
func (x DebugConsoleGetZones_PacketID) String() string {
	return proto.EnumName(DebugConsoleGetZones_PacketID_name, int32(x))
}
func (x *DebugConsoleGetZones_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleGetZones_PacketID_value, data, "DebugConsoleGetZones_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleGetZones_PacketID(value)
	return nil
}
func (DebugConsoleGetZones_PacketID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type DebugConsoleCmdList_PacketID int32

const (
	DebugConsoleCmdList_ID DebugConsoleCmdList_PacketID = 146
)

var DebugConsoleCmdList_PacketID_name = map[int32]string{
	146: "ID",
}
var DebugConsoleCmdList_PacketID_value = map[string]int32{
	"ID": 146,
}

func (x DebugConsoleCmdList_PacketID) Enum() *DebugConsoleCmdList_PacketID {
	p := new(DebugConsoleCmdList_PacketID)
	*p = x
	return p
}
func (x DebugConsoleCmdList_PacketID) String() string {
	return proto.EnumName(DebugConsoleCmdList_PacketID_name, int32(x))
}
func (x *DebugConsoleCmdList_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleCmdList_PacketID_value, data, "DebugConsoleCmdList_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleCmdList_PacketID(value)
	return nil
}
func (DebugConsoleCmdList_PacketID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

type DebugConsoleResponse_ResponseType int32

const (
	DebugConsoleResponse_CONSOLE_OUTPUT DebugConsoleResponse_ResponseType = 0
	DebugConsoleResponse_LOG_MESSAGE    DebugConsoleResponse_ResponseType = 1
)

var DebugConsoleResponse_ResponseType_name = map[int32]string{
	0: "CONSOLE_OUTPUT",
	1: "LOG_MESSAGE",
}
var DebugConsoleResponse_ResponseType_value = map[string]int32{
	"CONSOLE_OUTPUT": 0,
	"LOG_MESSAGE":    1,
}

func (x DebugConsoleResponse_ResponseType) Enum() *DebugConsoleResponse_ResponseType {
	p := new(DebugConsoleResponse_ResponseType)
	*p = x
	return p
}
func (x DebugConsoleResponse_ResponseType) String() string {
	return proto.EnumName(DebugConsoleResponse_ResponseType_name, int32(x))
}
func (x *DebugConsoleResponse_ResponseType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleResponse_ResponseType_value, data, "DebugConsoleResponse_ResponseType")
	if err != nil {
		return err
	}
	*x = DebugConsoleResponse_ResponseType(value)
	return nil
}
func (DebugConsoleResponse_ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type DebugConsoleResponse_PacketID int32

const (
	DebugConsoleResponse_ID DebugConsoleResponse_PacketID = 124
)

var DebugConsoleResponse_PacketID_name = map[int32]string{
	124: "ID",
}
var DebugConsoleResponse_PacketID_value = map[string]int32{
	"ID": 124,
}

func (x DebugConsoleResponse_PacketID) Enum() *DebugConsoleResponse_PacketID {
	p := new(DebugConsoleResponse_PacketID)
	*p = x
	return p
}
func (x DebugConsoleResponse_PacketID) String() string {
	return proto.EnumName(DebugConsoleResponse_PacketID_name, int32(x))
}
func (x *DebugConsoleResponse_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleResponse_PacketID_value, data, "DebugConsoleResponse_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleResponse_PacketID(value)
	return nil
}
func (DebugConsoleResponse_PacketID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 1}
}

type DebugPaneNewItems_PacketID int32

const (
	DebugPaneNewItems_ID DebugPaneNewItems_PacketID = 142
)

var DebugPaneNewItems_PacketID_name = map[int32]string{
	142: "ID",
}
var DebugPaneNewItems_PacketID_value = map[string]int32{
	"ID": 142,
}

func (x DebugPaneNewItems_PacketID) Enum() *DebugPaneNewItems_PacketID {
	p := new(DebugPaneNewItems_PacketID)
	*p = x
	return p
}
func (x DebugPaneNewItems_PacketID) String() string {
	return proto.EnumName(DebugPaneNewItems_PacketID_name, int32(x))
}
func (x *DebugPaneNewItems_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugPaneNewItems_PacketID_value, data, "DebugPaneNewItems_PacketID")
	if err != nil {
		return err
	}
	*x = DebugPaneNewItems_PacketID(value)
	return nil
}
func (DebugPaneNewItems_PacketID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

type DebugPaneDelItems_PacketID int32

const (
	DebugPaneDelItems_ID DebugPaneDelItems_PacketID = 143
)

var DebugPaneDelItems_PacketID_name = map[int32]string{
	143: "ID",
}
var DebugPaneDelItems_PacketID_value = map[string]int32{
	"ID": 143,
}

func (x DebugPaneDelItems_PacketID) Enum() *DebugPaneDelItems_PacketID {
	p := new(DebugPaneDelItems_PacketID)
	*p = x
	return p
}
func (x DebugPaneDelItems_PacketID) String() string {
	return proto.EnumName(DebugPaneDelItems_PacketID_name, int32(x))
}
func (x *DebugPaneDelItems_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugPaneDelItems_PacketID_value, data, "DebugPaneDelItems_PacketID")
	if err != nil {
		return err
	}
	*x = DebugPaneDelItems_PacketID(value)
	return nil
}
func (DebugPaneDelItems_PacketID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

type DebugConsoleZones_PacketID int32

const (
	DebugConsoleZones_ID DebugConsoleZones_PacketID = 148
)

var DebugConsoleZones_PacketID_name = map[int32]string{
	148: "ID",
}
var DebugConsoleZones_PacketID_value = map[string]int32{
	"ID": 148,
}

func (x DebugConsoleZones_PacketID) Enum() *DebugConsoleZones_PacketID {
	p := new(DebugConsoleZones_PacketID)
	*p = x
	return p
}
func (x DebugConsoleZones_PacketID) String() string {
	return proto.EnumName(DebugConsoleZones_PacketID_name, int32(x))
}
func (x *DebugConsoleZones_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleZones_PacketID_value, data, "DebugConsoleZones_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleZones_PacketID(value)
	return nil
}
func (DebugConsoleZones_PacketID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{9, 0}
}

type Deadend_PacketID int32

const (
	Deadend_ID Deadend_PacketID = 169
)

var Deadend_PacketID_name = map[int32]string{
	169: "ID",
}
var Deadend_PacketID_value = map[string]int32{
	"ID": 169,
}

func (x Deadend_PacketID) Enum() *Deadend_PacketID {
	p := new(Deadend_PacketID)
	*p = x
	return p
}
func (x Deadend_PacketID) String() string {
	return proto.EnumName(Deadend_PacketID_name, int32(x))
}
func (x *Deadend_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Deadend_PacketID_value, data, "Deadend_PacketID")
	if err != nil {
		return err
	}
	*x = Deadend_PacketID(value)
	return nil
}
func (Deadend_PacketID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{10, 0} }

type DeadendUtil_PacketID int32

const (
	DeadendUtil_ID DeadendUtil_PacketID = 167
)

var DeadendUtil_PacketID_name = map[int32]string{
	167: "ID",
}
var DeadendUtil_PacketID_value = map[string]int32{
	"ID": 167,
}

func (x DeadendUtil_PacketID) Enum() *DeadendUtil_PacketID {
	p := new(DeadendUtil_PacketID)
	*p = x
	return p
}
func (x DeadendUtil_PacketID) String() string {
	return proto.EnumName(DeadendUtil_PacketID_name, int32(x))
}
func (x *DeadendUtil_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeadendUtil_PacketID_value, data, "DeadendUtil_PacketID")
	if err != nil {
		return err
	}
	*x = DeadendUtil_PacketID(value)
	return nil
}
func (DeadendUtil_PacketID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

type PACKETTYPES struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PACKETTYPES) Reset()                    { *m = PACKETTYPES{} }
func (m *PACKETTYPES) String() string            { return proto.CompactTextString(m) }
func (*PACKETTYPES) ProtoMessage()               {}
func (*PACKETTYPES) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DebugConsoleGetCmdList struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DebugConsoleGetCmdList) Reset()                    { *m = DebugConsoleGetCmdList{} }
func (m *DebugConsoleGetCmdList) String() string            { return proto.CompactTextString(m) }
func (*DebugConsoleGetCmdList) ProtoMessage()               {}
func (*DebugConsoleGetCmdList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DebugConsoleCommand struct {
	Command          *string `protobuf:"bytes,1,req,name=command" json:"command,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugConsoleCommand) Reset()                    { *m = DebugConsoleCommand{} }
func (m *DebugConsoleCommand) String() string            { return proto.CompactTextString(m) }
func (*DebugConsoleCommand) ProtoMessage()               {}
func (*DebugConsoleCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DebugConsoleCommand) GetCommand() string {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return ""
}

type DebugConsoleUpdateFromPane struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugConsoleUpdateFromPane) Reset()                    { *m = DebugConsoleUpdateFromPane{} }
func (m *DebugConsoleUpdateFromPane) String() string            { return proto.CompactTextString(m) }
func (*DebugConsoleUpdateFromPane) ProtoMessage()               {}
func (*DebugConsoleUpdateFromPane) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DebugConsoleUpdateFromPane) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DebugConsoleUpdateFromPane) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type DebugConsoleGetZones struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DebugConsoleGetZones) Reset()                    { *m = DebugConsoleGetZones{} }
func (m *DebugConsoleGetZones) String() string            { return proto.CompactTextString(m) }
func (*DebugConsoleGetZones) ProtoMessage()               {}
func (*DebugConsoleGetZones) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DebugConsoleCmdList struct {
	Commands         []*DebugConsoleCmdList_DebugConsoleCmd `protobuf:"bytes,1,rep,name=commands" json:"commands,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *DebugConsoleCmdList) Reset()                    { *m = DebugConsoleCmdList{} }
func (m *DebugConsoleCmdList) String() string            { return proto.CompactTextString(m) }
func (*DebugConsoleCmdList) ProtoMessage()               {}
func (*DebugConsoleCmdList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DebugConsoleCmdList) GetCommands() []*DebugConsoleCmdList_DebugConsoleCmd {
	if m != nil {
		return m.Commands
	}
	return nil
}

type DebugConsoleCmdList_DebugConsoleCmd struct {
	Name             *string                                                     `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Params           []*DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam `protobuf:"bytes,2,rep,name=params" json:"params,omitempty"`
	XXX_unrecognized []byte                                                      `json:"-"`
}

func (m *DebugConsoleCmdList_DebugConsoleCmd) Reset()         { *m = DebugConsoleCmdList_DebugConsoleCmd{} }
func (m *DebugConsoleCmdList_DebugConsoleCmd) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleCmdList_DebugConsoleCmd) ProtoMessage()    {}
func (*DebugConsoleCmdList_DebugConsoleCmd) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0}
}

func (m *DebugConsoleCmdList_DebugConsoleCmd) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DebugConsoleCmdList_DebugConsoleCmd) GetParams() []*DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam {
	if m != nil {
		return m.Params
	}
	return nil
}

type DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam struct {
	ParamType        *string `protobuf:"bytes,1,req,name=param_type,json=paramType" json:"param_type,omitempty"`
	ParamName        *string `protobuf:"bytes,2,req,name=param_name,json=paramName" json:"param_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam) Reset() {
	*m = DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam{}
}
func (m *DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam) String() string {
	return proto.CompactTextString(m)
}
func (*DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam) ProtoMessage() {}
func (*DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 0, 0}
}

func (m *DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam) GetParamType() string {
	if m != nil && m.ParamType != nil {
		return *m.ParamType
	}
	return ""
}

func (m *DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam) GetParamName() string {
	if m != nil && m.ParamName != nil {
		return *m.ParamName
	}
	return ""
}

type DebugConsoleResponse struct {
	Response         *string                            `protobuf:"bytes,1,req,name=response" json:"response,omitempty"`
	ResponseType     *DebugConsoleResponse_ResponseType `protobuf:"varint,2,req,name=response_type,json=responseType,enum=bobnet.DebugConsoleResponse_ResponseType" json:"response_type,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *DebugConsoleResponse) Reset()                    { *m = DebugConsoleResponse{} }
func (m *DebugConsoleResponse) String() string            { return proto.CompactTextString(m) }
func (*DebugConsoleResponse) ProtoMessage()               {}
func (*DebugConsoleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DebugConsoleResponse) GetResponse() string {
	if m != nil && m.Response != nil {
		return *m.Response
	}
	return ""
}

func (m *DebugConsoleResponse) GetResponseType() DebugConsoleResponse_ResponseType {
	if m != nil && m.ResponseType != nil {
		return *m.ResponseType
	}
	return DebugConsoleResponse_CONSOLE_OUTPUT
}

type DebugPaneNewItems struct {
	Items            []*DebugPaneNewItems_DebugPaneNewItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *DebugPaneNewItems) Reset()                    { *m = DebugPaneNewItems{} }
func (m *DebugPaneNewItems) String() string            { return proto.CompactTextString(m) }
func (*DebugPaneNewItems) ProtoMessage()               {}
func (*DebugPaneNewItems) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DebugPaneNewItems) GetItems() []*DebugPaneNewItems_DebugPaneNewItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type DebugPaneNewItems_DebugPaneNewItem struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugPaneNewItems_DebugPaneNewItem) Reset()         { *m = DebugPaneNewItems_DebugPaneNewItem{} }
func (m *DebugPaneNewItems_DebugPaneNewItem) String() string { return proto.CompactTextString(m) }
func (*DebugPaneNewItems_DebugPaneNewItem) ProtoMessage()    {}
func (*DebugPaneNewItems_DebugPaneNewItem) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

func (m *DebugPaneNewItems_DebugPaneNewItem) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DebugPaneNewItems_DebugPaneNewItem) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type DebugPaneDelItems struct {
	Items            []*DebugPaneDelItems_DebugPaneDelItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *DebugPaneDelItems) Reset()                    { *m = DebugPaneDelItems{} }
func (m *DebugPaneDelItems) String() string            { return proto.CompactTextString(m) }
func (*DebugPaneDelItems) ProtoMessage()               {}
func (*DebugPaneDelItems) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DebugPaneDelItems) GetItems() []*DebugPaneDelItems_DebugPaneDelItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type DebugPaneDelItems_DebugPaneDelItem struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugPaneDelItems_DebugPaneDelItem) Reset()         { *m = DebugPaneDelItems_DebugPaneDelItem{} }
func (m *DebugPaneDelItems_DebugPaneDelItem) String() string { return proto.CompactTextString(m) }
func (*DebugPaneDelItems_DebugPaneDelItem) ProtoMessage()    {}
func (*DebugPaneDelItems_DebugPaneDelItem) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

func (m *DebugPaneDelItems_DebugPaneDelItem) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type DebugConsoleZones struct {
	Zones            []*DebugConsoleZones_DebugConsoleZone `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *DebugConsoleZones) Reset()                    { *m = DebugConsoleZones{} }
func (m *DebugConsoleZones) String() string            { return proto.CompactTextString(m) }
func (*DebugConsoleZones) ProtoMessage()               {}
func (*DebugConsoleZones) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DebugConsoleZones) GetZones() []*DebugConsoleZones_DebugConsoleZone {
	if m != nil {
		return m.Zones
	}
	return nil
}

type DebugConsoleZones_DebugConsoleZone struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Id               *uint32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugConsoleZones_DebugConsoleZone) Reset()         { *m = DebugConsoleZones_DebugConsoleZone{} }
func (m *DebugConsoleZones_DebugConsoleZone) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleZones_DebugConsoleZone) ProtoMessage()    {}
func (*DebugConsoleZones_DebugConsoleZone) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{9, 0}
}

func (m *DebugConsoleZones_DebugConsoleZone) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DebugConsoleZones_DebugConsoleZone) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type Deadend struct {
	Reply1           *string `protobuf:"bytes,1,opt,name=reply1" json:"reply1,omitempty"`
	Reply2           *string `protobuf:"bytes,2,opt,name=reply2" json:"reply2,omitempty"`
	Reply3           *string `protobuf:"bytes,3,opt,name=reply3" json:"reply3,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Deadend) Reset()                    { *m = Deadend{} }
func (m *Deadend) String() string            { return proto.CompactTextString(m) }
func (*Deadend) ProtoMessage()               {}
func (*Deadend) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Deadend) GetReply1() string {
	if m != nil && m.Reply1 != nil {
		return *m.Reply1
	}
	return ""
}

func (m *Deadend) GetReply2() string {
	if m != nil && m.Reply2 != nil {
		return *m.Reply2
	}
	return ""
}

func (m *Deadend) GetReply3() string {
	if m != nil && m.Reply3 != nil {
		return *m.Reply3
	}
	return ""
}

type DeadendUtil struct {
	Reply1           *string `protobuf:"bytes,1,opt,name=reply1" json:"reply1,omitempty"`
	Reply2           *string `protobuf:"bytes,2,opt,name=reply2" json:"reply2,omitempty"`
	Reply3           *string `protobuf:"bytes,3,opt,name=reply3" json:"reply3,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeadendUtil) Reset()                    { *m = DeadendUtil{} }
func (m *DeadendUtil) String() string            { return proto.CompactTextString(m) }
func (*DeadendUtil) ProtoMessage()               {}
func (*DeadendUtil) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeadendUtil) GetReply1() string {
	if m != nil && m.Reply1 != nil {
		return *m.Reply1
	}
	return ""
}

func (m *DeadendUtil) GetReply2() string {
	if m != nil && m.Reply2 != nil {
		return *m.Reply2
	}
	return ""
}

func (m *DeadendUtil) GetReply3() string {
	if m != nil && m.Reply3 != nil {
		return *m.Reply3
	}
	return ""
}

func init() {
	proto.RegisterType((*PACKETTYPES)(nil), "bobnet.PACKETTYPES")
	proto.RegisterType((*DebugConsoleGetCmdList)(nil), "bobnet.DebugConsoleGetCmdList")
	proto.RegisterType((*DebugConsoleCommand)(nil), "bobnet.DebugConsoleCommand")
	proto.RegisterType((*DebugConsoleUpdateFromPane)(nil), "bobnet.DebugConsoleUpdateFromPane")
	proto.RegisterType((*DebugConsoleGetZones)(nil), "bobnet.DebugConsoleGetZones")
	proto.RegisterType((*DebugConsoleCmdList)(nil), "bobnet.DebugConsoleCmdList")
	proto.RegisterType((*DebugConsoleCmdList_DebugConsoleCmd)(nil), "bobnet.DebugConsoleCmdList.DebugConsoleCmd")
	proto.RegisterType((*DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam)(nil), "bobnet.DebugConsoleCmdList.DebugConsoleCmd.DebugConsoleCmdParam")
	proto.RegisterType((*DebugConsoleResponse)(nil), "bobnet.DebugConsoleResponse")
	proto.RegisterType((*DebugPaneNewItems)(nil), "bobnet.DebugPaneNewItems")
	proto.RegisterType((*DebugPaneNewItems_DebugPaneNewItem)(nil), "bobnet.DebugPaneNewItems.DebugPaneNewItem")
	proto.RegisterType((*DebugPaneDelItems)(nil), "bobnet.DebugPaneDelItems")
	proto.RegisterType((*DebugPaneDelItems_DebugPaneDelItem)(nil), "bobnet.DebugPaneDelItems.DebugPaneDelItem")
	proto.RegisterType((*DebugConsoleZones)(nil), "bobnet.DebugConsoleZones")
	proto.RegisterType((*DebugConsoleZones_DebugConsoleZone)(nil), "bobnet.DebugConsoleZones.DebugConsoleZone")
	proto.RegisterType((*Deadend)(nil), "bobnet.Deadend")
	proto.RegisterType((*DeadendUtil)(nil), "bobnet.DeadendUtil")
	proto.RegisterEnum("bobnet.PACKETTYPES_BobNetCount", PACKETTYPES_BobNetCount_name, PACKETTYPES_BobNetCount_value)
	proto.RegisterEnum("bobnet.DebugConsoleGetCmdList_PacketID", DebugConsoleGetCmdList_PacketID_name, DebugConsoleGetCmdList_PacketID_value)
	proto.RegisterEnum("bobnet.DebugConsoleCommand_PacketID", DebugConsoleCommand_PacketID_name, DebugConsoleCommand_PacketID_value)
	proto.RegisterEnum("bobnet.DebugConsoleUpdateFromPane_PacketID", DebugConsoleUpdateFromPane_PacketID_name, DebugConsoleUpdateFromPane_PacketID_value)
	proto.RegisterEnum("bobnet.DebugConsoleGetZones_PacketID", DebugConsoleGetZones_PacketID_name, DebugConsoleGetZones_PacketID_value)
	proto.RegisterEnum("bobnet.DebugConsoleCmdList_PacketID", DebugConsoleCmdList_PacketID_name, DebugConsoleCmdList_PacketID_value)
	proto.RegisterEnum("bobnet.DebugConsoleResponse_ResponseType", DebugConsoleResponse_ResponseType_name, DebugConsoleResponse_ResponseType_value)
	proto.RegisterEnum("bobnet.DebugConsoleResponse_PacketID", DebugConsoleResponse_PacketID_name, DebugConsoleResponse_PacketID_value)
	proto.RegisterEnum("bobnet.DebugPaneNewItems_PacketID", DebugPaneNewItems_PacketID_name, DebugPaneNewItems_PacketID_value)
	proto.RegisterEnum("bobnet.DebugPaneDelItems_PacketID", DebugPaneDelItems_PacketID_name, DebugPaneDelItems_PacketID_value)
	proto.RegisterEnum("bobnet.DebugConsoleZones_PacketID", DebugConsoleZones_PacketID_name, DebugConsoleZones_PacketID_value)
	proto.RegisterEnum("bobnet.Deadend_PacketID", Deadend_PacketID_name, Deadend_PacketID_value)
	proto.RegisterEnum("bobnet.DeadendUtil_PacketID", DeadendUtil_PacketID_name, DeadendUtil_PacketID_value)
}

func init() {
	proto.RegisterFile("github.com/HearthSim/hs-proto-go/pegasus/bobnet/bobnet.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x66, 0x5d, 0x9a, 0xb6, 0x93, 0xfe, 0x84, 0x6d, 0x55, 0x99, 0x48, 0x48, 0xd5, 0x1e, 0x50,
	0xa0, 0x34, 0x11, 0xa9, 0xc4, 0xa9, 0x87, 0x16, 0x27, 0x84, 0x8a, 0x92, 0x44, 0x8e, 0x23, 0xf1,
	0x73, 0xa8, 0x36, 0xf5, 0x92, 0x1a, 0x62, 0xaf, 0x65, 0x6f, 0x40, 0xe5, 0xe7, 0x15, 0x40, 0x14,
	0x9e, 0x80, 0x0b, 0xf0, 0x2c, 0x3d, 0xf3, 0x26, 0x3c, 0x00, 0xb2, 0xbd, 0x0e, 0xdb, 0x74, 0x0f,
	0x20, 0xc4, 0x29, 0x33, 0xdf, 0x37, 0x3b, 0xdf, 0x37, 0x93, 0x31, 0xec, 0x0c, 0x3d, 0x71, 0x3c,
	0x1e, 0x54, 0x8f, 0xb8, 0x5f, 0xbb, 0xcf, 0x68, 0x24, 0x8e, 0x7b, 0x9e, 0x5f, 0x3b, 0x8e, 0xb7,
	0xc2, 0x88, 0x0b, 0xbe, 0x35, 0xe4, 0xb5, 0x90, 0x0d, 0x69, 0x3c, 0x8e, 0x6b, 0x03, 0x3e, 0x08,
	0x98, 0x90, 0x3f, 0xd5, 0x94, 0xc6, 0x85, 0x2c, 0x23, 0x15, 0x28, 0x76, 0xf7, 0xac, 0x07, 0x4d,
	0xc7, 0x79, 0xdc, 0x6d, 0xf6, 0xc8, 0x55, 0x28, 0xde, 0xe5, 0x83, 0x36, 0x13, 0x16, 0x1f, 0x07,
	0x02, 0x03, 0xcc, 0x5a, 0x9d, 0x7e, 0xdb, 0x29, 0xfd, 0x9c, 0x21, 0xb7, 0x60, 0xbd, 0xc1, 0x06,
	0xe3, 0xa1, 0xc5, 0x83, 0x98, 0x8f, 0x58, 0x8b, 0x09, 0xcb, 0x77, 0x0f, 0xbc, 0x58, 0x10, 0x0c,
	0xf3, 0x5d, 0x7a, 0xf4, 0x82, 0x89, 0xfd, 0x06, 0x2e, 0x80, 0xb1, 0xdf, 0x28, 0xbd, 0x23, 0x16,
	0xac, 0xaa, 0xd5, 0x16, 0xf7, 0x7d, 0x1a, 0xb8, 0xd8, 0x84, 0xb9, 0xa3, 0x2c, 0x34, 0xd1, 0x86,
	0x51, 0x59, 0xb0, 0xf3, 0x54, 0xd3, 0xe4, 0x0d, 0x79, 0x0a, 0x65, 0xb5, 0x49, 0x3f, 0x74, 0xa9,
	0x60, 0xf7, 0x22, 0xee, 0x77, 0x69, 0xc0, 0x30, 0x86, 0xcb, 0x01, 0xf5, 0x99, 0x6c, 0x94, 0xc6,
	0x78, 0x0d, 0x66, 0x5f, 0xd2, 0xd1, 0x98, 0x99, 0x46, 0x0a, 0x66, 0x09, 0x59, 0x55, 0x7a, 0xcf,
	0xa5, 0xbd, 0x3f, 0x22, 0xb2, 0x09, 0x6b, 0x53, 0xf3, 0x3c, 0xe1, 0x01, 0x8b, 0x35, 0xc5, 0x9f,
	0x10, 0x39, 0x33, 0xa6, 0xe6, 0xc9, 0x46, 0xc7, 0x2d, 0x98, 0x97, 0x03, 0xc4, 0x26, 0xda, 0x98,
	0xa9, 0x14, 0xeb, 0x9b, 0x55, 0xb9, 0x67, 0x4d, 0xf9, 0x34, 0x66, 0x4f, 0x1e, 0x97, 0x7f, 0x20,
	0x58, 0x99, 0x62, 0xb5, 0x03, 0x3e, 0x82, 0x42, 0x48, 0x23, 0xea, 0xc7, 0xa6, 0x91, 0xca, 0xed,
	0xfe, 0x85, 0xdc, 0x74, 0xde, 0x4d, 0x1a, 0xd9, 0xb2, 0x5f, 0xd9, 0x39, 0xbf, 0x8f, 0x9c, 0xc7,
	0xd7, 0x00, 0xd2, 0x8a, 0x43, 0x71, 0x12, 0xe6, 0x5e, 0x16, 0x52, 0xc4, 0x39, 0x09, 0xd9, 0x6f,
	0x3a, 0xb5, 0x6a, 0x28, 0x74, 0x9b, 0xfa, 0xba, 0xd5, 0x9f, 0x22, 0x72, 0x86, 0xce, 0x6b, 0xd9,
	0x2c, 0x0e, 0x79, 0x10, 0x33, 0x5c, 0x86, 0xf9, 0x48, 0xc6, 0x52, 0x69, 0x92, 0xe3, 0x36, 0x2c,
	0xe5, 0x71, 0x66, 0x25, 0xd1, 0x5a, 0xae, 0xdf, 0xd0, 0x2d, 0x20, 0x6f, 0x58, 0xcd, 0x83, 0xc4,
	0xaa, 0xbd, 0x18, 0x29, 0x19, 0xd9, 0x86, 0x45, 0x95, 0xc5, 0x18, 0x96, 0xad, 0x4e, 0xbb, 0xd7,
	0x39, 0x68, 0x1e, 0x76, 0xfa, 0x4e, 0xb7, 0xef, 0x94, 0x2e, 0xe1, 0x15, 0x28, 0x1e, 0x74, 0x5a,
	0x87, 0x0f, 0x9b, 0xbd, 0xde, 0x5e, 0xab, 0x59, 0x42, 0x9a, 0x2b, 0x7d, 0x4b, 0xbe, 0x21, 0xb8,
	0x92, 0x8a, 0x27, 0x57, 0xd9, 0x66, 0xaf, 0xf6, 0x05, 0xf3, 0x63, 0xbc, 0x0b, 0xb3, 0x5e, 0x12,
	0xc8, 0xb3, 0xb8, 0x79, 0xce, 0xa6, 0x5a, 0x79, 0x01, 0xb1, 0xb3, 0x87, 0xe5, 0x1d, 0x28, 0x4d,
	0x53, 0xff, 0x74, 0xf3, 0xef, 0x11, 0x39, 0x55, 0xad, 0x36, 0xd8, 0xe8, 0x0f, 0xad, 0xe6, 0x95,
	0x17, 0x90, 0xdc, 0xea, 0x75, 0xc5, 0xaa, 0xa4, 0x74, 0x56, 0x35, 0xa6, 0x3e, 0x20, 0xf2, 0x25,
	0x37, 0x25, 0xff, 0xbc, 0xf4, 0x33, 0x4c, 0x4c, 0xbd, 0x4e, 0x02, 0xad, 0x29, 0xb5, 0xf2, 0x02,
	0x62, 0x67, 0x0f, 0xcb, 0x77, 0xa4, 0x29, 0x85, 0xd2, 0xee, 0x6f, 0x19, 0x0c, 0xcf, 0x4d, 0x97,
	0xb7, 0x64, 0x1b, 0x9e, 0xab, 0x31, 0xf9, 0x19, 0x91, 0x67, 0x30, 0xd7, 0x60, 0xd4, 0x65, 0x81,
	0x8b, 0xd7, 0xa1, 0x10, 0xb1, 0x70, 0x74, 0x72, 0xdb, 0x44, 0x1b, 0xa8, 0xb2, 0x60, 0xcb, 0x6c,
	0x82, 0xd7, 0x4d, 0x43, 0xc1, 0xeb, 0x13, 0x7c, 0xdb, 0x9c, 0x51, 0xf0, 0x6d, 0x8d, 0xce, 0x77,
	0x44, 0x9e, 0x43, 0x51, 0xea, 0xf4, 0x85, 0x37, 0xfa, 0x8f, 0x5a, 0x5f, 0xd1, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x72, 0xf3, 0x65, 0x08, 0x42, 0x06, 0x00, 0x00,
}