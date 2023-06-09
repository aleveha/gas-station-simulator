package station

import (
	"fmt"
	"gas-station-simulator/src/car"
	"gas-station-simulator/src/constants"
	"gas-station-simulator/src/fuel"
	"math/rand"
	"sync"
	"time"
)

type GasStation struct {
	Checkouts chan *car.Car
	Pumps     map[fuel.Fuel]Pump
	wg        *sync.WaitGroup
}

func (gs GasStation) runPumps() {
	for _, pump := range gs.Pumps {
		go pump.Run(gs.Checkouts)
	}
}

func (gs GasStation) proceedCheckout() {
	for chC := range gs.Checkouts {
		time.Sleep(2 * constants.Time)
		chC.LeftAt = time.Now()
		fmt.Printf("✅   %v left the station after checkout at %v\n", chC, time.Now().Format("15:04:05.000"))
		gs.wg.Done()
	}
}

func (gs GasStation) addCarToQueue(c *car.Car) {
	select {
	case gs.Pumps[c.FuelType].Queue <- c:
		c.ArrivedAtPump = time.Now()
		fmt.Printf("🕓  %v arrived at %v pump queue\n", c, c.FuelType)
	case <-time.After(constants.Time / 1000):
		c.LeftAt = time.Now()
		fmt.Printf("❌  %v left the station because of long wating time at %v\n", c, time.Now().Format("15:04:05.000"))
		gs.wg.Done()
	}
}

func (gs GasStation) start(cars []*car.Car) {
	for _, c := range cars {
		c.ArrivedAtStation = time.Now()
		fmt.Printf("🚗  %v arrived to a gas station at %v\n", c, time.Now().Format("15:04:05.000"))
		go gs.addCarToQueue(c)
		time.Sleep(time.Duration(rand.Intn(6)) * constants.Time) // simulate time between car arrivals
	}
}

func (gs GasStation) Run(cars []*car.Car) {
	gs.wg = &sync.WaitGroup{}
	gs.wg.Add(len(cars))

	gs.runPumps()
	go gs.proceedCheckout()
	go gs.start(cars)

	gs.wg.Wait()
}
