// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/channel_service/channel_service.proto
// DO NOT EDIT!

/*
Package channel_service is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/channel_service/channel_service.proto

It has these top-level messages:
	AddMemberRequest
	RemoveMemberRequest
	UnsubscribeMemberRequest
	SendMessageRequest
	UpdateChannelStateRequest
	UpdateMemberStateRequest
	DissolveRequest
	SetRolesRequest
	AddNotification
	JoinNotification
	RemoveNotification
	LeaveNotification
	SendMessageNotification
	UpdateChannelStateNotification
	UpdateMemberStateNotification
*/
package channel_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type AddMemberRequest struct {
	AgentId          *entity.EntityId           `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	MemberIdentity   *entity.Identity           `protobuf:"bytes,2,req,name=member_identity,json=memberIdentity" json:"member_identity,omitempty"`
	MemberState      *channel_types.MemberState `protobuf:"bytes,3,req,name=member_state,json=memberState" json:"member_state,omitempty"`
	ObjectId         *uint64                    `protobuf:"varint,4,req,name=object_id,json=objectId" json:"object_id,omitempty"`
	Subscribe        *bool                      `protobuf:"varint,5,opt,name=subscribe,def=1" json:"subscribe,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *AddMemberRequest) Reset()                    { *m = AddMemberRequest{} }
func (m *AddMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*AddMemberRequest) ProtoMessage()               {}
func (*AddMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_AddMemberRequest_Subscribe bool = true

func (m *AddMemberRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *AddMemberRequest) GetMemberIdentity() *entity.Identity {
	if m != nil {
		return m.MemberIdentity
	}
	return nil
}

func (m *AddMemberRequest) GetMemberState() *channel_types.MemberState {
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

type RemoveMemberRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	MemberId         *entity.EntityId `protobuf:"bytes,2,req,name=member_id,json=memberId" json:"member_id,omitempty"`
	Reason           *uint32          `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *RemoveMemberRequest) Reset()                    { *m = RemoveMemberRequest{} }
func (m *RemoveMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveMemberRequest) ProtoMessage()               {}
func (*RemoveMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RemoveMemberRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *RemoveMemberRequest) GetMemberId() *entity.EntityId {
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

type UnsubscribeMemberRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	MemberId         *entity.EntityId `protobuf:"bytes,2,req,name=member_id,json=memberId" json:"member_id,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *UnsubscribeMemberRequest) Reset()                    { *m = UnsubscribeMemberRequest{} }
func (m *UnsubscribeMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*UnsubscribeMemberRequest) ProtoMessage()               {}
func (*UnsubscribeMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UnsubscribeMemberRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UnsubscribeMemberRequest) GetMemberId() *entity.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

type SendMessageRequest struct {
	AgentId            *entity.EntityId       `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Message            *channel_types.Message `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	RequiredPrivileges *uint64                `protobuf:"varint,3,opt,name=required_privileges,json=requiredPrivileges,def=0" json:"required_privileges,omitempty"`
	XXX_unrecognized   []byte                 `json:"-"`
}

func (m *SendMessageRequest) Reset()                    { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()               {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_SendMessageRequest_RequiredPrivileges uint64 = 0

func (m *SendMessageRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SendMessageRequest) GetMessage() *channel_types.Message {
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

type UpdateChannelStateRequest struct {
	AgentId          *entity.EntityId            `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	StateChange      *channel_types.ChannelState `protobuf:"bytes,2,req,name=state_change,json=stateChange" json:"state_change,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *UpdateChannelStateRequest) Reset()                    { *m = UpdateChannelStateRequest{} }
func (m *UpdateChannelStateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateChannelStateRequest) ProtoMessage()               {}
func (*UpdateChannelStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateChannelStateRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateChannelStateRequest) GetStateChange() *channel_types.ChannelState {
	if m != nil {
		return m.StateChange
	}
	return nil
}

type UpdateMemberStateRequest struct {
	AgentId          *entity.EntityId        `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	StateChange      []*channel_types.Member `protobuf:"bytes,2,rep,name=state_change,json=stateChange" json:"state_change,omitempty"`
	RemovedRole      []uint32                `protobuf:"varint,3,rep,packed,name=removed_role,json=removedRole" json:"removed_role,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdateMemberStateRequest) Reset()                    { *m = UpdateMemberStateRequest{} }
func (m *UpdateMemberStateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateMemberStateRequest) ProtoMessage()               {}
func (*UpdateMemberStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateMemberStateRequest) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateMemberStateRequest) GetStateChange() []*channel_types.Member {
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

type DissolveRequest struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Reason           *uint32          `protobuf:"varint,2,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *DissolveRequest) Reset()                    { *m = DissolveRequest{} }
func (m *DissolveRequest) String() string            { return proto.CompactTextString(m) }
func (*DissolveRequest) ProtoMessage()               {}
func (*DissolveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DissolveRequest) GetAgentId() *entity.EntityId {
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

type SetRolesRequest struct {
	AgentId          *entity.EntityId   `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Role             []uint32           `protobuf:"varint,2,rep,packed,name=role" json:"role,omitempty"`
	MemberId         []*entity.EntityId `protobuf:"bytes,3,rep,name=member_id,json=memberId" json:"member_id,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SetRolesRequest) Reset()                    { *m = SetRolesRequest{} }
func (m *SetRolesRequest) String() string            { return proto.CompactTextString(m) }
func (*SetRolesRequest) ProtoMessage()               {}
func (*SetRolesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SetRolesRequest) GetAgentId() *entity.EntityId {
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

func (m *SetRolesRequest) GetMemberId() []*entity.EntityId {
	if m != nil {
		return m.MemberId
	}
	return nil
}

type AddNotification struct {
	Self             *channel_types.Member       `protobuf:"bytes,1,opt,name=self" json:"self,omitempty"`
	Member           []*channel_types.Member     `protobuf:"bytes,2,rep,name=member" json:"member,omitempty"`
	ChannelState     *channel_types.ChannelState `protobuf:"bytes,3,req,name=channel_state,json=channelState" json:"channel_state,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *AddNotification) Reset()                    { *m = AddNotification{} }
func (m *AddNotification) String() string            { return proto.CompactTextString(m) }
func (*AddNotification) ProtoMessage()               {}
func (*AddNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AddNotification) GetSelf() *channel_types.Member {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *AddNotification) GetMember() []*channel_types.Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *AddNotification) GetChannelState() *channel_types.ChannelState {
	if m != nil {
		return m.ChannelState
	}
	return nil
}

type JoinNotification struct {
	Member           *channel_types.Member `protobuf:"bytes,1,req,name=member" json:"member,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *JoinNotification) Reset()                    { *m = JoinNotification{} }
func (m *JoinNotification) String() string            { return proto.CompactTextString(m) }
func (*JoinNotification) ProtoMessage()               {}
func (*JoinNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *JoinNotification) GetMember() *channel_types.Member {
	if m != nil {
		return m.Member
	}
	return nil
}

type RemoveNotification struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	MemberId         *entity.EntityId `protobuf:"bytes,2,req,name=member_id,json=memberId" json:"member_id,omitempty"`
	Reason           *uint32          `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *RemoveNotification) Reset()                    { *m = RemoveNotification{} }
func (m *RemoveNotification) String() string            { return proto.CompactTextString(m) }
func (*RemoveNotification) ProtoMessage()               {}
func (*RemoveNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RemoveNotification) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *RemoveNotification) GetMemberId() *entity.EntityId {
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

type LeaveNotification struct {
	AgentId          *entity.EntityId `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	MemberId         *entity.EntityId `protobuf:"bytes,2,req,name=member_id,json=memberId" json:"member_id,omitempty"`
	Reason           *uint32          `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *LeaveNotification) Reset()                    { *m = LeaveNotification{} }
func (m *LeaveNotification) String() string            { return proto.CompactTextString(m) }
func (*LeaveNotification) ProtoMessage()               {}
func (*LeaveNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *LeaveNotification) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *LeaveNotification) GetMemberId() *entity.EntityId {
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

type SendMessageNotification struct {
	AgentId            *entity.EntityId       `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	Message            *channel_types.Message `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	RequiredPrivileges *uint64                `protobuf:"varint,3,opt,name=required_privileges,json=requiredPrivileges,def=0" json:"required_privileges,omitempty"`
	XXX_unrecognized   []byte                 `json:"-"`
}

func (m *SendMessageNotification) Reset()                    { *m = SendMessageNotification{} }
func (m *SendMessageNotification) String() string            { return proto.CompactTextString(m) }
func (*SendMessageNotification) ProtoMessage()               {}
func (*SendMessageNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

const Default_SendMessageNotification_RequiredPrivileges uint64 = 0

func (m *SendMessageNotification) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *SendMessageNotification) GetMessage() *channel_types.Message {
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

type UpdateChannelStateNotification struct {
	AgentId          *entity.EntityId            `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	StateChange      *channel_types.ChannelState `protobuf:"bytes,2,req,name=state_change,json=stateChange" json:"state_change,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *UpdateChannelStateNotification) Reset()                    { *m = UpdateChannelStateNotification{} }
func (m *UpdateChannelStateNotification) String() string            { return proto.CompactTextString(m) }
func (*UpdateChannelStateNotification) ProtoMessage()               {}
func (*UpdateChannelStateNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UpdateChannelStateNotification) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateChannelStateNotification) GetStateChange() *channel_types.ChannelState {
	if m != nil {
		return m.StateChange
	}
	return nil
}

type UpdateMemberStateNotification struct {
	AgentId          *entity.EntityId        `protobuf:"bytes,1,opt,name=agent_id,json=agentId" json:"agent_id,omitempty"`
	StateChange      []*channel_types.Member `protobuf:"bytes,2,rep,name=state_change,json=stateChange" json:"state_change,omitempty"`
	RemovedRole      []uint32                `protobuf:"varint,3,rep,packed,name=removed_role,json=removedRole" json:"removed_role,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *UpdateMemberStateNotification) Reset()                    { *m = UpdateMemberStateNotification{} }
func (m *UpdateMemberStateNotification) String() string            { return proto.CompactTextString(m) }
func (*UpdateMemberStateNotification) ProtoMessage()               {}
func (*UpdateMemberStateNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *UpdateMemberStateNotification) GetAgentId() *entity.EntityId {
	if m != nil {
		return m.AgentId
	}
	return nil
}

func (m *UpdateMemberStateNotification) GetStateChange() []*channel_types.Member {
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

func init() {
	proto.RegisterType((*AddMemberRequest)(nil), "channel_service.AddMemberRequest")
	proto.RegisterType((*RemoveMemberRequest)(nil), "channel_service.RemoveMemberRequest")
	proto.RegisterType((*UnsubscribeMemberRequest)(nil), "channel_service.UnsubscribeMemberRequest")
	proto.RegisterType((*SendMessageRequest)(nil), "channel_service.SendMessageRequest")
	proto.RegisterType((*UpdateChannelStateRequest)(nil), "channel_service.UpdateChannelStateRequest")
	proto.RegisterType((*UpdateMemberStateRequest)(nil), "channel_service.UpdateMemberStateRequest")
	proto.RegisterType((*DissolveRequest)(nil), "channel_service.DissolveRequest")
	proto.RegisterType((*SetRolesRequest)(nil), "channel_service.SetRolesRequest")
	proto.RegisterType((*AddNotification)(nil), "channel_service.AddNotification")
	proto.RegisterType((*JoinNotification)(nil), "channel_service.JoinNotification")
	proto.RegisterType((*RemoveNotification)(nil), "channel_service.RemoveNotification")
	proto.RegisterType((*LeaveNotification)(nil), "channel_service.LeaveNotification")
	proto.RegisterType((*SendMessageNotification)(nil), "channel_service.SendMessageNotification")
	proto.RegisterType((*UpdateChannelStateNotification)(nil), "channel_service.UpdateChannelStateNotification")
	proto.RegisterType((*UpdateMemberStateNotification)(nil), "channel_service.UpdateMemberStateNotification")
}

func init() {
	proto.RegisterFile("github.com/HearthSim/hs-proto-go/bnet/channel_service/channel_service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x95, 0x51, 0x6f, 0x12, 0x41,
	0x10, 0xc7, 0xb3, 0x80, 0x2d, 0x0c, 0xb4, 0xd0, 0x6d, 0xac, 0x67, 0x1b, 0x0d, 0xb9, 0xc4, 0x04,
	0x63, 0x80, 0x86, 0x27, 0xdb, 0x44, 0x63, 0xad, 0x46, 0x51, 0xab, 0xe6, 0x48, 0x7d, 0x25, 0xc7,
	0xdd, 0x14, 0xd6, 0x70, 0x77, 0x74, 0x77, 0xc1, 0xf4, 0x03, 0x18, 0xf5, 0xc1, 0xcf, 0xa1, 0x0f,
	0xf6, 0x0b, 0xfa, 0x64, 0xd8, 0x3d, 0x0e, 0x0e, 0x94, 0xe8, 0x35, 0x56, 0x9f, 0x8e, 0xd9, 0xfb,
	0xcf, 0xec, 0x6f, 0x96, 0xd9, 0xff, 0xc1, 0xf3, 0x2e, 0x93, 0xbd, 0x61, 0xa7, 0xe6, 0x04, 0x5e,
	0xfd, 0x29, 0xda, 0x5c, 0xf6, 0x5a, 0xcc, 0xab, 0xf7, 0x44, 0x75, 0xc0, 0x03, 0x19, 0x54, 0xbb,
	0x41, 0xbd, 0xe3, 0xa3, 0xac, 0x3b, 0x3d, 0xdb, 0xf7, 0xb1, 0xdf, 0x16, 0xc8, 0x47, 0xcc, 0xc1,
	0xf9, 0xb8, 0xa6, 0xd4, 0xb4, 0x38, 0xb7, 0xbc, 0xfd, 0xe4, 0xcf, 0xaa, 0xcb, 0xb3, 0x01, 0x8a,
	0x78, 0xa4, 0x2b, 0x6f, 0xef, 0xfd, 0x5e, 0x21, 0xf4, 0x25, 0x93, 0x67, 0xe1, 0x43, 0xa7, 0x9a,
	0xdf, 0x09, 0x94, 0x0e, 0x5c, 0xf7, 0x08, 0xbd, 0x0e, 0x72, 0x0b, 0x4f, 0x87, 0x28, 0x24, 0xbd,
	0x03, 0x59, 0xbb, 0x8b, 0xbe, 0x6c, 0x33, 0xd7, 0x20, 0x65, 0x52, 0xc9, 0x37, 0x4a, 0xb5, 0x30,
	0xeb, 0xb1, 0x7a, 0x34, 0x5d, 0x6b, 0x55, 0x29, 0x9a, 0x2e, 0xdd, 0x83, 0xa2, 0xa7, 0xb2, 0xdb,
	0xcc, 0xd5, 0x22, 0x23, 0x55, 0x4e, 0xcd, 0xe6, 0x34, 0xc3, 0x75, 0x6b, 0x5d, 0x0b, 0x27, 0x31,
	0xbd, 0x07, 0x85, 0x30, 0x55, 0x48, 0x5b, 0xa2, 0x91, 0x56, 0x79, 0xdb, 0xb5, 0x78, 0x8f, 0x9a,
	0xad, 0x35, 0x56, 0x58, 0x79, 0x6f, 0x1a, 0xd0, 0x1d, 0xc8, 0x05, 0x9d, 0xb7, 0xe8, 0x28, 0xce,
	0x4c, 0x39, 0x55, 0xc9, 0x58, 0x59, 0xbd, 0xd0, 0x74, 0xa9, 0x09, 0x39, 0x31, 0xec, 0x08, 0x87,
	0xb3, 0x0e, 0x1a, 0x57, 0xca, 0xa4, 0x92, 0xdd, 0xcf, 0x48, 0x3e, 0x44, 0x6b, 0xba, 0x6c, 0x7e,
	0x22, 0xb0, 0x69, 0xa1, 0x17, 0x8c, 0xf0, 0x02, 0xfd, 0x57, 0x21, 0x17, 0xf5, 0x3f, 0xdf, 0x79,
	0xa4, 0xce, 0x4e, 0x3a, 0xa7, 0x5b, 0xb0, 0xc2, 0xd1, 0x16, 0x81, 0x6f, 0xa4, 0xcb, 0xa4, 0xb2,
	0x66, 0x85, 0x91, 0x39, 0x02, 0xe3, 0xd8, 0x8f, 0xd0, 0x2e, 0x8d, 0xc7, 0xfc, 0x42, 0x80, 0xb6,
	0xd0, 0x77, 0x8f, 0x50, 0x08, 0xbb, 0x8b, 0x89, 0xb6, 0xdc, 0x85, 0x55, 0x4f, 0xa7, 0x87, 0x1b,
	0x6e, 0x2d, 0xfc, 0x85, 0xba, 0xf8, 0x44, 0x46, 0x1b, 0xb0, 0xc9, 0xf1, 0x74, 0xc8, 0x38, 0xba,
	0xed, 0x01, 0x67, 0x23, 0xd6, 0xc7, 0x2e, 0x0a, 0x75, 0x24, 0x99, 0x7d, 0xb2, 0x6b, 0xd1, 0xc9,
	0xdb, 0xd7, 0xd1, 0x4b, 0xf3, 0x23, 0x81, 0xeb, 0xc7, 0x03, 0xd7, 0x96, 0x78, 0xa8, 0x8b, 0xeb,
	0x91, 0x48, 0x02, 0x7c, 0x1f, 0x0a, 0x6a, 0xe2, 0xda, 0x63, 0xcc, 0x88, 0x7a, 0x67, 0x8e, 0x3a,
	0xb6, 0x4d, 0x5e, 0x25, 0x1c, 0x2a, 0xbd, 0xf9, 0x95, 0x80, 0xa1, 0x51, 0x66, 0x87, 0x33, 0x09,
	0xc9, 0xdd, 0x05, 0x92, 0x74, 0x25, 0xdf, 0xb8, 0xfa, 0xd3, 0x2b, 0x10, 0x63, 0xa0, 0xb7, 0xa0,
	0xc0, 0xd5, 0xec, 0xba, 0x6d, 0x1e, 0xf4, 0xc7, 0x97, 0x27, 0x5d, 0x59, 0x7b, 0x98, 0x2a, 0x11,
	0x2b, 0x1f, 0xae, 0x5b, 0x41, 0x1f, 0xcd, 0x37, 0x50, 0x7c, 0xc4, 0x84, 0x08, 0xfa, 0xa3, 0x64,
	0x80, 0xd3, 0x79, 0x4d, 0xc5, 0xe6, 0xf5, 0x3d, 0x81, 0x62, 0x0b, 0xe5, 0x78, 0x0f, 0x91, 0xb0,
	0x70, 0x46, 0x71, 0xa7, 0x22, 0x6e, 0x15, 0xc7, 0xe7, 0x37, 0xad, 0x8e, 0x63, 0xd9, 0xfc, 0x9e,
	0x13, 0x28, 0x1e, 0xb8, 0xee, 0xcb, 0x40, 0xb2, 0x13, 0xe6, 0xd8, 0x92, 0x05, 0x3e, 0xbd, 0x0d,
	0x19, 0x81, 0xfd, 0x93, 0x90, 0xe1, 0x17, 0x87, 0xa9, 0x24, 0xb4, 0x0a, 0x2b, 0xba, 0xd4, 0xf2,
	0x93, 0x0f, 0x45, 0xf4, 0x01, 0xac, 0x45, 0x2e, 0x3e, 0x63, 0x59, 0x4b, 0x27, 0xa7, 0xe0, 0xcc,
	0x44, 0xe6, 0x01, 0x94, 0x9e, 0x05, 0xcc, 0x8f, 0xf1, 0x4e, 0x21, 0x88, 0x2a, 0xb7, 0x1c, 0x62,
	0x7c, 0x11, 0xa8, 0xb6, 0xad, 0x58, 0x95, 0x7f, 0xe1, 0x5a, 0x1f, 0x08, 0x6c, 0xbc, 0x40, 0xfb,
	0x3f, 0x20, 0xf9, 0x46, 0xe0, 0xda, 0x8c, 0x8f, 0x25, 0xe7, 0xb9, 0x1c, 0x33, 0xfb, 0x4c, 0xe0,
	0xe6, 0xa2, 0x99, 0x25, 0xa7, 0xbe, 0xa8, 0xa3, 0x9d, 0x13, 0xb8, 0xb1, 0xe0, 0x68, 0xc9, 0x71,
	0xfe, 0xb6, 0xad, 0x35, 0x72, 0xb0, 0x1a, 0x36, 0xd3, 0x58, 0x87, 0x42, 0xf8, 0xf3, 0xd5, 0x3b,
	0x1f, 0x79, 0x63, 0x13, 0x36, 0x26, 0x7d, 0x4e, 0x3e, 0xa7, 0xfc, 0x47, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x60, 0x26, 0xb4, 0x81, 0xca, 0x09, 0x00, 0x00,
}