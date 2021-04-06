// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: eventhandler/v1/resources.proto

package eventhandler

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EventObject_State int32

const (
	EventObject_NEW    EventObject_State = 0
	EventObject_OPEN   EventObject_State = 1
	EventObject_CLOSED EventObject_State = 2
)

// Enum value maps for EventObject_State.
var (
	EventObject_State_name = map[int32]string{
		0: "NEW",
		1: "OPEN",
		2: "CLOSED",
	}
	EventObject_State_value = map[string]int32{
		"NEW":    0,
		"OPEN":   1,
		"CLOSED": 2,
	}
)

func (x EventObject_State) Enum() *EventObject_State {
	p := new(EventObject_State)
	*p = x
	return p
}

func (x EventObject_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventObject_State) Descriptor() protoreflect.EnumDescriptor {
	return file_eventhandler_v1_resources_proto_enumTypes[0].Descriptor()
}

func (EventObject_State) Type() protoreflect.EnumType {
	return &file_eventhandler_v1_resources_proto_enumTypes[0]
}

func (x EventObject_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventObject_State.Descriptor instead.
func (EventObject_State) EnumDescriptor() ([]byte, []int) {
	return file_eventhandler_v1_resources_proto_rawDescGZIP(), []int{0, 0}
}

type EventObject_Severity int32

const (
	EventObject_INFO  EventObject_Severity = 0
	EventObject_WARN  EventObject_Severity = 1
	EventObject_ERROR EventObject_Severity = 2
)

// Enum value maps for EventObject_Severity.
var (
	EventObject_Severity_name = map[int32]string{
		0: "INFO",
		1: "WARN",
		2: "ERROR",
	}
	EventObject_Severity_value = map[string]int32{
		"INFO":  0,
		"WARN":  1,
		"ERROR": 2,
	}
)

func (x EventObject_Severity) Enum() *EventObject_Severity {
	p := new(EventObject_Severity)
	*p = x
	return p
}

func (x EventObject_Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventObject_Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_eventhandler_v1_resources_proto_enumTypes[1].Descriptor()
}

func (EventObject_Severity) Type() protoreflect.EnumType {
	return &file_eventhandler_v1_resources_proto_enumTypes[1]
}

func (x EventObject_Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventObject_Severity.Descriptor instead.
func (EventObject_Severity) EnumDescriptor() ([]byte, []int) {
	return file_eventhandler_v1_resources_proto_rawDescGZIP(), []int{0, 1}
}

type Activity_ChangeType int32

const (
	Activity_CREATED       Activity_ChangeType = 0
	Activity_VALUE_CHANGED Activity_ChangeType = 1
	Activity_ARRAY_ADD     Activity_ChangeType = 2
	Activity_ARRAY_DEL     Activity_ChangeType = 3
)

// Enum value maps for Activity_ChangeType.
var (
	Activity_ChangeType_name = map[int32]string{
		0: "CREATED",
		1: "VALUE_CHANGED",
		2: "ARRAY_ADD",
		3: "ARRAY_DEL",
	}
	Activity_ChangeType_value = map[string]int32{
		"CREATED":       0,
		"VALUE_CHANGED": 1,
		"ARRAY_ADD":     2,
		"ARRAY_DEL":     3,
	}
)

func (x Activity_ChangeType) Enum() *Activity_ChangeType {
	p := new(Activity_ChangeType)
	*p = x
	return p
}

func (x Activity_ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Activity_ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_eventhandler_v1_resources_proto_enumTypes[2].Descriptor()
}

func (Activity_ChangeType) Type() protoreflect.EnumType {
	return &file_eventhandler_v1_resources_proto_enumTypes[2]
}

func (x Activity_ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Activity_ChangeType.Descriptor instead.
func (Activity_ChangeType) EnumDescriptor() ([]byte, []int) {
	return file_eventhandler_v1_resources_proto_rawDescGZIP(), []int{2, 0}
}

type EventObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type     string               `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Source   string               `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	State    EventObject_State    `protobuf:"varint,4,opt,name=state,proto3,enum=veidemann.api.eventhandler.v1.EventObject_State" json:"state,omitempty"`
	Assignee string               `protobuf:"bytes,5,opt,name=assignee,proto3" json:"assignee,omitempty"`
	Data     []*Data              `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty"`
	Severity EventObject_Severity `protobuf:"varint,7,opt,name=severity,proto3,enum=veidemann.api.eventhandler.v1.EventObject_Severity" json:"severity,omitempty"`
	Label    []string             `protobuf:"bytes,8,rep,name=label,proto3" json:"label,omitempty"`
	Activity []*Activity          `protobuf:"bytes,9,rep,name=activity,proto3" json:"activity,omitempty"`
}

func (x *EventObject) Reset() {
	*x = EventObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventObject) ProtoMessage() {}

func (x *EventObject) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventObject.ProtoReflect.Descriptor instead.
func (*EventObject) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *EventObject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventObject) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EventObject) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *EventObject) GetState() EventObject_State {
	if x != nil {
		return x.State
	}
	return EventObject_NEW
}

