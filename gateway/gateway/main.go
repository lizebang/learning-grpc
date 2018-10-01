//go:generate protoc --proto_path=../hello --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:../hello ../hello/hello.proto
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	gw "github.com/lizebang/learning-grpc/gateway/hello"
)

const (
	endpoint = "www.hellotest.com:8080"
	address  = ":9090"
)

func main() {
	tc, err := credentials.NewClientTLSFromFile("../certificate/certificate.pem", "")
	if err != nil {
		log.Fatalf("failed to constructs TLS credentials: %v", err)
	}
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(tc))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	mux := runtime.NewServeMux()
	err = gw.RegisterHelloHandler(ctx, mux, conn)
	if err != nil {
		log.Fatalf("failed to register hello handler: %v", err)
	}

	err = http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatalf("failed to run gateway: %v", err)
	}
}
