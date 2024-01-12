package greet

import (
	"context"
	greetpb "gitlab.con/aionx/go-examples/grpc-in-go/pb-domain/greet"
	"go.uber.org/zap"
	"io"
	"log"
	"strconv"
	"time"
)

type Rpc struct {
	log *zap.Logger
	//orderService Servicer
}

func NewGreetRpc() *Rpc {
	return &Rpc{zap.L()}
}

func (rpc *Rpc) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	rpc.log.Info("Greet function was invoked with.", zap.String("request", req.String()))
	firstName := req.GetGreeting().GetFirstName()
	result := "Greetings " + firstName
	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func (rpc *Rpc) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	rpc.log.Info("GreetManyTimes function was invoked with.", zap.String("request", req.String()))
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Bonjour " + firstName + " " + strconv.Itoa(i) + " fois!"
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}

		err := stream.Send(res)
		if err != nil {
			rpc.log.Error("error while sending stream.", zap.Error(err))
		}

		time.Sleep(600 * time.Millisecond)
	}
	return nil
}

func (rpc *Rpc) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	rpc.log.Info("LongGreet function was invoked with a stream thing.")
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

func (rpc *Rpc) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	rpc.log.Info("GreetEveryone function was invoked with a stream request.")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal("Error while reading client stream", zap.Error(err))
			return err
		}

		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})

		if err != nil {
			log.Fatal("Error while sending data to client stream.", zap.Error(err))
			return err
		}
	}
}
