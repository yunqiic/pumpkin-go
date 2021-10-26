package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"

	pb "pumpkin-go/10-grpc_gateway/proto"
	"pumpkin-go/10-grpc_gateway/server/gateway"
	"pumpkin-go/10-grpc_gateway/server/middleware/auth"
	"pumpkin-go/10-grpc_gateway/server/middleware/cred"
	"pumpkin-go/10-grpc_gateway/server/middleware/recovery"
	"pumpkin-go/10-grpc_gateway/server/middleware/zap"
)

type SimpleService struct {
	pb.UnimplementedSimpleServer
}

func (s *SimpleService) Route(ctx context.Context, req *pb.InnerMessage) (*pb.OuterMessage, error) {
	res := pb.OuterMessage{
		ImportantString: "hello grpc validator",
		Inner:           req,
	}
	return &res, nil
}

const (
	Address string = ":8000"
	Network string = "tcp"
)

func main() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Lister err: %v", err)
	}

	grpcServer := grpc.NewServer(cred.TLSInterceptor(),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),
	)
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing with TLS and token...")
	/*err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}*/
	httpServer := gateway.ProvideHTTP(Address, grpcServer)
	if err = httpServer.Serve(tls.NewListener(listener, httpServer.TLSConfig)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	/*if err = httpServer.Serve(listener); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
}
