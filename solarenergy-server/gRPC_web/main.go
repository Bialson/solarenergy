package main

import (
	"log"
	"net"
	"net/http"
	"time"

	api "github.com/Bialson/solarenergy/proto"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

// SolarServiceServer struct
type solarServer struct {
	api.SolarServiceServer
}

func main() {
	//Generate a TLS certificates for grpc API server
	apiserver, err := GenerateTLSApi("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", "127.0.0.1:9990") // Start listening on a TCP Port
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
	server := &solarServer{}
	api.RegisterSolarServiceServer(apiserver, server) //Register sevice as solarServer
	log.Printf("Server started at: %v", lis.Addr())
	go func() {
		log.Fatalf("failed to start server: %v", apiserver.Serve(lis))
	}()
	grpcWebServer := grpcweb.WrapServer(apiserver) // Wrap the GRPC Server in grpc-web and also host the React app
	multiplex := GrpcMultiplexer{grpcWebServer}
	r := http.NewServeMux()
	// Load the webpage with a http fileserver
	webapp := http.FileServer(http.Dir("$PATH_TO_WEBAPP/build"))
	// Create a HTTP server and bind the router to it
	r.Handle("/", multiplex.Handler(webapp))
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//Listen and serve app on port 8080 over TLS
	log.Fatal(srv.ListenAndServeTLS("cert/server.crt", "cert/server.key"))
}
