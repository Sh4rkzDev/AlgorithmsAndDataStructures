package diccionario_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"
)

var cmpInts = func(a, b int) int { return a - b }

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestDiccionarioClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](cmpInts)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestReemplazoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	claves := [500]int{}
	for i := 0; i < 500; i++ {
		claves[i] = i
	}
	rand.Shuffle(500, func(i, j int) { claves[i], claves[j] = claves[j], claves[i] })
	for i := range claves {
		dic.Guardar(claves[i], claves[i])
	}
	for i := range claves {
		dic.Guardar(claves[i], claves[i]*2)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorrados(t *testing.T) {
	t.Log("Prueba para verificar que no haya problema reinsertar un elemento borrado")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](cmpInts)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dic := TDADiccionario.CrearABB[avanzado, int](func(a, b avanzado) int { return strings.Compare(a.z, b.z) })

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "chau", x: basico{a: "mundo", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) en orden con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cant := 0

	dic.Iterar(func(clave string, dato *int) bool {
		require.EqualValues(t, clave, claves[cant])
		cs[cant] = clave
		cant++
		return true
	})

	require.EqualValues(t, 3, cant)
	require.EqualValues(t, claves[0], cs[0])
	require.EqualValues(t, claves[1], cs[1])
	require.EqualValues(t, claves[2], cs[2])
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave1, 3)
	dic.Guardar(clave2, 5)
	dic.Guardar(clave3, 6)
	dic.Guardar(clave4, 2)
	dic.Guardar(clave5, 4)

	resultados := [5]int{2, 6, 24, 120, 720}
	factorial := 1
	cont := 0
	dic.Iterar(func(_ string, dato int) bool {
		factorial *= dato
		require.EqualValues(t, factorial, resultados[cont])
		cont++
		return true
	})

	require.EqualValues(t, 720, factorial)
	require.EqualValues(t, 5, cont)
}

func TestIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 3)
	dic.Guardar(clave2, 5)
	dic.Guardar(clave3, 6)
	dic.Guardar(clave4, 2)
	dic.Guardar(clave5, 4)

	dic.Borrar(clave0)

	resultados := [5]int{2, 6, 24, 120, 720}
	factorial := 1
	cont := 0
	dic.Iterar(func(_ string, dato int) bool {
		factorial *= dato
		require.EqualValues(t, factorial, resultados[cont])
		cont++
		return true
	})
	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumen(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)
	clavesOrd := make([]string, n)
	valoresOrd := make([]int, n)

	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		clavesOrd[i] = fmt.Sprintf("%08d", i)
		valoresOrd[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
		valores[i], valores[j] = valores[j], valores[i]
	})
	/* Inserta 'n' parejas en el ABB */
	for i := range claves {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(clavesOrd[i])
		require.EqualValues(b, dic.Obtener(clavesOrd[i]), valoresOrd[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(clavesOrd[i]) == valoresOrd[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionario(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario y que esten en orden. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primeroClave, primeroValor := iter.VerActual()
	require.EqualValues(t, "Gato", primeroClave)
	require.EqualValues(t, "miau", primeroValor)

	iter.Siguiente()
	segundoClave, segundoValor := iter.VerActual()
	require.EqualValues(t, "Perro", segundoClave)
	require.EqualValues(t, "guau", segundoValor)
	require.NotEqualValues(t, primeroClave, segundoClave)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	terceroClave, terceroValor := iter.VerActual()
	require.EqualValues(t, "Vaca", terceroClave)
	require.EqualValues(t, "moo", terceroValor)
	require.NotEqualValues(t, primeroClave, terceroClave)
	require.NotEqualValues(t, segundoClave, terceroClave)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[2], "")

	iter1 := dic.Iterador()
	clave1, _ := iter1.VerActual()
	require.EqualValues(t, "A", clave1)
	iter2 := dic.Iterador()
	iter2.Siguiente()
	clave1, _ = iter1.VerActual()
	require.EqualValues(t, "A", clave1)
	clave2, _ := iter2.VerActual()
	require.EqualValues(t, "B", clave2)
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.EqualValues(t, primero, clave1)
	require.EqualValues(t, segundo, clave2)
	require.EqualValues(t, tercero, "C")
}

func ejecutarPruebasVolumenIterador(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)
	clavesOrd := make([]string, n)
	valoresOrd := make([]int, n)

	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		clavesOrd[i] = fmt.Sprintf("%08d", i)
		valoresOrd[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
		valores[i], valores[j] = valores[j], valores[i]
	})
	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		valor = v1
		if clave == "" || valor == nil || clavesOrd[i] != clave || valoresOrd[i] != *valor {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	n := 10000

	claves := make([]int, n)
	clavesOrd := make([]int, n)

	for i := 0; i < n; i++ {
		claves[i] = i
		clavesOrd[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
	})
	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		dic.Guardar(claves[i], claves[i])
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false
	cont := 0

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c != 0 && c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		require.EqualValues(t, c, clavesOrd[cont])
		cont++
		return true
	})

	require.EqualValues(t, 100, cont)
	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIterarRangoCompleto(t *testing.T) {
	t.Log("Se itera sin poner rango alguno. Deberia iterar todos los elementos")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	cont := 0
	dic.IterarRango(nil, nil, func(a, b int) bool {
		cont++
		require.EqualValues(t, cont, a)
		require.EqualValues(t, cont, b)
		return true
	})
	require.EqualValues(t, cont, dic.Cantidad())
}

func TestIterarRangoCompletoCorte(t *testing.T) {
	t.Log("Se itera sin rango alguno y poniendo condicion de corte")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	cont := 0
	dic.IterarRango(nil, nil, func(a, b int) bool {
		if b > 5 {
			return false
		}
		require.EqualValues(t, a, cOrd[cont])
		require.EqualValues(t, b, cOrd[cont])
		cont++
		return true
	})
	require.EqualValues(t, 5, cont)
}

func TestIterarRangoMedio(t *testing.T) {
	t.Log("Prueba de iterar un cierto rango. Se deberia iterar desde el punto indicado hasta el otro, sin recorrer desde el principio y sin llegar al final")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], 10-c[i])
	}
	desde := 4
	hasta := 8
	suma := 0
	dic.IterarRango(&desde, &hasta, func(a, b int) bool {
		suma += b
		require.EqualValues(t, a, 10-b)
		require.EqualValues(t, a, cOrd[a-1])
		require.EqualValues(t, b, 10-cOrd[a-1])
		return true
	})
	require.EqualValues(t, 20, suma)
}

