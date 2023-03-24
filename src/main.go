package main

import (
	"gas-station-simulator/src/car"
	"gas-station-simulator/src/station"
)

func main() {
	cars := car.GenerateCars(10)
	gasStation := station.GasStation{
		Checkouts: make(chan *car.Car),
		Pumps:     station.GeneratePumps(),
	}

	gasStation.Run(cars)
}
