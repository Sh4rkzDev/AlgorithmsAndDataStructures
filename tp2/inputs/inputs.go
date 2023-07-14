package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	air "tp2/airport"
	err "tp2/errors"
)

func agregarArchivo(airport *air.Airport, path string) {
	er := (*airport).AgregarArchivo(path)
	if er != nil {
		os.Stderr.WriteString(er.Error())
	} else {
		fmt.Println("OK")
	}
}

func verTablero(airport *air.Airport, k int, mode, from, to string) {
	res, er := (*airport).VerTablero(k, mode, from, to)
	if er != nil {
		os.Stderr.WriteString(er.Error())
		return
	}
	for _, e := range res {
		fmt.Println(e)
	}
	fmt.Println("OK")
}

func infoVuelo(airport *air.Airport, id string) {
	res, er := (*airport).InfoVuelo(id)
	if er != nil {
		os.Stderr.WriteString(er.Error())
		return
	}
	fmt.Println(res)
	fmt.Println("OK")
}

func prioridadVuelos(airport *air.Airport, k int) {
	if k <= 0 {
		os.Stderr.WriteString("Error en comando prioridad_vuelos\n")
		return
	}
	res := (*airport).PrioridadVuelos(k)
	for _, e := range res {
		fmt.Println(e)
	}
	fmt.Println("OK")
}

func siguienteVuelo(airport *air.Airport, origin, dest, from string) {
	res := (*airport).SiguienteVuelo(origin, dest, from)
	if res != "" {
		fmt.Println(res)
	} else {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origin, dest, from)
	}
	fmt.Println("OK")
}

func borrar(airport *air.Airport, from, to string) {
	res, er := (*airport).Borrar(from, to)
	if er != nil {
		os.Stderr.WriteString(er.Error())
		return
	}
	for _, e := range res {
		fmt.Println(e)
	}
	fmt.Println("OK")
}

func ReadInput(airport *air.Airport) {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := strings.Fields(input.Text())
		cmd := line[0]
		switch cmd {
		case "agregar_archivo":
			agregarArchivo(airport, line[1])
		case "ver_tablero":
			if len(line) != 5 {
				er := err.ErrorTablero{}
				os.Stderr.WriteString(er.Error())
			} else {
				k, _ := strconv.Atoi(line[1])
				verTablero(airport, k, line[2], line[3], line[4])
			}
		case "info_vuelo":
			if len(line) != 2 {
				er := err.ErrorInfo{}
				os.Stderr.WriteString(er.Error())
				continue
			}
			infoVuelo(airport, line[1])
		case "prioridad_vuelos":
			if len(line) != 2 {
				er := err.ErrorPrior{}
				os.Stderr.WriteString(er.Error())
				continue
			}
			k, _ := strconv.Atoi(line[1])
			prioridadVuelos(airport, k)
		case "siguiente_vuelo":
			if len(line) != 4 {
				er := err.ErrorSig{}
				os.Stderr.WriteString(er.Error())
				continue
			}
			siguienteVuelo(airport, line[1], line[2], line[3])
		case "borrar":
			if len(line) != 3 {
				er := err.ErrorBorrar{}
				os.Stderr.WriteString(er.Error())
				continue
			}
			borrar(airport, line[1], line[2])
		default:
			er := err.ErrorCmd{Cmd: cmd}
			os.Stderr.WriteString(er.Error())
		}
	}
}
