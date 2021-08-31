package handler

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
	"todo-grpc/client/utils"
	"todo-grpc/pb"
)

type userServiceClient struct {
	c pb.UserServiceClient
}

func NewUserServiceClient(c pb.UserServiceClient) *userServiceClient {
	return &userServiceClient{c}
}

func (user *userServiceClient) CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var req pb.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	password := r.FormValue("password")
	newPassword, err := utils.Hash(password)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	req.Id = uuid.New().String()
	req.Fullname = r.FormValue("fullname")
	req.Username = r.FormValue("username")
	req.Email = r.FormValue("email")
	req.Password = string(newPassword)
	req.CreatedAt = time.Now().Unix()

	request := pb.UserCreateRequest{
		User: &req,
	}

	res, err := user.c.Create(ctx, &request)
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
		"data":    res.User,
	})
}

func (user *userServiceClient) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	r.ParseForm()
	var req pb.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	password := r.FormValue("password")
	newPassword, err := utils.Hash(password)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	req.Id = id
	req.Fullname = r.FormValue("fullname")
	req.Username = r.FormValue("username")
	req.Email = r.FormValue("email")
	req.Password = string(newPassword)
	req.CreatedAt = time.Now().Unix()

	request := pb.UserUpdateRequest{
		User: &req,
	}

	res, err := user.c.Update(ctx, &request)
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
		"data":    res.User,
	})
}

func (user *userServiceClient) GetAllUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := pb.UserReadAllRequest{}
	res, err := user.c.ReadAll(ctx, &request)
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
		"data":    res.User,
	})
}

func (user *userServiceClient) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := pb.UserReadRequest{
		Id: id,
	}

	res, err := user.c.Read(ctx, &request)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "User not found",
			"data":    "",
		})
		return
	}

	utils.Respond(w, http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "",
		"data":    res.User,
	})
}

func (user *userServiceClient) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	request := pb.UserDeleteRequest{
		Id: id,
	}

	res, err := user.c.Delete(ctx, &request)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": "User not found",
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