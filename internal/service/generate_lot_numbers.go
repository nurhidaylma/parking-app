package service

import (
	"fmt"

	"github.com/nurhidaylma/parking-app/internal/model"
	"github.com/nurhidaylma/parking-app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *parkingService) GenerateLotNumbers() error {
	var parkingNumbers []model.ParkingSpot
	letter := 'A'
	count := 0

	for i := 0; i < s.parkingCapacity; i++ {
		number := (count % 10) + 1
		parkingNumber := fmt.Sprintf("%c%d", letter, number)
		parkingNumbers = append(parkingNumbers, model.ParkingSpot{
			Number:      parkingNumber,
			IsAvailable: true,
		})
		count++
		if number == 10 {
			letter++
		}
	}

	// write parking spot
	err := s.repo.WriteParkingSpots(parkingNumbers)
	if err != nil {
		util.Logger.LogError(err.Error())
		return status.Error(codes.Internal, codes.Internal.String())
	}

	// write initial vehicle
	err = s.repo.WriteInitialVehicle()
	if err != nil {
		util.Logger.LogError(err.Error())
		return status.Error(codes.Internal, codes.Internal.String())
	}

	return nil
}
