package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"todo-grpc/pb"
)

type repoUser struct {
	db *gorm.DB
}

func NewRepoUserService(db *gorm.DB) pb.UserServiceServer {
	return &repoUser{
		db: db,
	}
}

func (r *repoUser) Create(_ context.Context, req *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	var err error
	if err = r.db.Debug().Model(&pb.User{}).Create(req.User).Error; err != nil {
		return nil, err
	}

	return &pb.UserCreateResponse{
		User: req.User,
	}, nil
}
func (r *repoUser) Read(_ context.Context, req *pb.UserReadRequest) (*pb.UserReadResponse, error){
	var err error
	var user pb.User
	if err = r.db.Debug().Model(&user).Where("id = ?", req.Id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &pb.UserReadResponse{
		User: &user,
	}, nil
}
func (r *repoUser) Update(_ context.Context, req *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	var err error
	var user pb.User
	if err = r.db.Debug().Model(&user).Where("id = ?", req.User.Id).Take(&user).UpdateColumns(&req.User).Error; err != nil {
		return nil, err
	}

	return &pb.UserUpdateResponse{
		User: req.User,
	}, nil
}
func (r *repoUser) Delete(_ context.Context, req *pb.UserDeleteRequest) (*pb.UserDeleteResponse, error) {
	var err error
	var user pb.User
	if err = r.db.Debug().Model(&user).Where("id = ?", req.Id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return &pb.UserDeleteResponse{Id: req.Id}, nil
}
func (r *repoUser) ReadAll(_ context.Context, _ *pb.UserReadAllRequest) (*pb.UserReadAllResponse, error) {
	var err error
	var user []*pb.User
	if err = r.db.Debug().Model(&pb.User{}).Find(&user).Error; err != nil {
		return nil, err
	}

	return &pb.UserReadAllResponse{
		User: user,
	}, nil
}
