package main

import (
	"flag"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	l "log"
	"net/http"
	"os"
	"os/signal"
	"todo-grpc/client/handler"
	"todo-grpc/pb"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		l.Println("Error getting env, not comming through ", err)
	}
}

func main() {
	BIND_ADDR := os.Getenv("BIND_ADDR_SERVER")

	address := flag.String("server", ":"+BIND_ADDR, "gRPC server in format host:port")
	flag.Parse()

	var logger log.Logger

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		level.Error(logger).Log("Can't Connect to ", err.Error())
	}
	defer conn.Close()

	connUser := pb.NewUserServiceClient(conn)
	connTodo := pb.NewToDoServiceClient(conn)

	handlerUser := handler.NewUserServiceClient(connUser)
	handlerTodo := handler.NewTodoServiceClient(connTodo)

	r := mux.NewRouter()

	// Method GET
	getR := r.Methods(http.MethodGet).Subrouter()
	// User
	getR.HandleFunc("/get/users", handlerUser.GetAllUser)
	getR.HandleFunc("/get/user", handlerUser.GetUser)
	getR.HandleFunc("/delete/user", handlerUser.DeleteUser)

	// Todo
	getR.HandleFunc("/get/todos", handlerTodo.GetAllTodos)
	getR.HandleFunc("/get/todo", handlerTodo.GetTodo)
	getR.HandleFunc("/delete/todo", handlerTodo.DeleteTodo)

	postR := r.Methods(http.MethodPost).Subrouter()
	// User
	postR.HandleFunc("/create/user", handlerUser.CreateUser)
	postR.HandleFunc("/update/user", handlerUser.UpdateUser)

	// Todo
	postR.HandleFunc("/create/todo", handlerTodo.CreateTodo)
	postR.HandleFunc("/update/todo", handlerTodo.UpdateTodo)

	srv := &http.Server{
		Handler:      handlers.CORS()(r),
		Addr:         os.Getenv("BIND_ADDR_CLIENT"),
	}

	go func() {
		l.Println("Starting server on port", os.Getenv("BIND_ADDR_CLIENT"))

		err = srv.ListenAndServe()
		if err != nil {
			l.Println("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	sig := <-ch
	l.Println("Got Signal:", sig)

	os.Exit(1)
}
