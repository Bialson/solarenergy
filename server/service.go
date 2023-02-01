package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	api "github.com/Bialson/solarenergy/proto"
)

const (
	DATA_CAT    = 1002
	PERIOD      = 282
	SECTION     = 156
	MAX_RESULTS = 500
)

var Variables = map[int]string{
	1002:    "Energia elektryczna",
	282:     "Rok - dane roczne",
	156:     "Polska, województwa; Charakter miejscowości",
	186:     "[MWh]",
	187:     "[kWh]",
	188:     "[kWh] - na 1 mieszkańca",
	189:     "[kwh] - na 1 odbiorcę",
	6655092: "Ogółem",
	6655093: "Miasto",
	6655153: "Wieś",
	33617:   "POLSKA",
	33619:   "MAMŁOPOLSKIE",
	33929:   "ŚLĄSKIE",
	34187:   "LUBUSKIE",
	34353:   "WIELKOPOLSKIE",
	34815:   "ZACHODNIOPOMORSKIE",
	35067:   "DOLNOŚLĄSKIE",
	35390:   "OPOLSKIE",
	35542:   "KUJAWEK-POMORSKIE",
	35786:   "POMORSKIE",
	35976:   "WARMIŃSKO-MAZURSKIE",
	36185:   "ŁÓDZKIE",
	36450:   "ŚWIĘTOKRZYSKIE",
	36627:   "LUBELSKIE",
	36924:   "PODKARPACKIE",
	37185:   "PODLASKIE",
	37380:   "MAZOWIECKIE",
}

type EnergyResponseElement struct {
	Rownumber                int64   `json:"rownumber"`
	IdZmienna                int64   `json:"id-zmienna"`
	IdPrzekroj               int64   `json:"id-przekroj"`
	IdWymiar1                int64   `json:"id-wymiar-1"`
	IdPozycja1               int64   `json:"id-pozycja-1"`
	IdWymiar2                int64   `json:"id-wymiar-2"`
	IdPozycja2               int64   `json:"id-pozycja-2"`
	IdOkres                  int64   `json:"id-okres"`
	IdSposobPrezentacjiMiara int64   `json:"id-sposob-prezentacji-miara"`
	IdDaty                   int64   `json:"id-daty"`
	IdBrakWartosci           int64   `json:"id-brak-wartosci"`
	IdTajnosci               int64   `json:"id-tajnosci"`
	IdFlaga                  int64   `json:"id-flaga"`
	Wartosc                  float32 `json:"wartosc"`
	Precyzja                 int64   `json:"precyzja"`
}

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
	dataURL := fmt.Sprintf("https://api-dbw.stat.gov.pl/api/1.1.0/variable/variable-data-section?id-zmienna=%v&id-przekroj=%v&id-rok=%d&id-okres=%v&ile-na-stronie=%d&numer-strony=0&lang=pl", DATA_CAT, SECTION, req.Year, PERIOD, resultsNumber)
	log.Printf("Requesting data from: %s", dataURL)
	dataReq, err := http.Get(dataURL)
	if err != nil {
		log.Fatalf("Could not get data from URL: %v", err)
	}
	defer dataReq.Body.Close()
	dataRes, err := ioutil.ReadAll(dataReq.Body)
	if err != nil {
		log.Fatalf("Could not read data: %v", err)
	}
	log.Printf("Data received: %s", dataRes)
	var dataJSON interface{}
	err = json.Unmarshal(dataRes, &dataJSON)
	if err != nil {
		log.Fatalf("Could not unmarshal data: %v", err)
	}
	energy := dataJSON.(map[string]interface{})["data"].([]interface{})
	for _, occurence := range energy {
		encodedJSON, _ := json.Marshal(occurence)
		var energyElement EnergyResponseElement
		err = json.Unmarshal([]byte(encodedJSON), &energyElement)
		if err != nil {
			log.Fatalf("Could not unmarshal data: %v", err)
		}
		log.Printf("Sending data: %v", energyElement)
		res := &api.PowerFromHomes{
			Value:     energyElement.Wartosc,
			Period:    Variables[int(energyElement.IdOkres)],
			Year:      energyElement.IdDaty,
			Unit:      Variables[int(energyElement.IdSposobPrezentacjiMiara)],
			Precision: energyElement.Precyzja,
			Region:    Variables[int(energyElement.IdPozycja2)],
			Character: Variables[int(energyElement.IdPozycja1)],
		}
		err = stream.Send(res)
		if err != nil {
			log.Fatalf("Could not send data: %v", err)
		}
	}
	return nil
}
