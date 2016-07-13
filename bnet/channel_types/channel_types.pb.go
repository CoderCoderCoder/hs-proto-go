// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/channel_types/channel_types.proto
// DO NOT EDIT!

/*
Package channel_types is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/channel_types/channel_types.proto

It has these top-level messages:
	Message
	FindChannelOptions
	ChannelDescription
	ChannelInfo
	ChannelState
	MemberState
	Member
*/
package channel_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import attribute "github.com/HearthSim/hs-proto-go/bnet/attribute"
import entity "github.com/HearthSim/hs-proto-go/bnet/entity"
import invitation_types "github.com/HearthSim/hs-proto-go/bnet/invitation_types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type ChannelState_PrivacyLevel int32

const (
	ChannelState_PRIVACY_LEVEL_OPEN                       ChannelState_PrivacyLevel = 1
	ChannelState_PRIVACY_LEVEL_OPEN_INVITATION_AND_FRIEND ChannelState_PrivacyLevel = 2
	ChannelState_PRIVACY_LEVEL_OPEN_INVITATION            ChannelState_PrivacyLevel = 3
	ChannelState_PRIVACY_LEVEL_CLOSED                     ChannelState_PrivacyLevel = 4
)

var ChannelState_PrivacyLevel_name = map[int32]string{
	1: "PRIVACY_LEVEL_OPEN",
	2: "PRIVACY_LEVEL_OPEN_INVITATION_AND_FRIEND",
	3: "PRIVACY_LEVEL_OPEN_INVITATION",
	4: "PRIVACY_LEVEL_CLOSED",
}
var ChannelState_PrivacyLevel_value = map[string]int32{
	"PRIVACY_LEVEL_OPEN":                       1,
	"PRIVACY_LEVEL_OPEN_INVITATION_AND_FRIEND": 2,
	"PRIVACY_LEVEL_OPEN_INVITATION":            3,
	"PRIVACY_LEVEL_CLOSED":                     4,
}

func (x ChannelState_PrivacyLevel) Enum() *ChannelState_PrivacyLevel {
	p := new(ChannelState_PrivacyLevel)
	*p = x
	return p
}
func (x ChannelState_PrivacyLevel) String() string {
	return proto.EnumName(ChannelState_PrivacyLevel_name, int32(x))
}
func (x *ChannelState_PrivacyLevel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ChannelState_PrivacyLevel_value, data, "ChannelState_PrivacyLevel")
	if err != nil {
		return err
	}
	*x = ChannelState_PrivacyLevel(value)
	return nil
}
func (ChannelState_PrivacyLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type Message struct {
	Attribute        []*attribute.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	Role             *uint32                `protobuf:"varint,2,opt,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *Message) GetRole() uint32 {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return 0
}

type FindChannelOptions struct {
	StartIndex       *uint32                    `protobuf:"varint,1,opt,name=start_index,def=0" json:"start_index,omitempty"`
	MaxResults       *uint32                    `protobuf:"varint,2,opt,name=max_results,def=16" json:"max_results,omitempty"`
	Name             *string                    `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Program          *uint32                    `protobuf:"fixed32,4,opt,name=program" json:"program,omitempty"`
	Locale           *uint32                    `protobuf:"fixed32,5,opt,name=locale" json:"locale,omitempty"`
	CapacityFull     *uint32                    `protobuf:"varint,6,opt,name=capacity_full" json:"capacity_full,omitempty"`
	AttributeFilter  *attribute.AttributeFilter `protobuf:"bytes,7,req,name=attribute_filter" json:"attribute_filter,omitempty"`
	ChannelType      *string                    `protobuf:"bytes,8,opt,name=channel_type" json:"channel_type,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *FindChannelOptions) Reset()                    { *m = FindChannelOptions{} }
func (m *FindChannelOptions) String() string            { return proto.CompactTextString(m) }
func (*FindChannelOptions) ProtoMessage()               {}
func (*FindChannelOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

const Default_FindChannelOptions_StartIndex uint32 = 0
const Default_FindChannelOptions_MaxResults uint32 = 16

func (m *FindChannelOptions) GetStartIndex() uint32 {
	if m != nil && m.StartIndex != nil {
		return *m.StartIndex
	}
	return Default_FindChannelOptions_StartIndex
}

func (m *FindChannelOptions) GetMaxResults() uint32 {
	if m != nil && m.MaxResults != nil {
		return *m.MaxResults
	}
	return Default_FindChannelOptions_MaxResults
}

func (m *FindChannelOptions) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FindChannelOptions) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *FindChannelOptions) GetLocale() uint32 {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return 0
}

func (m *FindChannelOptions) GetCapacityFull() uint32 {
	if m != nil && m.CapacityFull != nil {
		return *m.CapacityFull
	}
	return 0
}

func (m *FindChannelOptions) GetAttributeFilter() *attribute.AttributeFilter {
	if m != nil {
		return m.AttributeFilter
	}
	return nil
}

func (m *FindChannelOptions) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return ""
}

type ChannelDescription struct {
	ChannelId        *entity.EntityId `protobuf:"bytes,1,req,name=channel_id" json:"channel_id,omitempty"`
	CurrentMembers   *uint32          `protobuf:"varint,2,opt,name=current_members" json:"current_members,omitempty"`
	State            *ChannelState    `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ChannelDescription) Reset()                    { *m = ChannelDescription{} }
func (m *ChannelDescription) String() string            { return proto.CompactTextString(m) }
func (*ChannelDescription) ProtoMessage()               {}
func (*ChannelDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ChannelDescription) GetChannelId() *entity.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *ChannelDescription) GetCurrentMembers() uint32 {
	if m != nil && m.CurrentMembers != nil {
		return *m.CurrentMembers
	}
	return 0
}

