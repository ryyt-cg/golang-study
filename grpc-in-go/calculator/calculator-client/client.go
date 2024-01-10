package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"time"

	calculatorpb "gitlab.con/aionx/go-examples/grpc-in-go/pb-domain/calculator"
	"google.golang.org/grpc"
)

var log *zap.Logger

func main() {
	log, _ = zap.NewProduction()
	fmt.Println("Greetings, je suis une cliente calculatrice")
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("could not connect.", zap.Error(err))
	}

	defer func() {
		err := cc.Close()
		if err != nil {
			log.Fatal("error while closing.", zap.Error(err))
		}
	}()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to perform a Unary RPC...")
	req := &calculatorpb.SumRequest{
		A: 3,
		B: 10,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Error("error while calling Sum RPC", zap.Error(err))
	}

	log.Info("response from Sum.", zap.Int32("sum", res.SumResult))
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Initiating a server stream RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		InputNumber: 327,
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Error("error while calling Server Stream RPC.", zap.Error(err))
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			log.Info("Stream is done.")
			break // meaning nothing more is being streamed
		}
		if err != nil {
			log.Error("error while reading stream.", zap.Error(err))
		}
		log.Info("PrimeNumberDecomposition response.", zap.Int32("result", msg.GetResult()))
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to perform a Client Streaming RPC...")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Error("Error while calling ComputeAverage.", zap.Error(err))
	}

	numbers := []int64{11, 9, 7, 5, 4}

	for _, number := range numbers {
		fmt.Printf("A request is being sent with number %v\n", number)
		err := stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})

		if err != nil {
			log.Error("Error while sending.", zap.Error(err))
		}
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Error("Error while receiving response from ComputeAverage.", zap.Error(err))
	}

	fmt.Printf("Server computed an average of %v.\n", res.GetComputedAverage())
}
