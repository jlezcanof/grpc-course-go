package main

import (
	"context"
	"fmt"
	"grpc-course/hello/hellopb"
	"log"
	"time"

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

	goodbyeClientStreaming(c)

}

func goodbyeClientStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting goodbye function")

	requests := []*hellopb.HellosGoodbyeRequest{
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Mark",
				Prefix:    "Joven",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Jhon",
				Prefix:    "Dr",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Karla",
				Prefix:    "Ms",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Javier",
				Prefix:    "Sr",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Sandra",
				Prefix:    "Sra",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Jonathan",
				Prefix:    "Ing",
			},
		},
		&hellopb.HellosGoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Sol",
				Prefix:    "Sra",
			},
		},
	}

	stream, err := c.HellosGoodbye(context.Background())

	if err != nil {
		log.Printf("Error calling goodbye %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)

		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	goodbye, err := stream.CloseAndRecv()

	if err != nil {
		log.Printf("Error goodbye receive %v", err)
	}

	fmt.Println(goodbye)
}
