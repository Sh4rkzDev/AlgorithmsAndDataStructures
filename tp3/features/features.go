package features

import (
	"bufio"
	"fmt"
	"os"

	queue "tdas/cola"
	heap "tdas/cola_prioridad"
	dic "tdas/diccionario"
	"tdas/grafo"
	"tp3/archives"
	"tp3/cities"
)

type cityDist[T int | float32] struct {
	city string
	dist T
}

func cmpCityDistMin[T int | float32](a, b cityDist[T]) int {
	if a.dist > b.dist {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}

func cmpCityDistMax[T int | float32](a, b cityDist[T]) int {
	if a.dist > b.dist {
		return 1
	}
	if a == b {
		return 0
	}
	return -1
}

func CaminoMas[T int | float32](c cities.Cities, by int, origin string, dest *string, getWeight func(T, T) T) ([]string, dic.Diccionario[string, string], dic.Diccionario[string, T]) {
	var prev dic.Diccionario[string, string]
	var dist dic.Diccionario[string, T]
	g := c.GetGraph()

	firstFound := false
	airDest := ""
	var maxDist T
	var minVal T

	airportsOrigin := c.GetAirports(origin)

	for _, airOrigin := range airportsOrigin {
		actPrev := dic.CrearHash[string, string]()
		actDist := dic.CrearHash[string, T]()
		actVis := dic.CrearHash[string, int]()

		actPrev.Guardar(airOrigin, "")
		actDist.Guardar(airOrigin, minVal)

		o := cityDist[T]{airOrigin, minVal}
		var h heap.ColaPrioridad[cityDist[T]]
		h = heap.CrearHeap[cityDist[T]](cmpCityDistMin[T])
		h.Encolar(o)

		for !h.EstaVacia() {
			v := h.Desencolar()
			if firstFound && maxDist <= v.dist {
				break
			}
			if dest != nil && c.BelongsTo(v.city, *dest) {
				firstFound = true
				maxDist = v.dist
				airDest = v.city
				prev = actPrev
				dist = actDist
				break
			}
			if actVis.Pertenece(v.city) {
				continue
			}
			actVis.Guardar(v.city, 0)
			for _, w := range g.Adyacentes(v.city) {
				if actVis.Pertenece(w) {
					continue
				}
				distHere := getWeight(actDist.Obtener(v.city), T(g.Peso(v.city, w)[by]))
				if !actDist.Pertenece(w) || distHere < actDist.Obtener(w) {
					actDist.Guardar(w, distHere)
					actPrev.Guardar(w, v.city)
					h.Encolar(cityDist[T]{w, distHere})
				}
			}
			if h.EstaVacia() || actVis.Cantidad() == g.Cantidad() {
				prev = actPrev
				dist = actDist
				break
			}
		}
	}
	return getPath(airDest, prev), prev, dist
}

func getPath(dest string, prev dic.Diccionario[string, string]) []string {
	res := []string{}
	act := dest
	for act != "" {
		res = append(res, act)
		act = prev.Obtener(act)
	}
	invert(res)
	return res
}

func invert(arr []string) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}

func CaminoMenosEscalas(c cities.Cities, origin, dest string) []string {
	prev := dic.CrearHash[string, string]()
	g := c.GetGraph()

	airportsOrigin := c.GetAirports(origin)

	firstFound := false
	maxDist := 0
	airDest := ""

	for _, airOrigin := range airportsOrigin {
		q := queue.CrearColaEnlazada[string]()
		vis := dic.CrearHash[string, int]()
		actPrev := dic.CrearHash[string, string]()
		actDist := dic.CrearHash[string, int]()

		actPrev.Guardar(airOrigin, "")
		actDist.Guardar(airOrigin, 0)
		vis.Guardar(airOrigin, 0)
		q.Encolar(airOrigin)
		stop := false

		for !q.EstaVacia() {
			v := q.Desencolar()
			if firstFound && actDist.Obtener(v) == maxDist {
				break
			}
			for _, w := range g.Adyacentes(v) {
				if vis.Pertenece(w) {
					continue
				}
				vis.Guardar(w, 0)
				actPrev.Guardar(w, v)
				actDist.Guardar(w, actDist.Obtener(v)+1)
				if c.BelongsTo(w, dest) {
					firstFound = true
					maxDist = actDist.Obtener(w)
					prev = actPrev
					airDest = w
					stop = true
					break
				}
				q.Encolar(w)
			}
			if stop {
				break
			}
		}
	}
	return getPath(airDest, prev)
}

