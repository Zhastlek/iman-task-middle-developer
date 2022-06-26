package service

import (
	"iman-task/api/internal/models"
	"iman-task/connection_grpc"
)

type getPostService struct {
	getter connection_grpc.EditorClient
}

func NewGetterService(getter connection_grpc.EditorClient) ServiceGetter {
	return &getPostService{
		getter: getter,
	}
}

func (s *getPostService) GetOne(post *models.Post) {

}

func (s *getPostService) GetSome(ids []int) {

}
