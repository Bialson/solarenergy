package main

import (
	"log"
	"net"

	"github.com/Bialson/solarenergy/api/controller"
	"github.com/Bialson/solarenergy/assets"
	api "github.com/Bialson/solarenergy/proto"
	"google.golang.org/grpc"
)

func main() {
	//Server initialization
	lis, err := net.Listen("tcp", assets.PORT)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	server := grpc.NewServer()

	api.RegisterSolarServiceServer(server, &controller.SolarServer{})

	log.Printf("Server started at: %v", lis.Addr())

	if err := server.Serve(lis); err != nil { //Serving server
		log.Fatalf("failed to serve: %v", err)
	}
}
