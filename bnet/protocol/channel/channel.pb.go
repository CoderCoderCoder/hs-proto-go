// Code generated by protoc-gen-go.
// source: bnet/protocol/channel/channel.proto
// DO NOT EDIT!

package bnet_protocol_channel

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"
import bnet_protocol_attribute_468 "github.com/HearthSim/hs-proto-go/bnet/protocol/attribute_468"
import bnet_protocol_channel_extracted "github.com/HearthSim/hs-proto-go/bnet/protocol/channel_extracted"
import bnet_protocol "github.com/HearthSim/hs-proto-go/bnet/protocol"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// ref: bnet.protocol.channel.AddMemberRequest
type AddMemberRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	MemberIdentity   *bnet_protocol.Identity `protobuf:"bytes,2,req,name=member_identity" json:"member_identity,omitempty"`
	MemberState      *MemberState            `protobuf:"bytes,3,req,name=member_state" json:"member_state,omitempty"`
	ObjectId         *uint64                 `protobuf:"varint,4,req,name=object_id" json:"object_id,omitempty"`
	Subscribe        *bool                   `protobuf:"varint,5,opt,name=subscribe,def=1" json:"subscribe,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *AddMemberRequest) Reset()         { *m = AddMemberRequest{} }
func (m *AddMemberRequest) String() string { return proto.CompactTextString(m) }
func (*AddMemberRequest) ProtoMessage()    {}

const Default_AddMemberRequest_Subscribe bool = true

func (m *AddMemberRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *AddMemberRequest) GetMemberIdentity() *bnet_protocol.Identity {
	if m != nil {
		return m.MemberIdentity
	}
	return nil
}

func (m *AddMemberRequest) GetMemberState() *MemberState {
	if m != nil {
		return m.MemberState
	}
	return nil
}

func (m *AddMemberRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *AddMemberRequest) GetSubscribe() bool {
	if m != nil && m.Subscribe != nil {
		return *m.Subscribe
	}
	return Default_AddMemberRequest_Subscribe
}

// ref: bnet.protocol.channel.AddNotification
type AddNotification struct {
	Self             *Member                                       `protobuf:"bytes,1,opt,name=self" json:"self,omitempty"`
	Member           []*Member                                     `protobuf:"bytes,2,rep,name=member" json:"member,omitempty"`
	ChannelState     *bnet_protocol_channel_extracted.ChannelState `protobuf:"bytes,3,req,name=channel_state" json:"channel_state,omitempty"`
	XXX_unrecognized []byte                                        `json:"-"`
}

func (m *AddNotification) Reset()         { *m = AddNotification{} }
func (m *AddNotification) String() string { return proto.CompactTextString(m) }
func (*AddNotification) ProtoMessage()    {}

func (m *AddNotification) GetSelf() *Member {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *AddNotification) GetMember() []*Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *AddNotification) GetChannelState() *bnet_protocol_channel_extracted.ChannelState {
	if m != nil {
		return m.ChannelState
	}
	return nil
}

// ref: bnet.protocol.channel.ChannelInfo
type ChannelInfo struct {
	Description      *bnet_protocol_channel_extracted.ChannelDescription `protobuf:"bytes,1,req,name=description" json:"description,omitempty"`
	Member           []*Member                                           `protobuf:"bytes,2,rep,name=member" json:"member,omitempty"`
	XXX_unrecognized []byte                                              `json:"-"`
}

func (m *ChannelInfo) Reset()         { *m = ChannelInfo{} }
func (m *ChannelInfo) String() string { return proto.CompactTextString(m) }
func (*ChannelInfo) ProtoMessage()    {}

func (m *ChannelInfo) GetDescription() *bnet_protocol_channel_extracted.ChannelDescription {
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

// ref: bnet.protocol.channel.CreateChannelRequest
type CreateChannelRequest struct {
	AgentIdentity    *bnet_protocol.Identity                       `protobuf:"bytes,1,opt,name=agent_identity" json:"agent_identity,omitempty"`
	MemberState      *MemberState                                  `protobuf:"bytes,2,opt,name=member_state" json:"member_state,omitempty"`
	ChannelState     *bnet_protocol_channel_extracted.ChannelState `protobuf:"bytes,3,opt,name=channel_state" json:"channel_state,omitempty"`
	ChannelId        *bnet_protocol.EntityId                       `protobuf:"bytes,4,opt,name=channel_id" json:"channel_id,omitempty"`
	ObjectId         *uint64                                       `protobuf:"varint,5,opt,name=object_id" json:"object_id,omitempty"`
	LocalAgent       *bnet_protocol.EntityId                       `protobuf:"bytes,6,opt,name=local_agent" json:"local_agent,omitempty"`
	LocalMemberState *MemberState                                  `protobuf:"bytes,7,opt,name=local_member_state" json:"local_member_state,omitempty"`
	XXX_unrecognized []byte                                        `json:"-"`
}

func (m *CreateChannelRequest) Reset()         { *m = CreateChannelRequest{} }
func (m *CreateChannelRequest) String() string { return proto.CompactTextString(m) }
func (*CreateChannelRequest) ProtoMessage()    {}

func (m *CreateChannelRequest) GetAgentIdentity() *bnet_protocol.Identity {
	if m != nil {
		return m.AgentIdentity
	}
	return nil
}

func (m *CreateChannelRequest) GetMemberState() *MemberState {
	if m != nil {
		return m.MemberState
	}
	return nil
}

func (m *CreateChannelRequest) GetChannelState() *bnet_protocol_channel_extracted.ChannelState {
	if m != nil {
		return m.ChannelState
	}
	return nil
}

func (m *CreateChannelRequest) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *CreateChannelRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *CreateChannelRequest) GetLocalAgent() *bnet_protocol.EntityId {
	if m != nil {
		return m.LocalAgent
	}
	return nil
}

func (m *CreateChannelRequest) GetLocalMemberState() *MemberState {
	if m != nil {
		return m.LocalMemberState
	}
	return nil
}

// ref: bnet.protocol.channel.CreateChannelResponse
type CreateChannelResponse struct {
	ObjectId         *uint64                 `protobuf:"varint,1,req,name=object_id" json:"object_id,omitempty"`
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,2,opt,name=channel_id" json:"channel_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *CreateChannelResponse) Reset()         { *m = CreateChannelResponse{} }
func (m *CreateChannelResponse) String() string { return proto.CompactTextString(m) }
func (*CreateChannelResponse) ProtoMessage()    {}

func (m *CreateChannelResponse) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *CreateChannelResponse) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

// ref: bnet.protocol.channel.DissolveRequest
type DissolveRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	Reason           *uint32                 `protobuf:"varint,2,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *DissolveRequest) Reset()         { *m = DissolveRequest{} }
func (m *DissolveRequest) String() string { return proto.CompactTextString(m) }
func (*DissolveRequest) ProtoMessage()    {}

