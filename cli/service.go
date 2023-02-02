package main

import (
	"context"
	"fmt"
	"io"
	"log"

	api "github.com/Bialson/solarenergy/proto"
)

func givePowerByArea(client api.SolarServiceClient, params *api.PowerConsumptionRequest) {
	stream, err := client.GetSolarEnergyFromHomesByParams(context.Background(), params)
	if err != nil {
		log.Fatalf("Could not send params: %v", err)
	}
	log.Printf("Data streaming started! \n")
	fmt.Printf("Received data: \n\n")
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Oops, no data found! \n")
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming data: %v", err)
		}
		fmt.Println("-------------------------------------------------")
		fmt.Printf("\tValue: %v \n", message.Value)
		fmt.Printf("\tPeriod: %v \n", message.Period)
		fmt.Printf("\tYear: %v \n", message.Year)
		fmt.Printf("\tUnit: %v \n", message.Unit)
		fmt.Printf("\tPrecision: %v \n", message.Precision)
		fmt.Printf("\tRegion: %v \n", message.Region)
		fmt.Printf("\tCharacted: %v \n", message.Character)
	}
	log.Printf("Streaming finished!")
}
