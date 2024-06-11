package endpoint

import (
	"encoding/json"

	"net/http"

	"github.com/nurhidaylma/parking-app/internal/model"
	"github.com/nurhidaylma/parking-app/internal/service"
)

type ParkingEndpoints struct {
	service service.ParkingServiceInterface
}

func NewParkingEndpoints(service service.ParkingServiceInterface) ParkingEndpoints {
	return ParkingEndpoints{service: service}
}

func (pe *ParkingEndpoints) RegisterVehicle(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := pe.service.RegisterVehicle(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (pe *ParkingEndpoints) UnregisterVehicle(w http.ResponseWriter, r *http.Request) {
	var req model.UnregisterVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := pe.service.UnregisterVehicle(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (pe *ParkingEndpoints) GetVehiclesByType(w http.ResponseWriter, r *http.Request) {
	var req model.GetVehiclesByTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := pe.service.GetVehiclesByType(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (pe *ParkingEndpoints) GetVehiclesByColor(w http.ResponseWriter, r *http.Request) {
	var req model.GetVehiclesByColorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := pe.service.GetVehiclesByColor(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
