package client

import (
	"log"

	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	"google.golang.org/grpc"
)

func SetupClient() (applicationpb.ApplicationClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:10000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return applicationpb.NewApplicationClient(conn), conn
}
