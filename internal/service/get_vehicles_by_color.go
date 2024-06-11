package service

import (
	"strings"

	"github.com/nurhidaylma/parking-app/internal/model"
	"github.com/nurhidaylma/parking-app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *parkingService) GetVehiclesByColor(req model.GetVehiclesByColorRequest) (model.GetVehiclesByColorResponse, error) {
	vehiclesParked, err := s.repo.ReadVehicles()
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.GetVehiclesByColorResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}

	vehicles := []string{}
	req.VehicleColor = strings.ToLower(req.VehicleColor)
	for _, v := range vehiclesParked {
		vehicleColorLoweCase := strings.ToLower(v.Color)
		if vehicleColorLoweCase == req.VehicleColor {
			vehicles = append(vehicles, v.Number)
		}
	}

	return model.GetVehiclesByColorResponse{VehicleNumbers: vehicles}, nil
}
