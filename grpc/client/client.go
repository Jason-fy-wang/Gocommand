package main

import (
	"context"
	"fmt"

	"com.learn/command/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// create grpc connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()

	// client 
	client := pb.NewUserServiceClient(conn)

	// call GetUserInfo method
	resp, err := client.GetUserInfo(context.Background(), &pb.UserRequest{
		UserId: 12345,
		Token: "sometoken",
	})
	if err != nil {
		fmt.Printf("Error calling GetUserInfo: %v\n", err)
		return
	}
	
	fmt.Printf("Received response: UserId=%d, UserName=%s, Msg=%s, Age=%d\n", resp.GetUserId(), resp.GetUsername(), resp.GetMsg(), resp.GetAge())
}