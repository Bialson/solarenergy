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
	VARIABLE    = 1002
	PERIOD      = 282
	SECTION     = 156
	MAX_RESULTS = 500
)

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
	Wartosc                  float64 `json:"wartosc"`
	Precyzja                 int64   `json:"prezycja"`
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
	var dataJSON interface{}
	err = json.Unmarshal([]byte(dataRes), &dataJSON)
	if err != nil {
		log.Fatalf("Could not unmarshal data: %v", err)
	}
	energy := dataJSON.(map[string]interface{})["data"].([]interface{})
	for _, occurence := range energy {
		occurence = occurence.(map[string]interface{})
		decoded, err := json.Marshal(occurence)
		if err != nil {
			log.Fatalf("Could not encode data: %v", err)
		}
		var energyElement EnergyResponseElement
		err = json.Unmarshal(decoded, &energyElement)
		if err != nil {
			log.Fatalf("Could not unmarshal data: %v", err)
		}
		res := &api.PowerFromHomes{
			Value:     float32(energyElement.Wartosc),
			Period:    fmt.Sprint(energyElement.IdOkres),
			Year:      int64(energyElement.IdDaty),
			Unit:      fmt.Sprint(energyElement.IdSposobPrezentacjiMiara),
			Precision: int64(energyElement.Precyzja),
			Character: fmt.Sprint(energyElement.IdPozycja2),
		}
		err = stream.Send(res)
		if err != nil {
			log.Fatalf("Could not send data: %v", err)
		}
	}
	return nil
}
