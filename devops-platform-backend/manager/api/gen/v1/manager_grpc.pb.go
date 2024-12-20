// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: manager.proto

package managerpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	BizLineService_CreateBizLine_FullMethodName = "/manager.v1.BizLineService/CreateBizLine"
	BizLineService_DeleteBizLine_FullMethodName = "/manager.v1.BizLineService/DeleteBizLine"
	BizLineService_UpdateBizLine_FullMethodName = "/manager.v1.BizLineService/UpdateBizLine"
	BizLineService_GetBizLine_FullMethodName    = "/manager.v1.BizLineService/GetBizLine"
	BizLineService_GetBizLines_FullMethodName   = "/manager.v1.BizLineService/GetBizLines"
)

// BizLineServiceClient is the client API for BizLineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BizLineServiceClient interface {
	CreateBizLine(ctx context.Context, in *BizLine, opts ...grpc.CallOption) (*CreateBizLineResponse, error)
	DeleteBizLine(ctx context.Context, in *DeleteBizLineRequest, opts ...grpc.CallOption) (*DeleteBizLineResponse, error)
	UpdateBizLine(ctx context.Context, in *BizLineEntity, opts ...grpc.CallOption) (*UpdateBizLineResponse, error)
	GetBizLine(ctx context.Context, in *GetBizLineRequest, opts ...grpc.CallOption) (*BizLineEntity, error)
	GetBizLines(ctx context.Context, in *GetBizLinesRequest, opts ...grpc.CallOption) (*GetBizLinesResponse, error)
}

type bizLineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBizLineServiceClient(cc grpc.ClientConnInterface) BizLineServiceClient {
	return &bizLineServiceClient{cc}
}

func (c *bizLineServiceClient) CreateBizLine(ctx context.Context, in *BizLine, opts ...grpc.CallOption) (*CreateBizLineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBizLineResponse)
	err := c.cc.Invoke(ctx, BizLineService_CreateBizLine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizLineServiceClient) DeleteBizLine(ctx context.Context, in *DeleteBizLineRequest, opts ...grpc.CallOption) (*DeleteBizLineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBizLineResponse)
	err := c.cc.Invoke(ctx, BizLineService_DeleteBizLine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizLineServiceClient) UpdateBizLine(ctx context.Context, in *BizLineEntity, opts ...grpc.CallOption) (*UpdateBizLineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBizLineResponse)
	err := c.cc.Invoke(ctx, BizLineService_UpdateBizLine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizLineServiceClient) GetBizLine(ctx context.Context, in *GetBizLineRequest, opts ...grpc.CallOption) (*BizLineEntity, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BizLineEntity)
	err := c.cc.Invoke(ctx, BizLineService_GetBizLine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizLineServiceClient) GetBizLines(ctx context.Context, in *GetBizLinesRequest, opts ...grpc.CallOption) (*GetBizLinesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBizLinesResponse)
	err := c.cc.Invoke(ctx, BizLineService_GetBizLines_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BizLineServiceServer is the server API for BizLineService service.
// All implementations must embed UnimplementedBizLineServiceServer
// for forward compatibility
type BizLineServiceServer interface {
	CreateBizLine(context.Context, *BizLine) (*CreateBizLineResponse, error)
	DeleteBizLine(context.Context, *DeleteBizLineRequest) (*DeleteBizLineResponse, error)
	UpdateBizLine(context.Context, *BizLineEntity) (*UpdateBizLineResponse, error)
	GetBizLine(context.Context, *GetBizLineRequest) (*BizLineEntity, error)
	GetBizLines(context.Context, *GetBizLinesRequest) (*GetBizLinesResponse, error)
	mustEmbedUnimplementedBizLineServiceServer()
}

// UnimplementedBizLineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBizLineServiceServer struct {
}

func (UnimplementedBizLineServiceServer) CreateBizLine(context.Context, *BizLine) (*CreateBizLineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBizLine not implemented")
}
func (UnimplementedBizLineServiceServer) DeleteBizLine(context.Context, *DeleteBizLineRequest) (*DeleteBizLineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBizLine not implemented")
}
func (UnimplementedBizLineServiceServer) UpdateBizLine(context.Context, *BizLineEntity) (*UpdateBizLineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBizLine not implemented")
}
func (UnimplementedBizLineServiceServer) GetBizLine(context.Context, *GetBizLineRequest) (*BizLineEntity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBizLine not implemented")
}
func (UnimplementedBizLineServiceServer) GetBizLines(context.Context, *GetBizLinesRequest) (*GetBizLinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBizLines not implemented")
}
func (UnimplementedBizLineServiceServer) mustEmbedUnimplementedBizLineServiceServer() {}

// UnsafeBizLineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BizLineServiceServer will
// result in compilation errors.
type UnsafeBizLineServiceServer interface {
	mustEmbedUnimplementedBizLineServiceServer()
}

func RegisterBizLineServiceServer(s grpc.ServiceRegistrar, srv BizLineServiceServer) {
	s.RegisterService(&BizLineService_ServiceDesc, srv)
}

