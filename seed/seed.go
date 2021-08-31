package seed

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	"todo-grpc/pb"
)

func LoadSeed(db *gorm.DB) {
	var logger log.Logger

	if ok := db.HasTable(&pb.User{}); !ok {
		if err := db.Debug().CreateTable(&pb.User{}).Error; err != nil {
			level.Error(logger).Log("cannot create table: ", err.Error())
		}

		if err := db.Debug().AutoMigrate(&pb.User{}).Error; err != nil {
			level.Error(logger).Log("cannot migrate table: ", err.Error())
		}
	}

	if ok := db.HasTable(&pb.ToDo{}); !ok {
		if err := db.Debug().CreateTable(&pb.ToDo{}).Error; err != nil {
			level.Error(logger).Log("cannot create table: ", err.Error())
		}

		if err := db.Debug().AutoMigrate(&pb.ToDo{}).Error; err != nil {
			level.Error(logger).Log("cannot migrate table: ", err.Error())
		}

		//if err := db.Debug().Model(&pb.ToDo{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error; err != nil {
		//	level.Error(logger).Log("attaching foreign key error: %v", err.Error())
		//}
	}
}
