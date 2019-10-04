package main

import (
	"context"
	"log"
	"net"

	"github.com/lcullen/mardown/language/go/pb"

	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAddressBookServiceServer(s, &Reply{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Reply struct {
}

func (*Reply) GetAll(context.Context, *google_protobuf1.Empty) (*pb.AddressBook, error) {
	return &pb.AddressBook{
		People: []*pb.Person{
			{},
		},
	}, nil
}
