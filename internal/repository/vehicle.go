package repository

import (
	"encoding/json"
	"os"

	"github.com/nurhidaylma/parking-app/internal/model"
	"github.com/nurhidaylma/parking-app/util"
)

func (r *repository) ReadVehicles() ([]model.Vehicle, error) {
	const (
		fileName = "vehicle.go"
		funcName = "ReadVehicles"
	)

	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.Open(r.vehicleFileName)
	if err != nil {
		return nil, util.NewError(fileName, funcName, err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, util.NewError(fileName, funcName, err)
	}
	if fileInfo.Size() == 0 {
		return nil, util.NewError(fileName, funcName, err)
	}

	var vehicles []model.Vehicle
	err = json.NewDecoder(file).Decode(&vehicles)
	if err != nil {
		return nil, util.NewError(fileName, funcName, err)
	}

	return vehicles, nil
}

func (r *repository) WriteVehicles(vehicles []model.Vehicle) error {
	const (
		fileName = "vehicle.go"
		funcName = "WriteVehicles"
	)

	r.mutex.Lock()
	defer r.mutex.Unlock()

	data, err := json.MarshalIndent(vehicles, "", "  ")
	if err != nil {
		return util.NewError(fileName, funcName, err)
	}

	return os.WriteFile(r.vehicleFileName, data, 0644)
}

func (r *repository) WriteInitialVehicle() error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return os.WriteFile(r.vehicleFileName, []byte{}, 0644)
}
