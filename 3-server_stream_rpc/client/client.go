package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "pumpkin-go/3-server_stream_rpc/proto"
)

const Address string = ":8000"

var grpcClient pb.StreamServerClient

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	grpcClient = pb.NewStreamServerClient(conn)
	route()
	listValue()
}

func route() {
	req := pb.SimpleRequest{
		Data: "grpc",
	}
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	log.Println(res)
}

func listValue() {
	req := pb.SimpleRequest{
		Data: "stream server grpc",
	}

	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}

		log.Println(res.StreamValue)
	}
}
