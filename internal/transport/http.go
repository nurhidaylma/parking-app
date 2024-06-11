package transport

import (
	"net/http"

	"github.com/nurhidaylma/parking-app/internal/endpoint"
)

type ParkingTransport struct {
	endpoint endpoint.ParkingEndpoints
}

func NewHTTPHandler(ep endpoint.ParkingEndpoints) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", ep.RegisterVehicle)
	mux.HandleFunc("/unregister", ep.UnregisterVehicle)
	mux.HandleFunc("/vehicle/type", ep.GetVehiclesByType)
	mux.HandleFunc("/vehicle/color", ep.GetVehiclesByColor)

	return mux
}
