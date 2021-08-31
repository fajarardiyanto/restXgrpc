package server

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"todo-grpc/pb"
)

func RunServerGRPC(ctx context.Context, api pb.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterToDoServiceServer(server, api)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down server .....")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Println("starting server .....")
	return server.Serve(listen)
}
