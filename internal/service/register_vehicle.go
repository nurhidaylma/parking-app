package service

import (
	"time"

	"github.com/nurhidaylma/parking-app/internal/model"
	"github.com/nurhidaylma/parking-app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *parkingService) RegisterVehicle(req model.RegisterVehicleRequest) (model.RegisterVehicleResponse, error) {
	timeNow := time.Now()

	// get available parking spot
	parkingSpots, err := s.repo.ReadParkingSpots()
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.RegisterVehicleResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}
	availableSpot, index := findNextAvailableSpot(parkingSpots)
	if availableSpot == "" {
		util.Logger.LogWarning("parking lot is full")
		return model.RegisterVehicleResponse{}, status.Error(codes.FailedPrecondition, "parking lot is full")
	}

	vehiclesParked, err := s.repo.ReadVehicles()
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.RegisterVehicleResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}
	for _, v := range vehiclesParked {
		if v.Number == req.VehicleNumber {
			util.Logger.LogWarning("vehicle already registered")
			return model.RegisterVehicleResponse{}, status.Error(codes.InvalidArgument, "vehicle already registered")
		}
	}

	// mark parking spot to be unaivailable
	parkingSpots = markParkingSpotAvailability(parkingSpots, index, false)
	err = s.repo.WriteParkingSpots(parkingSpots)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.RegisterVehicleResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}

	vehiclesParked = append(vehiclesParked, model.Vehicle{
		Number:    req.VehicleNumber,
		Color:     req.VehicleColor,
		Type:      util.VehicleTypes(req.VehicleType),
		EntryTime: timeNow,
		LotNumber: availableSpot,
	})
	err = s.repo.WriteVehicles(vehiclesParked)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.RegisterVehicleResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}

	return model.RegisterVehicleResponse{
		VehicleNumber: req.VehicleNumber,
		LotNumber:     availableSpot,
		EntryTime:     timeNow.Format(time.DateTime),
	}, nil
}

func findNextAvailableSpot(parkingSpots []model.ParkingSpot) (string, int) {
	for i, spot := range parkingSpots {
		if spot.IsAvailable {
			return spot.Number, i
		}
	}
	return "", -1
}

func markParkingSpotAvailability(parkingSpots []model.ParkingSpot, index int, availability bool) []model.ParkingSpot {
	parkingSpots[index].IsAvailable = availability
	return parkingSpots
}
