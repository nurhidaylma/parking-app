package service

import (
	"github.com/nurhidaylma/parking-app/internal/model"
	"github.com/nurhidaylma/parking-app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *parkingService) GetVehiclesByType(req model.GetVehiclesByTypeRequest) (model.GetVehiclesByTypeResponse, error) {
	if !validVehicleType(req.VehicleType) {
		util.Logger.LogWarning("invalid vehicle type")
		return model.GetVehiclesByTypeResponse{}, status.Error(codes.FailedPrecondition, "invalid vehicle type")
	}

	vehiclesParked, err := s.repo.ReadVehicles()
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.GetVehiclesByTypeResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}

	var totalVehicles int
	for _, v := range vehiclesParked {
		if v.Type == util.VehicleTypes(req.VehicleType) {
			totalVehicles++
		}
	}

	return model.GetVehiclesByTypeResponse{TotalVehicle: totalVehicles}, nil
}

func validVehicleType(vehicleType string) bool {
	types := make(map[util.VehicleTypes]bool)
	types[util.VehicleTypeSUV] = true
	types[util.VehicleTypeMPV] = true

	return types[util.VehicleTypes(vehicleType)]
}
