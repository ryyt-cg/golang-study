package main

import (
	"context"
	"fmt"
	"gitlab.con/aionx/go-examples/grpc-in-go/pb-domain/calculator"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculator.SumRequest) (*calculator.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	a, b := req.GetA(), req.GetB()
	sum := a + b
	res := &calculator.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculator.PrimeNumberDecompositionRequest, stream calculator.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Prime Number Decomposition function was invoked with %v\n", req)

	n := req.GetInputNumber()
	var divider int32 = 2

	for n > 1 {
		if n%divider == 0 {
			res := &calculator.PrimeNumberDecompositionResponse{
				Result: divider,
			}
			stream.Send(res)
			n = n / divider
		} else {
			divider = divider + 1
		}
		fmt.Printf("current values: n=%v ; divider=%v\n", n, divider)
	}

	return nil
}

func (*server) ComputeAverage(stream calculator.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with a client stream mechanism.\n")
	counter := 0
	sumTotal := int64(0)
	for {
		fmt.Printf("count=%v, sumTotal=%v\n", counter, sumTotal)
		req, err := stream.Recv()

		if err == io.EOF { // client stream fully read
			return stream.SendAndClose(&calculator.ComputeAverageResponse{
				ComputedAverage: float64(sumTotal) / float64(counter),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		sumTotal += req.GetNumber()
		counter++
	}
}

func main() {
	fmt.Println("Calc server is a-go!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve?! %v", err)
	}
}
