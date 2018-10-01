package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/lizebang/learning-grpc/grpc/hello"
)

func main() {
	tc, err := credentials.NewClientTLSFromFile("../certificate/certificate.pem", "")
	if err != nil {
		log.Fatalf("failed to constructs TLS credentials: %v", err)
	}
	conn, err := grpc.Dial("www.hellotest.com:8080", grpc.WithTransportCredentials(tc))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	hello(c)
	helloTimeout(c)
	helloRequestStream(c)
	helloResponseStream(c)
	helloRequestResponseStream(c)
}

func hello(c pb.HelloClient) {
	log.Println("hello")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.Hello(ctx, &pb.HelloRequest{Data: now()})
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Printf("Server: %s\n", resp.GetData())
	}
}

func helloTimeout(c pb.HelloClient) {
	log.Println("helloTimeout")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.HelloTimeout(ctx, &pb.HelloRequest{Data: now()})
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Printf("Server: %s\n", resp.GetData())
	}
}

func helloRequestStream(c pb.HelloClient) {
	log.Println("helloRequestStream")
	stream, err := c.HelloRequestStream(context.TODO())
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for index := 0; index < 10; index++ {
		time.Sleep(time.Second)
		err := stream.Send(&pb.HelloRequest{
			Data: now(),
		})
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	log.Printf("Server: %s\n", resp.GetData())
}

func helloResponseStream(c pb.HelloClient) {
	log.Println("helloResponseStream")
	stream, err := c.HelloResponseStream(context.TODO(), &pb.HelloRequest{Data: now()})
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for {
		time.Sleep(time.Second)
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("helloResponseStream EOF")
			return
		}
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}
		log.Printf("Server: %s", resp.GetData())
	}
}

func helloRequestResponseStream(c pb.HelloClient) {
	log.Println("helloRequestResponseStream")
	stream, err := c.HelloRequestResponseStream(context.TODO())
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for index := 0; index < 10; index++ {
		time.Sleep(time.Second)

		err = stream.Send(&pb.HelloRequest{Data: now()})
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("helloRequestResponseStream EOF")
			return
		}
		if err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}
		log.Printf("Server: %s\n", resp.GetData())
	}
	err = stream.CloseSend()
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func now() string {
	return time.Now().Format(time.RFC3339Nano)
}
