package service

import (
	"iman-task/api/internal/models"
	"iman-task/connection_grpc"
)

type editPostService struct {
	editor connection_grpc.EditorClient
}

func NewEditorService(editor connection_grpc.EditorClient) ServiceEditor {
	return &editPostService{
		editor: editor,
	}
}

func (s *editPostService) UpdateOne(post *models.Post) {

}

func (s *editPostService) DeleteOne(id int) {

}
