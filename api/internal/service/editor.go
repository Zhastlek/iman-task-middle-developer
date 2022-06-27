package service

import (
	"context"
	"iman-task/api/internal/models"
	"iman-task/connection_grpc"
	"log"
)

type editPostService struct {
	editor connection_grpc.EditorClient
}

func NewEditorService(editor connection_grpc.EditorClient) ServiceEditor {
	return &editPostService{
		editor: editor,
	}
}

func (s *editPostService) UpdateOne(post *models.Post) (string, error) {
	resp, err := s.editor.UpdatePost(context.Background(), &connection_grpc.UpdatePostRequest{Id: int32(post.Id), UserId: int32(post.UserID), Title: post.Title, Body: post.Body})
	if err != nil {
		log.Printf("error update one post service:-->%v\n", err)
		return resp.GetStatus(), err
	}
	return resp.GetStatus(), nil
}

func (s *editPostService) DeleteOne(id int) (string, error) {
	resp, err := s.editor.DeletePost(context.Background(), &connection_grpc.DeletePostRequest{Id: int32(id)})
	if err != nil {
		log.Printf("error delete one post service:-->%v\n", err)
		return resp.GetStatus(), err
	}
	return resp.GetStatus(), nil
}
