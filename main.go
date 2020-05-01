package main

import (
	"fmt"

	"github.com/micro/go-micro"
	pb "github.com/soypita/shippy-service-user/proto/user"
)

func main() {

	db, err := CreateConnection()
	defer db.Close()

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

	repository := &UserRepository{db}
	tokenService := TokenService{repository}
	// Register handler
	h := &service{repository, tokenService}

	pb.RegisterUserServiceHandler(srv.Server(), h)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
