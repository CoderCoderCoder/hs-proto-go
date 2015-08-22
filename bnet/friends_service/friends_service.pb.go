// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/friends_service/friends_service.proto
// DO NOT EDIT!

/*
Package friends_service is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/friends_service/friends_service.proto

It has these top-level messages:
	SubscribeToFriendsRequest
	SubscribeToFriendsResponse
	UnsubscribeToFriendsRequest
	GenericFriendRequest
	GenericFriendResponse
	AssignRoleRequest
	ViewFriendsRequest
	ViewFriendsResponse
	UpdateFriendStateRequest
	FriendNotification
	UpdateFriendStateNotification
	InvitationNotification
*/
package friends_service

import proto "github.com/golang/protobuf/proto"
import math "math"
import attribute "github.com/HearthSim/hs-proto-go/bnet/attribute"
import entity "github.com/HearthSim/hs-proto-go/bnet/entity"
import friends_types "github.com/HearthSim/hs-proto-go/bnet/friends_types"
import invitation_types "github.com/HearthSim/hs-proto-go/bnet/invitation_types"
import role "github.com/HearthSim/hs-proto-go/bnet/role"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SubscribeToFriendsRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	ObjectId         *uint64          `protobuf:"varint,2,req,name=object_id" json:"object_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *SubscribeToFriendsRequest) Reset()         { *m = SubscribeToFriendsRequest{} }
func (m *SubscribeToFriendsRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeToFriendsRequest) ProtoMessage()    {}

func (m *SubscribeToFriendsRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SubscribeToFriendsRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

type SubscribeToFriendsResponse struct {
	MaxFriends             *uint32                        `protobuf:"varint,1,opt,name=max_friends" json:"max_friends,omitempty"`
	MaxReceivedInvitations *uint32                        `protobuf:"varint,2,opt,name=max_received_invitations" json:"max_received_invitations,omitempty"`
	MaxSentInvitations     *uint32                        `protobuf:"varint,3,opt,name=max_sent_invitations" json:"max_sent_invitations,omitempty"`
	Role                   []*role.Role                   `protobuf:"bytes,4,rep,name=role" json:"role,omitempty"`
	Friends                []*friends_types.Friend        `protobuf:"bytes,5,rep,name=friends" json:"friends,omitempty"`
	SentInvitations        []*invitation_types.Invitation `protobuf:"bytes,6,rep,name=sent_invitations" json:"sent_invitations,omitempty"`
	ReceivedInvitations    []*invitation_types.Invitation `protobuf:"bytes,7,rep,name=received_invitations" json:"received_invitations,omitempty"`
	XXX_unrecognized       []byte                         `json:"-"`
}

func (m *SubscribeToFriendsResponse) Reset()         { *m = SubscribeToFriendsResponse{} }
func (m *SubscribeToFriendsResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeToFriendsResponse) ProtoMessage()    {}

func (m *SubscribeToFriendsResponse) GetMaxFriends() uint32 {
	if m != nil && m.MaxFriends != nil {
		return *m.MaxFriends
	}
	return 0
}

func (m *SubscribeToFriendsResponse) GetMaxReceivedInvitations() uint32 {
	if m != nil && m.MaxReceivedInvitations != nil {
		return *m.MaxReceivedInvitations
	}
	return 0
}

func (m *SubscribeToFriendsResponse) GetMaxSentInvitations() uint32 {
	if m != nil && m.MaxSentInvitations != nil {
		return *m.MaxSentInvitations
	}
	return 0
}

func (m *SubscribeToFriendsResponse) GetRole() []*role.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *SubscribeToFriendsResponse) GetFriends() []*friends_types.Friend {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *SubscribeToFriendsResponse) GetSentInvitations() []*invitation_types.Invitation {
	if m != nil {
		return m.SentInvitations
	}
	return nil
}

func (m *SubscribeToFriendsResponse) GetReceivedInvitations() []*invitation_types.Invitation {
	if m != nil {
		return m.ReceivedInvitations
	}
	return nil
}

type UnsubscribeToFriendsRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	ObjectId         *uint64          `protobuf:"varint,2,opt,name=object_id" json:"object_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *UnsubscribeToFriendsRequest) Reset()         { *m = UnsubscribeToFriendsRequest{} }
func (m *UnsubscribeToFriendsRequest) String() string { return proto.CompactTextString(m) }
func (*UnsubscribeToFriendsRequest) ProtoMessage()    {}

func (m *UnsubscribeToFriendsRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UnsubscribeToFriendsRequest) GetObjectId() uint64 {
	if m != nil && m.ObjectId != nil {
		return *m.ObjectId
	}
	return 0
}

type GenericFriendRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	TargetId         *entity.EntityId `protobuf:"bytes,2,req,name=target_id" json:"target_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GenericFriendRequest) Reset()         { *m = GenericFriendRequest{} }
func (m *GenericFriendRequest) String() string { return proto.CompactTextString(m) }
func (*GenericFriendRequest) ProtoMessage()    {}

func (m *GenericFriendRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *GenericFriendRequest) GetTargetId() *entity.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

type GenericFriendResponse struct {
	TargetFriend     *friends_types.Friend `protobuf:"bytes,1,opt,name=target_friend" json:"target_friend,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *GenericFriendResponse) Reset()         { *m = GenericFriendResponse{} }
func (m *GenericFriendResponse) String() string { return proto.CompactTextString(m) }
func (*GenericFriendResponse) ProtoMessage()    {}

