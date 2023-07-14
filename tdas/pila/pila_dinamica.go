package pila

const (
	TAMINICIAL     = 10
	TAMREDIMENSION = 2
	CUARTO         = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	p := new(pilaDinamica[T])
	p.datos = make([]T, TAMINICIAL)
	return p
}

func (p *pilaDinamica[T]) verificarRedimension() (bool, int) {
	if p.cantidad == len(p.datos) {
		return true, len(p.datos) * TAMREDIMENSION
	}
	if p.cantidad <= len(p.datos)/CUARTO && len(p.datos) > TAMINICIAL {
		return true, len(p.datos) / TAMREDIMENSION
	}
	return false, 0
}

func (p *pilaDinamica[T]) redimensionar(tam int) {
	newP := make([]T, tam)
	copy(newP, p.datos)
	p.datos = newP
}

func (p *pilaDinamica[T]) Apilar(elem T) {
	p.datos[p.cantidad] = elem
	p.cantidad++
	verif, tam := p.verificarRedimension()
	if verif {
		p.redimensionar(tam)
	}
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	res := p.datos[p.cantidad-1]
	p.cantidad--
	verif, tam := p.verificarRedimension()
	if verif {
		p.redimensionar(tam)
	}
	return res
}

func (p pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}
