package votos

import "fmt"

type partidoImplementacion struct {
	nombre          string
	candidatos      [CANT_VOTACION]string
	votosCandidatos [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votosCandidatos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	p := new(partidoImplementacion)
	p.nombre = nombre
	for i, candidato := range candidatos {
		p.candidatos[i] = candidato
	}
	return p
}

func CrearVotosEnBlanco() Partido {
	return new(partidoEnBlanco)
}

func votos(k int) string {
	if k == 1 {
		return "voto"
	}
	return "votos"
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votosCandidatos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	return fmt.Sprintf("%s - %s: %d %s", partido.nombre, partido.candidatos[tipo], partido.votosCandidatos[tipo], votos(partido.votosCandidatos[tipo]))
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votosCandidatos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	return fmt.Sprintf("Votos en Blanco: %d %s", blanco.votosCandidatos[tipo], votos(blanco.votosCandidatos[tipo]))
}
