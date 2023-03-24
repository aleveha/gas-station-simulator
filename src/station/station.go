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
		time.Sleep(1 * time.Second)
		chC.LeftAt = time.Now()
		fmt.Printf("%v left the station after checkout at %v\n", chC, time.Now().Format("15:04:05.000"))
		wg.Done()
	}
}

func (gs GasStation) addCarToQueue(c *car.Car, wg *sync.WaitGroup) {
	select {
	case gs.Pumps[c.FuelType].Queue <- c:
		c.ArrivedAtPump = time.Now()
		fmt.Printf("%v arrived to a %v pump at %v\n", c, c.FuelType, time.Now().Format("15:04:05.000"))
	case <-time.After(1 * time.Second):
		c.LeftAt = time.Now()
		fmt.Printf("%v left the station because of long queue at %v\n", c, time.Now().Format("15:04:05.000"))
		wg.Done()
	}
}

func (gs GasStation) start(cars []car.Car, wg *sync.WaitGroup) {
	for i := range cars {
		cars[i].ArrivedAtStation = time.Now()
		go gs.addCarToQueue(&cars[i], wg)
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second) // simulate time between car arrivals
	}

	for _, pump := range gs.Pumps {
		close(pump.Queue)
	}
}

func (gs GasStation) Run(cars []car.Car) {
	wg := sync.WaitGroup{}
	wg.Add(len(cars))

	gs.runPumps()
	go gs.proceedCheckout(&wg)

	go gs.start(cars, &wg)

	wg.Wait()
}