func TestIterarRangoMedioCorte(t *testing.T) {
	t.Log("Prueba de iterar un cierto rango con corte")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], 10-c[i])
	}
	desde := 4
	hasta := 8
	suma := 0
	dic.IterarRango(&desde, &hasta, func(a, b int) bool {
		if a == 7 {
			return false
		}
		suma += b
		require.EqualValues(t, a, 10-b)
		require.EqualValues(t, a, cOrd[a-1])
		require.EqualValues(t, b, 10-cOrd[a-1])
		return true
	})
	require.EqualValues(t, 15, suma)
}

func TestIterarRangoFuera(t *testing.T) {
	t.Log("Se itera con un rango fuera de los elementos. No deberia iterar ni una vez")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	desde := 15
	cont := 0
	dic.IterarRango(&desde, nil, func(a, b int) bool {
		cont++
		return true
	})
	require.EqualValues(t, 0, cont)
	hasta := 0
	dic.IterarRango(nil, &hasta, func(a, b int) bool {
		cont++
		return true
	})
	require.EqualValues(t, 0, cont)
}

func TestIterarRangoFueraCorte(t *testing.T) {
	t.Log("Se itera con un rango fuera de los elementos con corte. No deberia iterar ni una vez")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	desde := 15
	cont := 0
	dic.IterarRango(&desde, nil, func(a, b int) bool {
		if a != 2 {
			cont++
			return true
		}
		return false
	})
	require.EqualValues(t, 0, cont)
	hasta := 0
	dic.IterarRango(nil, &hasta, func(a, b int) bool {
		if a != 2 {
			cont++
			return true
		}
		return false
	})
	require.EqualValues(t, 0, cont)
}

func TestIterarRangoSinInicio(t *testing.T) {
	t.Log("Se itera sin poner un desde. Se deberia iterar desde el primer elemento")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	hasta := 5
	suma := 0
	cont := 0
	dic.IterarRango(nil, &hasta, func(a, b int) bool {
		suma += b
		require.EqualValues(t, a, cOrd[cont])
		require.EqualValues(t, b, cOrd[cont])
		cont++
		return true
	})
	require.EqualValues(t, 5, cont)
	require.EqualValues(t, 15, suma)
}

