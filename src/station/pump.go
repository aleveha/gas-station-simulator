package station

import (
	"fmt"
	"gas-station-simulator/src/car"
	"gas-station-simulator/src/fuel"
	"time"
)

type Pump struct {
	Type  fuel.Fuel
	Queue chan *car.Car
}

func (p Pump) Run(checkouts chan *car.Car) {
	for c := range p.Queue {
		fmt.Printf("%v is refueling, waiting for %v..\n", c, c.RefuelDuration)
		c.Refuel()
		c.RefueledAt = time.Now()
		fmt.Printf("%v successfully refueled, moved to checkout..\n", c)
		checkouts <- c
	}
}

func GeneratePumps() map[fuel.Fuel]Pump {
	pumps := make(map[fuel.Fuel]Pump)
	pumps[fuel.Gas] = Pump{
		Type:  fuel.Gas,
		Queue: make(chan *car.Car, 4),
	}
	pumps[fuel.Diesel] = Pump{
		Type:  fuel.Diesel,
		Queue: make(chan *car.Car, 4),
	}
	pumps[fuel.LPG] = Pump{
		Type:  fuel.LPG,
		Queue: make(chan *car.Car, 1),
	}
	pumps[fuel.Electric] = Pump{
		Type:  fuel.LPG,
		Queue: make(chan *car.Car, 8),
	}
	return pumps
}
