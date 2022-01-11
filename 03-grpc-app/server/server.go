package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
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
