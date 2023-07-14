package cities

import "tdas/grafo"

type Cities interface {

	//Add the airport to the given city. If the city doesnt exist, it will create it. If it does, just add the airport to it
	AddAirport(code, city string, lat, long float64)

	//Return true if the Airport Code belongs to the given city
	BelongsTo(code, city string) bool

	//Connect both cities with the given cost, time, and quantity of flights that connect these two airports
	Connect(city1, city2 string, cost, time, quant int)

	//Return the number of flights shared by both airports
	GetNumberFlights(city1, city2 string) int

	//Returns the graph that has the cities as vertices and the slice [cost, time, freq] as edges
	GetGraph() grafo.Grafo[string, []int]

	//Returns a slice of all the airports codes of the city
	GetAirports(city string) []string

	//Returns the values of the lat and long of the airport
	GetCoords(code string) (float64, float64)

	//Returns the city where the airport is located
	GetCity(code string) string
}
