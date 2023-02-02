package main

import (
	"context"
	"errors"
	"io"
	"log"
	"testing"

	api "github.com/Bialson/solarenergy/proto"
)

func TestGetSolarEnergy(t *testing.T) {
	ctx := context.Background()
	client, conn := TestServer(ctx)
	defer conn()
	type expectation struct {
		out []*api.PowerFromHomes
		len int
		err error
	}

	tests := map[string]struct {
		in       *api.PowerConsumptionRequest
		expected expectation
	}{
		"Sort amount of res withour filters": {
			in: &api.PowerConsumptionRequest{
				Year:           2020,
				ResponseAmount: 4,
				Region:         "",
				Character:      "",
			},
			expected: expectation{
				out: []*api.PowerFromHomes{
					{
						Value:     879201.125000,
						Period:    "Rok - dane roczne",
						Year:      2020,
						Unit:      "[MWh]",
						Precision: 1,
						Region:    "DOLNOŚLĄSKIE",
						Character: "Wieś",
					},
					{
						Value:     879201152.000000,
						Period:    "Rok - dane roczne",
						Year:      2020,
						Unit:      "[kWh]",
						Precision: 1,
						Region:    "DOLNOŚLĄSKIE",
						Character: "Wieś",
					},
					{
						Value:     957.400024,
						Period:    "Rok - dane roczne",
						Year:      2020,
						Unit:      "[kWh] - na 1 mieszkańca",
						Precision: 1,
						Region:    "DOLNOŚLĄSKIE",
						Character: "Wieś",
					},
					{
						Value:     2588.199951,
						Period:    "Rok - dane roczne",
						Year:      2020,
						Unit:      "[kWh] - na 1 odbiorcę",
						Precision: 1,
						Region:    "DOLNOŚLĄSKIE",
						Character: "Wieś",
					},
				},
				len: 4,
				err: nil,
			},
		},
		"Responses with filters": {
			in: &api.PowerConsumptionRequest{
				Year:           2020,
				ResponseAmount: 2,
				Region:         "OPOLSKIE",
				Character:      "Miasto",
			},
			expected: expectation{
				out: []*api.PowerFromHomes{
					{
						Value:     416418.406250,
						Period:    "Rok - dane roczne",
						Year:      2020,
						Unit:      "[MWh]",
						Precision: 1,
						Region:    "OPOLSKIE",
						Character: "Miasto",
					},
					{
						Value:     416418432.000000,
						Period:    "Rok - dane roczne",
						Year:      2020,
						Unit:      "[kWh]",
						Precision: 1,
						Region:    "OPOLSKIE",
						Character: "Miasto",
					},
				},
				len: 2,
				err: nil,
			},
		},
	}
	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			log.Printf("Test: %s", scenario)
			stream, err := client.GetSolarEnergyFromHomesByParams(ctx, tt.in)
			var outs []*api.PowerFromHomes
			for {
				message, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					break
				}
				outs = append(outs, message)
			}
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Expected error: %v, got: %v", tt.expected.err, err)
				}
			} else {
				if len(outs) != len(tt.expected.out) && len(outs) != tt.expected.len {
					t.Errorf("Expected %d responses, got %d", len(tt.expected.out), len(outs))
				} else {
					for i, el := range outs {
						if el.Value != tt.expected.out[i].Value ||
							el.Period != tt.expected.out[i].Period ||
							el.Year != tt.expected.out[i].Year ||
							el.Unit != tt.expected.out[i].Unit ||
							el.Precision != tt.expected.out[i].Precision ||
							el.Region != tt.expected.out[i].Region ||
							el.Character != tt.expected.out[i].Character {
							t.Errorf("Expected: %v, got: %v", tt.expected.out, outs)
						}
					}
				}
			}
		})
	}
}
