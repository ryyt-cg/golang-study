package main

import (
	"context"
	"fmt"
	greetpb "gitlab.con/aionx/go-examples/grpc-in-go/pb-domain/greet"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"time"
)

var log *zap.Logger

func main() {
	log, _ = zap.NewProduction()
	log.Info("Greetings, je suis un client")

	// Create gRPC client connection
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect.", zap.Error(err))
	}

	// close connection right before exiting main function
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to perform a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Charles",
			LastName:  "Naylor",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Greet RPC.", zap.Error(err))
	}

	log.Info("response from Greet.", zap.String("result", res.Result))
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to perform a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Charles",
			LastName:  "Naylor",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling Server Stream RPC.", zap.Error(err))
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			log.Info("Stream is over.")
			break // meaning nothing more is being streamed
		}
		if err != nil {
			log.Fatal("error while reading stream.", zap.Error(err))
		}
		log.Info("GreetManyTimes response.", zap.String("result", msg.GetResult()))
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to perform a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Charles",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Diana",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sara",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mindy",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatal("Error while calling LongGreet.", zap.Error(err))
	}

	for _, req := range requests {
		fmt.Printf("A request is being sent: %v\n", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error while receiving response from LongGreet.", zap.Error(err))
	}

	fmt.Printf("The server responds with: \n%v", res.GetResult())
}
