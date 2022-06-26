package service

import (
	"iman-task/api/internal/models"
	"iman-task/connection_grpc"
)

type ServiceGetter interface {
	GetOne(post *models.Post) (*models.Post, error)
	GetSome(ids []int32) ([]*connection_grpc.GetPostByIdResponse, error)
}

type ServiceEditor interface {
	DeleteOne(id int) (string, error)
	UpdateOne(post *models.Post) (string, error)
}

type ServiceCreater interface {
	Create(url string) (string, error)
}
