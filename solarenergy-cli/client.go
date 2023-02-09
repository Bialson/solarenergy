package main

import (
	"fmt"

	api "github.com/Bialson/solarenergy/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateClient() (client api.SolarServiceClient, err error, conn *grpc.ClientConn) {
	conn, err = grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) //Connect to server
	if err != nil {
		fmt.Printf(" * Cannot connect to server: %v \n Exiting", err)
		return nil, err, nil
	} else {
		fmt.Println(" * Connected to server!")
	}
	client = api.NewSolarServiceClient(conn)
	return
}
