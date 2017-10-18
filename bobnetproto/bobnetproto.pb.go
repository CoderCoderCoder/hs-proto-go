// Code generated by protoc-gen-go.
// source: bobnetproto/bobnetproto.proto
// DO NOT EDIT!

package bobnetproto

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// ref: BobNetProto.Deadend/PacketID
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
func (x Deadend_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *Deadend_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Deadend_PacketID_value, data, "Deadend_PacketID")
	if err != nil {
		return err
	}
	*x = Deadend_PacketID(value)
	return nil
}

// ref: BobNetProto.DeadendUtil/PacketID
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
func (x DeadendUtil_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DeadendUtil_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DeadendUtil_PacketID_value, data, "DeadendUtil_PacketID")
	if err != nil {
		return err
	}
	*x = DeadendUtil_PacketID(value)
	return nil
}

// ref: BobNetProto.DebugConsoleCmdList/PacketID
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
func (x DebugConsoleCmdList_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugConsoleCmdList_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleCmdList_PacketID_value, data, "DebugConsoleCmdList_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleCmdList_PacketID(value)
	return nil
}

// ref: BobNetProto.DebugConsoleCommand/PacketID
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
func (x DebugConsoleCommand_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugConsoleCommand_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleCommand_PacketID_value, data, "DebugConsoleCommand_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleCommand_PacketID(value)
	return nil
}

// ref: BobNetProto.DebugConsoleGetCmdList/PacketID
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
func (x DebugConsoleGetCmdList_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugConsoleGetCmdList_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleGetCmdList_PacketID_value, data, "DebugConsoleGetCmdList_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleGetCmdList_PacketID(value)
	return nil
}

// ref: BobNetProto.DebugConsoleGetZones/PacketID
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
func (x DebugConsoleGetZones_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugConsoleGetZones_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleGetZones_PacketID_value, data, "DebugConsoleGetZones_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleGetZones_PacketID(value)
	return nil
}

// ref: BobNetProto.DebugConsoleResponse/PacketID
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
func (x DebugConsoleResponse_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugConsoleResponse_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleResponse_PacketID_value, data, "DebugConsoleResponse_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleResponse_PacketID(value)
	return nil
}

// ref: BobNetProto.DebugConsoleResponse/ResponseType
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
func (x DebugConsoleResponse_ResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugConsoleResponse_ResponseType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleResponse_ResponseType_value, data, "DebugConsoleResponse_ResponseType")
	if err != nil {
		return err
	}
	*x = DebugConsoleResponse_ResponseType(value)
	return nil
}

// ref: BobNetProto.DebugConsoleUpdateFromPane/PacketID
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
func (x DebugConsoleUpdateFromPane_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugConsoleUpdateFromPane_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleUpdateFromPane_PacketID_value, data, "DebugConsoleUpdateFromPane_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleUpdateFromPane_PacketID(value)
	return nil
}

// ref: BobNetProto.DebugConsoleZones/PacketID
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
func (x DebugConsoleZones_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugConsoleZones_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugConsoleZones_PacketID_value, data, "DebugConsoleZones_PacketID")
	if err != nil {
		return err
	}
	*x = DebugConsoleZones_PacketID(value)
	return nil
}

// ref: BobNetProto.DebugPaneDelItems/PacketID
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
func (x DebugPaneDelItems_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugPaneDelItems_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugPaneDelItems_PacketID_value, data, "DebugPaneDelItems_PacketID")
	if err != nil {
		return err
	}
	*x = DebugPaneDelItems_PacketID(value)
	return nil
}

// ref: BobNetProto.DebugPaneNewItems/PacketID
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
func (x DebugPaneNewItems_PacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *DebugPaneNewItems_PacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DebugPaneNewItems_PacketID_value, data, "DebugPaneNewItems_PacketID")
	if err != nil {
		return err
	}
	*x = DebugPaneNewItems_PacketID(value)
	return nil
}

// ref: BobNetProto.PACKETTYPES/BobNetCount
type PACKETTYPES_BobNetCount int32

const (
	PACKETTYPES_COUNT PACKETTYPES_BobNetCount = 600
)

