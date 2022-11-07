package main

import (
	"context"
	"fmt"
	"log"
	"session-3/common/config"
	"session-3/common/models"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func serviceUser() models.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	return models.NewUsersClient(conn)
}

func main() {
	userSvc := serviceUser()
	fmt.Println("user test")
	var user1 = &models.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "f0r Th3 H0rD3",
		Gender:   models.UserGender_FEMALE,
	}
	userSvc.Register(context.Background(), user1)

	var user2 = &models.User{
		Id:       "u002",
		Name:     "Windrunner",
		Password: "f0r H0rD3",
		Gender:   models.UserGender_MALE,
	}

	userSvc.Register(context.Background(), user2)

	users, _ := userSvc.List(context.Background(), new(empty.Empty))
	log.Printf("List Users %+v\n", users.GetList())
}
