// Code generated by protoc-gen-go. DO NOT EDIT.
// source: subscriberdb.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetMSISDNsRequest struct {
	// network_id of the subscriber
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// msisdns whose IMSIs should be retrieved
	// An empty list returns all tracked MSISDNs
	Msisdns              []string `protobuf:"bytes,2,rep,name=msisdns,proto3" json:"msisdns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMSISDNsRequest) Reset()         { *m = GetMSISDNsRequest{} }
func (m *GetMSISDNsRequest) String() string { return proto.CompactTextString(m) }
func (*GetMSISDNsRequest) ProtoMessage()    {}
func (*GetMSISDNsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{0}
}

func (m *GetMSISDNsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMSISDNsRequest.Unmarshal(m, b)
}
func (m *GetMSISDNsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMSISDNsRequest.Marshal(b, m, deterministic)
}
func (m *GetMSISDNsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMSISDNsRequest.Merge(m, src)
}
func (m *GetMSISDNsRequest) XXX_Size() int {
	return xxx_messageInfo_GetMSISDNsRequest.Size(m)
}
func (m *GetMSISDNsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMSISDNsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMSISDNsRequest proto.InternalMessageInfo

func (m *GetMSISDNsRequest) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *GetMSISDNsRequest) GetMsisdns() []string {
	if m != nil {
		return m.Msisdns
	}
	return nil
}

type GetMSISDNsResponse struct {
	// imsis_by_msisdn lists the requested imsis, keyed by their msisdn
	ImsisByMsisdn        map[string]string `protobuf:"bytes,1,rep,name=imsis_by_msisdn,json=imsisByMsisdn,proto3" json:"imsis_by_msisdn,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetMSISDNsResponse) Reset()         { *m = GetMSISDNsResponse{} }
func (m *GetMSISDNsResponse) String() string { return proto.CompactTextString(m) }
func (*GetMSISDNsResponse) ProtoMessage()    {}
func (*GetMSISDNsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{1}
}

