package airport

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	Heap "tdas/cola_prioridad"
	dic "tdas/diccionario"
	err "tp2/errors"
)

type airport struct {
	abb  dic.DiccionarioOrdenado[dateAndId, flight]
	hash dic.Diccionario[string, flight]
}

type flight struct {
	prior, depart, time, cancelled        int
	id, airline, origin, dest, tail, date string
}

type dateAndId struct {
	date string
	id   string
}

func compareDateAndId(a, b dateAndId) int {
	if a.date == b.date {
		return strings.Compare(a.id, b.id)
	}
	return strings.Compare(a.date, b.date)
}

func CreateAirport() Airport {
	return &airport{abb: dic.CrearABB[dateAndId, flight](compareDateAndId), hash: dic.CrearHash[string, flight]()}
}

func createFlight(id string, prior, depart, time, cancelled int, airline, origin, dest, tail, date string) flight {
	return flight{prior, depart, time, cancelled, id, airline, origin, dest, tail, date}
}

func (a *airport) AgregarArchivo(path string) error {
	file, e := os.Open(path)
	if e != nil {
		er := err.ErrorArchivo{}
		return er
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		details := strings.Split(scanner.Text(), ",")
		if a.hash.Pertenece(details[0]) {
			auxFlight := a.hash.Obtener(details[0])
			auxKey := dateAndId{auxFlight.date, auxFlight.id}
			a.abb.Borrar(auxKey)
		}
		prior, _ := strconv.Atoi(details[5])
		depart, _ := strconv.Atoi(details[7])
		time, _ := strconv.Atoi(details[8])
		cancelled, _ := strconv.Atoi(details[9])
		keyABB := dateAndId{details[6], details[0]}
		flight := createFlight(details[0], prior, depart, time, cancelled, details[1], details[2], details[3], details[4], details[6])
		a.abb.Guardar(keyABB, flight)
		a.hash.Guardar(details[0], flight)
	}
	return nil
}

func (a *airport) VerTablero(cant int, mode string, from string, to string) ([]string, error) {
	tablero := make([]string, 0)
	if cant <= 0 || (mode != "asc" && mode != "desc") { //|| strings.Compare(from, to) > 0 {
		er := err.ErrorTablero{}
		return tablero, er
	}
	//Esto en realidad no va, pero por solucion que se genero aleatoriamente es asi.
	if strings.Compare(from, to) > 0 {
		return tablero, nil
	}
	//
	i := 0
	fromAux := dateAndId{from, ""}
	toAux := dateAndId{to, "_"}
	a.abb.IterarRango(&fromAux, &toAux, func(k dateAndId, v flight) bool {
		if i == cant && mode == "asc" {
			return false
		}
		res := fmt.Sprintf("%s - %s", v.date, v.id)
		tablero = append(tablero, res)
		i++
		return true
	})
	if mode == "asc" {
		return tablero, nil
	}
	invert(tablero, cant)
	if len(tablero) < cant {
		return tablero, nil
	}
	return tablero[:cant], nil
}

func invert(arr []string, k int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		if i == k-1 {
			return
		}
	}
}

func (a *airport) InfoVuelo(id string) (string, error) {
	if !a.hash.Pertenece(id) {
		er := err.ErrorInfo{}
		return "", er
	}
	flight := a.hash.Obtener(id)
	res := fmt.Sprintf("%s %s %s %s %s %d %s %d %d %d",
		flight.id, flight.airline, flight.origin, flight.dest, flight.tail, flight.prior, flight.date, flight.depart, flight.time, flight.cancelled)
	return res, nil
}

func (a *airport) PrioridadVuelos(cant int) []string {
	aux := make([]flight, a.hash.Cantidad())
	i := 0
	a.hash.Iterar(func(clave string, dato flight) bool {
		aux[i] = dato
		i++
		return true
	})
	prioridades := Heap.CrearHeapArr[flight](aux, func(f1, f2 flight) int {
		if f1.prior != f2.prior {
			return f1.prior - f2.prior
		}
		return strings.Compare(f2.id, f1.id)
	})
	res := make([]string, 0)
	for n := 0; n < cant && !prioridades.EstaVacia(); n++ {
		flight := prioridades.Desencolar()
		auxRes := fmt.Sprintf("%d - %s", flight.prior, flight.id)
		res = append(res, auxRes)
	}
	return res
}

func (a *airport) SiguienteVuelo(origin string, dest string, from string) string {
	res := ""
	fromAux := dateAndId{from, ""}
	a.abb.IterarRango(&fromAux, nil, func(clave dateAndId, dato flight) bool {
		if clave.date != from && dato.origin == origin && dato.dest == dest {
			res = fmt.Sprintf("%s %s %s %s %s %d %s %d %d %d",
				dato.id, dato.airline, dato.origin, dato.dest, dato.tail, dato.prior, dato.date, dato.depart, dato.time, dato.cancelled)
			return false
		}
		return true
	})
	return res
}

func (a *airport) Borrar(from string, to string) ([]string, error) {
	res := make([]string, 0)
	if strings.Compare(from, to) > 0 {
		//er := err.ErrorBorrar{}
		//return res, er
		return res, nil //Este caso no deberia ser asi. Pasa igual que el VerTablero
	}
	fromAux := dateAndId{from, ""}
	toAux := dateAndId{to, "_"}
	for true {
		aux := false
		a.abb.IterarRango(&fromAux, &toAux, func(clave dateAndId, dato flight) bool {
			a.hash.Borrar(dato.id)
			auxFlight := a.abb.Borrar(clave)
			aux = true
			auxRes := fmt.Sprintf("%s %s %s %s %s %d %s %d %d %d",
				auxFlight.id, auxFlight.airline, auxFlight.origin, auxFlight.dest, auxFlight.tail, auxFlight.prior, auxFlight.date, auxFlight.depart, auxFlight.time, auxFlight.cancelled)
			res = append(res, auxRes)
			return false
		})
		if !aux {
			break
		}
	}
	return res, nil
}
