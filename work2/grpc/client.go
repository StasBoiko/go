package main

import (
	"context"
	"log"
	"work2/models/proto"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewTaskClient(conn)

	ctx := metadata.NewOutgoingContext(context.TODO(), metadata.New(map[string]string{
		"Authorization": "Bearer cm9tYW46cHdk",
	}))
	response, err := c.GetTasks(ctx, &proto.GetTasksRequest{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)
}

//func main() {
//	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	c := proto.NewTaskClient(conn)
//	res, err := c.GetTasks(context.TODO(), &proto.GetTasksRequest{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println(res.GetName(), res.GetEvent(), res.GetDescription())
//
//	cRes, err := c.CreateTask(context.TODO(), &proto.CreateRequest{
//		Name:        "Petya",
//		Event:       "some event3",
//		Date:        timestamppb.New(time.Now().UTC()),
//		Description: "some description",
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println(cRes.GetName(), cRes.GetEvent(), cRes.GetDescription())
//}
