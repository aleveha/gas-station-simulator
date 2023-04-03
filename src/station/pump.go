package station

import (
	"fmt"
	"gas-station-simulator/src/car"
	"gas-station-simulator/src/fuel"
	"sync"
	"time"
)

type Pump struct {
	Type    fuel.Fuel
	Queue   chan *car.Car
	Nipples int
}

func (p Pump) refuel(c *car.Car, checkouts chan *car.Car) {
	fmt.Printf("⛽️  %v started refueling at %v, wait for %v\n", c, time.Now().Format("15:04:05.000"), c.RefuelDuration)
	c.Refuel()
	c.RefueledAt = time.Now()
	fmt.Printf("🤑  %v successfully refueled, moved to checkout..\n", c)
	checkouts <- c
}

func (p Pump) Run(checkouts chan *car.Car) {
	wg := &sync.WaitGroup{}
	wg.Add(p.Nipples)
	for i := 0; i < p.Nipples; i++ {
		go func() {
			defer wg.Done()
			for c := range p.Queue {
				p.refuel(c, checkouts)
			}
		}()
	}
	wg.Wait()
}

func GeneratePumps() map[fuel.Fuel]Pump {
	pumps := make(map[fuel.Fuel]Pump)
	pumps[fuel.Gas] = Pump{
		Type:    fuel.Gas,
		Queue:   make(chan *car.Car, 3),
		Nipples: 4,
	}
	pumps[fuel.Diesel] = Pump{
		Type:    fuel.Diesel,
		Queue:   make(chan *car.Car, 3),
		Nipples: 4,
	}
	pumps[fuel.LPG] = Pump{
		Type:    fuel.LPG,
		Queue:   make(chan *car.Car, 2),
		Nipples: 4,
	}
	pumps[fuel.Electric] = Pump{
		Type:    fuel.LPG,
		Queue:   make(chan *car.Car, 2),
		Nipples: 4,
	}
	return pumps
}
