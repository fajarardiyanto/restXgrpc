package service

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"todo-grpc/auth"
	"todo-grpc/client/utils"
	"todo-grpc/pb"
)

type repoAuth struct {
	db *gorm.DB
}

func NewRepoAuthService(db *gorm.DB) pb.AuthServiceServer {
	return &repoAuth{
		db: db,
	}
}

func (r *repoAuth) Auth(_ context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	var err error
	var user pb.User

	if err = r.db.Debug().Model(&user).Where("username = ?", req.Auth.Username).Take(&user).Error; err != nil {
		return nil, errors.New("username not found")
	}

	if err = utils.VerifyPassword(user.Password, req.Auth.Password); err != nil {
		return nil, errors.New("password not matching")
	}

	token, err := auth.CreateToken(user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.AuthResponse{
		AccessToken: token,
	}, err
}

