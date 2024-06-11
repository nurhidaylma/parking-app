package service

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/nurhidaylma/parking-app/internal/model"
)

func Test_GetVehiclesByColor(t *testing.T) {
	initTest(t)

	tests := []struct {
		testName string
		request  model.GetVehiclesByColorRequest

		readVehiclesResp []model.Vehicle
		readVehiclesErr  error

		expectedResp model.GetVehiclesByColorResponse
		expectedErr  error
	}{
		{
			testName: "TC 1: Successful get vehicles by color",
			request: model.GetVehiclesByColorRequest{
				VehicleColor: "hitam",
			},
			readVehiclesResp: []model.Vehicle{
				{
					Number:    "T 123 TEST",
					Color:     "hitam",
					Type:      "SUV",
					EntryTime: time.Now(),
					LotNumber: "A1",
				},
			},
			readVehiclesErr: nil,
			expectedResp: model.GetVehiclesByColorResponse{
				VehicleNumbers: []string{"T 123 TEST"},
			},
			expectedErr: nil,
		},
		{
			testName: "TC 2: No mathcing color",
			request: model.GetVehiclesByColorRequest{
				VehicleColor: "hitam",
			},
			readVehiclesResp: []model.Vehicle{
				{
					Number:    "T 123 TEST",
					Color:     "hijau",
					Type:      "SUV",
					EntryTime: time.Now(),
					LotNumber: "A1",
				},
			},
			readVehiclesErr: nil,
			expectedResp:    model.GetVehiclesByColorResponse{},
			expectedErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockRepo.EXPECT().ReadVehicles().Return(tt.readVehiclesResp, tt.readVehiclesErr)

			gotResp, err := services.GetVehiclesByColor(tt.request)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Parking.GetVehiclesByColor() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.expectedResp) {
				t.Logf("Parking.GetVehiclesByColor() Response = %v, Login = %v", gotResp, tt.expectedResp)
				return
			}
		})
	}

}
