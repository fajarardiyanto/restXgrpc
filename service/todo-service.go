package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"todo-grpc/pb"
)

type repo struct {
	db *gorm.DB
}

func NewRepoServer(db *gorm.DB) pb.ToDoServiceServer {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(_ context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	var err error
	if err = r.db.Debug().Model(&pb.ToDo{}).Create(req.Todo).Error; err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Todo: req.Todo,
	}, nil
}

func (r *repo) Read(_ context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	var err error
	var todo pb.ToDo
	if err = r.db.Debug().Model(&todo).Where("id = ?", req.Id).Take(&todo).Error; err != nil {
		return nil, err
	}

	return &pb.ReadResponse{
		Todo: &todo,
	}, nil
}

func (r *repo) Update(_ context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	var err error
	var todo pb.ToDo
	if err = r.db.Debug().Model(&todo).Where("id = ?", req.Todo.Id).Take(&todo).UpdateColumns(&req.Todo).Error; err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{
		Todo: req.Todo,
	}, nil
}

func (r *repo) Delete(_ context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	var err error
	var todo pb.ToDo
	if err = r.db.Debug().Model(&todo).Where("id = ?", req.Id).Delete(&todo).Error; err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{Id: req.Id}, nil
}

func (r *repo) ReadAll(_ context.Context, _ *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	var err error
	var todo []*pb.ToDo
	if err = r.db.Debug().Model(&pb.ToDo{}).Find(&todo).Error; err != nil {
		return nil, err
	}

	return &pb.ReadAllResponse{
		Todo: todo,
	}, nil
}
