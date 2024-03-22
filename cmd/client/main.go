package main

import (
	"context"
	"github.com/creztfallen/dummy-bank/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewBankServiceClient(conn)
	req := &pb.BankRequest{
		BankName:      "Giancardi",
		AccountNumber: 123456789,
		AccountType:   pb.BankType_CHECKING,
	}

	res, err := client.GetBankInformation(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
