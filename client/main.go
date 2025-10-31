package main

import (
	"context"
	"log"
	pb "pb/protofile"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := pb.NewHelloSayerClient(conn)

	r, err := c.SayHello(ctx, &pb.HelloMessage{GreetingMessage: "WAW WAW"})

	if err != nil {
		log.Fatalf("Could not gree dis %v", err)
	}
	log.Printf("Greetings Traveler: %v", r.GetReplyMessage())
}
