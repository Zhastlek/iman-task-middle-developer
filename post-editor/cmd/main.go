package main

import (
	"iman-task/post-editor/app"
	"log"
	"net"
)

func main() {
	server := app.Initialize()
	l, err := net.Listen("tcp", ":8010")
	log.Println("Server run is 8010 port...")
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
