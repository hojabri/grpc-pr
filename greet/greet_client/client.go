package main

import (
	"context"
	"fmt"
	"github.com/hojabri/grpc-pr/greet/greetpb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Starting client...")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to greet!")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Omid",
			LastName:  "Hojabri",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet! :%v", err)
	}
	log.Printf("Response is: %v", res.Result)
}