var PACKETTYPES_BobNetCount_name = map[int32]string{
	600: "COUNT",
}
var PACKETTYPES_BobNetCount_value = map[string]int32{
	"COUNT": 600,
}

func (x PACKETTYPES_BobNetCount) Enum() *PACKETTYPES_BobNetCount {
	p := new(PACKETTYPES_BobNetCount)
	*p = x
	return p
}
func (x PACKETTYPES_BobNetCount) String() string {
	return proto.EnumName(PACKETTYPES_BobNetCount_name, int32(x))
}
func (x PACKETTYPES_BobNetCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *PACKETTYPES_BobNetCount) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PACKETTYPES_BobNetCount_value, data, "PACKETTYPES_BobNetCount")
	if err != nil {
		return err
	}
	*x = PACKETTYPES_BobNetCount(value)
	return nil
}

// ref: BobNetProto.Deadend
type Deadend struct {
	Reply1           *string `protobuf:"bytes,1,opt,name=reply1" json:"reply1,omitempty"`
	Reply2           *string `protobuf:"bytes,2,opt,name=reply2" json:"reply2,omitempty"`
	Reply3           *string `protobuf:"bytes,3,opt,name=reply3" json:"reply3,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Deadend) Reset()         { *m = Deadend{} }
func (m *Deadend) String() string { return proto.CompactTextString(m) }
func (*Deadend) ProtoMessage()    {}

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

// ref: BobNetProto.DeadendUtil
type DeadendUtil struct {
	Reply1           *string `protobuf:"bytes,1,opt,name=reply1" json:"reply1,omitempty"`
	Reply2           *string `protobuf:"bytes,2,opt,name=reply2" json:"reply2,omitempty"`
	Reply3           *string `protobuf:"bytes,3,opt,name=reply3" json:"reply3,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeadendUtil) Reset()         { *m = DeadendUtil{} }
func (m *DeadendUtil) String() string { return proto.CompactTextString(m) }
func (*DeadendUtil) ProtoMessage()    {}

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

// ref: BobNetProto.DebugConsoleCmdList
type DebugConsoleCmdList struct {
	Commands         []*DebugConsoleCmdList_DebugConsoleCmd `protobuf:"bytes,1,rep,name=commands" json:"commands,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *DebugConsoleCmdList) Reset()         { *m = DebugConsoleCmdList{} }
func (m *DebugConsoleCmdList) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleCmdList) ProtoMessage()    {}

func (m *DebugConsoleCmdList) GetCommands() []*DebugConsoleCmdList_DebugConsoleCmd {
	if m != nil {
		return m.Commands
	}
	return nil
}

// ref: BobNetProto.DebugConsoleCmdList/DebugConsoleCmd
type DebugConsoleCmdList_DebugConsoleCmd struct {
	Name             *string                                     `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Params           []*DebugConsoleCmdList_DebugConsoleCmdParam `protobuf:"bytes,2,rep,name=params" json:"params,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *DebugConsoleCmdList_DebugConsoleCmd) Reset()         { *m = DebugConsoleCmdList_DebugConsoleCmd{} }
func (m *DebugConsoleCmdList_DebugConsoleCmd) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleCmdList_DebugConsoleCmd) ProtoMessage()    {}

func (m *DebugConsoleCmdList_DebugConsoleCmd) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DebugConsoleCmdList_DebugConsoleCmd) GetParams() []*DebugConsoleCmdList_DebugConsoleCmdParam {
	if m != nil {
		return m.Params
	}
	return nil
}

// ref: BobNetProto.DebugConsoleCmdList/DebugConsoleCmd/DebugConsoleCmdParam
type DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam struct {
	ParamType        *string `protobuf:"bytes,1,req,name=param_type" json:"param_type,omitempty"`
	ParamName        *string `protobuf:"bytes,2,req,name=param_name" json:"param_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam) Reset() {
	*m = DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam{}
}
func (m *DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam) String() string {
	return proto.CompactTextString(m)
}
func (*DebugConsoleCmdList_DebugConsoleCmd_DebugConsoleCmdParam) ProtoMessage() {}

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

// ref: BobNetProto.DebugConsoleCmdList/DebugConsoleCmd/DebugConsoleCmdParam
type DebugConsoleCmdList_DebugConsoleCmdParam struct {
	ParamType        *string `protobuf:"bytes,1,req,name=param_type" json:"param_type,omitempty"`
	ParamName        *string `protobuf:"bytes,2,req,name=param_name" json:"param_name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugConsoleCmdList_DebugConsoleCmdParam) Reset() {
	*m = DebugConsoleCmdList_DebugConsoleCmdParam{}
}
func (m *DebugConsoleCmdList_DebugConsoleCmdParam) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleCmdList_DebugConsoleCmdParam) ProtoMessage()    {}

func (m *DebugConsoleCmdList_DebugConsoleCmdParam) GetParamType() string {
	if m != nil && m.ParamType != nil {
		return *m.ParamType
	}
	return ""
}

func (m *DebugConsoleCmdList_DebugConsoleCmdParam) GetParamName() string {
	if m != nil && m.ParamName != nil {
		return *m.ParamName
	}
	return ""
}

// ref: BobNetProto.DebugConsoleCommand
type DebugConsoleCommand struct {
	Command          *string `protobuf:"bytes,1,req,name=command" json:"command,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugConsoleCommand) Reset()         { *m = DebugConsoleCommand{} }
func (m *DebugConsoleCommand) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleCommand) ProtoMessage()    {}

func (m *DebugConsoleCommand) GetCommand() string {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return ""
}

// ref: BobNetProto.DebugConsoleGetCmdList
type DebugConsoleGetCmdList struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DebugConsoleGetCmdList) Reset()         { *m = DebugConsoleGetCmdList{} }
func (m *DebugConsoleGetCmdList) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleGetCmdList) ProtoMessage()    {}

