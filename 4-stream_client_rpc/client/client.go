package main

import (
	"context"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"

	pb "pumpkin-go/4-stream_client_rpc/proto"
)

const Address string = ":8000"

var streamClient pb.StreamClientClient

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}

	defer conn.Close()

	streamClient = pb.NewStreamClientClient(conn)
	route()
	routeList()
}

func route() {
	req := pb.SimpleRequest{
		Data: "grpc",
	}

	res, err := streamClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}

	log.Println(res)
}

func routeList() {
	stream, err := streamClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
	}
	for n := 0; n < 5; n++ {
		err := stream.Send(&pb.StreamRequest{StreamData: "stream client rpc " + strconv.Itoa(n)})
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}

		log.Println(n)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	log.Println(res)
}
