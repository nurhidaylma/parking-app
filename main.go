package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/nurhidaylma/parking-app/internal/endpoint"
	"github.com/nurhidaylma/parking-app/internal/repository"
	"github.com/nurhidaylma/parking-app/internal/service"
	"github.com/nurhidaylma/parking-app/internal/transport"
	"github.com/nurhidaylma/parking-app/util"
)

func main() {
	// check if any command-line arguments were provided
	if len(os.Args) < 2 {
		fmt.Println("please enter parking capacity: go run main.go [parking_capacity]")
		return
	}
	// extract the input from the command line
	inputStr := os.Args[1]
	// convert the input to an integer
	parkingCapacity, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println("input must be an integer")
		return
	}

	// initiate logger
	loggerInstance, err := util.NewCustomLogger("logfile.log")
	if err != nil {
		log.Fatal("failed to create logger: ", err.Error())
	}
	util.Logger = loggerInstance

	// get path
	path, err := os.Getwd()
	if err != nil {
		util.Logger.LogError(err.Error())
		return
	}
	vehicleFileName := path + "/config/vehicles.json"
	parkingLotFileName := path + "/config/parking_lots.json"

	repository := repository.NewRepository(vehicleFileName, parkingLotFileName)
	service := service.NewParkingService(repository, parkingCapacity)
	endpoint := endpoint.NewParkingEndpoints(service)

	// initialize parking lot numbers
	service.GenerateLotNumbers()

	handler := transport.NewHTTPHandler(endpoint)
	log.Println("Starting server at :9700")
	if err := http.ListenAndServe(":9700", handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
