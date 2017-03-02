// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/presence_types/presence_types.proto
// DO NOT EDIT!

/*
Package presence_types is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/presence_types/presence_types.proto

It has these top-level messages:
	RichPresence
	FieldKey
	Field
	FieldOperation
	ChannelState
*/
package presence_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import attribute "github.com/HearthSim/hs-proto-go/bnet/attribute"
import channel_types "github.com/HearthSim/hs-proto-go/bnet/channel_types"
import entity "github.com/HearthSim/hs-proto-go/bnet/entity"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FieldOperation_OperationType int32

const (
	FieldOperation_SET   FieldOperation_OperationType = 0
	FieldOperation_CLEAR FieldOperation_OperationType = 1
)

var FieldOperation_OperationType_name = map[int32]string{
	0: "SET",
	1: "CLEAR",
}
var FieldOperation_OperationType_value = map[string]int32{
	"SET":   0,
	"CLEAR": 1,
}

func (x FieldOperation_OperationType) Enum() *FieldOperation_OperationType {
	p := new(FieldOperation_OperationType)
	*p = x
	return p
}
func (x FieldOperation_OperationType) String() string {
	return proto.EnumName(FieldOperation_OperationType_name, int32(x))
}
func (x *FieldOperation_OperationType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FieldOperation_OperationType_value, data, "FieldOperation_OperationType")
	if err != nil {
		return err
	}
	*x = FieldOperation_OperationType(value)
	return nil
}
func (FieldOperation_OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type RichPresence struct {
	ProgramId        *uint32 `protobuf:"fixed32,1,req,name=program_id,json=programId" json:"program_id,omitempty"`
	StreamId         *uint32 `protobuf:"fixed32,2,req,name=stream_id,json=streamId" json:"stream_id,omitempty"`
	Index            *uint32 `protobuf:"varint,3,req,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RichPresence) Reset()                    { *m = RichPresence{} }
func (m *RichPresence) String() string            { return proto.CompactTextString(m) }
func (*RichPresence) ProtoMessage()               {}
func (*RichPresence) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RichPresence) GetProgramId() uint32 {
	if m != nil && m.ProgramId != nil {
		return *m.ProgramId
	}
	return 0
}

func (m *RichPresence) GetStreamId() uint32 {
	if m != nil && m.StreamId != nil {
		return *m.StreamId
	}
	return 0
}

func (m *RichPresence) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

type FieldKey struct {
	Program          *uint32 `protobuf:"varint,1,req,name=program" json:"program,omitempty"`
	Group            *uint32 `protobuf:"varint,2,req,name=group" json:"group,omitempty"`
	Field            *uint32 `protobuf:"varint,3,req,name=field" json:"field,omitempty"`
	Index            *uint64 `protobuf:"varint,4,opt,name=index,def=0" json:"index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FieldKey) Reset()                    { *m = FieldKey{} }
func (m *FieldKey) String() string            { return proto.CompactTextString(m) }
func (*FieldKey) ProtoMessage()               {}
func (*FieldKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_FieldKey_Index uint64 = 0

func (m *FieldKey) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *FieldKey) GetGroup() uint32 {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return 0
}

func (m *FieldKey) GetField() uint32 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *FieldKey) GetIndex() uint64 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return Default_FieldKey_Index
}

type Field struct {
	Key              *FieldKey          `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            *attribute.Variant `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Field) Reset()                    { *m = Field{} }
func (m *Field) String() string            { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()               {}
func (*Field) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Field) GetKey() *FieldKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Field) GetValue() *attribute.Variant {
	if m != nil {
		return m.Value
	}
	return nil
}

