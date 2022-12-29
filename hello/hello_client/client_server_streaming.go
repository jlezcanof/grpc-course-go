package main

import (
	"context"
	"fmt"
	"grpc-course/hello/hellopb"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Go client is running")

	certFile := "../../ssl/ca.crt"

	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")

	if sslErr != nil {
		log.Fatalf("Failed reading CA certificate %v", sslErr)
	}

	opts := grpc.WithTransportCredentials(creds)

	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)

	helloServerStreaming(c)
}

func helloServerStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting server streaming RPC Hello")

	req := &hellopb.HelloManyLanguagesRequest{
		Hello: &hellopb.Hello{
			FirstName: "Mark",
			Prefix:    "Joven",
		},
	}

	restStream, err := c.HelloManyLanguages(context.Background(), req)

	if err != nil {
		log.Printf("Error calling Hello Many languages %v", err)
	}

	for {
		msg, err := restStream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading stream %v", err)
		}

		log.Printf("Res from HML: %v\n", msg.GetHelloLanguage())
	}
}
