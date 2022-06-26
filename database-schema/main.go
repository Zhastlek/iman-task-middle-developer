package main

import (
	"database/sql"
	"fmt"
	"iman-task/config"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	log.Println(111)
	c, err := config.InitConfig()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(c)
	psqlInfo := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return
	}

	if err = db.Ping(); err != nil {
		return
	}
	if err = createTable(db); err != nil {
		return
	}
	log.Println("good")
}

func createTable(db *sql.DB) error {
	dir, err := ioutil.ReadDir("./database-schema/up/")
	if err != nil {
		return err
	}
	log.Println(22, dir)
	for _, v := range dir {
		log.Println(v.Name())
		body, err := ioutil.ReadFile("./database-schema/up/" + v.Name())
		if err != nil {
			return err
		}
		log.Println(string(body))
		_, err = db.Exec(string(body))
		if err != nil {
			log.Println("------", err)
			return err
		}
	}
	return nil
}