func (m *DissolveRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *DissolveRequest) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

// ref: bnet.protocol.channel.FindChannelOptions
type FindChannelOptions struct {
	StartIndex       *uint32                                      `protobuf:"varint,1,opt,name=start_index,def=0" json:"start_index,omitempty"`
	MaxResults       *uint32                                      `protobuf:"varint,2,opt,name=max_results,def=16" json:"max_results,omitempty"`
	Name             *string                                      `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Program          *uint32                                      `protobuf:"fixed32,4,opt,name=program" json:"program,omitempty"`
	Locale           *uint32                                      `protobuf:"fixed32,5,opt,name=locale" json:"locale,omitempty"`
	CapacityFull     *uint32                                      `protobuf:"varint,6,opt,name=capacity_full" json:"capacity_full,omitempty"`
	AttributeFilter  *bnet_protocol_attribute_468.AttributeFilter `protobuf:"bytes,7,req,name=attribute_filter" json:"attribute_filter,omitempty"`
	ChannelType      *string                                      `protobuf:"bytes,8,opt,name=channel_type" json:"channel_type,omitempty"`
	XXX_unrecognized []byte                                       `json:"-"`
}

func (m *FindChannelOptions) Reset()         { *m = FindChannelOptions{} }
func (m *FindChannelOptions) String() string { return proto.CompactTextString(m) }
func (*FindChannelOptions) ProtoMessage()    {}

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

func (m *FindChannelOptions) GetAttributeFilter() *bnet_protocol_attribute_468.AttributeFilter {
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

// ref: bnet.protocol.channel.FindChannelRequest
type FindChannelRequest struct {
	AgentIdentity    *bnet_protocol.Identity `protobuf:"bytes,1,opt,name=agent_identity" json:"agent_identity,omitempty"`
	Options          *FindChannelOptions     `protobuf:"bytes,2,req,name=options" json:"options,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *FindChannelRequest) Reset()         { *m = FindChannelRequest{} }
func (m *FindChannelRequest) String() string { return proto.CompactTextString(m) }
func (*FindChannelRequest) ProtoMessage()    {}

func (m *FindChannelRequest) GetAgentIdentity() *bnet_protocol.Identity {
	if m != nil {
		return m.AgentIdentity
	}
	return nil
}

func (m *FindChannelRequest) GetOptions() *FindChannelOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// ref: bnet.protocol.channel.FindChannelResponse
type FindChannelResponse struct {
	Channel          []*bnet_protocol_channel_extracted.ChannelDescription `protobuf:"bytes,1,rep,name=channel" json:"channel,omitempty"`
	XXX_unrecognized []byte                                                `json:"-"`
}

func (m *FindChannelResponse) Reset()         { *m = FindChannelResponse{} }
func (m *FindChannelResponse) String() string { return proto.CompactTextString(m) }
func (*FindChannelResponse) ProtoMessage()    {}

func (m *FindChannelResponse) GetChannel() []*bnet_protocol_channel_extracted.ChannelDescription {
	if m != nil {
		return m.Channel
	}
	return nil
}

// ref: bnet.protocol.channel.GetChannelIdRequest
type GetChannelIdRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetChannelIdRequest) Reset()         { *m = GetChannelIdRequest{} }
func (m *GetChannelIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetChannelIdRequest) ProtoMessage()    {}

