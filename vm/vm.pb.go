// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/vm/vm.proto

/*
Package vm is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/vm/vm.proto

It has these top-level messages:
	MessageTx
	DeployTx
	CallTx
	DeployResponse
	DeployResponseData
	TxHashList
*/
package vm

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type VMType int32

const (
	VMType_PLUGIN VMType = 0
	VMType_EVM    VMType = 1
)

var VMType_name = map[int32]string{
	0: "PLUGIN",
	1: "EVM",
}
var VMType_value = map[string]int32{
	"PLUGIN": 0,
	"EVM":    1,
}

func (x VMType) String() string {
	return proto.EnumName(VMType_name, int32(x))
}
func (VMType) EnumDescriptor() ([]byte, []int) { return fileDescriptorVm, []int{0} }

type MessageTx struct {
	To   *types.Address `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	From *types.Address `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	Data []byte         `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MessageTx) Reset()                    { *m = MessageTx{} }
func (m *MessageTx) String() string            { return proto.CompactTextString(m) }
func (*MessageTx) ProtoMessage()               {}
func (*MessageTx) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{0} }

func (m *MessageTx) GetTo() *types.Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *MessageTx) GetFrom() *types.Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *MessageTx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeployTx struct {
	VmType VMType `protobuf:"varint,1,opt,name=vm_type,json=vmType,proto3,enum=VMType" json:"vm_type,omitempty"`
	Code   []byte `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeployTx) Reset()                    { *m = DeployTx{} }
func (m *DeployTx) String() string            { return proto.CompactTextString(m) }
func (*DeployTx) ProtoMessage()               {}
func (*DeployTx) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{1} }

func (m *DeployTx) GetVmType() VMType {
	if m != nil {
		return m.VmType
	}
	return VMType_PLUGIN
}

func (m *DeployTx) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *DeployTx) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CallTx struct {
	VmType VMType         `protobuf:"varint,1,opt,name=vm_type,json=vmType,proto3,enum=VMType" json:"vm_type,omitempty"`
	Input  []byte         `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	Value  *types.BigUInt `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *CallTx) Reset()                    { *m = CallTx{} }
func (m *CallTx) String() string            { return proto.CompactTextString(m) }
func (*CallTx) ProtoMessage()               {}
func (*CallTx) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{2} }

func (m *CallTx) GetVmType() VMType {
	if m != nil {
		return m.VmType
	}
	return VMType_PLUGIN
}

func (m *CallTx) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *CallTx) GetValue() *types.BigUInt {
	if m != nil {
		return m.Value
	}
	return nil
}

type DeployResponse struct {
	Contract *types.Address `protobuf:"bytes,1,opt,name=contract" json:"contract,omitempty"`
	Output   []byte         `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (m *DeployResponse) Reset()                    { *m = DeployResponse{} }
func (m *DeployResponse) String() string            { return proto.CompactTextString(m) }
func (*DeployResponse) ProtoMessage()               {}
func (*DeployResponse) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{3} }

func (m *DeployResponse) GetContract() *types.Address {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *DeployResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

type DeployResponseData struct {
	TxHash   []byte `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Bytecode []byte `protobuf:"bytes,2,opt,name=bytecode,proto3" json:"bytecode,omitempty"`
}

func (m *DeployResponseData) Reset()                    { *m = DeployResponseData{} }
func (m *DeployResponseData) String() string            { return proto.CompactTextString(m) }
func (*DeployResponseData) ProtoMessage()               {}
func (*DeployResponseData) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{4} }

func (m *DeployResponseData) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *DeployResponseData) GetBytecode() []byte {
	if m != nil {
		return m.Bytecode
	}
	return nil
}

type TxHashList struct {
	TxHash [][]byte `protobuf:"bytes,1,rep,name=tx_hash,json=txHash" json:"tx_hash,omitempty"`
}

func (m *TxHashList) Reset()                    { *m = TxHashList{} }
func (m *TxHashList) String() string            { return proto.CompactTextString(m) }
func (*TxHashList) ProtoMessage()               {}
func (*TxHashList) Descriptor() ([]byte, []int) { return fileDescriptorVm, []int{5} }

func (m *TxHashList) GetTxHash() [][]byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageTx)(nil), "MessageTx")
	proto.RegisterType((*DeployTx)(nil), "DeployTx")
	proto.RegisterType((*CallTx)(nil), "CallTx")
	proto.RegisterType((*DeployResponse)(nil), "DeployResponse")
	proto.RegisterType((*DeployResponseData)(nil), "DeployResponseData")
	proto.RegisterType((*TxHashList)(nil), "TxHashList")
	proto.RegisterEnum("VMType", VMType_name, VMType_value)
}

func init() { proto.RegisterFile("github.com/loomnetwork/go-loom/vm/vm.proto", fileDescriptorVm) }

