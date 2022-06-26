// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: editor.proto

package connection_grpc

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

// EditorClient is the client API for Editor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EditorClient interface {
	GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error)
	GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error)
}

type editorClient struct {
	cc grpc.ClientConnInterface
}

func NewEditorClient(cc grpc.ClientConnInterface) EditorClient {
	return &editorClient{cc}
}

func (c *editorClient) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	out := new(GetPostsResponse)
	err := c.cc.Invoke(ctx, "/editer.Editor/GetPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *editorClient) GetPostById(ctx context.Context, in *GetPostByIdRequest, opts ...grpc.CallOption) (*GetPostByIdResponse, error) {
	out := new(GetPostByIdResponse)
	err := c.cc.Invoke(ctx, "/editer.Editor/GetPostById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *editorClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, "/editer.Editor/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *editorClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error) {
	out := new(UpdatePostResponse)
	err := c.cc.Invoke(ctx, "/editer.Editor/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EditorServer is the server API for Editor service.
// All implementations must embed UnimplementedEditorServer
// for forward compatibility
type EditorServer interface {
	GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error)
	GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error)
	mustEmbedUnimplementedEditorServer()
}

// UnimplementedEditorServer must be embedded to have forward compatible implementations.
type UnimplementedEditorServer struct {
}

func (UnimplementedEditorServer) GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedEditorServer) GetPostById(context.Context, *GetPostByIdRequest) (*GetPostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (UnimplementedEditorServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedEditorServer) UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedEditorServer) mustEmbedUnimplementedEditorServer() {}

// UnsafeEditorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EditorServer will
// result in compilation errors.
type UnsafeEditorServer interface {
	mustEmbedUnimplementedEditorServer()
}

func RegisterEditorServer(s grpc.ServiceRegistrar, srv EditorServer) {
	s.RegisterService(&Editor_ServiceDesc, srv)
}

func _Editor_GetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EditorServer).GetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/editer.Editor/GetPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EditorServer).GetPosts(ctx, req.(*GetPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Editor_GetPostById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EditorServer).GetPostById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/editer.Editor/GetPostById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EditorServer).GetPostById(ctx, req.(*GetPostByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Editor_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EditorServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/editer.Editor/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EditorServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Editor_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EditorServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/editer.Editor/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EditorServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Editor_ServiceDesc is the grpc.ServiceDesc for Editor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Editor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "editer.Editor",
	HandlerType: (*EditorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPosts",
			Handler:    _Editor_GetPosts_Handler,
		},
		{
			MethodName: "GetPostById",
			Handler:    _Editor_GetPostById_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Editor_DeletePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Editor_UpdatePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "editor.proto",
}
