package impl

import (
	"github.com/felistas/gRPC-Golang-REST-API/pkg/api/v1"
	"context"
)

type ToDoService struct {
}

func (t ToDoService) Create(context.Context, *v1.AddRequest) (*v1.AddResponse, error) {
	panic("implement me")
}

func (t ToDoService) Read(context.Context, *v1.ReadRequest) (*v1.ReadResponse, error) {
	panic("implement me")
}

func (t ToDoService) Update(context.Context, *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	panic("implement me")
}

func (t ToDoService) Delete(context.Context, *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	panic("implement me")
}

func (t ToDoService) ReadAll(context.Context, *v1.ReadAllRequest) (*v1.ReadAllResponse, error) {
	panic("implement me")
}

//NewToDoServiceImpl returns the pointer to the implementation.
func NewToDoServiceGrpcImpl() v1.ToDoServiceServer{
	return &ToDoService{}
}


