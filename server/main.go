package main

import (
	"log"
	"net"

	api "github.com/Bialson/solarenergy/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type solarServer struct {
	api.SolarServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	server := grpc.NewServer()
	api.RegisterSolarServiceServer(server, &solarServer{})
	log.Printf("Server started at: %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
