// Code generated by protoc-gen-go.
// source: bnet/protocol/channel_invitation/channel_invitation.proto
// DO NOT EDIT!

package bnet_protocol_channel_invitation

import proto "github.com/golang/protobuf/proto"
import json "encoding/json"
import math "math"
import bnet_protocol_channel "github.com/HearthSim/hs-proto-go/bnet/protocol/channel"
import bnet_protocol_channel_extracted "github.com/HearthSim/hs-proto-go/bnet/protocol/channel_extracted"
import bnet_protocol_friends "github.com/HearthSim/hs-proto-go/bnet/protocol/friends"
import bnet_protocol_invitation "github.com/HearthSim/hs-proto-go/bnet/protocol/invitation"
import bnet_protocol "github.com/HearthSim/hs-proto-go/bnet/protocol"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// ref: bnet.protocol.channel_invitation.AcceptInvitationRequest
type AcceptInvitationRequest struct {
	AgentId          *bnet_protocol.EntityId            `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	MemberState      *bnet_protocol_channel.MemberState `protobuf:"bytes,2,opt,name=member_state" json:"member_state,omitempty"`
	InvitationId     *uint64                            `protobuf:"fixed64,3,req,name=invitation_id" json:"invitation_id,omitempty"`
	ObjectId         *uint64                            `protobuf:"varint,4,req,name=object_id" json:"object_id,omitempty"`
	ChannelId        *bnet_protocol.EntityId            `protobuf:"bytes,5,opt,name=channel_id" json:"channel_id,omitempty"`
	ServiceType      *uint32                            `protobuf:"varint,6,opt,name=service_type" json:"service_type,omitempty"`
	LocalSubscriber  *bool                              `protobuf:"varint,7,opt,name=local_subscriber,def=1" json:"local_subscriber,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *AcceptInvitationRequest) Reset()         { *m = AcceptInvitationRequest{} }
func (m *AcceptInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*AcceptInvitationRequest) ProtoMessage()    {}

const Default_AcceptInvitationRequest_LocalSubscriber bool = true

func (m *AcceptInvitationRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *AcceptInvitationRequest) GetMemberState() *bnet_protocol_channel.MemberState {
	if m != nil {
		return m.MemberState
	}
	return nil
}

func (m *AcceptInvitationRequest) GetInvitationId() uint64 {
	if m != nil && m.InvitationId != nil {
		return *m.InvitationId
	}
	return 0
}

func (m *AcceptInvitationRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *AcceptInvitationRequest) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *AcceptInvitationRequest) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *AcceptInvitationRequest) GetLocalSubscriber() bool {
	if m != nil && m.LocalSubscriber != nil {
		return *m.LocalSubscriber
	}
	return Default_AcceptInvitationRequest_LocalSubscriber
}

