// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/pegasus/spectator/spectator.proto
// DO NOT EDIT!

/*
Package spectator is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/pegasus/spectator/spectator.proto

It has these top-level messages:
	JoinInfo
	Invite
	PartyServerInfo
	PartyQuestInfo
*/
package spectator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import shared "github.com/HearthSim/hs-proto-go/pegasus/shared"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type JoinInfo struct {
	ServerIpAddress      *string            `protobuf:"bytes,1,opt,name=server_ip_address" json:"server_ip_address,omitempty"`
	ServerPort           *uint32            `protobuf:"varint,2,opt,name=server_port" json:"server_port,omitempty"`
	GameHandle           *int32             `protobuf:"varint,3,opt,name=game_handle" json:"game_handle,omitempty"`
	SecretKey            *string            `protobuf:"bytes,4,opt,name=secret_key" json:"secret_key,omitempty"`
	IsJoinable           *bool              `protobuf:"varint,5,opt,name=is_joinable" json:"is_joinable,omitempty"`
	CurrentNumSpectators *int32             `protobuf:"varint,6,opt,name=current_num_spectators" json:"current_num_spectators,omitempty"`
	MaxNumSpectators     *int32             `protobuf:"varint,7,opt,name=max_num_spectators" json:"max_num_spectators,omitempty"`
	GameType             *shared.GameType   `protobuf:"varint,8,opt,name=game_type,enum=shared.GameType,def=0" json:"game_type,omitempty"`
	FormatType           *shared.FormatType `protobuf:"varint,12,opt,name=format_type,enum=shared.FormatType,def=0" json:"format_type,omitempty"`
	MissionId            *int32             `protobuf:"varint,9,opt,name=mission_id" json:"mission_id,omitempty"`
	SpectatedPlayers     []*shared.BnetId   `protobuf:"bytes,10,rep,name=spectated_players" json:"spectated_players,omitempty"`
	PartyId              *shared.BnetId     `protobuf:"bytes,11,opt,name=party_id" json:"party_id,omitempty"`
	XXX_unrecognized     []byte             `json:"-"`
}

