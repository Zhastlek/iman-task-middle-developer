package service

import (
	"errors"
	"iman-task/post-parser/internal/database"
	"iman-task/post-parser/internal/models"
	"log"
)

type store struct {
	storage database.CreatePostsStorage
}
type CreatePostsService interface {
	Create(posts []*models.Post) error
}

func NewCreateService(storage database.CreatePostsStorage) CreatePostsService {
	return &store{
		storage: storage,
	}
}

func (s *store) Create(posts []*models.Post) error {
	if len(posts) == 0 || posts == nil {
		log.Println("error in create posts service array posts is empty")
		return errors.New("error in create posts service array posts is empty")
	}
	if err := s.storage.CreatePosts(posts); err != nil {
		log.Printf("error in create posts service answer storage: --> %v\n", err)
		return err
	}
	return nil
}
