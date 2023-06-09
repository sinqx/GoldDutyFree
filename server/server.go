package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	pb "sg/proto"
)

const (
	grpcPort string = ":9000"
	httpPort string = ":8080"
)

func runHTTPServer() error {
	ctx := context.Background()
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

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(httpPort),
		Handler: mux,
	}
	log.Printf("HTTP server listening on %s", httpPort)
	return gwServer.ListenAndServe()
}

func runGRPCServer() error {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &pb.UnimplementedOrderServiceServer{})
	pb.RegisterUserServiceServer(s, &pb.UnimplementedUserServiceServer{})

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
