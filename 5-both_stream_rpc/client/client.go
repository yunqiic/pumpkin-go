package main

import (
	"context"
	"io"
	"log"
	"strconv"

	"google.golang.org/grpc"

	pb "pumpkin-go/5-both_stream_rpc/proto"
)

const Address string = ":8000"

var streamClient pb.StreamClient

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err:%v", err)
	}
	defer conn.Close()

	streamClient = pb.NewStreamClient(conn)
	route()
	conversations()
}

func route() {
	req := pb.SimpleRequest{
		Data: "grpc",
	}

	res, err := streamClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}

	log.Println(res.Value)
}

func conversations() {
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	for n := 0; n < 5; n++ {
		err := stream.Send(&pb.StreamRequest{Question: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}

		log.Println(res.Answer)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v", err)
	}
}
