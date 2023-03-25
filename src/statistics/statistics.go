package statistics

import (
	"fmt"
	"gas-station-simulator/src/car"
	"time"
)

func formatDateTime(t time.Time) string {
	return t.Format("02.01.2006 15:04:05")
}

func stationLineTime(c *car.Car) time.Duration {
	return c.ArrivedAtPump.Sub(c.ArrivedAtStation).Round(time.Millisecond)
}

func refuelTime(c *car.Car) time.Duration {
	return c.RefueledAt.Sub(c.ArrivedAtPump).Round(time.Second)
}

func checkoutLineTime(c *car.Car) time.Duration {
	return c.LeftAt.Sub(c.RefueledAt).Round(time.Second)
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

func Calculate(cars []*car.Car) {
	successfullyRefueled, notRefueledCars := filterCars(cars, func(c *car.Car) bool {
		return !c.RefueledAt.IsZero()
	})
	fmt.Printf("%v cars left station because of long waiting time!\n", len(notRefueledCars))
	fmt.Printf("%v cars was successfully refueled!\n\n", len(successfullyRefueled))

	for _, c := range successfullyRefueled {
		fmt.Println(getCarInfo(c))
	}
}
