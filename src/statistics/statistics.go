package statistics

import (
	"fmt"
	"gas-station-simulator/src/car"
	"gas-station-simulator/src/constants"
	"gas-station-simulator/src/fuel"
	"time"
)

func formatDateTime(t time.Time) string {
	return t.Format("02.01.2006 15:04:05")
}

func stationLineTime(c *car.Car) time.Duration {
	return c.ArrivedAtPump.Sub(c.ArrivedAtStation).Round(constants.Time)
}

func refuelTime(c *car.Car) time.Duration {
	return c.RefueledAt.Sub(c.ArrivedAtPump).Round(constants.Time)
}

func checkoutLineTime(c *car.Car) time.Duration {
	return c.LeftAt.Sub(c.RefueledAt).Round(constants.Time)
}

func getCarInfo(c *car.Car) string {
	allInfo := fmt.Sprintf("%v\n", c)
	waitingInLine := stationLineTime(c)
	refueling := refuelTime(c)
	checkoutTime := checkoutLineTime(c)

	allInfo += fmt.Sprintf("Arrived at station: %v\n", formatDateTime(c.ArrivedAtStation))
	allInfo += fmt.Sprintf("Arrived at pump: %v", formatDateTime(c.ArrivedAtPump))
	if waitingInLine > 0 {
		allInfo += fmt.Sprintf(" (waiting time: %v)\n", waitingInLine)
	} else {
		allInfo += "\n"
	}
	allInfo += fmt.Sprintf("Refueled at: %v (waiting + refueling time: %v)\n", formatDateTime(c.RefueledAt), refueling)
	allInfo += fmt.Sprintf("Left at: %v (time in checkout line: %v)\n", formatDateTime(c.LeftAt), checkoutTime)
	return allInfo
}

func filterCars(cars []*car.Car, filter func(*car.Car) bool) ([]*car.Car, []*car.Car) {
	var filteredCars []*car.Car
	var remainingCars []*car.Car
	for _, c := range cars {
		if filter(c) {
			filteredCars = append(filteredCars, c)
		} else {
			remainingCars = append(remainingCars, c)
		}
	}
	return filteredCars, remainingCars
}

func avgDuration(durations []time.Duration) time.Duration {
	var sum time.Duration
	for _, d := range durations {
		sum += d
	}
	return sum / time.Duration(len(durations))
}

func avgRefuelTimeByFuelType(cars []*car.Car, fuelType fuel.Fuel) time.Duration {
	var refuelTimes []time.Duration
	for _, c := range cars {
		if c.FuelType == fuelType {
			refuelTimes = append(refuelTimes, refuelTime(c))
		}
	}
	return avgDuration(refuelTimes)
}

func avgTimeAtStation(cars []*car.Car) time.Duration {
	var timesAtStation []time.Duration
	for _, c := range cars {
		timesAtStation = append(timesAtStation, c.LeftAt.Sub(c.ArrivedAtStation).Round(constants.Time))
	}
	return avgDuration(timesAtStation)
}

func avgTimeAtCheckoutLine(cars []*car.Car) time.Duration {
	var timesAtCheckoutLine []time.Duration
	for _, c := range cars {
		timesAtCheckoutLine = append(timesAtCheckoutLine, c.LeftAt.Sub(c.RefueledAt).Round(constants.Time))
	}
	return avgDuration(timesAtCheckoutLine)
}

func Calculate(cars []*car.Car) {
	successfullyRefueled, notRefueledCars := filterCars(cars, func(c *car.Car) bool {
		return !c.RefueledAt.IsZero()
	})
	fmt.Printf("%v cars left station because of long waiting time!\n", len(notRefueledCars))
	fmt.Printf("%v cars was successfully refueled!\n\n", len(successfullyRefueled))

	fmt.Printf("Average refueling time for %v cars: %v\n", fuel.Gas, avgRefuelTimeByFuelType(successfullyRefueled, fuel.Gas))
	fmt.Printf("Average refueling time for %v cars: %v\n", fuel.Diesel, avgRefuelTimeByFuelType(successfullyRefueled, fuel.Diesel))
	fmt.Printf("Average refueling time for %v cars: %v\n", fuel.LPG, avgRefuelTimeByFuelType(successfullyRefueled, fuel.LPG))
	fmt.Printf("Average refueling time for %v cars: %v\n", fuel.Electric, avgRefuelTimeByFuelType(successfullyRefueled, fuel.Electric))
	fmt.Printf("Average time at checkout line for %v cars: %v\n", len(successfullyRefueled), avgTimeAtCheckoutLine(successfullyRefueled))
	fmt.Printf("Average time at station for %v cars: %v\n\n", len(successfullyRefueled), avgTimeAtStation(successfullyRefueled))

	for _, c := range successfullyRefueled {
		fmt.Println(getCarInfo(c))
	}
}
