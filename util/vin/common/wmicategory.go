package common

type WMICategory int

const (
	NotSpecified WMICategory = iota
	PassengerCar
	MPV
	Truck
	Bus
	Trailer
	Motorcycle
	ATV
	IncompleteCar
	BasicChassis
	LowSpeedVehicle
)
