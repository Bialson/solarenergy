package main

import (
	"context"
	"log"
	"net"

	api "github.com/Bialson/solarenergy/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func TestServer(ctx context.Context) (api.SolarServiceClient, func()) {
	lis, err := net.Listen("tcp", port)
	baseServer := grpc.NewServer()
	api.RegisterSolarServiceServer(baseServer, &solarServer{})
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()
	log.Printf("Server started at: %v", lis.Addr())
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}
	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}
	client := api.NewSolarServiceClient(conn)

	return client, closer
}
