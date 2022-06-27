package app

import (
	"iman-task/config"
	"iman-task/connection_grpc"
	"iman-task/post-parser/internal/database"
	"iman-task/post-parser/internal/service"
	"iman-task/post-parser/provider"
	"log"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func Initialize() *grpc.Server {
	server := grpc.NewServer()

	config, err := config.InitConfig()
	if err != nil {
		log.Printf("error initialize config post parser : -->%v\n", err)
	}
	db, err := database.InitializeDatabase(config)
	if err != nil {
		log.Printf("error didn't initialize database:--> %v\n", err)
		log.Fatal()
	}
	createStorage := database.NewCreateStorage(db)
	createrService := service.NewCreateService(createStorage)
	createrProvider := provider.NewCreaterProvider(createrService)
	// log.Println("provider--->", createrProvider)

	connection_grpc.RegisterCreatorServer(server, createrProvider)
	return server
}
