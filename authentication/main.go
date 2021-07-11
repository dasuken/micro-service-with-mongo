package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"microservices/authentication/repository"
	"microservices/authentication/service"
	"microservices/db"
	"microservices/pb"
	"net"
)

var (
	local bool
	port int
)

func init() {
	flag.IntVar(&port, "port", 9001, "authentication service port")
	flag.BoolVar(&local, "local", true, "run service local")
	flag.Parse()
}

func main() {
	if local { 
		// load environment values
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	repo := repository.NewUsersRepository(conn)
	authService := service.NewAuthService(repo)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)

	log.Printf("Authentication service runnning on [::]:%d\n", port)

	grpcServer.Serve(lis)
}