// ref: bnet.protocol.channel.GetChannelIdResponse
type GetChannelIdResponse struct {
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=channel_id" json:"channel_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *GetChannelIdResponse) Reset()         { *m = GetChannelIdResponse{} }
func (m *GetChannelIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetChannelIdResponse) ProtoMessage()    {}

func (m *GetChannelIdResponse) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

// ref: bnet.protocol.channel.GetChannelInfoRequest
type GetChannelInfoRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=channel_id" json:"channel_id,omitempty"`
	FetchState       *bool                   `protobuf:"varint,3,opt,name=fetch_state,def=0" json:"fetch_state,omitempty"`
	FetchMembers     *bool                   `protobuf:"varint,4,opt,name=fetch_members,def=0" json:"fetch_members,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *GetChannelInfoRequest) Reset()         { *m = GetChannelInfoRequest{} }
func (m *GetChannelInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetChannelInfoRequest) ProtoMessage()    {}

const Default_GetChannelInfoRequest_FetchState bool = false
const Default_GetChannelInfoRequest_FetchMembers bool = false

func (m *GetChannelInfoRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *GetChannelInfoRequest) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *GetChannelInfoRequest) GetFetchState() bool {
	if m != nil && m.FetchState != nil {
		return *m.FetchState
	}
	return Default_GetChannelInfoRequest_FetchState
}

func (m *GetChannelInfoRequest) GetFetchMembers() bool {
	if m != nil && m.FetchMembers != nil {
		return *m.FetchMembers
	}
	return Default_GetChannelInfoRequest_FetchMembers
}

// ref: bnet.protocol.channel.GetChannelInfoResponse
type GetChannelInfoResponse struct {
	ChannelInfo      *ChannelInfo `protobuf:"bytes,1,opt,name=channel_info" json:"channel_info,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GetChannelInfoResponse) Reset()         { *m = GetChannelInfoResponse{} }
func (m *GetChannelInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetChannelInfoResponse) ProtoMessage()    {}

func (m *GetChannelInfoResponse) GetChannelInfo() *ChannelInfo {
	if m != nil {
		return m.ChannelInfo
	}
	return nil
}

