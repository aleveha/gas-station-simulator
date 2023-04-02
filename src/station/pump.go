package station

import (
	"fmt"
	"gas-station-simulator/src/car"
	"gas-station-simulator/src/fuel"
	"time"
)

type Pump struct {
	Type    fuel.Fuel
	Queue   chan *car.Car
	Nipples []chan *car.Car
}

func (p Pump) refuel(ch chan *car.Car, c *car.Car, checkouts chan *car.Car) {
	fmt.Printf("⛽️  %v started refueling at %v, wait for %v\n", c, time.Now().Format("15:04:05.000"), c.RefuelDuration)
	c.Refuel()
	c.RefueledAt = time.Now()
	fmt.Printf("🤑  %v successfully refueled, moved to checkout..\n", c)
	checkouts <- <-ch
}

func (p Pump) Run(checkouts chan *car.Car) {
	for c := range p.Queue {
		isWaiting := true
		for isWaiting {
			for _, ch := range p.Nipples {
				select {
				case ch <- c:
					isWaiting = false
					go p.refuel(ch, c, checkouts)
				default:
				}
				if !isWaiting {
					break
				}
			}
		}
	}
}

func genNipples(count int) []chan *car.Car {
	nipples := make([]chan *car.Car, count)
	for i := range nipples {
		nipples[i] = make(chan *car.Car, 1)
	}
	return nipples
}

func GeneratePumps() map[fuel.Fuel]Pump {
	pumps := make(map[fuel.Fuel]Pump)
	pumps[fuel.Gas] = Pump{
		Type:    fuel.Gas,
		Queue:   make(chan *car.Car, 3),
		Nipples: genNipples(4),
	}
	pumps[fuel.Diesel] = Pump{
		Type:    fuel.Diesel,
		Queue:   make(chan *car.Car, 3),
		Nipples: genNipples(4),
	}
	pumps[fuel.LPG] = Pump{
		Type:    fuel.LPG,
		Queue:   make(chan *car.Car, 2),
		Nipples: genNipples(1),
	}
	pumps[fuel.Electric] = Pump{
		Type:    fuel.LPG,
		Queue:   make(chan *car.Car, 2),
		Nipples: genNipples(8),
	}
	return pumps
}
