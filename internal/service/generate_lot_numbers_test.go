package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_GenerateLotNumbers(t *testing.T) {
	initTest(t)

	tests := []struct {
		testName               string
		writeParkingSpotsErr   error
		writeInitialVehicleErr error
		expectedErr            error
	}{
		{
			testName:               "TC 1: Successful generate parking lot numbers",
			writeParkingSpotsErr:   nil,
			writeInitialVehicleErr: nil,
			expectedErr:            nil,
		},
		{
			testName:             "TC 2: Error writing parking lot numbers",
			writeParkingSpotsErr: errors.New("error"),
			expectedErr:          status.Error(codes.Internal, codes.Internal.String()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			mockRepo.EXPECT().WriteParkingSpots(gomock.Any()).Return(tt.writeParkingSpotsErr)
			mockRepo.EXPECT().WriteInitialVehicle().Return(tt.writeInitialVehicleErr)

			err := services.GenerateLotNumbers()
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Parking.GenerateLotNumbers() Error = %v, WantError = %v", err, tt.expectedErr)
				return
			}
		})
	}
}
