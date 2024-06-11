package service

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/nurhidaylma/parking-app/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_GetVehiclesByType(t *testing.T) {
	initTest(t)

	tests := []struct {
		testName string
		request  model.GetVehiclesByTypeRequest

		readVehiclesResp []model.Vehicle
		readVehiclesErr  error

		expectedResp model.GetVehiclesByTypeResponse
		expectedErr  error
	}{
		{
			testName: "TC 1: Successful get vehicles by type",
			request: model.GetVehiclesByTypeRequest{
				VehicleType: "SUV",
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
			expectedResp: model.GetVehiclesByTypeResponse{
				TotalVehicle: 1,
			},
			expectedErr: nil,
		},
		{
			testName: "TC 2: Wrong vehicle type",
			request: model.GetVehiclesByTypeRequest{
				VehicleType: "MMP",
			},
			expectedResp: model.GetVehiclesByTypeResponse{
				TotalVehicle: 0,
			},
			expectedErr: status.Error(codes.FailedPrecondition, "invalid vehicle type"),
		},
		{
			testName: "TC 3: No matching vehicle type",
			request: model.GetVehiclesByTypeRequest{
				VehicleType: "MPV",
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
			expectedResp:    model.GetVehiclesByTypeResponse{TotalVehicle: 0},
			expectedErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockRepo.EXPECT().ReadVehicles().Return(tt.readVehiclesResp, tt.readVehiclesErr)

			gotResp, err := services.GetVehiclesByType(tt.request)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Parking.GetVehiclesByType() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.expectedResp) {
				t.Logf("Parking.GetVehiclesByType() Response = %v, Login = %v", gotResp, tt.expectedResp)
				return
			}
		})
	}
}