func MostImportant(c cities.Cities, n int) []string {
	cent := dic.CrearHash[string, int]()
	g := c.GetGraph()
	g.Iterar(func(v string) bool {
		cent.Guardar(v, 0)
		return true
	})
	g.Iterar(func(v string) bool {
		_, prev, dist := CaminoMas[float32](c, cities.FREQ, c.GetCity(v), nil, func(i1, i2 float32) float32 {
			return i1 + 1/i2
		})
		centAux := dic.CrearHash[string, int]()
		g.Iterar(func(w string) bool {
			centAux.Guardar(w, 0)
			return true
		})
		sortedVert := sortVert(dist)
		for _, w := range sortedVert {
			if prev.Obtener(w) == "" {
				continue
			}
			centAux.Guardar(prev.Obtener(w), 1+centAux.Obtener(w)+centAux.Obtener(prev.Obtener(w)))
		}
		g.Iterar(func(w string) bool {
			if w == v {
				return true
			}
			cent.Guardar(w, cent.Obtener(w)+centAux.Obtener(w))
			return true
		})
		return true
	})
	return getImportant(cent, n)
}

func sortVert(dist dic.Diccionario[string, float32]) []string {
	res := make([]string, dist.Cantidad())
	auxRes := make([]cityDist[float32], dist.Cantidad())
	i := 0
	dist.Iterar(func(clave string, dato float32) bool {
		aux := cityDist[float32]{clave, dato}
		auxRes[i] = aux
		i++
		return true
	})
	heap.HeapSort[cityDist[float32]](auxRes, cmpCityDistMin[float32])
	for ind, e := range auxRes {
		res[ind] = e.city
	}
	return res
}

func getImportant(d dic.Diccionario[string, int], n int) []string {
	aux := make([]cityDist[int], d.Cantidad())
	i := 0
	d.Iterar(func(clave string, dato int) bool {
		aux[i] = cityDist[int]{clave, dato}
		i++
		return true
	})
	res := make([]string, n)
	auxRes := heap.CrearHeapArr[cityDist[int]](aux, cmpCityDistMax[int])
	for i := 0; i < n; i++ {
		res[i] = auxRes.Desencolar().city
	}
	return res
}

type distOrigDest struct {
	origin, dest string
	dist         int
}

func cmpDistOrigDest(a, b distOrigDest) int {
	return b.dist - a.dist
}

func primRoutes(c cities.Cities) dic.Diccionario[string, string] {
	g := c.GetGraph()
	h := heap.CrearHeap[distOrigDest](cmpDistOrigDest)
	prev := dic.CrearHash[string, string]()
	vis := dic.CrearHash[string, int]()

	vIni := g.ObtenerVertice()
	vis.Guardar(vIni, 0)

	for _, w := range g.Adyacentes(vIni) {
		h.Encolar(distOrigDest{vIni, w, g.Peso(vIni, w)[cities.COST]})
	}

	for !h.EstaVacia() {
		aux := h.Desencolar()
		if vis.Pertenece(aux.dest) {
			continue
		}
		vis.Guardar(aux.dest, 0)
		prev.Guardar(aux.dest, aux.origin)
		if vis.Cantidad() == g.Cantidad() {
			break
		}
		for _, w := range g.Adyacentes(aux.dest) {
			if !vis.Pertenece(w) {
				h.Encolar(distOrigDest{aux.dest, w, g.Peso(aux.dest, w)[cities.COST]})
			}
		}
	}
	return prev
}

