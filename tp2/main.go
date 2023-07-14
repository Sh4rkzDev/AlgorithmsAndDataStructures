package main

import (
	air "tp2/airport"
	"tp2/inputs"
)

func main() {
	airport := air.CreateAirport()
	inputs.ReadInput(&airport)
}
