package model

import (
	"time"

	"github.com/nurhidaylma/parking-app/util"
)

type Vehicle struct {
	Number    string            `json:"number"`
	Color     string            `json:"color"`
	Type      util.VehicleTypes `json:"type"`
	EntryTime time.Time         `json:"entry_time"`
	LotNumber string            `json:"lot_number"`
}

type ParkingTicket struct {
	VehicleNumber string `json:"vehicle_number"`
	LotNumber     string `json:"lot_number"`
}

type ParkingSpot struct {
	Number      string `json:"number"`
	IsAvailable bool   `json:"is_available"`
}

type RegisterVehicleRequest struct {
	VehicleNumber string `json:"vehicle_number"`
	VehicleColor  string `json:"vehicle_color"`
	VehicleType   string `json:"vehicle_type"`
}

type RegisterVehicleResponse struct {
	VehicleNumber string `json:"vehicle_number"`
	LotNumber     string `json:"lot_number"`
	EntryTime     string `json:"entry_time"`
}

type UnregisterVehicleRequest struct {
	Number string `json:"vehicle_number"`
}

type UnregisterVehicleResponse struct {
	Number     string `json:"vehicle_number"`
	EntryTime  string `json:"entry_time"`
	ExitTime   string `json:"exit_time"`
	ParkingFee string `json:"parking_fee"`
}

type GetVehiclesByTypeRequest struct {
	VehicleType string `json:"vehicle_type"`
}

type GetVehiclesByTypeResponse struct {
	TotalVehicle int `json:"total_vehicle"`
}

type GetVehiclesByColorRequest struct {
	VehicleColor string `json:"vehicle_color"`
}

type GetVehiclesByColorResponse struct {
	VehicleNumbers []string `json:"vehicle_numbers"`
}
