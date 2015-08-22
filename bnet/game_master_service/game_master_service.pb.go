// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/game_master_service/game_master_service.proto
// DO NOT EDIT!

/*
Package game_master_service is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/game_master_service/game_master_service.proto

It has these top-level messages:
	JoinGameRequest
	JoinGameResponse
	ListFactoriesRequest
	ListFactoriesResponse
	FindGameRequest
	FindGameResponse
	GameEndedNotification
	PlayerLeftNotification
	RegisterServerRequest
	UnregisterServerRequest
	RegisterUtilitiesRequest
	UnregisterUtilitiesRequest
	SubscribeRequest
	SubscribeResponse
	UnsubscribeRequest
	ChangeGameRequest
	GetFactoryInfoRequest
	GetFactoryInfoResponse
	GetGameStatsRequest
	GetGameStatsResponse
	FactoryUpdateNotification
	GameFoundNotification
*/
package game_master_service

import proto "github.com/golang/protobuf/proto"
import math "math"
import attribute "github.com/HearthSim/hs-proto-go/bnet/attribute"
import entity "github.com/HearthSim/hs-proto-go/bnet/entity"
import game_factory "github.com/HearthSim/hs-proto-go/bnet/game_factory"
import game_master_types "github.com/HearthSim/hs-proto-go/bnet/game_master_types"
import server_pool_types "github.com/HearthSim/hs-proto-go/bnet/server_pool_types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FactoryUpdateNotification_Operation int32

const (
	FactoryUpdateNotification_ADD    FactoryUpdateNotification_Operation = 1
	FactoryUpdateNotification_REMOVE FactoryUpdateNotification_Operation = 2
	FactoryUpdateNotification_CHANGE FactoryUpdateNotification_Operation = 3
)

var FactoryUpdateNotification_Operation_name = map[int32]string{
	1: "ADD",
	2: "REMOVE",
	3: "CHANGE",
}
var FactoryUpdateNotification_Operation_value = map[string]int32{
	"ADD":    1,
	"REMOVE": 2,
	"CHANGE": 3,
}

func (x FactoryUpdateNotification_Operation) Enum() *FactoryUpdateNotification_Operation {
	p := new(FactoryUpdateNotification_Operation)
	*p = x
	return p
}
func (x FactoryUpdateNotification_Operation) String() string {
	return proto.EnumName(FactoryUpdateNotification_Operation_name, int32(x))
}
func (x *FactoryUpdateNotification_Operation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FactoryUpdateNotification_Operation_value, data, "FactoryUpdateNotification_Operation")
	if err != nil {
		return err
	}
	*x = FactoryUpdateNotification_Operation(value)
	return nil
}

type JoinGameRequest struct {
	GameHandle           *game_master_types.GameHandle `protobuf:"bytes,1,req,name=game_handle" json:"game_handle,omitempty"`
	Player               []*game_master_types.Player   `protobuf:"bytes,2,rep,name=player" json:"player,omitempty"`
	AdvancedNotification *bool                         `protobuf:"varint,3,opt,name=advanced_notification,def=0" json:"advanced_notification,omitempty"`
	XXX_unrecognized     []byte                        `json:"-"`
}

func (m *JoinGameRequest) Reset()         { *m = JoinGameRequest{} }
func (m *JoinGameRequest) String() string { return proto.CompactTextString(m) }
func (*JoinGameRequest) ProtoMessage()    {}

const Default_JoinGameRequest_AdvancedNotification bool = false

func (m *JoinGameRequest) GetGameHandle() *game_master_types.GameHandle {
	if m != nil {
		return m.GameHandle
	}
	return nil
}

func (m *JoinGameRequest) GetPlayer() []*game_master_types.Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *JoinGameRequest) GetAdvancedNotification() bool {
	if m != nil && m.AdvancedNotification != nil {
		return *m.AdvancedNotification
	}
	return Default_JoinGameRequest_AdvancedNotification
}

