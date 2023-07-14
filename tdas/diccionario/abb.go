package diccionario

import Pila "tdas/pila"

type nodoABB[K comparable, V any] struct {
	clave    K
	valor    V
	izq, der *nodoABB[K, V]
}

type abb[K comparable, V any] struct {
	root *nodoABB[K, V]
	cmp  func(K, K) int
	cant int
}

type iterABB[K comparable, V any] struct {
	nodos        *Pila.Pila[*nodoABB[K, V]]
	desde, hasta *K
	cmp          func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}

func crearNodoABB[K comparable, V any](clave K, valor V) *nodoABB[K, V] {
	return &nodoABB[K, V]{clave: clave, valor: valor}
}

func (a *abb[K, V]) Cantidad() int {
	return a.cant
}

func (a *abb[K, V]) buscar(nodo **nodoABB[K, V], aux func(**nodoABB[K, V]) *nodoABB[K, V], clave *K, valor *V) (bool, **nodoABB[K, V]) {
	if *nodo == nil {
		return false, nodo
	}
	if a.cmp(*clave, (*nodo).clave) == 0 {
		res := aux(nodo)
		return true, &res
	}
	if a.cmp(*clave, (*nodo).clave) < 0 {
		return a.buscar(&(*nodo).izq, aux, clave, valor)
	}
	return a.buscar(&(*nodo).der, aux, clave, valor)
}

func (a *abb[K, V]) Guardar(clave K, valor V) {
	found, act := a.buscar(&a.root, func(nodo **nodoABB[K, V]) *nodoABB[K, V] {
		(*nodo).valor = valor
		return nil
	}, &clave, &valor)
	if found {
		return
	}
	nodo := crearNodoABB[K, V](clave, valor)
	*act = nodo
	a.cant++
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	found, _ := a.buscar(&a.root, func(nodo **nodoABB[K, V]) *nodoABB[K, V] {
		return nil
	}, &clave, nil)
	return found
}

func (a *abb[K, V]) Obtener(clave K) V {
	found, res := a.buscar(&a.root, func(nodo **nodoABB[K, V]) *nodoABB[K, V] {
		return *nodo
	}, &clave, nil)
	if found {
		return (*res).valor
	}
	panic("La clave no pertenece al diccionario")
}

func (a *abb[K, V]) Borrar(clave K) V {
	found, res := a.buscar(&a.root, func(nodo **nodoABB[K, V]) *nodoABB[K, V] {
		res := **nodo
		if (*nodo).der == nil && (*nodo).izq == nil {
			//Ningun hijo
			*nodo = nil
		} else if (*nodo).der == nil {
			//Hijo Izq
			*nodo = (*nodo).izq
		} else if (*nodo).izq == nil {
			//Hijo Der
			*nodo = (*nodo).der
		} else if (*nodo).izq != nil && (*nodo).der != nil {
			//Dos Hijos
			claveRe := (*nodo).izq.reemplazante()
			valorRe := a.Borrar(claveRe)
			a.cant++ //para compensar la resta de despues y que no reste dos veces cuando solo se elimina un solo elemento
			(*nodo).clave = claveRe
			(*nodo).valor = valorRe
		}
		a.cant--
		return &res
	}, &clave, nil)
	if found {
		return (*res).valor
	}
	panic("La clave no pertenece al diccionario")
}

func (nodo *nodoABB[K, V]) reemplazante() K {
	res := nodo
	for res.der != nil {
		res = res.der
	}
	return res.clave
}

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iterar[K, V](visitar, a.root)
}

func iterar[K comparable, V any](visitar func(clave K, dato V) bool, nodo *nodoABB[K, V]) bool {
	if nodo == nil {
		return true
	}
	if !iterar(visitar, nodo.izq) || !visitar(nodo.clave, nodo.valor) || !iterar(visitar, nodo.der) {
		return false
	}
	return true
}

func (a *abb[K, V]) IterarRango(desde, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRango(a.root, desde, hasta, visitar, a.cmp)
}

func iterarRango[K comparable, V any](nodo *nodoABB[K, V], desde, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	cont := true
	if desde == nil || cmp(*desde, nodo.clave) < 0 {
		cont = iterarRango(nodo.izq, desde, hasta, visitar, cmp)
	}
	if cont && (desde == nil || cmp(*desde, nodo.clave) <= 0) && (hasta == nil || cmp(*hasta, nodo.clave) >= 0) {
		cont = visitar(nodo.clave, nodo.valor)
	}
	if !cont {
		return false
	}
	if hasta == nil || cmp(*hasta, nodo.clave) > 0 {
		cont = iterarRango(nodo.der, desde, hasta, visitar, cmp)
	}
	return cont
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := Pila.CrearPilaDinamica[*nodoABB[K, V]]()
	a.root.buscarEnRango(&pila, desde, hasta, a.cmp)
	return &iterABB[K, V]{&pila, desde, hasta, a.cmp}
}

func (nodo *nodoABB[K, V]) buscarEnRango(pila *Pila.Pila[*nodoABB[K, V]], desde, hasta *K, cmp func(K, K) int) {
	if nodo == nil {
		return
	}
	if desde == nil || cmp(*desde, nodo.clave) <= 0 {
		(*pila).Apilar(nodo)
		nodo.izq.buscarEnRango(pila, desde, hasta, cmp)
	} else if hasta == nil || cmp(*hasta, nodo.clave) > 0 {
		nodo.der.buscarEnRango(pila, desde, hasta, cmp)
	}
}

func (i *iterABB[K, V]) HaySiguiente() bool {
	return !(*i.nodos).EstaVacia() && (i.hasta == nil || i.cmp(*i.hasta, (*i.nodos).VerTope().clave) >= 0)
}

func (i *iterABB[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return (*i.nodos).VerTope().clave, (*i.nodos).VerTope().valor
}

func (i *iterABB[K, V]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elem := (*i.nodos).Desapilar()
	if elem.der == nil {
		return
	}
	elem.der.buscarEnRango(i.nodos, i.desde, i.hasta, i.cmp)
}
