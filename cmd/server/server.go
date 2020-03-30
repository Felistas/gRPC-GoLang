package main

import (
	"github.com/felistas/gRPC-Golang-REST-API/pkg/api/v1"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(7000)
	gRPCServer := grpc.NewServer()


	repositoryServiceImpl := v1.NewToDoServiceClient()
	v1.RegisterToDoServiceServer(gRPCServer, repositoryServiceImpl)

	// start the server
	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