func (m *JoinInfo) Reset()                    { *m = JoinInfo{} }
func (m *JoinInfo) String() string            { return proto.CompactTextString(m) }
func (*JoinInfo) ProtoMessage()               {}
func (*JoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_JoinInfo_GameType shared.GameType = shared.GameType_GT_UNKNOWN
const Default_JoinInfo_FormatType shared.FormatType = shared.FormatType_FT_UNKNOWN

func (m *JoinInfo) GetServerIpAddress() string {
	if m != nil && m.ServerIpAddress != nil {
		return *m.ServerIpAddress
	}
	return ""
}

func (m *JoinInfo) GetServerPort() uint32 {
	if m != nil && m.ServerPort != nil {
		return *m.ServerPort
	}
	return 0
}

func (m *JoinInfo) GetGameHandle() int32 {
	if m != nil && m.GameHandle != nil {
		return *m.GameHandle
	}
	return 0
}

func (m *JoinInfo) GetSecretKey() string {
	if m != nil && m.SecretKey != nil {
		return *m.SecretKey
	}
	return ""
}

func (m *JoinInfo) GetIsJoinable() bool {
	if m != nil && m.IsJoinable != nil {
		return *m.IsJoinable
	}
	return false
}

func (m *JoinInfo) GetCurrentNumSpectators() int32 {
	if m != nil && m.CurrentNumSpectators != nil {
		return *m.CurrentNumSpectators
	}
	return 0
}

func (m *JoinInfo) GetMaxNumSpectators() int32 {
	if m != nil && m.MaxNumSpectators != nil {
		return *m.MaxNumSpectators
	}
	return 0
}

func (m *JoinInfo) GetGameType() shared.GameType {
	if m != nil && m.GameType != nil {
		return *m.GameType
	}
	return Default_JoinInfo_GameType
}

func (m *JoinInfo) GetFormatType() shared.FormatType {
	if m != nil && m.FormatType != nil {
		return *m.FormatType
	}
	return Default_JoinInfo_FormatType
}

func (m *JoinInfo) GetMissionId() int32 {
	if m != nil && m.MissionId != nil {
		return *m.MissionId
	}
	return 0
}

func (m *JoinInfo) GetSpectatedPlayers() []*shared.BnetId {
	if m != nil {
		return m.SpectatedPlayers
	}
	return nil
}

func (m *JoinInfo) GetPartyId() *shared.BnetId {
	if m != nil {
		return m.PartyId
	}
	return nil
}

type Invite struct {
	InviterGameAccountId *shared.BnetId `protobuf:"bytes,1,req,name=inviter_game_account_id" json:"inviter_game_account_id,omitempty"`
	JoinInfo             *JoinInfo      `protobuf:"bytes,2,req,name=join_info" json:"join_info,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *Invite) Reset()                    { *m = Invite{} }
func (m *Invite) String() string            { return proto.CompactTextString(m) }
func (*Invite) ProtoMessage()               {}
func (*Invite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Invite) GetInviterGameAccountId() *shared.BnetId {
	if m != nil {
		return m.InviterGameAccountId
	}
	return nil
}

func (m *Invite) GetJoinInfo() *JoinInfo {
	if m != nil {
		return m.JoinInfo
	}
	return nil
}

type PartyServerInfo struct {
	ServerIpAddress  *string            `protobuf:"bytes,1,opt,name=server_ip_address" json:"server_ip_address,omitempty"`
	ServerPort       *uint32            `protobuf:"varint,2,opt,name=server_port" json:"server_port,omitempty"`
	GameHandle       *int32             `protobuf:"varint,3,opt,name=game_handle" json:"game_handle,omitempty"`
	SecretKey        *string            `protobuf:"bytes,4,opt,name=secret_key" json:"secret_key,omitempty"`
	GameType         *shared.GameType   `protobuf:"varint,5,opt,name=game_type,enum=shared.GameType,def=0" json:"game_type,omitempty"`
	FormatType       *shared.FormatType `protobuf:"varint,7,opt,name=format_type,enum=shared.FormatType,def=0" json:"format_type,omitempty"`
	MissionId        *int32             `protobuf:"varint,6,opt,name=mission_id" json:"mission_id,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *PartyServerInfo) Reset()                    { *m = PartyServerInfo{} }
func (m *PartyServerInfo) String() string            { return proto.CompactTextString(m) }
func (*PartyServerInfo) ProtoMessage()               {}
func (*PartyServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_PartyServerInfo_GameType shared.GameType = shared.GameType_GT_UNKNOWN
const Default_PartyServerInfo_FormatType shared.FormatType = shared.FormatType_FT_UNKNOWN

func (m *PartyServerInfo) GetServerIpAddress() string {
	if m != nil && m.ServerIpAddress != nil {
		return *m.ServerIpAddress
	}
	return ""
}

func (m *PartyServerInfo) GetServerPort() uint32 {
	if m != nil && m.ServerPort != nil {
		return *m.ServerPort
	}
	return 0
}

func (m *PartyServerInfo) GetGameHandle() int32 {
	if m != nil && m.GameHandle != nil {
		return *m.GameHandle
	}
	return 0
}

func (m *PartyServerInfo) GetSecretKey() string {
	if m != nil && m.SecretKey != nil {
		return *m.SecretKey
	}
	return ""
}

func (m *PartyServerInfo) GetGameType() shared.GameType {
	if m != nil && m.GameType != nil {
		return *m.GameType
	}
	return Default_PartyServerInfo_GameType
}

func (m *PartyServerInfo) GetFormatType() shared.FormatType {
	if m != nil && m.FormatType != nil {
		return *m.FormatType
	}
	return Default_PartyServerInfo_FormatType
}

func (m *PartyServerInfo) GetMissionId() int32 {
	if m != nil && m.MissionId != nil {
		return *m.MissionId
	}
	return 0
}

type PartyQuestInfo struct {
	QuestIds         []int32 `protobuf:"varint,1,rep,name=quest_ids" json:"quest_ids,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PartyQuestInfo) Reset()                    { *m = PartyQuestInfo{} }
func (m *PartyQuestInfo) String() string            { return proto.CompactTextString(m) }
func (*PartyQuestInfo) ProtoMessage()               {}
func (*PartyQuestInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PartyQuestInfo) GetQuestIds() []int32 {
	if m != nil {
		return m.QuestIds
	}
	return nil
}

func init() {
	proto.RegisterType((*JoinInfo)(nil), "spectator.JoinInfo")
	proto.RegisterType((*Invite)(nil), "spectator.Invite")
	proto.RegisterType((*PartyServerInfo)(nil), "spectator.PartyServerInfo")
	proto.RegisterType((*PartyQuestInfo)(nil), "spectator.PartyQuestInfo")
}

var fileDescriptor0 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x12, 0xd2, 0xc6, 0x13, 0x08, 0x74, 0x2b, 0x81, 0xe9, 0x01, 0x45, 0x41, 0x42, 0xe1,
	0x50, 0x5b, 0x2a, 0x37, 0xc4, 0xa9, 0x87, 0x96, 0x80, 0x14, 0x40, 0x2d, 0xe2, 0xb8, 0xda, 0xda,
	0x53, 0x7b, 0xa1, 0xde, 0x5d, 0x76, 0xd7, 0x15, 0xfe, 0xbf, 0xfc, 0x0b, 0x2e, 0x8c, 0xd7, 0x71,
	0x23, 0x3e, 0x0e, 0x80, 0xc4, 0xc9, 0xeb, 0x37, 0xf3, 0xde, 0xce, 0x9b, 0xb7, 0x70, 0x5c, 0x48,
	0x5f, 0xd6, 0x17, 0x49, 0xa6, 0xab, 0xf4, 0x25, 0x0a, 0xeb, 0xcb, 0x33, 0x59, 0xa5, 0xa5, 0x3b,
	0x34, 0x56, 0x7b, 0x7d, 0x58, 0xe8, 0xd4, 0x60, 0x21, 0x5c, 0xed, 0x52, 0x67, 0x30, 0xf3, 0xc2,
	0x6b, 0xbb, 0x3d, 0x25, 0xa1, 0x89, 0x45, 0x37, 0xc0, 0xc1, 0x8b, 0x3f, 0x97, 0x2b, 0x85, 0xc5,
	0x7c, 0xf3, 0xe9, 0x84, 0x16, 0xdf, 0x86, 0x30, 0x79, 0xa5, 0xa5, 0x5a, 0xa9, 0x4b, 0xcd, 0x1e,
	0xc2, 0x9e, 0x43, 0x7b, 0x8d, 0x96, 0x4b, 0xc3, 0x45, 0x9e, 0x5b, 0x74, 0x2e, 0x1e, 0xcc, 0x07,
	0xcb, 0x88, 0xed, 0xc3, 0x74, 0x53, 0x32, 0xda, 0xfa, 0x78, 0x48, 0xe0, 0x9d, 0x16, 0x2c, 0x44,
	0x85, 0xbc, 0x14, 0x2a, 0xbf, 0xc2, 0x78, 0x44, 0xe0, 0x98, 0x31, 0x00, 0x87, 0x99, 0x45, 0xcf,
	0x3f, 0x61, 0x13, 0xdf, 0xea, 0xd9, 0xd2, 0xf1, 0x8f, 0x74, 0x8f, 0xb8, 0xa0, 0xc6, 0x31, 0x81,
	0x13, 0xf6, 0x08, 0xee, 0x67, 0xb5, 0xb5, 0xa8, 0x3c, 0x57, 0x75, 0xc5, 0x6f, 0x1c, 0xb9, 0x78,
	0x27, 0x08, 0x1d, 0x00, 0xab, 0xc4, 0x97, 0x9f, 0x6b, 0xbb, 0xa1, 0x96, 0x42, 0x14, 0x6e, 0xf6,
	0x8d, 0xc1, 0x78, 0x42, 0xd0, 0xec, 0xe8, 0x5e, 0xb2, 0x31, 0x76, 0x4a, 0x85, 0x73, 0xc2, 0x9f,
	0xc3, 0xe9, 0x39, 0x7f, 0xbf, 0x7e, 0xbd, 0x7e, 0xf3, 0x61, 0xcd, 0x9e, 0xc1, 0xf4, 0x52, 0xdb,
	0x4a, 0xf8, 0x8e, 0x72, 0x3b, 0x50, 0x58, 0x4f, 0x39, 0x09, 0xa5, 0x8e, 0x74, 0xb2, 0x25, 0x91,
	0x95, 0x4a, 0x3a, 0x27, 0xb5, 0xe2, 0x32, 0x8f, 0xa3, 0x70, 0xf3, 0x53, 0xda, 0x51, 0x37, 0x0d,
	0xe6, 0xdc, 0x5c, 0x89, 0x06, 0x69, 0x28, 0x98, 0x8f, 0x96, 0xd3, 0xa3, 0x59, 0x2f, 0x77, 0xac,
	0xd0, 0xaf, 0x72, 0x36, 0x87, 0x89, 0xa1, 0x3c, 0x9a, 0x96, 0x3c, 0x25, 0xf2, 0x2f, 0x1d, 0x0b,
	0x01, 0x3b, 0x2b, 0x75, 0x2d, 0x3d, 0x92, 0xa1, 0x07, 0x32, 0x9c, 0x2c, 0x0f, 0xc6, 0x44, 0x96,
	0xe9, 0x9a, 0x36, 0x43, 0xd4, 0xc1, 0x7c, 0xf8, 0x1b, 0xf1, 0x27, 0x10, 0xb5, 0xfb, 0xe4, 0x92,
	0x82, 0xa3, 0x38, 0xda, 0x96, 0xfd, 0x64, 0xfb, 0x4c, 0xfa, 0x4c, 0x17, 0x5f, 0x07, 0x70, 0xf7,
	0x6d, 0x3b, 0xc5, 0x59, 0x88, 0xef, 0xff, 0xe6, 0xfc, 0x43, 0x2c, 0xe3, 0xbf, 0x8f, 0x65, 0xf7,
	0x1f, 0x62, 0x09, 0x8f, 0x65, 0xf1, 0x18, 0x66, 0xc1, 0xe5, 0xbb, 0x1a, 0x9d, 0x0f, 0x26, 0xf7,
	0x20, 0xfa, 0xdc, 0xfe, 0x50, 0x4f, 0x6b, 0x6e, 0xb4, 0x1c, 0x7f, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0x54, 0x84, 0x96, 0x7a, 0x03, 0x00, 0x00,
}
