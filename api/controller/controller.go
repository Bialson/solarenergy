package controller

import (
	"context"
	"log"

	"github.com/Bialson/solarenergy/api/service"
	"github.com/Bialson/solarenergy/assets"
	api "github.com/Bialson/solarenergy/proto"
)

type SolarServer struct {
	api.SolarServiceServer
}

func (s *SolarServer) SeyHello(ctx context.Context, req *api.HelloReq) (*api.HelloRes, error) {
	log.Printf("Received: %v", req.Name)
	return &api.HelloRes{Message: "Hello " + req.GetName()}, nil
}

func (s *SolarServer) GetEnergyFromHomesByParams(req *api.PowerConsumptionRequest, stream api.SolarService_GetEnergyFromHomesByParamsServer) error {
	log.Printf("Received params: %v", req)
	res := service.API.RequestDBWData(req.Year, assets.DATA_CAT_1, assets.SECTION_1)
	defer res.Body.Close()
	service.API.ExtractJSONData(res)
	if status := res.StatusCode; status != 200 { //Identyfing status code of response
		log.Printf("Bad request or server not responding, ERR_CODE: %v", status)
	} else {
		err := service.API.SendStreamData(stream)
		if err != nil {
			log.Printf("Error while sending data, ERR_CODE: %v", status)
		}
		// filters := map[string]string{"region": req.Region, "character": req.Character}
		// service.API.ApplyFilters(filters, req.ResponseAmount)
		// service.API.SortByRegion(0, len(assets.DataArray.Energy)-1)
		// for _, el := range assets.DataArray.Energy {
		// 	res := &api.PowerFromHomes{
		// 		Value:     el.Wartosc,
		// 		Period:    assets.Variables[int(el.IdOkres)],
		// 		Year:      el.IdDaty,
		// 		Unit:      assets.Units[int(el.IdSposobPrezentacjiMiara)],
		// 		Precision: el.Precyzja,
		// 		Region:    assets.Regions[int(el.IdPozycja1)],
		// 		Character: assets.Regions[int(el.IdPozycja2)],
		// 	}
		// 	err := stream.Send(res) //Sending response message to stream
		// 	if err != nil {
		// 		log.Fatalf("Could not send data: %v", err)
		// 	}
		// }
	}
	return nil
}

func (s *SolarServer) GetEcoEnergyByParams(req *api.EcoEnergyRequest, stream api.SolarService_GetEcoEnergyByParamsServer) error {
	log.Printf("Received params: %v", req)
	res := service.API.RequestDBWData(req.Year, assets.DATA_CAT_2, assets.SECTION_2)
	defer res.Body.Close()
	service.API.ExtractJSONData(res)
	if status := res.StatusCode; status != 200 { //Identyfing status code of response
		log.Printf("Bad request or server not responding, ERR_CODE: %v", status)
	} else {
		err := service.API.SendStreamData(stream)
		if err != nil {
			log.Printf("Error while sending data, ERR_CODE: %v", status)
		}
		// filters := map[string]string{"unit": req.Unit, "type": req.Type}
		// service.API.ApplyFilters(filters, req.ResponseAmount)
		// for _, el := range assets.DataArray.Energy {
		// 	//Generating response message
		// 	res := &api.EcoEnergy{
		// 		Value:     el.Wartosc,
		// 		Period:    assets.Variables[int(el.IdOkres)],
		// 		Year:      el.IdDaty,
		// 		Unit:      assets.Units[int(el.IdSposobPrezentacjiMiara)],
		// 		Precision: el.Precyzja,
		// 		Type:      assets.EnergyTypes[int(el.IdPozycja2)],
		// 		Region:    assets.Regions[int(el.IdPozycja1)],
		// 	}
		// 	err := stream.Send(res) //Sending response message to stream
		// 	if err != nil {
		// 		log.Fatalf("Could not send data: %v", err)
		// 	}
		// }
	}
	return nil
}
