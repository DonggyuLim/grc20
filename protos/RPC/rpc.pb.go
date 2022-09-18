// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: protos/RPC/rpc.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// /////////////////////////////////////////////////
// /////////////////////////////////////////////////
// Mutate
type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName string `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	To        string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	From      string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	Amount    uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *TransferRequest) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *TransferRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TransferRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TransferRequest) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToBalance   uint64 `protobuf:"varint,1,opt,name=toBalance,proto3" json:"toBalance,omitempty"`
	FromBalance uint64 `protobuf:"varint,2,opt,name=fromBalance,proto3" json:"fromBalance,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *TransferResponse) GetToBalance() uint64 {
	if x != nil {
		return x.ToBalance
	}
	return 0
}

func (x *TransferResponse) GetFromBalance() uint64 {
	if x != nil {
		return x.FromBalance
	}
	return 0
}

type ApproveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName string `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	Owner     string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Spender   string `protobuf:"bytes,3,opt,name=spender,proto3" json:"spender,omitempty"`
	Amount    uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ApproveRequest) Reset() {
	*x = ApproveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveRequest) ProtoMessage() {}

func (x *ApproveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveRequest.ProtoReflect.Descriptor instead.
func (*ApproveRequest) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *ApproveRequest) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *ApproveRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ApproveRequest) GetSpender() string {
	if x != nil {
		return x.Spender
	}
	return ""
}

func (x *ApproveRequest) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ApproveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allowance uint64 `protobuf:"varint,1,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (x *ApproveResponse) Reset() {
	*x = ApproveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApproveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveResponse) ProtoMessage() {}

func (x *ApproveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveResponse.ProtoReflect.Descriptor instead.
func (*ApproveResponse) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *ApproveResponse) GetAllowance() uint64 {
	if x != nil {
		return x.Allowance
	}
	return 0
}

type TransferFromRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName string `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	Owner     string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Spender   string `protobuf:"bytes,3,opt,name=spender,proto3" json:"spender,omitempty"`
	To        string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Amount    uint64 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferFromRequest) Reset() {
	*x = TransferFromRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFromRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFromRequest) ProtoMessage() {}

func (x *TransferFromRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFromRequest.ProtoReflect.Descriptor instead.
func (*TransferFromRequest) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *TransferFromRequest) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *TransferFromRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *TransferFromRequest) GetSpender() string {
	if x != nil {
		return x.Spender
	}
	return ""
}

func (x *TransferFromRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TransferFromRequest) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferFromResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToBalance uint64 `protobuf:"varint,1,opt,name=toBalance,proto3" json:"toBalance,omitempty"`
}

func (x *TransferFromResponse) Reset() {
	*x = TransferFromResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferFromResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferFromResponse) ProtoMessage() {}

func (x *TransferFromResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferFromResponse.ProtoReflect.Descriptor instead.
func (*TransferFromResponse) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *TransferFromResponse) GetToBalance() uint64 {
	if x != nil {
		return x.ToBalance
	}
	return 0
}

// /////////////////////////////////////////////
// /////////////////////////////////////////////
// Query
type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName string `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	Account   string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *GetBalanceRequest) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *GetBalanceRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance uint64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *GetBalanceResponse) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type TokenInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName string `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
}

func (x *TokenInfoRequest) Reset() {
	*x = TokenInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfoRequest) ProtoMessage() {}

func (x *TokenInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfoRequest.ProtoReflect.Descriptor instead.
func (*TokenInfoRequest) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *TokenInfoRequest) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

type TokenInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName   string `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	Symbol      string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimal     uint32 `protobuf:"varint,3,opt,name=decimal,proto3" json:"decimal,omitempty"`
	TotalSupply uint64 `protobuf:"varint,4,opt,name=totalSupply,proto3" json:"totalSupply,omitempty"`
}

func (x *TokenInfoResponse) Reset() {
	*x = TokenInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfoResponse) ProtoMessage() {}

func (x *TokenInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfoResponse.ProtoReflect.Descriptor instead.
func (*TokenInfoResponse) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{9}
}

func (x *TokenInfoResponse) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *TokenInfoResponse) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenInfoResponse) GetDecimal() uint32 {
	if x != nil {
		return x.Decimal
	}
	return 0
}

func (x *TokenInfoResponse) GetTotalSupply() uint64 {
	if x != nil {
		return x.TotalSupply
	}
	return 0
}

type AllowanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenName string `protobuf:"bytes,1,opt,name=tokenName,proto3" json:"tokenName,omitempty"`
	Owner     string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Spender   string `protobuf:"bytes,3,opt,name=spender,proto3" json:"spender,omitempty"`
}

func (x *AllowanceRequest) Reset() {
	*x = AllowanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowanceRequest) ProtoMessage() {}

func (x *AllowanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowanceRequest.ProtoReflect.Descriptor instead.
func (*AllowanceRequest) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{10}
}

func (x *AllowanceRequest) GetTokenName() string {
	if x != nil {
		return x.TokenName
	}
	return ""
}

func (x *AllowanceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AllowanceRequest) GetSpender() string {
	if x != nil {
		return x.Spender
	}
	return ""
}

type AllowanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allowance uint64 `protobuf:"varint,1,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (x *AllowanceResponse) Reset() {
	*x = AllowanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_RPC_rpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowanceResponse) ProtoMessage() {}

func (x *AllowanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_RPC_rpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowanceResponse.ProtoReflect.Descriptor instead.
func (*AllowanceResponse) Descriptor() ([]byte, []int) {
	return file_protos_RPC_rpc_proto_rawDescGZIP(), []int{11}
}

func (x *AllowanceResponse) GetAllowance() uint64 {
	if x != nil {
		return x.Allowance
	}
	return 0
}

var File_protos_RPC_rpc_proto protoreflect.FileDescriptor

var file_protos_RPC_rpc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x52, 0x50, 0x43, 0x2f, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x72, 0x63, 0x32, 0x30, 0x22, 0x6b, 0x0a,
	0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x6f, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x6f, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x66, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x76,
	0x0a, 0x0e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x6f, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x6f, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x4b, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x30, 0x0a, 0x10, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x22, 0x60, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x11, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x32, 0x8e, 0x03, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12,
	0x3b, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x65, 0x72,
	0x63, 0x32, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x2e, 0x65,
	0x72, 0x63, 0x32, 0x30, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x72, 0x6f,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x12, 0x15, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x2e,
	0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x17, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x65, 0x72,
	0x63, 0x32, 0x30, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x6e, 0x67, 0x67, 0x79, 0x75, 0x4c, 0x69,
	0x6d, 0x2f, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x52,
	0x50, 0x43, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_RPC_rpc_proto_rawDescOnce sync.Once
	file_protos_RPC_rpc_proto_rawDescData = file_protos_RPC_rpc_proto_rawDesc
)

func file_protos_RPC_rpc_proto_rawDescGZIP() []byte {
	file_protos_RPC_rpc_proto_rawDescOnce.Do(func() {
		file_protos_RPC_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_RPC_rpc_proto_rawDescData)
	})
	return file_protos_RPC_rpc_proto_rawDescData
}

var file_protos_RPC_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_protos_RPC_rpc_proto_goTypes = []interface{}{
	(*TransferRequest)(nil),      // 0: erc20.TransferRequest
	(*TransferResponse)(nil),     // 1: erc20.TransferResponse
	(*ApproveRequest)(nil),       // 2: erc20.ApproveRequest
	(*ApproveResponse)(nil),      // 3: erc20.ApproveResponse
	(*TransferFromRequest)(nil),  // 4: erc20.TransferFromRequest
	(*TransferFromResponse)(nil), // 5: erc20.TransferFromResponse
	(*GetBalanceRequest)(nil),    // 6: erc20.GetBalanceRequest
	(*GetBalanceResponse)(nil),   // 7: erc20.GetBalanceResponse
	(*TokenInfoRequest)(nil),     // 8: erc20.TokenInfoRequest
	(*TokenInfoResponse)(nil),    // 9: erc20.TokenInfoResponse
	(*AllowanceRequest)(nil),     // 10: erc20.AllowanceRequest
	(*AllowanceResponse)(nil),    // 11: erc20.AllowanceResponse
}
var file_protos_RPC_rpc_proto_depIdxs = []int32{
	0,  // 0: erc20.RPC.Transfer:input_type -> erc20.TransferRequest
	4,  // 1: erc20.RPC.TransferFrom:input_type -> erc20.TransferFromRequest
	2,  // 2: erc20.RPC.Approve:input_type -> erc20.ApproveRequest
	6,  // 3: erc20.RPC.GetBalance:input_type -> erc20.GetBalanceRequest
	8,  // 4: erc20.RPC.GetTokenInfo:input_type -> erc20.TokenInfoRequest
	10, // 5: erc20.RPC.GetAllowance:input_type -> erc20.AllowanceRequest
	1,  // 6: erc20.RPC.Transfer:output_type -> erc20.TransferResponse
	5,  // 7: erc20.RPC.TransferFrom:output_type -> erc20.TransferFromResponse
	3,  // 8: erc20.RPC.Approve:output_type -> erc20.ApproveResponse
	7,  // 9: erc20.RPC.GetBalance:output_type -> erc20.GetBalanceResponse
	9,  // 10: erc20.RPC.GetTokenInfo:output_type -> erc20.TokenInfoResponse
	11, // 11: erc20.RPC.GetAllowance:output_type -> erc20.AllowanceResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_protos_RPC_rpc_proto_init() }
func file_protos_RPC_rpc_proto_init() {
	if File_protos_RPC_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_RPC_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApproveResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFromRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferFromResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_RPC_rpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_RPC_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_RPC_rpc_proto_goTypes,
		DependencyIndexes: file_protos_RPC_rpc_proto_depIdxs,
		MessageInfos:      file_protos_RPC_rpc_proto_msgTypes,
	}.Build()
	File_protos_RPC_rpc_proto = out.File
	file_protos_RPC_rpc_proto_rawDesc = nil
	file_protos_RPC_rpc_proto_goTypes = nil
	file_protos_RPC_rpc_proto_depIdxs = nil
}
