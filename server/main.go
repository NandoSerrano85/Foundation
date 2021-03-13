package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"

	user "github.com/NandoSerrano85/Foundation/api/user"
)

type server struct {
	user.UnimplementedUserServer
	allUsers []*user.UserStruct

	mu sync.Mutex
}

func (s *server) ListAllUsers(empty *user.Empty, allUsers user.User_ListAllUsersServer) error {
	fmt.Println(empty)

	return nil
}

func (s *server) NewUser(ctx context.Context, request *user.UserStruct) (*user.Empty, error) {
	log.Printf("Creating new user with name: %s", request.FirstName)

	return &user.Empty{}, nil
}

func newServer() *server {
	s := &server{}
	return s
}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Listening...")
	s := grpc.NewServer()
	user.RegisterUserServer(s, newServer())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
