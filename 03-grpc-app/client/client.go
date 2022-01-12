package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	//doClientStreaming(client, ctx)
	//doBiDirectionalStreaming(client, ctx)
	doRequestResponseWithTimeout(client, ctx)
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

func doBiDirectionalStreaming(client proto.AppServiceClient, ctx context.Context) {
	stream, err := client.GreetEveryone(ctx)
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				done <- true
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println(res.GetGreeting())
		}
	}()

	users := []proto.UserName{
		proto.UserName{
			FirstName: "Magesh",
			LastName:  "Kuppan",
		},
		proto.UserName{
			FirstName: "Suresh",
			LastName:  "Kannan",
		},
		proto.UserName{
			FirstName: "Ramesh",
			LastName:  "Jayaraman",
		},
		proto.UserName{
			FirstName: "Rajesh",
			LastName:  "Pandit",
		},
		proto.UserName{
			FirstName: "Ganesh",
			LastName:  "Kumar",
		},
	}

	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range users {
		time.Sleep(time.Second * 2)
		fmt.Println("Sending User : ", user)
		req := &proto.GreetEveryoneRequest{
			User: &user,
		}
		err := stream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
	}

	<-done
}

func doRequestResponseWithTimeout(client proto.AppServiceClient, ctx context.Context) {
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()
	res, err := client.Add(timeoutCtx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Println("Timout error")
			} else {
				log.Fatalln(err)
			}
		}
		log.Fatalln(err)
	}
	log.Println(res.GetResult())
}
