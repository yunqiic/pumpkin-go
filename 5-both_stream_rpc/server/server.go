package main

import (
	"context"
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "pumpkin-go/5-both_stream_rpc/proto"
)

type StreamService struct {
	pb.UnimplementedStreamServer
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

	pb.RegisterStreamServer(grpcServer, &StreamService{})

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

func (s *StreamService) Conversations(srv pb.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&pb.StreamResponse{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + "question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++

		log.Printf("from stream client question: %s", req.Question)
	}
}
