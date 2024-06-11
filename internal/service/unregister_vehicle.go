package service

import (
	"fmt"
	"time"

	"github.com/nurhidaylma/parking-app/internal/model"
	"github.com/nurhidaylma/parking-app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *parkingService) UnregisterVehicle(req model.UnregisterVehicleRequest) (model.UnregisterVehicleResponse, error) {
	var vehicleIndex int
	var vehicle model.Vehicle
	timeNow := time.Now()

	vehiclesParked, err := s.repo.ReadVehicles()
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.UnregisterVehicleResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}
	for i, v := range vehiclesParked {
		if v.Number == req.Number {
			vehicleIndex = i
			vehicle = v

			break
		}
	}
	if vehicle.IsEmpty() {
		util.Logger.LogWarning("vehicle is not registered")
		return model.UnregisterVehicleResponse{}, status.Error(codes.InvalidArgument, "vehicle is not registered")
	}

	// calculate parking fee
	var parkingFee int64
	durationInHours := timeNow.Sub(vehicle.EntryTime).Hours()
	parkingFee = calculateParkingFee(vehicle.Type, durationInHours)
	parkingFeeInString := fmt.Sprintf("%d", parkingFee)

	// remove vehicle from parking
	vehiclesParked = append(vehiclesParked[:vehicleIndex], vehiclesParked[vehicleIndex+1:]...)
	err = s.repo.WriteVehicles(vehiclesParked)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.UnregisterVehicleResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}

	// mark parking lot number to be available
	var parkingIndex int
	parkingSpots, err := s.repo.ReadParkingSpots()
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.UnregisterVehicleResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}
	for i, p := range parkingSpots {
		if p.Number == vehicle.LotNumber {
			parkingIndex = i
		}
	}
	parkingSpots = markParkingSpotAvailability(parkingSpots, parkingIndex, true)
	err = s.repo.WriteParkingSpots(parkingSpots)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.UnregisterVehicleResponse{}, status.Error(codes.Internal, codes.Internal.String())
	}

	return model.UnregisterVehicleResponse{
		Number:     vehicle.Number,
		EntryTime:  vehicle.EntryTime.Format(time.DateTime),
		ExitTime:   timeNow.Format(time.DateTime),
		ParkingFee: parkingFeeInString,
	}, nil
}

func calculateParkingFee(vehicleType util.VehicleTypes, duration float64) int64 {
	var baseFee int64
	var additionalFee float64

	switch vehicleType {
	case util.VehicleTypeSUV:
		baseFee = int64(util.VehicleTypeFareSUV)
	case util.VehicleTypeMPV:
		baseFee = int64(util.VehicleTypeFareMPV)
	}

	if duration <= 1 {
		return baseFee
	}
	additionalFee = float64(baseFee) * util.AddinitionalRate

	return baseFee + int64(additionalFee)*(int64(duration)-1)
}