func (m *ChannelDescription) GetState() *ChannelState {
	if m != nil {
		return m.State
	}
	return nil
}

type ChannelInfo struct {
	Description      *ChannelDescription `protobuf:"bytes,1,req,name=description" json:"description,omitempty"`
	Member           []*Member           `protobuf:"bytes,2,rep,name=member" json:"member,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ChannelInfo) Reset()                    { *m = ChannelInfo{} }
func (m *ChannelInfo) String() string            { return proto.CompactTextString(m) }
func (*ChannelInfo) ProtoMessage()               {}
func (*ChannelInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ChannelInfo) GetDescription() *ChannelDescription {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *ChannelInfo) GetMember() []*Member {
	if m != nil {
		return m.Member
	}
	return nil
}

type ChannelState struct {
	MaxMembers          *uint32                        `protobuf:"varint,1,opt,name=max_members" json:"max_members,omitempty"`
	MinMembers          *uint32                        `protobuf:"varint,2,opt,name=min_members" json:"min_members,omitempty"`
	Attribute           []*attribute.Attribute         `protobuf:"bytes,3,rep,name=attribute" json:"attribute,omitempty"`
	Invitation          []*invitation_types.Invitation `protobuf:"bytes,4,rep,name=invitation" json:"invitation,omitempty"`
	MaxInvitations      *uint32                        `protobuf:"varint,5,opt,name=max_invitations" json:"max_invitations,omitempty"`
	Reason              *uint32                        `protobuf:"varint,6,opt,name=reason" json:"reason,omitempty"`
	PrivacyLevel        *ChannelState_PrivacyLevel     `protobuf:"varint,7,opt,name=privacy_level,enum=channel_types.ChannelState_PrivacyLevel,def=1" json:"privacy_level,omitempty"`
	Name                *string                        `protobuf:"bytes,8,opt,name=name" json:"name,omitempty"`
	DelegateName        *string                        `protobuf:"bytes,9,opt,name=delegate_name" json:"delegate_name,omitempty"`
	ChannelType         *string                        `protobuf:"bytes,10,opt,name=channel_type,def=default" json:"channel_type,omitempty"`
	Program             *uint32                        `protobuf:"fixed32,11,opt,name=program,def=0" json:"program,omitempty"`
	AllowOfflineMembers *bool                          `protobuf:"varint,12,opt,name=allow_offline_members,def=0" json:"allow_offline_members,omitempty"`
	SubscribeToPresence *bool                          `protobuf:"varint,13,opt,name=subscribe_to_presence,def=1" json:"subscribe_to_presence,omitempty"`
	XXX_extensions      map[int32]proto.Extension      `json:"-"`
	XXX_unrecognized    []byte                         `json:"-"`
}

func (m *ChannelState) Reset()                    { *m = ChannelState{} }
func (m *ChannelState) String() string            { return proto.CompactTextString(m) }
func (*ChannelState) ProtoMessage()               {}
func (*ChannelState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

var extRange_ChannelState = []proto.ExtensionRange{
	{100, 10000},
}

func (*ChannelState) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ChannelState
}
func (m *ChannelState) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_ChannelState_PrivacyLevel ChannelState_PrivacyLevel = ChannelState_PRIVACY_LEVEL_OPEN
const Default_ChannelState_ChannelType string = "default"
const Default_ChannelState_Program uint32 = 0
const Default_ChannelState_AllowOfflineMembers bool = false
const Default_ChannelState_SubscribeToPresence bool = true

func (m *ChannelState) GetMaxMembers() uint32 {
	if m != nil && m.MaxMembers != nil {
		return *m.MaxMembers
	}
	return 0
}

func (m *ChannelState) GetMinMembers() uint32 {
	if m != nil && m.MinMembers != nil {
		return *m.MinMembers
	}
	return 0
}

func (m *ChannelState) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *ChannelState) GetInvitation() []*invitation_types.Invitation {
	if m != nil {
		return m.Invitation
	}
	return nil
}

func (m *ChannelState) GetMaxInvitations() uint32 {
	if m != nil && m.MaxInvitations != nil {
		return *m.MaxInvitations
	}
	return 0
}

func (m *ChannelState) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

func (m *ChannelState) GetPrivacyLevel() ChannelState_PrivacyLevel {
	if m != nil && m.PrivacyLevel != nil {
		return *m.PrivacyLevel
	}
	return Default_ChannelState_PrivacyLevel
}

func (m *ChannelState) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChannelState) GetDelegateName() string {
	if m != nil && m.DelegateName != nil {
		return *m.DelegateName
	}
	return ""
}

func (m *ChannelState) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return Default_ChannelState_ChannelType
}

func (m *ChannelState) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return Default_ChannelState_Program
}

func (m *ChannelState) GetAllowOfflineMembers() bool {
	if m != nil && m.AllowOfflineMembers != nil {
		return *m.AllowOfflineMembers
	}
	return Default_ChannelState_AllowOfflineMembers
}

func (m *ChannelState) GetSubscribeToPresence() bool {
	if m != nil && m.SubscribeToPresence != nil {
		return *m.SubscribeToPresence
	}
	return Default_ChannelState_SubscribeToPresence
}

type MemberState struct {
	Attribute        []*attribute.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	Role             []uint32               `protobuf:"varint,2,rep,packed,name=role" json:"role,omitempty"`
	Privileges       *uint64                `protobuf:"varint,3,opt,name=privileges,def=0" json:"privileges,omitempty"`
	Info             *entity.AccountInfo    `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Hidden           *bool                  `protobuf:"varint,5,opt,name=hidden,def=0" json:"hidden,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *MemberState) Reset()                    { *m = MemberState{} }
func (m *MemberState) String() string            { return proto.CompactTextString(m) }
func (*MemberState) ProtoMessage()               {}
func (*MemberState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

const Default_MemberState_Privileges uint64 = 0
const Default_MemberState_Hidden bool = false

func (m *MemberState) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *MemberState) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *MemberState) GetPrivileges() uint64 {
	if m != nil && m.Privileges != nil {
		return *m.Privileges
	}
	return Default_MemberState_Privileges
}

func (m *MemberState) GetInfo() *entity.AccountInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *MemberState) GetHidden() bool {
	if m != nil && m.Hidden != nil {
		return *m.Hidden
	}
	return Default_MemberState_Hidden
}

type Member struct {
	Identity         *entity.Identity `protobuf:"bytes,1,req,name=identity" json:"identity,omitempty"`
	State            *MemberState     `protobuf:"bytes,2,req,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Member) Reset()                    { *m = Member{} }
func (m *Member) String() string            { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()               {}
func (*Member) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Member) GetIdentity() *entity.Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Member) GetState() *MemberState {
	if m != nil {
		return m.State
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "channel_types.Message")
	proto.RegisterType((*FindChannelOptions)(nil), "channel_types.FindChannelOptions")
	proto.RegisterType((*ChannelDescription)(nil), "channel_types.ChannelDescription")
	proto.RegisterType((*ChannelInfo)(nil), "channel_types.ChannelInfo")
	proto.RegisterType((*ChannelState)(nil), "channel_types.ChannelState")
	proto.RegisterType((*MemberState)(nil), "channel_types.MemberState")
	proto.RegisterType((*Member)(nil), "channel_types.Member")
	proto.RegisterEnum("channel_types.ChannelState_PrivacyLevel", ChannelState_PrivacyLevel_name, ChannelState_PrivacyLevel_value)
}

