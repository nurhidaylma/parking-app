package service

import (
	"github.com/nurhidaylma/parking-app/internal/model"
	"github.com/nurhidaylma/parking-app/internal/repository"
)

type parkingService struct {
	repo            repository.RepoInterface
	parkingCapacity int
}

func NewParkingService(repo repository.RepoInterface, parkingCapacity int) ParkingServiceInterface {
	return &parkingService{repo: repo, parkingCapacity: parkingCapacity}
}

type ParkingServiceInterface interface {
	GenerateLotNumbers() error

	RegisterVehicle(model.RegisterVehicleRequest) (model.RegisterVehicleResponse, error)
	UnregisterVehicle(model.UnregisterVehicleRequest) (model.UnregisterVehicleResponse, error)
	GetVehiclesByType(model.GetVehiclesByTypeRequest) (model.GetVehiclesByTypeResponse, error)
	GetVehiclesByColor(model.GetVehiclesByColorRequest) (model.GetVehiclesByColorResponse, error)
}
