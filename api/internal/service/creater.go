package service

import (
	"context"
	"iman-task/connection_grpc"
	"log"
)

type createrPostService struct {
	client connection_grpc.CreatorClient
}

func NewCreaterService(client connection_grpc.CreatorClient) ServiceCreater {
	return &createrPostService{
		client: client,
	}
}

func (service *createrPostService) Create(url string) (string, error) {
	ctx := context.Background()
	log.Println("uuuurrrrlll", url)
	status, err := service.client.Parse(ctx, &connection_grpc.Request{Url: url})
	if err != nil {
		log.Printf("error parse grpc method:-->%v\n", err)
		return status.GetStatus(), err
	}
	return status.GetStatus(), nil
}
