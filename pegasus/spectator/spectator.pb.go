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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JoinInfo struct {
	ServerIpAddress      *string            `protobuf:"bytes,1,opt,name=server_ip_address,json=serverIpAddress" json:"server_ip_address,omitempty"`
	ServerPort           *uint32            `protobuf:"varint,2,opt,name=server_port,json=serverPort" json:"server_port,omitempty"`
	GameHandle           *int32             `protobuf:"varint,3,opt,name=game_handle,json=gameHandle" json:"game_handle,omitempty"`
	SecretKey            *string            `protobuf:"bytes,4,opt,name=secret_key,json=secretKey" json:"secret_key,omitempty"`
	IsJoinable           *bool              `protobuf:"varint,5,opt,name=is_joinable,json=isJoinable" json:"is_joinable,omitempty"`
	CurrentNumSpectators *int32             `protobuf:"varint,6,opt,name=current_num_spectators,json=currentNumSpectators" json:"current_num_spectators,omitempty"`
	MaxNumSpectators     *int32             `protobuf:"varint,7,opt,name=max_num_spectators,json=maxNumSpectators" json:"max_num_spectators,omitempty"`
	GameType             *shared.GameType   `protobuf:"varint,8,opt,name=game_type,json=gameType,enum=shared.GameType,def=0" json:"game_type,omitempty"`
	FormatType           *shared.FormatType `protobuf:"varint,12,opt,name=format_type,json=formatType,enum=shared.FormatType,def=0" json:"format_type,omitempty"`
	MissionId            *int32             `protobuf:"varint,9,opt,name=mission_id,json=missionId" json:"mission_id,omitempty"`
	SpectatedPlayers     []*shared.BnetId   `protobuf:"bytes,10,rep,name=spectated_players,json=spectatedPlayers" json:"spectated_players,omitempty"`
	PartyId              *shared.BnetId     `protobuf:"bytes,11,opt,name=party_id,json=partyId" json:"party_id,omitempty"`
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
	InviterGameAccountId *shared.BnetId `protobuf:"bytes,1,req,name=inviter_game_account_id,json=inviterGameAccountId" json:"inviter_game_account_id,omitempty"`
	JoinInfo             *JoinInfo      `protobuf:"bytes,2,req,name=join_info,json=joinInfo" json:"join_info,omitempty"`
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
	ServerIpAddress  *string            `protobuf:"bytes,1,opt,name=server_ip_address,json=serverIpAddress" json:"server_ip_address,omitempty"`
	ServerPort       *uint32            `protobuf:"varint,2,opt,name=server_port,json=serverPort" json:"server_port,omitempty"`
	GameHandle       *int32             `protobuf:"varint,3,opt,name=game_handle,json=gameHandle" json:"game_handle,omitempty"`
	SecretKey        *string            `protobuf:"bytes,4,opt,name=secret_key,json=secretKey" json:"secret_key,omitempty"`
	GameType         *shared.GameType   `protobuf:"varint,5,opt,name=game_type,json=gameType,enum=shared.GameType,def=0" json:"game_type,omitempty"`
	FormatType       *shared.FormatType `protobuf:"varint,7,opt,name=format_type,json=formatType,enum=shared.FormatType,def=0" json:"format_type,omitempty"`
	MissionId        *int32             `protobuf:"varint,6,opt,name=mission_id,json=missionId" json:"mission_id,omitempty"`
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
	QuestIds         []int32 `protobuf:"varint,1,rep,name=quest_ids,json=questIds" json:"quest_ids,omitempty"`
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

func init() {
	proto.RegisterFile("github.com/HearthSim/hs-proto-go/pegasus/spectator/spectator.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x52, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xba, 0x75, 0x4b, 0x6e, 0x61, 0x1f, 0x66, 0x82, 0x08, 0x84, 0x88, 0xfa, 0x14, 0x10,
	0x6d, 0xd1, 0xc4, 0x0b, 0x1f, 0x12, 0xda, 0x24, 0xb6, 0x65, 0x93, 0x4a, 0x49, 0x87, 0x78, 0xb4,
	0xbc, 0xf8, 0xb6, 0xf5, 0x68, 0xe2, 0x60, 0x3b, 0xd3, 0xf2, 0xc8, 0x2f, 0xe1, 0x77, 0xf0, 0xef,
	0x50, 0x9c, 0xb4, 0xd5, 0xa6, 0x3d, 0x20, 0xed, 0x85, 0xa7, 0xdc, 0x9c, 0x7b, 0xce, 0xbd, 0xf6,
	0xf1, 0x81, 0xc3, 0xa9, 0x30, 0xb3, 0xe2, 0xa2, 0x9f, 0xc8, 0x74, 0x70, 0x82, 0x4c, 0x99, 0xd9,
	0x58, 0xa4, 0x83, 0x99, 0xee, 0xe5, 0x4a, 0x1a, 0xd9, 0x9b, 0xca, 0x41, 0x8e, 0x53, 0xa6, 0x0b,
	0x3d, 0xd0, 0x39, 0x26, 0x86, 0x19, 0xa9, 0x56, 0x55, 0xdf, 0x92, 0x88, 0xb7, 0x04, 0x9e, 0x7e,
	0xfc, 0xf7, 0x71, 0x33, 0xa6, 0x90, 0x37, 0x9f, 0x7a, 0x50, 0xf7, 0xf7, 0x3a, 0xb8, 0xa7, 0x52,
	0x64, 0x51, 0x36, 0x91, 0xe4, 0x15, 0xec, 0x6a, 0x54, 0x57, 0xa8, 0xa8, 0xc8, 0x29, 0xe3, 0x5c,
	0xa1, 0xd6, 0xbe, 0x13, 0x38, 0xa1, 0x17, 0x6f, 0xd7, 0x8d, 0x28, 0x3f, 0xa8, 0x61, 0xf2, 0x02,
	0x3a, 0x0d, 0x37, 0x97, 0xca, 0xf8, 0xad, 0xc0, 0x09, 0x1f, 0xc6, 0x50, 0x43, 0x23, 0xa9, 0x4c,
	0x45, 0x98, 0xb2, 0x14, 0xe9, 0x8c, 0x65, 0x7c, 0x8e, 0xfe, 0x5a, 0xe0, 0x84, 0xed, 0x18, 0x2a,
	0xe8, 0xc4, 0x22, 0xe4, 0x39, 0x80, 0xc6, 0x44, 0xa1, 0xa1, 0x3f, 0xb0, 0xf4, 0xd7, 0xed, 0x1a,
	0xaf, 0x46, 0xce, 0xb0, 0xac, 0xf4, 0x42, 0xd3, 0x4b, 0x29, 0x32, 0x76, 0x31, 0x47, 0xbf, 0x1d,
	0x38, 0xa1, 0x1b, 0x83, 0xd0, 0xa7, 0x0d, 0x42, 0xde, 0xc2, 0xe3, 0xa4, 0x50, 0x0a, 0x33, 0x43,
	0xb3, 0x22, 0xa5, 0x4b, 0x47, 0xb4, 0xbf, 0x61, 0x77, 0xed, 0x35, 0xdd, 0x61, 0x91, 0x8e, 0x97,
	0x3d, 0xf2, 0x1a, 0x48, 0xca, 0xae, 0x6f, 0x2b, 0x36, 0xad, 0x62, 0x27, 0x65, 0xd7, 0x37, 0xd9,
	0xef, 0xc0, 0xb3, 0x97, 0x30, 0x65, 0x8e, 0xbe, 0x1b, 0x38, 0xe1, 0xd6, 0xfe, 0x4e, 0xbf, 0x31,
	0xf0, 0x98, 0xa5, 0x78, 0x5e, 0xe6, 0xf8, 0x1e, 0x8e, 0xcf, 0xe9, 0xb7, 0xe1, 0xd9, 0xf0, 0xcb,
	0xf7, 0x61, 0xec, 0x4e, 0x1b, 0x94, 0x7c, 0x82, 0xce, 0x44, 0xaa, 0x94, 0x99, 0x5a, 0xfc, 0xc0,
	0x8a, 0xc9, 0x42, 0x7c, 0x64, 0x5b, 0xb5, 0xfc, 0x68, 0x25, 0x87, 0xc9, 0x12, 0xaf, 0xfc, 0x49,
	0x85, 0xd6, 0x42, 0x66, 0x54, 0x70, 0xdf, 0xb3, 0x27, 0xf4, 0x1a, 0x24, 0xe2, 0xe4, 0x03, 0xec,
	0x36, 0x17, 0x40, 0x4e, 0xf3, 0x39, 0x2b, 0x51, 0x69, 0x1f, 0x82, 0xb5, 0xb0, 0xb3, 0xbf, 0xb5,
	0xd8, 0x72, 0x98, 0xa1, 0x89, 0x78, 0xbc, 0xb3, 0x24, 0x8e, 0x6a, 0x1e, 0x79, 0x09, 0x6e, 0xce,
	0x94, 0x29, 0xab, 0xc9, 0x9d, 0xc0, 0xb9, 0x43, 0xb3, 0x69, 0xfb, 0x11, 0xef, 0xfe, 0x72, 0x60,
	0x23, 0xca, 0xae, 0x84, 0x41, 0xf2, 0x19, 0x9e, 0x08, 0x5b, 0x29, 0x6a, 0x5d, 0x61, 0x49, 0x22,
	0x8b, 0xcc, 0x54, 0x43, 0x9c, 0xa0, 0x75, 0xc7, 0x90, 0xbd, 0x86, 0x5e, 0x59, 0x75, 0x50, 0x93,
	0x23, 0x4e, 0xde, 0x80, 0x57, 0x3d, 0x2b, 0x15, 0xd9, 0x44, 0xfa, 0x2d, 0x2b, 0x7c, 0xd4, 0x5f,
	0x25, 0x7c, 0x11, 0xc7, 0xd8, 0xbd, 0x6c, 0xaa, 0xee, 0x9f, 0x16, 0x6c, 0x8f, 0xaa, 0xf3, 0x8c,
	0xeb, 0x14, 0xfe, 0x77, 0x61, 0xbd, 0x91, 0x93, 0xf6, 0x7d, 0x72, 0xb2, 0x79, 0xcf, 0x9c, 0x6c,
	0xdc, 0xca, 0x49, 0xb7, 0x07, 0x5b, 0xd6, 0xba, 0xaf, 0x05, 0x6a, 0x63, 0x9d, 0x7b, 0x06, 0xde,
	0xcf, 0xea, 0x87, 0x0a, 0x5e, 0x39, 0xb6, 0x16, 0xb6, 0x63, 0xd7, 0x02, 0x11, 0xd7, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0x4e, 0xea, 0xbb, 0x9e, 0x04, 0x00, 0x00,
}