package grafo_test

import (
	"github.com/stretchr/testify/require"
	"tdas/grafo"
	"testing"
)

func TestGrafoVacioND(t *testing.T) {
	t.Log("Pruebas de grafo vacio no dirigido")
	g := grafo.CrearGrafo[int, int](false)
	require.EqualValues(t, 0, g.Cantidad())
	require.PanicsWithValue(t, "El vertice 1 no pertenece al grafo", func() { g.BorrarVertice(1) })
	require.False(t, g.Pertenece(33))
	require.EqualValues(t, 0, len(g.Vertices()))
}

func TestGrafoVacioD(t *testing.T) {
	t.Log("Pruebas de grafo vacio dirigido")
	g := grafo.CrearGrafo[int, int](true)
	require.EqualValues(t, 0, g.Cantidad())
	require.PanicsWithValue(t, "El vertice 1 no pertenece al grafo", func() { g.BorrarVertice(1) })
	require.False(t, g.Pertenece(33))
	require.EqualValues(t, 0, len(g.Vertices()))
}

func TestGrafoUnVerticeND(t *testing.T) {
	t.Log("Pruebas de grafo no dirigido con un vertice")
	g := grafo.CrearGrafo[int, int](false)
	g.AgregarVertice(55)
	require.EqualValues(t, 1, g.Cantidad())
	require.True(t, g.Pertenece(55))
	require.EqualValues(t, 1, len(g.Vertices()))
	require.EqualValues(t, 0, len(g.Adyacentes(55)))
	g.BorrarVertice(55)
	require.EqualValues(t, 0, g.Cantidad())
	require.False(t, g.Pertenece(55))
}

func TestGrafoUnVerticeD(t *testing.T) {
	t.Log("Pruebas de grafo dirigido con un vertice")
	g := grafo.CrearGrafo[int, int](false)
	g.AgregarVertice(55)
	require.EqualValues(t, 1, g.Cantidad())
	require.True(t, g.Pertenece(55))
	require.EqualValues(t, 1, len(g.Vertices()))
	require.EqualValues(t, 0, len(g.Adyacentes(55)))
	g.BorrarVertice(55)
	require.EqualValues(t, 0, g.Cantidad())
	require.False(t, g.Pertenece(55))
	require.PanicsWithValue(t, "El vertice 55 no pertenece al grafo", func() { g.AgregarArista(55, 8, 1) })
}

func TestGrafoNDAlgunosVerticesConectar(t *testing.T) {
	t.Log("Pruebas de grafo no dirigido conectando algunos vertices")
	g := grafo.CrearGrafo[int, int](false)
	for i := 1; i < 10; i++ {
		g.AgregarVertice(i)
		require.True(t, g.Pertenece(i))
	}
	g.AgregarArista(1, 5, 1)
	g.AgregarArista(4, 8, 1)
	g.AgregarArista(3, 9, 1)
	g.AgregarArista(2, 1, 1)
	g.AgregarArista(8, 5, 1)
	g.AgregarArista(6, 7, 1)
	g.AgregarArista(3, 2, 1)
	g.AgregarArista(1, 4, 1)
	require.True(t, g.EstanUnidos(1, 5))
	require.True(t, g.EstanUnidos(5, 1))
	require.True(t, g.EstanUnidos(4, 8))
	require.True(t, g.EstanUnidos(8, 4))
	require.True(t, g.EstanUnidos(3, 9))
	require.True(t, g.EstanUnidos(9, 3))
	require.True(t, g.EstanUnidos(2, 1))
	require.True(t, g.EstanUnidos(1, 2))
	require.True(t, g.EstanUnidos(8, 5))
	require.True(t, g.EstanUnidos(5, 8))
	require.True(t, g.EstanUnidos(6, 7))
	require.True(t, g.EstanUnidos(7, 6))
	require.True(t, g.EstanUnidos(3, 2))
	require.True(t, g.EstanUnidos(2, 3))
	require.True(t, g.EstanUnidos(1, 4))
	require.True(t, g.EstanUnidos(4, 1))
	require.EqualValues(t, 9, len(g.Vertices()))
	require.EqualValues(t, 3, len(g.Adyacentes(1)))
	require.EqualValues(t, 2, len(g.Adyacentes(2)))
	require.EqualValues(t, 2, len(g.Adyacentes(3)))
	require.EqualValues(t, 2, len(g.Adyacentes(4)))
	require.EqualValues(t, 2, len(g.Adyacentes(5)))
	require.EqualValues(t, 1, len(g.Adyacentes(6)))
	require.EqualValues(t, 1, len(g.Adyacentes(7)))
	require.EqualValues(t, 2, len(g.Adyacentes(8)))
	require.EqualValues(t, 1, len(g.Adyacentes(9)))
	require.PanicsWithValue(t, "El vertice 10 no pertenece al grafo", func() { g.AgregarArista(10, 9, 1) })
	require.PanicsWithValue(t, "El vertice 15 no pertenece al grafo", func() { g.AgregarArista(15, 8, 1) })
	require.PanicsWithValue(t, "El vertice 17 no pertenece al grafo", func() { g.AgregarArista(4, 17, 1) })
}

