package main

import (
	"context"
	"fmt"
	"grpc-course/hello/hellopb"
	"io"
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

	goodbyeBidiStreaming(c)
}

func goodbyeBidiStreaming(c hellopb.HelloServiceClient) {
	fmt.Println("Starting goodbye bidi function")

	// create stream to call server
	stream, err := c.Goodbye(context.Background())
	requests := []*hellopb.GoodbyeRequest{
		&hellopb.GoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Mark",
				Prefix:    "Joven",
			},
		},
		&hellopb.GoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Jhon",
				Prefix:    "Dr",
			},
		},
		&hellopb.GoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Karla",
				Prefix:    "Ms",
			},
		},
		&hellopb.GoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Javier",
				Prefix:    "Sr",
			},
		},
		&hellopb.GoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Sandra",
				Prefix:    "Sra",
			},
		},
		&hellopb.GoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Jonathan",
				Prefix:    "Ing",
			},
		},
		&hellopb.GoodbyeRequest{
			Hello: &hellopb.Hello{
				FirstName: "Sol",
				Prefix:    "Sra",
			},
		},
	}
	if err != nil {
		log.Printf("Error creating stream %v", err)
	}

	waitc := make(chan struct{})
	// send many messages to the server (go routines)
	go func() {
		for _, req := range requests {
			log.Printf("Seending message %v", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// receive messages from the server (go routines)
	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error receiving stream %v", err)
				break
			}
			fmt.Printf("Gor it: %v\n", res.GetGoodbye())
		}
		close(waitc)
	}()
	//block when everthing is completed or closed
	<-waitc
}