func (x *EventObject) GetAssignee() string {
	if x != nil {
		return x.Assignee
	}
	return ""
}

func (x *EventObject) GetData() []*Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *EventObject) GetSeverity() EventObject_Severity {
	if x != nil {
		return x.Severity
	}
	return EventObject_INFO
}

func (x *EventObject) GetLabel() []string {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *EventObject) GetActivity() []*Activity {
	if x != nil {
		return x.Activity
	}
	return nil
}

type EventRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the EventObject to get
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EventRef) Reset() {
	*x = EventRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRef) ProtoMessage() {}

func (x *EventRef) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRef.ProtoReflect.Descriptor instead.
func (*EventRef) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_resources_proto_rawDescGZIP(), []int{1}
}

func (x *EventRef) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModifiedBy   string               `protobuf:"bytes,1,opt,name=modified_by,json=modifiedBy,proto3" json:"modified_by,omitempty"`
	ModifiedTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=modified_time,json=modifiedTime,proto3" json:"modified_time,omitempty"`
	Description  []*Activity_Change   `protobuf:"bytes,3,rep,name=description,proto3" json:"description,omitempty"`
	Comment      string               `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_resources_proto_rawDescGZIP(), []int{2}
}

func (x *Activity) GetModifiedBy() string {
	if x != nil {
		return x.ModifiedBy
	}
	return ""
}

func (x *Activity) GetModifiedTime() *timestamp.Timestamp {
	if x != nil {
		return x.ModifiedTime
	}
	return nil
}

func (x *Activity) GetDescription() []*Activity_Change {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Activity) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_resources_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Data) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Activity_Change struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   Activity_ChangeType `protobuf:"varint,1,opt,name=type,proto3,enum=veidemann.api.eventhandler.v1.Activity_ChangeType" json:"type,omitempty"`
	Field  string              `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	OldVal string              `protobuf:"bytes,3,opt,name=old_val,json=oldVal,proto3" json:"old_val,omitempty"`
	NewVal string              `protobuf:"bytes,4,opt,name=new_val,json=newVal,proto3" json:"new_val,omitempty"`
}

func (x *Activity_Change) Reset() {
	*x = Activity_Change{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventhandler_v1_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity_Change) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity_Change) ProtoMessage() {}

func (x *Activity_Change) ProtoReflect() protoreflect.Message {
	mi := &file_eventhandler_v1_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity_Change.ProtoReflect.Descriptor instead.
func (*Activity_Change) Descriptor() ([]byte, []int) {
	return file_eventhandler_v1_resources_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Activity_Change) GetType() Activity_ChangeType {
	if x != nil {
		return x.Type
	}
	return Activity_CREATED
}

func (x *Activity_Change) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *Activity_Change) GetOldVal() string {
	if x != nil {
		return x.OldVal
	}
	return ""
}

func (x *Activity_Change) GetNewVal() string {
	if x != nil {
		return x.NewVal
	}
	return ""
}

var File_eventhandler_v1_resources_proto protoreflect.FileDescriptor

