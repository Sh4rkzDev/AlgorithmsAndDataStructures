package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	t.Log("Se hacen pruebas con pila vacia")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaConUnElemento(t *testing.T) {
	t.Log("Se hacen pruebas con un elemento")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(5)
	require.EqualValues(t, 5, pila.VerTope())
	require.False(t, pila.EstaVacia())
	elem := pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.EqualValues(t, 5, elem)
}

func TestPilaConAlgunosElementos(t *testing.T) {
	t.Log("Se hacen pruebas con algunos elementos")
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 1; i <= 5; i++ {
		pila.Apilar(i)
	}
	for i := 5; i >= 1; i-- {
		require.EqualValues(t, i, pila.VerTope())
		elem := pila.Desapilar()
		require.EqualValues(t, i, elem)
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaConMuchosElementos(t *testing.T) {
	t.Log("Se hacen pruebas con muchos elementos")
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 1; i <= 1000; i++ {
		pila.Apilar(i)
	}
	require.False(t, pila.EstaVacia())
	for i := 1000; i >= 1; i-- {
		require.EqualValues(t, i, pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaVolumen(t *testing.T) {
	t.Log("Se hacen pruebas con nivel alto de volumen")
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 1; i <= 50000; i++ {
		pila.Apilar(i)
	}
	for i := 50000; i >= 1; i-- {
		require.EqualValues(t, i, pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaConCadenas(t *testing.T) {
	t.Log("Se hacen pruebas con pila de cadenas")
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Hola")
	pila.Apilar("Como")
	pila.Apilar("Estas")
	require.EqualValues(t, "Estas", pila.Desapilar())
	require.EqualValues(t, "Como", pila.Desapilar())
	require.EqualValues(t, "Hola", pila.Desapilar())
}