func _BizLineService_CreateBizLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BizLine)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizLineServiceServer).CreateBizLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BizLineService_CreateBizLine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizLineServiceServer).CreateBizLine(ctx, req.(*BizLine))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizLineService_DeleteBizLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBizLineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizLineServiceServer).DeleteBizLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BizLineService_DeleteBizLine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizLineServiceServer).DeleteBizLine(ctx, req.(*DeleteBizLineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizLineService_UpdateBizLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BizLineEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizLineServiceServer).UpdateBizLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BizLineService_UpdateBizLine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizLineServiceServer).UpdateBizLine(ctx, req.(*BizLineEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizLineService_GetBizLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBizLineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizLineServiceServer).GetBizLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BizLineService_GetBizLine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizLineServiceServer).GetBizLine(ctx, req.(*GetBizLineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizLineService_GetBizLines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBizLinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizLineServiceServer).GetBizLines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BizLineService_GetBizLines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizLineServiceServer).GetBizLines(ctx, req.(*GetBizLinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BizLineService_ServiceDesc is the grpc.ServiceDesc for BizLineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BizLineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.v1.BizLineService",
	HandlerType: (*BizLineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBizLine",
			Handler:    _BizLineService_CreateBizLine_Handler,
		},
		{
			MethodName: "DeleteBizLine",
			Handler:    _BizLineService_DeleteBizLine_Handler,
		},
		{
			MethodName: "UpdateBizLine",
			Handler:    _BizLineService_UpdateBizLine_Handler,
		},
		{
			MethodName: "GetBizLine",
			Handler:    _BizLineService_GetBizLine_Handler,
		},
		{
			MethodName: "GetBizLines",
			Handler:    _BizLineService_GetBizLines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}

const (
	ProjectService_CreateProject_FullMethodName = "/manager.v1.ProjectService/CreateProject"
	ProjectService_DeleteProject_FullMethodName = "/manager.v1.ProjectService/DeleteProject"
	ProjectService_UpdateProject_FullMethodName = "/manager.v1.ProjectService/UpdateProject"
	ProjectService_GetProject_FullMethodName    = "/manager.v1.ProjectService/GetProject"
	ProjectService_GetProjects_FullMethodName   = "/manager.v1.ProjectService/GetProjects"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error)
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error)
	GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_CreateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_DeleteProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_UpdateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetProjects(ctx context.Context, in *GetProjectsRequest, opts ...grpc.CallOption) (*GetProjectsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProjectsResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetProjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error)
	UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error)
	GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectServiceServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectServiceServer) UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedProjectServiceServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectServiceServer) GetProjects(context.Context, *GetProjectsRequest) (*GetProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_UpdateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).UpdateProject(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetProjects(ctx, req.(*GetProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.v1.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectService_CreateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _ProjectService_DeleteProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _ProjectService_UpdateProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _ProjectService_GetProject_Handler,
		},
		{
			MethodName: "GetProjects",
			Handler:    _ProjectService_GetProjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}

const (
	AppService_CreateProject_FullMethodName = "/manager.v1.AppService/CreateProject"
	AppService_DeleteProject_FullMethodName = "/manager.v1.AppService/DeleteProject"
	AppService_UpdateProject_FullMethodName = "/manager.v1.AppService/UpdateProject"
	AppService_GetProject_FullMethodName    = "/manager.v1.AppService/GetProject"
	AppService_GetProjects_FullMethodName   = "/manager.v1.AppService/GetProjects"
)

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	CreateProject(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error)
	DeleteProject(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*DeleteAppResponse, error)
	UpdateProject(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error)
	GetProject(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error)
	GetProjects(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) CreateProject(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAppResponse)
	err := c.cc.Invoke(ctx, AppService_CreateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) DeleteProject(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*DeleteAppResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAppResponse)
	err := c.cc.Invoke(ctx, AppService_DeleteProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) UpdateProject(ctx context.Context, in *UpdateAppRequest, opts ...grpc.CallOption) (*UpdateAppResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAppResponse)
	err := c.cc.Invoke(ctx, AppService_UpdateProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetProject(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*GetAppResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppResponse)
	err := c.cc.Invoke(ctx, AppService_GetProject_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) GetProjects(ctx context.Context, in *GetAppsRequest, opts ...grpc.CallOption) (*GetAppsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAppsResponse)
	err := c.cc.Invoke(ctx, AppService_GetProjects_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility
type AppServiceServer interface {
	CreateProject(context.Context, *CreateAppRequest) (*CreateAppResponse, error)
	DeleteProject(context.Context, *DeleteAppRequest) (*DeleteAppResponse, error)
	UpdateProject(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error)
	GetProject(context.Context, *GetAppRequest) (*GetAppResponse, error)
	GetProjects(context.Context, *GetAppsRequest) (*GetAppsResponse, error)
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppServiceServer struct {
}

func (UnimplementedAppServiceServer) CreateProject(context.Context, *CreateAppRequest) (*CreateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedAppServiceServer) DeleteProject(context.Context, *DeleteAppRequest) (*DeleteAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedAppServiceServer) UpdateProject(context.Context, *UpdateAppRequest) (*UpdateAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedAppServiceServer) GetProject(context.Context, *GetAppRequest) (*GetAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedAppServiceServer) GetProjects(context.Context, *GetAppsRequest) (*GetAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjects not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).CreateProject(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).DeleteProject(ctx, req.(*DeleteAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_UpdateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).UpdateProject(ctx, req.(*UpdateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_GetProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetProject(ctx, req.(*GetAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_GetProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).GetProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_GetProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).GetProjects(ctx, req.(*GetAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.v1.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _AppService_CreateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _AppService_DeleteProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _AppService_UpdateProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _AppService_GetProject_Handler,
		},
		{
			MethodName: "GetProjects",
			Handler:    _AppService_GetProjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}
