package main

import (
	"github.com/felistas/gRPC-Golang-REST-API/pkg/service/v1"
	"fmt"
	"log"
	"net"

	"google.golang.org/gRPC"
)

func main() {
	netListener := getNetListener(7000)
	gRPCServer := gRPC.NewServer()


	repositoryServiceImpl := v1.NewToDoServiceServer()
	service.RegisterRepositoryServiceServer(gRPCServer, repositoryServiceImpl)

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
