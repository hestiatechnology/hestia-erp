// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: proto/company_management.proto

package company

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

type CreateCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates whether the company is a sole trader or not
	IsSoleTrader   bool      `protobuf:"varint,2,opt,name=isSoleTrader,proto3" json:"isSoleTrader,omitempty"`
	CommercialName *string   `protobuf:"bytes,3,opt,name=commercialName,proto3,oneof" json:"commercialName,omitempty"`
	VatId          int32     `protobuf:"varint,4,opt,name=vatId,proto3" json:"vatId,omitempty"`
	Ssn            int32     `protobuf:"varint,5,opt,name=ssn,proto3" json:"ssn,omitempty"`
	Location       *Location `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CreateCompanyRequest) Reset() {
	*x = CreateCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanyRequest) ProtoMessage() {}

func (x *CreateCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompanyRequest.ProtoReflect.Descriptor instead.
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return file_proto_company_management_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCompanyRequest) GetIsSoleTrader() bool {
	if x != nil {
		return x.IsSoleTrader
	}
	return false
}

func (x *CreateCompanyRequest) GetCommercialName() string {
	if x != nil && x.CommercialName != nil {
		return *x.CommercialName
	}
	return ""
}

func (x *CreateCompanyRequest) GetVatId() int32 {
	if x != nil {
		return x.VatId
	}
	return 0
}

func (x *CreateCompanyRequest) GetSsn() int32 {
	if x != nil {
		return x.Ssn
	}
	return 0
}

func (x *CreateCompanyRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Locality   string `protobuf:"bytes,2,opt,name=locality,proto3" json:"locality,omitempty"`
	PostalCode string `protobuf:"bytes,3,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	// Must be a Alpha-2 code
	Country string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_proto_company_management_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Location) GetLocality() string {
	if x != nil {
		return x.Locality
	}
	return ""
}

func (x *Location) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Location) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_proto_company_management_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateCompanyRequest) Reset() {
	*x = UpdateCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCompanyRequest) ProtoMessage() {}

func (x *UpdateCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCompanyRequest.ProtoReflect.Descriptor instead.
func (*UpdateCompanyRequest) Descriptor() ([]byte, []int) {
	return file_proto_company_management_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCompanyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCompanyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateCompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateCompanyResponse) Reset() {
	*x = UpdateCompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_management_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCompanyResponse) ProtoMessage() {}

func (x *UpdateCompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_management_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCompanyResponse.ProtoReflect.Descriptor instead.
func (*UpdateCompanyResponse) Descriptor() ([]byte, []int) {
	return file_proto_company_management_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCompanyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListCompaniesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Companies []*Company `protobuf:"bytes,1,rep,name=companies,proto3" json:"companies,omitempty"`
}

func (x *ListCompaniesResponse) Reset() {
	*x = ListCompaniesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_management_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCompaniesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCompaniesResponse) ProtoMessage() {}

func (x *ListCompaniesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_management_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCompaniesResponse.ProtoReflect.Descriptor instead.
func (*ListCompaniesResponse) Descriptor() ([]byte, []int) {
	return file_proto_company_management_proto_rawDescGZIP(), []int{5}
}

func (x *ListCompaniesResponse) GetCompanies() []*Company {
	if x != nil {
		return x.Companies
	}
	return nil
}

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CommercialName string `protobuf:"bytes,3,opt,name=commercialName,proto3" json:"commercialName,omitempty"`
	Description    string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Logo           string `protobuf:"bytes,5,opt,name=logo,proto3" json:"logo,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_management_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_management_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_proto_company_management_proto_rawDescGZIP(), []int{6}
}

func (x *Company) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetCommercialName() string {
	if x != nil {
		return x.CommercialName
	}
	return ""
}

func (x *Company) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Company) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

type AddUserToCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Email of the user to add to the company
	// If the user does not exist, a temporary account will be created until the user logs in
	Email      string  `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	CompanyId  string  `protobuf:"bytes,2,opt,name=companyId,proto3" json:"companyId,omitempty"`
	EmployeeId *string `protobuf:"bytes,3,opt,name=employeeId,proto3,oneof" json:"employeeId,omitempty"`
}

func (x *AddUserToCompanyRequest) Reset() {
	*x = AddUserToCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_company_management_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToCompanyRequest) ProtoMessage() {}

func (x *AddUserToCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_company_management_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToCompanyRequest.ProtoReflect.Descriptor instead.
func (*AddUserToCompanyRequest) Descriptor() ([]byte, []int) {
	return file_proto_company_management_proto_rawDescGZIP(), []int{7}
}

func (x *AddUserToCompanyRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddUserToCompanyRequest) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *AddUserToCompanyRequest) GetEmployeeId() string {
	if x != nil && x.EmployeeId != nil {
		return *x.EmployeeId
	}
	return ""
}

var File_proto_company_management_proto protoreflect.FileDescriptor

var file_proto_company_management_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x53, 0x6f, 0x6c, 0x65, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x53,
	0x6f, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x74, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x73, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x73, 0x6e, 0x12, 0x2f,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x7a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x14,
	0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6c, 0x6f, 0x67, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x32, 0x9f, 0x03, 0x0a, 0x11, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3f,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12,
	0x1f, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x12,
	0x48, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x0d, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x48, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1f, 0x2e, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x22,
	0x2e, 0x68, 0x65, 0x73, 0x74, 0x69, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x17, 0x5a, 0x15, 0x68, 0x65,
	0x73, 0x74, 0x69, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_company_management_proto_rawDescOnce sync.Once
	file_proto_company_management_proto_rawDescData = file_proto_company_management_proto_rawDesc
)

func file_proto_company_management_proto_rawDescGZIP() []byte {
	file_proto_company_management_proto_rawDescOnce.Do(func() {
		file_proto_company_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_company_management_proto_rawDescData)
	})
	return file_proto_company_management_proto_rawDescData
}

