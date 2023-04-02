package main

import (
	"fmt"
	"gas-station-simulator/src/car"
	"gas-station-simulator/src/constants"
	"gas-station-simulator/src/station"
	"gas-station-simulator/src/statistics"
	"time"
)

func main() {
	cars := car.GenerateCars(constants.CarsCount)
	gasStation := station.GasStation{
		Checkouts: make(chan *car.Car),
		Pumps:     station.GeneratePumps(),
	}

	println("\nStarting simulation..\n")
	start := time.Now()
	gasStation.Run(cars)
	end := time.Now()

	fmt.Println("\nSimulation finished, printing statistics..")
	fmt.Printf("Simulation took %v\n", end.Sub(start).Round(constants.Time))
	statistics.Calculate(cars)
}
