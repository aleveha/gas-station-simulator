package car

import (
	"fmt"
	"gas-station-simulator/src/fuel"
	"math/rand"
	"time"
)

type Car struct {
	ArrivedAtStation time.Time
	ArrivedAtPump    time.Time
	FuelType         fuel.Fuel
	Id               int
	LeftAt           time.Time
	RefuelDuration   time.Duration
	RefueledAt       time.Time
}

func (car Car) Refuel() {
	time.Sleep(car.RefuelDuration)
}

func (car Car) String() string {
	return fmt.Sprintf("%v car (id: %d)", car.FuelType, car.Id)
}

func randomFuelType() fuel.Fuel {
	return fuel.Types[rand.Intn(len(fuel.Types))]
}

func randomRefuelTime() int {
	return rand.Intn(3) + 3
}

func GenerateCars(count int) []*Car {
	var cars []*Car
	for i := 0; i < count; i++ {
		carType := randomFuelType()
		fillUpDuration := randomRefuelTime()
		if carType == fuel.Electric {
			fillUpDuration *= 3
		}
		cars = append(cars, &Car{
			Id:             i + 1,
			RefuelDuration: time.Second * time.Duration(fillUpDuration),
			FuelType:       carType,
		})
	}
	return cars
}
