package app

import (
	"iman-task/config"
	"iman-task/connection_grpc"
	"iman-task/post-editor/internal/database"
	"iman-task/post-editor/internal/service"
	"iman-task/post-editor/provider"
	"log"

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
	getterStorage := database.NewGetterStorage(db)
	editorStorage := database.NewEditorStorage(db)

	getterService := service.NewGetterService(getterStorage)
	editorService := service.NewEditorService(editorStorage)

	editorProvider := provider.NewProvider(editorService, getterService)
	connection_grpc.RegisterEditorServer(server, editorProvider)
	return server
}
