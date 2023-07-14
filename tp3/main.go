package main

import (
	"tp3/archives"
	"tp3/inputs"
)

func main() {
	city := archives.ReadParams()
	inputs.ReadInputs(city)
}
