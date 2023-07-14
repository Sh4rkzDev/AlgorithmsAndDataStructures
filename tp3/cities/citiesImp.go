package cities

import (
	dic "tdas/diccionario"
	"tdas/grafo"
)

const (
	COST = iota
	TIME
	FREQ
)

type cities struct {
	hash   dic.Diccionario[string, dic.Diccionario[string, [2]float64]]
	codes  dic.Diccionario[string, string]
	cities grafo.Grafo[string, []int]
}

func CreateDB() Cities {
	cit := dic.CrearHash[string, dic.Diccionario[string, [2]float64]]()
	codes := dic.CrearHash[string, string]()
	citiesG := grafo.CrearGrafo[string, []int](false)
	return &cities{cit, codes, citiesG}
}

func (c *cities) AddAirport(code string, city string, lat, long float64) {
	c.codes.Guardar(code, city)
	if !c.hash.Pertenece(city) {
		airs := dic.CrearHash[string, [2]float64]()
		c.hash.Guardar(city, airs)
	}
	coords := [2]float64{lat, long}
	c.hash.Obtener(city).Guardar(code, coords)
	c.cities.AgregarVertice(code)
}

func (c *cities) BelongsTo(code string, city string) bool {
	return c.codes.Obtener(code) == city
}

func (c *cities) Connect(c1, c2 string, cost, time, freq int) {
	c.cities.AgregarArista(c1, c2, []int{cost, time, freq})
}

func (c *cities) GetNumberFlights(c1, c2 string) int {
	return c.cities.Peso(c1, c2)[FREQ]
}

func (c *cities) GetGraph() grafo.Grafo[string, []int] {
	return c.cities
}

func (c *cities) GetAirports(city string) []string {
	res := make([]string, c.hash.Obtener(city).Cantidad())
	i := 0
	c.hash.Obtener(city).Iterar(func(clave string, dato [2]float64) bool {
		res[i] = clave
		i++
		return true
	})
	return res
}

func (c *cities) GetCoords(code string) (float64, float64) {
	city := c.codes.Obtener(code)
	latLong := c.hash.Obtener(city).Obtener(code)
	return latLong[0], latLong[1]
}

func (c *cities) GetCity(code string) string {
	return c.codes.Obtener(code)
}
