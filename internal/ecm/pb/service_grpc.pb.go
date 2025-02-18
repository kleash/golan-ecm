// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// ContentEngineClient is the client API for ContentEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentEngineClient interface {
	// CreateWorkspace creates a new workspace.
	CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error)
	GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error)
	CreateDocumentClass(ctx context.Context, in *CreateDocumentClassRequest, opts ...grpc.CallOption) (*CreateDocumentClassResponse, error)
}

type contentEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewContentEngineClient(cc grpc.ClientConnInterface) ContentEngineClient {
	return &contentEngineClient{cc}
}

func (c *contentEngineClient) CreateWorkspace(ctx context.Context, in *CreateWorkspaceRequest, opts ...grpc.CallOption) (*CreateWorkspaceResponse, error) {
	out := new(CreateWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/ecm.ContentEngine/CreateWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentEngineClient) GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error) {
	out := new(GetWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/ecm.ContentEngine/GetWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentEngineClient) CreateDocumentClass(ctx context.Context, in *CreateDocumentClassRequest, opts ...grpc.CallOption) (*CreateDocumentClassResponse, error) {
	out := new(CreateDocumentClassResponse)
	err := c.cc.Invoke(ctx, "/ecm.ContentEngine/CreateDocumentClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentEngineServer is the server API for ContentEngine service.
// All implementations must embed UnimplementedContentEngineServer
// for forward compatibility
type ContentEngineServer interface {
	// CreateWorkspace creates a new workspace.
	CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error)
	GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error)
	CreateDocumentClass(context.Context, *CreateDocumentClassRequest) (*CreateDocumentClassResponse, error)
	mustEmbedUnimplementedContentEngineServer()
}

// UnimplementedContentEngineServer must be embedded to have forward compatible implementations.
type UnimplementedContentEngineServer struct {
}

func (UnimplementedContentEngineServer) CreateWorkspace(context.Context, *CreateWorkspaceRequest) (*CreateWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspace not implemented")
}
func (UnimplementedContentEngineServer) GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspace not implemented")
}
func (UnimplementedContentEngineServer) CreateDocumentClass(context.Context, *CreateDocumentClassRequest) (*CreateDocumentClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocumentClass not implemented")
}
func (UnimplementedContentEngineServer) mustEmbedUnimplementedContentEngineServer() {}

// UnsafeContentEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentEngineServer will
// result in compilation errors.
type UnsafeContentEngineServer interface {
	mustEmbedUnimplementedContentEngineServer()
}

func RegisterContentEngineServer(s grpc.ServiceRegistrar, srv ContentEngineServer) {
	s.RegisterService(&ContentEngine_ServiceDesc, srv)
}

func _ContentEngine_CreateWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentEngineServer).CreateWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecm.ContentEngine/CreateWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentEngineServer).CreateWorkspace(ctx, req.(*CreateWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentEngine_GetWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentEngineServer).GetWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecm.ContentEngine/GetWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentEngineServer).GetWorkspace(ctx, req.(*GetWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentEngine_CreateDocumentClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentEngineServer).CreateDocumentClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ecm.ContentEngine/CreateDocumentClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentEngineServer).CreateDocumentClass(ctx, req.(*CreateDocumentClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentEngine_ServiceDesc is the grpc.ServiceDesc for ContentEngine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentEngine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecm.ContentEngine",
	HandlerType: (*ContentEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkspace",
			Handler:    _ContentEngine_CreateWorkspace_Handler,
		},
		{
			MethodName: "GetWorkspace",
			Handler:    _ContentEngine_GetWorkspace_Handler,
		},
		{
			MethodName: "CreateDocumentClass",
			Handler:    _ContentEngine_CreateDocumentClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
