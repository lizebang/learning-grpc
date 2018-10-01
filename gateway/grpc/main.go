//go:generate protoc --proto_path=../hello --proto_path=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:../hello ../hello/hello.proto
package main

import (
	"context"
	"log"

	"github.com/lizebang/learning-grpc/gateway/server"
)

const (
	network = "tcp"
	address = ":8080"
)

func main() {
	ctx := context.Background()
	if err := server.Run(ctx, network, address); err != nil {
		log.Fatalln(err)
	}
}
