package repository

import (
	"sync"

	"github.com/nurhidaylma/parking-app/internal/model"
)

type repository struct {
	vehicleFileName    string
	parkingLotFileName string
	mutex              sync.Mutex
}

func NewRepository(vehicleFileName, parkingLotFileName string) RepoInterface {
	return &repository{vehicleFileName: vehicleFileName, parkingLotFileName: parkingLotFileName}
}

type RepoInterface interface {
	WriteParkingSpots([]model.ParkingSpot) error
	ReadParkingSpots() ([]model.ParkingSpot, error)

	ReadVehicles() ([]model.Vehicle, error)
	WriteVehicles([]model.Vehicle) error
	WriteInitialVehicle() error
}
