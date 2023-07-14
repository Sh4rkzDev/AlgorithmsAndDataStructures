package cola

type nodo[T any] struct {
	data T
	next *nodo[T]
}

type colaEnlazada[T any] struct {
	prim *nodo[T]
	ult  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(colaEnlazada[T])
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.prim == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.prim.data
}

func crearNodo[T any](elem T) *nodo[T] {
	nod := new(nodo[T])
	nod.data = elem
	return nod
}

func (c *colaEnlazada[T]) Encolar(elem T) {
	nod := crearNodo[T](elem)
	if !c.EstaVacia() {
		c.ult.next = nod
	} else {
		c.prim = nod
	}
	c.ult = nod
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	res := c.VerPrimero()
	c.prim = c.prim.next
	return res
}
