# Parking Lot Management System
## Overview
The Parking Lot Management System is a simple application to manage a parking lot for cars, specifically designed to handle the operations of parking, retrieving, and generating reports about parked vehicles. The system is implemented in Go and does not use a traditional database; instead, it uses a file-based repository for data storage.

## Features
1. Register Vehicle: Register a vehicle entering the parking lot and assign it a parking slot.
2. Unregister Vehicle: Calculate parking fees and vacate the slot upon vehicle exit.
3. Generate Report: Obtain a summary of the number of vehicles by type.
4. Get Vehicles by Color: List vehicle numbers filtered by their color.

## Project Structure
parking-lot/
├── config/
│   └── parking_lots.json             # Storage of parking lots
│   └── vehicles.json                 # Storage of vehicles parked
├── internal/
│   ├── endpoints/
│   │   └── endpoints.go              # API endpoints
│   ├── model/
│   │   └── model.go                  # Data models
│   ├── repository/
│   │   ├── Makefile                  # Script to mocking
│   │   ├── repository.go             # Repository interface
│   │   ├── vehicle.go                # File-based repository implementation for vehicle 
│   │   └── parking_lot.go            # File-based repository implementation for parking lot
│   ├── service/
│   │   ├── service.go                # Business logic
│   │   └── service_test.go           # Unit tests for service layer
│   └── transport/
│       └── http.go                   # HTTP handlers      
└── go.mod                            # Go module file

## Setup

To run the application locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/nurhidaylma/parking-app.git
   cd parking-app

2. Install dependencies:
   ```bash
   go mod tidy

3. Start the server:
   ```bash
   go run main.go [insert parking-capacity here]