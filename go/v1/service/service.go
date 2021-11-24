package service

import (
	"context"
	"database/sql"

	pb "github.com/inhun/grpc-study/go/v1/proto"
)

type UserServer struct {
	pb.UserServer
	DB *sql.DB
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	userID := req.Id

	var userMessage *pb.UserMessage

	row, _ := s.DB.Query(`SELECT id, name, email FROM users WHERE id = ?`, userID)

	if row.Next() {
		_ = row.Scan(&userMessage.Id, &userMessage.Name, &userMessage.Email)

	}

	return &pb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}
