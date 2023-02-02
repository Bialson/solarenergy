package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	api "github.com/Bialson/solarenergy/proto"
)

func givePower(client api.SolarServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.GetSolarEnergy(ctx, &api.NoParam{})
	if err != nil {
		log.Fatalf("could not get power: %v", err)
	}
	log.Printf("Power: %v", res)
}

func givePowerByArea(client api.SolarServiceClient, params *api.PowerConsumptionRequest) {
	log.Printf("Streaming started")
	stream, err := client.GetSolarEnergyFromHomesByParams(context.Background(), params)
	if err != nil {
		log.Fatalf("Could not send params: %v", err)
	}
	fmt.Println("Received data: ")
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming data: %v", err)
		}
		fmt.Printf(`
	Value: %10f
	Period: %10s
	Year: %d
	Unit: %s
	Precision: %d
	Region: %s
	Character: %s
`, message.Value, message.Period, message.Year, message.Unit, message.Precision, message.Region, message.Character)
	}
	log.Println("Streaming finished")
}
