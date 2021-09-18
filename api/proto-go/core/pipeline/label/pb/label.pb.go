// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: label.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/erda-project/erda-proto-go/common/pb"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PipelineLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              uint64                 `protobuf:"varint,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Type            string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	TargetID        uint64                 `protobuf:"varint,3,opt,name=targetID,proto3" json:"targetID,omitempty"`
	PipelineSource  string                 `protobuf:"bytes,4,opt,name=pipelineSource,proto3" json:"pipelineSource,omitempty"`
	PipelineYmlName string                 `protobuf:"bytes,5,opt,name=pipelineYmlName,proto3" json:"pipelineYmlName,omitempty"`
	Key             string                 `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Value           string                 `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	TimeCreated     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=timeCreated,proto3" json:"timeCreated,omitempty"`
	TimeUpdated     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=timeUpdated,proto3" json:"timeUpdated,omitempty"`
}

func (x *PipelineLabel) Reset() {
	*x = PipelineLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineLabel) ProtoMessage() {}

func (x *PipelineLabel) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineLabel.ProtoReflect.Descriptor instead.
func (*PipelineLabel) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{0}
}

func (x *PipelineLabel) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *PipelineLabel) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PipelineLabel) GetTargetID() uint64 {
	if x != nil {
		return x.TargetID
	}
	return 0
}

func (x *PipelineLabel) GetPipelineSource() string {
	if x != nil {
		return x.PipelineSource
	}
	return ""
}

func (x *PipelineLabel) GetPipelineYmlName() string {
	if x != nil {
		return x.PipelineYmlName
	}
	return ""
}

func (x *PipelineLabel) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PipelineLabel) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PipelineLabel) GetTimeCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeCreated
	}
	return nil
}

func (x *PipelineLabel) GetTimeUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeUpdated
	}
	return nil
}

type PipelineLabelBatchInsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels []*PipelineLabel `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *PipelineLabelBatchInsertRequest) Reset() {
	*x = PipelineLabelBatchInsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineLabelBatchInsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineLabelBatchInsertRequest) ProtoMessage() {}

func (x *PipelineLabelBatchInsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineLabelBatchInsertRequest.ProtoReflect.Descriptor instead.
func (*PipelineLabelBatchInsertRequest) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{1}
}

func (x *PipelineLabelBatchInsertRequest) GetLabels() []*PipelineLabel {
	if x != nil {
		return x.Labels
	}
	return nil
}

type PipelineLabelBatchInsertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PipelineLabelBatchInsertResponse) Reset() {
	*x = PipelineLabelBatchInsertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineLabelBatchInsertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineLabelBatchInsertResponse) ProtoMessage() {}

func (x *PipelineLabelBatchInsertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineLabelBatchInsertResponse.ProtoReflect.Descriptor instead.
func (*PipelineLabelBatchInsertResponse) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{2}
}

type PipelineLabelListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineSource  string   `protobuf:"bytes,1,opt,name=pipelineSource,proto3" json:"pipelineSource,omitempty"`
	PipelineYmlName string   `protobuf:"bytes,2,opt,name=pipelineYmlName,proto3" json:"pipelineYmlName,omitempty"`
	TargetIDs       []uint64 `protobuf:"varint,3,rep,packed,name=targetIDs,proto3" json:"targetIDs,omitempty"`
	MatchKeys       []string `protobuf:"bytes,4,rep,name=matchKeys,proto3" json:"matchKeys,omitempty"`
	UnMatchKeys     []string `protobuf:"bytes,5,rep,name=unMatchKeys,proto3" json:"unMatchKeys,omitempty"`
}

func (x *PipelineLabelListRequest) Reset() {
	*x = PipelineLabelListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineLabelListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineLabelListRequest) ProtoMessage() {}

func (x *PipelineLabelListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineLabelListRequest.ProtoReflect.Descriptor instead.
func (*PipelineLabelListRequest) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{3}
}

func (x *PipelineLabelListRequest) GetPipelineSource() string {
	if x != nil {
		return x.PipelineSource
	}
	return ""
}

func (x *PipelineLabelListRequest) GetPipelineYmlName() string {
	if x != nil {
		return x.PipelineYmlName
	}
	return ""
}

func (x *PipelineLabelListRequest) GetTargetIDs() []uint64 {
	if x != nil {
		return x.TargetIDs
	}
	return nil
}

func (x *PipelineLabelListRequest) GetMatchKeys() []string {
	if x != nil {
		return x.MatchKeys
	}
	return nil
}

func (x *PipelineLabelListRequest) GetUnMatchKeys() []string {
	if x != nil {
		return x.UnMatchKeys
	}
	return nil
}

type PipelineLabelListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels []*PipelineLabel `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Total  int64            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PipelineLabelListResponse) Reset() {
	*x = PipelineLabelListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_label_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineLabelListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineLabelListResponse) ProtoMessage() {}

