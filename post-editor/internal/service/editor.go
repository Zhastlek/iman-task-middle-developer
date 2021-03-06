package service

import (
	"errors"
	"iman-task/post-editor/internal/database"
	"iman-task/post-editor/internal/models"
	"log"
)

func NewEditorService(storage database.EditorStorage) ServiceEditor {
	return &editorPost{
		storage: storage,
	}
}

func (s *editorPost) Update(post *models.Post) error {
	if post == nil {
		return errors.New("error invalid update post information")
	}
	// log.Println(post.Id, post.UserID)
	if err := s.storage.Update(post); err != nil {
		log.Printf("error in update service method: -- %v\n", err)
		return err
	}
	return nil
}

func (s *editorPost) Delete(id int32) error {
	if id < 0 {
		return errors.New("error invalid delete post id")
	}
	if err := s.storage.Delete(id); err != nil {
		log.Printf("error in delete service method: -- %v\n", err)
		return err
	}
	return nil
}
