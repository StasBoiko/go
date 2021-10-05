package main

import (
	"log"
	"net"
	"work2/models/proto"

	"google.golang.org/grpc"
)

//func main() {
//	fmt.Println("Go gRPC Beginners Tutorial!")
//
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	s := chat.Server{}
//
//	grpcServer := grpc.NewServer()
//
//	chat.RegisterChatServiceServer(grpcServer, &s)
//
//	if err := grpcServer.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %s", err)
//	}
//}

func main() {
	s := grpc.NewServer()
	server := &proto.GRPCserver{}
	proto.RegisterTaskServer(s, server)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

	//lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8081))
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//authMD := middlewares.AuthMD{}
	//opts := make([]grpc.ServerOption, 0)
	//opts = append(opts, grpc.ChainUnaryInterceptor(authMD.UnaryInterceptor()))
	//grpcServer := grpc.NewServer(opts...)
	//chatHandler := proto.GRPCserver{}
	//// registering specific handlers for this server
	//proto.RegisterTaskServer(grpcServer, &chatHandler)
	//log.Println("starting server")
	//
	//if err := grpcServer.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %s", err)
	//}
}
