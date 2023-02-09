package main

import (
	"context"
	"fmt"
	"io"
	"log"

	api "github.com/Bialson/solarenergy/proto"
)

func GetEnergyFromHomes(client api.SolarServiceClient, params *api.PowerConsumptionRequest) {
	stream, err := client.GetEnergyFromHomesByParams(context.Background(), params) //Calling request for gathering energy data from gRPC server
	if err != nil {
		log.Fatalf("Could not send params: %v", err)
	}
	log.Printf("Data streaming started! \n")
	fmt.Printf("Received data: \n\n")
	for {
		message, err := stream.Recv() //Receiving message from stream
		if err == io.EOF {            //If stream is finished
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming data: %v", err)
		}
		//Printing results to console
		fmt.Println("-------------------------------------------------")
		fmt.Printf("\tValue: %v \n", message.Value)
		fmt.Printf("\tPeriod: %v \n", message.Period)
		fmt.Printf("\tYear: %v \n", message.Year)
		fmt.Printf("\tUnit: %v \n", message.Unit)
		fmt.Printf("\tPrecision: %v \n", message.Precision)
		fmt.Printf("\tRegion: %v \n", message.Region)
		fmt.Printf("\tCharacted: %v \n", message.Character)
		fmt.Printf("-------------------------------------------------\n\n")
	}
	log.Println("Streaming finished!")
}

func GetEcoEnergy(client api.SolarServiceClient, params *api.EcoEnergyRequest) {
	stream, err := client.GetEcoEnergyByParams(context.Background(), params) //Calling request for gathering energy data from gRPC server
	if err != nil {
		log.Fatalf("Could not send params: %v", err)
	}
	log.Printf("Data streaming started! \n")
	fmt.Printf("Received data: \n\n")
	for {
		message, err := stream.Recv() //Receiving message from stream
		if err == io.EOF {            //If stream is finished
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming data: %v", err)
		}
		//Printing results to console
		fmt.Println("-------------------------------------------------")
		fmt.Printf("\tValue: %v \n", message.Value)
		fmt.Printf("\tPeriod: %v \n", message.Period)
		fmt.Printf("\tYear: %v \n", message.Year)
		fmt.Printf("\tUnit: %v \n", message.Unit)
		fmt.Printf("\tPrecision: %v \n", message.Precision)
		fmt.Printf("\tType: %v \n", message.Type)
		fmt.Printf("\tRegion: %v \n", message.Region)
		fmt.Printf("-------------------------------------------------\n\n")
	}
	log.Println("Streaming finished!")
}
