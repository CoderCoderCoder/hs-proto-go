// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/game_utilities_service/game_utilities_service.proto
// DO NOT EDIT!

/*
Package game_utilities_service is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/game_utilities_service/game_utilities_service.proto

It has these top-level messages:
	ClientRequest
	ClientResponse
	ServerRequest
	ServerResponse
	PresenceChannelCreatedRequest
	GetPlayerVariablesRequest
	GetPlayerVariablesResponse
	GameAccountOnlineNotification
	GameAccountOfflineNotification
*/
package game_utilities_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import attribute "github.com/HearthSim/hs-proto-go/bnet/attribute"
import entity "github.com/HearthSim/hs-proto-go/bnet/entity"
import game_utilities_types "github.com/HearthSim/hs-proto-go/bnet/game_utilities_types"
import rpc "github.com/HearthSim/hs-proto-go/bnet/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientRequest struct {
	Attribute        []*attribute.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	Host             *rpc.ProcessId         `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	BnetAccountId    *entity.EntityId       `protobuf:"bytes,3,opt,name=bnet_account_id,json=bnetAccountId" json:"bnet_account_id,omitempty"`
	GameAccountId    *entity.EntityId       `protobuf:"bytes,4,opt,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ClientRequest) Reset()                    { *m = ClientRequest{} }
func (m *ClientRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientRequest) ProtoMessage()               {}
func (*ClientRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientRequest) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *ClientRequest) GetHost() *rpc.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ClientRequest) GetBnetAccountId() *entity.EntityId {
	if m != nil {
		return m.BnetAccountId
	}
	return nil
}

func (m *ClientRequest) GetGameAccountId() *entity.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

type ClientResponse struct {
	Attribute        []*attribute.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ClientResponse) Reset()                    { *m = ClientResponse{} }
