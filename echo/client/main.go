package main

import (
	"context"
	"github.com/DrmagicE/grpc-practice/echo/echo"
	"google.golang.org/grpc"
	"log"
	"time"
	"io"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := echo.NewEchoServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r, err := c.UnaryEcho(ctx, &echo.Message{Payload: "UnaryEcho", MessageId: 1})
	if err != nil {
		log.Fatalf("UnaryEcho err: %v \n", err)
	}
	log.Printf("UnaryEcho returns: %v \n", r)

	clientStream, err := c.ClientStreamingEcho(context.Background())
	if err != nil {
		log.Fatalf("ClientStreamingEcho err: %v \n", err)
	}
	err = clientStream.Send(&echo.Message{Payload: "abc", MessageId: 1})
	if err != nil {
		log.Fatalf("ClientStreamingEcho clientStream.Send() err: %v \n", err)
	}
	m, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("ClientStreamingEcho clientStream.CloseAndRecv() err: %v \n", err)
	}
	log.Printf("ClientStreamingEcho returns: %v \n", m)

	serverStream, err := c.ServerStreamingEcho(ctx, &echo.Message{Payload: "ServerStreamingEcho", MessageId: 1})
	for {
		m, err := serverStream.Recv()
		if err == nil {
			log.Printf("ServerStreamingEcho returns: %v \n", m)
		}
		if err == io.EOF {
			break;
		}
		if err != nil {
			log.Fatalf("ServerStreamingEcho serverStream.Recv() err: %v \n", err)
		}
	}

	bStream, err := c.BidirectionalStreaming(ctx)
	err = bStream.Send(&echo.Message{Payload:"BidirectionalStreaming", MessageId:1})
	if err != nil {
		log.Fatalf("BidirectionalStreaming bStream.Send() err: %v \n", err)
	}
	m, err = bStream.Recv()
	if err != nil {
		log.Fatalf("ServerStreamingEcho serverStream.Recv() err: %v \n", err)
	}
	log.Printf("BidirectionalStreaming returns: %v \n", m)
	err = bStream.Send(&echo.Message{Payload:"BidirectionalStreaming", MessageId:2})
	if err != nil {
		log.Fatalf("BidirectionalStreaming bStream.Send() err: %v \n", err)
	}
	m, err = bStream.Recv()
	if err != nil {
		log.Fatalf("ServerStreamingEcho serverStream.Recv() err: %v \n", err)
	}
	log.Printf("BidirectionalStreaming returns: %v \n", m)

}
