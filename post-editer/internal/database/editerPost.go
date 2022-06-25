package database

import (
	"database/sql"
	"errors"
	"fmt"
	"iman-task/post-editer/internal/models"
	"log"
)

type editerStorage struct {
	db *sql.DB
}

type EditerStorage interface {
	Update(post *models.Post) error
	Delete(post *models.Post) error
}

func NewEditerStorage(db *sql.DB) EditerStorage {
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
	defer tx.Rollback()
	_, err = tx.Exec(updateText)
	if err != nil {
		log.Printf("error update post in update post method: %v\n", err)
		// tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		log.Printf("error in tx commit update post method: %v\n", err)
		return err
	}
	return nil
}

func (s *editerStorage) Delete(post *models.Post) error {
	if post == nil {
		return errors.New("array post is empty")
	}
	deleteText := fmt.Sprintf("DELETE FROM posts WHERE id='%d' and user_id='%d'", post.Id, post.UserID)
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
