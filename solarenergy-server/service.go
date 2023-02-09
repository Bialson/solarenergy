package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	api "github.com/Bialson/solarenergy/proto"
)

//Structures of each message is defined in proto file (proto/energy.proto)

//GetSolarEnergyFromHomesByParams method implementation, request -> PowerConsumptionRequest message, response -> stream of PowerConsumptionResponse message, error

//Method is responsible for getting energy consumption data from API DBW, filtering it and sending response to client

func (s *solarServer) GetEnergyFromHomesByParams(req *api.PowerConsumptionRequest, stream api.SolarService_GetEnergyFromHomesByParamsServer) error {
	log.Printf("Received params: %v", req)
	dataURL := fmt.Sprintf("https://api-dbw.stat.gov.pl/api/1.1.0/variable/variable-data-section?sorts=id-pozycja-2&id-zmienna=%v&id-przekroj=%v&id-rok=%d&id-okres=%v&ile-na-stronie=%d&numer-strony=0&lang=pl", DATA_CAT, SECTION_1, req.Year, PERIOD, MAX_RESULTS) //URL for data request based on request parameters
	log.Printf("Requesting data from: %s", dataURL)
	dataReq, err := http.Get(dataURL) //Requesting data from URL
	if err != nil {
		log.Fatalf("Could not get data from URL: %v", err)
	}
	defer dataReq.Body.Close()
	//Indetyfying response status
	status := dataReq.StatusCode
	if status != 200 {
		log.Printf("Bad request or server not responding, ERR_CODE: %v", status)
	} else {
		dataRes, err := io.ReadAll(dataReq.Body) //Reading response body
		if err != nil {
			log.Fatalf("Could not read data: %v", err)
		}
		dataJSON := EnergyData{}
		//Unmarshalling JSON data to EnergyData struct
		err = json.Unmarshal([]byte(dataRes), &dataJSON)
		if err != nil {
			log.Fatalf("Could not unmarshal data: %v", err)
		}
		log.Printf("Data received count: %v", len(dataJSON.Energy))
		EnergyDataArr = dataJSON.Energy
		for _, el := range EnergyDataArr {
			fmt.Println(el)
		}
		// Filtering data based on request parameters
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
		//Limiting number of response elements
		if req.ResponseAmount != 0 && int(req.ResponseAmount) < len(EnergyDataArrFiltered) {
			EnergyDataArrFiltered = EnergyDataArrFiltered[:req.ResponseAmount]
		}
		for _, el := range EnergyDataArrFiltered {
			//Generating response message
			res := &api.PowerFromHomes{
				Value:     el.Wartosc,
				Period:    Variables[int(el.IdOkres)],
				Year:      el.IdDaty,
				Unit:      Units[int(el.IdSposobPrezentacjiMiara)],
				Precision: el.Precyzja,
				Region:    Regions[int(el.IdPozycja1)],
				Character: Regions[int(el.IdPozycja2)],
			}
			err = stream.Send(res) //Sending response message to stream
			if err != nil {
				log.Fatalf("Could not send data: %v", err)
			}
		}
		EnergyDataArr = nil
		EnergyDataArrFiltered = nil
	}
	return nil
}
