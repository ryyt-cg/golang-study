package main

import (
	"gitlab.con/aionx/go-examples/grpc-in-go/greet/greet-server/rpc/greet"
	greetpb "gitlab.con/aionx/go-examples/grpc-in-go/pb-domain/greet"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var log *zap.Logger

func main() {
	log, _ = zap.NewProduction()
	log.Info("Greet server is a-go!")

	// Create TCP listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen", zap.Error(err))
	}

	// create grpc and register service
	s := grpc.NewServer()
	greetServer := greet.NewGreetRpc()
	greetpb.RegisterGreetServiceServer(s, greetServer)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve?!.", zap.Error(err))
	}
}
