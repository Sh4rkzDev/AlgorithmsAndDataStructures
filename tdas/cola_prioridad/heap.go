package cola_prioridad

const (
	tamRed = 2
	tamIni = 7
	tamMin = 4
)

type heap[T any] struct {
	datos []T
	cmp   func(T, T) int
	cant  int
}

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	arr := make([]T, tamIni)
	return &heap[T]{datos: arr, cmp: cmp}
}

func upheap[T any](arr []T, pos int, cmp func(T, T) int) {
	if pos == 0 {
		return
	}
	padre := (pos - 1) / 2
	if cmp(arr[padre], arr[pos]) < 0 {
		arr[padre], arr[pos] = arr[pos], arr[padre]
		upheap(arr, padre, cmp)
	}
}

func downheap[T any](arr []T, pos, cant int, cmp func(T, T) int) {
	hijoIzq := 2*pos + 1
	hijoDer := 2*pos + 2
	if (hijoIzq >= cant || cmp(arr[pos], arr[hijoIzq]) > 0) && (hijoDer >= cant || cmp(arr[pos], arr[hijoDer]) > 0) {
		return
	}
	hMax := hijoMax(arr, hijoIzq, hijoDer, cant, cmp)
	arr[pos], arr[hMax] = arr[hMax], arr[pos]
	downheap(arr, hMax, cant, cmp)
}

func hijoMax[T any](arr []T, a, b, cant int, cmp func(T, T) int) int {
	if b >= cant || cmp(arr[a], arr[b]) > 0 {
		return a
	}
	return b
}

func heapify[T any](arr []T, cant int, cmp func(T, T) int) {
	mitad := cant / 2
	for i := mitad; i >= 0; i-- {
		downheap(arr, i, cant, cmp)
	}
}

func CrearHeapArr[T any](arr []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	cop := make([]T, len(arr))
	copy(cop, arr)
	heapify(cop, len(cop), funcion_cmp)
	if len(cop) < tamIni {
		aux := make([]T, tamIni)
		copy(aux, cop)
		return &heap[T]{aux, funcion_cmp, len(cop)}
	}
	return &heap[T]{cop, funcion_cmp, len(cop)}
}

func (h *heap[T]) Cantidad() int {
	return h.cant
}

func (h *heap[T]) EstaVacia() bool {
	return h.Cantidad() == 0
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) verificarRedimension() (bool, int) {
	if h.cant == len(h.datos) {
		return true, len(h.datos) * tamRed
	}
	if h.cant <= len(h.datos)/tamMin && len(h.datos) > tamIni {
		return true, len(h.datos) / tamRed
	}
	return false, 0
}

func (h *heap[T]) redimensionar(tam int) {
	newH := make([]T, tam)
	copy(newH, h.datos)
	h.datos = newH
}

func (h *heap[T]) Encolar(elem T) {
	h.datos[h.cant] = elem
	upheap(h.datos, h.cant, h.cmp)
	h.cant++
	verif, tam := h.verificarRedimension()
	if verif {
		h.redimensionar(tam)
	}
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	h.datos[0], h.datos[h.cant-1] = h.datos[h.cant-1], h.datos[0]
	res := h.datos[h.cant-1]
	h.cant--
	downheap(h.datos, 0, h.cant, h.cmp)
	verif, tam := h.verificarRedimension()
	if verif {
		h.redimensionar(tam)
	}
	return res
}

func HeapSort[T any](elems []T, cmp func(T, T) int) {
	heapify(elems, len(elems), cmp)
	for i := len(elems) - 1; i > 0; i-- {
		downheap(elems, 0, i+1, cmp)
		elems[0], elems[i] = elems[i], elems[0]
	}
}
