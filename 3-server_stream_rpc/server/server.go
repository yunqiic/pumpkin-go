package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "pumpkin-go/3-server_stream_rpc/proto"
)

type StreamService struct {
	pb.UnimplementedStreamServerServer
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

	pb.RegisterStreamServerServer(grpcServer, &StreamService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}

func (s *StreamService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

func (s *StreamService) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for n := 0; n < 5; n++ {
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
