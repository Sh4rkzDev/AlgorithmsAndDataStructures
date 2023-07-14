package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	Heap "tdas/cola_prioridad"
	"testing"
)

const volumen = 100000

func cmpInt(a, b int) int {
	return a - b
}

func TestHeapVacio(t *testing.T) {
	t.Log("Pruebas de un heap vacio")
	h := Heap.CrearHeap(cmpInt)
	require.True(t, h.EstaVacia())
	require.EqualValues(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
}

func TestHeapUnElemento(t *testing.T) {
	t.Log("Pruebas con un elemento")
	h := Heap.CrearHeap(cmpInt)
	h.Encolar(1)
	require.False(t, h.EstaVacia())
	require.EqualValues(t, 1, h.Cantidad())
	require.EqualValues(t, 1, h.VerMax())
	require.EqualValues(t, 1, h.Desencolar())
	require.True(t, h.EstaVacia())
	require.EqualValues(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
}

func TestHeapPocosElementos(t *testing.T) {
	t.Log("Pruebas con pocos elementos")
	h := Heap.CrearHeap(cmpInt)
	for i := 1; i <= 5; i++ {
		h.Encolar(i)
		require.EqualValues(t, i, h.Cantidad())
		require.EqualValues(t, i, h.VerMax())
	}
	for i := 5; i >= 1; i-- {
		require.EqualValues(t, i, h.VerMax())
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i-1, h.Cantidad())
	}
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestHeapVariosElementos(t *testing.T) {
	t.Log("Pruebas con varios elementos")
	h := Heap.CrearHeap(cmpInt)
	for i := 1; i <= 50; i++ {
		h.Encolar(i)
		require.EqualValues(t, i, h.Cantidad())
		require.EqualValues(t, i, h.VerMax())
	}
	require.EqualValues(t, 50, h.Cantidad())
	require.EqualValues(t, 50, h.VerMax())
	for i := 50; i >= 1; i-- {
		require.EqualValues(t, i, h.VerMax())
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i-1, h.Cantidad())
	}
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestHeapVariosElementosDesord(t *testing.T) {
	t.Log("Pruebas con varios elementos encolando desordenadamente")
	aux := make([]int, 50)
	for i := 0; i < 50; i++ {
		aux[i] = i + 1
	}
	rand.Shuffle(50, func(a, b int) { aux[a], aux[b] = aux[b], aux[a] })
	h := Heap.CrearHeap(cmpInt)
	for i := 0; i < 50; i++ {
		h.Encolar(aux[i])
		require.EqualValues(t, i+1, h.Cantidad())
	}
	require.EqualValues(t, 50, h.Cantidad())
	require.EqualValues(t, 50, h.VerMax())
	for i := 50; i >= 1; i-- {
		require.EqualValues(t, i, h.VerMax())
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i-1, h.Cantidad())
	}
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestHeapVolumen(t *testing.T) {
	t.Log("Pruebas de Volumen")
	aux := make([]int, volumen)
	for i := 0; i < volumen; i++ {
		aux[i] = i + 1
	}
	rand.Shuffle(volumen, func(a, b int) { aux[a], aux[b] = aux[b], aux[a] })

	h := Heap.CrearHeap(cmpInt)
	for i := 0; i < volumen; i++ {
		h.Encolar(aux[i])
		require.EqualValues(t, i+1, h.Cantidad())
	}
	for i := volumen; i >= 1; i-- {
		require.EqualValues(t, i, h.VerMax())
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i-1, h.Cantidad())
	}
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestHeapVariado(t *testing.T) {
	t.Log("Pruebas de encolar y desencolar muchas veces")
	h := Heap.CrearHeap(cmpInt)
	h.Encolar(9)
	h.Encolar(3)
	h.Encolar(5)
	h.Encolar(1)
	h.Encolar(2)
	h.Encolar(8)
	require.EqualValues(t, 9, h.Desencolar())
	require.EqualValues(t, 8, h.Desencolar())
	require.EqualValues(t, 5, h.Desencolar())
	h.Encolar(10)
	require.EqualValues(t, 10, h.VerMax())
	require.EqualValues(t, 10, h.Desencolar())
	require.EqualValues(t, 3, h.Desencolar())
	h.Encolar(7)
	h.Encolar(7)
	h.Encolar(7)
	require.EqualValues(t, 7, h.Desencolar())
	require.EqualValues(t, 7, h.Desencolar())
	require.EqualValues(t, 7, h.Desencolar())
	require.EqualValues(t, 2, h.Desencolar())
	require.EqualValues(t, 1, h.Desencolar())
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestHeapStrings(t *testing.T) {
	t.Log("Pruebas de un Heap con Strings")
	h := Heap.CrearHeap(strings.Compare)
	h.Encolar("abcdef")
	h.Encolar("abc")
	h.Encolar("abcdefg")
	h.Encolar("a")
	h.Encolar("abcd")
	require.EqualValues(t, "abcdefg", h.Desencolar())
	require.EqualValues(t, "abcdef", h.Desencolar())
	require.EqualValues(t, "abcd", h.Desencolar())
	require.EqualValues(t, "abc", h.Desencolar())
	require.EqualValues(t, "a", h.Desencolar())
}

func TestHeapConArregloVacio(t *testing.T) {
	t.Log("Pruebas de crear un Heap a partir de un arreglo vacio")
	arr := make([]int, 0)
	h := Heap.CrearHeapArr(arr, cmpInt)
	require.True(t, h.EstaVacia())
	require.EqualValues(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
}

func TestHeapConArregloVacioInsertando(t *testing.T) {
	t.Log("Pruebas de crear un Heap a partir de un arreglo vacio y despues insertando elementos")
	arr := make([]int, 0)
	h := Heap.CrearHeapArr(arr, cmpInt)
	h.Encolar(50)
	h.Encolar(39)
	h.Encolar(-7)
	h.Encolar(84)
	require.EqualValues(t, 84, h.VerMax())
	require.EqualValues(t, 4, h.Cantidad())
	require.False(t, h.EstaVacia())
	require.EqualValues(t, 84, h.Desencolar())
	require.EqualValues(t, 3, h.Cantidad())
	require.EqualValues(t, 50, h.VerMax())
	require.EqualValues(t, 50, h.Desencolar())
	require.EqualValues(t, 2, h.Cantidad())
	require.EqualValues(t, 39, h.VerMax())
	require.EqualValues(t, 39, h.Desencolar())
	require.EqualValues(t, -7, h.VerMax())
	require.EqualValues(t, 1, h.Cantidad())
	require.EqualValues(t, -7, h.Desencolar())
	require.True(t, h.EstaVacia())
	require.EqualValues(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
}

func TestHeapConArreglo(t *testing.T) {
	t.Log("Pruebas de crear un Heap a partir de un arreglo")
	arr := make([]int, 15)
	for i := 1; i <= 15; i++ {
		arr[i-1] = i
	}
	h := Heap.CrearHeapArr(arr, cmpInt)
	require.EqualValues(t, 15, h.VerMax())
	require.EqualValues(t, 15, h.Cantidad())
	for i := 15; i >= 1; i-- {
		require.EqualValues(t, i, h.VerMax())
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i-1, h.Cantidad())
	}
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestHeapConArregloDesord(t *testing.T) {
	t.Log("Pruebas de crear un Heap a partir de un arreglo desordenado")
	arr := make([]int, 15)
	for i := 1; i <= 15; i++ {
		arr[i-1] = i
	}
	rand.Shuffle(15, func(a, b int) { arr[a], arr[b] = arr[b], arr[a] })
	h := Heap.CrearHeapArr(arr, cmpInt)
	require.EqualValues(t, 15, h.VerMax())
	require.EqualValues(t, 15, h.Cantidad())
	for i := 15; i >= 1; i-- {
		require.EqualValues(t, i, h.VerMax())
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i-1, h.Cantidad())
	}
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestHeapConArregloNoMod(t *testing.T) {
	t.Log("Verifica que crear un heap a partir de un arreglo no modifica el arreglo original")
	arr := []int{8, 4, 5, 3, 1, 9, 2, 7, 6, 10}
	_ = Heap.CrearHeapArr(arr, cmpInt)
	require.EqualValues(t, 8, arr[0])
	require.EqualValues(t, 4, arr[1])
	require.EqualValues(t, 5, arr[2])
	require.EqualValues(t, 3, arr[3])
	require.EqualValues(t, 1, arr[4])
	require.EqualValues(t, 9, arr[5])
	require.EqualValues(t, 2, arr[6])
	require.EqualValues(t, 7, arr[7])
	require.EqualValues(t, 6, arr[8])
	require.EqualValues(t, 10, arr[9])
}

func TestHeapConArregloVolumen(t *testing.T) {
	t.Log("Pruebas de volumen de crear un Heap a partir de un arreglo")
	arr := make([]int, volumen)
	for i := 1; i <= volumen; i++ {
		arr[i-1] = i
	}
	rand.Shuffle(volumen, func(a, b int) { arr[a], arr[b] = arr[b], arr[a] })
	h := Heap.CrearHeapArr(arr, cmpInt)
	require.EqualValues(t, volumen, h.VerMax())
	require.EqualValues(t, volumen, h.Cantidad())
	for i := volumen; i >= 1; i-- {
		require.EqualValues(t, i, h.VerMax())
		require.EqualValues(t, i, h.Desencolar())
		require.EqualValues(t, i-1, h.Cantidad())
	}
	require.EqualValues(t, 0, h.Cantidad())
	require.True(t, h.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	t.Log("Pruebas de ordenar un arreglo con el algoritmo de HeapSort")
	arr := make([]int, 15)
	for i := 1; i <= 15; i++ {
		arr[i-1] = i
	}
	rand.Shuffle(15, func(a, b int) { arr[a], arr[b] = arr[b], arr[a] })
	Heap.HeapSort(arr, cmpInt)
	for i := 1; i <= 15; i++ {
		require.EqualValues(t, i, arr[i-1])
	}
}
