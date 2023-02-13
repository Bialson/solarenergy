package main

import (
	"log"

	api "github.com/Bialson/solarenergy/proto"
)

// Creating a new instance of EnergyDataArr and assigning it to EnergyService instance of Energy interface.
var EnergyDataArr = EnergyData{}
var EnergyService Energy = &EnergyDataArr

//Structures of each message is defined in proto file (proto/energy.proto)

//GetEnergyFromHomesByParams method implementation, request -> PowerConsumptionRequest message, response -> stream of PowerConsumptionResponse message, error

//Method is responsible for getting energy consumption data from API DBW, filtering it and sending response to client

func (s *solarServer) GetEnergyFromHomesByParams(req *api.PowerConsumptionRequest, stream api.SolarService_GetEnergyFromHomesByParamsServer) error {
	log.Printf("Received params: %v", req)
	res := EnergyService.RequestDBWData(req.Year, DATA_CAT_1, SECTION_1)
	defer res.Body.Close()
	EnergyService.ExtractJSONData(res)
	if status := res.StatusCode; status != 200 { //Identyfing status code of response
		log.Printf("Bad request or server not responding, ERR_CODE: %v", status)
	} else {
		filters := map[string]string{"region": req.Region, "character": req.Character}
		EnergyService.ApplyFilters(filters, req.ResponseAmount)
		EnergyService.SortByRegion(0, len(EnergyDataArr.Energy)-1)
		for _, el := range EnergyDataArr.Energy {
			res := &api.PowerFromHomes{
				Value:     el.Wartosc,
				Period:    Variables[int(el.IdOkres)],
				Year:      el.IdDaty,
				Unit:      Units[int(el.IdSposobPrezentacjiMiara)],
				Precision: el.Precyzja,
				Region:    Regions[int(el.IdPozycja1)],
				Character: Regions[int(el.IdPozycja2)],
			}
			err := stream.Send(res) //Sending response message to stream
			if err != nil {
				log.Fatalf("Could not send data: %v", err)
			}
		}
	}
	return nil
}

func (s *solarServer) GetEcoEnergyByParams(req *api.EcoEnergyRequest, stream api.SolarService_GetEcoEnergyByParamsServer) error {
	log.Printf("Received params: %v", req)
	res := EnergyService.RequestDBWData(req.Year, DATA_CAT_2, SECTION_2)
	defer res.Body.Close()
	EnergyService.ExtractJSONData(res)
	if status := res.StatusCode; status != 200 { //Identyfing status code of response
		log.Printf("Bad request or server not responding, ERR_CODE: %v", status)
	} else {
		filters := map[string]string{"unit": req.Unit, "type": req.Type}
		EnergyService.ApplyFilters(filters, req.ResponseAmount)
		for _, el := range EnergyDataArr.Energy {
			//Generating response message
			res := &api.EcoEnergy{
				Value:     el.Wartosc,
				Period:    Variables[int(el.IdOkres)],
				Year:      el.IdDaty,
				Unit:      Units[int(el.IdSposobPrezentacjiMiara)],
				Precision: el.Precyzja,
				Type:      Types[int(el.IdPozycja2)],
				Region:    Regions[int(el.IdPozycja1)],
			}
			err := stream.Send(res) //Sending response message to stream
			if err != nil {
				log.Fatalf("Could not send data: %v", err)
			}
		}
	}
	return nil
}