// ref: bnet.protocol.channel.JoinChannelRequest
type JoinChannelRequest struct {
	AgentIdentity    *bnet_protocol.Identity   `protobuf:"bytes,1,opt,name=agent_identity" json:"agent_identity,omitempty"`
	MemberState      *MemberState              `protobuf:"bytes,2,opt,name=member_state" json:"member_state,omitempty"`
	ChannelId        *bnet_protocol.EntityId   `protobuf:"bytes,3,req,name=channel_id" json:"channel_id,omitempty"`
	ObjectId         *uint64                   `protobuf:"varint,4,req,name=object_id" json:"object_id,omitempty"`
	FriendAccountId  []*bnet_protocol.EntityId `protobuf:"bytes,5,rep,name=friend_account_id" json:"friend_account_id,omitempty"`
	LocalSubscriber  *bool                     `protobuf:"varint,6,opt,name=local_subscriber,def=1" json:"local_subscriber,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *JoinChannelRequest) Reset()         { *m = JoinChannelRequest{} }
func (m *JoinChannelRequest) String() string { return proto.CompactTextString(m) }
func (*JoinChannelRequest) ProtoMessage()    {}

const Default_JoinChannelRequest_LocalSubscriber bool = true

func (m *JoinChannelRequest) GetAgentIdentity() *bnet_protocol.Identity {
	if m != nil {
		return m.AgentIdentity
	}
	return nil
}

func (m *JoinChannelRequest) GetMemberState() *MemberState {
	if m != nil {
		return m.MemberState
	}
	return nil
}

func (m *JoinChannelRequest) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *JoinChannelRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *JoinChannelRequest) GetFriendAccountId() []*bnet_protocol.EntityId {
	if m != nil {
		return m.FriendAccountId
	}
	return nil
}

func (m *JoinChannelRequest) GetLocalSubscriber() bool {
	if m != nil && m.LocalSubscriber != nil {
		return *m.LocalSubscriber
	}
	return Default_JoinChannelRequest_LocalSubscriber
}

// ref: bnet.protocol.channel.JoinChannelResponse
type JoinChannelResponse struct {
	ObjectId                *uint64                   `protobuf:"varint,1,opt,name=object_id" json:"object_id,omitempty"`
	RequireFriendValidation *bool                     `protobuf:"varint,2,opt,name=require_friend_validation,def=0" json:"require_friend_validation,omitempty"`
	PrivilegedAccount       []*bnet_protocol.EntityId `protobuf:"bytes,3,rep,name=privileged_account" json:"privileged_account,omitempty"`
	XXX_unrecognized        []byte                    `json:"-"`
}

func (m *JoinChannelResponse) Reset()         { *m = JoinChannelResponse{} }
func (m *JoinChannelResponse) String() string { return proto.CompactTextString(m) }
func (*JoinChannelResponse) ProtoMessage()    {}

const Default_JoinChannelResponse_RequireFriendValidation bool = false

func (m *JoinChannelResponse) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *JoinChannelResponse) GetRequireFriendValidation() bool {
	if m != nil && m.RequireFriendValidation != nil {
		return *m.RequireFriendValidation
	}
	return Default_JoinChannelResponse_RequireFriendValidation
}

func (m *JoinChannelResponse) GetPrivilegedAccount() []*bnet_protocol.EntityId {
	if m != nil {
		return m.PrivilegedAccount
	}
	return nil
}

// ref: bnet.protocol.channel.JoinNotification
type JoinNotification struct {
	Member           *Member `protobuf:"bytes,1,req,name=member" json:"member,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *JoinNotification) Reset()         { *m = JoinNotification{} }
func (m *JoinNotification) String() string { return proto.CompactTextString(m) }
func (*JoinNotification) ProtoMessage()    {}

func (m *JoinNotification) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

// ref: bnet.protocol.channel.LeaveNotification
type LeaveNotification struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	MemberId         *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=member_id" json:"member_id,omitempty"`
	Reason           *uint32                 `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *LeaveNotification) Reset()         { *m = LeaveNotification{} }
func (m *LeaveNotification) String() string { return proto.CompactTextString(m) }
func (*LeaveNotification) ProtoMessage()    {}

func (m *LeaveNotification) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *LeaveNotification) GetMemberId() *bnet_protocol.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *LeaveNotification) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

// ref: bnet.protocol.channel.Member
type Member struct {
	Identity         *bnet_protocol.Identity `protobuf:"bytes,1,req,name=identity" json:"identity,omitempty"`
	State            *MemberState            `protobuf:"bytes,2,req,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}

func (m *Member) GetIdentity() *bnet_protocol.Identity {
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

// ref: bnet.protocol.channel.MemberState
type MemberState struct {
	Attribute        []*bnet_protocol.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	Role             []uint32                   `protobuf:"varint,2,rep,packed,name=role" json:"role,omitempty"`
	Privileges       *uint64                    `protobuf:"varint,3,opt,name=privileges,def=0" json:"privileges,omitempty"`
	Info             *bnet_protocol.AccountInfo `protobuf:"bytes,4,opt,name=info" json:"info,omitempty"`
	Hidden           *bool                      `protobuf:"varint,5,opt,name=hidden,def=0" json:"hidden,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *MemberState) Reset()         { *m = MemberState{} }
func (m *MemberState) String() string { return proto.CompactTextString(m) }
func (*MemberState) ProtoMessage()    {}

const Default_MemberState_Privileges uint64 = 0
const Default_MemberState_Hidden bool = false

func (m *MemberState) GetAttribute() []*bnet_protocol.Attribute {
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

func (m *MemberState) GetInfo() *bnet_protocol.AccountInfo {
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

// ref: bnet.protocol.channel.Message
type Message struct {
	Attribute        []*bnet_protocol.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	Role             *uint32                    `protobuf:"varint,2,opt,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func (m *Message) GetAttribute() []*bnet_protocol.Attribute {
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

// ref: bnet.protocol.channel.RemoveMemberRequest
type RemoveMemberRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	MemberId         *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=member_id" json:"member_id,omitempty"`
	Reason           *uint32                 `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *RemoveMemberRequest) Reset()         { *m = RemoveMemberRequest{} }
func (m *RemoveMemberRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveMemberRequest) ProtoMessage()    {}

func (m *RemoveMemberRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *RemoveMemberRequest) GetMemberId() *bnet_protocol.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *RemoveMemberRequest) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

// ref: bnet.protocol.channel.RemoveNotification
type RemoveNotification struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	MemberId         *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=member_id" json:"member_id,omitempty"`
	Reason           *uint32                 `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *RemoveNotification) Reset()         { *m = RemoveNotification{} }
func (m *RemoveNotification) String() string { return proto.CompactTextString(m) }
func (*RemoveNotification) ProtoMessage()    {}

func (m *RemoveNotification) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *RemoveNotification) GetMemberId() *bnet_protocol.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *RemoveNotification) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

// ref: bnet.protocol.channel.SendMessageNotification
type SendMessageNotification struct {
	AgentId            *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	Message            *Message                `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	RequiredPrivileges *uint64                 `protobuf:"varint,3,opt,name=required_privileges,def=0" json:"required_privileges,omitempty"`
	XXX_unrecognized   []byte                  `json:"-"`
}

func (m *SendMessageNotification) Reset()         { *m = SendMessageNotification{} }
func (m *SendMessageNotification) String() string { return proto.CompactTextString(m) }
func (*SendMessageNotification) ProtoMessage()    {}

const Default_SendMessageNotification_RequiredPrivileges uint64 = 0

func (m *SendMessageNotification) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SendMessageNotification) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SendMessageNotification) GetRequiredPrivileges() uint64 {
	if m != nil && m.RequiredPrivileges != nil {
		return *m.RequiredPrivileges
	}
	return Default_SendMessageNotification_RequiredPrivileges
}

