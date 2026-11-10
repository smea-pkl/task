package main

import (
	"context"
	"fmt"
	"net"

	usersv1 "pkl/backend/gen/pkl/users/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserServer struct {
	usersv1.UnimplementedUserServiceServer
}

func (s *UserServer) Login(
	ctx context.Context,
	req *usersv1.LoginRequest,
) (*usersv1.LoginResponse, error) {
	username := req.Username
	password := req.Password

	// Cek kredensial: username "heri" dan password "heri7777"
	if username == "heri" && password == "heri7777" {
		return &usersv1.LoginResponse{
			Token: "123456",
		}, nil
	}

	// Jika username/password tidak cocok, kembalikan token "0"
	return &usersv1.LoginResponse{
		Token: "0",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	usersv1.RegisterUserServiceServer(
		grpcServer,
		&UserServer{},
	)

	reflection.Register(grpcServer)

	fmt.Println("Jalan")

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
