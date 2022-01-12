package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Processing %d and %d\n", x, y)
	time.Sleep(2 * time.Second)
	result := x + y
	fmt.Printf("Sending result %d\n", result)
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func (s *server) GeneratePrimes(req *proto.PrimeRequest, stream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Generating primes from %d and %d\n", start, end)
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Printf("Sending Prime No : %d\n", i)
			res := &proto.PrimeResponse{
				PrimeNo: i,
			}
			err := stream.Send(res)
			if err != nil {
				return err
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func (s *server) ComputeAverage(serverStream proto.AppService_ComputeAverageServer) error {
	var sum int32
	var count int32
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.GetNo()
		fmt.Println("Received No : ", req.GetNo())
		count++
	}
	avg := sum / count
	res := &proto.AverageResponse{
		Average: avg,
	}
	fmt.Println("Sending Average : ", avg)
	return serverStream.SendAndClose(res)
}

func isPrime(no int32) bool {
	if no < 2 {
		return false
	}
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (s *server) GreetEveryone(stream proto.AppService_GreetEveryoneServer) error {
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		user := req.GetUser()
		firstName := user.GetFirstName()
		lastName := user.GetLastName()
		fmt.Printf("Received %s %s\n", firstName, lastName)
		msg := fmt.Sprintf("Hello %s %s", firstName, lastName)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Sending msg : ", msg)
		res := &proto.GreetEveryoneResponse{
			Greeting: msg,
		}
		err = stream.Send(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	s := &server{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, s)
	grpcServer.Serve(listener)
}
