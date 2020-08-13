package main

import (
	"context"
	"fmt"
	"github.com/hojabri/grpc-pr/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Starting calculator client...")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := calculatorpb.NewCalculatorServiceClient(conn)

	doUnary(c, 35, 67)

}

func doUnary(c calculatorpb.CalculatorServiceClient, firstNumber int32, secondNumber int32) {

	fmt.Println("Starting to Sum!")
	req := &calculatorpb.SumRequest{
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
	}
	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet! :%v", err)
	}
	log.Printf("The sum of %d and %d is: %v", firstNumber, secondNumber, res.SumResult)
}