func TestGrafoDAlgunosVerticesConectar(t *testing.T) {
	t.Log("Pruebas de grafo dirigido conectando algunos vertices")
	g := grafo.CrearGrafo[int, int](true)
	for i := 1; i < 10; i++ {
		g.AgregarVertice(i)
		require.True(t, g.Pertenece(i))
	}
	g.AgregarArista(1, 5, 1)
	g.AgregarArista(4, 8, 1)
	g.AgregarArista(3, 9, 1)
	g.AgregarArista(2, 1, 1)
	g.AgregarArista(8, 5, 1)
	g.AgregarArista(6, 7, 1)
	g.AgregarArista(3, 2, 1)
	g.AgregarArista(1, 4, 1)
	require.True(t, g.EstanUnidos(1, 5))
	require.False(t, g.EstanUnidos(5, 1))
	require.True(t, g.EstanUnidos(4, 8))
	require.False(t, g.EstanUnidos(8, 4))
	require.True(t, g.EstanUnidos(3, 9))
	require.False(t, g.EstanUnidos(9, 3))
	require.True(t, g.EstanUnidos(2, 1))
	require.False(t, g.EstanUnidos(1, 2))
	require.True(t, g.EstanUnidos(8, 5))
	require.False(t, g.EstanUnidos(5, 8))
	require.True(t, g.EstanUnidos(6, 7))
	require.False(t, g.EstanUnidos(7, 6))
	require.True(t, g.EstanUnidos(3, 2))
	require.False(t, g.EstanUnidos(2, 3))
	require.True(t, g.EstanUnidos(1, 4))
	require.False(t, g.EstanUnidos(4, 1))
	require.EqualValues(t, 9, len(g.Vertices()))
	require.EqualValues(t, 9, g.Cantidad())
	require.EqualValues(t, 2, len(g.Adyacentes(1)))
	require.EqualValues(t, 1, len(g.Adyacentes(2)))
	require.EqualValues(t, 2, len(g.Adyacentes(3)))
	require.EqualValues(t, 1, len(g.Adyacentes(4)))
	require.EqualValues(t, 0, len(g.Adyacentes(5)))
	require.EqualValues(t, 1, len(g.Adyacentes(6)))
	require.EqualValues(t, 0, len(g.Adyacentes(7)))
	require.EqualValues(t, 1, len(g.Adyacentes(8)))
	require.EqualValues(t, 0, len(g.Adyacentes(9)))
}

func TestGrafoNDAlgunosVerticesDesconectar(t *testing.T) {
	t.Log("Pruebas de grafo no dirigido desconectando algunos vertices")
	g := grafo.CrearGrafo[int, int](false)
	for i := 1; i < 10; i++ {
		g.AgregarVertice(i)
		require.True(t, g.Pertenece(i))
	}
	g.AgregarArista(1, 5, 1)
	g.AgregarArista(4, 8, 1)
	g.AgregarArista(3, 9, 1)
	g.AgregarArista(2, 1, 1)
	g.AgregarArista(8, 5, 1)
	g.AgregarArista(6, 7, 1)
	g.AgregarArista(3, 2, 1)
	g.AgregarArista(1, 4, 1)
	g.BorrarArista(5, 1)
	require.False(t, g.EstanUnidos(1, 5))
	g.BorrarArista(8, 4)
	require.False(t, g.EstanUnidos(4, 8))
	g.BorrarArista(3, 9)
	require.False(t, g.EstanUnidos(9, 3))
	g.BorrarArista(2, 1)
	require.False(t, g.EstanUnidos(2, 1))
	g.BorrarArista(5, 8)
	require.False(t, g.EstanUnidos(8, 5))
	g.BorrarArista(7, 6)
	require.False(t, g.EstanUnidos(7, 6))
	g.BorrarArista(2, 3)
	require.False(t, g.EstanUnidos(2, 3))
	g.BorrarArista(1, 4)
	require.False(t, g.EstanUnidos(1, 4))
	require.EqualValues(t, 9, g.Cantidad())
	require.EqualValues(t, 0, len(g.Adyacentes(1)))
	require.EqualValues(t, 0, len(g.Adyacentes(2)))
	require.EqualValues(t, 0, len(g.Adyacentes(3)))
	require.EqualValues(t, 0, len(g.Adyacentes(4)))
	require.EqualValues(t, 0, len(g.Adyacentes(5)))
	require.EqualValues(t, 0, len(g.Adyacentes(6)))
	require.EqualValues(t, 0, len(g.Adyacentes(7)))
	require.EqualValues(t, 0, len(g.Adyacentes(8)))
	require.EqualValues(t, 0, len(g.Adyacentes(9)))
	require.PanicsWithValue(t, "El vertice 10 no pertenece al grafo", func() { g.BorrarArista(10, 9) })
	require.PanicsWithValue(t, "El vertice 15 no pertenece al grafo", func() { g.BorrarArista(15, 8) })
	require.PanicsWithValue(t, "El vertice 17 no pertenece al grafo", func() { g.BorrarArista(4, 17) })
}

