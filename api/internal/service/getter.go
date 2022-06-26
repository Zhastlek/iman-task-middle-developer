package service

import (
	"context"
	"iman-task/api/internal/models"
	"iman-task/connection_grpc"
	"log"
)

type getPostService struct {
	getter connection_grpc.EditorClient
}

func NewGetterService(getter connection_grpc.EditorClient) ServiceGetter {
	return &getPostService{
		getter: getter,
	}
}

func (s *getPostService) GetOne(post *models.Post) (*models.Post, error) {
	log.Println("id post--", post.Id)
	resp, err := s.getter.GetPostById(context.Background(), &connection_grpc.GetPostByIdRequest{Id: int32(post.Id)})
	if err != nil {
		log.Printf("error get one post service:-->%v\n", err)
		return nil, err
	}
	p := &models.Post{
		Id:     int(resp.GetId()),
		UserID: int(resp.GetUserId()),
		Title:  resp.GetTitle(),
		Body:   resp.GetBody(),
	}
	return p, nil
}

func (s *getPostService) GetSome(ids []int32) ([]*connection_grpc.GetPostByIdResponse, error) {
	resp, err := s.getter.GetPosts(context.Background(), &connection_grpc.GetPostsRequest{PostsId: ids})
	if err != nil {
		log.Printf("error get some posts service:-->%v\n", err)
		return nil, err
	}
	return resp.GetPosts(), nil
}
