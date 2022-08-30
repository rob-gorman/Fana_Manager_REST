// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: responses.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Flags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flags []*FlagSparseResp `protobuf:"bytes,1,rep,name=Flags,proto3" json:"Flags,omitempty"`
}

func (x *Flags) Reset() {
	*x = Flags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flags) ProtoMessage() {}

func (x *Flags) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flags.ProtoReflect.Descriptor instead.
func (*Flags) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{0}
}

func (x *Flags) GetFlags() []*FlagSparseResp {
	if x != nil {
		return x.Flags
	}
	return nil
}

type FlagSparseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Key         string                 `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	DisplayName string                 `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	Status      bool                   `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *FlagSparseResp) Reset() {
	*x = FlagSparseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagSparseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagSparseResp) ProtoMessage() {}

func (x *FlagSparseResp) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagSparseResp.ProtoReflect.Descriptor instead.
func (*FlagSparseResp) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{1}
}

func (x *FlagSparseResp) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *FlagSparseResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FlagSparseResp) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *FlagSparseResp) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *FlagSparseResp) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FlagSparseResp) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type FlagFullResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Key         string                 `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	DisplayName string                 `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	Status      bool                   `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Audiences   []*AudienceSparseResp  `protobuf:"bytes,7,rep,name=Audiences,proto3" json:"Audiences,omitempty"`
}

func (x *FlagFullResp) Reset() {
	*x = FlagFullResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagFullResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagFullResp) ProtoMessage() {}

func (x *FlagFullResp) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagFullResp.ProtoReflect.Descriptor instead.
func (*FlagFullResp) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{2}
}

func (x *FlagFullResp) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *FlagFullResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FlagFullResp) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *FlagFullResp) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *FlagFullResp) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *FlagFullResp) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *FlagFullResp) GetAudiences() []*AudienceSparseResp {
	if x != nil {
		return x.Audiences
	}
	return nil
}

type Audiences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audiences []*AudienceSparseResp `protobuf:"bytes,1,rep,name=Audiences,proto3" json:"Audiences,omitempty"`
}

func (x *Audiences) Reset() {
	*x = Audiences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audiences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audiences) ProtoMessage() {}

func (x *Audiences) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audiences.ProtoReflect.Descriptor instead.
func (*Audiences) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{3}
}

func (x *Audiences) GetAudiences() []*AudienceSparseResp {
	if x != nil {
		return x.Audiences
	}
	return nil
}

type AudienceSparseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Key         string                 `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	DisplayName string                 `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *AudienceSparseResp) Reset() {
	*x = AudienceSparseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudienceSparseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudienceSparseResp) ProtoMessage() {}

func (x *AudienceSparseResp) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudienceSparseResp.ProtoReflect.Descriptor instead.
func (*AudienceSparseResp) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{4}
}

func (x *AudienceSparseResp) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AudienceSparseResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AudienceSparseResp) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AudienceSparseResp) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AudienceSparseResp) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type AudienceFullResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Key         string                 `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	DisplayName string                 `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	Combine     string                 `protobuf:"bytes,4,opt,name=Combine,proto3" json:"Combine,omitempty"`
	Conditions  []*ConditionEmbedded   `protobuf:"bytes,5,rep,name=Conditions,proto3" json:"Conditions,omitempty"`
	Flags       []*FlagSparseResp      `protobuf:"bytes,6,rep,name=Flags,proto3" json:"Flags,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *AudienceFullResp) Reset() {
	*x = AudienceFullResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudienceFullResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudienceFullResp) ProtoMessage() {}

func (x *AudienceFullResp) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudienceFullResp.ProtoReflect.Descriptor instead.
func (*AudienceFullResp) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{5}
}

func (x *AudienceFullResp) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AudienceFullResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AudienceFullResp) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AudienceFullResp) GetCombine() string {
	if x != nil {
		return x.Combine
	}
	return ""
}

func (x *AudienceFullResp) GetConditions() []*ConditionEmbedded {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *AudienceFullResp) GetFlags() []*FlagSparseResp {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *AudienceFullResp) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AudienceFullResp) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ConditionEmbedded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attribute *AttributeResp `protobuf:"bytes,1,opt,name=Attribute,proto3" json:"Attribute,omitempty"`
	Operator  string         `protobuf:"bytes,2,opt,name=Operator,proto3" json:"Operator,omitempty"`
	Negate    bool           `protobuf:"varint,3,opt,name=Negate,proto3" json:"Negate,omitempty"`
	Vals      string         `protobuf:"bytes,4,opt,name=Vals,proto3" json:"Vals,omitempty"`
}

func (x *ConditionEmbedded) Reset() {
	*x = ConditionEmbedded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConditionEmbedded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConditionEmbedded) ProtoMessage() {}

func (x *ConditionEmbedded) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConditionEmbedded.ProtoReflect.Descriptor instead.
func (*ConditionEmbedded) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{6}
}

func (x *ConditionEmbedded) GetAttribute() *AttributeResp {
	if x != nil {
		return x.Attribute
	}
	return nil
}

func (x *ConditionEmbedded) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *ConditionEmbedded) GetNegate() bool {
	if x != nil {
		return x.Negate
	}
	return false
}

func (x *ConditionEmbedded) GetVals() string {
	if x != nil {
		return x.Vals
	}
	return ""
}

type Attributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*AttributeResp `protobuf:"bytes,1,rep,name=Attributes,proto3" json:"Attributes,omitempty"`
}

func (x *Attributes) Reset() {
	*x = Attributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attributes) ProtoMessage() {}

func (x *Attributes) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attributes.ProtoReflect.Descriptor instead.
func (*Attributes) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{7}
}

func (x *Attributes) GetAttributes() []*AttributeResp {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type AttributeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Key         string                 `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	DisplayName string                 `protobuf:"bytes,3,opt,name=DisplayName,proto3" json:"DisplayName,omitempty"`
	Type        string                 `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Audiences   []*AudienceSparseResp  `protobuf:"bytes,5,rep,name=Audiences,proto3" json:"Audiences,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *AttributeResp) Reset() {
	*x = AttributeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeResp) ProtoMessage() {}

