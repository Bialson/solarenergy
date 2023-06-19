package main

import (
	"fmt"

	api "github.com/Bialson/solarenergy/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateClient() (client api.SolarServiceClient, conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials())) //Connect to server
	if err != nil {
		fmt.Printf(" * Cannot connect to server: %v \n Exiting", err)
		return nil, nil, err
	} else {
		fmt.Println(" * Connected to server!")
	}
	client = api.NewSolarServiceClient(conn)
	return client, conn, nil
}
