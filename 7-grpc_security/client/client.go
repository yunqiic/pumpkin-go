package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"pumpkin-go/7-grpc_security/pkg/auth"
	pb "pumpkin-go/7-grpc_security/proto"
)

const Address string = ":8000"

var grpcClient pb.SimpleClient

func main() {
	credential, err := credentials.NewClientTLSFromFile("../pkg/tls/server.pem", "pumpkin-go")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials: %v", err)
	}

	token := auth.Token{
		AppID:     "grpc_token",
		AppSecret: "123456",
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(credential), grpc.WithPerRPCCredentials(&token))
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
