package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/lizebang/learning-grpc/gateway/hello"
)

// Run starts the gRPC service.
func Run(ctx context.Context, network, address string) error {
	tc, err := credentials.NewServerTLSFromFile("../certificate/certificate.pem", "../certificate/key.pem")
	if err != nil {
		log.Printf("failed to constructs TLS credentials: %v", err)
		return err
	}
	s := grpc.NewServer(grpc.Creds(tc))

	lis, err := net.Listen(network, address)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}
	defer func() {
		if err := lis.Close(); err != nil {
			log.Printf("fail to close: %v", err)
		}
	}()

	pb.RegisterHelloServer(s, newServer())

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()

	return s.Serve(lis)
}
