package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	api "github.com/Bialson/solarenergy/proto"
)

const (
	VARIABLE    = 1002
	PERIOD      = 282
	SECTION     = 156
	MAX_RESULTS = 500
)

func (s *solarServer) GetSolarEnergy(ctx context.Context, req *api.NoParam) (*api.PowerResponse, error) {
	return &api.PowerResponse{Value: 100}, nil
}

func (s *solarServer) GetSolarEnergyFromHomesByParams(req *api.PowerConsumptionRequest, stream api.SolarService_GetSolarEnergyFromHomesByParamsServer) error {
	log.Printf("Received params: %v", req)
	resultsNumber := req.ResponseAmount
	if resultsNumber > MAX_RESULTS {
		log.Printf("Requested amount of results is too big, max is %d", MAX_RESULTS)
		resultsNumber = MAX_RESULTS
	}
	dataURL := fmt.Sprintf("https://api-dbw.stat.gov.pl/api/1.1.0/variable/variable-data-section?id-zmienna=%v&id-przekroj=%v&id-rok=%d&id-okres=%v&ile-na-stronie=%d&numer-strony=0&lang=pl", VARIABLE, SECTION, req.Year, PERIOD, resultsNumber)
	log.Printf("Requesting data from: %s", dataURL)
	dataReq, err := http.Get(dataURL)
	if err != nil {
		log.Fatalf("Could not get data from URL: %v", err)
	}
	defer dataReq.Body.Close()
	dataRes, err := ioutil.ReadAll(dataReq.Body)
	if err != nil {
		return err
	}
	log.Printf("Data received: %s", dataRes)
	return nil
}
