// Code generated by protoc-gen-go.
// source: spectatorproto/spectatorproto.proto
// DO NOT EDIT!

/*
Package spectatorproto is a generated protocol buffer package.

It is generated from these files:
	spectatorproto/spectatorproto.proto

It has these top-level messages:
	Invite
	JoinInfo
	PartyQuestInfo
	PartyServerInfo
*/
package spectatorproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pegasus_pegasusshared "pegasus/pegasusshared"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ref: SpectatorProto.Invite
type Invite struct {
	InviterGameAccountId *pegasus_pegasusshared.BnetId `protobuf:"bytes,1,opt,name=inviter_game_account_id,json=inviterGameAccountId" json:"inviter_game_account_id,omitempty"`
	JoinInfo             *JoinInfo                     `protobuf:"bytes,2,opt,name=join_info,json=joinInfo" json:"join_info,omitempty"`
}

func (m *Invite) Reset()                    { *m = Invite{} }
func (m *Invite) String() string            { return proto.CompactTextString(m) }
func (*Invite) ProtoMessage()               {}
func (*Invite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Invite) GetInviterGameAccountId() *pegasus_pegasusshared.BnetId {
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

// ref: SpectatorProto.JoinInfo
type JoinInfo struct {
	ServerIpAddress      string                           `protobuf:"bytes,1,opt,name=server_ip_address,json=serverIpAddress" json:"server_ip_address,omitempty"`
	ServerPort           uint32                           `protobuf:"varint,2,opt,name=server_port,json=serverPort" json:"server_port,omitempty"`
	GameHandle           int32                            `protobuf:"varint,3,opt,name=game_handle,json=gameHandle" json:"game_handle,omitempty"`
	SecretKey            string                           `protobuf:"bytes,4,opt,name=secret_key,json=secretKey" json:"secret_key,omitempty"`
	IsJoinable           bool                             `protobuf:"varint,5,opt,name=is_joinable,json=isJoinable" json:"is_joinable,omitempty"`
	CurrentNumSpectators int32                            `protobuf:"varint,6,opt,name=current_num_spectators,json=currentNumSpectators" json:"current_num_spectators,omitempty"`
	MaxNumSpectators     int32                            `protobuf:"varint,7,opt,name=max_num_spectators,json=maxNumSpectators" json:"max_num_spectators,omitempty"`
	GameType             pegasus_pegasusshared.GameType   `protobuf:"varint,8,opt,name=game_type,json=gameType,enum=pegasus.pegasusshared.GameType" json:"game_type,omitempty"`
	MissionId            int32                            `protobuf:"varint,9,opt,name=mission_id,json=missionId" json:"mission_id,omitempty"`
	SpectatedPlayers     []*pegasus_pegasusshared.BnetId  `protobuf:"bytes,10,rep,name=spectated_players,json=spectatedPlayers" json:"spectated_players,omitempty"`
	PartyId              *pegasus_pegasusshared.BnetId    `protobuf:"bytes,11,opt,name=party_id,json=partyId" json:"party_id,omitempty"`
	FormatType           pegasus_pegasusshared.FormatType `protobuf:"varint,12,opt,name=format_type,json=formatType,enum=pegasus.pegasusshared.FormatType" json:"format_type,omitempty"`
}

func (m *JoinInfo) Reset()                    { *m = JoinInfo{} }
func (m *JoinInfo) String() string            { return proto.CompactTextString(m) }
func (*JoinInfo) ProtoMessage()               {}
func (*JoinInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JoinInfo) GetServerIpAddress() string {
	if m != nil {
		return m.ServerIpAddress
	}
	return ""
}

func (m *JoinInfo) GetServerPort() uint32 {
	if m != nil {
		return m.ServerPort
	}
	return 0
}

func (m *JoinInfo) GetGameHandle() int32 {
	if m != nil {
		return m.GameHandle
	}
	return 0
}

func (m *JoinInfo) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *JoinInfo) GetIsJoinable() bool {
	if m != nil {
		return m.IsJoinable
	}
	return false
}

