// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: proto/Textile.proto

package textile

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

type TechnicalModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is given after the model is created
	Id          *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type        string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Created     string  `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated     string  `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *TechnicalModel) Reset() {
	*x = TechnicalModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Textile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TechnicalModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TechnicalModel) ProtoMessage() {}

func (x *TechnicalModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Textile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TechnicalModel.ProtoReflect.Descriptor instead.
func (*TechnicalModel) Descriptor() ([]byte, []int) {
	return file_proto_Textile_proto_rawDescGZIP(), []int{0}
}

func (x *TechnicalModel) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *TechnicalModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TechnicalModel) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TechnicalModel) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TechnicalModel) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *TechnicalModel) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

var File_proto_Textile_proto protoreflect.FileDescriptor

var file_proto_Textile_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x69, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31,
	0x2e, 0x74, 0x65, 0x78, 0x74, 0x69, 0x6c, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x54, 0x65, 0x63,
	0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x13, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x69, 0x64, 0x32, 0x67, 0x0a, 0x07, 0x54, 0x65, 0x78, 0x74, 0x69, 0x6c, 0x65,
	0x12, 0x5c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69,
	0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x69, 0x6c, 0x65, 0x2e, 0x54, 0x65, 0x63,
	0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x21, 0x2e, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x69, 0x6c, 0x65, 0x2e,
	0x54, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x10,
	0x5a, 0x0e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x69, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_Textile_proto_rawDescOnce sync.Once
	file_proto_Textile_proto_rawDescData = file_proto_Textile_proto_rawDesc
)

func file_proto_Textile_proto_rawDescGZIP() []byte {
	file_proto_Textile_proto_rawDescOnce.Do(func() {
		file_proto_Textile_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Textile_proto_rawDescData)
	})
	return file_proto_Textile_proto_rawDescData
}

var file_proto_Textile_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_Textile_proto_goTypes = []interface{}{
	(*TechnicalModel)(nil), // 0: hestia.v1.textile.TechnicalModel
}
var file_proto_Textile_proto_depIdxs = []int32{
	0, // 0: hestia.v1.textile.Textile.CreateTechnicalModel:input_type -> hestia.v1.textile.TechnicalModel
	0, // 1: hestia.v1.textile.Textile.CreateTechnicalModel:output_type -> hestia.v1.textile.TechnicalModel
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_Textile_proto_init() }
func file_proto_Textile_proto_init() {
	if File_proto_Textile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_Textile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TechnicalModel); i {
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
	file_proto_Textile_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_Textile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_Textile_proto_goTypes,
		DependencyIndexes: file_proto_Textile_proto_depIdxs,
		MessageInfos:      file_proto_Textile_proto_msgTypes,
	}.Build()
	File_proto_Textile_proto = out.File
	file_proto_Textile_proto_rawDesc = nil
	file_proto_Textile_proto_goTypes = nil
	file_proto_Textile_proto_depIdxs = nil
}
