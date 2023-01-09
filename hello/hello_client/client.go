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

	//helloUnary(c)

	//helloServerStreaming(c)

	// goodbyeClientStreaming(c)

	goodbyeBidiStreaming(c)
}

func helloUnary(c hellopb.HelloServiceClient) {
	fmt.Println("Starting unary RPC Hello")

	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{
			FirstName: "Mark",
			Prefix:    "Joven",
		},
	}

	res, err := c.Hello(context.Background(), req)

	if err != nil {
		log.Fatalf("Error, calling Hello RPC: \n%v", err)
	}

	log.Printf("Response Hello: %v", res.CustomHello)
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