func TestIterarRangoSinInicioCorte(t *testing.T) {
	t.Log("Se itera sin poner un desde y con corte. Se deberia iterar desde el primer elemento")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	hasta := 5
	suma := 0
	cont := 0
	dic.IterarRango(nil, &hasta, func(a, b int) bool {
		if a == 1 {
			return false
		}
		suma += b
		require.EqualValues(t, a, cOrd[cont])
		require.EqualValues(t, b, cOrd[cont])
		cont++
		return true
	})
	require.EqualValues(t, 0, cont)
	require.EqualValues(t, 0, suma)
	dic.IterarRango(nil, &hasta, func(a, b int) bool {
		if a == 3 {
			return false
		}
		suma += b
		require.EqualValues(t, a, cOrd[cont])
		require.EqualValues(t, b, cOrd[cont])
		cont++
		return true
	})
	require.EqualValues(t, 2, cont)
	require.EqualValues(t, 3, suma)
}

func TestIterarRangoSinFinal(t *testing.T) {
	t.Log("Se itera sin poner un final. Se deberia iterar desde el indicado hasta el ultimo")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	desde := 7
	cont := 0
	suma := 0
	dic.IterarRango(&desde, nil, func(a, b int) bool {
		cont++
		require.EqualValues(t, a, cOrd[a-1])
		require.EqualValues(t, b, cOrd[a-1])
		suma += b
		return true
	})
	require.EqualValues(t, 4, cont)
	require.EqualValues(t, 34, suma)
}

func TestIterarRangoSinFinalCorte(t *testing.T) {
	t.Log("Se itera sin poner un final y con corte. Se deberia iterar desde el indicado hasta el ultimo o se cumpla la condicion de corte")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	desde := 7
	cont := 0
	suma := 0
	dic.IterarRango(&desde, nil, func(a, b int) bool {
		if a == 10 {
			return false
		}
		cont++
		require.EqualValues(t, a, cOrd[a-1])
		require.EqualValues(t, b, cOrd[a-1])
		suma += b
		return true
	})
	require.EqualValues(t, 3, cont)
	require.EqualValues(t, 24, suma)
}
func TestIteradorRangoCompleto(t *testing.T) {
	t.Log("Se usa el iterador externo sin poner rango alguno. Deberia iterar todos los elementos")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	i := 0
	iter := dic.IteradorRango(nil, nil)
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		require.EqualValues(t, clave, cOrd[i])
		require.EqualValues(t, valor, cOrd[i])
		i++
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoMedio(t *testing.T) {
	t.Log("Prueba de usar el iterador externo con un cierto rango. Se deberia iterar desde el punto indicado hasta el otro, sin recorrer desde el principio y sin llegar al final")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	desde := 2
	hasta := 6
	i := desde
	iter := dic.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		require.EqualValues(t, clave, cOrd[i-1])
		require.EqualValues(t, valor, cOrd[i-1])
		i += 1
		iter.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoFuera(t *testing.T) {
	t.Log("Se usa el iterador externo con un rango fuera de los elementos. No deberia iterar ni una vez")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}

	cont := 0

	desde := 11
	iter := dic.IteradorRango(&desde, nil)
	for iter.HaySiguiente() {
		cont++
		iter.Siguiente()
	}
	require.EqualValues(t, 0, cont)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })

	hasta := -3
	iter = dic.IteradorRango(nil, &hasta)
	for iter.HaySiguiente() {
		cont++
		iter.Siguiente()
	}
	require.EqualValues(t, 0, cont)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoSinInicio(t *testing.T) {
	t.Log("Se usa el iterador externo sin poner un desde. Se deberia iterar desde el primer elemento")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	hasta := 8
	cont := 0
	iter := dic.IteradorRango(nil, &hasta)
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		require.EqualValues(t, clave, cOrd[cont])
		require.EqualValues(t, valor, cOrd[cont])
		cont++
		iter.Siguiente()
	}
	require.EqualValues(t, 8, cont)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoSinFinal(t *testing.T) {
	t.Log("Se usa el iterador externo sin poner un final. Se deberia iterar desde el indicado hasta el ultimo")
	dic := TDADiccionario.CrearABB[int, int](cmpInts)
	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cOrd := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(10, func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	for i := range c {
		dic.Guardar(c[i], c[i])
	}
	desde := 9
	cont := 0
	aux := desde
	iter := dic.IteradorRango(&desde, nil)
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		require.EqualValues(t, clave, cOrd[aux-1])
		require.EqualValues(t, valor, cOrd[aux-1])
		aux++
		cont++
		iter.Siguiente()
	}
	require.EqualValues(t, 2, cont)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
