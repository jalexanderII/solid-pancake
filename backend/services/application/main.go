package main

import (
	"log"
	"net"

	"github.com/hashicorp/go-hclog"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	"github.com/jalexanderII/solid-pancake/services/application/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var applicationAddr = "localhost:9093"

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	grpcServer, lis := setupApplicationServer()
	// start service's server
	log.Println("starting application rpc service on", applicationAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func setupApplicationServer() (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", applicationAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Set options, here we can configure things like TLS support
	var opts []grpc.ServerOption
	// setup and register currency service
	// create a new gRPC server, use WithInsecure to allow http connections and (blank) options
	grpcServer := grpc.NewServer(opts...)

	// Register the service with the server
	applicationpb.RegisterApplicationServer(grpcServer, server.NewApplicationServer(hclog.Default()))

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	return grpcServer, lis
}
