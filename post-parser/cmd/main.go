package main

import (
	"iman-task/post-parser/app"
	"log"
	"net"
)

func main() {
	server := app.Initialize()
	l, err := net.Listen("tcp", ":8008")
	log.Println("Server run is 8008 port...")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
