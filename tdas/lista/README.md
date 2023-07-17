# Lista enlazada

Primitivas de la lista

```
type Lista[T any] interface {
	EstaVacia() bool
	InsertarPrimero(T)
	InsertarUltimo(T)
	BorrarPrimero() T
	VerPrimero() T
	VerUltimo() T
	Largo() int
	Iterar(visitar func(T) bool)
	Iterador() IteradorLista[T]
}
```

En caso que se invoque a BorrarPrimero, VerPrimero o VerUltimo sobre una lista vacía, todas deben entrar en pánico con un mensaje "La lista esta vacia".

Además, es necesario tener la primitiva de creación de la lista enlazada (en lista_enlazada.go):

```
func CrearListaEnlazada[T any]() Lista[T] {
	//...
}
```

### Primitiva del iterador interno

Como está indicado entre las primitivas, se debe implementar el iterador interno cuya firma es:

`Iterar(visitar func(T) bool)`

Dicha función debe aplicarse a cada uno de los datos de la lista (de primero a último), hasta que la lista se termine o la función visitar devuelva false (lo que ocurra primero).
Primitivas del iterador externo

La primitiva Iterador de la lista debe devolver un IteradorLista, cuyas primitivas son:

```
type IteradorLista[T any] interface {
	VerActual() T
	HaySiguiente() bool
	Siguiente()
	Insertar(T)
	Borrar() T
}
```

En caso que se invoque a VerActual, Siguiente o Borrar sobre un iterador que ya haya iterado todos los elementos, debe entrar en pánico con un mensaje "El iterador termino de iterar".

## Pruebas

Considerar que todas las primitivas (exceptuando Iterar) deben funcionar en tiempo constante.

Las pruebas deben incluir los casos básicos de TDA similares a los contemplados para la pila y la cola, y adicionalmente debe verificar los siguientes casos del iterador externo:

  1. Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
  2. Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
  3. Insertar un elemento en el medio se hace en la posición correcta.
  4. Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
  5. Remover el último elemento con el iterador cambia el último de la lista.
  6. Verificar que al remover un elemento del medio, este no está.
  7. Otros casos borde que pueden encontrarse al utilizar el iterador externo.
  8. Casos del iterador interno, incluyendo casos con corte (la función visitar devuelve false eventualmente).
