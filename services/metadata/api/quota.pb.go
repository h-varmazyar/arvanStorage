// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: services/metadata/api/src/quota.proto

package api

import (
	api "github.com/h-varmazyar/arvanStorage/api"
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

type Quota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"user_id"
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"user_id"`
	// @inject_tag: json:"max_request_per_minute"
	MaxRequestPerMinute int64 `protobuf:"varint,3,opt,name=MaxRequestPerMinute,proto3" json:"max_request_per_minute"`
	// @inject_tag: json:"max_volume_in_month"
	MaxVolumeInMonth int64 `protobuf:"varint,4,opt,name=MaxVolumeInMonth,proto3" json:"max_volume_in_month"`
	// @inject_tag: json:"total_used_volume"
	TotalUsedVolume int64 `protobuf:"varint,5,opt,name=TotalUsedVolume,proto3" json:"total_used_volume"`
}

func (x *Quota) Reset() {
	*x = Quota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_metadata_api_src_quota_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quota) ProtoMessage() {}

func (x *Quota) ProtoReflect() protoreflect.Message {
	mi := &file_services_metadata_api_src_quota_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quota.ProtoReflect.Descriptor instead.
func (*Quota) Descriptor() ([]byte, []int) {
	return file_services_metadata_api_src_quota_proto_rawDescGZIP(), []int{0}
}

func (x *Quota) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Quota) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Quota) GetMaxRequestPerMinute() int64 {
	if x != nil {
		return x.MaxRequestPerMinute
	}
	return 0
}

func (x *Quota) GetMaxVolumeInMonth() int64 {
	if x != nil {
		return x.MaxVolumeInMonth
	}
	return 0
}

func (x *Quota) GetTotalUsedVolume() int64 {
	if x != nil {
		return x.TotalUsedVolume
	}
	return 0
}

type QuotaRemainingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"account_id"
	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"account_id"`
}

func (x *QuotaRemainingReq) Reset() {
	*x = QuotaRemainingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_metadata_api_src_quota_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaRemainingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaRemainingReq) ProtoMessage() {}

func (x *QuotaRemainingReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_metadata_api_src_quota_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaRemainingReq.ProtoReflect.Descriptor instead.
func (*QuotaRemainingReq) Descriptor() ([]byte, []int) {
	return file_services_metadata_api_src_quota_proto_rawDescGZIP(), []int{1}
}

func (x *QuotaRemainingReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type QuotaRemainingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"total_remaining_volume"
	TotalRemainingVolume int64 `protobuf:"varint,1,opt,name=TotalRemainingVolume,proto3" json:"total_remaining_volume"`
}

func (x *QuotaRemainingResp) Reset() {
	*x = QuotaRemainingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_metadata_api_src_quota_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaRemainingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaRemainingResp) ProtoMessage() {}

func (x *QuotaRemainingResp) ProtoReflect() protoreflect.Message {
	mi := &file_services_metadata_api_src_quota_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaRemainingResp.ProtoReflect.Descriptor instead.
func (*QuotaRemainingResp) Descriptor() ([]byte, []int) {
	return file_services_metadata_api_src_quota_proto_rawDescGZIP(), []int{2}
}

func (x *QuotaRemainingResp) GetTotalRemainingVolume() int64 {
	if x != nil {
		return x.TotalRemainingVolume
	}
	return 0
}

type QuotaRequestLimitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"account_id"
	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"account_id"`
}

func (x *QuotaRequestLimitReq) Reset() {
	*x = QuotaRequestLimitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_metadata_api_src_quota_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaRequestLimitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaRequestLimitReq) ProtoMessage() {}

func (x *QuotaRequestLimitReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_metadata_api_src_quota_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaRequestLimitReq.ProtoReflect.Descriptor instead.
func (*QuotaRequestLimitReq) Descriptor() ([]byte, []int) {
	return file_services_metadata_api_src_quota_proto_rawDescGZIP(), []int{3}
}

func (x *QuotaRequestLimitReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type QuotaRequestLimitResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"count"
	Count int64 `protobuf:"varint,1,opt,name=Count,proto3" json:"count"`
	// @inject_tag: json:"duration"
	Duration int64 `protobuf:"varint,2,opt,name=Duration,proto3" json:"duration"`
}

func (x *QuotaRequestLimitResp) Reset() {
	*x = QuotaRequestLimitResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_metadata_api_src_quota_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaRequestLimitResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaRequestLimitResp) ProtoMessage() {}

func (x *QuotaRequestLimitResp) ProtoReflect() protoreflect.Message {
	mi := &file_services_metadata_api_src_quota_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaRequestLimitResp.ProtoReflect.Descriptor instead.
func (*QuotaRequestLimitResp) Descriptor() ([]byte, []int) {
	return file_services_metadata_api_src_quota_proto_rawDescGZIP(), []int{4}
}

func (x *QuotaRequestLimitResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *QuotaRequestLimitResp) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"account_id"
	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"account_id"`
	// @inject_tag: json:"volume"
	Volume int64 `protobuf:"varint,2,opt,name=Volume,proto3" json:"volume"`
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_metadata_api_src_quota_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_services_metadata_api_src_quota_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_services_metadata_api_src_quota_proto_rawDescGZIP(), []int{5}
}

func (x *Volume) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Volume) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

