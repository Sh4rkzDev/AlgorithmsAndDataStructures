package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	t.Log("Se hacen pruebas con cola vacia")
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestColaConUnElemento(t *testing.T) {
	t.Log("Se hacen pruebas con un elemento")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(5)
	require.EqualValues(t, 5, cola.VerPrimero())
	require.False(t, cola.EstaVacia())
	elem := cola.Desencolar()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.EqualValues(t, 5, elem)
}

func TestColaConAlgunosElementos(t *testing.T) {
	t.Log("Se hacen pruebas con algunos elementos")
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		cola.Encolar(i)
	}
	for i := 1; i <= 5; i++ {
		require.EqualValues(t, i, cola.VerPrimero())
		elem := cola.Desencolar()
		require.EqualValues(t, i, elem)
	}
	require.True(t, cola.EstaVacia())
}

func TestColaConMuchosElementos(t *testing.T) {
	t.Log("Se hacen pruebas con muchos elementos")
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 1; i <= 1000; i++ {
		cola.Encolar(i)
	}
	require.False(t, cola.EstaVacia())
	for i := 1; i <= 1000; i++ {
		require.EqualValues(t, i, cola.VerPrimero())
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
}

func TestColaVolumen(t *testing.T) {
	t.Log("Se hacen pruebas con nivel alto de volumen")
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 1; i <= 50000; i++ {
		cola.Encolar(i)
	}
	for i := 1; i <= 50000; i++ {
		require.EqualValues(t, i, cola.VerPrimero())
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
}

func TestColaConCadenas(t *testing.T) {
	t.Log("Se hacen pruebas con cola de cadenas")
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("Hola")
	cola.Encolar("Como")
	cola.Encolar("Estas")
	require.EqualValues(t, "Hola", cola.Desencolar())
	require.EqualValues(t, "Como", cola.Desencolar())
	require.EqualValues(t, "Estas", cola.Desencolar())
}