func (m *ClientResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientResponse) ProtoMessage()               {}
func (*ClientResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientResponse) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type ServerRequest struct {
	Attribute        []*attribute.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	Program          *uint32                `protobuf:"fixed32,2,req,name=program" json:"program,omitempty"`
	Host             *rpc.ProcessId         `protobuf:"bytes,3,opt,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ServerRequest) Reset()                    { *m = ServerRequest{} }
func (m *ServerRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerRequest) ProtoMessage()               {}
func (*ServerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServerRequest) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *ServerRequest) GetProgram() uint32 {
	if m != nil && m.Program != nil {
		return *m.Program
	}
	return 0
}

func (m *ServerRequest) GetHost() *rpc.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

type ServerResponse struct {
	Attribute        []*attribute.Attribute `protobuf:"bytes,1,rep,name=attribute" json:"attribute,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ServerResponse) Reset()                    { *m = ServerResponse{} }
func (m *ServerResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerResponse) ProtoMessage()               {}
func (*ServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ServerResponse) GetAttribute() []*attribute.Attribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

type PresenceChannelCreatedRequest struct {
	Id               *entity.EntityId `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	GameAccountId    *entity.EntityId `protobuf:"bytes,3,opt,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	BnetAccountId    *entity.EntityId `protobuf:"bytes,4,opt,name=bnet_account_id,json=bnetAccountId" json:"bnet_account_id,omitempty"`
	Host             *rpc.ProcessId   `protobuf:"bytes,5,opt,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *PresenceChannelCreatedRequest) Reset()                    { *m = PresenceChannelCreatedRequest{} }
func (m *PresenceChannelCreatedRequest) String() string            { return proto.CompactTextString(m) }
func (*PresenceChannelCreatedRequest) ProtoMessage()               {}
func (*PresenceChannelCreatedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PresenceChannelCreatedRequest) GetId() *entity.EntityId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PresenceChannelCreatedRequest) GetGameAccountId() *entity.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

func (m *PresenceChannelCreatedRequest) GetBnetAccountId() *entity.EntityId {
	if m != nil {
		return m.BnetAccountId
	}
	return nil
}

func (m *PresenceChannelCreatedRequest) GetHost() *rpc.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

type GetPlayerVariablesRequest struct {
	PlayerVariables  []*game_utilities_types.PlayerVariables `protobuf:"bytes,1,rep,name=player_variables,json=playerVariables" json:"player_variables,omitempty"`
	Host             *rpc.ProcessId                          `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte                                  `json:"-"`
}

func (m *GetPlayerVariablesRequest) Reset()                    { *m = GetPlayerVariablesRequest{} }
func (m *GetPlayerVariablesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPlayerVariablesRequest) ProtoMessage()               {}
func (*GetPlayerVariablesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetPlayerVariablesRequest) GetPlayerVariables() []*game_utilities_types.PlayerVariables {
	if m != nil {
		return m.PlayerVariables
	}
	return nil
}

func (m *GetPlayerVariablesRequest) GetHost() *rpc.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

type GetPlayerVariablesResponse struct {
	PlayerVariables  []*game_utilities_types.PlayerVariables `protobuf:"bytes,1,rep,name=player_variables,json=playerVariables" json:"player_variables,omitempty"`
	XXX_unrecognized []byte                                  `json:"-"`
}

func (m *GetPlayerVariablesResponse) Reset()                    { *m = GetPlayerVariablesResponse{} }
func (m *GetPlayerVariablesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetPlayerVariablesResponse) ProtoMessage()               {}
func (*GetPlayerVariablesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetPlayerVariablesResponse) GetPlayerVariables() []*game_utilities_types.PlayerVariables {
	if m != nil {
		return m.PlayerVariables
	}
	return nil
}

type GameAccountOnlineNotification struct {
	GameAccountId    *entity.EntityId `protobuf:"bytes,1,req,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	Host             *rpc.ProcessId   `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GameAccountOnlineNotification) Reset()                    { *m = GameAccountOnlineNotification{} }
func (m *GameAccountOnlineNotification) String() string            { return proto.CompactTextString(m) }
func (*GameAccountOnlineNotification) ProtoMessage()               {}
func (*GameAccountOnlineNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GameAccountOnlineNotification) GetGameAccountId() *entity.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

func (m *GameAccountOnlineNotification) GetHost() *rpc.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

type GameAccountOfflineNotification struct {
	GameAccountId    *entity.EntityId `protobuf:"bytes,1,req,name=game_account_id,json=gameAccountId" json:"game_account_id,omitempty"`
	Host             *rpc.ProcessId   `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *GameAccountOfflineNotification) Reset()                    { *m = GameAccountOfflineNotification{} }
func (m *GameAccountOfflineNotification) String() string            { return proto.CompactTextString(m) }
func (*GameAccountOfflineNotification) ProtoMessage()               {}
func (*GameAccountOfflineNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GameAccountOfflineNotification) GetGameAccountId() *entity.EntityId {
	if m != nil {
		return m.GameAccountId
	}
	return nil
}

func (m *GameAccountOfflineNotification) GetHost() *rpc.ProcessId {
	if m != nil {
		return m.Host
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientRequest)(nil), "game_utilities_service.ClientRequest")
	proto.RegisterType((*ClientResponse)(nil), "game_utilities_service.ClientResponse")
	proto.RegisterType((*ServerRequest)(nil), "game_utilities_service.ServerRequest")
	proto.RegisterType((*ServerResponse)(nil), "game_utilities_service.ServerResponse")
	proto.RegisterType((*PresenceChannelCreatedRequest)(nil), "game_utilities_service.PresenceChannelCreatedRequest")
	proto.RegisterType((*GetPlayerVariablesRequest)(nil), "game_utilities_service.GetPlayerVariablesRequest")
	proto.RegisterType((*GetPlayerVariablesResponse)(nil), "game_utilities_service.GetPlayerVariablesResponse")
	proto.RegisterType((*GameAccountOnlineNotification)(nil), "game_utilities_service.GameAccountOnlineNotification")
	proto.RegisterType((*GameAccountOfflineNotification)(nil), "game_utilities_service.GameAccountOfflineNotification")
}

func init() {
	proto.RegisterFile("github.com/HearthSim/hs-proto-go/bnet/game_utilities_service/game_utilities_service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x53, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x25, 0x69, 0x65, 0x71, 0x96, 0x6e, 0x97, 0x22, 0x12, 0x0b, 0x2b, 0x21, 0x20, 0xf4, 0x65,
	0x53, 0xa8, 0x2f, 0xfa, 0x24, 0x4b, 0x95, 0xb5, 0x2f, 0x5a, 0xb2, 0xac, 0xe0, 0x53, 0x98, 0x26,
	0xb7, 0xcd, 0x85, 0x74, 0x26, 0xce, 0xdc, 0x14, 0xfa, 0xa0, 0xe0, 0x27, 0xf8, 0x79, 0xe2, 0xcf,
	0x48, 0xd2, 0x4c, 0xbb, 0x2e, 0x89, 0x84, 0xad, 0xf8, 0x50, 0x66, 0x3a, 0xf7, 0x9e, 0x73, 0xe7,
	0x9c, 0x33, 0x61, 0x9f, 0x57, 0x48, 0x49, 0xbe, 0xf0, 0x23, 0xb9, 0x1e, 0xbf, 0x07, 0xae, 0x28,
	0xb9, 0xc1, 0xf5, 0x38, 0xd1, 0x97, 0x99, 0x92, 0x24, 0x2f, 0x57, 0x72, 0xbc, 0x10, 0x40, 0xe3,
	0x15, 0x5f, 0x43, 0x98, 0x13, 0xa6, 0x48, 0x08, 0x3a, 0xd4, 0xa0, 0x36, 0x18, 0x41, 0xc3, 0xb1,
	0x5f, 0x62, 0x07, 0x4f, 0xeb, 0xab, 0xc3, 0x37, 0xed, 0x46, 0x72, 0x22, 0x85, 0x8b, 0x9c, 0xe0,
	0xb0, 0xdb, 0x11, 0x0f, 0x5f, 0xb7, 0x23, 0x00, 0x41, 0x48, 0xdb, 0x6a, 0xa9, 0xa0, 0xb7, 0x0f,
	0x92, 0x4b, 0xdb, 0x0c, 0x74, 0xed, 0x61, 0x45, 0xfb, 0xb2, 0x1d, 0xad, 0xca, 0xa2, 0xe2, 0xb7,
	0x03, 0x79, 0xbf, 0x2c, 0xd6, 0x9b, 0xa6, 0x08, 0x82, 0x02, 0xf8, 0x92, 0x83, 0xa6, 0xc1, 0x84,
	0x3d, 0xde, 0x6b, 0x75, 0x2c, 0xb7, 0x33, 0x3a, 0x9d, 0x3c, 0xf1, 0x0f, 0xea, 0xaf, 0xcc, 0x2e,
	0x38, 0xb4, 0x0d, 0x3c, 0xd6, 0x4d, 0xa4, 0x26, 0xc7, 0x76, 0xad, 0xd1, 0xe9, 0xe4, 0xcc, 0x2f,
	0xf8, 0xe7, 0x4a, 0x46, 0xa0, 0xf5, 0x2c, 0x0e, 0xca, 0xda, 0xe0, 0x15, 0xeb, 0x17, 0xf3, 0x43,
	0x1e, 0x45, 0x32, 0x17, 0x14, 0x62, 0xec, 0x74, 0xca, 0xf6, 0x73, 0xbf, 0x72, 0xe7, 0x5d, 0xb9,
	0xcc, 0xe2, 0xa0, 0x57, 0x34, 0x5e, 0xed, 0xfa, 0x66, 0x71, 0x81, 0x2c, 0x65, 0xdf, 0x41, 0x76,
	0x9b, 0x90, 0x45, 0xe3, 0x1e, 0xe9, 0xbd, 0x65, 0x67, 0x46, 0x9c, 0xce, 0xa4, 0xd0, 0xf0, 0x10,
	0x75, 0xde, 0x77, 0x8b, 0xf5, 0x6e, 0x40, 0x6d, 0x40, 0x1d, 0xe3, 0x91, 0xc3, 0x4e, 0x32, 0x25,
	0x57, 0x8a, 0xaf, 0x1d, 0xdb, 0xb5, 0x47, 0x27, 0x81, 0xf9, 0xbb, 0x77, 0xaf, 0xd3, 0xec, 0x5e,
	0xa1, 0xc4, 0x5c, 0xe1, 0x08, 0x25, 0x3f, 0x2d, 0x76, 0x31, 0x57, 0xa0, 0x41, 0x44, 0x30, 0x4d,
	0xb8, 0x10, 0x90, 0x4e, 0x15, 0x70, 0x82, 0xd8, 0x28, 0x73, 0x99, 0x8d, 0xb1, 0x63, 0xb9, 0x76,
	0xad, 0xbd, 0x36, 0xd6, 0xa6, 0xd1, 0x69, 0x95, 0x46, 0xdd, 0x0b, 0xe8, 0xb6, 0x7b, 0x01, 0xc6,
	0xa1, 0x47, 0x7f, 0x71, 0xe8, 0x87, 0xc5, 0x9e, 0x5d, 0x03, 0xcd, 0x53, 0xbe, 0x05, 0xf5, 0x89,
	0x2b, 0xe4, 0x8b, 0x14, 0xb4, 0xd1, 0x35, 0x67, 0xe7, 0x59, 0x59, 0x09, 0x37, 0xa6, 0x54, 0x99,
	0xf6, 0xc2, 0xaf, 0xfd, 0xa6, 0xee, 0xf3, 0xf4, 0xb3, 0x3f, 0x0f, 0xda, 0xbc, 0x79, 0x4f, 0xb0,
	0x61, 0xdd, 0x95, 0xaa, 0x04, 0xff, 0xf9, 0x9d, 0xbc, 0xaf, 0xec, 0xe2, 0xfa, 0x60, 0xf9, 0x47,
	0x91, 0xa2, 0x80, 0x0f, 0x92, 0x70, 0x89, 0x11, 0x27, 0x94, 0xa2, 0x2e, 0xbc, 0xa6, 0xac, 0xef,
	0x85, 0xd7, 0x46, 0xee, 0x37, 0xf6, 0xfc, 0xee, 0xf8, 0xe5, 0xf2, 0xff, 0xce, 0x9f, 0xf4, 0x59,
	0xaf, 0x98, 0x7f, 0x6b, 0x6c, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x25, 0x1a, 0x19, 0x19, 0x59,
	0x06, 0x00, 0x00,
}