var File_services_metadata_api_src_quota_proto protoreflect.FileDescriptor

var file_services_metadata_api_src_quota_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x12, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xb7, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x13, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x65, 0x72, 0x4d, 0x69,
	0x6e, 0x75, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x4d, 0x61, 0x78, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x49, 0x6e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x4d, 0x61, 0x78, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x49, 0x6e, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x28, 0x0a, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x73, 0x65, 0x64, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x55, 0x73, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x51, 0x75,
	0x6f, 0x74, 0x61, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x12, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a,
	0x14, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x22, 0x2e, 0x0a, 0x14, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x49, 0x0a, 0x15, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x06,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x32, 0xbf, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x6f, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x64, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x1a, 0x09, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72, 0x6d, 0x61, 0x7a, 0x79,
	0x61, 0x72, 0x2f, 0x61, 0x72, 0x76, 0x61, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_metadata_api_src_quota_proto_rawDescOnce sync.Once
	file_services_metadata_api_src_quota_proto_rawDescData = file_services_metadata_api_src_quota_proto_rawDesc
)

func file_services_metadata_api_src_quota_proto_rawDescGZIP() []byte {
	file_services_metadata_api_src_quota_proto_rawDescOnce.Do(func() {
		file_services_metadata_api_src_quota_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_metadata_api_src_quota_proto_rawDescData)
	})
	return file_services_metadata_api_src_quota_proto_rawDescData
}

var file_services_metadata_api_src_quota_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_metadata_api_src_quota_proto_goTypes = []interface{}{
	(*Quota)(nil),                 // 0: api.Quota
	(*QuotaRemainingReq)(nil),     // 1: api.QuotaRemainingReq
	(*QuotaRemainingResp)(nil),    // 2: api.QuotaRemainingResp
	(*QuotaRequestLimitReq)(nil),  // 3: api.QuotaRequestLimitReq
	(*QuotaRequestLimitResp)(nil), // 4: api.QuotaRequestLimitResp
	(*Volume)(nil),                // 5: api.Volume
	(*api.Void)(nil),              // 6: api.Void
}
var file_services_metadata_api_src_quota_proto_depIdxs = []int32{
	1, // 0: api.QuotaService.Remaining:input_type -> api.QuotaRemainingReq
	3, // 1: api.QuotaService.RequestLimit:input_type -> api.QuotaRequestLimitReq
	5, // 2: api.QuotaService.UpdateUsedVolume:input_type -> api.Volume
	2, // 3: api.QuotaService.Remaining:output_type -> api.QuotaRemainingResp
	4, // 4: api.QuotaService.RequestLimit:output_type -> api.QuotaRequestLimitResp
	6, // 5: api.QuotaService.UpdateUsedVolume:output_type -> api.Void
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_metadata_api_src_quota_proto_init() }
func file_services_metadata_api_src_quota_proto_init() {
	if File_services_metadata_api_src_quota_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_metadata_api_src_quota_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quota); i {
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
		file_services_metadata_api_src_quota_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaRemainingReq); i {
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
		file_services_metadata_api_src_quota_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaRemainingResp); i {
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
		file_services_metadata_api_src_quota_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaRequestLimitReq); i {
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
		file_services_metadata_api_src_quota_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaRequestLimitResp); i {
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
		file_services_metadata_api_src_quota_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
			RawDescriptor: file_services_metadata_api_src_quota_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_metadata_api_src_quota_proto_goTypes,
		DependencyIndexes: file_services_metadata_api_src_quota_proto_depIdxs,
		MessageInfos:      file_services_metadata_api_src_quota_proto_msgTypes,
	}.Build()
	File_services_metadata_api_src_quota_proto = out.File
	file_services_metadata_api_src_quota_proto_rawDesc = nil
	file_services_metadata_api_src_quota_proto_goTypes = nil
	file_services_metadata_api_src_quota_proto_depIdxs = nil
}
