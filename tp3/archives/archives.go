package archives

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"tdas/grafo"
	"tp3/cities"
)

func ReadParams() cities.Cities {
	params := os.Args
	fAirports, _ := os.Open(params[1])
	fFlights, _ := os.Open(params[2])
	defer fAirports.Close()
	defer fFlights.Close()
	cit := cities.CreateDB()
	scanAirports(bufio.NewScanner(fAirports), &cit)
	scanFlights(bufio.NewScanner(fFlights), &cit)
	return cit
}

func scanAirports(s *bufio.Scanner, c *cities.Cities) {
	for s.Scan() {
		details := strings.Split(s.Text(), ",")
		lat, _ := strconv.ParseFloat(details[2], 64)
		long, _ := strconv.ParseFloat(details[3], 64)
		(*c).AddAirport(details[1], details[0], lat, long)
	}
}

func scanFlights(s *bufio.Scanner, c *cities.Cities) {
	for s.Scan() {
		details := strings.Split(s.Text(), ",")
		cost, _ := strconv.Atoi(details[3])
		time, _ := strconv.Atoi(details[2])
		quant, _ := strconv.Atoi(details[4])
		(*c).Connect(details[0], details[1], cost, time, quant)
	}
}

func GetGraphItinerary(path string) grafo.Grafo[string, int] {
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	g := grafo.CrearGrafo[string, int](true)
	first := true

	for scanner.Scan() {
		if first {
			cit := strings.Split(scanner.Text(), ",")
			for _, city := range cit {
				g.AgregarVertice(city)
			}
			first = false
			continue
		}
		flights := strings.Split(scanner.Text(), ",")
		g.AgregarArista(flights[0], flights[1], 1)
	}
	return g
}
