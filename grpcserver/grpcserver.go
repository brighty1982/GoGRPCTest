package main

import (
	"context"
	"log"
	"net"

	pb "github.com/brighty1982/GoGRPCTest/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SubmitMeterRead(ctx context.Context, in *pb.MeterRead) (*pb.MeterReadReply, error) {

	///logic here

	return &pb.MeterReadReply{Message: "valid read for " + in.SerialNo}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterMeterReadServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
