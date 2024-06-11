package service

import (
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	repoMock "github.com/nurhidaylma/parking-app/internal/repository"
	"github.com/nurhidaylma/parking-app/util"
)

var (
	mockRepo *repoMock.MockRepoInterface

	repos    *parkingService
	services ParkingServiceInterface
)

func initTest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	loggerInstance, err := util.NewCustomLogger("logfile_test.log")
	if err != nil {
		log.Fatal("failed to create logger: ", err.Error())
	}
	util.Logger = loggerInstance

	mockRepo = repoMock.NewMockRepoInterface(mockCtrl)

	repos = &parkingService{
		repo:            mockRepo,
		parkingCapacity: 5,
	}

	services = NewParkingService(mockRepo, 5)
}
