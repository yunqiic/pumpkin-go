package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "pumpkin-go/2-simple_rpc/proto"
)

const Address string = ":8000"

var grpcClient pb.SimpleClient

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	grpcClient = pb.NewSimpleClient(conn)
	route()
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