type JoinGameResponse struct {
	RequestId        *uint64                          `protobuf:"fixed64,1,opt,name=request_id" json:"request_id,omitempty"`
	Queued           *bool                            `protobuf:"varint,2,opt,name=queued,def=0" json:"queued,omitempty"`
	ConnectInfo      []*game_master_types.ConnectInfo `protobuf:"bytes,3,rep,name=connect_info" json:"connect_info,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *JoinGameResponse) Reset()         { *m = JoinGameResponse{} }
func (m *JoinGameResponse) String() string { return proto.CompactTextString(m) }
func (*JoinGameResponse) ProtoMessage()    {}

const Default_JoinGameResponse_Queued bool = false

func (m *JoinGameResponse) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *JoinGameResponse) GetQueued() bool {
	if m != nil && m.Queued != nil {
		return *m.Queued
	}
	return Default_JoinGameResponse_Queued
}

func (m *JoinGameResponse) GetConnectInfo() []*game_master_types.ConnectInfo {
	if m != nil {
		return m.ConnectInfo
	}
	return nil
}

type ListFactoriesRequest struct {
	Filter           *attribute.AttributeFilter `protobuf:"bytes,1,req,name=filter" json:"filter,omitempty"`
	StartIndex       *uint32                    `protobuf:"varint,2,opt,name=start_index,def=0" json:"start_index,omitempty"`
	MaxResults       *uint32                    `protobuf:"varint,3,opt,name=max_results,def=100" json:"max_results,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ListFactoriesRequest) Reset()         { *m = ListFactoriesRequest{} }
func (m *ListFactoriesRequest) String() string { return proto.CompactTextString(m) }
func (*ListFactoriesRequest) ProtoMessage()    {}

const Default_ListFactoriesRequest_StartIndex uint32 = 0
const Default_ListFactoriesRequest_MaxResults uint32 = 100

func (m *ListFactoriesRequest) GetFilter() *attribute.AttributeFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ListFactoriesRequest) GetStartIndex() uint32 {
	if m != nil && m.StartIndex != nil {
		return *m.StartIndex
	}
	return Default_ListFactoriesRequest_StartIndex
}

func (m *ListFactoriesRequest) GetMaxResults() uint32 {
	if m != nil && m.MaxResults != nil {
		return *m.MaxResults
	}
	return Default_ListFactoriesRequest_MaxResults
}

type ListFactoriesResponse struct {
	Description      []*game_master_types.GameFactoryDescription `protobuf:"bytes,1,rep,name=description" json:"description,omitempty"`
	TotalResults     *uint32                                     `protobuf:"varint,2,opt,name=total_results" json:"total_results,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *ListFactoriesResponse) Reset()         { *m = ListFactoriesResponse{} }
func (m *ListFactoriesResponse) String() string { return proto.CompactTextString(m) }
func (*ListFactoriesResponse) ProtoMessage()    {}

func (m *ListFactoriesResponse) GetDescription() []*game_master_types.GameFactoryDescription {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *ListFactoriesResponse) GetTotalResults() uint32 {
	if m != nil && m.TotalResults != nil {
		return *m.TotalResults
	}
	return 0
}

type FindGameRequest struct {
	Player               []*game_master_types.Player  `protobuf:"bytes,1,rep,name=player" json:"player,omitempty"`
	FactoryId            *uint64                      `protobuf:"fixed64,2,opt,name=factory_id" json:"factory_id,omitempty"`
	Properties           *game_factory.GameProperties `protobuf:"bytes,3,opt,name=properties" json:"properties,omitempty"`
	ObjectId             *uint64                      `protobuf:"varint,4,opt,name=object_id" json:"object_id,omitempty"`
	RequestId            *uint64                      `protobuf:"fixed64,5,opt,name=request_id" json:"request_id,omitempty"`
	AdvancedNotification *bool                        `protobuf:"varint,6,opt,name=advanced_notification,def=0" json:"advanced_notification,omitempty"`
	XXX_unrecognized     []byte                       `json:"-"`
}

func (m *FindGameRequest) Reset()         { *m = FindGameRequest{} }
func (m *FindGameRequest) String() string { return proto.CompactTextString(m) }
func (*FindGameRequest) ProtoMessage()    {}

const Default_FindGameRequest_AdvancedNotification bool = false

func (m *FindGameRequest) GetPlayer() []*game_master_types.Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *FindGameRequest) GetFactoryId() uint64 {
	if m != nil && m.FactoryId != nil {
		return *m.FactoryId
	}
	return 0
}

func (m *FindGameRequest) GetProperties() *game_factory.GameProperties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *FindGameRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *FindGameRequest) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *FindGameRequest) GetAdvancedNotification() bool {
	if m != nil && m.AdvancedNotification != nil {
		return *m.AdvancedNotification
	}
	return Default_FindGameRequest_AdvancedNotification
}

type FindGameResponse struct {
	RequestId        *uint64 `protobuf:"fixed64,1,opt,name=request_id" json:"request_id,omitempty"`
	FactoryId        *uint64 `protobuf:"fixed64,2,opt,name=factory_id" json:"factory_id,omitempty"`
	Queued           *bool   `protobuf:"varint,3,opt,name=queued,def=0" json:"queued,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FindGameResponse) Reset()         { *m = FindGameResponse{} }
func (m *FindGameResponse) String() string { return proto.CompactTextString(m) }
func (*FindGameResponse) ProtoMessage()    {}

const Default_FindGameResponse_Queued bool = false

func (m *FindGameResponse) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *FindGameResponse) GetFactoryId() uint64 {
	if m != nil && m.FactoryId != nil {
		return *m.FactoryId
	}
	return 0
}

