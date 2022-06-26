package provider

import (
	"context"
	"iman-task/connection_grpc"
	"iman-task/post-editor/internal/models"
	"iman-task/post-editor/internal/service"
	"log"
)

type GRPCServer struct {
	getter service.ServiceGetter
	editor service.ServiceEditor
	connection_grpc.UnimplementedEditorServer
}

func NewProvider(editor service.ServiceEditor, getter service.ServiceGetter) connection_grpc.EditorServer {
	return &GRPCServer{
		getter: getter,
		editor: editor,
	}
}

func (g *GRPCServer) GetPosts(ctx context.Context, req *connection_grpc.GetPostsRequest) (*connection_grpc.GetPostsResponse, error) {
	posts, err := g.getter.GetSome(req.GetPostsId())
	if err != nil {
		log.Printf("error in get some posts in post editor provider:--->%v\n", err)
		return posts, err
	}
	return posts, nil
}

func (g *GRPCServer) GetPostById(ctx context.Context, req *connection_grpc.GetPostByIdRequest) (*connection_grpc.GetPostByIdResponse, error) {
	post, err := g.getter.GetOneById(req.GetId())
	if err != nil {
		log.Printf("error in get one post in post editor provider:--->%v\n", err)
		return post, err
	}
	return post, nil
}

func (g *GRPCServer) DeletePost(ctx context.Context, req *connection_grpc.DeletePostRequest) (*connection_grpc.DeletePostResponse, error) {
	if err := g.editor.Delete(req.GetId()); err != nil {
		log.Printf("error in delete post in post editor provider:--->%v\n", err)
		return &connection_grpc.DeletePostResponse{Status: "invalid id"}, err
	}
	return &connection_grpc.DeletePostResponse{Status: "success"}, nil
}
func (g *GRPCServer) UpdatePost(ctx context.Context, req *connection_grpc.UpdatePostRequest) (*connection_grpc.UpdatePostResponse, error) {
	post := &models.Post{
		Id:     int(req.GetId()),
		UserID: int(req.GetUserId()),
		Title:  req.GetTitle(),
		Body:   req.GetBody(),
	}
	if err := g.editor.Update(post); err != nil {
		log.Printf("error in update post in post editor provider:--->%v\n", err)
		return &connection_grpc.UpdatePostResponse{Status: "failed"}, err
	}
	return &connection_grpc.UpdatePostResponse{Status: "success post updated"}, nil
}
