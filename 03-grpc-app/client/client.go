package main

import (
	"context"
	"grpc-app/proto"
	"io"
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
	ctx := context.Background()

	//doRequestResponse(client, ctx)
	doServerStreaming(client, ctx)
}

func doRequestResponse(client proto.AppServiceClient, ctx context.Context) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.GetResult())
}

func doServerStreaming(client proto.AppServiceClient, ctx context.Context) {
	req := &proto.PrimeRequest{
		Start: 6,
		End:   99,
	}
	stream, err := client.GeneratePrimes(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Prime No = ", res.GetPrimeNo())
	}
}
