package service

import "iman-task/connection_grpc"

type createrPostService struct {
	client connection_grpc.CreatorClient
}

func NewCreaterService(client connection_grpc.CreatorClient) ServiceCreater {
	return &createrPostService{
		client: client,
	}
}

func (service *createrPostService) Create() {

}
