package grafo

type Grafo[V comparable, W any] interface {

	//Devuelve la cantidad de Vertices del Grafo
	Cantidad() int

	//Agrega un vertice al grafo. En caso que ya exista, no hace nada y mantiene sus Adyacentes
	AgregarVertice(V)

	//Borra el vertice del grafo y todas las aristas que salen como que entran de el
	//En caso de no existir, entrara en panico con el mensaje "El vertice <nombre> no pertenece al grafo"
	BorrarVertice(V)

	//Agrega una arista entre los Vertices con el peso indicado. En caso de grafo dirigido, la arista tiene origen en el primer vertice y apunta al segundo vertice
	//En caso que ya exista dicha arista, actualizara su valor en caso de grafo pesado
	//En caso que no exista alguno de los vertices, entrara en panico con el mensaje "El vertice <nombre> no pertenece al grafo"
	AgregarArista(V, V, W)

	//Borra la arista entre los vertices indicados. En caso de grafo dirigido, solo borrara la arista que empieza en el primer vertice y termina en el segundo.
	//En caso de no existir alguno de los vertices o la arista, entrara en panico con el mensaje "El vertice <nombre> no pertenece al grafo" o
	//"La arista uniendo a los vertices indicados no existe" respectivamente
	BorrarArista(V, V)

	//Devuelve True en caso de que haya una arista uniendo dichos vertices. En caso de grafo dirigido, que una el primer vertice con el segundo
	//En caso que no exista alguno de los vertices, entrara en panico con el mensaje "El vertice <nombre> no pertenece al grafo"
	EstanUnidos(V, V) bool

	//Devuelve el peso de la arista que une los Vertices. En caso de grafo dirigido, que una el primer vertice con el segundo
	//En caso de no existir alguno de los vertices o la arista, entrara en panico con el mensaje "El vertice <nombre> no pertenece al grafo" o
	//"La arista uniendo a los vertices indicados no existe" respectivamente
	Peso(V, V) W

	//Devuelve True si existe un vertice con la clave ingresada
	Pertenece(V) bool

	//Devuelve un array de todos los vertices del Grafo
	Vertices() []V

	//Devuelve un array de todos los adyacentes al vertice. En caso de que no exista dicho vertice, entrara en panico
	Adyacentes(V) []V

	//Devuelve un vertice aleatorio. En caso de no haber vertices, entrara en panico con el mensaje "El grafo se encuentra vacio"
	ObtenerVertice() V

	//Itera todos los vertices del grafo aplicandoles la funcion pasada por parametro. En caso que la funcion devuelva false, termina la iteracion
	Iterar(func(V) bool)
}
