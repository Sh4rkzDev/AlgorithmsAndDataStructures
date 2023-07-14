package lista

type Lista[T any] interface {

	//EstaVacia devuelve True si la lista no tiene elementos
	EstaVacia() bool

	//InsertarPrimero inserta el elemento al inicio de la lista
	InsertarPrimero(T)

	//InsertarUltimo inserta el elemento al final de la lista
	InsertarUltimo(T)

	//BorrarPrimero borra el primer elemento de la lista. En caso de la lista estar vacia, entrara en panico.
	BorrarPrimero() T

	//VerPrimero devuelve el valor del primer elemento de la lista. En caso de la lista estar vacia, entrara en panico.
	VerPrimero() T

	//VerUltimo devuelve el valor del ultimo elemento de la lista. En caso de la lista estar vacia, entrara en panico.
	VerUltimo() T

	//Largo devuelve la cantidad de elementos almacenados en la lista.
	Largo() int

	//Iterar itera toda la lista, aplicando la funcion pasada por parametro a cada uno de los elementos. Termina de iterar cuando se termine la lista o la funcion por parametro devuelva false.
	Iterar(visitar func(T) bool)

	//Iterador crea un iterador el cual se puede manipular.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	//VerAcutal devuelve el valor del elemento en el que esta situado el iterador. En caso de haber terminado de iterar toda la lista, entrara en panico.
	VerActual() T

	//HaySiguiente devuelve true en caso de que no se haya terminado de iterar la lista.
	HaySiguiente() bool

	//Siguiente itera al siguiente elemento de la lista. En caso de haber terminado de iterar la lista, entrara en panico
	Siguiente()

	//Insertar insertara un elemento en la lista en la posicion actual que se encuentre el iterador.
	Insertar(T)

	//Borrar borrara el elemento en el cual este posicionado el iterador y lo devolvera. En caso de haber terminado de iterar la lista, entrara en panico.
	Borrar() T
}
