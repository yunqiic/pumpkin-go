package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "pumpkin-go/6-grpc_deadlines/proto"
)

const Address string = ":8000"

var grpcClient pb.SimpleClient

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	grpcClient = pb.NewSimpleClient(conn)
	route(ctx, 2)
}

func route(ctx context.Context, deadlines time.Duration) {
	clientDeadline := time.Now().Add(time.Duration(deadlines * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()

	req := pb.SimpleRequest{
		Data: "grpc",
	}

	res, err := grpcClient.Route(ctx, &req)
	if err != nil {
		status, ok := status.FromError(err)
		if ok {
			if status.Code() == codes.DeadlineExceeded {
				log.Fatalf("Route timeout!")
			}
		}
		log.Fatalf("Call Route err: %v", err)
	}
	log.Println(res.Value)
}
