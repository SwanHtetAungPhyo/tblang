
package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	Provider_GetSchema_FullMethodName              = "/plugin.Provider/GetSchema"
	Provider_Configure_FullMethodName              = "/plugin.Provider/Configure"
	Provider_PlanResourceChange_FullMethodName     = "/plugin.Provider/PlanResourceChange"
	Provider_ApplyResourceChange_FullMethodName    = "/plugin.Provider/ApplyResourceChange"
	Provider_ReadResource_FullMethodName           = "/plugin.Provider/ReadResource"
	Provider_ImportResource_FullMethodName         = "/plugin.Provider/ImportResource"
	Provider_ValidateResourceConfig_FullMethodName = "/plugin.Provider/ValidateResourceConfig"
)

type ProviderClient interface {
	GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error)
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
	PlanResourceChange(ctx context.Context, in *PlanResourceChangeRequest, opts ...grpc.CallOption) (*PlanResourceChangeResponse, error)
	ApplyResourceChange(ctx context.Context, in *ApplyResourceChangeRequest, opts ...grpc.CallOption) (*ApplyResourceChangeResponse, error)
	ReadResource(ctx context.Context, in *ReadResourceRequest, opts ...grpc.CallOption) (*ReadResourceResponse, error)
	ImportResource(ctx context.Context, in *ImportResourceRequest, opts ...grpc.CallOption) (*ImportResourceResponse, error)
	ValidateResourceConfig(ctx context.Context, in *ValidateResourceConfigRequest, opts ...grpc.CallOption) (*ValidateResourceConfigResponse, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSchemaResponse)
	err := c.cc.Invoke(ctx, Provider_GetSchema_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfigureResponse)
	err := c.cc.Invoke(ctx, Provider_Configure_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) PlanResourceChange(ctx context.Context, in *PlanResourceChangeRequest, opts ...grpc.CallOption) (*PlanResourceChangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlanResourceChangeResponse)
	err := c.cc.Invoke(ctx, Provider_PlanResourceChange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ApplyResourceChange(ctx context.Context, in *ApplyResourceChangeRequest, opts ...grpc.CallOption) (*ApplyResourceChangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApplyResourceChangeResponse)
	err := c.cc.Invoke(ctx, Provider_ApplyResourceChange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ReadResource(ctx context.Context, in *ReadResourceRequest, opts ...grpc.CallOption) (*ReadResourceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadResourceResponse)
	err := c.cc.Invoke(ctx, Provider_ReadResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ImportResource(ctx context.Context, in *ImportResourceRequest, opts ...grpc.CallOption) (*ImportResourceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImportResourceResponse)
	err := c.cc.Invoke(ctx, Provider_ImportResource_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ValidateResourceConfig(ctx context.Context, in *ValidateResourceConfigRequest, opts ...grpc.CallOption) (*ValidateResourceConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateResourceConfigResponse)
	err := c.cc.Invoke(ctx, Provider_ValidateResourceConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ProviderServer interface {
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
	PlanResourceChange(context.Context, *PlanResourceChangeRequest) (*PlanResourceChangeResponse, error)
	ApplyResourceChange(context.Context, *ApplyResourceChangeRequest) (*ApplyResourceChangeResponse, error)
	ReadResource(context.Context, *ReadResourceRequest) (*ReadResourceResponse, error)
	ImportResource(context.Context, *ImportResourceRequest) (*ImportResourceResponse, error)
	ValidateResourceConfig(context.Context, *ValidateResourceConfigRequest) (*ValidateResourceConfigResponse, error)
	mustEmbedUnimplementedProviderServer()
}

type UnimplementedProviderServer struct{}

func (UnimplementedProviderServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedProviderServer) Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedProviderServer) PlanResourceChange(context.Context, *PlanResourceChangeRequest) (*PlanResourceChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlanResourceChange not implemented")
}
func (UnimplementedProviderServer) ApplyResourceChange(context.Context, *ApplyResourceChangeRequest) (*ApplyResourceChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyResourceChange not implemented")
}
func (UnimplementedProviderServer) ReadResource(context.Context, *ReadResourceRequest) (*ReadResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadResource not implemented")
}
func (UnimplementedProviderServer) ImportResource(context.Context, *ImportResourceRequest) (*ImportResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportResource not implemented")
}
func (UnimplementedProviderServer) ValidateResourceConfig(context.Context, *ValidateResourceConfigRequest) (*ValidateResourceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateResourceConfig not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}
func (UnimplementedProviderServer) testEmbeddedByValue()                  {}

type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_GetSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).GetSchema(ctx, req.(*GetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_Configure_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_PlanResourceChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanResourceChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).PlanResourceChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_PlanResourceChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).PlanResourceChange(ctx, req.(*PlanResourceChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ApplyResourceChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyResourceChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ApplyResourceChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_ApplyResourceChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ApplyResourceChange(ctx, req.(*ApplyResourceChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ReadResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ReadResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_ReadResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ReadResource(ctx, req.(*ReadResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ImportResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ImportResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_ImportResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ImportResource(ctx, req.(*ImportResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ValidateResourceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateResourceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ValidateResourceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Provider_ValidateResourceConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ValidateResourceConfig(ctx, req.(*ValidateResourceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugin.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchema",
			Handler:    _Provider_GetSchema_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Provider_Configure_Handler,
		},
		{
			MethodName: "PlanResourceChange",
			Handler:    _Provider_PlanResourceChange_Handler,
		},
		{
			MethodName: "ApplyResourceChange",
			Handler:    _Provider_ApplyResourceChange_Handler,
		},
		{
			MethodName: "ReadResource",
			Handler:    _Provider_ReadResource_Handler,
		},
		{
			MethodName: "ImportResource",
			Handler:    _Provider_ImportResource_Handler,
		},
		{
			MethodName: "ValidateResourceConfig",
			Handler:    _Provider_ValidateResourceConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/plugin/proto/plugin.proto",
}
