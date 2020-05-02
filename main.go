package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	pb "github.com/soypita/shippy-service-user/proto/user"
)

func main() {

	db, err := CreateConnection()
	defer db.Close()
	log.Println("Success establish connection to database")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(User{})

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
	)

	// Init will parse the command line flags.
	srv.Init()

	pubsub := micro.NewPublisher("user.created", srv.Client())

	repository := &UserRepository{db}
	tokenService := TokenService{repository}
	// Register handler
	h := &service{repository, tokenService, pubsub}

	pb.RegisterUserServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
