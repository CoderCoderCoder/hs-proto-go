// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/resource_service/resource_service.proto
// DO NOT EDIT!

/*
Package resource_service is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/resource_service/resource_service.proto

It has these top-level messages:
	ContentHandleRequest
*/
package resource_service

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ContentHandleRequest struct {
	ProgramId        *uint32 `protobuf:"fixed32,1,req,name=program_id" json:"program_id,omitempty"`
	StreamId         *uint32 `protobuf:"fixed32,2,req,name=stream_id" json:"stream_id,omitempty"`
	Locale           *uint32 `protobuf:"fixed32,3,opt,name=locale,def=1701729619" json:"locale,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ContentHandleRequest) Reset()         { *m = ContentHandleRequest{} }
func (m *ContentHandleRequest) String() string { return proto.CompactTextString(m) }
func (*ContentHandleRequest) ProtoMessage()    {}

const Default_ContentHandleRequest_Locale uint32 = 1701729619

func (m *ContentHandleRequest) GetProgramId() uint32 {
	if m != nil && m.ProgramId != nil {
		return *m.ProgramId
	}
	return 0
}

func (m *ContentHandleRequest) GetStreamId() uint32 {
	if m != nil && m.StreamId != nil {
		return *m.StreamId
	}
	return 0
}

func (m *ContentHandleRequest) GetLocale() uint32 {
	if m != nil && m.Locale != nil {
		return *m.Locale
	}
	return Default_ContentHandleRequest_Locale
}
