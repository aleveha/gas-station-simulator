package car

import (
	"fmt"
	"time"
)

type Type int

const (
	Gas Type = iota
	Diesel
	LPG
	Electric
)

type Car struct {
	Id             int
	Type           Type
	FillUpDuration time.Duration
}

type Refueler interface {
	FillUp() time.Duration
}

func (car Car) FillUp() time.Duration {
	fmt.Printf("%v car (id: %d) is filling up, please wait for %v..\n", carTypeToString(car.Type), car.Id, car.FillUpDuration)
	defer fmt.Printf("%v car (id: %d) successfully refueled\n", carTypeToString(car.Type), car.Id)
	time.Sleep(car.FillUpDuration)
	return car.FillUpDuration
}

func GetCarType(n int) Type {
	switch n {
	case 3:
		return Electric
	case 2:
		return LPG
	case 1:
		return Diesel
	default:
		return Gas
	}
}

func carTypeToString(t Type) string {
	switch t {
	case Electric:
		return "Electric"
	case LPG:
		return "LPG"
	case Diesel:
		return "Diesel"
	case Gas:
		return "Gas"
	default:
		return "Unknown type"
	}
}
