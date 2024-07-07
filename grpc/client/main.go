package main

import (
	"context"
	"fmt"
	"github.com/yashwant-shrivastava/tech-session/grpc/customer_rpc"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := customer_rpc.NewCustomerAPIClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//testUnaryCall(c, ctx)
	//testClientStreamingCall(c, ctx)
	//testServerStreamingCall(c, ctx)
	testBirectionalStreamingCall(c, ctx)
}

func testUnaryCall(c customer_rpc.CustomerAPIClient, ctx context.Context) {
	res, err := c.GetCustomerById(ctx, &customer_rpc.GetCustomerByIdRequest{Id: "1234"})
	if err != nil {
		log.Fatalf("Could not call: %v", err)
	}
	log.Printf("Response: %+v", res)
}

func testClientStreamingCall(c customer_rpc.CustomerAPIClient, ctx context.Context) {
	stream, err := c.GetCustomerByIds(ctx)
	if err != nil {
		log.Fatalf("Could not call: %v", err)
	}

	for i := 0; i < 10; i++ {
		err := stream.Send(&customer_rpc.GetCustomerByIdRequest{Id: fmt.Sprint(i)})
		if err != nil {
			log.Fatalf("Could not send: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not receive: %v", err)
	}
	log.Printf("Response: %+v", res)
}

func testServerStreamingCall(c customer_rpc.CustomerAPIClient, ctx context.Context) {
	stream, err := c.GetCustomerInterests(ctx, &customer_rpc.GetCustomerInterestsRequest{Id: "1234"})
	if err != nil {
		log.Fatalf("Could not call: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("Could not receive: %v", err)
		}
		log.Printf("Response: %+v", res)
	}
}

func testBirectionalStreamingCall(c customer_rpc.CustomerAPIClient, ctx context.Context) {
	stream, err := c.GetCustomerInterestsByIds(ctx)
	if err != nil {
		log.Fatalf("Could not call: %v", err)
	}

	ch := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			err := stream.Send(&customer_rpc.GetCustomerByIdRequest{Id: fmt.Sprint(i)})
			if err != nil {
				log.Fatalf("Could not send: %v", err)
			}
			time.Sleep(500 * time.Millisecond)
		}

		// Close the send direction of the stream
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("Failed to close send stream: %v", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// Server has finished sending messages
				close(ch)
				return
			}
			if err != nil {
				log.Fatalf("Could not receive: %v", err)
			}

			log.Printf("Response: %+v", res)
		}
	}()
	<-ch
}
