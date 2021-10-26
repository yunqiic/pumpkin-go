package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"pumpkin-go/10-grpc_gateway/client/auth"
	pb "pumpkin-go/10-grpc_gateway/proto"
)

const Address string = ":8000"

var grpcClient pb.SimpleClient

func main() {
	credential, err := credentials.NewClientTLSFromFile("../tls/test.crt", "*.yunqiic.com")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials: %v", err)
	}
	token := auth.Token{
		Value: "bearer grpc.auth.token",
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
	req := pb.InnerMessage{
		SomeInteger: 99,
		SomeFloat:   1,
	}

	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	log.Println(res)
}
