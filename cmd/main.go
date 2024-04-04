package main

import (
	"net"
	"os"

	"github.com/dkshi/shortener/internal/repository"
	"github.com/dkshi/shortener/internal/serverapi"
	"github.com/dkshi/shortener/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env: %s", err.Error())
	}

	repo := repository.NewRepository()
	repo.StartGC()
	
	service := service.NewService(repo)

	serverAPI := serverapi.NewServerAPI(service)
	grpcServer := grpc.NewServer()
	serverAPI.RegisterServer(grpcServer)

	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		logrus.Fatalf("error listening to tcp: %s", err.Error())
	}

	if err = grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("error serving server: %s", err.Error())
	}
}
