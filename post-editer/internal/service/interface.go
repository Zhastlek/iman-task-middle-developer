package service

import (
	"iman-task/post-editer/internal/database"
	"iman-task/post-editer/internal/models"
)

type ServiceGetter interface {
	GetOneById(id int) (*models.Post, error)
	GetSome(id []int) ([]*models.Post, error)
}

type ServiceEditer interface {
	Update(post *models.Post) error
	Delete(post *models.Post) error
}

type getterPost struct {
	storage database.GetterStorage
}

type editerPost struct {
	storage database.EditerStorage
}
