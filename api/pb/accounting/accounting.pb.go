// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: accounting.proto

package accounting

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VatRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Rate    uint32 `protobuf:"varint,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Code    string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Country string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *VatRate) Reset() {
	*x = VatRate{}
	mi := &file_accounting_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VatRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VatRate) ProtoMessage() {}

func (x *VatRate) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VatRate.ProtoReflect.Descriptor instead.
func (*VatRate) Descriptor() ([]byte, []int) {
	return file_accounting_proto_rawDescGZIP(), []int{0}
}

func (x *VatRate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VatRate) GetRate() uint32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *VatRate) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VatRate) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type VatRates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VatRates []*VatRate `protobuf:"bytes,1,rep,name=vat_rates,json=vatRates,proto3" json:"vat_rates,omitempty"`
}

func (x *VatRates) Reset() {
	*x = VatRates{}
	mi := &file_accounting_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VatRates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VatRates) ProtoMessage() {}

func (x *VatRates) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VatRates.ProtoReflect.Descriptor instead.
func (*VatRates) Descriptor() ([]byte, []int) {
	return file_accounting_proto_rawDescGZIP(), []int{1}
}

func (x *VatRates) GetVatRates() []*VatRate {
	if x != nil {
		return x.VatRates
	}
	return nil
}

type VatExemption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *VatExemption) Reset() {
	*x = VatExemption{}
	mi := &file_accounting_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VatExemption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VatExemption) ProtoMessage() {}

func (x *VatExemption) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VatExemption.ProtoReflect.Descriptor instead.
func (*VatExemption) Descriptor() ([]byte, []int) {
	return file_accounting_proto_rawDescGZIP(), []int{2}
}

func (x *VatExemption) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VatExemption) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *VatExemption) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type VatExemptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VatExemptions []*VatExemption `protobuf:"bytes,1,rep,name=vat_exemptions,json=vatExemptions,proto3" json:"vat_exemptions,omitempty"`
}

func (x *VatExemptions) Reset() {
	*x = VatExemptions{}
	mi := &file_accounting_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VatExemptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VatExemptions) ProtoMessage() {}

func (x *VatExemptions) ProtoReflect() protoreflect.Message {
	mi := &file_accounting_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VatExemptions.ProtoReflect.Descriptor instead.
func (*VatExemptions) Descriptor() ([]byte, []int) {
	return file_accounting_proto_rawDescGZIP(), []int{3}
}

func (x *VatExemptions) GetVatExemptions() []*VatExemption {
	if x != nil {
		return x.VatExemptions
	}
	return nil
}

var File_accounting_proto protoreflect.FileDescriptor

var file_accounting_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x07, 0x56, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x72, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x22, 0x46, 0x0a, 0x08, 0x56, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x3a,
	0x0a, 0x09, 0x76, 0x61, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x08, 0x76, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x0c, 0x56, 0x61,
	0x74, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x5a, 0x0a, 0x0d, 0x56, 0x61, 0x74, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x49, 0x0a, 0x0e, 0x76, 0x61, 0x74, 0x5f, 0x65, 0x78, 0x65, 0x6d, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x68, 0x65, 0x73, 0x74,
	0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x61, 0x74, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x76,
	0x61, 0x74, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x9d, 0x01, 0x0a,
	0x03, 0x54, 0x61, 0x78, 0x12, 0x45, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x56, 0x61, 0x74, 0x52, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x56, 0x61, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x56, 0x61, 0x74, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56,
	0x61, 0x74, 0x45, 0x78, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x1a, 0x5a, 0x18,
	0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounting_proto_rawDescOnce sync.Once
	file_accounting_proto_rawDescData = file_accounting_proto_rawDesc
)

func file_accounting_proto_rawDescGZIP() []byte {
	file_accounting_proto_rawDescOnce.Do(func() {
		file_accounting_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounting_proto_rawDescData)
	})
	return file_accounting_proto_rawDescData
}

var file_accounting_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_accounting_proto_goTypes = []any{
	(*VatRate)(nil),       // 0: hestia.v1.accounting.VatRate
	(*VatRates)(nil),      // 1: hestia.v1.accounting.VatRates
	(*VatExemption)(nil),  // 2: hestia.v1.accounting.VatExemption
	(*VatExemptions)(nil), // 3: hestia.v1.accounting.VatExemptions
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_accounting_proto_depIdxs = []int32{
	0, // 0: hestia.v1.accounting.VatRates.vat_rates:type_name -> hestia.v1.accounting.VatRate
	2, // 1: hestia.v1.accounting.VatExemptions.vat_exemptions:type_name -> hestia.v1.accounting.VatExemption
	4, // 2: hestia.v1.accounting.Tax.GetVatRates:input_type -> google.protobuf.Empty
	4, // 3: hestia.v1.accounting.Tax.GetVatExemptions:input_type -> google.protobuf.Empty
	1, // 4: hestia.v1.accounting.Tax.GetVatRates:output_type -> hestia.v1.accounting.VatRates
	3, // 5: hestia.v1.accounting.Tax.GetVatExemptions:output_type -> hestia.v1.accounting.VatExemptions
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_accounting_proto_init() }
func file_accounting_proto_init() {
	if File_accounting_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accounting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounting_proto_goTypes,
		DependencyIndexes: file_accounting_proto_depIdxs,
		MessageInfos:      file_accounting_proto_msgTypes,
	}.Build()
	File_accounting_proto = out.File
	file_accounting_proto_rawDesc = nil
	file_accounting_proto_goTypes = nil
	file_accounting_proto_depIdxs = nil
}