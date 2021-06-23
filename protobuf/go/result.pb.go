// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: result.proto

package pb

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

type Error_Code int32

const (
	Error_UNKNOWN_CODE Error_Code = 0
	Error_NOT_FOUND    Error_Code = 1
	Error_INVALID      Error_Code = 2
	Error_BLOCKED      Error_Code = 3
)

// Enum value maps for Error_Code.
var (
	Error_Code_name = map[int32]string{
		0: "UNKNOWN_CODE",
		1: "NOT_FOUND",
		2: "INVALID",
		3: "BLOCKED",
	}
	Error_Code_value = map[string]int32{
		"UNKNOWN_CODE": 0,
		"NOT_FOUND":    1,
		"INVALID":      2,
		"BLOCKED":      3,
	}
)

func (x Error_Code) Enum() *Error_Code {
	p := new(Error_Code)
	*p = x
	return p
}

func (x Error_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_result_proto_enumTypes[0].Descriptor()
}

func (Error_Code) Type() protoreflect.EnumType {
	return &file_result_proto_enumTypes[0]
}

func (x Error_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Code.Descriptor instead.
func (Error_Code) EnumDescriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{3, 0}
}

type Outputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txos         []*Output  `protobuf:"bytes,1,rep,name=txos,proto3" json:"txos"`
	ExtraTxos    []*Output  `protobuf:"bytes,2,rep,name=extra_txos,json=extraTxos,proto3" json:"extra_txos"`
	Total        uint32     `protobuf:"varint,3,opt,name=total,proto3" json:"total"`
	Offset       uint32     `protobuf:"varint,4,opt,name=offset,proto3" json:"offset"`
	Blocked      []*Blocked `protobuf:"bytes,5,rep,name=blocked,proto3" json:"blocked"`
	BlockedTotal uint32     `protobuf:"varint,6,opt,name=blocked_total,json=blockedTotal,proto3" json:"blocked_total"`
}

func (x *Outputs) Reset() {
	*x = Outputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outputs) ProtoMessage() {}

func (x *Outputs) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outputs.ProtoReflect.Descriptor instead.
func (*Outputs) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{0}
}

func (x *Outputs) GetTxos() []*Output {
	if x != nil {
		return x.Txos
	}
	return nil
}

func (x *Outputs) GetExtraTxos() []*Output {
	if x != nil {
		return x.ExtraTxos
	}
	return nil
}

func (x *Outputs) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Outputs) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Outputs) GetBlocked() []*Blocked {
	if x != nil {
		return x.Blocked
	}
	return nil
}

