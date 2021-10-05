package proto

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCserver struct{}

func (s *GRPCserver) GetTasks(ctx context.Context, req *GetTasksRequest) (*GetTasksResponse, error) {
	return &GetTasksResponse{
		Id: 1, Name: "Vasya",
		Event:       "some event name",
		Date:        timestamppb.New(time.Now().UTC()),
		Description: "some description",
		UserId:      "1",
		UpdatedAt:   timestamppb.New(time.Now().UTC()),
		CreatedAt:   timestamppb.New(time.Now().UTC()),
		DeletedAt:   timestamppb.New(time.Now().UTC()),
	}, nil
}

func (s *GRPCserver) CreateTask(context.Context, *CreateRequest) (*CreateResponse, error) {
	return &CreateResponse{
		Id: 1, Name: "Petya",
		Event:       "some event name2",
		Date:        timestamppb.New(time.Now().UTC()),
		Description: "some description2",
		UserId:      "2",
		UpdatedAt:   timestamppb.New(time.Now().UTC()),
		CreatedAt:   timestamppb.New(time.Now().UTC()),
		DeletedAt:   timestamppb.New(time.Now().UTC()),
	}, nil
}

//func (s *GRPCserver) GetTasks(ctx context.Context, req *GetTasksRequest) (*GetTasksResponse, error) {
//	return &GetTasksResponse{Id: 1, Name: "Vasya", Event: "some event name", Description: "some description", UserId: "1"}, nil
//}
