package service

import (
	"iman-task/api/internal/models"
)

type getPostService struct {
}

func NewGetterService() ServiceGetter {
	return &getPostService{}
}

func (s *getPostService) GetOne(post *models.Post) {

}

func (s *getPostService) GetSome(ids []int) {

}