func (m *JoinInfo) GetCurrentNumSpectators() int32 {
	if m != nil {
		return m.CurrentNumSpectators
	}
	return 0
}

func (m *JoinInfo) GetMaxNumSpectators() int32 {
	if m != nil {
		return m.MaxNumSpectators
	}
	return 0
}

func (m *JoinInfo) GetGameType() pegasus_pegasusshared.GameType {
	if m != nil {
		return m.GameType
	}
	return pegasus_pegasusshared.GameType_GT_UNKNOWN
}

func (m *JoinInfo) GetMissionId() int32 {
	if m != nil {
		return m.MissionId
	}
	return 0
}

func (m *JoinInfo) GetSpectatedPlayers() []*pegasus_pegasusshared.BnetId {
	if m != nil {
		return m.SpectatedPlayers
	}
	return nil
}

func (m *JoinInfo) GetPartyId() *pegasus_pegasusshared.BnetId {
	if m != nil {
		return m.PartyId
	}
	return nil
}

func (m *JoinInfo) GetFormatType() pegasus_pegasusshared.FormatType {
	if m != nil {
		return m.FormatType
	}
	return pegasus_pegasusshared.FormatType_FT_UNKNOWN
}

// ref: SpectatorProto.PartyQuestInfo
type PartyQuestInfo struct {
	QuestIds []int32 `protobuf:"varint,1,rep,name=quest_ids,json=questIds" json:"quest_ids,omitempty"`
}

func (m *PartyQuestInfo) Reset()                    { *m = PartyQuestInfo{} }
func (m *PartyQuestInfo) String() string            { return proto.CompactTextString(m) }
func (*PartyQuestInfo) ProtoMessage()               {}
func (*PartyQuestInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PartyQuestInfo) GetQuestIds() []int32 {
	if m != nil {
		return m.QuestIds
	}
	return nil
}

// ref: SpectatorProto.PartyServerInfo
type PartyServerInfo struct {
	ServerIpAddress string                           `protobuf:"bytes,1,opt,name=server_ip_address,json=serverIpAddress" json:"server_ip_address,omitempty"`
	ServerPort      uint32                           `protobuf:"varint,2,opt,name=server_port,json=serverPort" json:"server_port,omitempty"`
	GameHandle      int32                            `protobuf:"varint,3,opt,name=game_handle,json=gameHandle" json:"game_handle,omitempty"`
	SecretKey       string                           `protobuf:"bytes,4,opt,name=secret_key,json=secretKey" json:"secret_key,omitempty"`
	GameType        pegasus_pegasusshared.GameType   `protobuf:"varint,5,opt,name=game_type,json=gameType,enum=pegasus.pegasusshared.GameType" json:"game_type,omitempty"`
	MissionId       int32                            `protobuf:"varint,6,opt,name=mission_id,json=missionId" json:"mission_id,omitempty"`
	FormatType      pegasus_pegasusshared.FormatType `protobuf:"varint,7,opt,name=format_type,json=formatType,enum=pegasus.pegasusshared.FormatType" json:"format_type,omitempty"`
}

func (m *PartyServerInfo) Reset()                    { *m = PartyServerInfo{} }
func (m *PartyServerInfo) String() string            { return proto.CompactTextString(m) }
func (*PartyServerInfo) ProtoMessage()               {}
func (*PartyServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PartyServerInfo) GetServerIpAddress() string {
	if m != nil {
		return m.ServerIpAddress
	}
	return ""
}

func (m *PartyServerInfo) GetServerPort() uint32 {
	if m != nil {
		return m.ServerPort
	}
	return 0
}

func (m *PartyServerInfo) GetGameHandle() int32 {
	if m != nil {
		return m.GameHandle
	}
	return 0
}

func (m *PartyServerInfo) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *PartyServerInfo) GetGameType() pegasus_pegasusshared.GameType {
	if m != nil {
		return m.GameType
	}
	return pegasus_pegasusshared.GameType_GT_UNKNOWN
}

