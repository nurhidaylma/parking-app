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

func Test_UnregisterVehicle(t *testing.T) {
	initTest(t)

	tests := []struct {
		testName string
		request  model.UnregisterVehicleRequest

		readVehiclesResp []model.Vehicle
		readVehiclesErr  error

		writeVehiclesErr error

		readParkingSpotsResp []model.ParkingSpot
		readParkingSpotsErr  error

		writeParkingSpotsErr error

		expectedResp model.UnregisterVehicleResponse
		expectedErr  error
	}{
		{
			testName: "TC 1: Successful unregister vehicle",
			request: model.UnregisterVehicleRequest{
				Number: "T 123 TEST",
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
			readVehiclesErr:  nil,
			writeVehiclesErr: nil,
			readParkingSpotsResp: []model.ParkingSpot{
				{
					Number:      "A1",
					IsAvailable: false,
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
			writeParkingSpotsErr: nil,
			expectedResp: model.UnregisterVehicleResponse{
				Number:     "T 123 TEST",
				EntryTime:  time.Now().Format(time.DateTime),
				ExitTime:   time.Now().Add(time.Hour * 2).Format(time.DateTime),
				ParkingFee: "30000",
			},
			expectedErr: nil,
		},
		{
			testName: "TC 2: Vehicle is not registered",
			request: model.UnregisterVehicleRequest{
				Number: "T 122 TEST",
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
			expectedResp: model.UnregisterVehicleResponse{},
			expectedErr:  status.Error(codes.InvalidArgument, "vehicle is not registered"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockRepo.EXPECT().ReadVehicles().Return(tt.readVehiclesResp, tt.readVehiclesErr)
			mockRepo.EXPECT().WriteVehicles(gomock.Any()).Return(tt.writeVehiclesErr)
			mockRepo.EXPECT().ReadParkingSpots().Return(tt.readParkingSpotsResp, tt.readParkingSpotsErr)
			mockRepo.EXPECT().WriteParkingSpots(gomock.Any()).Return(tt.writeParkingSpotsErr)

			gotResp, err := services.UnregisterVehicle(tt.request)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Parking.UnregisterVehicle() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.expectedResp) {
				t.Logf("Parking.UnregisterVehicle() Response = %v, Login = %v", gotResp, tt.expectedResp)
				return
			}
		})
	}
}
