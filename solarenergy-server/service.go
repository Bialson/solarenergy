package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	api "github.com/Bialson/solarenergy/proto"
)

func (s *solarServer) GetSolarEnergyFromHomesByParams(req *api.PowerConsumptionRequest, stream api.SolarService_GetSolarEnergyFromHomesByParamsServer) error {
	log.Printf("Received params: %v", req)
	dataURL := fmt.Sprintf("https://api-dbw.stat.gov.pl/api/1.1.0/variable/variable-data-section?sorts=id-pozycja-2&id-zmienna=%v&id-przekroj=%v&id-rok=%d&id-okres=%v&ile-na-stronie=%d&numer-strony=0&lang=pl", DATA_CAT, SECTION, req.Year, PERIOD, MAX_RESULTS)
	log.Printf("Requesting data from: %s", dataURL)
	dataReq, err := http.Get(dataURL)
	if err != nil {
		log.Fatalf("Could not get data from URL: %v", err)
	}
	defer dataReq.Body.Close()
	status := dataReq.StatusCode
	if status != 200 {
		log.Printf("Bad request or server not responding, ERR_CODE: %v", status)
	} else {
		dataRes, err := ioutil.ReadAll(dataReq.Body)
		if err != nil {
			log.Fatalf("Could not read data: %v", err)
		}
		var dataJSON interface{}
		err = json.Unmarshal(dataRes, &dataJSON)
		if err != nil {
			log.Fatalf("Could not unmarshal data: %v", err)
		}
		energyJSON := dataJSON.(map[string]interface{})["data"].([]interface{})
		log.Printf("Data received count: %v", len(energyJSON))
		for _, occurence := range energyJSON {
			encodedEnergyJSONElement, _ := json.Marshal(occurence)
			var el EnergyElement
			err = json.Unmarshal([]byte(encodedEnergyJSONElement), &el)
			if err != nil {
				log.Fatalf("Could not unmarshal data: %v", err)
			}
			EnergyDataArr = append(EnergyDataArr, el)
		}
		if req.Region != "" && req.Character != "" {
			EnergyDataArrFiltered = FilterByCharacterAndRegion(req.Character, req.Region)
		} else if req.Character != "" {
			EnergyDataArrFiltered = FilterByCharacter(req.Character)
			QuickSortByRegion(EnergyDataArrFiltered, 0, len(EnergyDataArrFiltered)-1)
		} else if req.Region != "" {
			EnergyDataArrFiltered = FilterByRegion(req.Region)
		} else {
			EnergyDataArrFiltered = EnergyDataArr
			QuickSortByRegion(EnergyDataArrFiltered, 0, len(EnergyDataArrFiltered)-1)
		}
		log.Printf("Filtered data count: %v", len(EnergyDataArrFiltered))
		if req.ResponseAmount != 0 && int(req.ResponseAmount) < len(EnergyDataArrFiltered) {
			EnergyDataArrFiltered = EnergyDataArrFiltered[:req.ResponseAmount]
		}
		for _, el := range EnergyDataArrFiltered {
			res := &api.PowerFromHomes{
				Value:     el.Wartosc,
				Period:    Variables[int(el.IdOkres)],
				Year:      el.IdDaty,
				Unit:      Variables[int(el.IdSposobPrezentacjiMiara)],
				Precision: el.Precyzja,
				Region:    Variables[int(el.IdPozycja1)],
				Character: Variables[int(el.IdPozycja2)],
			}
			err = stream.Send(res)
			if err != nil {
				log.Fatalf("Could not send data: %v", err)
			}
		}
		EnergyDataArr = nil
		EnergyDataArrFiltered = nil
	}
	return nil
}
