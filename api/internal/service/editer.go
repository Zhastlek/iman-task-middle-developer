package service

import (
	"iman-task/api/internal/models"
)

type editPostService struct {
}

func NewEditerService() ServiceEditer {
	return &editPostService{}
}

func (s *editPostService) UpdateOne(post *models.Post) {

}

func (s *editPostService) DeleteOne(id int) {

}
