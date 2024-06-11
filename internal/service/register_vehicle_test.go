package service

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/nurhidaylma/parking-app/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_RegisterVehicle(t *testing.T) {
	initTest(t)

	tests := []struct {
		testName string
		request  model.RegisterVehicleRequest

		readParkingSpotsResp []model.ParkingSpot
		readParkingSpotsErr  error

		readVehiclesResp []model.Vehicle
		readVehiclesErr  error

		writeParkingSpotsErr error
		writeVehiclesErr     error

		expectedResp model.RegisterVehicleResponse
		expectedErr  error
	}{
		{
			testName: "TC 1: Successful register vehicle",
			request: model.RegisterVehicleRequest{
				VehicleNumber: "T 123 TEST",
				VehicleColor:  "hitam",
				VehicleType:   "SUV",
			},
			readParkingSpotsResp: []model.ParkingSpot{
				{
					Number:      "A1",
					IsAvailable: true,
				},
				{
					Number:      "A2",
					IsAvailable: true,
				},
				{
					Number:      "A3",
					IsAvailable: true,
				},
				{
					Number:      "A4",
					IsAvailable: true,
				},
				{
					Number:      "A5",
					IsAvailable: true,
				},
			},
			readParkingSpotsErr:  nil,
			readVehiclesResp:     []model.Vehicle{},
			readVehiclesErr:      nil,
			writeParkingSpotsErr: nil,
			writeVehiclesErr:     nil,
			expectedResp: model.RegisterVehicleResponse{
				VehicleNumber: "T 123 TEST",
				LotNumber:     "A1",
				EntryTime:     time.Now().Format(time.DateTime),
			},
			expectedErr: nil,
		},
		{
			testName: "TC 2: Parking lot is full",
			request: model.RegisterVehicleRequest{
				VehicleNumber: "T 123 TEST",
				VehicleColor:  "hitam",
				VehicleType:   "SUV",
			},
			readParkingSpotsResp: []model.ParkingSpot{
				{
					Number:      "A1",
					IsAvailable: false,
				},
				{
					Number:      "A2",
					IsAvailable: false,
				},
				{
					Number:      "A3",
					IsAvailable: false,
				},
				{
					Number:      "A4",
					IsAvailable: false,
				},
				{
					Number:      "A5",
					IsAvailable: false,
				},
			},
			readParkingSpotsErr: nil,
			expectedResp:        model.RegisterVehicleResponse{},
			expectedErr:         status.Error(codes.FailedPrecondition, "parking lot is full"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockRepo.EXPECT().ReadParkingSpots().Return(tt.readParkingSpotsResp, tt.readParkingSpotsErr)
			mockRepo.EXPECT().ReadVehicles().Return(tt.readVehiclesResp, tt.readVehiclesErr)
			mockRepo.EXPECT().WriteParkingSpots(gomock.Any()).Return(tt.writeParkingSpotsErr)
			mockRepo.EXPECT().WriteVehicles(gomock.Any()).Return(tt.writeVehiclesErr)

			gotResp, err := services.RegisterVehicle(tt.request)
			t.Logf(gotResp.LotNumber)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Parking.RegisterVehicle() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.expectedResp) {
				t.Logf("Parking.RegisterVehicle() Response = %v, Login = %v", gotResp, tt.expectedResp)
				return
			}
		})
	}
}