// ref: bnet.protocol.channel.SendMessageRequest
type SendMessageRequest struct {
	AgentId            *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	Message            *Message                `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	RequiredPrivileges *uint64                 `protobuf:"varint,3,opt,name=required_privileges,def=0" json:"required_privileges,omitempty"`
	XXX_unrecognized   []byte                  `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}

const Default_SendMessageRequest_RequiredPrivileges uint64 = 0

func (m *SendMessageRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SendMessageRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SendMessageRequest) GetRequiredPrivileges() uint64 {
	if m != nil && m.RequiredPrivileges != nil {
		return *m.RequiredPrivileges
	}
	return Default_SendMessageRequest_RequiredPrivileges
}

// ref: bnet.protocol.channel.SetRolesRequest
type SetRolesRequest struct {
	AgentId          *bnet_protocol.EntityId   `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	Role             []uint32                  `protobuf:"varint,2,rep,packed,name=role" json:"role,omitempty"`
	MemberId         []*bnet_protocol.EntityId `protobuf:"bytes,3,rep,name=member_id" json:"member_id,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *SetRolesRequest) Reset()         { *m = SetRolesRequest{} }
func (m *SetRolesRequest) String() string { return proto.CompactTextString(m) }
func (*SetRolesRequest) ProtoMessage()    {}

func (m *SetRolesRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SetRolesRequest) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *SetRolesRequest) GetMemberId() []*bnet_protocol.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

// ref: bnet.protocol.channel.SubscribeChannelRequest
type SubscribeChannelRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=channel_id" json:"channel_id,omitempty"`
	ObjectId         *uint64                 `protobuf:"varint,3,req,name=object_id" json:"object_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *SubscribeChannelRequest) Reset()         { *m = SubscribeChannelRequest{} }
func (m *SubscribeChannelRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeChannelRequest) ProtoMessage()    {}

func (m *SubscribeChannelRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SubscribeChannelRequest) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *SubscribeChannelRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

// ref: bnet.protocol.channel.SubscribeChannelResponse
type SubscribeChannelResponse struct {
	ObjectId         *uint64 `protobuf:"varint,1,opt,name=object_id" json:"object_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SubscribeChannelResponse) Reset()         { *m = SubscribeChannelResponse{} }
