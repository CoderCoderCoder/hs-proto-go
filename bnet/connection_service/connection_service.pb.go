// Code generated by protoc-gen-go.
// source: github.com/HearthSim/hs-proto-go/bnet/connection_service/connection_service.proto
// DO NOT EDIT!

/*
Package connection_service is a generated protocol buffer package.

It is generated from these files:
	github.com/HearthSim/hs-proto-go/bnet/connection_service/connection_service.proto

It has these top-level messages:
	ConnectRequest
	ConnectionMeteringContentHandles
	ConnectResponse
	BoundService
	BindRequest
	BindResponse
	EchoRequest
	EchoResponse
	DisconnectRequest
	DisconnectNotification
	EncryptRequest
*/
package connection_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import content_handle "github.com/HearthSim/hs-proto-go/bnet/content_handle"
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

type ConnectRequest struct {
	ClientId         *rpc.ProcessId `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	BindRequest      *BindRequest   `protobuf:"bytes,2,opt,name=bind_request,json=bindRequest" json:"bind_request,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ConnectRequest) Reset()                    { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string            { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()               {}
func (*ConnectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConnectRequest) GetClientId() *rpc.ProcessId {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *ConnectRequest) GetBindRequest() *BindRequest {
	if m != nil {
		return m.BindRequest
	}
	return nil
}

type ConnectionMeteringContentHandles struct {
	ContentHandle    []*content_handle.ContentHandle `protobuf:"bytes,1,rep,name=content_handle,json=contentHandle" json:"content_handle,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *ConnectionMeteringContentHandles) Reset()         { *m = ConnectionMeteringContentHandles{} }
func (m *ConnectionMeteringContentHandles) String() string { return proto.CompactTextString(m) }
func (*ConnectionMeteringContentHandles) ProtoMessage()    {}
func (*ConnectionMeteringContentHandles) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1}
}

func (m *ConnectionMeteringContentHandles) GetContentHandle() []*content_handle.ContentHandle {
	if m != nil {
		return m.ContentHandle
	}
	return nil
}

type ConnectResponse struct {
	ServerId           *rpc.ProcessId                    `protobuf:"bytes,1,req,name=server_id,json=serverId" json:"server_id,omitempty"`
	ClientId           *rpc.ProcessId                    `protobuf:"bytes,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	BindResult         *uint32                           `protobuf:"varint,3,opt,name=bind_result,json=bindResult" json:"bind_result,omitempty"`
	BindResponse       *BindResponse                     `protobuf:"bytes,4,opt,name=bind_response,json=bindResponse" json:"bind_response,omitempty"`
	ContentHandleArray *ConnectionMeteringContentHandles `protobuf:"bytes,5,opt,name=content_handle_array,json=contentHandleArray" json:"content_handle_array,omitempty"`
	ServerTime         *uint64                           `protobuf:"varint,6,opt,name=server_time,json=serverTime" json:"server_time,omitempty"`
	XXX_unrecognized   []byte                            `json:"-"`
}

func (m *ConnectResponse) Reset()                    { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string            { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()               {}
func (*ConnectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ConnectResponse) GetServerId() *rpc.ProcessId {
	if m != nil {
		return m.ServerId
	}
	return nil
}

func (m *ConnectResponse) GetClientId() *rpc.ProcessId {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *ConnectResponse) GetBindResult() uint32 {
	if m != nil && m.BindResult != nil {
		return *m.BindResult
	}
	return 0
}

func (m *ConnectResponse) GetBindResponse() *BindResponse {
	if m != nil {
		return m.BindResponse
	}
	return nil
}

func (m *ConnectResponse) GetContentHandleArray() *ConnectionMeteringContentHandles {
	if m != nil {
		return m.ContentHandleArray
	}
	return nil
}

func (m *ConnectResponse) GetServerTime() uint64 {
	if m != nil && m.ServerTime != nil {
		return *m.ServerTime
	}
	return 0
}

type BoundService struct {
	Hash             *uint32 `protobuf:"fixed32,1,req,name=hash" json:"hash,omitempty"`
	Id               *uint32 `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BoundService) Reset()                    { *m = BoundService{} }
func (m *BoundService) String() string            { return proto.CompactTextString(m) }
func (*BoundService) ProtoMessage()               {}
func (*BoundService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BoundService) GetHash() uint32 {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return 0
}

func (m *BoundService) GetId() uint32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type BindRequest struct {
	ImportedServiceHash []uint32        `protobuf:"fixed32,1,rep,name=imported_service_hash,json=importedServiceHash" json:"imported_service_hash,omitempty"`
	ExportedService     []*BoundService `protobuf:"bytes,2,rep,name=exported_service,json=exportedService" json:"exported_service,omitempty"`
	XXX_unrecognized    []byte          `json:"-"`
}

func (m *BindRequest) Reset()                    { *m = BindRequest{} }
func (m *BindRequest) String() string            { return proto.CompactTextString(m) }
func (*BindRequest) ProtoMessage()               {}
func (*BindRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BindRequest) GetImportedServiceHash() []uint32 {
	if m != nil {
		return m.ImportedServiceHash
	}
	return nil
}

func (m *BindRequest) GetExportedService() []*BoundService {
	if m != nil {
		return m.ExportedService
	}
	return nil
}

type BindResponse struct {
	ImportedServiceId []uint32 `protobuf:"varint,1,rep,packed,name=imported_service_id,json=importedServiceId" json:"imported_service_id,omitempty"`
	XXX_unrecognized  []byte   `json:"-"`
}

func (m *BindResponse) Reset()                    { *m = BindResponse{} }
func (m *BindResponse) String() string            { return proto.CompactTextString(m) }
func (*BindResponse) ProtoMessage()               {}
func (*BindResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BindResponse) GetImportedServiceId() []uint32 {
	if m != nil {
		return m.ImportedServiceId
	}
	return nil
}

type EchoRequest struct {
	Time             *uint64 `protobuf:"fixed64,1,opt,name=time" json:"time,omitempty"`
	NetworkOnly      *bool   `protobuf:"varint,2,opt,name=network_only,json=networkOnly,def=0" json:"network_only,omitempty"`
	Payload          []byte  `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

const Default_EchoRequest_NetworkOnly bool = false

func (m *EchoRequest) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *EchoRequest) GetNetworkOnly() bool {
	if m != nil && m.NetworkOnly != nil {
		return *m.NetworkOnly
	}
	return Default_EchoRequest_NetworkOnly
}

func (m *EchoRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type EchoResponse struct {
	Time             *uint64 `protobuf:"fixed64,1,opt,name=time" json:"time,omitempty"`
	Payload          []byte  `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *EchoResponse) GetTime() uint64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *EchoResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type DisconnectRequest struct {
	ErrorCode        *uint32 `protobuf:"varint,1,req,name=error_code,json=errorCode" json:"error_code,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DisconnectRequest) Reset()                    { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string            { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()               {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DisconnectRequest) GetErrorCode() uint32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

type DisconnectNotification struct {
	ErrorCode        *uint32 `protobuf:"varint,1,req,name=error_code,json=errorCode" json:"error_code,omitempty"`
	Reason           *string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DisconnectNotification) Reset()                    { *m = DisconnectNotification{} }
func (m *DisconnectNotification) String() string            { return proto.CompactTextString(m) }
func (*DisconnectNotification) ProtoMessage()               {}
func (*DisconnectNotification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DisconnectNotification) GetErrorCode() uint32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *DisconnectNotification) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

type EncryptRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *EncryptRequest) Reset()                    { *m = EncryptRequest{} }
func (m *EncryptRequest) String() string            { return proto.CompactTextString(m) }
func (*EncryptRequest) ProtoMessage()               {}
func (*EncryptRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*ConnectRequest)(nil), "connection_service.ConnectRequest")
	proto.RegisterType((*ConnectionMeteringContentHandles)(nil), "connection_service.ConnectionMeteringContentHandles")
	proto.RegisterType((*ConnectResponse)(nil), "connection_service.ConnectResponse")
	proto.RegisterType((*BoundService)(nil), "connection_service.BoundService")
	proto.RegisterType((*BindRequest)(nil), "connection_service.BindRequest")
	proto.RegisterType((*BindResponse)(nil), "connection_service.BindResponse")
	proto.RegisterType((*EchoRequest)(nil), "connection_service.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "connection_service.EchoResponse")
	proto.RegisterType((*DisconnectRequest)(nil), "connection_service.DisconnectRequest")
	proto.RegisterType((*DisconnectNotification)(nil), "connection_service.DisconnectNotification")
	proto.RegisterType((*EncryptRequest)(nil), "connection_service.EncryptRequest")
}

func init() {
	proto.RegisterFile("github.com/HearthSim/hs-proto-go/bnet/connection_service/connection_service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xdb, 0x3a,
	0x10, 0x84, 0x64, 0xe7, 0x6b, 0x25, 0x3b, 0x09, 0xf3, 0x5e, 0x20, 0x04, 0x08, 0x62, 0xe8, 0x24,
	0xe0, 0x21, 0x36, 0xa0, 0xd7, 0x53, 0xd1, 0x4b, 0xed, 0x04, 0x88, 0x51, 0xb4, 0x69, 0x99, 0xde,
	0x05, 0x99, 0x64, 0x22, 0xa2, 0x32, 0xe9, 0x92, 0x74, 0x5b, 0x1f, 0xfb, 0x07, 0xfa, 0x0b, 0xfb,
	0x63, 0x0a, 0x91, 0x52, 0x62, 0x25, 0x2e, 0x92, 0x83, 0x01, 0xed, 0x6a, 0x67, 0x76, 0x66, 0xbc,
	0x82, 0x4f, 0x77, 0xdc, 0x14, 0xcb, 0xd9, 0x90, 0xc8, 0xf9, 0xe8, 0x8a, 0xe5, 0xca, 0x14, 0x37,
	0x7c, 0x3e, 0x2a, 0xf4, 0xf9, 0x42, 0x49, 0x23, 0xcf, 0xef, 0xe4, 0x68, 0x26, 0x98, 0x19, 0x11,
	0x29, 0x04, 0x23, 0x86, 0x4b, 0x91, 0x69, 0xa6, 0xbe, 0x71, 0xc2, 0x36, 0xb4, 0x86, 0x16, 0x83,
	0xd0, 0xd3, 0x37, 0x27, 0xd3, 0x17, 0xaf, 0x31, 0x4c, 0x98, 0xac, 0xc8, 0x05, 0x2d, 0xd9, 0xa3,
	0xd2, 0xd1, 0x9f, 0xfc, 0xff, 0x32, 0x2a, 0xb5, 0x20, 0xd5, 0xcf, 0x81, 0xe2, 0x9f, 0x1e, 0xf4,
	0x27, 0x4e, 0x16, 0x66, 0x5f, 0x97, 0x4c, 0x1b, 0xf4, 0x1f, 0xec, 0x91, 0x92, 0x57, 0xf4, 0x9c,
	0x46, 0xde, 0xc0, 0x4b, 0x82, 0xb4, 0x3f, 0xac, 0x10, 0x1f, 0x95, 0x24, 0x4c, 0xeb, 0x29, 0xc5,
	0xbb, 0x6e, 0x60, 0x4a, 0xd1, 0x18, 0xc2, 0x19, 0x17, 0x34, 0x53, 0x0e, 0x1c, 0xf9, 0x76, 0xfe,
	0x6c, 0xb8, 0x21, 0x84, 0x31, 0x17, 0xb4, 0xde, 0x81, 0x83, 0xd9, 0x43, 0x11, 0x17, 0x30, 0x98,
	0xdc, 0x8f, 0xbf, 0x67, 0x86, 0x29, 0x2e, 0xee, 0x26, 0xce, 0xe2, 0x95, 0x75, 0xa8, 0xd1, 0x05,
	0xf4, 0xdb, 0xa6, 0x23, 0x6f, 0xd0, 0x49, 0x82, 0xf4, 0x74, 0xf8, 0x28, 0x8b, 0x16, 0x0e, 0xf7,
	0xc8, 0x7a, 0x19, 0xff, 0xf6, 0x61, 0xff, 0xde, 0xad, 0x5e, 0x48, 0xa1, 0x59, 0x65, 0xb7, 0x52,
	0xc8, 0x94, 0xb3, 0xeb, 0x6f, 0xb2, 0xeb, 0x06, 0xa6, 0xb4, 0x9d, 0x8d, 0xff, 0x4c, 0x36, 0x67,
	0x10, 0xd4, 0xd9, 0xe8, 0x65, 0x69, 0xa2, 0xce, 0xc0, 0x4b, 0x7a, 0x18, 0x9c, 0xf3, 0xaa, 0x83,
	0x2e, 0xa1, 0xd7, 0x0c, 0x58, 0x2d, 0x51, 0xd7, 0x32, 0x0e, 0xfe, 0x9e, 0x9e, 0x9b, 0xc3, 0xe1,
	0x6c, 0xad, 0x42, 0xb7, 0xf0, 0x4f, 0x3b, 0x84, 0x2c, 0x57, 0x2a, 0x5f, 0x45, 0x5b, 0x96, 0xed,
	0xd5, 0x26, 0xb6, 0xe7, 0xf2, 0xc6, 0xa8, 0x15, 0xdc, 0xdb, 0x8a, 0xaf, 0xf2, 0x53, 0x27, 0x65,
	0xf8, 0x9c, 0x45, 0xdb, 0x03, 0x2f, 0xe9, 0x62, 0x70, 0xad, 0xcf, 0x7c, 0xce, 0xe2, 0x14, 0xc2,
	0xb1, 0x5c, 0x0a, 0x7a, 0xe3, 0xb6, 0x20, 0x04, 0xdd, 0x22, 0xd7, 0x85, 0x4d, 0x75, 0x07, 0xdb,
	0x67, 0xd4, 0x07, 0xdf, 0x46, 0xe7, 0x27, 0x3d, 0xec, 0x73, 0x1a, 0xff, 0xf2, 0x20, 0x58, 0xbb,
	0x0c, 0x94, 0xc2, 0xbf, 0x7c, 0xbe, 0x90, 0xca, 0x30, 0xda, 0xa8, 0xcd, 0x6a, 0x92, 0x4e, 0xb2,
	0x83, 0x8f, 0x9a, 0x97, 0xf5, 0x8e, 0xab, 0x8a, 0xf3, 0x1d, 0x1c, 0xb0, 0x1f, 0x6d, 0x4c, 0xe4,
	0xdb, 0xf3, 0xd8, 0x1c, 0xe5, 0x9a, 0x46, 0xbc, 0xdf, 0x20, 0xeb, 0x46, 0x3c, 0x86, 0x70, 0x3d,
	0x6b, 0x94, 0xc2, 0xd1, 0x13, 0x41, 0xf6, 0x52, 0x3a, 0x49, 0x6f, 0xec, 0x1f, 0x78, 0xf8, 0xf0,
	0x91, 0xa4, 0x29, 0x8d, 0x19, 0x04, 0x97, 0xa4, 0x90, 0x8d, 0x27, 0x04, 0x5d, 0x9b, 0x58, 0xf5,
	0x31, 0x6d, 0x63, 0xfb, 0x8c, 0x12, 0x08, 0x05, 0x33, 0xdf, 0xa5, 0xfa, 0x92, 0x49, 0x51, 0xae,
	0xec, 0x31, 0xed, 0xbe, 0xde, 0xba, 0xcd, 0x4b, 0xcd, 0x70, 0x50, 0xbf, 0xba, 0x16, 0xe5, 0x0a,
	0x45, 0xb0, 0xb3, 0xc8, 0x57, 0xa5, 0xcc, 0xa9, 0x3d, 0xa1, 0x10, 0x37, 0x65, 0xfc, 0x06, 0x42,
	0xb7, 0xa6, 0x96, 0xba, 0x69, 0xcf, 0x1a, 0xda, 0x6f, 0xa3, 0x53, 0x38, 0xbc, 0xe0, 0x9a, 0xb4,
	0x3f, 0xfe, 0x53, 0x00, 0xa6, 0x94, 0x54, 0x19, 0x91, 0x94, 0xd9, 0x3f, 0xae, 0x87, 0xf7, 0x6c,
	0x67, 0x22, 0x29, 0x8b, 0xaf, 0xe1, 0xf8, 0x01, 0xf3, 0x41, 0x1a, 0x7e, 0xcb, 0x49, 0x5e, 0x85,
	0xfb, 0x0c, 0x10, 0x1d, 0xc3, 0xb6, 0x62, 0xb9, 0x96, 0xc2, 0xaa, 0xd8, 0xc3, 0x75, 0x15, 0x1f,
	0x40, 0xff, 0x52, 0x10, 0xb5, 0x5a, 0x34, 0x0a, 0xd2, 0x23, 0x38, 0x7c, 0xb8, 0xce, 0x3a, 0xd2,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x66, 0xe6, 0xd5, 0x8e, 0x05, 0x00, 0x00,
}
