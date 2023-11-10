// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sprinkles/v1/api.proto

package sprinklesv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Api_UpsertHellos_FullMethodName            = "/sprinkles.v1.Api/UpsertHellos"
	Api_DeleteHellos_FullMethodName            = "/sprinkles.v1.Api/DeleteHellos"
	Api_ListHellos_FullMethodName              = "/sprinkles.v1.Api/ListHellos"
	Api_GetHellos_FullMethodName               = "/sprinkles.v1.Api/GetHellos"
	Api_GetOptionDefinitions_FullMethodName    = "/sprinkles.v1.Api/GetOptionDefinitions"
	Api_ListOptionDefinitions_FullMethodName   = "/sprinkles.v1.Api/ListOptionDefinitions"
	Api_DeleteOptionDefinitions_FullMethodName = "/sprinkles.v1.Api/DeleteOptionDefinitions"
	Api_UpsertOptionDefinitions_FullMethodName = "/sprinkles.v1.Api/UpsertOptionDefinitions"
	Api_GetOptionOverrides_FullMethodName      = "/sprinkles.v1.Api/GetOptionOverrides"
	Api_ListOptionOverrides_FullMethodName     = "/sprinkles.v1.Api/ListOptionOverrides"
	Api_DeleteOptionOverrides_FullMethodName   = "/sprinkles.v1.Api/DeleteOptionOverrides"
	Api_UpsertOptionOverrides_FullMethodName   = "/sprinkles.v1.Api/UpsertOptionOverrides"
	Api_GetOptionsByGroup_FullMethodName       = "/sprinkles.v1.Api/GetOptionsByGroup"
	Api_GetOptionValue_FullMethodName          = "/sprinkles.v1.Api/GetOptionValue"
	Api_ListGroups_FullMethodName              = "/sprinkles.v1.Api/ListGroups"
	Api_DeleteGroups_FullMethodName            = "/sprinkles.v1.Api/DeleteGroups"
	Api_UpsertGroups_FullMethodName            = "/sprinkles.v1.Api/UpsertGroups"
	Api_Healthy_FullMethodName                 = "/sprinkles.v1.Api/Healthy"
	Api_Ready_FullMethodName                   = "/sprinkles.v1.Api/Ready"
)

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	UpsertHellos(ctx context.Context, in *UpsertHellosRequest, opts ...grpc.CallOption) (*Hellos, error)
	DeleteHellos(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	ListHellos(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Hellos, error)
	GetHellos(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hellos, error)
	GetOptionDefinitions(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*OptionDefinitions, error)
	ListOptionDefinitions(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*OptionDefinitions, error)
	DeleteOptionDefinitions(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	UpsertOptionDefinitions(ctx context.Context, in *UpsertOptionDefinitionsRequest, opts ...grpc.CallOption) (*OptionDefinitions, error)
	GetOptionOverrides(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*OptionOverrides, error)
	ListOptionOverrides(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*OptionOverrides, error)
	DeleteOptionOverrides(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	UpsertOptionOverrides(ctx context.Context, in *UpsertOptionOverridesRequest, opts ...grpc.CallOption) (*OptionOverrides, error)
	GetOptionsByGroup(ctx context.Context, in *OptionGroupRequest, opts ...grpc.CallOption) (*OptionOverrides, error)
	GetOptionValue(ctx context.Context, in *OptionValueRequest, opts ...grpc.CallOption) (*OptionValueResponse, error)
	ListGroups(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Groups, error)
	DeleteGroups(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	UpsertGroups(ctx context.Context, in *UpsertGroupsRequest, opts ...grpc.CallOption) (*Groups, error)
	// Health check
	Healthy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Readiness check
	Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) UpsertHellos(ctx context.Context, in *UpsertHellosRequest, opts ...grpc.CallOption) (*Hellos, error) {
	out := new(Hellos)
	err := c.cc.Invoke(ctx, Api_UpsertHellos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteHellos(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Api_DeleteHellos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListHellos(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Hellos, error) {
	out := new(Hellos)
	err := c.cc.Invoke(ctx, Api_ListHellos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetHellos(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hellos, error) {
	out := new(Hellos)
	err := c.cc.Invoke(ctx, Api_GetHellos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetOptionDefinitions(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*OptionDefinitions, error) {
	out := new(OptionDefinitions)
	err := c.cc.Invoke(ctx, Api_GetOptionDefinitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListOptionDefinitions(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*OptionDefinitions, error) {
	out := new(OptionDefinitions)
	err := c.cc.Invoke(ctx, Api_ListOptionDefinitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteOptionDefinitions(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Api_DeleteOptionDefinitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpsertOptionDefinitions(ctx context.Context, in *UpsertOptionDefinitionsRequest, opts ...grpc.CallOption) (*OptionDefinitions, error) {
	out := new(OptionDefinitions)
	err := c.cc.Invoke(ctx, Api_UpsertOptionDefinitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetOptionOverrides(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*OptionOverrides, error) {
	out := new(OptionOverrides)
	err := c.cc.Invoke(ctx, Api_GetOptionOverrides_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListOptionOverrides(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*OptionOverrides, error) {
	out := new(OptionOverrides)
	err := c.cc.Invoke(ctx, Api_ListOptionOverrides_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteOptionOverrides(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Api_DeleteOptionOverrides_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpsertOptionOverrides(ctx context.Context, in *UpsertOptionOverridesRequest, opts ...grpc.CallOption) (*OptionOverrides, error) {
	out := new(OptionOverrides)
	err := c.cc.Invoke(ctx, Api_UpsertOptionOverrides_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetOptionsByGroup(ctx context.Context, in *OptionGroupRequest, opts ...grpc.CallOption) (*OptionOverrides, error) {
	out := new(OptionOverrides)
	err := c.cc.Invoke(ctx, Api_GetOptionsByGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetOptionValue(ctx context.Context, in *OptionValueRequest, opts ...grpc.CallOption) (*OptionValueResponse, error) {
	out := new(OptionValueResponse)
	err := c.cc.Invoke(ctx, Api_GetOptionValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListGroups(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Groups, error) {
	out := new(Groups)
	err := c.cc.Invoke(ctx, Api_ListGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteGroups(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Api_DeleteGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpsertGroups(ctx context.Context, in *UpsertGroupsRequest, opts ...grpc.CallOption) (*Groups, error) {
	out := new(Groups)
	err := c.cc.Invoke(ctx, Api_UpsertGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Healthy(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Api_Healthy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Api_Ready_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations should embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	UpsertHellos(context.Context, *UpsertHellosRequest) (*Hellos, error)
	DeleteHellos(context.Context, *DeleteRequest) (*DeleteResponse, error)
	ListHellos(context.Context, *ListRequest) (*Hellos, error)
	GetHellos(context.Context, *GetRequest) (*Hellos, error)
	GetOptionDefinitions(context.Context, *GetRequest) (*OptionDefinitions, error)
	ListOptionDefinitions(context.Context, *ListRequest) (*OptionDefinitions, error)
	DeleteOptionDefinitions(context.Context, *DeleteRequest) (*DeleteResponse, error)
	UpsertOptionDefinitions(context.Context, *UpsertOptionDefinitionsRequest) (*OptionDefinitions, error)
	GetOptionOverrides(context.Context, *GetRequest) (*OptionOverrides, error)
	ListOptionOverrides(context.Context, *ListRequest) (*OptionOverrides, error)
	DeleteOptionOverrides(context.Context, *DeleteRequest) (*DeleteResponse, error)
	UpsertOptionOverrides(context.Context, *UpsertOptionOverridesRequest) (*OptionOverrides, error)
	GetOptionsByGroup(context.Context, *OptionGroupRequest) (*OptionOverrides, error)
	GetOptionValue(context.Context, *OptionValueRequest) (*OptionValueResponse, error)
	ListGroups(context.Context, *ListRequest) (*Groups, error)
	DeleteGroups(context.Context, *DeleteGroupRequest) (*DeleteResponse, error)
	UpsertGroups(context.Context, *UpsertGroupsRequest) (*Groups, error)
	// Health check
	Healthy(context.Context, *Empty) (*Empty, error)
	// Readiness check
	Ready(context.Context, *Empty) (*Empty, error)
}

// UnimplementedApiServer should be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) UpsertHellos(context.Context, *UpsertHellosRequest) (*Hellos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertHellos not implemented")
}
func (UnimplementedApiServer) DeleteHellos(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHellos not implemented")
}
func (UnimplementedApiServer) ListHellos(context.Context, *ListRequest) (*Hellos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHellos not implemented")
}
func (UnimplementedApiServer) GetHellos(context.Context, *GetRequest) (*Hellos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHellos not implemented")
}
func (UnimplementedApiServer) GetOptionDefinitions(context.Context, *GetRequest) (*OptionDefinitions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptionDefinitions not implemented")
}
func (UnimplementedApiServer) ListOptionDefinitions(context.Context, *ListRequest) (*OptionDefinitions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOptionDefinitions not implemented")
}
func (UnimplementedApiServer) DeleteOptionDefinitions(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOptionDefinitions not implemented")
}
func (UnimplementedApiServer) UpsertOptionDefinitions(context.Context, *UpsertOptionDefinitionsRequest) (*OptionDefinitions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertOptionDefinitions not implemented")
}
func (UnimplementedApiServer) GetOptionOverrides(context.Context, *GetRequest) (*OptionOverrides, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptionOverrides not implemented")
}
func (UnimplementedApiServer) ListOptionOverrides(context.Context, *ListRequest) (*OptionOverrides, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOptionOverrides not implemented")
}
func (UnimplementedApiServer) DeleteOptionOverrides(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOptionOverrides not implemented")
}
func (UnimplementedApiServer) UpsertOptionOverrides(context.Context, *UpsertOptionOverridesRequest) (*OptionOverrides, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertOptionOverrides not implemented")
}
func (UnimplementedApiServer) GetOptionsByGroup(context.Context, *OptionGroupRequest) (*OptionOverrides, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptionsByGroup not implemented")
}
func (UnimplementedApiServer) GetOptionValue(context.Context, *OptionValueRequest) (*OptionValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOptionValue not implemented")
}
func (UnimplementedApiServer) ListGroups(context.Context, *ListRequest) (*Groups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedApiServer) DeleteGroups(context.Context, *DeleteGroupRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroups not implemented")
}
func (UnimplementedApiServer) UpsertGroups(context.Context, *UpsertGroupsRequest) (*Groups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertGroups not implemented")
}
func (UnimplementedApiServer) Healthy(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthy not implemented")
}
func (UnimplementedApiServer) Ready(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_UpsertHellos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertHellosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpsertHellos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_UpsertHellos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpsertHellos(ctx, req.(*UpsertHellosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteHellos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteHellos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_DeleteHellos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteHellos(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListHellos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListHellos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ListHellos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListHellos(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetHellos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetHellos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetHellos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetHellos(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetOptionDefinitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetOptionDefinitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetOptionDefinitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetOptionDefinitions(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListOptionDefinitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListOptionDefinitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ListOptionDefinitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListOptionDefinitions(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteOptionDefinitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteOptionDefinitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_DeleteOptionDefinitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteOptionDefinitions(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpsertOptionDefinitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertOptionDefinitionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpsertOptionDefinitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_UpsertOptionDefinitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpsertOptionDefinitions(ctx, req.(*UpsertOptionDefinitionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetOptionOverrides_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetOptionOverrides(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetOptionOverrides_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetOptionOverrides(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListOptionOverrides_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListOptionOverrides(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ListOptionOverrides_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListOptionOverrides(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteOptionOverrides_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteOptionOverrides(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_DeleteOptionOverrides_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteOptionOverrides(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpsertOptionOverrides_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertOptionOverridesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpsertOptionOverrides(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_UpsertOptionOverrides_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpsertOptionOverrides(ctx, req.(*UpsertOptionOverridesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetOptionsByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OptionGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetOptionsByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetOptionsByGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetOptionsByGroup(ctx, req.(*OptionGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetOptionValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OptionValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetOptionValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_GetOptionValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetOptionValue(ctx, req.(*OptionValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_ListGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ListGroups(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_DeleteGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteGroups(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpsertGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpsertGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_UpsertGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpsertGroups(ctx, req.(*UpsertGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Healthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Healthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_Healthy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Healthy(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_Ready_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Ready(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sprinkles.v1.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertHellos",
			Handler:    _Api_UpsertHellos_Handler,
		},
		{
			MethodName: "DeleteHellos",
			Handler:    _Api_DeleteHellos_Handler,
		},
		{
			MethodName: "ListHellos",
			Handler:    _Api_ListHellos_Handler,
		},
		{
			MethodName: "GetHellos",
			Handler:    _Api_GetHellos_Handler,
		},
		{
			MethodName: "GetOptionDefinitions",
			Handler:    _Api_GetOptionDefinitions_Handler,
		},
		{
			MethodName: "ListOptionDefinitions",
			Handler:    _Api_ListOptionDefinitions_Handler,
		},
		{
			MethodName: "DeleteOptionDefinitions",
			Handler:    _Api_DeleteOptionDefinitions_Handler,
		},
		{
			MethodName: "UpsertOptionDefinitions",
			Handler:    _Api_UpsertOptionDefinitions_Handler,
		},
		{
			MethodName: "GetOptionOverrides",
			Handler:    _Api_GetOptionOverrides_Handler,
		},
		{
			MethodName: "ListOptionOverrides",
			Handler:    _Api_ListOptionOverrides_Handler,
		},
		{
			MethodName: "DeleteOptionOverrides",
			Handler:    _Api_DeleteOptionOverrides_Handler,
		},
		{
			MethodName: "UpsertOptionOverrides",
			Handler:    _Api_UpsertOptionOverrides_Handler,
		},
		{
			MethodName: "GetOptionsByGroup",
			Handler:    _Api_GetOptionsByGroup_Handler,
		},
		{
			MethodName: "GetOptionValue",
			Handler:    _Api_GetOptionValue_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _Api_ListGroups_Handler,
		},
		{
			MethodName: "DeleteGroups",
			Handler:    _Api_DeleteGroups_Handler,
		},
		{
			MethodName: "UpsertGroups",
			Handler:    _Api_UpsertGroups_Handler,
		},
		{
			MethodName: "Healthy",
			Handler:    _Api_Healthy_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _Api_Ready_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sprinkles/v1/api.proto",
}
