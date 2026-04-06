package main

import (
	"context"
	"fmt"
	"net"

	"com.learn/command/grpc/pb"
	"google.golang.org/grpc"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServiceServer) GetUserInfo(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	fmt.Printf("Received request for User Id: %d with Token: %s\n", req.GetUserId(), req.GetToken())


	return &pb.UserResponse{
		UserId:   req.GetUserId(),
		Username: fmt.Sprintf("User%d", req.GetUserId()),
		Msg:    fmt.Sprintf("user%d@example.com", req.GetUserId()),
		Age: 10,
	},nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}

	// create grpc server
	grpcServer := grpc.NewServer()
	
	// register service
	pb.RegisterUserServiceServer(grpcServer, &userServiceServer{})
	fmt.Println("gRPC server is running on port 50051...")

	// start server
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}