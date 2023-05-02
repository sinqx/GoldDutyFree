package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	db "sg/database"
	pb "sg/proto"
	"sg/service"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
}
type UserServer struct {
	pb.UnimplementedUserServiceServer
}

var database = db.Connect().Debug()

const (
	grpcPort = ":9000"
	httpPort = ":8080"
)

func runHTTPServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts)
	if err != nil {
		return err
	}

	err = pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts)
	if err != nil {
		return err
	}

	http.Handle("/", mux)
	http.Handle("/gold", mux)

	log.Printf("HTTP server listening on %s", httpPort)
	return http.ListenAndServe(httpPort, nil)
}

func runGRPCServer() error {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &OrderServer{})
	pb.RegisterUserServiceServer(s, &UserServer{})
	log.Printf("gRPC server listening on %s", grpcPort)
	return s.Serve(lis)
}

func main() {
	go func() {
		err := runGRPCServer()
		if err != nil {
			log.Fatalf("Failed to run gRPC server: %v", err)
		}
	}()

	err := runHTTPServer()
	if err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}
}

func (s *OrderServer) CreateOrder(_ context.Context, OrderModel *pb.OrderModel) (*pb.Order, error) {
	return service.CreateOrder(OrderModel), nil
}
