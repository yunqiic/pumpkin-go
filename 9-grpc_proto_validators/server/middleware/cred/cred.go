package cred

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TLSInterceptor() grpc.ServerOption {
	credential, err := credentials.NewServerTLSFromFile("../tls/test.crt", "../tls/ca.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	return grpc.Creds(credential)
}
