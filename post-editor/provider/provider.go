package provider

import (
	"context"
	"iman-task/connection_grpc"
	"iman-task/post-editor/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	getter service.ServiceGetter
	editor service.ServiceEditor
}

func NewProvider(editor service.ServiceEditor, getter service.ServiceGetter) connection_grpc.EditorServer {
	return &GRPCServer{
		getter: getter,
		editor: editor,
	}
}

func (g *GRPCServer) GetPosts(ctx context.Context, req *connection_grpc.GetPostsRequest) (*connection_grpc.GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (g *GRPCServer) GetPostById(ctx context.Context, req *connection_grpc.GetPostByIdRequest) (*connection_grpc.GetPostByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostById not implemented")
}
func (g *GRPCServer) DeletePost(ctx context.Context, req *connection_grpc.DeletePostRequest) (*connection_grpc.DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (g *GRPCServer) UpdatePost(ctx context.Context, req *connection_grpc.UpdatePostRequest) (*connection_grpc.UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
