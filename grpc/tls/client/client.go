package main

import (
	pb "github.com/fengchunjian/goexamples/grpc/tls/proto" //引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	//TLS连接
	creds, err := credentials.NewClientTLSFromFile("../keys/server.pem", "server name")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := pb.NewHelloClient(conn)
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r.Message)
}