func (x *AttributeResp) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeResp.ProtoReflect.Descriptor instead.
func (*AttributeResp) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{8}
}

func (x *AttributeResp) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AttributeResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AttributeResp) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *AttributeResp) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AttributeResp) GetAudiences() []*AudienceSparseResp {
	if x != nil {
		return x.Audiences
	}
	return nil
}

func (x *AttributeResp) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type AuditLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlagLogs      []*LogMsg `protobuf:"bytes,1,rep,name=FlagLogs,proto3" json:"FlagLogs,omitempty"`
	AudienceLogs  []*LogMsg `protobuf:"bytes,2,rep,name=AudienceLogs,proto3" json:"AudienceLogs,omitempty"`
	AttributeLogs []*LogMsg `protobuf:"bytes,3,rep,name=AttributeLogs,proto3" json:"AttributeLogs,omitempty"`
}

func (x *AuditLogResp) Reset() {
	*x = AuditLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLogResp) ProtoMessage() {}

func (x *AuditLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLogResp.ProtoReflect.Descriptor instead.
func (*AuditLogResp) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{9}
}

func (x *AuditLogResp) GetFlagLogs() []*LogMsg {
	if x != nil {
		return x.FlagLogs
	}
	return nil
}

func (x *AuditLogResp) GetAudienceLogs() []*LogMsg {
	if x != nil {
		return x.AudienceLogs
	}
	return nil
}

func (x *AuditLogResp) GetAttributeLogs() []*LogMsg {
	if x != nil {
		return x.AttributeLogs
	}
	return nil
}

type LogMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogID     int32                  `protobuf:"varint,1,opt,name=LogID,proto3" json:"LogID,omitempty"`
	ID        int32                  `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Key       string                 `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Action    string                 `protobuf:"bytes,4,opt,name=Action,proto3" json:"Action,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *LogMsg) Reset() {
	*x = LogMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_responses_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogMsg) ProtoMessage() {}

func (x *LogMsg) ProtoReflect() protoreflect.Message {
	mi := &file_responses_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogMsg.ProtoReflect.Descriptor instead.
func (*LogMsg) Descriptor() ([]byte, []int) {
	return file_responses_proto_rawDescGZIP(), []int{10}
}

func (x *LogMsg) GetLogID() int32 {
	if x != nil {
		return x.LogID
	}
	return 0
}

func (x *LogMsg) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *LogMsg) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LogMsg) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *LogMsg) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_responses_proto protoreflect.FileDescriptor

var file_responses_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x46, 0x6c, 0x61,
	0x67, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x05, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x0e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x70, 0x61, 0x72, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x91, 0x02, 0x0a, 0x0c, 0x46, 0x6c, 0x61, 0x67, 0x46, 0x75,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x09,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x3e, 0x0a, 0x09, 0x41, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x09, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x41, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x09,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x12, 0x41, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b,
	0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xbf, 0x02, 0x0a, 0x10, 0x41, 0x75, 0x64,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x25, 0x0a, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x46, 0x6c, 0x61, 0x67, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x11, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x2c, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x52, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x4e, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x56, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x56, 0x61, 0x6c, 0x73, 0x22, 0x3c, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31,
	0x0a, 0x09, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x09, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x0c,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x08,
	0x46, 0x6c, 0x61, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x73, 0x67, 0x52, 0x08, 0x46, 0x6c, 0x61, 0x67, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x2b, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x73, 0x67,
	0x52, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x2d,
	0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x73, 0x67, 0x52, 0x0d,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x92, 0x01,
	0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_responses_proto_rawDescOnce sync.Once
	file_responses_proto_rawDescData = file_responses_proto_rawDesc
)

func file_responses_proto_rawDescGZIP() []byte {
	file_responses_proto_rawDescOnce.Do(func() {
		file_responses_proto_rawDescData = protoimpl.X.CompressGZIP(file_responses_proto_rawDescData)
	})
	return file_responses_proto_rawDescData
}

var file_responses_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_responses_proto_goTypes = []interface{}{
	(*Flags)(nil),                 // 0: Flags
	(*FlagSparseResp)(nil),        // 1: FlagSparseResp
	(*FlagFullResp)(nil),          // 2: FlagFullResp
	(*Audiences)(nil),             // 3: Audiences
	(*AudienceSparseResp)(nil),    // 4: AudienceSparseResp
	(*AudienceFullResp)(nil),      // 5: AudienceFullResp
	(*ConditionEmbedded)(nil),     // 6: ConditionEmbedded
	(*Attributes)(nil),            // 7: Attributes
	(*AttributeResp)(nil),         // 8: AttributeResp
	(*AuditLogResp)(nil),          // 9: AuditLogResp
	(*LogMsg)(nil),                // 10: LogMsg
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_responses_proto_depIdxs = []int32{
	1,  // 0: Flags.Flags:type_name -> FlagSparseResp
	11, // 1: FlagSparseResp.CreatedAt:type_name -> google.protobuf.Timestamp
	11, // 2: FlagSparseResp.UpdatedAt:type_name -> google.protobuf.Timestamp
	11, // 3: FlagFullResp.CreatedAt:type_name -> google.protobuf.Timestamp
	11, // 4: FlagFullResp.UpdatedAt:type_name -> google.protobuf.Timestamp
	4,  // 5: FlagFullResp.Audiences:type_name -> AudienceSparseResp
	4,  // 6: Audiences.Audiences:type_name -> AudienceSparseResp
	11, // 7: AudienceSparseResp.CreatedAt:type_name -> google.protobuf.Timestamp
	11, // 8: AudienceSparseResp.UpdatedAt:type_name -> google.protobuf.Timestamp
	6,  // 9: AudienceFullResp.Conditions:type_name -> ConditionEmbedded
	1,  // 10: AudienceFullResp.Flags:type_name -> FlagSparseResp
	11, // 11: AudienceFullResp.CreatedAt:type_name -> google.protobuf.Timestamp
	11, // 12: AudienceFullResp.UpdatedAt:type_name -> google.protobuf.Timestamp
	8,  // 13: ConditionEmbedded.Attribute:type_name -> AttributeResp
	8,  // 14: Attributes.Attributes:type_name -> AttributeResp
	4,  // 15: AttributeResp.Audiences:type_name -> AudienceSparseResp
	11, // 16: AttributeResp.CreatedAt:type_name -> google.protobuf.Timestamp
	10, // 17: AuditLogResp.FlagLogs:type_name -> LogMsg
	10, // 18: AuditLogResp.AudienceLogs:type_name -> LogMsg
	10, // 19: AuditLogResp.AttributeLogs:type_name -> LogMsg
	11, // 20: LogMsg.CreatedAt:type_name -> google.protobuf.Timestamp
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_responses_proto_init() }
func file_responses_proto_init() {
	if File_responses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_responses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flags); i {
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
		file_responses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagSparseResp); i {
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
		file_responses_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagFullResp); i {
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
		file_responses_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audiences); i {
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
		file_responses_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudienceSparseResp); i {
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
		file_responses_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudienceFullResp); i {
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
		file_responses_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConditionEmbedded); i {
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
		file_responses_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attributes); i {
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
		file_responses_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeResp); i {
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
		file_responses_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLogResp); i {
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
		file_responses_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogMsg); i {
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
			RawDescriptor: file_responses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_responses_proto_goTypes,
		DependencyIndexes: file_responses_proto_depIdxs,
		MessageInfos:      file_responses_proto_msgTypes,
	}.Build()
	File_responses_proto = out.File
	file_responses_proto_rawDesc = nil
	file_responses_proto_goTypes = nil
	file_responses_proto_depIdxs = nil
}