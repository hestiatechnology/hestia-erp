// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/company_management.proto

package company

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CompanyManagement_CreateCompany_FullMethodName    = "/hestia.v1.com.CompanyManagement/CreateCompany"
	CompanyManagement_GetCompany_FullMethodName       = "/hestia.v1.com.CompanyManagement/GetCompany"
	CompanyManagement_UpdateCompany_FullMethodName    = "/hestia.v1.com.CompanyManagement/UpdateCompany"
	CompanyManagement_CreateLocation_FullMethodName   = "/hestia.v1.com.CompanyManagement/CreateLocation"
	CompanyManagement_AddUserToCompany_FullMethodName = "/hestia.v1.com.CompanyManagement/AddUserToCompany"
)

// CompanyManagementClient is the client API for CompanyManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyManagementClient interface {
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*IdResponse, error)
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error)
	// rpc DeleteCompany(DeleteCompanyRequest) returns (DeleteCompanyResponse);
	// rpc ListCompanies(ListCompaniesRequest) returns (ListCompaniesResponse);
	CreateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*IdResponse, error)
	AddUserToCompany(ctx context.Context, in *AddUserToCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type companyManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyManagementClient(cc grpc.ClientConnInterface) CompanyManagementClient {
	return &companyManagementClient{cc}
}

func (c *companyManagementClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, CompanyManagement_CreateCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyManagementClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyManagement_GetCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyManagementClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error) {
	out := new(UpdateCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyManagement_UpdateCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyManagementClient) CreateLocation(ctx context.Context, in *Location, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, CompanyManagement_CreateLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyManagementClient) AddUserToCompany(ctx context.Context, in *AddUserToCompanyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CompanyManagement_AddUserToCompany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyManagementServer is the server API for CompanyManagement service.
// All implementations must embed UnimplementedCompanyManagementServer
// for forward compatibility
type CompanyManagementServer interface {
	CreateCompany(context.Context, *CreateCompanyRequest) (*IdResponse, error)
	GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error)
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error)
	// rpc DeleteCompany(DeleteCompanyRequest) returns (DeleteCompanyResponse);
	// rpc ListCompanies(ListCompaniesRequest) returns (ListCompaniesResponse);
	CreateLocation(context.Context, *Location) (*IdResponse, error)
	AddUserToCompany(context.Context, *AddUserToCompanyRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCompanyManagementServer()
}

// UnimplementedCompanyManagementServer must be embedded to have forward compatible implementations.
type UnimplementedCompanyManagementServer struct {
}

func (UnimplementedCompanyManagementServer) CreateCompany(context.Context, *CreateCompanyRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedCompanyManagementServer) GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (UnimplementedCompanyManagementServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedCompanyManagementServer) CreateLocation(context.Context, *Location) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLocation not implemented")
}
func (UnimplementedCompanyManagementServer) AddUserToCompany(context.Context, *AddUserToCompanyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToCompany not implemented")
}
func (UnimplementedCompanyManagementServer) mustEmbedUnimplementedCompanyManagementServer() {}

// UnsafeCompanyManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyManagementServer will
// result in compilation errors.
type UnsafeCompanyManagementServer interface {
	mustEmbedUnimplementedCompanyManagementServer()
}

func RegisterCompanyManagementServer(s grpc.ServiceRegistrar, srv CompanyManagementServer) {
	s.RegisterService(&CompanyManagement_ServiceDesc, srv)
}

func _CompanyManagement_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyManagementServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyManagement_CreateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyManagementServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyManagement_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyManagementServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyManagement_GetCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyManagementServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyManagement_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyManagementServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyManagement_UpdateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyManagementServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyManagement_CreateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Location)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyManagementServer).CreateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyManagement_CreateLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyManagementServer).CreateLocation(ctx, req.(*Location))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyManagement_AddUserToCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyManagementServer).AddUserToCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyManagement_AddUserToCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyManagementServer).AddUserToCompany(ctx, req.(*AddUserToCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyManagement_ServiceDesc is the grpc.ServiceDesc for CompanyManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hestia.v1.com.CompanyManagement",
	HandlerType: (*CompanyManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyManagement_CreateCompany_Handler,
		},
		{
			MethodName: "GetCompany",
			Handler:    _CompanyManagement_GetCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CompanyManagement_UpdateCompany_Handler,
		},
		{
			MethodName: "CreateLocation",
			Handler:    _CompanyManagement_CreateLocation_Handler,
		},
		{
			MethodName: "AddUserToCompany",
			Handler:    _CompanyManagement_AddUserToCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/company_management.proto",
}
