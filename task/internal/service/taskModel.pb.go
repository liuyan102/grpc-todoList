// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: taskModel.proto

package service

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

type TaskModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"task_id"
	TaskID uint32 `protobuf:"varint,1,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	// @inject_tag: json:"user_id"
	UserID uint32 `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	// @inject_tag: json:"status"
	Status uint32 `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,4,opt,name=Title,proto3" json:"Title,omitempty"`
	// @inject_tag: json:"content"
	Content string `protobuf:"bytes,5,opt,name=Content,proto3" json:"Content,omitempty"`
	// @inject_tag: json:"start_item"
	StartTime uint32 `protobuf:"varint,6,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	// @inject_tag: json:"end_time"
	EndTime uint32 `protobuf:"varint,7,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (x *TaskModel) Reset() {
	*x = TaskModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_taskModel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskModel) ProtoMessage() {}

func (x *TaskModel) ProtoReflect() protoreflect.Message {
	mi := &file_taskModel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskModel.ProtoReflect.Descriptor instead.
func (*TaskModel) Descriptor() ([]byte, []int) {
	return file_taskModel_proto_rawDescGZIP(), []int{0}
}

func (x *TaskModel) GetTaskID() uint32 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *TaskModel) GetUserID() uint32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *TaskModel) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TaskModel) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TaskModel) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *TaskModel) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

var File_taskModel_proto protoreflect.FileDescriptor

var file_taskModel_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42,
	0x1b, 0x5a, 0x19, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_taskModel_proto_rawDescOnce sync.Once
	file_taskModel_proto_rawDescData = file_taskModel_proto_rawDesc
)

func file_taskModel_proto_rawDescGZIP() []byte {
	file_taskModel_proto_rawDescOnce.Do(func() {
		file_taskModel_proto_rawDescData = protoimpl.X.CompressGZIP(file_taskModel_proto_rawDescData)
	})
	return file_taskModel_proto_rawDescData
}

var file_taskModel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_taskModel_proto_goTypes = []interface{}{
	(*TaskModel)(nil), // 0: TaskModel
}
var file_taskModel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_taskModel_proto_init() }
func file_taskModel_proto_init() {
	if File_taskModel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_taskModel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskModel); i {
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
			RawDescriptor: file_taskModel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_taskModel_proto_goTypes,
		DependencyIndexes: file_taskModel_proto_depIdxs,
		MessageInfos:      file_taskModel_proto_msgTypes,
	}.Build()
	File_taskModel_proto = out.File
	file_taskModel_proto_rawDesc = nil
	file_taskModel_proto_goTypes = nil
	file_taskModel_proto_depIdxs = nil
}
