package main

import (
	"context"
	"log"
	"net"
	"runtime"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "pumpkin-go/6-grpc_deadlines/proto"
)

type SimpleService struct {
	pb.UnimplementedSimpleServer
}

const (
	Address string = ":8000"
	Network string = "tcp"
)

func main() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")

	grpcServer := grpc.NewServer()

	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Server err: %v", err)
	}
}

func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	data := make(chan *pb.SimpleResponse, 1)
	go handle(ctx, req, data)
	select {
	case res := <-data:
		return res, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
}

func handle(ctx context.Context, req *pb.SimpleRequest, data chan<- *pb.SimpleResponse) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		runtime.Goexit()
	case <-time.After(4 * time.Second):
		res := pb.SimpleResponse{
			Code:  200,
			Value: "hello " + req.Data,
		}
		data <- &res
	}
}
