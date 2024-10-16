package main

import (
	"context"
	"fmt"
	"log"

	pbServer "github.com/kodinggo/grpc-server-gb2/pb/grpc_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// connect to grpc server without credentials
	conn, err := grpc.NewClient("localhost:5555", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("failed to open connection grpc server, error %v", err)
	}

	// init grpc client as package dependency from grpc-server repository
	client := pbServer.NewStoryServiceClient(conn)

	results, _ := client.FindAll(context.TODO(), &pbServer.StoryFindAllRequest{})

	for _, story := range results.Stories {
		fmt.Println(story.Id)
		fmt.Println(story.Title)
	}
}
