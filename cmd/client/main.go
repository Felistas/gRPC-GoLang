package main

import (
	"github.com/felistas/gRPC-Golang-REST-API/pkg/api/v1"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:7000"

	conn, e := gRPC.Dial(serverAddress, gRPC.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := v1.ToDoServiceClient(conn)


	//for i := range [10]int{} {

		x := v1.ToDo{
			id:        int64(1),
			title: "hele",
			Description:      string("Grpc-Demo"),

		}
		ToDoModel :=v1.AddRequest{x}

		if responseMessage, e :=client.Create(context.Background(), &ToDoModel); e != nil {
			panic(fmt.Sprintf("Was not able to insert Record %v", e))
		} else {
			fmt.Println("Record Inserted..")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	//}
}