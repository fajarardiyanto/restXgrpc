package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	l "log"
	"os"
	"todo-grpc/cmd/server"
	"todo-grpc/seed"
	"todo-grpc/service"

	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		l.Println("Error getting env, not comming through ", err)
	}
}

func RunServer() error {
	ctx := context.Background()

	BIND_ADDR := os.Getenv("BIND_ADDR_SERVER")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "users",
			"time: ", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started at port "+BIND_ADDR)
	defer level.Info(logger).Log("msg", "service ended")

	DB_URI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DbHost"), os.Getenv("DbPort"), os.Getenv("DbUser"), os.Getenv("DbName"), os.Getenv("DbPassword"))
	var db *gorm.DB
	{
		var err error

		db, err = gorm.Open("postgres", DB_URI)
		if err != nil {
			level.Error(logger).Log("Exit", err.Error())
			os.Exit(-1)
		}

		defer db.Close()
	}

	seed.LoadSeed(db)
	api := service.NewRepoServer(db)

	return server.RunServerGRPC(ctx, api, BIND_ADDR)
}

func main() {
	if err := RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
