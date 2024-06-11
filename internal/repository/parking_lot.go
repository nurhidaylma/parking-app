package repository

import (
	"encoding/json"
	"os"

	"github.com/nurhidaylma/parking-app/internal/model"
	"github.com/nurhidaylma/parking-app/util"
)

func (r *repository) WriteParkingSpots(req []model.ParkingSpot) error {
	const (
		fileName = "parking_lot.go"
		funcName = "WriteParkingSpots"
	)

	r.mutex.Lock()
	defer r.mutex.Unlock()

	data, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		return util.NewError(fileName, funcName, err)
	}

	err = os.WriteFile(r.parkingLotFileName, data, 0644)
	if err != nil {
		return util.NewError(fileName, funcName, err)
	}

	return nil
}

func (r *repository) ReadParkingSpots() ([]model.ParkingSpot, error) {
	const (
		fileName = "parking_lot.go"
		funcName = "ReadParkingSpots"
	)

	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.Open(r.parkingLotFileName)
	if err != nil {
		return nil, util.NewError(fileName, funcName, err)
	}
	defer file.Close()

	var parkingSpots []model.ParkingSpot
	err = json.NewDecoder(file).Decode(&parkingSpots)
	if err != nil {
		return nil, util.NewError(fileName, funcName, err)
	}

	return parkingSpots, nil
}
