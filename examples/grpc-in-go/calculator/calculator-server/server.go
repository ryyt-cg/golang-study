package main

import (
	"fmt"
	"gitlab.con/aionx/go-examples/grpc-in-go/calculator/calculator-server/rpc/calculator"
	calculatorpb "gitlab.con/aionx/go-examples/grpc-in-go/pb-domain/calculator"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Calc server is a-go!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorRpc := calculator.NewCalculatorRpc()
	calculatorpb.RegisterCalculatorServiceServer(s, calculatorRpc)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve?! %v", err)
	}
}
