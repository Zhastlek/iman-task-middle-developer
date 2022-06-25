package database

import (
	"database/sql"
	"fmt"
	"iman-task/post-editor/internal/models"
	"log"
	"strconv"
)

type getterStorage struct {
	db *sql.DB
}

type GetterStorage interface {
	GetOneById(id int) (*models.Post, error)
	GetSome(id []int) ([]*models.Post, error)
}

func NewGetterStorage(db *sql.DB) GetterStorage {
	return &getterStorage{
		db: db,
	}
}

func (s *getterStorage) GetOneById(id int) (*models.Post, error) {
	getOneText := fmt.Sprintf("SELECT * FROM posts WHERE id='%d'", id)
	post := &models.Post{}
	row := s.db.QueryRow(getOneText)
	err := row.Scan(&post.Id, &post.UserID, &post.Title, &post.Body)
	switch err {
	case sql.ErrNoRows:
		log.Println("No rows were returned!")
		return nil, err
	case nil:
		return post, nil
	default:
		log.Printf("error in get one by id method storage: --%v\n", err)
		return nil, err
	}
}

func (s *getterStorage) GetSome(id []int) ([]*models.Post, error) {
	getOneText := "SELECT * FROM posts WHERE"
	ids := ""
	text := " id="
	for i, value := range id {
		if i == len(id)-1 {
			ids = ids + text + strconv.Itoa(value)
			continue
		}
		ids = ids + text + strconv.Itoa(value) + " or"
	}
	getOneText += ids
	rows, err := s.db.Query(getOneText)
	if err != nil {
		log.Printf("error query in get some method: -- %v\n", err)
		return nil, err
	}
	defer rows.Close()
	posts := []*models.Post{}
	for rows.Next() {
		onePost := &models.Post{}
		err := rows.Scan(&onePost.Id, &onePost.UserID, &onePost.Title, &onePost.Body)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("error post storage get some method :--> %v\n", err)
				return nil, err
			}
			log.Printf("error post storage get some method :--> %v\n", err)
			return nil, err
		}
		posts = append(posts, onePost)
	}
	return posts, nil
}
