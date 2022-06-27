package database

import (
	"database/sql"
	"errors"
	"fmt"
	"iman-task/post-editor/internal/models"
	"log"
)

type editerStorage struct {
	db *sql.DB
}

type EditorStorage interface {
	Update(post *models.Post) error
	Delete(id int32) error
}

func NewEditorStorage(db *sql.DB) EditorStorage {
	return &editerStorage{
		db: db,
	}
}

func (s *editerStorage) Update(post *models.Post) error {
	if post == nil {
		return errors.New("array post is empty")
	}
	updateText := fmt.Sprintf("UPDATE posts SET title='%s', body='%s' WHERE id='%d' and user_id='%d'", post.Title, post.Body, post.Id, post.UserID)
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	// log.Println(post.Id, post.UserID)
	// log.Println(updateText)
	defer tx.Rollback()
	res, err := tx.Exec(updateText)
	if err != nil {
		log.Printf("error update post in update post method: %v\n", err)
		// tx.Rollback()
		return err
	}
	status, _ := res.RowsAffected()
	if status == 0 {
		return errors.New("error invalid id post")
	}
	if err = tx.Commit(); err != nil {
		log.Printf("error in tx commit update post method: %v\n", err)
		return err
	}
	return nil
}

func (s *editerStorage) Delete(id int32) error {
	deleteText := fmt.Sprintf("DELETE FROM posts WHERE id='%d'", id)
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	res, err := tx.Exec(deleteText)
	if err != nil {
		log.Printf("error delete post in delete post method: %v\n", err)
		// tx.Rollback()
		return err
	}
	status, _ := res.RowsAffected()
	if status == 0 {
		return errors.New("error invalid id post")
	}
	if err = tx.Commit(); err != nil {
		log.Printf("error in tx commit delete post method: %v\n", err)
		return err
	}
	return nil
}