var file_proto_company_management_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_company_management_proto_goTypes = []interface{}{
	(*CreateCompanyRequest)(nil),    // 0: hestia.v1.CreateCompanyRequest
	(*Location)(nil),                // 1: hestia.v1.Location
	(*Id)(nil),                      // 2: hestia.v1.Id
	(*UpdateCompanyRequest)(nil),    // 3: hestia.v1.UpdateCompanyRequest
	(*UpdateCompanyResponse)(nil),   // 4: hestia.v1.UpdateCompanyResponse
	(*ListCompaniesResponse)(nil),   // 5: hestia.v1.ListCompaniesResponse
	(*Company)(nil),                 // 6: hestia.v1.Company
	(*AddUserToCompanyRequest)(nil), // 7: hestia.v1.AddUserToCompanyRequest
	(*emptypb.Empty)(nil),           // 8: google.protobuf.Empty
}
var file_proto_company_management_proto_depIdxs = []int32{
	1, // 0: hestia.v1.CreateCompanyRequest.location:type_name -> hestia.v1.Location
	6, // 1: hestia.v1.ListCompaniesResponse.companies:type_name -> hestia.v1.Company
	0, // 2: hestia.v1.CompanyManagement.CreateCompany:input_type -> hestia.v1.CreateCompanyRequest
	8, // 3: hestia.v1.CompanyManagement.GetCompanies:input_type -> google.protobuf.Empty
	2, // 4: hestia.v1.CompanyManagement.GetCompany:input_type -> hestia.v1.Id
	3, // 5: hestia.v1.CompanyManagement.UpdateCompany:input_type -> hestia.v1.UpdateCompanyRequest
	1, // 6: hestia.v1.CompanyManagement.CreateLocation:input_type -> hestia.v1.Location
	7, // 7: hestia.v1.CompanyManagement.AddUserToCompany:input_type -> hestia.v1.AddUserToCompanyRequest
	2, // 8: hestia.v1.CompanyManagement.CreateCompany:output_type -> hestia.v1.Id
	5, // 9: hestia.v1.CompanyManagement.GetCompanies:output_type -> hestia.v1.ListCompaniesResponse
	6, // 10: hestia.v1.CompanyManagement.GetCompany:output_type -> hestia.v1.Company
	8, // 11: hestia.v1.CompanyManagement.UpdateCompany:output_type -> google.protobuf.Empty
	2, // 12: hestia.v1.CompanyManagement.CreateLocation:output_type -> hestia.v1.Id
	8, // 13: hestia.v1.CompanyManagement.AddUserToCompany:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_company_management_proto_init() }
func file_proto_company_management_proto_init() {
	if File_proto_company_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_company_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompanyRequest); i {
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
		file_proto_company_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_proto_company_management_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_proto_company_management_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCompanyRequest); i {
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
		file_proto_company_management_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCompanyResponse); i {
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
		file_proto_company_management_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCompaniesResponse); i {
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
		file_proto_company_management_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Company); i {
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
		file_proto_company_management_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToCompanyRequest); i {
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
	file_proto_company_management_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_company_management_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_company_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_company_management_proto_goTypes,
		DependencyIndexes: file_proto_company_management_proto_depIdxs,
		MessageInfos:      file_proto_company_management_proto_msgTypes,
	}.Build()
	File_proto_company_management_proto = out.File
	file_proto_company_management_proto_rawDesc = nil
	file_proto_company_management_proto_goTypes = nil
	file_proto_company_management_proto_depIdxs = nil
}