// ref: bnet.protocol.channel_invitation.AcceptInvitationResponse
type AcceptInvitationResponse struct {
	ObjectId         *uint64 `protobuf:"varint,1,req,name=object_id" json:"object_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AcceptInvitationResponse) Reset()         { *m = AcceptInvitationResponse{} }
func (m *AcceptInvitationResponse) String() string { return proto.CompactTextString(m) }
func (*AcceptInvitationResponse) ProtoMessage()    {}

func (m *AcceptInvitationResponse) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

// ref: bnet.protocol.channel_invitation.ChannelCount
type ChannelCount struct {
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=channel_id" json:"channel_id,omitempty"`
	ChannelType      *string                 `protobuf:"bytes,2,opt,name=channel_type,def=default" json:"channel_type,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ChannelCount) Reset()         { *m = ChannelCount{} }
func (m *ChannelCount) String() string { return proto.CompactTextString(m) }
func (*ChannelCount) ProtoMessage()    {}

const Default_ChannelCount_ChannelType string = "default"

func (m *ChannelCount) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *ChannelCount) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return Default_ChannelCount_ChannelType
}

// ref: bnet.protocol.channel_invitation.ChannelCountDescription
type ChannelCountDescription struct {
	ServiceType      *uint32                 `protobuf:"varint,1,req,name=service_type" json:"service_type,omitempty"`
	Program          *uint32                 `protobuf:"fixed32,2,req,name=program" json:"program,omitempty"`
	ChannelType      *string                 `protobuf:"bytes,3,opt,name=channel_type,def=default" json:"channel_type,omitempty"`
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,4,opt,name=channel_id" json:"channel_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ChannelCountDescription) Reset()         { *m = ChannelCountDescription{} }
func (m *ChannelCountDescription) String() string { return proto.CompactTextString(m) }
func (*ChannelCountDescription) ProtoMessage()    {}

const Default_ChannelCountDescription_ChannelType string = "default"

func (m *ChannelCountDescription) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *ChannelCountDescription) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *ChannelCountDescription) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return Default_ChannelCountDescription_ChannelType
}

func (m *ChannelCountDescription) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.ChannelInvitationParams
type ChannelInvitationParams struct {
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,1,req,name=channel_id" json:"channel_id,omitempty"`
	Reserved         *bool                   `protobuf:"varint,2,opt,name=reserved" json:"reserved,omitempty"`
	Rejoin           *bool                   `protobuf:"varint,3,opt,name=rejoin" json:"rejoin,omitempty"`
	ServiceType      *uint32                 `protobuf:"varint,4,req,name=service_type" json:"service_type,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ChannelInvitationParams) Reset()         { *m = ChannelInvitationParams{} }
func (m *ChannelInvitationParams) String() string { return proto.CompactTextString(m) }
func (*ChannelInvitationParams) ProtoMessage()    {}

func (m *ChannelInvitationParams) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *ChannelInvitationParams) GetReserved() bool {
	if m != nil && m.Reserved != nil {
		return *m.Reserved
	}
	return false
}

func (m *ChannelInvitationParams) GetRejoin() bool {
	if m != nil && m.Rejoin != nil {
		return *m.Rejoin
	}
	return false
}

func (m *ChannelInvitationParams) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

// ref: bnet.protocol.channel_invitation.DecrementChannelCountRequest
type DecrementChannelCountRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,req,name=agent_id" json:"agent_id,omitempty"`
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,2,opt,name=channel_id" json:"channel_id,omitempty"`
	ReservationToken *uint64                 `protobuf:"varint,3,opt,name=reservation_token" json:"reservation_token,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *DecrementChannelCountRequest) Reset()         { *m = DecrementChannelCountRequest{} }
func (m *DecrementChannelCountRequest) String() string { return proto.CompactTextString(m) }
func (*DecrementChannelCountRequest) ProtoMessage()    {}

func (m *DecrementChannelCountRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *DecrementChannelCountRequest) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *DecrementChannelCountRequest) GetReservationToken() uint64 {
	if m != nil && m.ReservationToken != nil {
		return *m.ReservationToken
	}
	return 0
}

// ref: bnet.protocol.channel_invitation.HasRoomForInvitationRequest
type HasRoomForInvitationRequest struct {
	ServiceType      *uint32 `protobuf:"varint,1,req,name=service_type" json:"service_type,omitempty"`
	Program          *uint32 `protobuf:"fixed32,2,opt,name=program" json:"program,omitempty"`
	ChannelType      *string `protobuf:"bytes,3,opt,name=channel_type,def=default" json:"channel_type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *HasRoomForInvitationRequest) Reset()         { *m = HasRoomForInvitationRequest{} }
func (m *HasRoomForInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*HasRoomForInvitationRequest) ProtoMessage()    {}

const Default_HasRoomForInvitationRequest_ChannelType string = "default"

func (m *HasRoomForInvitationRequest) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *HasRoomForInvitationRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *HasRoomForInvitationRequest) GetChannelType() string {
	if m != nil && m.ChannelType != nil {
		return *m.ChannelType
	}
	return Default_HasRoomForInvitationRequest_ChannelType
}

