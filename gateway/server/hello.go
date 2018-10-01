package server

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/lizebang/learning-grpc/gateway/hello"
)

type server struct{}

func newServer() *server {
	return &server{}
}

func (s *server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("Hello")
	log.Printf("Client: %s\n", in.GetData())
	return &pb.HelloResponse{
		Data: now(),
	}, nil
}

func (s *server) HelloTimeout(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("HelloTimeout")
	log.Printf("Client: %s\n", in.GetData())
	time.Sleep(10 * time.Second)
	return &pb.HelloResponse{
		Data: now(),
	}, nil
}

func (s *server) HelloRequestStream(stream pb.Hello_HelloRequestStreamServer) error {
	log.Println("HelloRequestStream")
	ns := []string(nil)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("HelloRequestStream EOF")
			return stream.SendAndClose(&pb.HelloResponse{
				Data: now(),
			})
		}
		if err != nil {
			return err
		}
		log.Printf("Client: %s\n", in.GetData())
		ns = append(ns, in.GetData())
	}
}

func (s *server) HelloResponseStream(in *pb.HelloRequest, stream pb.Hello_HelloResponseStreamServer) error {
	log.Println("HelloResponseStream")
	log.Printf("Client: %s\n", in.GetData())
	for index := 0; index < 10; index++ {
		if err := stream.Send(&pb.HelloResponse{
			Data: now(),
		}); err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}
	}
	return nil
}

func (s *server) HelloRequestResponseStream(stream pb.Hello_HelloRequestResponseStreamServer) error {
	log.Println("HelloRequestResponseStream")
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("HelloRequestResponseStream EOF")
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Client: %s\n", in.GetData())
		err = stream.Send(&pb.HelloResponse{
			Data: now(),
		})
		if err != nil {
			return err
		}
	}
}

func now() string {
	return time.Now().Format(time.RFC3339Nano)
}
