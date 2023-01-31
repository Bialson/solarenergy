package main

import (
	"log"

	api "github.com/Bialson/solarenergy/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := api.NewSolarServiceClient(conn)
	params := &api.PowerConsumptionRequest{
		Year:           2020,
		ResponseAmount: 10,
	}
	givePowerByArea(client, params)
}
