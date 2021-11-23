package service

import (

	pb "github.com/inhun/grpc-study/go/v1/proto"
	"database/sql"
)


type UserServer struct {
	pb.UserServer
	DB *sql.DB
}

func (s *UserServer) GetUser()