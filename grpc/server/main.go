package main

import (
	"context"
	"fmt"
	"github.com/yashwant-shrivastava/tech-session/grpc/customer_rpc"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type CustomerAPIServer struct {
	customer_rpc.UnimplementedCustomerAPIServer
}

func (c CustomerAPIServer) GetCustomerById(ctx context.Context, request *customer_rpc.GetCustomerByIdRequest) (*customer_rpc.GetCustomerByIdResponse, error) {
	return &customer_rpc.GetCustomerByIdResponse{
		Status: true,
		Id:     request.GetId(),
	}, nil
}

func (c CustomerAPIServer) GetCustomerByIds(stream customer_rpc.CustomerAPI_GetCustomerByIdsServer) error {
	var ids []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&customer_rpc.GetCustomerByIdsResponse{Status: true, Ids: ids})
		}
		if err != nil {
			return err
		}
		ids = append(ids, req.GetId())
	}
}

func (c CustomerAPIServer) GetCustomerInterests(request *customer_rpc.GetCustomerInterestsRequest, server customer_rpc.CustomerAPI_GetCustomerInterestsServer) error {
	for i := 0; i < 5; i++ {
		if err := server.Send(&customer_rpc.GetCustomerInterestsResponse{Interest: fmt.Sprint(i)}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (c CustomerAPIServer) GetCustomerInterestsByIds(server customer_rpc.CustomerAPI_GetCustomerInterestsByIdsServer) error {
	for {
		req, err := server.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		if err := server.Send(&customer_rpc.GetCustomerInterestsResponse{Interest: fmt.Sprint(req.GetId())}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	customer_rpc.RegisterCustomerAPIServer(s, &CustomerAPIServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
}
