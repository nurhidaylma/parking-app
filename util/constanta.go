package util

const (
	AddinitionalRate = 0.2
)

type VehicleTypes string

const (
	VehicleTypeSUV VehicleTypes = "SUV"
	VehicleTypeMPV VehicleTypes = "MPV"
)

type VehicleTypeFare int64

const (
	VehicleTypeFareSUV VehicleTypeFare = 25000
	VehicleTypeFareMPV VehicleTypeFare = 35000
)
