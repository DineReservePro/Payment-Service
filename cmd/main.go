package main

import (
	"net"
	"payment-service/config"
	pb "payment-service/generated/payment_service"
	"payment-service/service"
	"payment-service/storage/postgres"

	"google.golang.org/grpc"
)

func main(){
	db,err := postgres.ConnectDB()
	if err != nil{
		panic(err)
	}
	defer db.Close()

	config := config.Load()

	listener,err := net.Listen("tcp",":"+config.URL_PORT)
	if err != nil{
		panic(err)
	}
	defer listener.Close()

	s := service.NewPaymentService(*postgres.NewPaymentRepo(db))
	server := grpc.NewServer()
	pb.RegisterPaymentServiceServer(server,s)
	if err := server.Serve(listener);err != nil{
		panic(err)
	}
}