// ref: BobNetProto.DebugConsoleGetZones
type DebugConsoleGetZones struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DebugConsoleGetZones) Reset()         { *m = DebugConsoleGetZones{} }
func (m *DebugConsoleGetZones) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleGetZones) ProtoMessage()    {}

// ref: BobNetProto.DebugConsoleResponse
type DebugConsoleResponse struct {
	Response         *string                            `protobuf:"bytes,1,req,name=response" json:"response,omitempty"`
	ResponseType     *DebugConsoleResponse_ResponseType `protobuf:"varint,2,req,name=response_type,enum=bobnetproto.DebugConsoleResponse_ResponseType" json:"response_type,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *DebugConsoleResponse) Reset()         { *m = DebugConsoleResponse{} }
func (m *DebugConsoleResponse) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleResponse) ProtoMessage()    {}

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
	return 0
}

// ref: BobNetProto.DebugConsoleUpdateFromPane
type DebugConsoleUpdateFromPane struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugConsoleUpdateFromPane) Reset()         { *m = DebugConsoleUpdateFromPane{} }
func (m *DebugConsoleUpdateFromPane) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleUpdateFromPane) ProtoMessage()    {}

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

// ref: BobNetProto.DebugConsoleZones
type DebugConsoleZones struct {
	Zones            []*DebugConsoleZones_DebugConsoleZone `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *DebugConsoleZones) Reset()         { *m = DebugConsoleZones{} }
func (m *DebugConsoleZones) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleZones) ProtoMessage()    {}

func (m *DebugConsoleZones) GetZones() []*DebugConsoleZones_DebugConsoleZone {
	if m != nil {
		return m.Zones
	}
	return nil
}

// ref: BobNetProto.DebugConsoleZones/DebugConsoleZone
type DebugConsoleZones_DebugConsoleZone struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Id               *uint32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugConsoleZones_DebugConsoleZone) Reset()         { *m = DebugConsoleZones_DebugConsoleZone{} }
func (m *DebugConsoleZones_DebugConsoleZone) String() string { return proto.CompactTextString(m) }
func (*DebugConsoleZones_DebugConsoleZone) ProtoMessage()    {}

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

// ref: BobNetProto.DebugPaneDelItems
type DebugPaneDelItems struct {
	Items            []*DebugPaneDelItems_DebugPaneDelItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *DebugPaneDelItems) Reset()         { *m = DebugPaneDelItems{} }
func (m *DebugPaneDelItems) String() string { return proto.CompactTextString(m) }
func (*DebugPaneDelItems) ProtoMessage()    {}

