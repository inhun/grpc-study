package service

import (

	pb "github.com/inhun/grpc-study/go/v1/proto"
	"database/sql"
)


type UserServer struct {
	pb.UserServer
	DB *sql.DB
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUsrRequest) (*pb.GetUserResponse, error) {
	userID := req.UserId

	var userMessage *pb.UserMessage

	row, err := s.DB.Query(`SELECT id, name, email FROM users WHERE id = ?`, userID)
	
	if row.Next() {
		_ := row.Scan(&userMessage.Id, &userMessage.Name, &userMessage.Email)
		
	}

	return &pb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}