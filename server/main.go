package main

import (
	"context"
	"fmt"
	"net"
	pb "pb/protofile"

	"google.golang.org/grpc"
)

var address string = ":50051"

type server struct {
	pb.UnimplementedHelloSayerServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloMessage) (*pb.HelloReply, error) {
	fmt.Println("We Received Message !!")
	fmt.Println(in.GetGreetingMessage())
	return &pb.HelloReply{ReplyMessage: in.GetGreetingMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("failed to listen")
	}

	s := grpc.NewServer()
	pb.RegisterHelloSayerServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve")
	}
}
