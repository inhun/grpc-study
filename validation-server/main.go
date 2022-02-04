package main

import (
	"fmt"
	"log"
	"net"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	pb "github.com/inhun/grpc-study/validation-server/v1/proto"
	"github.com/inhun/grpc-study/validation-server/v1/service"
	"github.com/inhun/grpc-study/validation-server/v1/utils"
	"google.golang.org/grpc"
)

const portNumber = "9000"

func main() {
	cfg, _ := utils.LoadConfig("config.json")

	db, err := sql.Open("mysql", cfg.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	
	lis, _ := net.Listen("tcp", ":"+portNumber)
	grpcServer := grpc.NewServer()

	pb.RegisterUserServer(grpcServer, &service.UserServer{DB: db})

	
	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("errorerrorerrorerrorerrorerrorerrorerror")
		log.Fatalf("failed")
	}
}
