//go:generate protoc --proto_path=../hello --go_out=plugins=grpc:../hello ../hello/hello.proto
package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/lizebang/learning-grpc/grpc/hello"
)

type server struct{}

func main() {
	tc, err := credentials.NewServerTLSFromFile("../certificate/certificate.pem", "../certificate/key.pem")
	if err != nil {
		log.Fatalf("failed to constructs TLS credentials: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(tc))

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func() {
		if err := lis.Close(); err != nil {
			log.Fatalf("fail to close: %v", err)
		}
	}()

	pb.RegisterHelloServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
