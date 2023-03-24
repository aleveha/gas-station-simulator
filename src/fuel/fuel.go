package fuel

type Fuel string

var Types = []Fuel{Gas, Diesel, LPG, Electric}

const (
	Diesel   Fuel = "Diesel"
	Electric      = "Electric"
	Gas           = "Gas"
	LPG           = "LPG"
)
