// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/sample_go_contract/sample_go_contract.proto

package sample_go_contract

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SampleGoContractInitRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleGoContractInitRequest) Reset()         { *m = SampleGoContractInitRequest{} }
func (m *SampleGoContractInitRequest) String() string { return proto.CompactTextString(m) }
func (*SampleGoContractInitRequest) ProtoMessage()    {}
func (*SampleGoContractInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sample_go_contract_52cdfc5560053bb3, []int{0}
}
func (m *SampleGoContractInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleGoContractInitRequest.Unmarshal(m, b)
}
func (m *SampleGoContractInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleGoContractInitRequest.Marshal(b, m, deterministic)
}
func (dst *SampleGoContractInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleGoContractInitRequest.Merge(dst, src)
}
func (m *SampleGoContractInitRequest) XXX_Size() int {
	return xxx_messageInfo_SampleGoContractInitRequest.Size(m)
}
func (m *SampleGoContractInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleGoContractInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SampleGoContractInitRequest proto.InternalMessageInfo

type SampleGoContractNestedEvmRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SampleGoContractNestedEvmRequest) Reset()         { *m = SampleGoContractNestedEvmRequest{} }
func (m *SampleGoContractNestedEvmRequest) String() string { return proto.CompactTextString(m) }
func (*SampleGoContractNestedEvmRequest) ProtoMessage()    {}
func (*SampleGoContractNestedEvmRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_sample_go_contract_52cdfc5560053bb3, []int{1}
}
func (m *SampleGoContractNestedEvmRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SampleGoContractNestedEvmRequest.Unmarshal(m, b)
}
func (m *SampleGoContractNestedEvmRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SampleGoContractNestedEvmRequest.Marshal(b, m, deterministic)
}
func (dst *SampleGoContractNestedEvmRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SampleGoContractNestedEvmRequest.Merge(dst, src)
}
func (m *SampleGoContractNestedEvmRequest) XXX_Size() int {
	return xxx_messageInfo_SampleGoContractNestedEvmRequest.Size(m)
}
func (m *SampleGoContractNestedEvmRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SampleGoContractNestedEvmRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SampleGoContractNestedEvmRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SampleGoContractInitRequest)(nil), "SampleGoContractInitRequest")
	proto.RegisterType((*SampleGoContractNestedEvmRequest)(nil), "SampleGoContractNestedEvmRequest")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/sample_go_contract/sample_go_contract.proto", fileDescriptor_sample_go_contract_52cdfc5560053bb3)
}

var fileDescriptor_sample_go_contract_52cdfc5560053bb3 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x8a, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xc9, 0xcf, 0xcf, 0xcd, 0x4b, 0x2d, 0x29, 0xcf,
	0x2f, 0xca, 0xd6, 0x4f, 0xcf, 0xd7, 0x05, 0x71, 0xf5, 0x93, 0x4a, 0x33, 0x73, 0x4a, 0x32, 0xf3,
	0xf4, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x8b, 0x13, 0x73, 0x0b, 0x72, 0x52, 0xe3, 0xd3, 0xf3,
	0xe3, 0x93, 0xf3, 0xf3, 0x4a, 0x8a, 0x12, 0x93, 0x4b, 0xb0, 0x08, 0xe9, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x4b, 0x19, 0x10, 0x30, 0x19, 0x62, 0x22, 0x98, 0x84, 0xe8, 0x50, 0x92, 0xe5, 0x92, 0x0e,
	0x06, 0x9b, 0xe6, 0x9e, 0xef, 0x0c, 0x35, 0xcb, 0x33, 0x2f, 0xb3, 0x24, 0x28, 0xb5, 0xb0, 0x34,
	0xb5, 0xb8, 0x44, 0x49, 0x89, 0x4b, 0x01, 0x5d, 0xda, 0x2f, 0xb5, 0xb8, 0x24, 0x35, 0xc5, 0xb5,
	0x2c, 0x17, 0xaa, 0x26, 0x89, 0x0d, 0x6c, 0x92, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x48,
	0x93, 0xf8, 0xd7, 0x00, 0x00, 0x00,
}