package main

import (
	"context"

	api "github.com/Bialson/solarenergy/proto"
)

func (s *solarServer) GetSolarEnergy(ctx context.Context, req *api.NoParam) (*api.Response, error) {
	return &api.Response{Value: 100}, nil
}
