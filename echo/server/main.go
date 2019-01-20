package main

import (
	"context"
	"log"
	"net"

	"github.com/DrmagicE/grpc-practice/echo/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
)

const (
	port = ":50051"
)

type server struct{}

//UnaryEcho 从客户端接收一个Message，在payload中增加"reply"前缀后返回给客户端
func (s *server) UnaryEcho(ctx context.Context, message *echo.Message) (*echo.Message, error) {
	log.Printf("UnaryEcho Called:\n")
	log.Printf("Received message %v \n", message)
	return &echo.Message{Payload: "reply " + message.Payload, MessageId: message.MessageId}, nil
}

//ClientStreamingEcho 从客户端读取一个Message流，返回最后一个Message
func (s *server) ClientStreamingEcho(stream echo.EchoService_ClientStreamingEchoServer) error {
	log.Printf("ClientStreamingEcho Called:\n")
	var lastMsg *echo.Message
	for {
		m, err := stream.Recv()
		if err == nil {
			log.Printf("Received message %v \n", m)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		lastMsg = m
	}
	err := stream.SendAndClose(lastMsg)
	return err
}

//ServerStreamingEcho 从客户端接收到Message后，返回一个Message流。
func (s *server) ServerStreamingEcho(message *echo.Message, stream echo.EchoService_ServerStreamingEchoServer) error {
	log.Printf("ServerStreamingEcho Called:\n")
	log.Printf("Received message %v \n", message)
	for i := 0; i < 10; i++ {
		m := &echo.Message{Payload: message.Payload, MessageId: message.MessageId + int32(i)}
		err := stream.Send(m)
		if err != nil {
			return err
		}
	}
	return nil
}

//请求响应模式，一次请求对应一次响应。
func (s *server) BidirectionalStreaming(stream echo.EchoService_BidirectionalStreamingServer) error {
	log.Printf("BidirectionalStreaming Called:\n")
	for {
		m, err := stream.Recv()
		if err == nil {
			log.Printf("Received message %v \n", m)
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(m); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