func (m *FindGameResponse) GetQueued() bool {
	if m != nil && m.Queued != nil {
		return *m.Queued
	}
	return Default_FindGameResponse_Queued
}

type GameEndedNotification struct {
	GameHandle       *game_master_types.GameHandle `protobuf:"bytes,1,req,name=game_handle" json:"game_handle,omitempty"`
	Reason           *uint32                       `protobuf:"varint,2,opt,name=reason,def=0" json:"reason,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *GameEndedNotification) Reset()         { *m = GameEndedNotification{} }
func (m *GameEndedNotification) String() string { return proto.CompactTextString(m) }
func (*GameEndedNotification) ProtoMessage()    {}

const Default_GameEndedNotification_Reason uint32 = 0

func (m *GameEndedNotification) GetGameHandle() *game_master_types.GameHandle {
	if m != nil {
		return m.GameHandle
	}
	return nil
}

func (m *GameEndedNotification) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return Default_GameEndedNotification_Reason
}

type PlayerLeftNotification struct {
	GameHandle       *game_master_types.GameHandle `protobuf:"bytes,1,req,name=game_handle" json:"game_handle,omitempty"`
	MemberId         *entity.EntityId              `protobuf:"bytes,2,req,name=member_id" json:"member_id,omitempty"`
	Reason           *uint32                       `protobuf:"varint,3,opt,name=reason,def=1" json:"reason,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *PlayerLeftNotification) Reset()         { *m = PlayerLeftNotification{} }
func (m *PlayerLeftNotification) String() string { return proto.CompactTextString(m) }
func (*PlayerLeftNotification) ProtoMessage()    {}

const Default_PlayerLeftNotification_Reason uint32 = 1

func (m *PlayerLeftNotification) GetGameHandle() *game_master_types.GameHandle {
	if m != nil {
		return m.GameHandle
	}
	return nil
}

func (m *PlayerLeftNotification) GetMemberId() *entity.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *PlayerLeftNotification) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return Default_PlayerLeftNotification_Reason
}

type RegisterServerRequest struct {
	Attribute        []*attribute.Attribute         `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	State            *server_pool_types.ServerState `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	ProgramId        *uint32                        `protobuf:"fixed32,3,req,name=program_id" json:"program_id,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *RegisterServerRequest) Reset()         { *m = RegisterServerRequest{} }
func (m *RegisterServerRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterServerRequest) ProtoMessage()    {}

func (m *RegisterServerRequest) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *RegisterServerRequest) GetState() *server_pool_types.ServerState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *RegisterServerRequest) GetProgramId() uint32 {
	if m != nil && m.ProgramId != nil {
		return *m.ProgramId
	}
	return 0
}

type UnregisterServerRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *UnregisterServerRequest) Reset()         { *m = UnregisterServerRequest{} }
func (m *UnregisterServerRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterServerRequest) ProtoMessage()    {}

type RegisterUtilitiesRequest struct {
	Attribute        []*attribute.Attribute         `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	State            *server_pool_types.ServerState `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	ProgramId        *uint32                        `protobuf:"fixed32,3,req,name=program_id" json:"program_id,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *RegisterUtilitiesRequest) Reset()         { *m = RegisterUtilitiesRequest{} }
func (m *RegisterUtilitiesRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterUtilitiesRequest) ProtoMessage()    {}

func (m *RegisterUtilitiesRequest) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *RegisterUtilitiesRequest) GetState() *server_pool_types.ServerState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *RegisterUtilitiesRequest) GetProgramId() uint32 {
	if m != nil && m.ProgramId != nil {
		return *m.ProgramId
	}
	return 0
}

type UnregisterUtilitiesRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *UnregisterUtilitiesRequest) Reset()         { *m = UnregisterUtilitiesRequest{} }
func (m *UnregisterUtilitiesRequest) String() string { return proto.CompactTextString(m) }
func (*UnregisterUtilitiesRequest) ProtoMessage()    {}

type SubscribeRequest struct {
	ObjectId         *uint64 `protobuf:"varint,1,req,name=object_id" json:"object_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}

func (m *SubscribeRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

type SubscribeResponse struct {
	SubscriptionId   *uint64 `protobuf:"varint,1,opt,name=subscription_id" json:"subscription_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}

func (m *SubscribeResponse) GetSubscriptionId() uint64 {
	if m != nil && m.SubscriptionId != nil {
		return *m.SubscriptionId
	}
	return 0
}

type UnsubscribeRequest struct {
	SubscriptionId   *uint64 `protobuf:"varint,1,req,name=subscription_id" json:"subscription_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UnsubscribeRequest) Reset()         { *m = UnsubscribeRequest{} }
func (m *UnsubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeRequest) ProtoMessage()    {}

func (m *UnsubscribeRequest) GetSubscriptionId() uint64 {
	if m != nil && m.SubscriptionId != nil {
		return *m.SubscriptionId
	}
	return 0
}

type ChangeGameRequest struct {
	GameHandle       *game_master_types.GameHandle `protobuf:"bytes,1,req,name=game_handle" json:"game_handle,omitempty"`
	Open             *bool                         `protobuf:"varint,2,opt,name=open" json:"open,omitempty"`
	Attribute        []*attribute.Attribute        `protobuf:"bytes,3,rep,name=attribute" json:"attribute,omitempty"`
	Replace          *bool                         `protobuf:"varint,4,opt,name=replace,def=0" json:"replace,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *ChangeGameRequest) Reset()         { *m = ChangeGameRequest{} }
func (m *ChangeGameRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeGameRequest) ProtoMessage()    {}

const Default_ChangeGameRequest_Replace bool = false

func (m *ChangeGameRequest) GetGameHandle() *game_master_types.GameHandle {
	if m != nil {
		return m.GameHandle
	}
	return nil
}

func (m *ChangeGameRequest) GetOpen() bool {
	if m != nil && m.Open != nil {
		return *m.Open
	}
	return false
}

func (m *ChangeGameRequest) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *ChangeGameRequest) GetReplace() bool {
	if m != nil && m.Replace != nil {
		return *m.Replace
	}
	return Default_ChangeGameRequest_Replace
}

type GetFactoryInfoRequest struct {
	FactoryId        *uint64 `protobuf:"fixed64,1,req,name=factory_id" json:"factory_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetFactoryInfoRequest) Reset()         { *m = GetFactoryInfoRequest{} }
func (m *GetFactoryInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetFactoryInfoRequest) ProtoMessage()    {}

func (m *GetFactoryInfoRequest) GetFactoryId() uint64 {
	if m != nil && m.FactoryId != nil {
		return *m.FactoryId
	}
	return 0
}

type GetFactoryInfoResponse struct {
	Attribute        []*attribute.Attribute               `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	StatsBucket      []*game_master_types.GameStatsBucket `protobuf:"bytes,2,rep,name=stats_bucket" json:"stats_bucket,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *GetFactoryInfoResponse) Reset()         { *m = GetFactoryInfoResponse{} }
func (m *GetFactoryInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetFactoryInfoResponse) ProtoMessage()    {}

func (m *GetFactoryInfoResponse) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *GetFactoryInfoResponse) GetStatsBucket() []*game_master_types.GameStatsBucket {
	if m != nil {
		return m.StatsBucket
	}
	return nil
}

type GetGameStatsRequest struct {
	FactoryId        *uint64                    `protobuf:"fixed64,1,req,name=factory_id" json:"factory_id,omitempty"`
	Filter           *attribute.AttributeFilter `protobuf:"bytes,2,req,name=filter" json:"filter,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *GetGameStatsRequest) Reset()         { *m = GetGameStatsRequest{} }
func (m *GetGameStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetGameStatsRequest) ProtoMessage()    {}

func (m *GetGameStatsRequest) GetFactoryId() uint64 {
	if m != nil && m.FactoryId != nil {
		return *m.FactoryId
	}
	return 0
}

func (m *GetGameStatsRequest) GetFilter() *attribute.AttributeFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type GetGameStatsResponse struct {
	StatsBucket      []*game_master_types.GameStatsBucket `protobuf:"bytes,1,rep,name=stats_bucket" json:"stats_bucket,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *GetGameStatsResponse) Reset()         { *m = GetGameStatsResponse{} }
func (m *GetGameStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetGameStatsResponse) ProtoMessage()    {}

func (m *GetGameStatsResponse) GetStatsBucket() []*game_master_types.GameStatsBucket {
	if m != nil {
		return m.StatsBucket
	}
	return nil
}

type FactoryUpdateNotification struct {
	Op               *FactoryUpdateNotification_Operation      `protobuf:"varint,1,req,name=op,enum=game_master_service.FactoryUpdateNotification_Operation" json:"op,omitempty"`
	Description      *game_master_types.GameFactoryDescription `protobuf:"bytes,2,req,name=description" json:"description,omitempty"`
	ProgramId        *uint32                                   `protobuf:"fixed32,3,opt,name=program_id" json:"program_id,omitempty"`
	XXX_unrecognized []byte                                    `json:"-"`
}

func (m *FactoryUpdateNotification) Reset()         { *m = FactoryUpdateNotification{} }
func (m *FactoryUpdateNotification) String() string { return proto.CompactTextString(m) }
func (*FactoryUpdateNotification) ProtoMessage()    {}

func (m *FactoryUpdateNotification) GetOp() FactoryUpdateNotification_Operation {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return FactoryUpdateNotification_ADD
}

func (m *FactoryUpdateNotification) GetDescription() *game_master_types.GameFactoryDescription {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *FactoryUpdateNotification) GetProgramId() uint32 {
	if m != nil && m.ProgramId != nil {
		return *m.ProgramId
	}
	return 0
}

type GameFoundNotification struct {
	RequestId        *uint64                          `protobuf:"fixed64,1,req,name=request_id" json:"request_id,omitempty"`
	ErrorCode        *uint32                          `protobuf:"varint,2,opt,name=error_code,def=0" json:"error_code,omitempty"`
	GameHandle       *game_master_types.GameHandle    `protobuf:"bytes,3,opt,name=game_handle" json:"game_handle,omitempty"`
	ConnectInfo      []*game_master_types.ConnectInfo `protobuf:"bytes,4,rep,name=connect_info" json:"connect_info,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *GameFoundNotification) Reset()         { *m = GameFoundNotification{} }
func (m *GameFoundNotification) String() string { return proto.CompactTextString(m) }
func (*GameFoundNotification) ProtoMessage()    {}

const Default_GameFoundNotification_ErrorCode uint32 = 0

func (m *GameFoundNotification) GetRequestId() uint64 {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return 0
}

func (m *GameFoundNotification) GetErrorCode() uint32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return Default_GameFoundNotification_ErrorCode
}

func (m *GameFoundNotification) GetGameHandle() *game_master_types.GameHandle {
	if m != nil {
		return m.GameHandle
	}
	return nil
}

func (m *GameFoundNotification) GetConnectInfo() []*game_master_types.ConnectInfo {
	if m != nil {
		return m.ConnectInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("game_master_service.FactoryUpdateNotification_Operation", FactoryUpdateNotification_Operation_name, FactoryUpdateNotification_Operation_value)
}
