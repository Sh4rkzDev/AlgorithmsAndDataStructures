package votos

import (
	pila "tdas/pila"
	err "tp1/errores"
)

type votanteImplementacion struct {
	Voto
	realizado bool
	dni       int
	pMov      pila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	v := new(votanteImplementacion)
	v.VotoPorTipo = [CANT_VOTACION]int{}
	v.dni = dni
	v.pMov = pila.CrearPilaDinamica[Voto]()
	return v
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) verificarFraude() error {
	if votante.realizado {
		er := new(err.ErrorVotanteFraudulento)
		er.Dni = votante.LeerDNI()
		return er
	}
	return nil
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	errorFraude := votante.verificarFraude()
	if errorFraude != nil {
		return errorFraude
	}
	if alternativa == LISTA_IMPUGNA || votante.Impugnado {
		votante.Impugnado = true
	}
	votante.VotoPorTipo[tipo] = alternativa
	cop := votante.Voto
	votante.pMov.Apilar(cop)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	errorFraude := votante.verificarFraude()
	if errorFraude != nil {
		return errorFraude
	}
	if votante.pMov.EstaVacia() {
		er := new(err.ErrorNoHayVotosAnteriores)
		return er
	}
	votante.pMov.Desapilar()
	if votante.pMov.EstaVacia() {
		votante.VotoPorTipo[PRESIDENTE] = 0
		votante.VotoPorTipo[GOBERNADOR] = 0
		votante.VotoPorTipo[INTENDENTE] = 0
		votante.Impugnado = false
	} else {
		votante.Voto = votante.pMov.VerTope()
	}
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	errorFraude := votante.verificarFraude()
	if errorFraude != nil {
		return votante.Voto, errorFraude
	}
	votante.realizado = true
	return votante.Voto, nil
}
