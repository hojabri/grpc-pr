package main

import (
	"context"
	"fmt"
	"github.com/hojabri/grpc-pr/greet/greetpb"
	"google.golang.org/grpc"
	"io"
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
	doServerStreaming(c)

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
func doServerStreaming(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to greet many times!")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Omid",
			LastName:  "Hojabri",
		},
	}
	res, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet many times! :%v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			//We have reached to the end of stream
			log.Println("---End of receiving the stream---")
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive message :%v", err)
		}

		log.Printf("Response from the server: %v\n", msg.GetResult())
	}
}