func (m *SubscribeChannelResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeChannelResponse) ProtoMessage()    {}

func (m *SubscribeChannelResponse) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

// ref: bnet.protocol.channel.UnsubscribeMemberRequest
type UnsubscribeMemberRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	MemberId         *bnet_protocol.EntityId `protobuf:"bytes,2,req,name=member_id" json:"member_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UnsubscribeMemberRequest) Reset()         { *m = UnsubscribeMemberRequest{} }
func (m *UnsubscribeMemberRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeMemberRequest) ProtoMessage()    {}

func (m *UnsubscribeMemberRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UnsubscribeMemberRequest) GetMemberId() *bnet_protocol.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

// ref: bnet.protocol.channel.UpdateChannelStateNotification
type UpdateChannelStateNotification struct {
	AgentId          *bnet_protocol.EntityId                       `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	StateChange      *bnet_protocol_channel_extracted.ChannelState `protobuf:"bytes,2,req,name=state_change" json:"state_change,omitempty"`
	XXX_unrecognized []byte                                        `json:"-"`
}

func (m *UpdateChannelStateNotification) Reset()         { *m = UpdateChannelStateNotification{} }
func (m *UpdateChannelStateNotification) String() string { return proto.CompactTextString(m) }
func (*UpdateChannelStateNotification) ProtoMessage()    {}

func (m *UpdateChannelStateNotification) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateChannelStateNotification) GetStateChange() *bnet_protocol_channel_extracted.ChannelState {
	if m != nil {
		return m.StateChange
	}
	return nil
}

// ref: bnet.protocol.channel.UpdateChannelStateRequest
type UpdateChannelStateRequest struct {
	AgentId          *bnet_protocol.EntityId                       `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	StateChange      *bnet_protocol_channel_extracted.ChannelState `protobuf:"bytes,2,req,name=state_change" json:"state_change,omitempty"`
	XXX_unrecognized []byte                                        `json:"-"`
}

func (m *UpdateChannelStateRequest) Reset()         { *m = UpdateChannelStateRequest{} }
func (m *UpdateChannelStateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateChannelStateRequest) ProtoMessage()    {}

func (m *UpdateChannelStateRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateChannelStateRequest) GetStateChange() *bnet_protocol_channel_extracted.ChannelState {
	if m != nil {
		return m.StateChange
	}
	return nil
}

// ref: bnet.protocol.channel.UpdateMemberStateNotification
type UpdateMemberStateNotification struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	StateChange      []*Member               `protobuf:"bytes,2,rep,name=state_change" json:"state_change,omitempty"`
	RemovedRole      []uint32                `protobuf:"varint,3,rep,packed,name=removed_role" json:"removed_role,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdateMemberStateNotification) Reset()         { *m = UpdateMemberStateNotification{} }
func (m *UpdateMemberStateNotification) String() string { return proto.CompactTextString(m) }
func (*UpdateMemberStateNotification) ProtoMessage()    {}

func (m *UpdateMemberStateNotification) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateMemberStateNotification) GetStateChange() []*Member {
	if m != nil {
		return m.StateChange
	}
	return nil
}

func (m *UpdateMemberStateNotification) GetRemovedRole() []uint32 {
	if m != nil {
		return m.RemovedRole
	}
	return nil
}

// ref: bnet.protocol.channel.UpdateMemberStateRequest
type UpdateMemberStateRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	StateChange      []*Member               `protobuf:"bytes,2,rep,name=state_change" json:"state_change,omitempty"`
	RemovedRole      []uint32                `protobuf:"varint,3,rep,packed,name=removed_role" json:"removed_role,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdateMemberStateRequest) Reset()         { *m = UpdateMemberStateRequest{} }
func (m *UpdateMemberStateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMemberStateRequest) ProtoMessage()    {}

func (m *UpdateMemberStateRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateMemberStateRequest) GetStateChange() []*Member {
	if m != nil {
		return m.StateChange
	}
	return nil
}

func (m *UpdateMemberStateRequest) GetRemovedRole() []uint32 {
	if m != nil {
		return m.RemovedRole
	}
	return nil
}

func init() {
}
