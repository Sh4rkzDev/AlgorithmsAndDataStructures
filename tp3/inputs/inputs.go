package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tp3/cities"
	"tp3/features"
)

func createString(arr []string) string {
	res := ""
	for i, city := range arr {
		res += city
		if i != len(arr)-1 {
			res += " -> "
		}
	}
	res += "\n"
	return res
}

func caminoMas(c cities.Cities, mode, origin, dest string, lastRouteCmd *[]string) {
	by := 0
	if mode == "rapido" {
		by = cities.TIME
	} else {
		by = cities.COST
	}
	order, _, _ := features.CaminoMas[int](c, by, origin, &dest, func(i1, i2 int) int {
		return i1 + i2
	})
	*lastRouteCmd = order
	res := createString(order)
	fmt.Printf("%s", res)
}

func caminoEscalas(c cities.Cities, origin, dest string, lastRouteCmd *[]string) {
	order := features.CaminoMenosEscalas(c, origin, dest)
	*lastRouteCmd = order
	res := createString(order)
	fmt.Printf("%s", res)
}

func centralidad(c cities.Cities, n int) {
	res := features.MostImportant(c, n)
	aux := ""
	for i, e := range res {
		aux += e
		if e == "" {
			break
		}
		if i != len(res)-1 {
			aux += ", "
		}
	}
	fmt.Println(aux)
}

func nuevaAerolinea(c cities.Cities, r string) {
	features.RouteOptimization(r, c)
	fmt.Println("OK")
}

func itinerario(c cities.Cities, r string) {
	order, sOver := features.Itinerary(c, r)
	res := ""
	for i, city := range order {
		res += city
		if i != len(order)-1 {
			res += ", "
		}
	}
	fmt.Printf("%s\n", res)
	for _, e := range sOver {
		fmt.Printf("%s", createString(e))
	}
}

func exportarKML(c cities.Cities, lastRouteCmd []string, r string) {
	features.ExportKML(c, lastRouteCmd, r)
	fmt.Println("OK")
}

func getCMD(line string) (string, []string) {
	var res string
	var resArr []string
	if len(line) < 3 {
		return res, resArr
	}
	cmd := line[:3]
	if cmd == "cam" {
		if line[7:8] == "m" {
			res = "camino_mas"
			aux := line[11:]
			resArr = strings.Split(aux, ",")
		} else {
			res = "camino_escalas"
			aux := line[15:]
			resArr = strings.Split(aux, ",")
		}
	} else if cmd == "cen" {
		res = "centralidad"
		resArr = []string{line[12:]}
	} else if cmd == "nue" {
		res = "nueva_aerolinea"
		resArr = []string{line[16:]}
	} else if cmd == "iti" {
		res = "itinerario"
		resArr = []string{line[11:]}
	} else if cmd == "exp" {
		res = "exportar_kml"
		resArr = []string{line[13:]}
	}
	return res, resArr
}

func ReadInputs(c cities.Cities) {
	input := bufio.NewScanner(os.Stdin)
	lastRouteCmd := []string{}
	for input.Scan() {
		cmd, params := getCMD(input.Text())
		switch cmd {
		case "camino_mas":
			caminoMas(c, params[0], params[1], params[2], &lastRouteCmd)
		case "camino_escalas":
			caminoEscalas(c, params[0], params[1], &lastRouteCmd)
		case "centralidad":
			n, _ := strconv.Atoi(params[0])
			centralidad(c, n)
		case "nueva_aerolinea":
			nuevaAerolinea(c, params[0])
		case "itinerario":
			itinerario(c, params[0])
		case "exportar_kml":
			exportarKML(c, lastRouteCmd, params[0])
		default:
			fmt.Println("Que pone' bobo? Ponelo bien bobo")
		}
	}
}
