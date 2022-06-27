package app

import (
	"iman-task/api/internal/handlers"
	"iman-task/api/internal/service"
	"iman-task/connection_grpc"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Initialize() *gin.Engine {
	// ctx := con
	router := gin.Default()
	connParser, err := grpc.Dial(":8008", grpc.WithInsecure())
	if err != nil {
		log.Println("connection parser port 8008 :-->", err)
	}
	connEditor, err := grpc.Dial(":8010", grpc.WithInsecure())
	if err != nil {
		log.Println("connection editor port 8010 :-->", err)
	}

	clientParser := connection_grpc.NewCreatorClient(connParser)
	clientEditor := connection_grpc.NewEditorClient(connEditor)
	clientGetter := connection_grpc.NewEditorClient(connEditor)
	// log.Println("client Parser--", clientParser)

	createrService := service.NewCreaterService(clientParser)
	editorService := service.NewEditorService(clientEditor)
	getterService := service.NewGetterService(clientGetter)

	handler := handlers.NewHandler(createrService, getterService, editorService)
	handler.Register(router)
	return router
}
