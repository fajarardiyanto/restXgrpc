package handler

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
	"todo-grpc/client/utils"
	"todo-grpc/pb"
)

type todoServiceClient struct {
	c pb.ToDoServiceClient
}

func NewTodoServiceClient(c pb.ToDoServiceClient) *todoServiceClient {
	return &todoServiceClient{c}
}

func (todo *todoServiceClient) CreateTodo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var req pb.ToDo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req.Id = uuid.New().String()
	req.Title = r.FormValue("title")
	req.Description = r.FormValue("description")
	req.CreatedAt = time.Now().Unix()

	request := pb.CreateRequest{
		Todo: &req,
	}

	res, err := todo.c.Create(ctx, &request)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	utils.Respond(w, http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "",
		"data":    res.Todo,
	})
}

func (todo *todoServiceClient) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	r.ParseForm()
	var req pb.ToDo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req.Id = id
	req.Title = r.FormValue("title")
	req.Description = r.FormValue("description")
	req.CreatedAt = time.Now().Unix()

	request := pb.UpdateRequest{
		Todo: &req,
	}

	res, err := todo.c.Update(ctx, &request)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	utils.Respond(w, http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "",
		"data":    res.Todo,
	})
}

func (todo *todoServiceClient) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := pb.ReadAllRequest{}
	res, err := todo.c.ReadAll(ctx, &request)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	utils.Respond(w, http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "",
		"data":    res.Todo,
	})
}

func (todo *todoServiceClient) GetTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := pb.ReadRequest{
		Id: id,
	}

	res, err := todo.c.Read(ctx, &request)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Todos not found",
			"data":    "",
		})
		return
	}

	utils.Respond(w, http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "",
		"data":    res.Todo,
	})
}

func (todo *todoServiceClient) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := pb.DeleteRequest{
		Id: id,
	}

	res, err := todo.c.Delete(ctx, &request)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "Todos not found",
			"data":    "",
		})
		return
	}

	utils.Respond(w, http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "",
		"data":    res.Id,
	})
}