func (x *Outputs) GetBlockedTotal() uint32 {
	if x != nil {
		return x.BlockedTotal
	}
	return 0
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash []byte `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash"`
	Nout   uint32 `protobuf:"varint,2,opt,name=nout,proto3" json:"nout"`
	Height uint32 `protobuf:"varint,3,opt,name=height,proto3" json:"height"`
	// Types that are assignable to Meta:
	//	*Output_Claim
	//	*Output_Error
	Meta isOutput_Meta `protobuf_oneof:"meta"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetTxHash() []byte {
	if x != nil {
		return x.TxHash
	}
	return nil
}

func (x *Output) GetNout() uint32 {
	if x != nil {
		return x.Nout
	}
	return 0
}

func (x *Output) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (m *Output) GetMeta() isOutput_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (x *Output) GetClaim() *ClaimMeta {
	if x, ok := x.GetMeta().(*Output_Claim); ok {
		return x.Claim
	}
	return nil
}

func (x *Output) GetError() *Error {
	if x, ok := x.GetMeta().(*Output_Error); ok {
		return x.Error
	}
	return nil
}

type isOutput_Meta interface {
	isOutput_Meta()
}

type Output_Claim struct {
	Claim *ClaimMeta `protobuf:"bytes,7,opt,name=claim,proto3,oneof"`
}

type Output_Error struct {
	Error *Error `protobuf:"bytes,15,opt,name=error,proto3,oneof"`
}

func (*Output_Claim) isOutput_Meta() {}

func (*Output_Error) isOutput_Meta() {}

type ClaimMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel          *Output `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel"`
	Repost           *Output `protobuf:"bytes,2,opt,name=repost,proto3" json:"repost"`
	ShortUrl         string  `protobuf:"bytes,3,opt,name=short_url,json=shortUrl,proto3" json:"short_url"`
	CanonicalUrl     string  `protobuf:"bytes,4,opt,name=canonical_url,json=canonicalUrl,proto3" json:"canonical_url"`
	IsControlling    bool    `protobuf:"varint,5,opt,name=is_controlling,json=isControlling,proto3" json:"is_controlling"`
	TakeOverHeight   uint32  `protobuf:"varint,6,opt,name=take_over_height,json=takeOverHeight,proto3" json:"take_over_height"`
	CreationHeight   uint32  `protobuf:"varint,7,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height"`
	ActivationHeight uint32  `protobuf:"varint,8,opt,name=activation_height,json=activationHeight,proto3" json:"activation_height"`
	ExpirationHeight uint32  `protobuf:"varint,9,opt,name=expiration_height,json=expirationHeight,proto3" json:"expiration_height"`
	ClaimsInChannel  uint32  `protobuf:"varint,10,opt,name=claims_in_channel,json=claimsInChannel,proto3" json:"claims_in_channel"`
	Reposted         uint32  `protobuf:"varint,11,opt,name=reposted,proto3" json:"reposted"`
	EffectiveAmount  uint64  `protobuf:"varint,20,opt,name=effective_amount,json=effectiveAmount,proto3" json:"effective_amount"`
	SupportAmount    uint64  `protobuf:"varint,21,opt,name=support_amount,json=supportAmount,proto3" json:"support_amount"`
	TrendingGroup    uint32  `protobuf:"varint,22,opt,name=trending_group,json=trendingGroup,proto3" json:"trending_group"`
	TrendingMixed    float32 `protobuf:"fixed32,23,opt,name=trending_mixed,json=trendingMixed,proto3" json:"trending_mixed"`
	TrendingLocal    float32 `protobuf:"fixed32,24,opt,name=trending_local,json=trendingLocal,proto3" json:"trending_local"`
	TrendingGlobal   float32 `protobuf:"fixed32,25,opt,name=trending_global,json=trendingGlobal,proto3" json:"trending_global"`
}

func (x *ClaimMeta) Reset() {
	*x = ClaimMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimMeta) ProtoMessage() {}

func (x *ClaimMeta) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimMeta.ProtoReflect.Descriptor instead.
func (*ClaimMeta) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{2}
}

func (x *ClaimMeta) GetChannel() *Output {
	if x != nil {
		return x.Channel
	}
	return nil
}

func (x *ClaimMeta) GetRepost() *Output {
	if x != nil {
		return x.Repost
	}
	return nil
}

func (x *ClaimMeta) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *ClaimMeta) GetCanonicalUrl() string {
	if x != nil {
		return x.CanonicalUrl
	}
	return ""
}

func (x *ClaimMeta) GetIsControlling() bool {
	if x != nil {
		return x.IsControlling
	}
	return false
}

func (x *ClaimMeta) GetTakeOverHeight() uint32 {
	if x != nil {
		return x.TakeOverHeight
	}
	return 0
}

func (x *ClaimMeta) GetCreationHeight() uint32 {
	if x != nil {
		return x.CreationHeight
	}
	return 0
}

func (x *ClaimMeta) GetActivationHeight() uint32 {
	if x != nil {
		return x.ActivationHeight
	}
	return 0
}

func (x *ClaimMeta) GetExpirationHeight() uint32 {
	if x != nil {
		return x.ExpirationHeight
	}
	return 0
}

func (x *ClaimMeta) GetClaimsInChannel() uint32 {
	if x != nil {
		return x.ClaimsInChannel
	}
	return 0
}

func (x *ClaimMeta) GetReposted() uint32 {
	if x != nil {
		return x.Reposted
	}
	return 0
}

func (x *ClaimMeta) GetEffectiveAmount() uint64 {
	if x != nil {
		return x.EffectiveAmount
	}
	return 0
}

func (x *ClaimMeta) GetSupportAmount() uint64 {
	if x != nil {
		return x.SupportAmount
	}
	return 0
}

func (x *ClaimMeta) GetTrendingGroup() uint32 {
	if x != nil {
		return x.TrendingGroup
	}
	return 0
}

func (x *ClaimMeta) GetTrendingMixed() float32 {
	if x != nil {
		return x.TrendingMixed
	}
	return 0
}

func (x *ClaimMeta) GetTrendingLocal() float32 {
	if x != nil {
		return x.TrendingLocal
	}
	return 0
}

