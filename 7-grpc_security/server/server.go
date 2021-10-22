package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "pumpkin-go/7-grpc_security/proto"
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
	credential, err := credentials.NewServerTLSFromFile("../pkg/tls/test.pem", "../pkg/tls/test.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials: %v", err)
	}

	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = Check(ctx)
		if err != nil {
			return
		}
		return handler(ctx, req)
	}

	grpcServer := grpc.NewServer(grpc.Creds(credential), grpc.UnaryInterceptor(interceptor))
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing with TLS and token")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.serve err: %v", err)
	}
}

func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

func Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "get token fail")
	}
	var (
		appID     string
		appSecret string
	)
	if value, ok := md["app_id"]; ok {
		appID = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[1]
	}
	if appID != "grpc_token" || appSecret != "123456" {
		return status.Errorf(codes.Unauthenticated, "token err: app_id=%s, app_secret=%s", appID, appSecret)
	}
	return nil
}