func TestGrafoDAlgunosVerticesDesconectar(t *testing.T) {
	t.Log("Pruebas de grafo dirigido desconectando algunos vertices")
	g := grafo.CrearGrafo[int, int](true)
	for i := 1; i < 10; i++ {
		g.AgregarVertice(i)
		require.True(t, g.Pertenece(i))
	}
	g.AgregarArista(1, 5, 1)
	g.AgregarArista(4, 8, 1)
	g.AgregarArista(3, 9, 1)
	g.AgregarArista(2, 1, 1)
	g.AgregarArista(8, 5, 1)
	g.AgregarArista(6, 7, 1)
	g.AgregarArista(3, 2, 1)
	g.AgregarArista(1, 4, 1)
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.BorrarArista(5, 1) })
	g.BorrarArista(1, 5)
	require.False(t, g.EstanUnidos(1, 5))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.BorrarArista(8, 4) })
	g.BorrarArista(4, 8)
	require.False(t, g.EstanUnidos(4, 8))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.BorrarArista(9, 3) })
	g.BorrarArista(3, 9)
	require.False(t, g.EstanUnidos(3, 9))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.BorrarArista(1, 2) })
	g.BorrarArista(2, 1)
	require.False(t, g.EstanUnidos(2, 1))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.BorrarArista(5, 8) })
	g.BorrarArista(8, 5)
	require.False(t, g.EstanUnidos(8, 5))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.BorrarArista(7, 6) })
	g.BorrarArista(6, 7)
	require.False(t, g.EstanUnidos(6, 7))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.BorrarArista(2, 3) })
	g.BorrarArista(3, 2)
	require.False(t, g.EstanUnidos(3, 2))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.BorrarArista(4, 1) })
	g.BorrarArista(1, 4)
	require.False(t, g.EstanUnidos(1, 4))
	require.EqualValues(t, 9, g.Cantidad())
	require.EqualValues(t, 0, len(g.Adyacentes(1)))
	require.EqualValues(t, 0, len(g.Adyacentes(2)))
	require.EqualValues(t, 0, len(g.Adyacentes(3)))
	require.EqualValues(t, 0, len(g.Adyacentes(4)))
	require.EqualValues(t, 0, len(g.Adyacentes(5)))
	require.EqualValues(t, 0, len(g.Adyacentes(6)))
	require.EqualValues(t, 0, len(g.Adyacentes(7)))
	require.EqualValues(t, 0, len(g.Adyacentes(8)))
	require.EqualValues(t, 0, len(g.Adyacentes(9)))
}

