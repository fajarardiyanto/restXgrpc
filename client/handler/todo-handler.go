package handler

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
	"todo-grpc/auth"
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

	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.Respond(w, http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": "Unauthorization",
			"data":    "",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req.Id = uuid.New().String()
	req.Title = r.FormValue("title")
	req.Description = r.FormValue("description")
	req.AuthorId = userID
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

	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.Respond(w, http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": "Unauthorization",
			"data":    "",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req.Id = id
	req.AuthorId = userID
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

	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.Respond(w, http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": "Unauthorization",
			"data":    "",
		})
		return
	}

	request := pb.ReadAllRequest{
		AuthorId: userID,
	}
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

	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.Respond(w, http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": "Unauthorization",
			"data":    "",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := pb.ReadRequest{
		Id:       id,
		AuthorId: userID,
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

	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.Respond(w, http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": "Unauthorization",
			"data":    "",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := pb.DeleteRequest{
		Id: id,
		AuthorId: userID,
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