func (m *DebugPaneDelItems) GetItems() []*DebugPaneDelItems_DebugPaneDelItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// ref: BobNetProto.DebugPaneDelItems/DebugPaneDelItem
type DebugPaneDelItems_DebugPaneDelItem struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugPaneDelItems_DebugPaneDelItem) Reset()         { *m = DebugPaneDelItems_DebugPaneDelItem{} }
func (m *DebugPaneDelItems_DebugPaneDelItem) String() string { return proto.CompactTextString(m) }
func (*DebugPaneDelItems_DebugPaneDelItem) ProtoMessage()    {}

func (m *DebugPaneDelItems_DebugPaneDelItem) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

// ref: BobNetProto.DebugPaneNewItems
type DebugPaneNewItems struct {
	Items            []*DebugPaneNewItems_DebugPaneNewItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *DebugPaneNewItems) Reset()         { *m = DebugPaneNewItems{} }
func (m *DebugPaneNewItems) String() string { return proto.CompactTextString(m) }
func (*DebugPaneNewItems) ProtoMessage()    {}

func (m *DebugPaneNewItems) GetItems() []*DebugPaneNewItems_DebugPaneNewItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// ref: BobNetProto.DebugPaneNewItems/DebugPaneNewItem
type DebugPaneNewItems_DebugPaneNewItem struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DebugPaneNewItems_DebugPaneNewItem) Reset()         { *m = DebugPaneNewItems_DebugPaneNewItem{} }
func (m *DebugPaneNewItems_DebugPaneNewItem) String() string { return proto.CompactTextString(m) }
func (*DebugPaneNewItems_DebugPaneNewItem) ProtoMessage()    {}

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

// ref: BobNetProto.PACKETTYPES
type PACKETTYPES struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PACKETTYPES) Reset()         { *m = PACKETTYPES{} }
func (m *PACKETTYPES) String() string { return proto.CompactTextString(m) }
func (*PACKETTYPES) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("bobnetproto.Deadend_PacketID", Deadend_PacketID_name, Deadend_PacketID_value)
	proto.RegisterEnum("bobnetproto.DeadendUtil_PacketID", DeadendUtil_PacketID_name, DeadendUtil_PacketID_value)
	proto.RegisterEnum("bobnetproto.DebugConsoleCmdList_PacketID", DebugConsoleCmdList_PacketID_name, DebugConsoleCmdList_PacketID_value)
	proto.RegisterEnum("bobnetproto.DebugConsoleCommand_PacketID", DebugConsoleCommand_PacketID_name, DebugConsoleCommand_PacketID_value)
	proto.RegisterEnum("bobnetproto.DebugConsoleGetCmdList_PacketID", DebugConsoleGetCmdList_PacketID_name, DebugConsoleGetCmdList_PacketID_value)
	proto.RegisterEnum("bobnetproto.DebugConsoleGetZones_PacketID", DebugConsoleGetZones_PacketID_name, DebugConsoleGetZones_PacketID_value)
	proto.RegisterEnum("bobnetproto.DebugConsoleResponse_PacketID", DebugConsoleResponse_PacketID_name, DebugConsoleResponse_PacketID_value)
	proto.RegisterEnum("bobnetproto.DebugConsoleResponse_ResponseType", DebugConsoleResponse_ResponseType_name, DebugConsoleResponse_ResponseType_value)
	proto.RegisterEnum("bobnetproto.DebugConsoleUpdateFromPane_PacketID", DebugConsoleUpdateFromPane_PacketID_name, DebugConsoleUpdateFromPane_PacketID_value)
	proto.RegisterEnum("bobnetproto.DebugConsoleZones_PacketID", DebugConsoleZones_PacketID_name, DebugConsoleZones_PacketID_value)
	proto.RegisterEnum("bobnetproto.DebugPaneDelItems_PacketID", DebugPaneDelItems_PacketID_name, DebugPaneDelItems_PacketID_value)
	proto.RegisterEnum("bobnetproto.DebugPaneNewItems_PacketID", DebugPaneNewItems_PacketID_name, DebugPaneNewItems_PacketID_value)
	proto.RegisterEnum("bobnetproto.PACKETTYPES_BobNetCount", PACKETTYPES_BobNetCount_name, PACKETTYPES_BobNetCount_value)
}