func (m *GetMSISDNsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMSISDNsResponse.Unmarshal(m, b)
}
func (m *GetMSISDNsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMSISDNsResponse.Marshal(b, m, deterministic)
}
func (m *GetMSISDNsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMSISDNsResponse.Merge(m, src)
}
func (m *GetMSISDNsResponse) XXX_Size() int {
	return xxx_messageInfo_GetMSISDNsResponse.Size(m)
}
func (m *GetMSISDNsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMSISDNsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMSISDNsResponse proto.InternalMessageInfo

func (m *GetMSISDNsResponse) GetImsisByMsisdn() map[string]string {
	if m != nil {
		return m.ImsisByMsisdn
	}
	return nil
}

type SetMSISDNRequest struct {
	// network_id of the subscriber
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// msisdn to set
	Msisdn string `protobuf:"bytes,2,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	// imsi whose MSISDN should be set
	Imsi                 string   `protobuf:"bytes,3,opt,name=imsi,proto3" json:"imsi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetMSISDNRequest) Reset()         { *m = SetMSISDNRequest{} }
func (m *SetMSISDNRequest) String() string { return proto.CompactTextString(m) }
func (*SetMSISDNRequest) ProtoMessage()    {}
func (*SetMSISDNRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{2}
}

func (m *SetMSISDNRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetMSISDNRequest.Unmarshal(m, b)
}
func (m *SetMSISDNRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetMSISDNRequest.Marshal(b, m, deterministic)
}
func (m *SetMSISDNRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMSISDNRequest.Merge(m, src)
}
func (m *SetMSISDNRequest) XXX_Size() int {
	return xxx_messageInfo_SetMSISDNRequest.Size(m)
}
func (m *SetMSISDNRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMSISDNRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetMSISDNRequest proto.InternalMessageInfo

func (m *SetMSISDNRequest) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *SetMSISDNRequest) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

func (m *SetMSISDNRequest) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

type SetMSISDNResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetMSISDNResponse) Reset()         { *m = SetMSISDNResponse{} }
func (m *SetMSISDNResponse) String() string { return proto.CompactTextString(m) }
func (*SetMSISDNResponse) ProtoMessage()    {}
func (*SetMSISDNResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{3}
}

func (m *SetMSISDNResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetMSISDNResponse.Unmarshal(m, b)
}
func (m *SetMSISDNResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetMSISDNResponse.Marshal(b, m, deterministic)
}
func (m *SetMSISDNResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMSISDNResponse.Merge(m, src)
}
func (m *SetMSISDNResponse) XXX_Size() int {
	return xxx_messageInfo_SetMSISDNResponse.Size(m)
}
func (m *SetMSISDNResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMSISDNResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetMSISDNResponse proto.InternalMessageInfo

type DeleteMSISDNRequest struct {
	// network_id of the subscriber
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// msisdn to delete
	Msisdn               string   `protobuf:"bytes,2,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMSISDNRequest) Reset()         { *m = DeleteMSISDNRequest{} }
func (m *DeleteMSISDNRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMSISDNRequest) ProtoMessage()    {}
func (*DeleteMSISDNRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{4}
}

func (m *DeleteMSISDNRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMSISDNRequest.Unmarshal(m, b)
}
func (m *DeleteMSISDNRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMSISDNRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMSISDNRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMSISDNRequest.Merge(m, src)
}
func (m *DeleteMSISDNRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMSISDNRequest.Size(m)
}
func (m *DeleteMSISDNRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMSISDNRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMSISDNRequest proto.InternalMessageInfo

func (m *DeleteMSISDNRequest) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *DeleteMSISDNRequest) GetMsisdn() string {
	if m != nil {
		return m.Msisdn
	}
	return ""
}

type DeleteMSISDNResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMSISDNResponse) Reset()         { *m = DeleteMSISDNResponse{} }
func (m *DeleteMSISDNResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteMSISDNResponse) ProtoMessage()    {}
func (*DeleteMSISDNResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{5}
}

func (m *DeleteMSISDNResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMSISDNResponse.Unmarshal(m, b)
}
func (m *DeleteMSISDNResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMSISDNResponse.Marshal(b, m, deterministic)
}
func (m *DeleteMSISDNResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMSISDNResponse.Merge(m, src)
}
func (m *DeleteMSISDNResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteMSISDNResponse.Size(m)
}
func (m *DeleteMSISDNResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMSISDNResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMSISDNResponse proto.InternalMessageInfo

type GetIPsRequest struct {
	// network_id of the subscriber
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// ips whose IMSIs should be retrieved
	// An empty list returns all tracked IPs
	Ips                  []string `protobuf:"bytes,2,rep,name=ips,proto3" json:"ips,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIPsRequest) Reset()         { *m = GetIPsRequest{} }
func (m *GetIPsRequest) String() string { return proto.CompactTextString(m) }
func (*GetIPsRequest) ProtoMessage()    {}
func (*GetIPsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{6}
}

func (m *GetIPsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIPsRequest.Unmarshal(m, b)
}
func (m *GetIPsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIPsRequest.Marshal(b, m, deterministic)
}
func (m *GetIPsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIPsRequest.Merge(m, src)
}
func (m *GetIPsRequest) XXX_Size() int {
	return xxx_messageInfo_GetIPsRequest.Size(m)
}
func (m *GetIPsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIPsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIPsRequest proto.InternalMessageInfo

func (m *GetIPsRequest) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *GetIPsRequest) GetIps() []string {
	if m != nil {
		return m.Ips
	}
	return nil
}

type GetIPsResponse struct {
	// ip_mappings found
	IpMappings           []*IPMapping `protobuf:"bytes,1,rep,name=ip_mappings,json=ipMappings,proto3" json:"ip_mappings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetIPsResponse) Reset()         { *m = GetIPsResponse{} }
func (m *GetIPsResponse) String() string { return proto.CompactTextString(m) }
func (*GetIPsResponse) ProtoMessage()    {}
func (*GetIPsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{7}
}

func (m *GetIPsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIPsResponse.Unmarshal(m, b)
}
func (m *GetIPsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIPsResponse.Marshal(b, m, deterministic)
}
func (m *GetIPsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIPsResponse.Merge(m, src)
}
func (m *GetIPsResponse) XXX_Size() int {
	return xxx_messageInfo_GetIPsResponse.Size(m)
}
func (m *GetIPsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIPsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIPsResponse proto.InternalMessageInfo

func (m *GetIPsResponse) GetIpMappings() []*IPMapping {
	if m != nil {
		return m.IpMappings
	}
	return nil
}

type SetIPsRequest struct {
	// network_id of the subscriber
	NetworkId string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// ip_mappings to set
	IpMappings           []*IPMapping `protobuf:"bytes,2,rep,name=ip_mappings,json=ipMappings,proto3" json:"ip_mappings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SetIPsRequest) Reset()         { *m = SetIPsRequest{} }
func (m *SetIPsRequest) String() string { return proto.CompactTextString(m) }
func (*SetIPsRequest) ProtoMessage()    {}
func (*SetIPsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{8}
}

func (m *SetIPsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetIPsRequest.Unmarshal(m, b)
}
func (m *SetIPsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetIPsRequest.Marshal(b, m, deterministic)
}
func (m *SetIPsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetIPsRequest.Merge(m, src)
}
func (m *SetIPsRequest) XXX_Size() int {
	return xxx_messageInfo_SetIPsRequest.Size(m)
}
func (m *SetIPsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetIPsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetIPsRequest proto.InternalMessageInfo

func (m *SetIPsRequest) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *SetIPsRequest) GetIpMappings() []*IPMapping {
	if m != nil {
		return m.IpMappings
	}
	return nil
}

type SetIPsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetIPsResponse) Reset()         { *m = SetIPsResponse{} }
func (m *SetIPsResponse) String() string { return proto.CompactTextString(m) }
func (*SetIPsResponse) ProtoMessage()    {}
func (*SetIPsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{9}
}

func (m *SetIPsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetIPsResponse.Unmarshal(m, b)
}
func (m *SetIPsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetIPsResponse.Marshal(b, m, deterministic)
}
func (m *SetIPsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetIPsResponse.Merge(m, src)
}
func (m *SetIPsResponse) XXX_Size() int {
	return xxx_messageInfo_SetIPsResponse.Size(m)
}
func (m *SetIPsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetIPsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetIPsResponse proto.InternalMessageInfo

type IPMapping struct {
	// ip to set
	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	// imsi whose IP should be set
	Imsi string `protobuf:"bytes,2,opt,name=imsi,proto3" json:"imsi,omitempty"`
	// apn under which this IMSI is assigned the IP
	Apn                  string   `protobuf:"bytes,3,opt,name=apn,proto3" json:"apn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPMapping) Reset()         { *m = IPMapping{} }
func (m *IPMapping) String() string { return proto.CompactTextString(m) }
func (*IPMapping) ProtoMessage()    {}
func (*IPMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_7926c2bb91580e5a, []int{10}
}

func (m *IPMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPMapping.Unmarshal(m, b)
}
func (m *IPMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPMapping.Marshal(b, m, deterministic)
}
func (m *IPMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPMapping.Merge(m, src)
}
func (m *IPMapping) XXX_Size() int {
	return xxx_messageInfo_IPMapping.Size(m)
}
func (m *IPMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_IPMapping.DiscardUnknown(m)
}

var xxx_messageInfo_IPMapping proto.InternalMessageInfo

func (m *IPMapping) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *IPMapping) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *IPMapping) GetApn() string {
	if m != nil {
		return m.Apn
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMSISDNsRequest)(nil), "magma.lte.subscriberdb.GetMSISDNsRequest")
	proto.RegisterType((*GetMSISDNsResponse)(nil), "magma.lte.subscriberdb.GetMSISDNsResponse")
	proto.RegisterMapType((map[string]string)(nil), "magma.lte.subscriberdb.GetMSISDNsResponse.ImsisByMsisdnEntry")
	proto.RegisterType((*SetMSISDNRequest)(nil), "magma.lte.subscriberdb.SetMSISDNRequest")
	proto.RegisterType((*SetMSISDNResponse)(nil), "magma.lte.subscriberdb.SetMSISDNResponse")
	proto.RegisterType((*DeleteMSISDNRequest)(nil), "magma.lte.subscriberdb.DeleteMSISDNRequest")
	proto.RegisterType((*DeleteMSISDNResponse)(nil), "magma.lte.subscriberdb.DeleteMSISDNResponse")
	proto.RegisterType((*GetIPsRequest)(nil), "magma.lte.subscriberdb.GetIPsRequest")
	proto.RegisterType((*GetIPsResponse)(nil), "magma.lte.subscriberdb.GetIPsResponse")
	proto.RegisterType((*SetIPsRequest)(nil), "magma.lte.subscriberdb.SetIPsRequest")
	proto.RegisterType((*SetIPsResponse)(nil), "magma.lte.subscriberdb.SetIPsResponse")
	proto.RegisterType((*IPMapping)(nil), "magma.lte.subscriberdb.IPMapping")
}

func init() { proto.RegisterFile("subscriberdb.proto", fileDescriptor_7926c2bb91580e5a) }

var fileDescriptor_7926c2bb91580e5a = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0xdd, 0x24, 0x6b, 0x35, 0x77, 0x6d, 0xcd, 0xde, 0x5d, 0x4a, 0x08, 0x08, 0x75, 0x40, 0xe9,
	0xaa, 0xe4, 0x61, 0x7d, 0x11, 0x41, 0x58, 0xcb, 0xca, 0x12, 0x68, 0x65, 0xc9, 0xf8, 0xa2, 0x20,
	0x25, 0xb1, 0xc3, 0x32, 0xb4, 0x4d, 0xc6, 0xcc, 0x54, 0xe9, 0x9b, 0x3f, 0xe6, 0xbf, 0x49, 0x92,
	0x49, 0x9a, 0xed, 0x6e, 0xdd, 0x28, 0x3e, 0xf5, 0xce, 0xf4, 0xdc, 0x73, 0xe6, 0xde, 0x73, 0x5a,
	0x40, 0xb9, 0x8a, 0xe5, 0xd7, 0x8c, 0xc7, 0x2c, 0x9b, 0xc5, 0xbe, 0xc8, 0x52, 0x95, 0x62, 0x7f,
	0x19, 0x5d, 0x2d, 0x23, 0x7f, 0xa1, 0x98, 0xdf, 0xfc, 0x96, 0x8c, 0xe1, 0xf0, 0x82, 0xa9, 0x09,
	0x0d, 0xe8, 0xf9, 0x07, 0x19, 0xb2, 0x6f, 0x2b, 0x26, 0x15, 0x3e, 0x06, 0x48, 0x98, 0xfa, 0x91,
	0x66, 0xf3, 0x29, 0x9f, 0xb9, 0xc6, 0xc0, 0x18, 0xda, 0xa1, 0xad, 0x6f, 0x82, 0x19, 0xba, 0x70,
	0x7f, 0x29, 0xb9, 0x9c, 0x25, 0xd2, 0x35, 0x07, 0xd6, 0xd0, 0x0e, 0xab, 0x23, 0xf9, 0x65, 0x00,
	0x36, 0xe9, 0xa4, 0x48, 0x13, 0xc9, 0x90, 0xc1, 0x23, 0x9e, 0x43, 0xa6, 0xf1, 0x7a, 0x5a, 0x42,
	0x5d, 0x63, 0x60, 0x0d, 0x0f, 0x4e, 0xdf, 0xfa, 0xb7, 0x3f, 0xcb, 0xbf, 0x49, 0xe2, 0x07, 0x79,
	0xe7, 0x68, 0x3d, 0x29, 0xfa, 0xdf, 0x27, 0x2a, 0x5b, 0x87, 0x5d, 0xde, 0xbc, 0xf3, 0xce, 0x00,
	0x6f, 0x82, 0xd0, 0x01, 0x6b, 0xce, 0xd6, 0x7a, 0x8a, 0xbc, 0xc4, 0x63, 0xb8, 0xf7, 0x3d, 0x5a,
	0xac, 0x98, 0x6b, 0x16, 0x77, 0xe5, 0xe1, 0x8d, 0xf9, 0xda, 0x20, 0x5f, 0xc0, 0xa1, 0x95, 0x72,
	0xcb, 0x65, 0xf4, 0xa1, 0xa3, 0x47, 0x2a, 0xd9, 0xf4, 0x09, 0x11, 0xf6, 0xf3, 0xd7, 0xb9, 0x56,
	0x71, 0x5b, 0xd4, 0xe4, 0x08, 0x0e, 0x1b, 0xf4, 0xe5, 0x5c, 0x64, 0x0c, 0x47, 0xe7, 0x6c, 0xc1,
	0x14, 0xfb, 0x1f, 0xb2, 0xa4, 0x0f, 0xc7, 0xd7, 0xd9, 0xb4, 0xca, 0x19, 0x74, 0x2f, 0x98, 0x0a,
	0x2e, 0xdb, 0x7a, 0xec, 0x80, 0xc5, 0x45, 0xe5, 0x6f, 0x5e, 0x92, 0x8f, 0xd0, 0xab, 0x18, 0xb4,
	0xad, 0x23, 0x38, 0xe0, 0x62, 0xba, 0x8c, 0x84, 0xe0, 0xc9, 0x95, 0xd4, 0x96, 0x3e, 0xd9, 0x65,
	0x69, 0x70, 0x39, 0x29, 0x91, 0x21, 0x70, 0xa1, 0x4b, 0x49, 0x32, 0xe8, 0xd2, 0xbf, 0x79, 0xd7,
	0x96, 0xa6, 0xf9, 0x2f, 0x9a, 0x0e, 0xf4, 0xe8, 0xb5, 0x49, 0xc8, 0x3b, 0xb0, 0x6b, 0x28, 0xf6,
	0xc0, 0xe4, 0x42, 0x2b, 0x9b, 0x5c, 0xd4, 0x4e, 0x9a, 0x1b, 0x27, 0xf3, 0xf5, 0x44, 0x22, 0xd1,
	0xe6, 0xe6, 0xe5, 0xe9, 0xcf, 0x7d, 0x70, 0x68, 0xad, 0x3d, 0x4e, 0xd3, 0xf9, 0x4a, 0x20, 0x03,
	0xd8, 0x24, 0x19, 0x4f, 0xda, 0xa4, 0xbd, 0xd8, 0x82, 0xf7, 0xbc, 0xfd, 0x0f, 0x83, 0xec, 0x61,
	0x0c, 0x76, 0x9d, 0x2b, 0x1c, 0xee, 0x6a, 0xdd, 0x4e, 0xb6, 0x77, 0xd2, 0x02, 0x59, 0x6b, 0xcc,
	0xe1, 0x61, 0x33, 0x58, 0xf8, 0x62, 0x57, 0xf3, 0x2d, 0x61, 0xf6, 0x5e, 0xb6, 0x03, 0xd7, 0x62,
	0x9f, 0xa0, 0x53, 0x66, 0x0d, 0x9f, 0xfe, 0x61, 0x11, 0x9b, 0xd4, 0x78, 0xcf, 0xee, 0x82, 0x35,
	0xa9, 0xe9, 0x1d, 0xd4, 0xb4, 0x1d, 0xf5, 0x56, 0x86, 0xf6, 0x46, 0x0f, 0x3e, 0x77, 0x8a, 0x3f,
	0x5b, 0x19, 0x97, 0x9f, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x86, 0x0c, 0x92, 0x79, 0x8a,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SubscriberLookupClient is the client API for SubscriberLookup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscriberLookupClient interface {
	// GetMSISDNs returns MSISDN -> IMSI mappings.
	GetMSISDNs(ctx context.Context, in *GetMSISDNsRequest, opts ...grpc.CallOption) (*GetMSISDNsResponse, error)
	// SetMSISDN creates a MSISDN -> IMSI mapping.
	// Error if MSISDN has already been assigned.
	SetMSISDN(ctx context.Context, in *SetMSISDNRequest, opts ...grpc.CallOption) (*SetMSISDNResponse, error)
	// DeleteMSISDN removes the MSISDN -> IMSI mapping.
	DeleteMSISDN(ctx context.Context, in *DeleteMSISDNRequest, opts ...grpc.CallOption) (*DeleteMSISDNResponse, error)
	// GetIPs returns IP -> IMSI mappings.
	GetIPs(ctx context.Context, in *GetIPsRequest, opts ...grpc.CallOption) (*GetIPsResponse, error)
	// SetIPs creates an IP -> IMSI mapping.
	// Error if IP has already been assigned.
	SetIPs(ctx context.Context, in *SetIPsRequest, opts ...grpc.CallOption) (*SetIPsResponse, error)
}

type subscriberLookupClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriberLookupClient(cc grpc.ClientConnInterface) SubscriberLookupClient {
	return &subscriberLookupClient{cc}
}

func (c *subscriberLookupClient) GetMSISDNs(ctx context.Context, in *GetMSISDNsRequest, opts ...grpc.CallOption) (*GetMSISDNsResponse, error) {
	out := new(GetMSISDNsResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.subscriberdb.SubscriberLookup/GetMSISDNs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberLookupClient) SetMSISDN(ctx context.Context, in *SetMSISDNRequest, opts ...grpc.CallOption) (*SetMSISDNResponse, error) {
	out := new(SetMSISDNResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.subscriberdb.SubscriberLookup/SetMSISDN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberLookupClient) DeleteMSISDN(ctx context.Context, in *DeleteMSISDNRequest, opts ...grpc.CallOption) (*DeleteMSISDNResponse, error) {
	out := new(DeleteMSISDNResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.subscriberdb.SubscriberLookup/DeleteMSISDN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberLookupClient) GetIPs(ctx context.Context, in *GetIPsRequest, opts ...grpc.CallOption) (*GetIPsResponse, error) {
	out := new(GetIPsResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.subscriberdb.SubscriberLookup/GetIPs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriberLookupClient) SetIPs(ctx context.Context, in *SetIPsRequest, opts ...grpc.CallOption) (*SetIPsResponse, error) {
	out := new(SetIPsResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.subscriberdb.SubscriberLookup/SetIPs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriberLookupServer is the server API for SubscriberLookup service.
type SubscriberLookupServer interface {
	// GetMSISDNs returns MSISDN -> IMSI mappings.
	GetMSISDNs(context.Context, *GetMSISDNsRequest) (*GetMSISDNsResponse, error)
	// SetMSISDN creates a MSISDN -> IMSI mapping.
	// Error if MSISDN has already been assigned.
	SetMSISDN(context.Context, *SetMSISDNRequest) (*SetMSISDNResponse, error)
	// DeleteMSISDN removes the MSISDN -> IMSI mapping.
	DeleteMSISDN(context.Context, *DeleteMSISDNRequest) (*DeleteMSISDNResponse, error)
	// GetIPs returns IP -> IMSI mappings.
	GetIPs(context.Context, *GetIPsRequest) (*GetIPsResponse, error)
	// SetIPs creates an IP -> IMSI mapping.
	// Error if IP has already been assigned.
	SetIPs(context.Context, *SetIPsRequest) (*SetIPsResponse, error)
}

// UnimplementedSubscriberLookupServer can be embedded to have forward compatible implementations.
type UnimplementedSubscriberLookupServer struct {
}

func (*UnimplementedSubscriberLookupServer) GetMSISDNs(ctx context.Context, req *GetMSISDNsRequest) (*GetMSISDNsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMSISDNs not implemented")
}
func (*UnimplementedSubscriberLookupServer) SetMSISDN(ctx context.Context, req *SetMSISDNRequest) (*SetMSISDNResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMSISDN not implemented")
}
func (*UnimplementedSubscriberLookupServer) DeleteMSISDN(ctx context.Context, req *DeleteMSISDNRequest) (*DeleteMSISDNResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMSISDN not implemented")
}
func (*UnimplementedSubscriberLookupServer) GetIPs(ctx context.Context, req *GetIPsRequest) (*GetIPsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIPs not implemented")
}
func (*UnimplementedSubscriberLookupServer) SetIPs(ctx context.Context, req *SetIPsRequest) (*SetIPsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIPs not implemented")
}

func RegisterSubscriberLookupServer(s *grpc.Server, srv SubscriberLookupServer) {
	s.RegisterService(&_SubscriberLookup_serviceDesc, srv)
}

func _SubscriberLookup_GetMSISDNs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMSISDNsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberLookupServer).GetMSISDNs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.subscriberdb.SubscriberLookup/GetMSISDNs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberLookupServer).GetMSISDNs(ctx, req.(*GetMSISDNsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriberLookup_SetMSISDN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMSISDNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberLookupServer).SetMSISDN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.subscriberdb.SubscriberLookup/SetMSISDN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberLookupServer).SetMSISDN(ctx, req.(*SetMSISDNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriberLookup_DeleteMSISDN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMSISDNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberLookupServer).DeleteMSISDN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.subscriberdb.SubscriberLookup/DeleteMSISDN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberLookupServer).DeleteMSISDN(ctx, req.(*DeleteMSISDNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriberLookup_GetIPs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIPsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberLookupServer).GetIPs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.subscriberdb.SubscriberLookup/GetIPs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberLookupServer).GetIPs(ctx, req.(*GetIPsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriberLookup_SetIPs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetIPsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriberLookupServer).SetIPs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.subscriberdb.SubscriberLookup/SetIPs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriberLookupServer).SetIPs(ctx, req.(*SetIPsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SubscriberLookup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.subscriberdb.SubscriberLookup",
	HandlerType: (*SubscriberLookupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMSISDNs",
			Handler:    _SubscriberLookup_GetMSISDNs_Handler,
		},
		{
			MethodName: "SetMSISDN",
			Handler:    _SubscriberLookup_SetMSISDN_Handler,
		},
		{
			MethodName: "DeleteMSISDN",
			Handler:    _SubscriberLookup_DeleteMSISDN_Handler,
		},
		{
			MethodName: "GetIPs",
			Handler:    _SubscriberLookup_GetIPs_Handler,
		},
		{
			MethodName: "SetIPs",
			Handler:    _SubscriberLookup_SetIPs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscriberdb.proto",
}
