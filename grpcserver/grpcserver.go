package main

import (
	"context"
	"log"
	"net"

	pb "github.com/brighty1982/GoGRPCTest/proto"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SubmitMeterRead(ctx context.Context, in *pb.MeterRead) (*pb.MeterReadReply, error) {

	//logic here
	//lookup and validate content against meter
	//store in db
	//validate actual read.

	readDate, _ := ptypes.Timestamp(in.Date)

	log.Printf("Serial Number: %s", in.SerialNo)
	log.Printf("Timestamp: %s", readDate)
	log.Printf("Reg1: %d", in.Reg1Read)
	log.Printf("Reg2: %d", in.Reg2Read)

	// return message
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