func (x *ClaimMeta) GetTrendingGlobal() float32 {
	if x != nil {
		return x.TrendingGlobal
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Error_Code `protobuf:"varint,1,opt,name=code,proto3,enum=pb.Error_Code" json:"code"`
	Text    string     `protobuf:"bytes,2,opt,name=text,proto3" json:"text"`
	Blocked *Blocked   `protobuf:"bytes,3,opt,name=blocked,proto3" json:"blocked"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetCode() Error_Code {
	if x != nil {
		return x.Code
	}
	return Error_UNKNOWN_CODE
}

func (x *Error) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Error) GetBlocked() *Blocked {
	if x != nil {
		return x.Blocked
	}
	return nil
}

type Blocked struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   uint32  `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	Channel *Output `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel"`
}

func (x *Blocked) Reset() {
	*x = Blocked{}
	if protoimpl.UnsafeEnabled {
		mi := &file_result_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blocked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blocked) ProtoMessage() {}

func (x *Blocked) ProtoReflect() protoreflect.Message {
	mi := &file_result_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blocked.ProtoReflect.Descriptor instead.
func (*Blocked) Descriptor() ([]byte, []int) {
	return file_result_proto_rawDescGZIP(), []int{4}
}

func (x *Blocked) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Blocked) GetChannel() *Output {
	if x != nil {
		return x.Channel
	}
	return nil
}

var File_result_proto protoreflect.FileDescriptor

var file_result_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0xce, 0x01, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x1e,
	0x0a, 0x04, 0x74, 0x78, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x04, 0x74, 0x78, 0x6f, 0x73, 0x12, 0x29,
	0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x74, 0x78, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x09,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x54, 0x78, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x9f, 0x01, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6e, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x65, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x06, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0xa3, 0x05, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61,
	0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12,
	0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x6f,
	0x76, 0x65, 0x72, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x74, 0x61, 0x6b, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5f, 0x69, 0x6e,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x49, 0x6e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x6d, 0x69, 0x78, 0x65, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x74, 0x72,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x69, 0x78, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0d, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x19, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x74, 0x72, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x22, 0xa9, 0x01, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x25, 0x0a,
	0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x07, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x22, 0x41, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x0c,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4c,
	0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x03, 0x22, 0x45, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x62, 0x72,
	0x79, 0x69, 0x6f, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_result_proto_rawDescOnce sync.Once
	file_result_proto_rawDescData = file_result_proto_rawDesc
)

func file_result_proto_rawDescGZIP() []byte {
	file_result_proto_rawDescOnce.Do(func() {
		file_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_result_proto_rawDescData)
	})
	return file_result_proto_rawDescData
}

var file_result_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_result_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_result_proto_goTypes = []interface{}{
	(Error_Code)(0),   // 0: pb.Error.Code
	(*Outputs)(nil),   // 1: pb.Outputs
	(*Output)(nil),    // 2: pb.Output
	(*ClaimMeta)(nil), // 3: pb.ClaimMeta
	(*Error)(nil),     // 4: pb.Error
	(*Blocked)(nil),   // 5: pb.Blocked
}
var file_result_proto_depIdxs = []int32{
	2,  // 0: pb.Outputs.txos:type_name -> pb.Output
	2,  // 1: pb.Outputs.extra_txos:type_name -> pb.Output
	5,  // 2: pb.Outputs.blocked:type_name -> pb.Blocked
	3,  // 3: pb.Output.claim:type_name -> pb.ClaimMeta
	4,  // 4: pb.Output.error:type_name -> pb.Error
	2,  // 5: pb.ClaimMeta.channel:type_name -> pb.Output
	2,  // 6: pb.ClaimMeta.repost:type_name -> pb.Output
	0,  // 7: pb.Error.code:type_name -> pb.Error.Code
	5,  // 8: pb.Error.blocked:type_name -> pb.Blocked
	2,  // 9: pb.Blocked.channel:type_name -> pb.Output
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_result_proto_init() }
func file_result_proto_init() {
	if File_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outputs); i {
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
		file_result_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_result_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimMeta); i {
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
		file_result_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_result_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blocked); i {
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
	file_result_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Output_Claim)(nil),
		(*Output_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_result_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_result_proto_goTypes,
		DependencyIndexes: file_result_proto_depIdxs,
		EnumInfos:         file_result_proto_enumTypes,
		MessageInfos:      file_result_proto_msgTypes,
	}.Build()
	File_result_proto = out.File
	file_result_proto_rawDesc = nil
	file_result_proto_goTypes = nil
	file_result_proto_depIdxs = nil
}
