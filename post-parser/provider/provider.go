package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"iman-task/connection_grpc"
	"iman-task/post-parser/internal/models"
	"iman-task/post-parser/internal/service"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type GRPCServer struct {
	creater service.CreatePostsService
	connection_grpc.UnimplementedCreatorServer
}

func NewCreaterProvider(creater service.CreatePostsService) connection_grpc.CreatorServer {
	return &GRPCServer{
		creater: creater,
	}
}

func (g *GRPCServer) Parse(ctx context.Context, req *connection_grpc.Request) (*connection_grpc.Response, error) {
	reqString := req.GetUrl()
	log.Println("praser provider", reqString)
	i := 1
	var wg sync.WaitGroup
	//client := http.Client{Timeout: 30 * time.Second}
	//var body *http.Response
	// link := req.GetUrl()
	t := time.Now()
	// str := "https://gorest.co.in/public/v1/posts?page="
	//req, _ := http.Get(link)
	for i < 51 {
		j := 0
		for j < 6 && i < 51 {
			wg.Add(1)
			link := reqString + strconv.Itoa(i)
			go func(linknum string) {
				defer wg.Done()
				r, _ := http.Get(linknum)
				generator := &models.AutoGenerated{}
				log.Printf("body----------> %v\n", r)
				data := []*models.Post{}
				json.NewDecoder(r.Body).Decode(&generator)
				defer r.Body.Close()
				data = generator.Posts
				log.Println(data)
				if err := g.creater.Create(data); err != nil {
					log.Println("error create posts method--->", err)
				}
			}(link)
			i++
			log.Println("num gorutines ---->", i)
			j++
		}
		wg.Wait()
	}
	fmt.Printf("%s", time.Since(t))
	return &connection_grpc.Response{Status: "Posts created"}, nil
}
