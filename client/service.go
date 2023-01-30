package main

import (
	"context"
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
	log.Printf("Power: %v", res.Value)
}
