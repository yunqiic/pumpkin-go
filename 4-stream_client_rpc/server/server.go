package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "pumpkin-go/4-stream_client_rpc/proto"
)

type SimpleService struct {
	pb.UnimplementedStreamClientServer
}

func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

func (s *SimpleService) RouteList(srv pb.StreamClient_RouteListServer) error {
	for {
		res, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
		}
		if err != nil {
			return err
		}
		log.Println(res.StreamData)
	}
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

	pb.RegisterStreamClientServer(grpcServer, &SimpleService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Server err: %v", err)
	}
}