func (m *GenericFriendResponse) GetTargetFriend() *friends_types.Friend {
	if m != nil {
		return m.TargetFriend
	}
	return nil
}

type AssignRoleRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	TargetId         *entity.EntityId `protobuf:"bytes,2,req,name=target_id" json:"target_id,omitempty"`
	Role             []int32          `protobuf:"varint,3,rep,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *AssignRoleRequest) Reset()         { *m = AssignRoleRequest{} }
func (m *AssignRoleRequest) String() string { return proto.CompactTextString(m) }
func (*AssignRoleRequest) ProtoMessage()    {}

func (m *AssignRoleRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *AssignRoleRequest) GetTargetId() *entity.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *AssignRoleRequest) GetRole() []int32 {
	if m != nil {
		return m.Role
	}
	return nil
}

type ViewFriendsRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	TargetId         *entity.EntityId `protobuf:"bytes,2,req,name=target_id" json:"target_id,omitempty"`
	Role             []uint32         `protobuf:"varint,3,rep,packed,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ViewFriendsRequest) Reset()         { *m = ViewFriendsRequest{} }
func (m *ViewFriendsRequest) String() string { return proto.CompactTextString(m) }
func (*ViewFriendsRequest) ProtoMessage()    {}

func (m *ViewFriendsRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *ViewFriendsRequest) GetTargetId() *entity.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *ViewFriendsRequest) GetRole() []uint32 {
	if m != nil {
		return m.Role
	}
	return nil
}

type ViewFriendsResponse struct {
	Friends          []*friends_types.Friend `protobuf:"bytes,1,rep,name=friends" json:"friends,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *ViewFriendsResponse) Reset()         { *m = ViewFriendsResponse{} }
func (m *ViewFriendsResponse) String() string { return proto.CompactTextString(m) }
func (*ViewFriendsResponse) ProtoMessage()    {}

func (m *ViewFriendsResponse) GetFriends() []*friends_types.Friend {
	if m != nil {
		return m.Friends
	}
	return nil
}

type UpdateFriendStateRequest struct {
	AgentId          *entity.EntityId       `protobuf:"bytes,1,opt,name=agent_id" json:"agent_id,omitempty"`
	TargetId         *entity.EntityId       `protobuf:"bytes,2,req,name=target_id" json:"target_id,omitempty"`
	Attribute        []*attribute.Attribute `protobuf:"bytes,3,rep,name=attribute" json:"attribute,omitempty"`
	AttributesEpoch  *uint64                `protobuf:"varint,4,opt,name=attributes_epoch" json:"attributes_epoch,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *UpdateFriendStateRequest) Reset()         { *m = UpdateFriendStateRequest{} }
func (m *UpdateFriendStateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFriendStateRequest) ProtoMessage()    {}

func (m *UpdateFriendStateRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateFriendStateRequest) GetTargetId() *entity.EntityId {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *UpdateFriendStateRequest) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *UpdateFriendStateRequest) GetAttributesEpoch() uint64 {
	if m != nil && m.AttributesEpoch != nil {
		return *m.AttributesEpoch
	}
	return 0
}

type FriendNotification struct {
	Target           *friends_types.Friend `protobuf:"bytes,1,req,name=target" json:"target,omitempty"`
	GameAccountId    *entity.EntityId      `protobuf:"bytes,2,opt,name=game_account_id" json:"game_account_id,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *FriendNotification) Reset()         { *m = FriendNotification{} }
func (m *FriendNotification) String() string { return proto.CompactTextString(m) }
func (*FriendNotification) ProtoMessage()    {}

func (m *FriendNotification) GetTarget() *friends_types.Friend {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *FriendNotification) GetGameAccountId() *entity.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

type UpdateFriendStateNotification struct {
	ChangedFriend    *friends_types.Friend `protobuf:"bytes,1,req,name=changed_friend" json:"changed_friend,omitempty"`
	GameAccountId    *entity.EntityId      `protobuf:"bytes,2,opt,name=game_account_id" json:"game_account_id,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *UpdateFriendStateNotification) Reset()         { *m = UpdateFriendStateNotification{} }
func (m *UpdateFriendStateNotification) String() string { return proto.CompactTextString(m) }
func (*UpdateFriendStateNotification) ProtoMessage()    {}

func (m *UpdateFriendStateNotification) GetChangedFriend() *friends_types.Friend {
	if m != nil {
		return m.ChangedFriend
	}
	return nil
}

func (m *UpdateFriendStateNotification) GetGameAccountId() *entity.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

type InvitationNotification struct {
	Invitation       *invitation_types.Invitation `protobuf:"bytes,1,req,name=invitation" json:"invitation,omitempty"`
	GameAccountId    *entity.EntityId             `protobuf:"bytes,2,opt,name=game_account_id" json:"game_account_id,omitempty"`
	Reason           *uint32                      `protobuf:"varint,3,opt,name=reason,def=0" json:"reason,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *InvitationNotification) Reset()         { *m = InvitationNotification{} }
func (m *InvitationNotification) String() string { return proto.CompactTextString(m) }
func (*InvitationNotification) ProtoMessage()    {}

const Default_InvitationNotification_Reason uint32 = 0

func (m *InvitationNotification) GetInvitation() *invitation_types.Invitation {
	if m != nil {
		return m.Invitation
	}
	return nil
}

func (m *InvitationNotification) GetGameAccountId() *entity.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

func (m *InvitationNotification) GetReason() uint32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return Default_InvitationNotification_Reason
}
