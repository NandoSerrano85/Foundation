package main

import (
	"context"
	"fmt"
	"log"
	"net"

	user "backend/proto/user"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) ListAllUsers(ctx context.Context, request *user.Request) (*user.Response, error) {
	fmt.Println(request)

	result := "this is a test for all users"

	return &user.Response{Result: result}, nil
}

func (s *server) NewUser(ctx context.Context, request *user.Request) (*user.Response, error) {
	fmt.Println(request)

	result := "this is a test for new users"

	return &user.Response{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Listening...")
	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
