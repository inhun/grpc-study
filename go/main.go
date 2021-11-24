package main

import (
	"log"
	"net"
	"fmt"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/inhun/grpc-study/go/v1/utils"
	"github.com/inhun/grpc-study/go/v1/service"
	pb "github.com/inhun/grpc-study/go/v1/proto"
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