func (x *PipelineLabelListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_label_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineLabelListResponse.ProtoReflect.Descriptor instead.
func (*PipelineLabelListResponse) Descriptor() ([]byte, []int) {
	return file_label_proto_rawDescGZIP(), []int{4}
}

func (x *PipelineLabelListResponse) GetLabels() []*PipelineLabel {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *PipelineLabelListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_label_proto protoreflect.FileDescriptor

var file_label_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65,
	0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x0d, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x22, 0x62, 0x0a, 0x1f, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x22, 0x22, 0x0a, 0x20, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x18, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x59, 0x6d, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x59, 0x6d, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x44, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65,
	0x79, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4b,
	0x65, 0x79, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65,
	0x79, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x72, 0x0a, 0x19, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xfb, 0x02, 0x0a, 0x0c, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xcc, 0x01, 0x0a, 0x18, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x39, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x22, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x69, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x2d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x9b, 0x01, 0x0a, 0x11, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x32, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x2d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_label_proto_rawDescOnce sync.Once
	file_label_proto_rawDescData = file_label_proto_rawDesc
)

func file_label_proto_rawDescGZIP() []byte {
	file_label_proto_rawDescOnce.Do(func() {
		file_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_label_proto_rawDescData)
	})
	return file_label_proto_rawDescData
}

var file_label_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_label_proto_goTypes = []interface{}{
	(*PipelineLabel)(nil),                    // 0: erda.core.pipeline.label.PipelineLabel
	(*PipelineLabelBatchInsertRequest)(nil),  // 1: erda.core.pipeline.label.PipelineLabelBatchInsertRequest
	(*PipelineLabelBatchInsertResponse)(nil), // 2: erda.core.pipeline.label.PipelineLabelBatchInsertResponse
	(*PipelineLabelListRequest)(nil),         // 3: erda.core.pipeline.label.PipelineLabelListRequest
	(*PipelineLabelListResponse)(nil),        // 4: erda.core.pipeline.label.PipelineLabelListResponse
	(*timestamppb.Timestamp)(nil),            // 5: google.protobuf.Timestamp
}
var file_label_proto_depIdxs = []int32{
	5, // 0: erda.core.pipeline.label.PipelineLabel.timeCreated:type_name -> google.protobuf.Timestamp
	5, // 1: erda.core.pipeline.label.PipelineLabel.timeUpdated:type_name -> google.protobuf.Timestamp
	0, // 2: erda.core.pipeline.label.PipelineLabelBatchInsertRequest.labels:type_name -> erda.core.pipeline.label.PipelineLabel
	0, // 3: erda.core.pipeline.label.PipelineLabelListResponse.labels:type_name -> erda.core.pipeline.label.PipelineLabel
	1, // 4: erda.core.pipeline.label.LabelService.PipelineLabelBatchInsert:input_type -> erda.core.pipeline.label.PipelineLabelBatchInsertRequest
	3, // 5: erda.core.pipeline.label.LabelService.PipelineLabelList:input_type -> erda.core.pipeline.label.PipelineLabelListRequest
	2, // 6: erda.core.pipeline.label.LabelService.PipelineLabelBatchInsert:output_type -> erda.core.pipeline.label.PipelineLabelBatchInsertResponse
	4, // 7: erda.core.pipeline.label.LabelService.PipelineLabelList:output_type -> erda.core.pipeline.label.PipelineLabelListResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_label_proto_init() }
func file_label_proto_init() {
	if File_label_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_label_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineLabel); i {
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
		file_label_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineLabelBatchInsertRequest); i {
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
		file_label_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineLabelBatchInsertResponse); i {
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
		file_label_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineLabelListRequest); i {
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
		file_label_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineLabelListResponse); i {
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
			RawDescriptor: file_label_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_label_proto_goTypes,
		DependencyIndexes: file_label_proto_depIdxs,
		MessageInfos:      file_label_proto_msgTypes,
	}.Build()
	File_label_proto = out.File
	file_label_proto_rawDesc = nil
	file_label_proto_goTypes = nil
	file_label_proto_depIdxs = nil
}