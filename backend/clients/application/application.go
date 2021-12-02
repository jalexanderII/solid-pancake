package application

import (
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	"google.golang.org/grpc"
)

func NewApplClient() applicationpb.ApplicationClient {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return applicationpb.NewApplicationClient(conn)
}


