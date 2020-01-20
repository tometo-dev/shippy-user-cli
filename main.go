package main

import (
	"context"
	micro "github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	pb "github.com/tsuki42/shippy-user-service/proto/auth"
	"log"
	"os"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags
	srv.Init()

	client := pb.NewAuthService("go.micro.srv.user", microclient.DefaultClient)

	name := "Sudhanshu Ranjan"
	email := "perfectsudh@gmail.com"
	password := "password"
	company := "NASA"

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Company:  company,
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Count not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// just exit
	os.Exit(0)

}