var fileDescriptor0 = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xea, 0x46,
	0x10, 0x95, 0xf9, 0xce, 0x18, 0xee, 0xb5, 0xf6, 0x86, 0xd6, 0x4a, 0x1b, 0x29, 0x71, 0x53, 0x95,
	0x46, 0x0d, 0xa4, 0xa8, 0x8a, 0x54, 0x5e, 0x5a, 0x1a, 0x48, 0x6b, 0x89, 0x40, 0x14, 0xa2, 0x54,
	0x7d, 0x5a, 0x2d, 0xf6, 0x02, 0x2b, 0x19, 0x1b, 0xd9, 0xeb, 0x34, 0x79, 0xea, 0x6f, 0xc8, 0x63,
	0x7f, 0x55, 0x1f, 0xfb, 0x77, 0xee, 0xee, 0xda, 0x7c, 0x06, 0x45, 0x3c, 0xd9, 0x1e, 0x9f, 0x99,
	0x3d, 0x67, 0xe6, 0xcc, 0xc2, 0xef, 0x13, 0xc6, 0xa7, 0xf1, 0xa8, 0xee, 0x04, 0xb3, 0xc6, 0x1f,
	0x94, 0x84, 0x7c, 0x3a, 0x64, 0xb3, 0xc6, 0x34, 0xba, 0x98, 0x87, 0x01, 0x0f, 0x2e, 0x26, 0x41,
	0x63, 0xe4, 0x53, 0xde, 0x70, 0xa6, 0xc4, 0xf7, 0xa9, 0x87, 0xf9, 0xcb, 0x9c, 0x46, 0x9b, 0x5f,
	0x75, 0x85, 0x44, 0x95, 0x8d, 0xe0, 0xd1, 0x2f, 0xfb, 0xd5, 0x25, 0x9c, 0x87, 0x6c, 0x14, 0x73,
	0xba, 0x7a, 0x4b, 0xea, 0x1d, 0xfd, 0xbc, 0x5f, 0x01, 0xea, 0x73, 0xc6, 0x5f, 0xd2, 0x47, 0x9a,
	0x7a, 0xbb, 0x5f, 0x2a, 0xf3, 0x9f, 0x18, 0x27, 0x9c, 0x05, 0x7e, 0x2a, 0x6b, 0x3b, 0x90, 0x94,
	0xb3, 0x7e, 0x85, 0xe2, 0x2d, 0x8d, 0x22, 0x32, 0xa1, 0xe8, 0x3b, 0x38, 0x58, 0xf2, 0x34, 0xb5,
	0x93, 0x6c, 0x4d, 0x6f, 0x1e, 0xd6, 0x57, 0xcc, 0xdb, 0x8b, 0x37, 0x54, 0x86, 0x5c, 0x18, 0x78,
	0xd4, 0xcc, 0x9c, 0x68, 0xb5, 0x8a, 0xf5, 0xbf, 0x06, 0xe8, 0x86, 0xf9, 0xee, 0x75, 0xd2, 0xa2,
	0xc1, 0x5c, 0x9e, 0x11, 0xa1, 0x2f, 0x40, 0x8f, 0xb8, 0xa0, 0x87, 0xc5, 0x2f, 0xfa, 0x2c, 0xea,
	0x09, 0x6c, 0x4b, 0xbb, 0x44, 0x5f, 0x82, 0x3e, 0x23, 0xcf, 0x38, 0xa4, 0x51, 0xec, 0xf1, 0x28,
	0xa9, 0xd1, 0xca, 0xfc, 0x78, 0x25, 0xab, 0xfa, 0x64, 0x46, 0xcd, 0xac, 0x88, 0x1c, 0xa0, 0x8f,
	0x50, 0x14, 0x04, 0x27, 0x21, 0x99, 0x99, 0x39, 0x11, 0x28, 0xa2, 0x0f, 0x50, 0xf0, 0x02, 0x87,
	0x88, 0x63, 0xf3, 0xea, 0xbb, 0x0a, 0x15, 0x87, 0xcc, 0x89, 0x23, 0x3a, 0x83, 0xc7, 0xb1, 0xe7,
	0x99, 0x05, 0x59, 0x09, 0xfd, 0x04, 0xc6, 0x92, 0x32, 0x1e, 0x33, 0x8f, 0xd3, 0xd0, 0x2c, 0x9e,
	0x64, 0x84, 0x96, 0xa3, 0x5d, 0x5a, 0x6e, 0x14, 0x02, 0x1d, 0x42, 0x79, 0x7d, 0xc2, 0x66, 0x49,
	0x72, 0xb0, 0xfe, 0x01, 0x94, 0x8a, 0xea, 0xd0, 0xc8, 0x09, 0x99, 0x52, 0x86, 0xce, 0x00, 0x16,
	0x58, 0xe6, 0x0a, 0x5d, 0xb2, 0xb6, 0x51, 0x4f, 0x67, 0xd4, 0x55, 0x0f, 0xdb, 0x15, 0x32, 0x3f,
	0x3a, 0x71, 0x18, 0x8a, 0x30, 0x9e, 0xd1, 0xd9, 0x88, 0x86, 0xa9, 0x54, 0x74, 0x0e, 0x79, 0xd1,
	0x17, 0x9e, 0xe8, 0xd4, 0x9b, 0x5f, 0xd5, 0x37, 0xfd, 0x96, 0x1e, 0x38, 0x94, 0x10, 0xcb, 0x03,
	0x3d, 0xfd, 0xb6, 0xfd, 0x71, 0x80, 0xae, 0x40, 0x77, 0x57, 0x44, 0xd2, 0xa3, 0x4f, 0x77, 0x17,
	0x58, 0x67, 0xfc, 0x2d, 0x14, 0x12, 0x0e, 0x82, 0x82, 0x9c, 0x6a, 0x75, 0x2b, 0xe5, 0x56, 0xfd,
	0xb4, 0xfe, 0xcb, 0x41, 0x79, 0xfd, 0x78, 0xf4, 0x29, 0x19, 0xd5, 0x82, 0xbf, 0x1a, 0xa1, 0x0a,
	0x32, 0x7f, 0x4b, 0xd4, 0x86, 0x75, 0xb2, 0xef, 0x58, 0xe7, 0x12, 0x60, 0x65, 0x44, 0x31, 0x59,
	0x89, 0xfc, 0xba, 0xfe, 0xc6, 0x9b, 0xf6, 0x32, 0x20, 0x1b, 0x29, 0x49, 0xac, 0x20, 0x91, 0x32,
	0x40, 0x45, 0x1a, 0x22, 0xa4, 0x24, 0x12, 0x65, 0x92, 0xc9, 0x0f, 0xa1, 0x32, 0x0f, 0xd9, 0x13,
	0x71, 0x5e, 0xb0, 0x47, 0x9f, 0xa8, 0x27, 0xc6, 0xae, 0xd5, 0x3e, 0x34, 0x6b, 0xef, 0x34, 0xb8,
	0x7e, 0x97, 0x24, 0xf4, 0x24, 0xbe, 0x85, 0xee, 0xee, 0xed, 0xc7, 0xf6, 0xf5, 0x5f, 0xb8, 0xd7,
	0x7d, 0xec, 0xf6, 0xf0, 0xe0, 0xae, 0xdb, 0x5f, 0x9a, 0x52, 0x19, 0x42, 0x7a, 0xce, 0xa5, 0x1e,
	0x9d, 0x88, 0x54, 0xac, 0xc2, 0x07, 0x2a, 0x7c, 0xbc, 0xe5, 0x1e, 0x90, 0xd1, 0x56, 0xd1, 0xa5,
	0x63, 0x22, 0x3c, 0x8e, 0xd0, 0xca, 0xca, 0xba, 0xb4, 0xae, 0xdc, 0x82, 0x33, 0xa8, 0x12, 0xcf,
	0x0b, 0xfe, 0xc6, 0xc1, 0x78, 0xec, 0x31, 0x9f, 0x2e, 0xfb, 0x59, 0x16, 0x88, 0x52, 0x2b, 0x3f,
	0x26, 0x5e, 0x44, 0xd1, 0x37, 0x50, 0x8d, 0xe2, 0x91, 0x1c, 0xe4, 0x88, 0x62, 0x1e, 0xe0, 0xb9,
	0xd8, 0x1a, 0xea, 0x3b, 0xd4, 0xac, 0x28, 0x54, 0x8e, 0x87, 0x31, 0xb5, 0x5e, 0x35, 0x28, 0xaf,
	0xeb, 0x10, 0x9b, 0xb7, 0x43, 0x89, 0xa1, 0xa1, 0x1f, 0xa0, 0xf6, 0x36, 0x8e, 0xed, 0xfe, 0xa3,
	0xfd, 0xd0, 0x7e, 0xb0, 0x07, 0x7d, 0xdc, 0xee, 0x77, 0xf0, 0xcd, 0xbd, 0xdd, 0xed, 0x77, 0x8c,
	0x0c, 0x3a, 0x85, 0xe3, 0x77, 0xd1, 0x46, 0x16, 0x99, 0x70, 0xb8, 0x09, 0xb9, 0xee, 0x0d, 0x86,
	0xdd, 0x8e, 0x91, 0x3b, 0xcf, 0x97, 0x5c, 0xe3, 0xb5, 0x6f, 0xfd, 0xab, 0x81, 0x9e, 0x98, 0x2b,
	0x31, 0xd4, 0xde, 0x37, 0x8c, 0xb1, 0xbc, 0x61, 0xb2, 0xb5, 0xca, 0x6f, 0x19, 0x41, 0xbe, 0x0a,
	0x20, 0xa7, 0xcb, 0x44, 0xf7, 0x69, 0xa4, 0x76, 0x27, 0x27, 0xfb, 0x78, 0x0a, 0x39, 0x26, 0x56,
	0x43, 0xdd, 0x11, 0x7a, 0xf3, 0xd3, 0x62, 0x0d, 0xdb, 0x8e, 0x13, 0xc4, 0x3e, 0x57, 0x5b, 0x53,
	0x85, 0xc2, 0x94, 0xb9, 0x2e, 0xf5, 0x95, 0x6f, 0x16, 0xbd, 0xb5, 0xfe, 0x84, 0x42, 0x42, 0x0d,
	0x59, 0x50, 0x62, 0x6e, 0x92, 0xb8, 0xbd, 0xce, 0x76, 0x1a, 0x47, 0xdf, 0x2f, 0xb6, 0x36, 0x93,
	0xde, 0x25, 0xbb, 0x36, 0x48, 0x89, 0xfc, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x27, 0xcf, 0x74, 0xfa,
	0x75, 0x06, 0x00, 0x00,
}
