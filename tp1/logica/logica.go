package logica

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"tdas/cola"
	err "tp1/errores"
	"tp1/votos"
)

func DNIValido(arr []votos.Votante, dni int) (bool, votos.Votante) {
	largo := len(arr)
	if arr[0].LeerDNI() > dni || arr[largo-1].LeerDNI() < dni {
		return false, nil
	}
	if arr[0].LeerDNI() == dni {
		return true, arr[0]
	}
	if arr[largo-1].LeerDNI() == dni {
		return true, arr[largo-1]
	}
	return dniValido(arr, dni, 0, largo)
}

func dniValido(arr []votos.Votante, dni, izq, der int) (bool, votos.Votante) {
	if izq > der {
		return false, nil
	}
	med := (izq + der) / 2
	if arr[med].LeerDNI() == dni {
		return true, arr[med]
	}
	if arr[med].LeerDNI() > dni {
		return dniValido(arr, dni, izq, med-1)
	}
	return dniValido(arr, dni, med+1, der)
}

func hayAlguien(act votos.Votante, fila *cola.Cola[votos.Votante]) bool {
	return act != nil || !(*fila).EstaVacia()
}

func ingresar(fila *cola.Cola[votos.Votante], dniS string, padrones []votos.Votante) {
	dni, _ := strconv.Atoi(dniS)
	if dni <= 0 {
		er := new(err.DNIError)
		fmt.Println(er.Error())
		return
	}
	dniEmpadronado, votante := DNIValido(padrones, dni)
	if !dniEmpadronado {
		er := new(err.DNIFueraPadron)
		fmt.Println(er.Error())
		return
	}
	(*fila).Encolar(votante)
	fmt.Println("OK")
}

func votar(act *votos.Votante, fila *cola.Cola[votos.Votante], tipo, alt string, partidos []votos.Partido) {
	if !hayAlguien(*act, fila) {
		er := new(err.FilaVacia)
		fmt.Println(er.Error())
		return
	} else if *act == nil {
		*act = (*fila).Desencolar()
	}
	if tipo != "Presidente" && tipo != "Gobernador" && tipo != "Intendente" {
		er := new(err.ErrorTipoVoto)
		fmt.Println(er.Error())
		return
	}
	alter, errorConversion := strconv.Atoi(alt)
	if errorConversion != nil || len(partidos)-1 < alter || alter < 0 {
		er := new(err.ErrorAlternativaInvalida)
		fmt.Println(er.Error())
		return
	}
	var er error
	switch tipo {
	case "Presidente":
		er = (*act).Votar(votos.PRESIDENTE, alter)
	case "Gobernador":
		er = (*act).Votar(votos.GOBERNADOR, alter)
	default:
		er = (*act).Votar(votos.INTENDENTE, alter)
	}
	if er != nil {
		fmt.Println(er.Error())
		*act = nil
	} else {
		fmt.Println("OK")
	}
}

func deshacer(act *votos.Votante, fila *cola.Cola[votos.Votante]) {
	if !hayAlguien(*act, fila) {
		er := new(err.FilaVacia)
		fmt.Println(er.Error())
		return
	} else if *act == nil {
		*act = (*fila).Desencolar()
	}
	er := (*act).Deshacer()
	if er != nil {
		if fmt.Sprintf("%T", er) == "*errores.ErrorVotanteFraudulento" {
			*act = nil
		}
		fmt.Println(er.Error())
	} else {
		fmt.Println("OK")
	}
}

func finVotar(act *votos.Votante, fila *cola.Cola[votos.Votante], tiposDeVotos [votos.CANT_VOTACION]votos.TipoVoto, partidos []votos.Partido, impugnados *int) {
	if !hayAlguien(*act, fila) {
		er := new(err.FilaVacia)
		fmt.Println(er.Error())
		return
	} else if *act == nil {
		*act = (*fila).Desencolar()
	}
	votoFinal, er := (*act).FinVoto()
	if er != nil {
		fmt.Println(er.Error())
		*act = nil
		return
	}
	if votoFinal.Impugnado {
		*impugnados++
	} else {
		for _, tipo := range tiposDeVotos {
			partidos[votoFinal.VotoPorTipo[tipo]].VotadoPara(tipo)
		}
	}
	fmt.Println("OK")
	*act = nil
}

func LeerEntrada(fila *cola.Cola[votos.Votante], padrones []votos.Votante, partidos []votos.Partido, impugnados *int, tiposDeVotos [votos.CANT_VOTACION]votos.TipoVoto) {
	input := bufio.NewScanner(os.Stdin)
	var act votos.Votante
	for input.Scan() {
		line := strings.Fields(input.Text())
		cmd := line[0]
		switch cmd {
		case "ingresar":
			ingresar(fila, line[1], padrones)
		case "votar":
			votar(&act, fila, line[1], line[2], partidos)
		case "deshacer":
			deshacer(&act, fila)
		case "fin-votar":
			finVotar(&act, fila, tiposDeVotos, partidos, impugnados)
		default:
			fmt.Println("")
		}
	}
	if hayAlguien(act, fila) {
		er := new(err.ErrorCiudadanosSinVotar)
		fmt.Println(er.Error())
	}
}
