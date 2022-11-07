package main

import (
	"context"
	"log"
	"net"
	"session-3/common/config"
	"session-3/common/models"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var localStorage *models.UserList

func init() {
	localStorage = new(models.UserList)
	localStorage.List = make([]*models.User, 0)
}

type UsersServer struct{}

func (UsersServer) Register(ctx context.Context, param *models.User) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)
	log.Println("Registering User - ", param.String())
	return new(empty.Empty), nil
}

func (UsersServer) List(ctx context.Context, void *empty.Empty) (*models.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	userSrv := UsersServer{}
	models.RegisterUsersServer(srv, userSrv)

	log.Println("Starting User Server at ", config.SERVICE_USER_PORT)

	l, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
