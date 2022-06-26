package database

import (
	"database/sql"
	"fmt"
	"iman-task/config"

	_ "github.com/lib/pq"
)

func InitializeDatabase(c *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