var fileDescriptorVm = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x6f, 0xab, 0xd3, 0x30,
	0x14, 0xc6, 0xed, 0xfe, 0xb4, 0xdd, 0x71, 0x8c, 0x11, 0x44, 0xcb, 0x50, 0x19, 0x45, 0x61, 0x0c,
	0xec, 0x64, 0x7e, 0x02, 0x75, 0xa2, 0x83, 0x6d, 0x48, 0xe8, 0xe6, 0xcb, 0x99, 0xb5, 0xb1, 0x2d,
	0xb6, 0x4d, 0x69, 0x4e, 0x6b, 0xfb, 0xed, 0xa5, 0xe9, 0xd8, 0xee, 0xee, 0x7d, 0x71, 0x2f, 0x94,
	0xe6, 0x3c, 0xe7, 0x9c, 0xfc, 0x4e, 0xf2, 0x04, 0xe6, 0x41, 0x84, 0x61, 0x71, 0x72, 0x3c, 0x91,
	0x2c, 0x62, 0x21, 0x92, 0x94, 0xe3, 0x3f, 0x91, 0xff, 0x5d, 0x04, 0xe2, 0x43, 0x23, 0x17, 0x65,
	0xf3, 0x39, 0x59, 0x2e, 0x50, 0x4c, 0x3e, 0x3e, 0xd2, 0x8b, 0x75, 0xc6, 0x65, 0xfb, 0x6f, 0x77,
	0xd8, 0xbf, 0x60, 0xb0, 0xe5, 0x52, 0xb2, 0x80, 0xbb, 0x15, 0xb1, 0xa0, 0x83, 0xc2, 0xd2, 0xa6,
	0xda, 0xec, 0xf9, 0xd2, 0x74, 0x3e, 0xfb, 0x7e, 0xce, 0xa5, 0xa4, 0x1d, 0x14, 0xe4, 0x35, 0xf4,
	0xfe, 0xe4, 0x22, 0xb1, 0x3a, 0xf7, 0x6a, 0x2a, 0x4b, 0x08, 0xf4, 0x7c, 0x86, 0xcc, 0xea, 0x4e,
	0xb5, 0xd9, 0x90, 0xaa, 0xd8, 0x76, 0xc1, 0x5c, 0xf1, 0x2c, 0x16, 0xb5, 0x5b, 0x91, 0x29, 0x18,
	0x65, 0x72, 0x6c, 0xc6, 0x2a, 0xf8, 0x68, 0x69, 0x38, 0x87, 0xad, 0x5b, 0x67, 0x9c, 0xea, 0x65,
	0xd2, 0xac, 0x0d, 0xc1, 0x13, 0x3e, 0x57, 0xfc, 0x21, 0x55, 0x71, 0x93, 0x4b, 0x59, 0xc2, 0x15,
	0x75, 0x40, 0x55, 0x6c, 0xff, 0x06, 0xfd, 0x2b, 0x8b, 0xe3, 0x27, 0x31, 0x5f, 0x40, 0x3f, 0x4a,
	0xb3, 0x02, 0xcf, 0xd0, 0x56, 0x90, 0xb7, 0xd0, 0x2f, 0x59, 0x5c, 0xb4, 0xd8, 0xe6, 0x2a, 0x5f,
	0xa2, 0x60, 0xbf, 0x4e, 0x91, 0xb6, 0x69, 0x7b, 0x07, 0xa3, 0xf6, 0xdc, 0x94, 0xcb, 0x4c, 0xa4,
	0x92, 0x93, 0x77, 0x60, 0x7a, 0x22, 0xc5, 0x9c, 0x79, 0xf8, 0xc0, 0x9b, 0x4b, 0x85, 0xbc, 0x04,
	0x5d, 0x14, 0x78, 0x1d, 0x77, 0x56, 0xf6, 0x1a, 0xc8, 0x2d, 0x6f, 0xc5, 0x90, 0x91, 0x57, 0x60,
	0x60, 0x75, 0x0c, 0x99, 0x0c, 0x15, 0x72, 0x48, 0x75, 0xac, 0x7e, 0x30, 0x19, 0x92, 0x09, 0x98,
	0xa7, 0x1a, 0xf9, 0x1d, 0x33, 0x2e, 0xda, 0x7e, 0x0f, 0xe0, 0xaa, 0xae, 0x4d, 0x24, 0xf1, 0x16,
	0xd1, 0xbd, 0x22, 0xe6, 0x6f, 0x40, 0x6f, 0x9d, 0x20, 0x00, 0xfa, 0xcf, 0xcd, 0xfe, 0xfb, 0x7a,
	0x37, 0x7e, 0x46, 0x0c, 0xe8, 0x7e, 0x3b, 0x6c, 0xc7, 0xda, 0x49, 0x57, 0x0f, 0xff, 0xe9, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x23, 0xe4, 0xd6, 0x58, 0x02, 0x00, 0x00,
}