func RouteOptimization(r string, c cities.Cities) {
	f, _ := os.Create(r)
	defer f.Close()
	w := bufio.NewWriter(f)
	g := c.GetGraph()

	prev := primRoutes(c)
	prev.Iterar(func(clave, dato string) bool {
		w.WriteString(fmt.Sprintf("%s,%s,%d,%d,%d\n", dato, clave, g.Peso(clave, dato)[cities.TIME], g.Peso(clave, dato)[cities.COST], c.GetNumberFlights(clave, dato)))
		return true
	})
	w.Flush()
}

func Itinerary(c cities.Cities, path string) ([]string, [][]string) {
	g := archives.GetGraphItinerary(path)
	incEd := getGrades(g)
	q := queue.CrearColaEnlazada[string]()
	g.Iterar(func(v string) bool {
		if !incEd.Pertenece(v) {
			q.Encolar(v)
		}
		return true
	})

	order := make([]string, g.Cantidad())
	i := 0

	for !q.EstaVacia() {
		v := q.Desencolar()
		order[i] = v
		for _, w := range g.Adyacentes(v) {
			incEd.Guardar(w, incEd.Obtener(w)-1)
			if incEd.Obtener(w) == 0 {
				q.Encolar(w)
			}
		}
		i++
	}
	res := make([][]string, len(order)-1)
	for i, city := range order {
		auxRes := CaminoMenosEscalas(c, city, order[i+1])
		res[i] = auxRes

		if i == len(order)-2 {
			break
		}
	}
	return order, res
}

func getGrades(g grafo.Grafo[string, int]) dic.Diccionario[string, int] {
	res := dic.CrearHash[string, int]()
	g.Iterar(func(v string) bool {
		for _, w := range g.Adyacentes(v) {
			if !res.Pertenece(w) {
				res.Guardar(w, 0)
			}
			res.Guardar(w, res.Obtener(w)+1)
		}
		return true
	})
	return res
}

func ExportKML(c cities.Cities, arr []string, path string) {
	f, _ := os.Create(path)
	defer f.Close()
	w := bufio.NewWriter(f)

	w.WriteString("" +
		"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
		"<kml xmlns=\"http://earth.google.com/kml/2.1\">\n" +
		"	<Document>\n" +
		"		<name>KML File created by Messi</name>\n" +
		"		<description>MUCHAAAAAAAAACHOOOOOOOOOOOOOOOS</description>\n\n")

	lines := ""

	for i, code := range arr {
		lat, long := c.GetCoords(code)
		w.WriteString(createPlaceMark(code, "", lat, long))
		if i == 0 {
			continue
		}
		city2 := arr[i-1]
		lat2, long2 := c.GetCoords(city2)
		lines += createPlaceMarkLine(lat2, long2, lat, long)
	}

	w.WriteString(lines)

	w.WriteString("" +
		"	</Document>\n" +
		"</kml>\n")
	w.Flush()
}

func createPlaceMark(name, desc string, lat, long float64) string {
	return fmt.Sprintf(""+
		"		<Placemark>\n"+
		"			<name>%s</name>\n"+
		"			<description>%s</description>\n"+
		"			<Point>\n"+
		"				<coordinates>%f, %f</coordinates>\n"+
		"			</Point>\n"+
		"		</Placemark>\n\n",
		name, desc, long, lat)
}

func createPlaceMarkLine(lat1, long1, lat2, long2 float64) string {
	return fmt.Sprintf(""+
		"		<Placemark>\n"+
		"			<LineString>\n"+
		"				<coordinates>%f, %f %f, %f</coordinates>\n"+
		"			</LineString>\n"+
		"		</Placemark>\n\n",
		long1, lat1, long2, lat2)
}
