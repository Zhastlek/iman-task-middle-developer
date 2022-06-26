package provider

import (
	"context"
	"iman-task/connection_grpc"
	"iman-task/post-parser/internal/service"
)

type GRPCServer struct {
	creater service.CreatePostsService
}

func NewCreaterProvider(creater service.CreatePostsService) connection_grpc.CreatorServer {
	return &GRPCServer{
		creater: creater,
	}
}

func (g *GRPCServer) Parse(ctx context.Context, req *connection_grpc.Request) (*connection_grpc.Response, error) {
	// reqString := req.GetUrl()

	return &connection_grpc.Response{
		Status: "OK",
	}, nil
}
