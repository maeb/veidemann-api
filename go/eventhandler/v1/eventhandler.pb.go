// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: eventhandler/v1/eventhandler.proto

package eventhandler

import (
	v1 "github.com/maeb/veidemann-api/go/commons/v1"
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

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Select objects by one or more id's
	Id                 []string      `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	QueryTemplate      *EventObject  `protobuf:"bytes,2,opt,name=query_template,json=queryTemplate,proto3" json:"query_template,omitempty"`
	QueryMask          *v1.FieldMask `protobuf:"bytes,3,opt,name=query_mask,json=queryMask,proto3" json:"query_mask,omitempty"`
	ReturnedFieldsMask *v1.FieldMask `protobuf:"bytes,4,opt,name=returned_fields_mask,json=returnedFieldsMask,proto3" json:"returned_fields_mask,omitempty"`
	PageSize           int32         `protobuf:"varint,7,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Offset             int32         `protobuf:"varint,8,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_eventhandler_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ListRequest) GetQueryTemplate() *EventObject {
	if x != nil {
		return x.QueryTemplate
	}
	return nil
}

func (x *ListRequest) GetQueryMask() *v1.FieldMask {
	if x != nil {
		return x.QueryMask
	}
	return nil
}

func (x *ListRequest) GetReturnedFieldsMask() *v1.FieldMask {
	if x != nil {
		return x.ReturnedFieldsMask
	}
	return nil
}

func (x *ListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListRequest    *ListRequest  `protobuf:"bytes,1,opt,name=list_request,json=listRequest,proto3" json:"list_request,omitempty"`
	UpdateMask     *v1.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	UpdateTemplate *EventObject  `protobuf:"bytes,4,opt,name=update_template,json=updateTemplate,proto3" json:"update_template,omitempty"`
	Comment        string        `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_eventhandler_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRequest) GetListRequest() *ListRequest {
	if x != nil {
		return x.ListRequest
	}
	return nil
}

func (x *UpdateRequest) GetUpdateMask() *v1.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateRequest) GetUpdateTemplate() *EventObject {
	if x != nil {
		return x.UpdateTemplate
	}
	return nil
}

func (x *UpdateRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object  *EventObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Comment string       `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_eventhandler_proto_rawDescGZIP(), []int{2}
}