func TestGrafoNDPesos(t *testing.T) {
	t.Log("Pruebas de grafo no dirigido de obtener pesos de las aristas")
	g := grafo.CrearGrafo[int, int](false)
	g.AgregarVertice(1)
	g.AgregarVertice(2)
	g.AgregarVertice(3)
	g.AgregarVertice(4)
	g.AgregarArista(1, 2, 5)
	g.AgregarArista(1, 3, 3)
	g.AgregarArista(3, 2, 9)
	g.AgregarArista(1, 4, 1)
	g.AgregarArista(4, 2, 2)
	require.EqualValues(t, 5, g.Peso(1, 2))
	require.EqualValues(t, 5, g.Peso(2, 1))
	require.EqualValues(t, 3, g.Peso(1, 3))
	require.EqualValues(t, 3, g.Peso(3, 1))
	require.EqualValues(t, 9, g.Peso(3, 2))
	require.EqualValues(t, 9, g.Peso(2, 3))
	require.EqualValues(t, 1, g.Peso(1, 4))
	require.EqualValues(t, 1, g.Peso(4, 1))
	require.EqualValues(t, 2, g.Peso(4, 2))
	require.EqualValues(t, 2, g.Peso(2, 4))
	g.BorrarArista(2, 1)
	g.BorrarArista(3, 1)
	g.BorrarArista(2, 4)
	require.True(t, g.Pertenece(1))
	require.True(t, g.Pertenece(2))
	require.True(t, g.Pertenece(3))
	require.True(t, g.Pertenece(4))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.Peso(1, 2) })
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.Peso(4, 2) })
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.Peso(1, 3) })
}

func TestGrafoDPesos(t *testing.T) {
	t.Log("Pruebas de grafo dirigido de obtener pesos de las aristas")
	g := grafo.CrearGrafo[int, int](true)
	g.AgregarVertice(1)
	g.AgregarVertice(2)
	g.AgregarVertice(3)
	g.AgregarVertice(4)
	g.AgregarArista(1, 2, 5)
	g.AgregarArista(1, 3, 3)
	g.AgregarArista(3, 2, 9)
	g.AgregarArista(1, 4, 1)
	g.AgregarArista(4, 2, 2)
	require.EqualValues(t, 5, g.Peso(1, 2))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.Peso(2, 1) })
	require.EqualValues(t, 3, g.Peso(1, 3))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.Peso(3, 1) })
	require.EqualValues(t, 9, g.Peso(3, 2))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.Peso(2, 3) })
	require.EqualValues(t, 1, g.Peso(1, 4))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.Peso(4, 1) })
	require.EqualValues(t, 2, g.Peso(4, 2))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.Peso(2, 4) })
}

func TestGrafoNDVariado(t *testing.T) {
	t.Log("Pruebas de agregar y borrar vertices y aristas")
	g := grafo.CrearGrafo[int, int](false)
	g.AgregarVertice(1)
	g.AgregarVertice(2)
	g.AgregarVertice(3)
	g.AgregarArista(2, 3, 4)
	g.AgregarArista(1, 2, 5)
	g.AgregarArista(1, 3, 2)
	g.BorrarVertice(1)
	require.False(t, g.Pertenece(1))
	require.PanicsWithValue(t, "El vertice 1 no pertenece al grafo", func() { g.EstanUnidos(1, 2) })
	require.PanicsWithValue(t, "El vertice 1 no pertenece al grafo", func() { g.EstanUnidos(1, 3) })
	require.PanicsWithValue(t, "El vertice 1 no pertenece al grafo", func() { g.EstanUnidos(3, 1) })
	require.PanicsWithValue(t, "El vertice 1 no pertenece al grafo", func() { g.BorrarArista(1, 3) })
	require.PanicsWithValue(t, "El vertice 1 no pertenece al grafo", func() { g.Peso(1, 3) })
	require.PanicsWithValue(t, "El vertice 1 no pertenece al grafo", func() { g.Peso(2, 1) })
	g.BorrarArista(3, 2)
	require.False(t, g.EstanUnidos(2, 3))
	require.PanicsWithValue(t, "La arista uniendo a los vertices indicados no existe", func() { g.Peso(3, 2) })
}

type miStruct struct {
	n, m int
}

type miStructError struct {
	n int
	s string
}

func TestGrafoComplejo(t *testing.T) {
	t.Log("Pruebas con un grafo con pesos de estructuras complejas")
	g := grafo.CrearGrafo[int, miStruct](false)
	g.AgregarVertice(1)
	g.AgregarVertice(2)
	g.AgregarVertice(3)
	g.AgregarArista(1, 2, miStruct{5, 6})
	g.AgregarArista(1, 3, miStruct{2, 9})
	g.AgregarArista(3, 2, miStruct{1, 4})

	a1 := g.Peso(2, 1)
	require.EqualValues(t, 5, a1.n)
	require.EqualValues(t, 6, a1.m)

	a2 := g.Peso(3, 1)
	require.EqualValues(t, 2, a2.n)
	require.EqualValues(t, 9, a2.m)

	a3 := g.Peso(3, 2)
	require.EqualValues(t, 1, a3.n)
	require.EqualValues(t, 4, a3.m)

	gB := grafo.CrearGrafo[int, miStructError](false)
	gB.AgregarVertice(1)
	gB.AgregarVertice(2)
	require.PanicsWithValue(t, "La estructura debe contener solo campos de tipo Int", func() { gB.AgregarArista(1, 2, miStructError{1, ""}) })
}

