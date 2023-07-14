package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

// https://golangprojectstructure.com/hash-functions-go-code/#how-are-hash-functions-used-in-the-map-data-structure
const (
	uint64Offset uint64 = 0xcbf29ce484222325
	uint64Prime  uint64 = 0x00000100000001b3
	tamInicial   int    = 11
	factorPred   int    = 10 / 9
)

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashing(data []byte) (hash uint64) {
	hash = uint64Offset

	for _, b := range data {
		hash ^= uint64(b)
		hash *= uint64Prime
	}

	return
}

type parClaveValor[K comparable, V any] struct {
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct {
	tabla []TDALista.Lista[*parClaveValor[K, V]]
	cant  int
}

type iteradorDiccionario[K comparable, V any] struct {
	dic      *hashAbierto[K, V]
	act      int
	actLista TDALista.Lista[*parClaveValor[K, V]]
	actPar   TDALista.IteradorLista[*parClaveValor[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	aux := make([]TDALista.Lista[*parClaveValor[K, V]], tamInicial)
	for i := range aux {
		aux[i] = TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
	}
	return &hashAbierto[K, V]{tabla: aux}
}

func (iter *iteradorDiccionario[K, V]) buscarProx(pos int) {
	for pos < len(iter.dic.tabla) {
		if pos != len(iter.dic.tabla)-1 && iter.dic.tabla[pos].Largo() == 0 {
			pos++
			continue
		}
		iter.act = pos
		iter.actLista = iter.dic.tabla[pos]
		iter.actPar = iter.actLista.Iterador()
		break
	}
}

func (h *hashAbierto[K, V]) crearIterDiccionario() IterDiccionario[K, V] {
	iter := &iteradorDiccionario[K, V]{dic: h}
	iter.buscarProx(0)
	return iter
}

func crearParClaveValor[K comparable, V any](clave K, valor V) *parClaveValor[K, V] {
	return &parClaveValor[K, V]{clave, valor}
}

func obtenerPos[K comparable](clave K, tam int) uint64 {
	return hashing(convertirABytes(clave)) % uint64(tam)
}

func (h *hashAbierto[K, V]) redimensionar(tam int) {
	nueva := make([]TDALista.Lista[*parClaveValor[K, V]], tam)
	for i := range nueva {
		nueva[i] = TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
	}
	for _, lista := range h.tabla {
		for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			par := iter.VerActual()
			pos := obtenerPos(par.clave, tam)
			nueva[pos].InsertarUltimo(crearParClaveValor(par.clave, par.valor))
		}
	}
	h.tabla = nueva
}

func (h *hashAbierto[K, V]) buscar(clave K, valor *V, iter TDALista.IteradorLista[*parClaveValor[K, V]], aux func(*parClaveValor[K, V]) *V) (bool, *V) {
	for ; iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			res := aux(iter.VerActual())
			return true, res
		}
	}
	return false, nil
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	pos := obtenerPos(clave, len(h.tabla))
	iter := h.tabla[pos].Iterador()
	found, _ := h.buscar(clave, &dato, iter, func(par *parClaveValor[K, V]) *V {
		par.valor = dato
		return nil
	})
	if found {
		return
	}
	h.tabla[pos].InsertarUltimo(crearParClaveValor(clave, dato))
	h.cant++
	if h.Cantidad()/len(h.tabla) >= 2 {
		h.redimensionar(h.Cantidad() * factorPred)
	}
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	pos := obtenerPos(clave, len(h.tabla))
	iter := h.tabla[pos].Iterador()
	found, _ := h.buscar(clave, nil, iter, func(_ *parClaveValor[K, V]) *V {
		return nil
	})
	return found
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	pos := obtenerPos(clave, len(h.tabla))
	iter := h.tabla[pos].Iterador()
	found, res := h.buscar(clave, nil, iter, func(par *parClaveValor[K, V]) *V {
		return &par.valor
	})
	if !found {
		panic("La clave no pertenece al diccionario")
	}
	return *res
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	pos := obtenerPos(clave, len(h.tabla))
	iter := h.tabla[pos].Iterador()
	found, res := h.buscar(clave, nil, iter, func(par *parClaveValor[K, V]) *V {
		res := iter.Borrar()
		h.cant--
		factor := float32(h.Cantidad()) / float32(len(h.tabla))
		aux := h.Cantidad() * factorPred
		if factor <= 0.2 && aux > tamInicial {
			h.redimensionar(aux)
		}
		return &res.valor
	})
	if !found {
		panic("La clave no pertenece al diccionario")
	}
	return *res
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cant
}

func (h *hashAbierto[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	seguir := true
	i := 0
	for i < len(h.tabla) && seguir {
		h.tabla[i].Iterar(func(par *parClaveValor[K, V]) bool {
			if !visitar(par.clave, par.valor) {
				seguir = false
				return false
			}
			return true
		})
		i++
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	return h.crearIterDiccionario()
}

func (iter *iteradorDiccionario[K, V]) HaySiguiente() bool {
	return iter.actLista.Largo() != 0 && iter.actPar.HaySiguiente()
}

func (iter *iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actPar.VerActual().clave, iter.actPar.VerActual().valor
}

func (iter *iteradorDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.actPar.Siguiente()
	if !iter.actPar.HaySiguiente() {
		iter.buscarProx(iter.act + 1)
	}
}
