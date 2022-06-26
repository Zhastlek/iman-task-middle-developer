package service

import (
	"iman-task/connection_grpc"
	"iman-task/post-editor/internal/database"
	"iman-task/post-editor/internal/models"
)

type ServiceGetter interface {
	GetOneById(id int32) (*connection_grpc.GetPostByIdResponse, error)
	GetSome(id []int32) (*connection_grpc.GetPostsResponse, error)
}

type ServiceEditor interface {
	Update(post *models.Post) error
	Delete(id int32) error
}

type getterPost struct {
	storage database.GetterStorage
}

type editorPost struct {
	storage database.EditorStorage
}
