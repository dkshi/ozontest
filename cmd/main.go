package main

import (
	"net"
	"os"
	"time"

	"github.com/dkshi/shortener/internal/repository"
	"github.com/dkshi/shortener/internal/serverapi"
	"github.com/dkshi/shortener/internal/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	logrus.Fatalf("error loading .env: %s", err.Error())
	// }

	if err := loadTimeLocation(os.Getenv("TIME_LOCATION")); err != nil {
		logrus.Fatalf("error loading time location: %s", err.Error())
	}

	repo := chooseStorage(os.Getenv("STORAGE"))
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

// Loads time zone, UTC prefered because it's postgres default timezone
func loadTimeLocation(location string) error {
	utcLocation, err := time.LoadLocation(location)
	if err != nil {
		return err
	}
	time.Local = utcLocation
	return nil
}

// Chose storage to work with
func chooseStorage(storage string) *repository.Repository {
	if storage == "postgres" {

		db, err := repository.NewPostgresDB(repository.Config{
			Host:     os.Getenv("DB_HOST"),
			DBName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		})

		if err != nil {
			logrus.Fatalf("error connecting to the database: %s", err.Error())
		}

		return repository.NewRepository(db)
	}

	return repository.NewRepository(nil)
}
