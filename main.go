package main

import (
	carPackage "gas-station-simulator/src"
	"math/rand"
	"sync"
	"time"
)

func generateCars(count int) []carPackage.Car {
	var cars []carPackage.Car
	for i := 0; i < count; i++ {
		carType := carPackage.GetCarType(rand.Intn(4))
		fillUpDuration := rand.Intn(5) + 1
		if carType == carPackage.Electric {
			fillUpDuration *= 3
		}
		cars = append(cars, carPackage.Car{
			FillUpDuration: time.Second * time.Duration(fillUpDuration),
			Id:             i + 1,
			Type:           carType,
		})
	}
	return cars
}

func main() {
	cars := generateCars(5)

	wg := sync.WaitGroup{}
	wg.Add(len(cars))

	for _, car := range cars {
		car := car
		go func() {
			defer wg.Done()
			car.FillUp()
		}()
	}

	wg.Wait()
}
