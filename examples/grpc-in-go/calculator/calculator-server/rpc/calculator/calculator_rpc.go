package calculator

import (
	"context"
	"fmt"
	"io"
	"log"

	calculatorpb "gitlab.con/aionx/go-examples/grpc-in-go/pb-domain/calculator"
	"go.uber.org/zap"
)

type Rpc struct {
	log *zap.Logger
}

func NewCalculatorRpc() *Rpc {
	return &Rpc{zap.L()}
}

func (rpc *Rpc) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v\n", req)
	a, b := req.GetA(), req.GetB()
	sum := a + b
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (rpc *Rpc) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Prime Number Decomposition function was invoked with %v\n", req)

	n := req.GetInputNumber()
	var divider int32 = 2

	for n > 1 {
		if n%divider == 0 {
			res := &calculatorpb.PrimeNumberDecompositionResponse{
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

func (rpc *Rpc) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with a client stream mechanism.\n")
	counter := 0
	sumTotal := int64(0)
	for {
		fmt.Printf("count=%v, sumTotal=%v\n", counter, sumTotal)
		req, err := stream.Recv()

		if err == io.EOF { // client stream fully read
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
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
