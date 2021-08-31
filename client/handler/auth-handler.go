package handler

import (
	"context"
	"net/http"
	"time"
	"todo-grpc/client/utils"
	"todo-grpc/pb"
)

type authServiceClient struct {
	c pb.AuthServiceClient
}

func NewAuthServiceClient(c pb.AuthServiceClient) *authServiceClient {
	return &authServiceClient{c}
}

func (auth *authServiceClient) Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var req pb.Auth

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req.Username = r.FormValue("username")
	req.Password = r.FormValue("password")

	request := pb.AuthRequest{
		Auth: &req,
	}

	res, err := auth.c.Auth(ctx, &request)
	if err != nil {
		utils.Respond(w, http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	cookie := &http.Cookie{
		Name:"user_todos",
		Value: res.AccessToken,
		Path: "/",
		MaxAge: 86400,
		Secure: true,
		HttpOnly: true,
	}
	r.AddCookie(cookie)

	//for _, c := range r.Cookies() {
	//	fmt.Println(c.Value)
	//}
	utils.Respond(w, http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "",
		"data":    res.AccessToken,
	})
}
