// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/game_master_types/game_master_types.proto
// DO NOT EDIT!

/*
Package game_master_types is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/game_master_types/game_master_types.proto

It has these top-level messages:
	Player
	ConnectInfo
	GameStatsBucket
	GameFactoryDescription
	GameHandle
	CancelGameEntryRequest
*/
package game_master_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import attribute "github.com/HearthSim/hs-proto-go/bnet/attribute"
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

type Player struct {
	Identity         *entity.Identity       `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,2,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Player) GetIdentity() *entity.Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Player) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type ConnectInfo struct {
	MemberId         *entity.EntityId       `protobuf:"bytes,1,req,name=member_id,json=memberId" json:"member_id,omitempty"`
	Host             *string                `protobuf:"bytes,2,req,name=host" json:"host,omitempty"`
	Port             *int32                 `protobuf:"varint,3,req,name=port" json:"port,omitempty"`
	Token            []byte                 `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,5,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ConnectInfo) Reset()                    { *m = ConnectInfo{} }
func (m *ConnectInfo) String() string            { return proto.CompactTextString(m) }
func (*ConnectInfo) ProtoMessage()               {}
func (*ConnectInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConnectInfo) GetMemberId() *entity.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *ConnectInfo) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return ""
}

func (m *ConnectInfo) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ConnectInfo) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *ConnectInfo) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type GameStatsBucket struct {
	BucketMin                  *float32 `protobuf:"fixed32,1,opt,name=bucket_min,json=bucketMin,def=0" json:"bucket_min,omitempty"`
	BucketMax                  *float32 `protobuf:"fixed32,2,opt,name=bucket_max,json=bucketMax,def=4.294967e+009" json:"bucket_max,omitempty"`
	WaitMilliseconds           *uint32  `protobuf:"varint,3,opt,name=wait_milliseconds,json=waitMilliseconds,def=0" json:"wait_milliseconds,omitempty"`
	GamesPerHour               *uint32  `protobuf:"varint,4,opt,name=games_per_hour,json=gamesPerHour,def=0" json:"games_per_hour,omitempty"`
	ActiveGames                *uint32  `protobuf:"varint,5,opt,name=active_games,json=activeGames,def=0" json:"active_games,omitempty"`
	ActivePlayers              *uint32  `protobuf:"varint,6,opt,name=active_players,json=activePlayers,def=0" json:"active_players,omitempty"`
	FormingGames               *uint32  `protobuf:"varint,7,opt,name=forming_games,json=formingGames,def=0" json:"forming_games,omitempty"`
	WaitingPlayers             *uint32  `protobuf:"varint,8,opt,name=waiting_players,json=waitingPlayers,def=0" json:"waiting_players,omitempty"`
	OpenJoinableGames          *uint32  `protobuf:"varint,9,opt,name=open_joinable_games,json=openJoinableGames,def=0" json:"open_joinable_games,omitempty"`
	PlayersInOpenJoinableGames *uint32  `protobuf:"varint,10,opt,name=players_in_open_joinable_games,json=playersInOpenJoinableGames,def=0" json:"players_in_open_joinable_games,omitempty"`
	OpenGamesTotal             *uint32  `protobuf:"varint,11,opt,name=open_games_total,json=openGamesTotal,def=0" json:"open_games_total,omitempty"`
	PlayersInOpenGamesTotal    *uint32  `protobuf:"varint,12,opt,name=players_in_open_games_total,json=playersInOpenGamesTotal,def=0" json:"players_in_open_games_total,omitempty"`
	XXX_unrecognized           []byte   `json:"-"`
}

