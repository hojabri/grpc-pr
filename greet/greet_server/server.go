package main

import (
	"context"
	"fmt"
	"github.com/hojabri/grpc-pr/greet/greetpb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function is calling from client: %v \n", req)
	firstname := req.GetGreeting().FirstName
	lastname := req.GetGreeting().LastName

	result := "Hello " + firstname + " " + lastname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function is calling from client: %v \n", req)
	firstName := req.Greeting.FirstName
	lastName := req.Greeting.LastName

	for i := 0; i < 10; i++ {
		res := greetpb.GreetManyTimesResponse{
			Result: "Hello " + firstName + " " + lastName + " number:" + strconv.Itoa(i),
		}
		stream.Send(&res)
		time.Sleep(1 * time.Second)
	}

	return nil

}

func main() {

	fmt.Println("Server is running...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
