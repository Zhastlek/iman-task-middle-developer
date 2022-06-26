package service

import (
	"errors"
	"iman-task/connection_grpc"
	"iman-task/post-editor/internal/database"
	"log"
)

func NewGetterService(storage database.GetterStorage) ServiceGetter {
	return &getterPost{
		storage: storage,
	}
}

func (s *getterPost) GetOneById(id int32) (*connection_grpc.GetPostByIdResponse, error) {
	if id < 1 {
		return nil, errors.New("error invalid id post")
	}
	post, err := s.storage.GetOneById(id)
	if err != nil {
		log.Printf("error in getOneById getter service method: -- %v\n", err)
		return nil, err
	}
	return post, nil
}

func (s *getterPost) GetSome(id []int32) (*connection_grpc.GetPostsResponse, error) {
	for _, val := range id {
		if val < 1 {
			return nil, errors.New("error invalid id post")
		}
	}
	posts, err := s.storage.GetSome(id)
	if err != nil {
		log.Printf("error in GetSome getter service method: -- %v\n", err)
		return nil, err
	}
	return posts, nil
}
