package main

import (
	"context"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	//proxy instance
	client := proto.NewAppServiceClient(clientConn)
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.GetResult())
}
