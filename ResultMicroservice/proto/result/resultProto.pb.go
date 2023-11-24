// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: resultProto.proto

package result

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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ExerciseID  string `protobuf:"bytes,2,opt,name=ExerciseID,proto3" json:"ExerciseID,omitempty"`
	UserID      string `protobuf:"bytes,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ClassID     string `protobuf:"bytes,4,opt,name=ClassID,proto3" json:"ClassID,omitempty"`
	ModuleID    string `protobuf:"bytes,5,opt,name=ModuleID,proto3" json:"ModuleID,omitempty"`
	Input       string `protobuf:"bytes,6,opt,name=Input,proto3" json:"Input,omitempty"`
	Result      string `protobuf:"bytes,7,opt,name=Result,proto3" json:"Result,omitempty"`
	CreatedAt   string `protobuf:"bytes,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	SoftDeleted bool   `protobuf:"varint,10,opt,name=SoftDeleted,proto3" json:"SoftDeleted,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resultProto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_resultProto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_resultProto_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Result) GetExerciseID() string {
	if x != nil {
		return x.ExerciseID
	}
	return ""
}

func (x *Result) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Result) GetClassID() string {
	if x != nil {
		return x.ClassID
	}
	return ""
}

func (x *Result) GetModuleID() string {
	if x != nil {
		return x.ModuleID
	}
	return ""
}

func (x *Result) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *Result) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *Result) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Result) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Result) GetSoftDeleted() bool {
	if x != nil {
		return x.SoftDeleted
	}
	return false
}

type ModuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModuleID    string `protobuf:"bytes,1,opt,name=ModuleID,proto3" json:"ModuleID,omitempty"`
	BearerToken string `protobuf:"bytes,2,opt,name=BearerToken,proto3" json:"BearerToken,omitempty"`
}

func (x *ModuleRequest) Reset() {
	*x = ModuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resultProto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleRequest) ProtoMessage() {}

func (x *ModuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_resultProto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleRequest.ProtoReflect.Descriptor instead.
func (*ModuleRequest) Descriptor() ([]byte, []int) {
	return file_resultProto_proto_rawDescGZIP(), []int{1}
}

func (x *ModuleRequest) GetModuleID() string {
	if x != nil {
		return x.ModuleID
	}
	return ""
}

func (x *ModuleRequest) GetBearerToken() string {
	if x != nil {
		return x.BearerToken
	}
	return ""
}

var File_resultProto_proto protoreflect.FileDescriptor

var file_resultProto_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x6f, 0x66, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x53, 0x6f, 0x66,
	0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x4d, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x55, 0x0a, 0x09, 0x4d, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e,
	0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x53, 0x6f, 0x66, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x09,
	0x5a, 0x07, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_resultProto_proto_rawDescOnce sync.Once
	file_resultProto_proto_rawDescData = file_resultProto_proto_rawDesc
)

func file_resultProto_proto_rawDescGZIP() []byte {
	file_resultProto_proto_rawDescOnce.Do(func() {
		file_resultProto_proto_rawDescData = protoimpl.X.CompressGZIP(file_resultProto_proto_rawDescData)
	})
	return file_resultProto_proto_rawDescData
}

var file_resultProto_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resultProto_proto_goTypes = []interface{}{
	(*Result)(nil),        // 0: Result
	(*ModuleRequest)(nil), // 1: ModuleRequest
}
var file_resultProto_proto_depIdxs = []int32{
	1, // 0: MyService.Delete:input_type -> ModuleRequest
	1, // 1: MyService.SoftDelete:input_type -> ModuleRequest
	0, // 2: MyService.Delete:output_type -> Result
	0, // 3: MyService.SoftDelete:output_type -> Result
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resultProto_proto_init() }
func file_resultProto_proto_init() {
	if File_resultProto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resultProto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_resultProto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleRequest); i {
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
			RawDescriptor: file_resultProto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resultProto_proto_goTypes,
		DependencyIndexes: file_resultProto_proto_depIdxs,
		MessageInfos:      file_resultProto_proto_msgTypes,
	}.Build()
	File_resultProto_proto = out.File
	file_resultProto_proto_rawDesc = nil
	file_resultProto_proto_goTypes = nil
	file_resultProto_proto_depIdxs = nil
}