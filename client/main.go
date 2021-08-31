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
	connAuth := pb.NewAuthServiceClient(conn)

	handlerUser := handler.NewUserServiceClient(connUser)
	handlerTodo := handler.NewTodoServiceClient(connTodo)
	handlerAuth := handler.NewAuthServiceClient(connAuth)

	r := mux.NewRouter()

	// Auth
	r.HandleFunc("/login", handlerAuth.Login).Methods("POST")

	// User
	r.HandleFunc("/get/user", handlerUser.GetUser).Methods("GET")
	r.HandleFunc("/get/users", handlerUser.GetAllUser).Methods("GET")
	r.HandleFunc("/delete/user", handlerUser.DeleteUser).Methods("GET")
	r.HandleFunc("/create/user", handlerUser.CreateUser).Methods("POST")
	r.HandleFunc("/update/user", handlerUser.UpdateUser).Methods("POST")

	// Todo
	r.HandleFunc("/get/todo", handlerTodo.GetTodo).Methods("GET")
	r.HandleFunc("/get/todos", handlerTodo.GetAllTodos).Methods("GET")
	r.HandleFunc("/delete/todo", handlerTodo.DeleteTodo).Methods("GET")
	r.HandleFunc("/create/todo", handlerTodo.CreateTodo).Methods("POST")
	r.HandleFunc("/update/todo", handlerTodo.UpdateTodo).Methods("POST")

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
