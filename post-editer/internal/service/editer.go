package service

import (
	"errors"
	"iman-task/post-editer/internal/database"
	"iman-task/post-editer/internal/models"
	"log"
)

func NewEditerService(storage database.EditerStorage) ServiceEditer {
	return &editerPost{
		storage: storage,
	}
}

func (s *editerPost) Update(post *models.Post) error {
	if post == nil {
		return errors.New("error invalid update post information")
	}
	if err := s.storage.Update(post); err != nil {
		log.Printf("error in update service method: -- %v\n", err)
		return err
	}
	return nil
}

func (s *editerPost) Delete(post *models.Post) error {
	if post == nil {
		return errors.New("error invalid delete post information")
	}
	if err := s.storage.Delete(post); err != nil {
		log.Printf("error in delete service method: -- %v\n", err)
		return err
	}
	return nil
}
