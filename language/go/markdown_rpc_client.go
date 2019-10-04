package main

import (
	"context"
	"fmt"

	"github.com/lcullen/mardown/language/go/pb"

	"github.com/golang/protobuf/ptypes/empty"

	grpc "google.golang.org/grpc"
)

const address = "127.0.0.1:5001"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	client := pb.NewAddressBookServiceClient(conn)
	ret, err := client.GetAll(context.Background(), &empty.Empty{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(ret)
}
