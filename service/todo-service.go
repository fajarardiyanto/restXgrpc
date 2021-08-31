package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"todo-grpc/pb"
)

type repoTodo struct {
	db *gorm.DB
}

func NewRepoTodoServer(db *gorm.DB) pb.ToDoServiceServer {
	return &repoTodo{
		db: db,
	}
}

func (r *repoTodo) Create(_ context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	var err error
	if err = r.db.Debug().Model(&pb.ToDo{}).Create(req.Todo).Error; err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Todo: req.Todo,
	}, nil
}

func (r *repoTodo) Read(_ context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	var err error
	var todo pb.ToDo
	if err = r.db.Debug().Model(&todo).Where("id = ? AND author_id = ?", req.Id, req.AuthorId).Take(&todo).Error; err != nil {
		return nil, err
	}

	return &pb.ReadResponse{
		Todo: &todo,
	}, nil
}

func (r *repoTodo) Update(_ context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	var err error
	var todo pb.ToDo
	if err = r.db.Debug().Model(&todo).Where("id = ? AND author_id = ?", req.Todo.Id, req.Todo.AuthorId).Take(&todo).UpdateColumns(&req.Todo).Error; err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{
		Todo: req.Todo,
	}, nil
}

func (r *repoTodo) Delete(_ context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	var err error
	var todo pb.ToDo
	if err = r.db.Debug().Model(&todo).Where("id = ? AND author_id = ?", req.Id, req.AuthorId).Delete(&todo).Error; err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{Id: req.Id}, nil
}

func (r *repoTodo) ReadAll(_ context.Context, req *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	var err error
	var todo []*pb.ToDo
	if err = r.db.Debug().Model(&pb.ToDo{}).Where("author_id = ?", req.AuthorId).Find(&todo).Error; err != nil {
		return nil, err
	}

	return &pb.ReadAllResponse{
		Todo: todo,
	}, nil
}
