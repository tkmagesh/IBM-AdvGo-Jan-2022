package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

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
	//doServerStreaming(client, ctx)
	doClientStreaming(client, ctx)
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
		if res.GetPrimeNo() == 29 {
			break
		}
	}
}

func doClientStreaming(client proto.AppServiceClient, ctx context.Context) {
	nos := []int32{5, 1, 6, 3, 8, 2, 4, 9, 7}
	stream, err := client.ComputeAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		fmt.Println("Sending No : ", no)
		req := &proto.AverageRequest{
			No: no,
		}
		err := stream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Average = ", res.GetAverage())
}
