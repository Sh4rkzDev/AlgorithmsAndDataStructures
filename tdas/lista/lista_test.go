package lista_test

import (
	"github.com/stretchr/testify/require"
	TDALista "tdas/lista"
	"testing"
)

func TestListaVacia(t *testing.T) {
	t.Log("Se hacen pruebas con una lista vacia")
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestListaUnElemento(t *testing.T) {
	t.Log("Se hacen pruebas con un elemento")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	lista.InsertarUltimo(-87)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, -87, lista.VerPrimero())
	require.EqualValues(t, -87, lista.VerUltimo())
	require.EqualValues(t, -87, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestListaAlgunosElementosPrimero(t *testing.T) {
	t.Log("Se hacen pruebas con algunos elementos")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarPrimero(i)
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i+1, lista.Largo())
	}
	require.EqualValues(t, 0, lista.VerUltimo())
	require.EqualValues(t, 10, lista.Largo())
	for i := 9; i >= 0; i-- {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.EqualValues(t, i, lista.Largo())
	}
	require.True(t, lista.EstaVacia())
}

func TestListaAlgunosElementosUltimo(t *testing.T) {
	t.Log("Se hacen pruebas con algunos elementos")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 10; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
		require.EqualValues(t, i+1, lista.Largo())
	}
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 10, lista.Largo())
	for i := 0; i < 10; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.EqualValues(t, 9-i, lista.Largo())
	}
	require.True(t, lista.EstaVacia())
}

func TestListaVolumen(t *testing.T) {
	t.Log("Se hacen pruebas de volumen")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 50000; i++ {
		lista.InsertarUltimo(i)
		require.EqualValues(t, i, lista.VerUltimo())
		require.EqualValues(t, i+1, lista.Largo())
	}
	require.EqualValues(t, 0, lista.VerPrimero())
	require.EqualValues(t, 50000, lista.Largo())
	for i := 0; i < 50000; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.EqualValues(t, 49999-i, lista.Largo())
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestListaStrings(t *testing.T) {
	t.Log("Se hacen pruebas con una lista de strings")
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("estas")
	lista.InsertarPrimero("como")
	lista.InsertarPrimero("Hola")
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, "Hola", lista.VerPrimero())
	require.EqualValues(t, "estas", lista.VerUltimo())
	require.EqualValues(t, "Hola", lista.BorrarPrimero())
	require.EqualValues(t, "como", lista.VerPrimero())
	require.EqualValues(t, "como", lista.BorrarPrimero())
	require.EqualValues(t, "estas", lista.VerPrimero())
	require.EqualValues(t, "estas", lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestIteradorListaVacia(t *testing.T) {
	t.Log("Se hacen pruebas con el iterador en una lista vacia")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.False(t, iter.HaySiguiente())
}

func TestIteradorIterar(t *testing.T) {
	t.Log("Se hacen pruebas verificando que el iterador externo itera correctamente la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < 100; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for i := 0; i < 100; i++ {
		require.EqualValues(t, i, iter.VerActual())
		iter.Siguiente()
	}
}

func TestIteradorInsertarPrincipioVacio(t *testing.T) {
	t.Log("Se hacen pruebas insertando un elemento al principio de una lista vacia con el iterador externo")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(-53)
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, -53, lista.VerPrimero())
	require.EqualValues(t, -53, lista.VerUltimo())
	require.EqualValues(t, -53, iter.VerActual())
	require.True(t, iter.HaySiguiente())
	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.False(t, iter.HaySiguiente())
}

func TestIteradorInsertarPrincipioVarios(t *testing.T) {
	t.Log("Se hacen pruebas insertando un elemento al principio de una lista con elementos con el iterador externo")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(9)
	iter := lista.Iterador()
	iter.Insertar(-53)
	require.EqualValues(t, -53, lista.VerPrimero())
	require.EqualValues(t, 3, lista.Largo())
}

func TestIteradorInsetarMedio(t *testing.T) {
	t.Log("Se hacen pruebas insertando un elemento en el medio con el iterador externo")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(6)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(8)
	iter.Insertar(9)
	iter2 := lista.Iterador()
	for i := 10; i >= 6; i-- {
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
	require.EqualValues(t, 5, lista.Largo())
}

func TestIteradorInsertarFinal(t *testing.T) {
	t.Log("Se hacen pruebas insertando un elemento al final de la lista con el iterador externo")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(9)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	iter.Insertar(8)
	require.EqualValues(t, 8, iter.VerActual())
	iter2 := lista.Iterador()
	for i := 10; i >= 8; i-- {
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
	require.EqualValues(t, 8, lista.VerUltimo())
}

func TestIteradorBorrarPrimero(t *testing.T) {
	t.Log("Se hacen pruebas borrando el primero elemento con el iterador externo")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(9)
	iter := lista.Iterador()
	require.EqualValues(t, 9, iter.Borrar())
	require.EqualValues(t, 10, lista.VerPrimero())
	require.EqualValues(t, 10, lista.VerUltimo())
	require.EqualValues(t, 10, iter.VerActual())
}

func TestIteradorBorrarMedio(t *testing.T) {
	t.Log("Se hacen pruebas borrando un elemento del medio con el iterador externo")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(9)
	lista.InsertarPrimero(8)
	iter := lista.Iterador()
	iter.Siguiente()
	require.EqualValues(t, 9, iter.Borrar())
	require.EqualValues(t, 8, lista.VerPrimero())
	require.EqualValues(t, 10, lista.VerUltimo())
	require.EqualValues(t, 10, iter.VerActual())
}

func TestIteradorBorrarFinal(t *testing.T) {
	t.Log("Se hacen pruebas borrando el ultimo elemento con el iterador externo")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(9)
	lista.InsertarPrimero(8)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	require.EqualValues(t, 10, iter.Borrar())
	require.EqualValues(t, 8, lista.VerPrimero())
	require.EqualValues(t, 9, lista.VerUltimo())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIteradorVolumen(t *testing.T) {
	t.Log("Se hacen pruebas de volumen para el iterador externo")
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	for i := 0; i < 50000; i++ {
		iter.Insertar(i)
		require.EqualValues(t, i+1, lista.Largo())
		require.EqualValues(t, i, lista.VerPrimero())
	}
	require.EqualValues(t, 50000, lista.Largo())
	for i := 50000; i > 0; i-- {
		require.EqualValues(t, i-1, lista.VerPrimero())
		require.EqualValues(t, i-1, iter.Borrar())
		require.EqualValues(t, i-1, lista.Largo())
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.False(t, iter.HaySiguiente())
}

func TestIterarTodo(t *testing.T) {
	t.Log("Se hacen pruebas iterando todos los elementos de la lista con el iterador interno")
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(9)
	lista.InsertarPrimero(8)
	suma := 0
	lista.Iterar(func(x int) bool {
		suma += x
		return true
	})
	require.EqualValues(t, 27, suma)
}

func TestIterarCorte(t *testing.T) {
	t.Log("Se hacen pruebas con un corte en la iteracion de la lista con el iterador interno")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 100; i++ {
		lista.InsertarUltimo(i)
	}
	contador := 0
	lista.Iterar(func(x int) bool {
		if x >= 50 {
			return false
		}
		contador++
		return true
	})
	require.EqualValues(t, 49, contador)
}

func TestIterarVolumen(t *testing.T) {
	t.Log("Se hacen pruebas de volumen con el iterador interno")
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 20000; i++ {
		lista.InsertarUltimo(i)
	}
	pares := 0
	lista.Iterar(func(x int) bool {
		if x%2 == 0 {
			pares++
		}
		return true
	})
	require.EqualValues(t, 10000, pares)
}
