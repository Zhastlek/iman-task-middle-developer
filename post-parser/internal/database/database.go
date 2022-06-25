package database

import (
	"errors"
	"fmt"
	"iman-task/post-parser/internal/models"
	"log"
)

func (s *storage) CreatePosts(posts []*models.Post) error {
	if len(posts) == 0 || posts == nil {
		return errors.New("array post is empty")
	}
	insertText := "INSERT INTO posts(id, user_id, title, body) VALUES"
	allPosts := ""
	for i, value := range posts {
		if i == len(posts)-1 {
			allPosts += fmt.Sprintf("('%d','%d','%s','%s')", value.Id, value.UserID, value.Title, value.Body)
			continue
		}
		allPosts += fmt.Sprintf("('%d','%d','%s','%s'),", value.Id, value.UserID, value.Title, value.Body)
	}
	insertText += allPosts
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.Exec(insertText)
	if err != nil {
		log.Printf("error insert posts in create posts method: %v\n", err)
		// tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		log.Printf("error in tx commit create posts method: %v\n", err)
		return err
	}
	return nil
}
