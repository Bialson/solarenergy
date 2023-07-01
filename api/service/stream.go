package service

import (
	"log"

	"github.com/Bialson/solarenergy/assets"
	api "github.com/Bialson/solarenergy/proto"
)

func EnergyConsumptionStream(data []ResponseElement, stream api.SolarService_GetEnergyFromHomesByParamsServer) error {
	for _, el := range data {
		res := &api.PowerFromHomes{
			Value:     el.Wartosc,
			Period:    assets.Variables[int(el.IdOkres)],
			Year:      el.IdDaty,
			Unit:      assets.Units[int(el.IdSposobPrezentacjiMiara)],
			Precision: el.Precyzja,
			Region:    assets.Regions[int(el.IdPozycja1)],
			Character: assets.Regions[int(el.IdPozycja2)],
		}
		err := stream.Send(res)
		if err != nil {
			log.Fatalf("Could not send data: %v", err)
			return err
		}
	}
	return nil
}
