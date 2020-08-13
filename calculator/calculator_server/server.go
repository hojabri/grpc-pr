package main

import (
	"context"
	"fmt"
	"github.com/hojabri/grpc-pr/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function is calling from client: %v \n", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	result := firstNumber + secondNumber

	res := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return res, nil
}

func main() {

	fmt.Println("Calculator Server is running...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
