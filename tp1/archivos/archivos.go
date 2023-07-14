package archivos

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	err "tp1/errores"
	"tp1/votos"
)

func scanPartidos(s *bufio.Scanner) []votos.Partido {
	res := make([]votos.Partido, 1)
	res[0] = votos.CrearVotosEnBlanco()
	for s.Scan() {
		lista := strings.Split(s.Text(), ",")
		candidatos := [votos.CANT_VOTACION]string{lista[1], lista[2], lista[3]}
		partido := votos.CrearPartido(lista[0], candidatos)
		res = append(res, partido)
	}
	return res
}

func scanPadrones(s *bufio.Scanner) []votos.Votante {
	res := make([]votos.Votante, 0)
	for s.Scan() {
		dni, _ := strconv.Atoi(s.Text())
		votante := votos.CrearVotante(dni)
		res = append(res, votante)
	}
	return res
}

func ordenarPadrones(arr []votos.Votante) []votos.Votante {
	res := make([]votos.Votante, len(arr))
	copy(res, arr)
	k := 1
	for k <= 10000000 {
		res = ordenar(res, k)
		k *= 10
	}
	return res
}

func ordenar(arr []votos.Votante, k int) []votos.Votante {
	res := make([]votos.Votante, len(arr))
	freq := [10]int{}
	pos := [10]int{}
	for _, elem := range arr {
		dni := elem.LeerDNI()
		freq[(dni/k)%10]++
	}
	for i := 1; i < 10; i++ {
		pos[i] = freq[i-1] + pos[i-1]
	}
	for _, elem := range arr {
		dni := elem.LeerDNI()
		digito := (dni / k) % 10
		res[pos[digito]] = elem
		pos[digito]++
	}
	return res
}

func ObtenerData() ([]votos.Votante, []votos.Partido, error) {
	params := os.Args
	padrones := make([]votos.Votante, 0)
	partidos := make([]votos.Partido, 0)
	if len(params) != 3 {
		er := new(err.ErrorParametros)
		return padrones, partidos, er
	}
	arcPartidos, errPart := os.Open(params[1])
	arcPadrones, errPad := os.Open(params[2])
	if errPart != nil || errPad != nil {
		er := new(err.ErrorLeerArchivo)
		return padrones, partidos, er
	}
	defer arcPartidos.Close()
	defer arcPadrones.Close()
	partidos = scanPartidos(bufio.NewScanner(arcPartidos))
	padrones = scanPadrones(bufio.NewScanner(arcPadrones))
	padronesOrdenados := ordenarPadrones(padrones)
	return padronesOrdenados, partidos, nil
}
