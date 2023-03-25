package station

import (
	"fmt"
	"gas-station-simulator/src/car"
	"gas-station-simulator/src/fuel"
	"math/rand"
	"sync"
	"time"
)

type GasStation struct {
	Checkouts chan *car.Car
	Pumps     map[fuel.Fuel]Pump
}

func (gs GasStation) runPumps() {
	for _, pump := range gs.Pumps {
		go pump.Run(gs.Checkouts)
	}
}

func (gs GasStation) proceedCheckout(wg *sync.WaitGroup) {
	for chC := range gs.Checkouts {
		time.Sleep(time.Second)
		chC.LeftAt = time.Now()
		fmt.Printf("✅   %v left the station after checkout at %v\n", chC, time.Now().Format("15:04:05.000"))
		wg.Done()
	}
}

func (gs GasStation) addCarToQueue(c *car.Car, wg *sync.WaitGroup) {
	select {
	case gs.Pumps[c.FuelType].Queue <- c:
		c.ArrivedAtPump = time.Now()
		fmt.Printf("🕓  %v arrived to a pump and waiting in a queue\n", c)
	case <-time.After(time.Second):
		c.LeftAt = time.Now()
		fmt.Printf("❌  %v left the station because of long wating time at %v\n", c, time.Now().Format("15:04:05.000"))
		wg.Done()
	}
}

func (gs GasStation) start(cars []*car.Car, wg *sync.WaitGroup) {
	for _, c := range cars {
		c.ArrivedAtStation = time.Now()
		fmt.Printf("🚗  %v arrived to a gas station at %v\n", c, time.Now().Format("15:04:05.000"))
		go gs.addCarToQueue(c, wg)
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second) // simulate time between car arrivals
	}
}

func (gs GasStation) Run(cars []*car.Car) {
	wg := sync.WaitGroup{}
	wg.Add(len(cars))

	gs.runPumps()
	go gs.proceedCheckout(&wg)

	go gs.start(cars, &wg)

	wg.Wait()
}