func (x *SaveRequest) GetObject() *EventObject {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *SaveRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated int64 `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_eventhandler_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResponse) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type ListCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Approximate bool  `protobuf:"varint,2,opt,name=approximate,proto3" json:"approximate,omitempty"`
}

func (x *ListCountResponse) Reset() {
	*x = ListCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountResponse) ProtoMessage() {}

func (x *ListCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountResponse.ProtoReflect.Descriptor instead.
func (*ListCountResponse) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_eventhandler_proto_rawDescGZIP(), []int{4}
}

func (x *ListCountResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListCountResponse) GetApproximate() bool {
	if x != nil {
		return x.Approximate
	}
	return false
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_eventhandler_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type ListLabelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ListLabelRequest) Reset() {
	*x = ListLabelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLabelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLabelRequest) ProtoMessage() {}

func (x *ListLabelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLabelRequest.ProtoReflect.Descriptor instead.
func (*ListLabelRequest) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_eventhandler_proto_rawDescGZIP(), []int{6}
}

func (x *ListLabelRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ListLabelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label []string `protobuf:"bytes,1,rep,name=label,proto3" json:"label,omitempty"`
}

func (x *ListLabelResponse) Reset() {
	*x = ListLabelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLabelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLabelResponse) ProtoMessage() {}

func (x *ListLabelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_eventhandler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLabelResponse.ProtoReflect.Descriptor instead.
func (*ListLabelResponse) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_eventhandler_proto_rawDescGZIP(), []int{7}
}

func (x *ListLabelResponse) GetLabel() []string {
	if x != nil {
		return x.Label
	}
	return nil
}

var File_eventhandler_v1_eventhandler_proto protoreflect.FileDescriptor

var file_eventhandler_v1_eventhandler_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc0, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x51, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x55, 0x0a, 0x14, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x12, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x22, 0x93, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x53, 0x0a, 0x0f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x0b, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x22, 0x4b, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x22,
	0x2a, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x26, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x32, 0xa1,
	0x06, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12,
	0x67, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x27, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x1a, 0x2a, 0x2e, 0x76, 0x65, 0x69,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x73, 0x0a, 0x11, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x2a, 0x2e,
	0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a,
	0x0f, 0x53, 0x61, 0x76, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x12, 0x2c, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x70, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x2a, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x1a, 0x2d, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x2f, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x7f, 0x0a, 0x27, 0x6e, 0x6f, 0x2e, 0x6e, 0x62, 0x2e, 0x6e, 0x6e, 0x61, 0x2e,
	0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x61, 0x65, 0x62, 0x2f, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventhandler_v1_eventhandler_proto_rawDescOnce sync.Once
	file_eventhandler_v1_eventhandler_proto_rawDescData = file_eventhandler_v1_eventhandler_proto_rawDesc
)

func file_eventhandler_v1_eventhandler_proto_rawDescGZIP() []byte {
	file_eventhandler_v1_eventhandler_proto_rawDescOnce.Do(func() {
		file_eventhandler_v1_eventhandler_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventhandler_v1_eventhandler_proto_rawDescData)
	})
	return file_eventhandler_v1_eventhandler_proto_rawDescData
}

var file_eventhandler_v1_eventhandler_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_eventhandler_v1_eventhandler_proto_goTypes = []interface{}{
	(*ListRequest)(nil),       // 0: veidemann.api.eventhandler.v1.ListRequest
	(*UpdateRequest)(nil),     // 1: veidemann.api.eventhandler.v1.UpdateRequest
	(*SaveRequest)(nil),       // 2: veidemann.api.eventhandler.v1.SaveRequest
	(*UpdateResponse)(nil),    // 3: veidemann.api.eventhandler.v1.UpdateResponse
	(*ListCountResponse)(nil), // 4: veidemann.api.eventhandler.v1.ListCountResponse
	(*DeleteResponse)(nil),    // 5: veidemann.api.eventhandler.v1.DeleteResponse
	(*ListLabelRequest)(nil),  // 6: veidemann.api.eventhandler.v1.ListLabelRequest
	(*ListLabelResponse)(nil), // 7: veidemann.api.eventhandler.v1.ListLabelResponse
	(*EventObject)(nil),       // 8: veidemann.api.eventhandler.v1.EventObject
	(*v1.FieldMask)(nil),      // 9: veidemann.api.commons.v1.FieldMask
	(*EventRef)(nil),          // 10: veidemann.api.eventhandler.v1.EventRef
}
var file_eventhandler_v1_eventhandler_proto_depIdxs = []int32{
	8,  // 0: veidemann.api.eventhandler.v1.ListRequest.query_template:type_name -> veidemann.api.eventhandler.v1.EventObject
	9,  // 1: veidemann.api.eventhandler.v1.ListRequest.query_mask:type_name -> veidemann.api.commons.v1.FieldMask
	9,  // 2: veidemann.api.eventhandler.v1.ListRequest.returned_fields_mask:type_name -> veidemann.api.commons.v1.FieldMask
	0,  // 3: veidemann.api.eventhandler.v1.UpdateRequest.list_request:type_name -> veidemann.api.eventhandler.v1.ListRequest
	9,  // 4: veidemann.api.eventhandler.v1.UpdateRequest.update_mask:type_name -> veidemann.api.commons.v1.FieldMask
	8,  // 5: veidemann.api.eventhandler.v1.UpdateRequest.update_template:type_name -> veidemann.api.eventhandler.v1.EventObject
	8,  // 6: veidemann.api.eventhandler.v1.SaveRequest.object:type_name -> veidemann.api.eventhandler.v1.EventObject
	10, // 7: veidemann.api.eventhandler.v1.EventHandler.GetEventObject:input_type -> veidemann.api.eventhandler.v1.EventRef
	0,  // 8: veidemann.api.eventhandler.v1.EventHandler.ListEventObjects:input_type -> veidemann.api.eventhandler.v1.ListRequest
	0,  // 9: veidemann.api.eventhandler.v1.EventHandler.CountEventObjects:input_type -> veidemann.api.eventhandler.v1.ListRequest
	2,  // 10: veidemann.api.eventhandler.v1.EventHandler.SaveEventObject:input_type -> veidemann.api.eventhandler.v1.SaveRequest
	1,  // 11: veidemann.api.eventhandler.v1.EventHandler.UpdateEventObjects:input_type -> veidemann.api.eventhandler.v1.UpdateRequest
	8,  // 12: veidemann.api.eventhandler.v1.EventHandler.DeleteEventObject:input_type -> veidemann.api.eventhandler.v1.EventObject
	6,  // 13: veidemann.api.eventhandler.v1.EventHandler.ListLabels:input_type -> veidemann.api.eventhandler.v1.ListLabelRequest
	8,  // 14: veidemann.api.eventhandler.v1.EventHandler.GetEventObject:output_type -> veidemann.api.eventhandler.v1.EventObject
	8,  // 15: veidemann.api.eventhandler.v1.EventHandler.ListEventObjects:output_type -> veidemann.api.eventhandler.v1.EventObject
	4,  // 16: veidemann.api.eventhandler.v1.EventHandler.CountEventObjects:output_type -> veidemann.api.eventhandler.v1.ListCountResponse
	8,  // 17: veidemann.api.eventhandler.v1.EventHandler.SaveEventObject:output_type -> veidemann.api.eventhandler.v1.EventObject
	3,  // 18: veidemann.api.eventhandler.v1.EventHandler.UpdateEventObjects:output_type -> veidemann.api.eventhandler.v1.UpdateResponse
	5,  // 19: veidemann.api.eventhandler.v1.EventHandler.DeleteEventObject:output_type -> veidemann.api.eventhandler.v1.DeleteResponse
	7,  // 20: veidemann.api.eventhandler.v1.EventHandler.ListLabels:output_type -> veidemann.api.eventhandler.v1.ListLabelResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_eventhandler_v1_eventhandler_proto_init() }
func file_eventhandler_v1_eventhandler_proto_init() {
	if File_eventhandler_v1_eventhandler_proto != nil {
		return
	}
	file_eventhandler_v1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eventhandler_v1_eventhandler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_eventhandler_v1_eventhandler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_eventhandler_v1_eventhandler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest); i {
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
		file_eventhandler_v1_eventhandler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_eventhandler_v1_eventhandler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCountResponse); i {
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
		file_eventhandler_v1_eventhandler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_eventhandler_v1_eventhandler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLabelRequest); i {
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
		file_eventhandler_v1_eventhandler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLabelResponse); i {
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
			RawDescriptor: file_eventhandler_v1_eventhandler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eventhandler_v1_eventhandler_proto_goTypes,
		DependencyIndexes: file_eventhandler_v1_eventhandler_proto_depIdxs,
		MessageInfos:      file_eventhandler_v1_eventhandler_proto_msgTypes,
	}.Build()
	File_eventhandler_v1_eventhandler_proto = out.File
	file_eventhandler_v1_eventhandler_proto_rawDesc = nil
	file_eventhandler_v1_eventhandler_proto_goTypes = nil
	file_eventhandler_v1_eventhandler_proto_depIdxs = nil
}