func (m *GameStatsBucket) Reset()                    { *m = GameStatsBucket{} }
func (m *GameStatsBucket) String() string            { return proto.CompactTextString(m) }
func (*GameStatsBucket) ProtoMessage()               {}
func (*GameStatsBucket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_GameStatsBucket_BucketMin float32 = 0
const Default_GameStatsBucket_BucketMax float32 = 4.294967e+009
const Default_GameStatsBucket_WaitMilliseconds uint32 = 0
const Default_GameStatsBucket_GamesPerHour uint32 = 0
const Default_GameStatsBucket_ActiveGames uint32 = 0
const Default_GameStatsBucket_ActivePlayers uint32 = 0
const Default_GameStatsBucket_FormingGames uint32 = 0
const Default_GameStatsBucket_WaitingPlayers uint32 = 0
const Default_GameStatsBucket_OpenJoinableGames uint32 = 0
const Default_GameStatsBucket_PlayersInOpenJoinableGames uint32 = 0
const Default_GameStatsBucket_OpenGamesTotal uint32 = 0
const Default_GameStatsBucket_PlayersInOpenGamesTotal uint32 = 0

func (m *GameStatsBucket) GetBucketMin() float32 {
	if m != nil && m.BucketMin != nil {
		return *m.BucketMin
	}
	return Default_GameStatsBucket_BucketMin
}

func (m *GameStatsBucket) GetBucketMax() float32 {
	if m != nil && m.BucketMax != nil {
		return *m.BucketMax
	}
	return Default_GameStatsBucket_BucketMax
}

func (m *GameStatsBucket) GetWaitMilliseconds() uint32 {
	if m != nil && m.WaitMilliseconds != nil {
		return *m.WaitMilliseconds
	}
	return Default_GameStatsBucket_WaitMilliseconds
}

func (m *GameStatsBucket) GetGamesPerHour() uint32 {
	if m != nil && m.GamesPerHour != nil {
		return *m.GamesPerHour
	}
	return Default_GameStatsBucket_GamesPerHour
}

func (m *GameStatsBucket) GetActiveGames() uint32 {
	if m != nil && m.ActiveGames != nil {
		return *m.ActiveGames
	}
	return Default_GameStatsBucket_ActiveGames
}

func (m *GameStatsBucket) GetActivePlayers() uint32 {
	if m != nil && m.ActivePlayers != nil {
		return *m.ActivePlayers
	}
	return Default_GameStatsBucket_ActivePlayers
}

func (m *GameStatsBucket) GetFormingGames() uint32 {
	if m != nil && m.FormingGames != nil {
		return *m.FormingGames
	}
	return Default_GameStatsBucket_FormingGames
}

func (m *GameStatsBucket) GetWaitingPlayers() uint32 {
	if m != nil && m.WaitingPlayers != nil {
		return *m.WaitingPlayers
	}
	return Default_GameStatsBucket_WaitingPlayers
}

func (m *GameStatsBucket) GetOpenJoinableGames() uint32 {
	if m != nil && m.OpenJoinableGames != nil {
		return *m.OpenJoinableGames
	}
	return Default_GameStatsBucket_OpenJoinableGames
}

func (m *GameStatsBucket) GetPlayersInOpenJoinableGames() uint32 {
	if m != nil && m.PlayersInOpenJoinableGames != nil {
		return *m.PlayersInOpenJoinableGames
	}
	return Default_GameStatsBucket_PlayersInOpenJoinableGames
}

func (m *GameStatsBucket) GetOpenGamesTotal() uint32 {
	if m != nil && m.OpenGamesTotal != nil {
		return *m.OpenGamesTotal
	}
	return Default_GameStatsBucket_OpenGamesTotal
}

func (m *GameStatsBucket) GetPlayersInOpenGamesTotal() uint32 {
	if m != nil && m.PlayersInOpenGamesTotal != nil {
		return *m.PlayersInOpenGamesTotal
	}
	return Default_GameStatsBucket_PlayersInOpenGamesTotal
}

type GameFactoryDescription struct {
	Id               *uint64                `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	Name             *string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,3,rep,name=attribute" json:"attribute,omitempty"`
	StatsBucket      []*GameStatsBucket     `protobuf:"bytes,4,rep,name=stats_bucket,json=statsBucket" json:"stats_bucket,omitempty"`
	UnseededId       *uint64                `protobuf:"fixed64,5,opt,name=unseeded_id,json=unseededId,def=0" json:"unseeded_id,omitempty"`
	AllowQueueing    *bool                  `protobuf:"varint,6,opt,name=allow_queueing,json=allowQueueing,def=1" json:"allow_queueing,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *GameFactoryDescription) Reset()                    { *m = GameFactoryDescription{} }
func (m *GameFactoryDescription) String() string            { return proto.CompactTextString(m) }
func (*GameFactoryDescription) ProtoMessage()               {}
func (*GameFactoryDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_GameFactoryDescription_UnseededId uint64 = 0
const Default_GameFactoryDescription_AllowQueueing bool = true

func (m *GameFactoryDescription) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *GameFactoryDescription) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *GameFactoryDescription) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *GameFactoryDescription) GetStatsBucket() []*GameStatsBucket {
	if m != nil {
		return m.StatsBucket
	}
	return nil
}

func (m *GameFactoryDescription) GetUnseededId() uint64 {
	if m != nil && m.UnseededId != nil {
		return *m.UnseededId
	}
	return Default_GameFactoryDescription_UnseededId
}

func (m *GameFactoryDescription) GetAllowQueueing() bool {
	if m != nil && m.AllowQueueing != nil {
		return *m.AllowQueueing
	}
	return Default_GameFactoryDescription_AllowQueueing
}

type GameHandle struct {
	FactoryId        *uint64          `protobuf:"fixed64,1,req,name=factory_id,json=factoryId" json:"factory_id,omitempty"`
	GameId           *entity.EntityId `protobuf:"bytes,2,req,name=game_id,json=gameId" json:"game_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GameHandle) Reset()                    { *m = GameHandle{} }
func (m *GameHandle) String() string            { return proto.CompactTextString(m) }
func (*GameHandle) ProtoMessage()               {}
func (*GameHandle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GameHandle) GetFactoryId() uint64 {
	if m != nil && m.FactoryId != nil {
		return *m.FactoryId
	}
	return 0
}

func (m *GameHandle) GetGameId() *entity.EntityId {
	if m != nil {
		return m.GameId
	}
	return nil
}

type CancelGameEntryRequest struct {
	RequestId        *uint64   `protobuf:"fixed64,1,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	FactoryId        *uint64   `protobuf:"fixed64,2,opt,name=factory_id,json=factoryId" json:"factory_id,omitempty"`
	Player           []*Player `protobuf:"bytes,3,rep,name=player" json:"player,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *CancelGameEntryRequest) Reset()                    { *m = CancelGameEntryRequest{} }
func (m *CancelGameEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelGameEntryRequest) ProtoMessage()               {}
func (*CancelGameEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CancelGameEntryRequest) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *CancelGameEntryRequest) GetFactoryId() uint64 {
	if m != nil && m.FactoryId != nil {
		return *m.FactoryId
	}
	return 0
}

func (m *CancelGameEntryRequest) GetPlayer() []*Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func init() {
	proto.RegisterType((*Player)(nil), "game_master_types.Player")
	proto.RegisterType((*ConnectInfo)(nil), "game_master_types.ConnectInfo")
	proto.RegisterType((*GameStatsBucket)(nil), "game_master_types.GameStatsBucket")
	proto.RegisterType((*GameFactoryDescription)(nil), "game_master_types.GameFactoryDescription")
	proto.RegisterType((*GameHandle)(nil), "game_master_types.GameHandle")
	proto.RegisterType((*CancelGameEntryRequest)(nil), "game_master_types.CancelGameEntryRequest")
}

func init() {
	proto.RegisterFile("github.com/HearthSim/hs-proto-go/bnet/game_master_types/game_master_types.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x05, 0x69, 0x59, 0xb1, 0x46, 0x97, 0xd8, 0xdb, 0x20, 0x65, 0x5d, 0xb4, 0x20, 0x84, 0xa2,
	0x55, 0xeb, 0x58, 0x76, 0x8c, 0xa0, 0x85, 0xf3, 0x12, 0xb4, 0xa9, 0x5b, 0xb3, 0x40, 0x90, 0x94,
	0x29, 0xfa, 0x4a, 0xac, 0xc8, 0xb1, 0xb4, 0x09, 0xb9, 0xab, 0xec, 0x2e, 0x1b, 0xeb, 0x0f, 0xda,
	0x9f, 0xe9, 0x67, 0xf4, 0xbb, 0x8a, 0xbd, 0x50, 0x17, 0xcb, 0x0f, 0x7a, 0xe2, 0x70, 0xce, 0x99,
	0x33, 0xb3, 0xc3, 0xc3, 0x85, 0xd7, 0x53, 0xa6, 0x67, 0xf5, 0x64, 0x9c, 0x8b, 0xea, 0xec, 0x1a,
	0xa9, 0xd4, 0xb3, 0xb7, 0xac, 0x3a, 0x9b, 0xa9, 0xd3, 0xb9, 0x14, 0x5a, 0x9c, 0x4e, 0xc5, 0xd9,
	0x84, 0xa3, 0x3e, 0x9b, 0xd2, 0x0a, 0xb3, 0x8a, 0x2a, 0x8d, 0x32, 0xd3, 0x8b, 0x39, 0xaa, 0xed,
	0xcc, 0xd8, 0x56, 0x90, 0xa3, 0x2d, 0xe0, 0xf8, 0xc5, 0x6e, 0x3d, 0xa8, 0xd6, 0x92, 0x4d, 0x6a,
	0x8d, 0xab, 0xc8, 0x69, 0x1e, 0x5f, 0xee, 0x26, 0x80, 0x5c, 0x33, 0xbd, 0xf0, 0x0f, 0x57, 0x3a,
	0x7c, 0x07, 0xed, 0x37, 0x25, 0x5d, 0xa0, 0x24, 0x4f, 0xe0, 0x80, 0x15, 0x0e, 0x8b, 0x82, 0x38,
	0x18, 0x75, 0x2f, 0x0e, 0xc7, 0x9e, 0x9a, 0xf8, 0x7c, 0xba, 0x64, 0x90, 0x0b, 0xe8, 0x2c, 0xa7,
	0x88, 0xc2, 0x78, 0x6f, 0xd4, 0xbd, 0x78, 0x34, 0x5e, 0xcd, 0xf5, 0x63, 0x13, 0xa5, 0x2b, 0xda,
	0xf0, 0xdf, 0x00, 0xba, 0x2f, 0x05, 0xe7, 0x98, 0xeb, 0x84, 0xdf, 0x08, 0x72, 0x0a, 0x9d, 0x0a,
	0xab, 0x09, 0xca, 0x8c, 0x15, 0x51, 0x10, 0x87, 0xeb, 0x2d, 0xaf, 0xec, 0x23, 0x29, 0xd2, 0x03,
	0x47, 0x49, 0x0a, 0x42, 0xa0, 0x35, 0x13, 0x4a, 0x47, 0x61, 0x1c, 0x8e, 0x3a, 0xa9, 0x8d, 0x4d,
	0x6e, 0x2e, 0xa4, 0x8e, 0xf6, 0xe2, 0x70, 0xb4, 0x9f, 0xda, 0x98, 0x3c, 0x82, 0x7d, 0x2d, 0xde,
	0x23, 0x8f, 0x5a, 0x71, 0x30, 0xea, 0xa5, 0xee, 0x65, 0x73, 0xe0, 0xfd, 0xdd, 0x06, 0xfe, 0xaf,
	0x05, 0x0f, 0x7f, 0xa5, 0x15, 0xbe, 0xd5, 0x54, 0xab, 0x9f, 0xea, 0xfc, 0x3d, 0x6a, 0x12, 0x03,
	0x4c, 0x6c, 0x94, 0x55, 0x8c, 0xdb, 0x45, 0x85, 0xcf, 0x83, 0xf3, 0xb4, 0xe3, 0x92, 0xaf, 0x18,
	0x27, 0x4f, 0x56, 0x0c, 0x7a, 0x1b, 0x85, 0x96, 0xd1, 0x7f, 0x36, 0xbe, 0xb8, 0x7c, 0x76, 0xf9,
	0xfd, 0x0f, 0x78, 0x72, 0x7e, 0x7e, 0xb9, 0x64, 0xd3, 0x5b, 0x32, 0x86, 0xa3, 0x8f, 0x94, 0x19,
	0xb5, 0xb2, 0x64, 0x0a, 0x73, 0xc1, 0x0b, 0x15, 0xed, 0xc5, 0xc1, 0xa8, 0x6f, 0x64, 0x0f, 0x0d,
	0xf6, 0x6a, 0x0d, 0x22, 0xdf, 0xc0, 0xc0, 0x38, 0x48, 0x65, 0x73, 0x94, 0xd9, 0x4c, 0xd4, 0xd2,
	0x1e, 0xd3, 0x92, 0x7b, 0x16, 0x78, 0x83, 0xf2, 0x5a, 0xd4, 0x92, 0x7c, 0x05, 0x3d, 0x9a, 0x6b,
	0xf6, 0x17, 0x66, 0x36, 0x1d, 0xed, 0x37, 0xb4, 0xae, 0x4b, 0x9b, 0x83, 0x29, 0x32, 0x82, 0x81,
	0x67, 0xcd, 0xad, 0x0d, 0x54, 0xd4, 0x6e, 0x78, 0x7d, 0x07, 0x38, 0x7b, 0x28, 0xf2, 0x35, 0xf4,
	0x6f, 0x84, 0xac, 0x18, 0x9f, 0x7a, 0xc1, 0x07, 0xcb, 0xbe, 0x3e, 0xef, 0x14, 0xbf, 0x83, 0x87,
	0x66, 0x68, 0xc3, 0x6b, 0x24, 0x0f, 0x1a, 0xe6, 0xc0, 0x23, 0x8d, 0xe6, 0x53, 0xf8, 0x44, 0xcc,
	0x91, 0x67, 0xef, 0x04, 0xe3, 0x74, 0x52, 0x36, 0xa3, 0x76, 0x1a, 0xfe, 0x91, 0x41, 0x7f, 0xf3,
	0xa0, 0x93, 0xbf, 0x82, 0x2f, 0xbd, 0x6c, 0xc6, 0x78, 0x76, 0x5f, 0x35, 0x34, 0xd5, 0xc7, 0x9e,
	0x98, 0xf0, 0xd7, 0x5b, 0x32, 0x27, 0x70, 0x68, 0x6b, 0xdd, 0x2e, 0xb5, 0xd0, 0xb4, 0x8c, 0xba,
	0xcb, 0x31, 0x0d, 0x64, 0x79, 0x7f, 0x18, 0x80, 0xbc, 0x80, 0xcf, 0xef, 0xf6, 0x5c, 0xaf, 0xeb,
	0x35, 0x75, 0x9f, 0x6e, 0x34, 0x5c, 0x09, 0x0c, 0xff, 0x0e, 0xe1, 0xb1, 0x79, 0xfd, 0x85, 0xe6,
	0x5a, 0xc8, 0xc5, 0xcf, 0xa8, 0x72, 0xc9, 0xe6, 0x9a, 0x09, 0x4e, 0x06, 0x10, 0x7a, 0xf7, 0xb7,
	0xd3, 0x90, 0x59, 0x97, 0x73, 0x5a, 0xa1, 0xf5, 0x4d, 0x27, 0xb5, 0xf1, 0xa6, 0x77, 0xf7, 0x76,
	0xf2, 0x2e, 0xb9, 0x82, 0x9e, 0x32, 0xb6, 0xcd, 0x9c, 0xd5, 0xa2, 0x96, 0x2d, 0x1b, 0x8e, 0xb7,
	0xef, 0xa5, 0x3b, 0x0e, 0x4f, 0xbb, 0x6a, 0xcd, 0xee, 0x43, 0xe8, 0xd6, 0x5c, 0x21, 0x16, 0x58,
	0x98, 0xbf, 0xd4, 0x98, 0xa8, 0x6d, 0x8e, 0x0a, 0x4d, 0x36, 0x29, 0xc8, 0x09, 0x0c, 0x68, 0x59,
	0x8a, 0x8f, 0xd9, 0x87, 0x1a, 0x6b, 0x64, 0x7c, 0x6a, 0x3d, 0x74, 0xf0, 0xbc, 0xa5, 0x65, 0x8d,
	0x69, 0xdf, 0x62, 0xbf, 0x7b, 0x68, 0xf8, 0x27, 0x80, 0x69, 0x78, 0x4d, 0x79, 0x51, 0x22, 0xf9,
	0x02, 0xe0, 0xc6, 0xed, 0x24, 0x5b, 0x6e, 0xa1, 0xe3, 0x33, 0x49, 0x41, 0xbe, 0x85, 0x07, 0x76,
	0x5e, 0x56, 0xd8, 0xbf, 0xfe, 0xbe, 0xfb, 0xa1, 0x6d, 0x08, 0x49, 0x31, 0xfc, 0x27, 0x80, 0xc7,
	0x2f, 0x29, 0xcf, 0xb1, 0x34, 0xf2, 0x57, 0x5c, 0xcb, 0x45, 0x8a, 0x1f, 0x6a, 0x54, 0xda, 0x34,
	0x91, 0x2e, 0x5c, 0x6b, 0xe2, 0x33, 0x49, 0x71, 0x67, 0x06, 0xb3, 0xf7, 0x8d, 0x19, 0x9e, 0x42,
	0xdb, 0x7d, 0x56, 0xbf, 0xf9, 0xcf, 0xee, 0x59, 0xa1, 0xf3, 0x73, 0xea, 0x89, 0xff, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0x34, 0xaa, 0xbf, 0x35, 0x06, 0x00, 0x00,
}