type FieldOperation struct {
	Field            *Field                        `protobuf:"bytes,1,req,name=field" json:"field,omitempty"`
	Operation        *FieldOperation_OperationType `protobuf:"varint,2,opt,name=operation,enum=presence_types.FieldOperation_OperationType,def=0" json:"operation,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *FieldOperation) Reset()                    { *m = FieldOperation{} }
func (m *FieldOperation) String() string            { return proto.CompactTextString(m) }
func (*FieldOperation) ProtoMessage()               {}
func (*FieldOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_FieldOperation_Operation FieldOperation_OperationType = FieldOperation_SET

func (m *FieldOperation) GetField() *Field {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *FieldOperation) GetOperation() FieldOperation_OperationType {
	if m != nil && m.Operation != nil {
		return *m.Operation
	}
	return Default_FieldOperation_Operation
}

type ChannelState struct {
	EntityId         *entity.EntityId  `protobuf:"bytes,1,opt,name=entity_id,json=entityId" json:"entity_id,omitempty"`
	FieldOperation   []*FieldOperation `protobuf:"bytes,2,rep,name=field_operation,json=fieldOperation" json:"field_operation,omitempty"`
	Healing          *bool             `protobuf:"varint,3,opt,name=healing,def=0" json:"healing,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *ChannelState) Reset()                    { *m = ChannelState{} }
func (m *ChannelState) String() string            { return proto.CompactTextString(m) }
func (*ChannelState) ProtoMessage()               {}
func (*ChannelState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

const Default_ChannelState_Healing bool = false

func (m *ChannelState) GetEntityId() *entity.EntityId {
	if m != nil {
		return m.EntityId
	}
	return nil
}

func (m *ChannelState) GetFieldOperation() []*FieldOperation {
	if m != nil {
		return m.FieldOperation
	}
	return nil
}

func (m *ChannelState) GetHealing() bool {
	if m != nil && m.Healing != nil {
		return *m.Healing
	}
	return Default_ChannelState_Healing
}

var E_ChannelState_Presence = &proto.ExtensionDesc{
	ExtendedType:  (*channel_types.ChannelState)(nil),
	ExtensionType: (*ChannelState)(nil),
	Field:         101,
	Name:          "presence_types.ChannelState.presence",
	Tag:           "bytes,101,opt,name=presence",
	Filename:      "github.com/HearthSim/hs-proto-go/bnet/presence_types/presence_types.proto",
}

func init() {
	proto.RegisterType((*RichPresence)(nil), "presence_types.RichPresence")
	proto.RegisterType((*FieldKey)(nil), "presence_types.FieldKey")
	proto.RegisterType((*Field)(nil), "presence_types.Field")
	proto.RegisterType((*FieldOperation)(nil), "presence_types.FieldOperation")
	proto.RegisterType((*ChannelState)(nil), "presence_types.ChannelState")
	proto.RegisterEnum("presence_types.FieldOperation_OperationType", FieldOperation_OperationType_name, FieldOperation_OperationType_value)
	proto.RegisterExtension(E_ChannelState_Presence)
}

func init() {
	proto.RegisterFile("github.com/HearthSim/hs-proto-go/bnet/presence_types/presence_types.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x8f, 0xd2, 0x40,
	0x10, 0xc7, 0x2d, 0x1c, 0x01, 0x86, 0x03, 0xc9, 0x46, 0x63, 0x73, 0xe7, 0x0f, 0x82, 0x2f, 0x8d,
	0x4a, 0x31, 0x7d, 0x93, 0x17, 0x63, 0x2e, 0x78, 0x12, 0x4d, 0xce, 0x2c, 0xa7, 0x6f, 0x06, 0x17,
	0x3a, 0xb4, 0x1b, 0x4b, 0xdb, 0x6c, 0x17, 0x63, 0xff, 0x2e, 0xff, 0x3f, 0x63, 0xba, 0xdb, 0x05,
	0x4a, 0x2e, 0x17, 0x9e, 0xba, 0x33, 0x3b, 0xf3, 0xf9, 0x7e, 0xb7, 0x33, 0x30, 0x0b, 0xb8, 0x0c,
	0xb7, 0x4b, 0x77, 0x95, 0x6c, 0xc6, 0x9f, 0x90, 0x09, 0x19, 0xce, 0xf9, 0x66, 0x1c, 0x66, 0xa3,
	0x54, 0x24, 0x32, 0x19, 0x05, 0xc9, 0x78, 0x19, 0xa3, 0x1c, 0xa7, 0x02, 0x33, 0x8c, 0x57, 0xb8,
	0x90, 0x79, 0x8a, 0xd9, 0x51, 0xe8, 0xaa, 0x5a, 0xd2, 0xab, 0x66, 0x2f, 0xde, 0x9f, 0x86, 0x66,
	0x52, 0x0a, 0xbe, 0xdc, 0x4a, 0xdc, 0x9f, 0x34, 0xf0, 0xe2, 0xfa, 0x34, 0xc0, 0x2a, 0x64, 0x71,
	0x8c, 0x51, 0x69, 0xad, 0x12, 0x95, 0xa0, 0x77, 0xa7, 0x81, 0x30, 0x96, 0x5c, 0xe6, 0xe5, 0x47,
	0xb7, 0x0e, 0x7f, 0xc2, 0x39, 0xe5, 0xab, 0xf0, 0x6b, 0xf9, 0x34, 0xf2, 0x0c, 0x20, 0x15, 0x49,
	0x20, 0xd8, 0x66, 0xc1, 0x7d, 0xdb, 0x1a, 0xd4, 0x9c, 0x26, 0x6d, 0x97, 0x99, 0x99, 0x4f, 0x2e,
	0xa1, 0x9d, 0x49, 0x81, 0xfa, 0xb6, 0xa6, 0x6e, 0x5b, 0x3a, 0x31, 0xf3, 0xc9, 0x23, 0x68, 0xf0,
	0xd8, 0xc7, 0x3f, 0x76, 0x7d, 0x50, 0x73, 0xba, 0x54, 0x07, 0x43, 0x0e, 0xad, 0x8f, 0x1c, 0x23,
	0xff, 0x33, 0xe6, 0xc4, 0x86, 0x66, 0xc9, 0x52, 0xe8, 0x2e, 0x35, 0x61, 0xd1, 0x1b, 0x88, 0x64,
	0x9b, 0x2a, 0x68, 0x97, 0xea, 0xa0, 0xc8, 0xae, 0x8b, 0x5e, 0x43, 0x54, 0x01, 0x79, 0x62, 0x74,
	0xce, 0x06, 0x96, 0x73, 0x36, 0xb1, 0xde, 0x1a, 0xa9, 0x1f, 0xd0, 0x50, 0x52, 0xe4, 0x15, 0xd4,
	0x7f, 0x61, 0xae, 0x34, 0x3a, 0x9e, 0xed, 0x1e, 0x8d, 0xd3, 0xd8, 0xa1, 0x45, 0x11, 0x71, 0xa0,
	0xf1, 0x9b, 0x45, 0x5b, 0x54, 0xca, 0x1d, 0x8f, 0xb8, 0xfb, 0x31, 0x7d, 0x67, 0x82, 0xb3, 0x58,
	0x52, 0x5d, 0x30, 0xfc, 0x6b, 0x41, 0x4f, 0xf5, 0xde, 0xa4, 0x28, 0x98, 0xe4, 0x49, 0x4c, 0x5e,
	0x1b, 0x83, 0x5a, 0xea, 0xf1, 0x9d, 0x52, 0xc6, 0xf7, 0x0d, 0xb4, 0x13, 0xd3, 0x69, 0xd7, 0x06,
	0x96, 0xd3, 0xf3, 0xde, 0xdc, 0xd9, 0xb0, 0xe3, 0xbb, 0xbb, 0xd3, 0x6d, 0x9e, 0xe2, 0xa4, 0x3e,
	0x9f, 0xde, 0xd2, 0x3d, 0x63, 0xf8, 0x12, 0xba, 0x95, 0x02, 0xd2, 0x84, 0xa2, 0xa4, 0xff, 0x80,
	0xb4, 0xa1, 0x71, 0xf5, 0x65, 0xfa, 0x81, 0xf6, 0xad, 0xe1, 0x3f, 0x0b, 0xce, 0xaf, 0xf4, 0xd2,
	0xcc, 0x25, 0x93, 0x48, 0x46, 0xd0, 0xd6, 0x2b, 0xa0, 0x27, 0x6c, 0x39, 0x1d, 0xaf, 0xef, 0x96,
	0x4b, 0x31, 0x55, 0x9f, 0x99, 0x4f, 0x5b, 0x58, 0x9e, 0xc8, 0x35, 0x3c, 0x54, 0xf6, 0x17, 0x87,
	0xde, 0xeb, 0x4e, 0xc7, 0x7b, 0x7e, 0xbf, 0x77, 0xda, 0x5b, 0x57, 0xff, 0xd5, 0x0b, 0x68, 0x86,
	0xc8, 0x22, 0x1e, 0x07, 0x76, 0x7d, 0x60, 0x39, 0xad, 0x49, 0x63, 0xcd, 0xa2, 0x0c, 0xa9, 0xc9,
	0x7a, 0xdf, 0xa0, 0x65, 0x88, 0xe4, 0xd2, 0xad, 0x2e, 0xfa, 0xe1, 0x0b, 0x6c, 0x54, 0xa6, 0x9f,
	0x1e, 0xeb, 0x1f, 0xd6, 0xd0, 0x1d, 0xea, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xab, 0xfa,
	0x57, 0x03, 0x04, 0x00, 0x00,
}