// ref: bnet.protocol.channel_invitation.IncrementChannelCountRequest
type IncrementChannelCountRequest struct {
	AgentId          *bnet_protocol.EntityId    `protobuf:"bytes,1,req,name=agent_id" json:"agent_id,omitempty"`
	Descriptions     []*ChannelCountDescription `protobuf:"bytes,2,rep,name=descriptions" json:"descriptions,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *IncrementChannelCountRequest) Reset()         { *m = IncrementChannelCountRequest{} }
func (m *IncrementChannelCountRequest) String() string { return proto.CompactTextString(m) }
func (*IncrementChannelCountRequest) ProtoMessage()    {}

func (m *IncrementChannelCountRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *IncrementChannelCountRequest) GetDescriptions() []*ChannelCountDescription {
	if m != nil {
		return m.Descriptions
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.IncrementChannelCountResponse
type IncrementChannelCountResponse struct {
	ReservationTokens []uint64 `protobuf:"varint,1,rep,name=reservation_tokens" json:"reservation_tokens,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *IncrementChannelCountResponse) Reset()         { *m = IncrementChannelCountResponse{} }
func (m *IncrementChannelCountResponse) String() string { return proto.CompactTextString(m) }
func (*IncrementChannelCountResponse) ProtoMessage()    {}

func (m *IncrementChannelCountResponse) GetReservationTokens() []uint64 {
	if m != nil {
		return m.ReservationTokens
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.InvitationAddedNotification
type InvitationAddedNotification struct {
	Invitation       *bnet_protocol_channel_extracted.Invitation `protobuf:"bytes,1,req,name=invitation" json:"invitation,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *InvitationAddedNotification) Reset()         { *m = InvitationAddedNotification{} }
func (m *InvitationAddedNotification) String() string { return proto.CompactTextString(m) }
func (*InvitationAddedNotification) ProtoMessage()    {}

func (m *InvitationAddedNotification) GetInvitation() *bnet_protocol_channel_extracted.Invitation {
	if m != nil {
		return m.Invitation
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.InvitationCollection
type InvitationCollection struct {
	ServiceType            *uint32                                       `protobuf:"varint,1,opt,name=service_type" json:"service_type,omitempty"`
	MaxReceivedInvitations *uint32                                       `protobuf:"varint,2,opt,name=max_received_invitations" json:"max_received_invitations,omitempty"`
	ObjectId               *uint64                                       `protobuf:"varint,3,opt,name=object_id" json:"object_id,omitempty"`
	ReceivedInvitation     []*bnet_protocol_channel_extracted.Invitation `protobuf:"bytes,4,rep,name=received_invitation" json:"received_invitation,omitempty"`
	XXX_unrecognized       []byte                                        `json:"-"`
}

func (m *InvitationCollection) Reset()         { *m = InvitationCollection{} }
func (m *InvitationCollection) String() string { return proto.CompactTextString(m) }
func (*InvitationCollection) ProtoMessage()    {}

func (m *InvitationCollection) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *InvitationCollection) GetMaxReceivedInvitations() uint32 {
	if m != nil && m.MaxReceivedInvitations != nil {
		return *m.MaxReceivedInvitations
	}
	return 0
}

func (m *InvitationCollection) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

func (m *InvitationCollection) GetReceivedInvitation() []*bnet_protocol_channel_extracted.Invitation {
	if m != nil {
		return m.ReceivedInvitation
	}
	return nil
}

// ref: bnet.protocol.invitation.InvitationParams
type InvitationParams struct {
	InvitationMessage *string                                       `protobuf:"bytes,1,opt,name=invitation_message" json:"invitation_message,omitempty"`
	ExpirationTime    *uint64                                       `protobuf:"varint,2,opt,name=expiration_time,def=0" json:"expiration_time,omitempty"`
	FriendParams      *bnet_protocol_friends.FriendInvitationParams `protobuf:"bytes,103,opt,name=friend_params" json:"friend_params,omitempty"`
	ChannelParams     *ChannelInvitationParams                      `protobuf:"bytes,105,opt,name=channel_params" json:"channel_params,omitempty"`
	XXX_unrecognized  []byte                                        `json:"-"`
}

func (m *InvitationParams) Reset()         { *m = InvitationParams{} }
func (m *InvitationParams) String() string { return proto.CompactTextString(m) }
func (*InvitationParams) ProtoMessage()    {}

const Default_InvitationParams_ExpirationTime uint64 = 0

func (m *InvitationParams) GetInvitationMessage() string {
	if m != nil && m.InvitationMessage != nil {
		return *m.InvitationMessage
	}
	return ""
}

func (m *InvitationParams) GetExpirationTime() uint64 {
	if m != nil && m.ExpirationTime != nil {
		return *m.ExpirationTime
	}
	return Default_InvitationParams_ExpirationTime
}

func (m *InvitationParams) GetFriendParams() *bnet_protocol_friends.FriendInvitationParams {
	if m != nil {
		return m.FriendParams
	}
	return nil
}

func (m *InvitationParams) GetChannelParams() *ChannelInvitationParams {
	if m != nil {
		return m.ChannelParams
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.InvitationRemovedNotification
type InvitationRemovedNotification struct {
	Invitation       *bnet_protocol_channel_extracted.Invitation `protobuf:"bytes,1,req,name=invitation" json:"invitation,omitempty"`
	Reason           *uint32                                     `protobuf:"varint,2,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte                                      `json:"-"`
}

func (m *InvitationRemovedNotification) Reset()         { *m = InvitationRemovedNotification{} }
func (m *InvitationRemovedNotification) String() string { return proto.CompactTextString(m) }
func (*InvitationRemovedNotification) ProtoMessage()    {}

func (m *InvitationRemovedNotification) GetInvitation() *bnet_protocol_channel_extracted.Invitation {
	if m != nil {
		return m.Invitation
	}
	return nil
}

func (m *InvitationRemovedNotification) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

// ref: bnet.protocol.channel_invitation.ListChannelCountRequest
type ListChannelCountRequest struct {
	MemberId         *bnet_protocol.EntityId `protobuf:"bytes,1,req,name=member_id" json:"member_id,omitempty"`
	ServiceType      *uint32                 `protobuf:"varint,2,req,name=service_type" json:"service_type,omitempty"`
	Program          *uint32                 `protobuf:"fixed32,3,opt,name=program" json:"program,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ListChannelCountRequest) Reset()         { *m = ListChannelCountRequest{} }
func (m *ListChannelCountRequest) String() string { return proto.CompactTextString(m) }
func (*ListChannelCountRequest) ProtoMessage()    {}

func (m *ListChannelCountRequest) GetMemberId() *bnet_protocol.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

func (m *ListChannelCountRequest) GetServiceType() uint32 {
	if m != nil && m.ServiceType != nil {
		return *m.ServiceType
	}
	return 0
}

func (m *ListChannelCountRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

// ref: bnet.protocol.channel_invitation.ListChannelCountResponse
type ListChannelCountResponse struct {
	Channel          []*ChannelCount `protobuf:"bytes,1,rep,name=channel" json:"channel,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ListChannelCountResponse) Reset()         { *m = ListChannelCountResponse{} }
func (m *ListChannelCountResponse) String() string { return proto.CompactTextString(m) }
func (*ListChannelCountResponse) ProtoMessage()    {}

func (m *ListChannelCountResponse) GetChannel() []*ChannelCount {
	if m != nil {
		return m.Channel
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.RevokeInvitationRequest
type RevokeInvitationRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	TargetId         *bnet_protocol.EntityId `protobuf:"bytes,2,opt,name=target_id" json:"target_id,omitempty"`
	InvitationId     *uint64                 `protobuf:"fixed64,3,req,name=invitation_id" json:"invitation_id,omitempty"`
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,4,req,name=channel_id" json:"channel_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *RevokeInvitationRequest) Reset()         { *m = RevokeInvitationRequest{} }
func (m *RevokeInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeInvitationRequest) ProtoMessage()    {}

func (m *RevokeInvitationRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *RevokeInvitationRequest) GetTargetId() *bnet_protocol.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *RevokeInvitationRequest) GetInvitationId() uint64 {
	if m != nil && m.InvitationId != nil {
		return *m.InvitationId
	}
	return 0
}

func (m *RevokeInvitationRequest) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

// ref: bnet.protocol.invitation.SendInvitationRequest
type SendInvitationRequest struct {
	AgentIdentity    *bnet_protocol.Identity                    `protobuf:"bytes,1,opt,name=agent_identity" json:"agent_identity,omitempty"`
	TargetId         *bnet_protocol.EntityId                    `protobuf:"bytes,2,req,name=target_id" json:"target_id,omitempty"`
	Params           *InvitationParams                          `protobuf:"bytes,3,req,name=params" json:"params,omitempty"`
	AgentInfo        *bnet_protocol.AccountInfo                 `protobuf:"bytes,4,opt,name=agent_info" json:"agent_info,omitempty"`
	Target           *bnet_protocol_invitation.InvitationTarget `protobuf:"bytes,5,opt,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *SendInvitationRequest) Reset()         { *m = SendInvitationRequest{} }
func (m *SendInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*SendInvitationRequest) ProtoMessage()    {}

func (m *SendInvitationRequest) GetAgentIdentity() *bnet_protocol.Identity {
	if m != nil {
		return m.AgentIdentity
	}
	return nil
}

func (m *SendInvitationRequest) GetTargetId() *bnet_protocol.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *SendInvitationRequest) GetParams() *InvitationParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *SendInvitationRequest) GetAgentInfo() *bnet_protocol.AccountInfo {
	if m != nil {
		return m.AgentInfo
	}
	return nil
}

func (m *SendInvitationRequest) GetTarget() *bnet_protocol_invitation.InvitationTarget {
	if m != nil {
		return m.Target
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.SubscribeRequest
type SubscribeRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	ObjectId         *uint64                 `protobuf:"varint,2,req,name=object_id" json:"object_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}

func (m *SubscribeRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SubscribeRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

// ref: bnet.protocol.channel_invitation.SubscribeResponse
type SubscribeResponse struct {
	Collection         []*InvitationCollection                       `protobuf:"bytes,1,rep,name=collection" json:"collection,omitempty"`
	ReceivedInvitation []*bnet_protocol_channel_extracted.Invitation `protobuf:"bytes,2,rep,name=received_invitation" json:"received_invitation,omitempty"`
	XXX_unrecognized   []byte                                        `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}

func (m *SubscribeResponse) GetCollection() []*InvitationCollection {
	if m != nil {
		return m.Collection
	}
	return nil
}

func (m *SubscribeResponse) GetReceivedInvitation() []*bnet_protocol_channel_extracted.Invitation {
	if m != nil {
		return m.ReceivedInvitation
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.SuggestInvitationRequest
type SuggestInvitationRequest struct {
	AgentId          *bnet_protocol.EntityId    `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	ChannelId        *bnet_protocol.EntityId    `protobuf:"bytes,2,req,name=channel_id" json:"channel_id,omitempty"`
	TargetId         *bnet_protocol.EntityId    `protobuf:"bytes,3,req,name=target_id" json:"target_id,omitempty"`
	ApprovalId       *bnet_protocol.EntityId    `protobuf:"bytes,4,opt,name=approval_id" json:"approval_id,omitempty"`
	AgentIdentity    *bnet_protocol.Identity    `protobuf:"bytes,5,opt,name=agent_identity" json:"agent_identity,omitempty"`
	AgentInfo        *bnet_protocol.AccountInfo `protobuf:"bytes,6,opt,name=agent_info" json:"agent_info,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *SuggestInvitationRequest) Reset()         { *m = SuggestInvitationRequest{} }
func (m *SuggestInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*SuggestInvitationRequest) ProtoMessage()    {}

func (m *SuggestInvitationRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SuggestInvitationRequest) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *SuggestInvitationRequest) GetTargetId() *bnet_protocol.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *SuggestInvitationRequest) GetApprovalId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ApprovalId
	}
	return nil
}

func (m *SuggestInvitationRequest) GetAgentIdentity() *bnet_protocol.Identity {
	if m != nil {
		return m.AgentIdentity
	}
	return nil
}

func (m *SuggestInvitationRequest) GetAgentInfo() *bnet_protocol.AccountInfo {
	if m != nil {
		return m.AgentInfo
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.SuggestionAddedNotification
type SuggestionAddedNotification struct {
	Suggestion       *bnet_protocol_invitation.Suggestion `protobuf:"bytes,1,req,name=suggestion" json:"suggestion,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *SuggestionAddedNotification) Reset()         { *m = SuggestionAddedNotification{} }
func (m *SuggestionAddedNotification) String() string { return proto.CompactTextString(m) }
func (*SuggestionAddedNotification) ProtoMessage()    {}

func (m *SuggestionAddedNotification) GetSuggestion() *bnet_protocol_invitation.Suggestion {
	if m != nil {
		return m.Suggestion
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.UnsubscribeRequest
type UnsubscribeRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UnsubscribeRequest) Reset()         { *m = UnsubscribeRequest{} }
func (m *UnsubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeRequest) ProtoMessage()    {}

func (m *UnsubscribeRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

// ref: bnet.protocol.channel_invitation.UpdateChannelCountRequest
type UpdateChannelCountRequest struct {
	AgentId          *bnet_protocol.EntityId `protobuf:"bytes,1,req,name=agent_id" json:"agent_id,omitempty"`
	ReservationToken *uint64                 `protobuf:"varint,2,opt,name=reservation_token" json:"reservation_token,omitempty"`
	ChannelId        *bnet_protocol.EntityId `protobuf:"bytes,3,req,name=channel_id" json:"channel_id,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdateChannelCountRequest) Reset()         { *m = UpdateChannelCountRequest{} }
func (m *UpdateChannelCountRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateChannelCountRequest) ProtoMessage()    {}

func (m *UpdateChannelCountRequest) GetAgentId() *bnet_protocol.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateChannelCountRequest) GetReservationToken() uint64 {
	if m != nil && m.ReservationToken != nil {
		return *m.ReservationToken
	}
	return 0
}

func (m *UpdateChannelCountRequest) GetChannelId() *bnet_protocol.EntityId {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

// ref: bnet.protocol.invitation.UpdateInvitationRequest
type UpdateInvitationRequest struct {
	AgentIdentity    *bnet_protocol.Identity `protobuf:"bytes,1,opt,name=agent_identity" json:"agent_identity,omitempty"`
	InvitationId     *uint64                 `protobuf:"fixed64,2,req,name=invitation_id" json:"invitation_id,omitempty"`
	Params           *InvitationParams       `protobuf:"bytes,3,req,name=params" json:"params,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdateInvitationRequest) Reset()         { *m = UpdateInvitationRequest{} }
func (m *UpdateInvitationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInvitationRequest) ProtoMessage()    {}

func (m *UpdateInvitationRequest) GetAgentIdentity() *bnet_protocol.Identity {
	if m != nil {
		return m.AgentIdentity
	}
	return nil
}

func (m *UpdateInvitationRequest) GetInvitationId() uint64 {
	if m != nil && m.InvitationId != nil {
		return *m.InvitationId
	}
	return 0
}

func (m *UpdateInvitationRequest) GetParams() *InvitationParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
}
