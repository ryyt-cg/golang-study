package main

import (
	"context"
	greetpb "gitlab.con/aionx/go-examples/grpc-in-go/pb-domain/greet"
	"go.uber.org/zap"
	"io"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

var log *zap.Logger

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Info("Greet function was invoked with.", zap.String("request", req.String()))
	firstName := req.GetGreeting().GetFirstName()
	result := "Greetings " + firstName
	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Info("GreetManyTimes function was invoked with.", zap.String("request", req.String()))
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Bonjour " + firstName + " " + strconv.Itoa(i) + " fois!"
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(600 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log, _ = zap.NewProduction()
	log.Info("LongGreet function was invoked with a stream thing.")
	result := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF { // client stream fully read
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatal("Error while reading client stream", zap.Error(err))
		}

		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "!\n"
	}
}

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
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve?!.", zap.Error(err))
	}
}
