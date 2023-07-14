package main

import (
	"fmt"

	"tdas/cola"
	"tp1/archivos"
	"tp1/logica"
	"tp1/votos"
)

func imprimirTipo(tipo votos.TipoVoto) {
	switch tipo {
	case votos.PRESIDENTE:
		fmt.Println("Presidente:")
	case votos.GOBERNADOR:
		fmt.Println("Gobernador:")
	default:
		fmt.Println("Intendente:")
	}
}

func main() {
	padrones, partidos, errorAlLeer := archivos.ObtenerData()
	if errorAlLeer != nil {
		fmt.Println(errorAlLeer.Error())
		return
	}
	fila := cola.CrearColaEnlazada[votos.Votante]()
	impugnados := 0
	tiposDeVotos := [votos.CANT_VOTACION]votos.TipoVoto{votos.PRESIDENTE, votos.GOBERNADOR, votos.INTENDENTE}
	logica.LeerEntrada(&fila, padrones, partidos, &impugnados, tiposDeVotos)
	for _, tipo := range tiposDeVotos {
		imprimirTipo(tipo)
		for _, partido := range partidos {
			fmt.Println(partido.ObtenerResultado(tipo))
		}
		fmt.Println()
	}
	if impugnados != 1 {
		fmt.Println("Votos Impugnados:", impugnados, "votos")
	} else {
		fmt.Println("Votos Impugnados:", impugnados, "voto")
	}
}
