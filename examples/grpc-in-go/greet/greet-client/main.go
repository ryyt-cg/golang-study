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
	defer func() {
		err := cc.Close()
		if err != nil {
			log.Error("Error while close gRPC connection.", zap.Error(err))
		}
	}()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)
	doServerStreaming(c)
	doClientStreaming(c)
	doBidiStreaming(c)
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
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Charles",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Diana",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sara",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},
		{
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
		err := stream.Send(req)
		if err != nil {
			log.Error("Error while sending", zap.Error(err))
		}

		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Error while receiving response from LongGreet.", zap.Error(err))
	}

	fmt.Printf("The server responds with: \n%v", res.GetResult())
}

func doBidiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to perform a Bidirectional Streaming RPC...")

	// create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatal("Error while creating stream.", zap.Error(err))
	}

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Stephane",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Lucy",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Mark",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Piper",
			},
		},
	}

	waitCh := make(chan struct{})

	// we send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			err := stream.Send(req)
			if err != nil {
				log.Error("Error while sending", zap.Error(err))
			}
			time.Sleep(1000 * time.Millisecond)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Error("Error while close sending stream", zap.Error(err))
		}
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("Error while receiving", zap.Error(err))
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitCh)
	}()

	// block until everything is done
	<-waitCh
}