func (m *PartyServerInfo) GetMissionId() int32 {
	if m != nil {
		return m.MissionId
	}
	return 0
}

func (m *PartyServerInfo) GetFormatType() pegasus_pegasusshared.FormatType {
	if m != nil {
		return m.FormatType
	}
	return pegasus_pegasusshared.FormatType_FT_UNKNOWN
}

func init() {
	proto.RegisterType((*Invite)(nil), "spectatorproto.Invite")
	proto.RegisterType((*JoinInfo)(nil), "spectatorproto.JoinInfo")
	proto.RegisterType((*PartyQuestInfo)(nil), "spectatorproto.PartyQuestInfo")
	proto.RegisterType((*PartyServerInfo)(nil), "spectatorproto.PartyServerInfo")
}

func init() { proto.RegisterFile("spectatorproto/spectatorproto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x65, 0x9b, 0x26, 0xd9, 0x9d, 0x40, 0x9a, 0x5a, 0x15, 0xac, 0x90, 0xaa, 0x84, 0x70, 0x09,
	0x08, 0xa5, 0xa2, 0x80, 0xc4, 0x81, 0x4b, 0x73, 0x00, 0xb6, 0x48, 0x55, 0xd8, 0xf6, 0x6e, 0xb9,
	0xf1, 0x24, 0x75, 0xc9, 0xda, 0x8b, 0xed, 0xad, 0xba, 0x1f, 0xc2, 0xf7, 0x70, 0xe4, 0xb7, 0x90,
	0xbd, 0x4b, 0xa4, 0x8d, 0x54, 0x81, 0xe0, 0xc2, 0x25, 0x19, 0x3f, 0xcf, 0x9b, 0x99, 0xe7, 0x79,
	0x0b, 0x4f, 0x4d, 0x8e, 0x0b, 0xcb, 0xac, 0xd2, 0xb9, 0x56, 0x56, 0x1d, 0x35, 0x8f, 0x53, 0xff,
	0x4b, 0xfa, 0x4d, 0xf4, 0xf1, 0xb3, 0x1c, 0x57, 0xcc, 0x14, 0xe6, 0xa8, 0xfe, 0x37, 0x57, 0x4c,
	0x23, 0x6f, 0x9e, 0x2a, 0xea, 0xf8, 0x5b, 0x00, 0x9d, 0x44, 0xde, 0x08, 0x8b, 0xe4, 0x02, 0x1e,
	0x09, 0x1f, 0x69, 0xba, 0x62, 0x19, 0x52, 0xb6, 0x58, 0xa8, 0x42, 0x5a, 0x2a, 0x78, 0x1c, 0x8c,
	0x82, 0x49, 0xef, 0xf8, 0x70, 0x5a, 0x57, 0x98, 0x36, 0x2b, 0xcd, 0x24, 0xda, 0x84, 0xa7, 0x07,
	0x35, 0xfb, 0x03, 0xcb, 0xf0, 0xa4, 0xe2, 0x26, 0x9c, 0xbc, 0x81, 0xe8, 0x5a, 0x09, 0x49, 0x85,
	0x5c, 0xaa, 0x78, 0xc7, 0xd7, 0x89, 0xa7, 0x5b, 0x2a, 0x4e, 0x95, 0x90, 0x89, 0x5c, 0xaa, 0x34,
	0xbc, 0xae, 0xa3, 0xf1, 0x8f, 0x5d, 0x08, 0x7f, 0xc1, 0xe4, 0x39, 0xec, 0x1b, 0xd4, 0x37, 0xa8,
	0xa9, 0xc8, 0x29, 0xe3, 0x5c, 0xa3, 0x31, 0x7e, 0xa6, 0x28, 0xdd, 0xab, 0x2e, 0x92, 0xfc, 0xa4,
	0x82, 0xc9, 0x10, 0x7a, 0x75, 0x6e, 0xae, 0xb4, 0xf5, 0x1d, 0x1f, 0xa4, 0x50, 0x41, 0x73, 0xa5,
	0xad, 0x4b, 0xf0, 0xf2, 0xae, 0x98, 0xe4, 0x6b, 0x8c, 0x5b, 0xa3, 0x60, 0xd2, 0x4e, 0xc1, 0x41,
	0x1f, 0x3d, 0x42, 0x0e, 0x01, 0x0c, 0x2e, 0x34, 0x5a, 0xfa, 0x05, 0xcb, 0x78, 0xd7, 0xb7, 0x89,
	0x2a, 0xe4, 0x13, 0x96, 0x8e, 0x2f, 0x0c, 0x75, 0x83, 0xb2, 0xcb, 0x35, 0xc6, 0xed, 0x51, 0x30,
	0x09, 0x53, 0x10, 0xe6, 0xb4, 0x46, 0xc8, 0x6b, 0x78, 0xb8, 0x28, 0xb4, 0x46, 0x69, 0xa9, 0x2c,
	0x32, 0xba, 0xd1, 0x6a, 0xe2, 0x8e, 0xef, 0x75, 0x50, 0xdf, 0x9e, 0x15, 0xd9, 0xf9, 0xe6, 0x8e,
	0xbc, 0x00, 0x92, 0xb1, 0xdb, 0x6d, 0x46, 0xd7, 0x33, 0x06, 0x19, 0xbb, 0x6d, 0x66, 0xbf, 0x83,
	0xc8, 0x8b, 0xb0, 0x65, 0x8e, 0x71, 0x38, 0x0a, 0x26, 0xfd, 0xe3, 0xe1, 0x1d, 0xdb, 0x71, 0xeb,
	0xb8, 0x28, 0x73, 0x4c, 0xc3, 0x55, 0x1d, 0x39, 0x85, 0x99, 0x30, 0x46, 0x28, 0xe9, 0x96, 0x1b,
	0xf9, 0x1e, 0x51, 0x8d, 0x24, 0x9c, 0x9c, 0xc1, 0x7e, 0x3d, 0x02, 0x72, 0x9a, 0xaf, 0x59, 0x89,
	0xda, 0xc4, 0x30, 0x6a, 0xfd, 0xd6, 0x02, 0xb3, 0x9d, 0xc1, 0xbd, 0x74, 0xb0, 0xe1, 0xce, 0x2b,
	0x2a, 0x79, 0x0b, 0x61, 0xce, 0xb4, 0x2d, 0x5d, 0xb3, 0xde, 0x9f, 0x38, 0xa9, 0xeb, 0xd3, 0x13,
	0x4e, 0x66, 0xd0, 0x5b, 0x2a, 0x9d, 0x31, 0x5b, 0x09, 0xbd, 0xef, 0x85, 0x3e, 0xb9, 0x83, 0xfc,
	0xde, 0x67, 0x7a, 0xa9, 0xb0, 0xdc, 0xc4, 0xe3, 0x97, 0xd0, 0x9f, 0xbb, 0x72, 0x9f, 0x0b, 0x34,
	0xd6, 0xdb, 0x69, 0x08, 0xd1, 0x57, 0x77, 0xa0, 0x82, 0x3b, 0x1b, 0xb5, 0x26, 0x6d, 0x3f, 0x78,
	0xe8, 0xc1, 0x84, 0x9b, 0xf1, 0xf7, 0x1d, 0xd8, 0xf3, 0x9c, 0xf3, 0xca, 0x5c, 0xff, 0x9d, 0x07,
	0x1b, 0xeb, 0x6f, 0xff, 0xdb, 0xfa, 0x3b, 0xdb, 0xeb, 0xdf, 0x7a, 0xf4, 0xee, 0x5f, 0x3c, 0xfa,
	0x65, 0xc7, 0x7f, 0xd8, 0xaf, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x7c, 0xe4, 0xab, 0xbf,
	0x04, 0x00, 0x00,
}
