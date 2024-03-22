package main

import (
	"context"
	"github.com/creztfallen/dummy-bank/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	*pb.UnimplementedBankServiceServer
}

func (s *Server) GetBankInformation(ctx context.Context, in *pb.BankRequest) (*pb.BankResponse, error) {
	return &pb.BankResponse{
		BankName:       "Giancardi Bank",
		AccountHolder:  "Geralt",
		AccountNumber:  123456789,
		AccountBalance: 253.32,
	}, nil
}

func main() {
	println("starting new grpc server")

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterBankServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalf("")
	}
}