var file_eventhandler_v1_resources_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe5, 0x03, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x65, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4f, 0x0a, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x43, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x08, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x02, 0x22, 0x29,
	0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e,
	0x46, 0x4f, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x22, 0x1a, 0x0a, 0x08, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbf, 0x03, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x3f, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x98, 0x01, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x6c, 0x64,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6c, 0x64, 0x56,
	0x61, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x22, 0x4a, 0x0a, 0x0a, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x52, 0x52,
	0x41, 0x59, 0x5f, 0x41, 0x44, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x52, 0x52, 0x41,
	0x59, 0x5f, 0x44, 0x45, 0x4c, 0x10, 0x03, 0x22, 0x2e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x82, 0x01, 0x0a, 0x27, 0x6e, 0x6f, 0x2e, 0x6e,
	0x62, 0x2e, 0x6e, 0x6e, 0x61, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x42, 0x15, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventhandler_v1_resources_proto_rawDescOnce sync.Once
	file_eventhandler_v1_resources_proto_rawDescData = file_eventhandler_v1_resources_proto_rawDesc
)

func file_eventhandler_v1_resources_proto_rawDescGZIP() []byte {
	file_eventhandler_v1_resources_proto_rawDescOnce.Do(func() {
		file_eventhandler_v1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventhandler_v1_resources_proto_rawDescData)
	})
	return file_eventhandler_v1_resources_proto_rawDescData
}

var file_eventhandler_v1_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_eventhandler_v1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_eventhandler_v1_resources_proto_goTypes = []interface{}{
	(EventObject_State)(0),      // 0: veidemann.api.eventhandler.v1.EventObject.State
	(EventObject_Severity)(0),   // 1: veidemann.api.eventhandler.v1.EventObject.Severity
	(Activity_ChangeType)(0),    // 2: veidemann.api.eventhandler.v1.Activity.ChangeType
	(*EventObject)(nil),         // 3: veidemann.api.eventhandler.v1.EventObject
	(*EventRef)(nil),            // 4: veidemann.api.eventhandler.v1.EventRef
	(*Activity)(nil),            // 5: veidemann.api.eventhandler.v1.Activity
	(*Data)(nil),                // 6: veidemann.api.eventhandler.v1.Data
	(*Activity_Change)(nil),     // 7: veidemann.api.eventhandler.v1.Activity.Change
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_eventhandler_v1_resources_proto_depIdxs = []int32{
	0, // 0: veidemann.api.eventhandler.v1.EventObject.state:type_name -> veidemann.api.eventhandler.v1.EventObject.State
	6, // 1: veidemann.api.eventhandler.v1.EventObject.data:type_name -> veidemann.api.eventhandler.v1.Data
	1, // 2: veidemann.api.eventhandler.v1.EventObject.severity:type_name -> veidemann.api.eventhandler.v1.EventObject.Severity
	5, // 3: veidemann.api.eventhandler.v1.EventObject.activity:type_name -> veidemann.api.eventhandler.v1.Activity
	8, // 4: veidemann.api.eventhandler.v1.Activity.modified_time:type_name -> google.protobuf.Timestamp
	7, // 5: veidemann.api.eventhandler.v1.Activity.description:type_name -> veidemann.api.eventhandler.v1.Activity.Change
	2, // 6: veidemann.api.eventhandler.v1.Activity.Change.type:type_name -> veidemann.api.eventhandler.v1.Activity.ChangeType
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_eventhandler_v1_resources_proto_init() }
func file_eventhandler_v1_resources_proto_init() {
	if File_eventhandler_v1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventhandler_v1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventObject); i {
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
		file_eventhandler_v1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRef); i {
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
		file_eventhandler_v1_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_eventhandler_v1_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_eventhandler_v1_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity_Change); i {
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
			RawDescriptor: file_eventhandler_v1_resources_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eventhandler_v1_resources_proto_goTypes,
		DependencyIndexes: file_eventhandler_v1_resources_proto_depIdxs,
		EnumInfos:         file_eventhandler_v1_resources_proto_enumTypes,
		MessageInfos:      file_eventhandler_v1_resources_proto_msgTypes,
	}.Build()
	File_eventhandler_v1_resources_proto = out.File
	file_eventhandler_v1_resources_proto_rawDesc = nil
	file_eventhandler_v1_resources_proto_goTypes = nil
	file_eventhandler_v1_resources_proto_depIdxs = nil
}
