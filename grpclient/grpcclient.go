package main

import (
	"log"
	"time"

	pb "github.com/brighty1982/GoGRPCTest/proto"
	"github.com/golang/protobuf/ptypes"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMeterReadServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	meteread := new(pb.MeterRead)
	meteread.SerialNo = "SG628162H"
	meteread.Date, _ = ptypes.TimestampProto(time.Now().UTC())
	meteread.Reg1Read = 12745
	meteread.Reg2Read = 2516

	r, err := c.SubmitMeterRead(ctx, meteread)
	if err != nil {
		log.Fatalf("error submitting meter read %v", err)
	} else {
		log.Printf("Response: %s", r.Message)
	}

}
