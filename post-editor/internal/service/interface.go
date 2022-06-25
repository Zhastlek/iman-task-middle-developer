package service

import (
	"iman-task/post-editor/internal/database"
	"iman-task/post-editor/internal/models"
)

type ServiceGetter interface {
	GetOneById(id int) (*models.Post, error)
	GetSome(id []int) ([]*models.Post, error)
}

type ServiceEditor interface {
	Update(post *models.Post) error
	Delete(post *models.Post) error
}

type getterPost struct {
	storage database.GetterStorage
}

type editorPost struct {
	storage database.EditorStorage
}