func TestGrafoSlice(t *testing.T) {
	t.Log("Pruebas con un grafo con pesos de slices")
	g := grafo.CrearGrafo[int, []int](false)
	g.AgregarVertice(1)
	g.AgregarVertice(2)
	g.AgregarVertice(3)
	g.AgregarArista(1, 2, []int{5, 6})
	g.AgregarArista(1, 3, []int{2, 9})
	g.AgregarArista(3, 2, []int{1, 4})

	a1 := g.Peso(2, 1)
	require.EqualValues(t, 5, a1[0])
	require.EqualValues(t, 6, a1[1])

	a2 := g.Peso(3, 1)
	require.EqualValues(t, 2, a2[0])
	require.EqualValues(t, 9, a2[1])

	a3 := g.Peso(3, 2)
	require.EqualValues(t, 1, a3[0])
	require.EqualValues(t, 4, a3[1])

	gB := grafo.CrearGrafo[int, []string](false)
	gB.AgregarVertice(1)
	gB.AgregarVertice(2)
	require.PanicsWithValue(t, "El arreglo debe ser de tipo Int", func() { gB.AgregarArista(1, 2, []string{"aaa", "bbb"}) })
}

func TestGrafoCadenas(t *testing.T) {
	t.Log("Pruebas con un grafo de strings")
	g := grafo.CrearGrafo[string, int](false)
	g.AgregarVertice("Hola")
	g.AgregarVertice("Mundo")
	g.AgregarVertice("Hello")
	g.AgregarVertice("World")
	g.AgregarArista("Hola", "Mundo", 1)
	g.AgregarArista("Hola", "World", 1)
	g.AgregarArista("Hello", "Mundo", 1)
	g.AgregarArista("Hello", "World", 1)
	g.AgregarArista("Hola", "Hello", 1)
	require.True(t, g.EstanUnidos("Hello", "World"))
	require.True(t, g.Pertenece("Hola"))
	require.False(t, g.EstanUnidos("Mundo", "World"))
	g.BorrarArista("Hello", "Hola")
	require.False(t, g.EstanUnidos("Hello", "Hola"))
}

func TestGrafoVertAzar(t *testing.T) {
	t.Log("Pruebas de obtener un vertice aleatorio valido")
	g := grafo.CrearGrafo[int, int](false)
	require.PanicsWithValue(t, "El grafo se encuentra vacio", func() { g.ObtenerVertice() })
	g.AgregarVertice(5)
	require.EqualValues(t, 5, g.ObtenerVertice())
	g.BorrarVertice(5)
	require.PanicsWithValue(t, "El grafo se encuentra vacio", func() { g.ObtenerVertice() })
}

func TestGrafoIterar(t *testing.T) {
	t.Log("Pruebas de iterar un grafo con su iterador, asegurandose de visitar todos los vertices una sola vez. No hay condicion de corte")
	g := grafo.CrearGrafo[int, int](false)
	g.AgregarVertice(1)
	g.AgregarVertice(2)
	g.AgregarVertice(3)
	g.AgregarVertice(4)
	g.AgregarVertice(5)
	g.AgregarVertice(6)
	cont := 0
	g.Iterar(func(i int) bool {
		cont++
		return true
	})
	require.EqualValues(t, g.Cantidad(), cont)
}

func TestGrafoIterarConCorte(t *testing.T) {
	t.Log("Pruebas de iterar un grafo con su iterador con condicion de corte")
	g := grafo.CrearGrafo[int, int](false)
	g.AgregarVertice(1)
	g.AgregarVertice(2)
	g.AgregarVertice(3)
	g.AgregarVertice(4)
	g.AgregarVertice(5)
	g.AgregarVertice(6)
	seguir := true
	noSeguir := false
	vis := 0
	noVis := g.Cantidad()
	g.Iterar(func(i int) bool {
		require.False(t, noSeguir)
		if i == 3 {
			seguir = false
			noSeguir = true
			return false
		}
		vis++
		noVis--
		return true
	})
	require.False(t, seguir)
	require.EqualValues(t, g.Cantidad(), vis+noVis)
	require.True(t, noSeguir)
}
