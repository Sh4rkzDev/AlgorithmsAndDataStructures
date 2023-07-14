package lista

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	prim *nodoLista[T]
	ult  *nodoLista[T]
	cant int
}

type iterListaEnlazada[T any] struct {
	ant  *nodoLista[T]
	act  *nodoLista[T]
	list *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return new(listaEnlazada[T])
}

func crearNodo[T any](elem T, prox *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{elem, prox}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.prim == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevoNodo := crearNodo[T](elem, l.prim)
	if l.EstaVacia() {
		l.ult = nuevoNodo
	}
	l.prim = nuevoNodo
	l.cant++
}

func (l *listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevoNodo := crearNodo[T](elem, nil)
	if l.EstaVacia() {
		l.prim = nuevoNodo
	} else {
		l.ult.prox = nuevoNodo
	}
	l.ult = nuevoNodo
	l.cant++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	res := l.prim.dato
	l.prim = l.prim.prox
	l.cant--
	if l.prim == nil {
		l.ult = nil
	}
	return res
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.prim.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ult.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.cant
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	act := l.prim
	for act != nil && visitar(act.dato) {
		act = act.prox
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{act: l.prim, list: l}
}

func (i *iterListaEnlazada[T]) VerActual() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.act.dato
}

func (i *iterListaEnlazada[T]) HaySiguiente() bool {
	return i.act != nil
}

func (i *iterListaEnlazada[T]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	i.ant = i.act
	i.act = i.act.prox
}

func (i *iterListaEnlazada[T]) Insertar(elem T) {
	nuevoNodo := crearNodo[T](elem, i.act)
	if !i.HaySiguiente() {
		i.list.ult = nuevoNodo
	}
	if i.ant == nil {
		i.list.prim = nuevoNodo
	} else {
		i.ant.prox = nuevoNodo
	}
	i.act = nuevoNodo
	i.list.cant++
}

func (i *iterListaEnlazada[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	res := i.act.dato
	i.act = i.act.prox
	if !i.HaySiguiente() {
		i.list.ult = i.ant
	}
	if i.ant == nil {
		i.list.prim = i.act
	} else {
		i.ant.prox = i.act
	}
	i.list.cant--
	return res
}
