# Cola enlazada

La entrega es muy similar a la realizada para el TDA Pila. La función para crear la cola debe ser:

```
func CrearColaEnlazada[T any]() Cola[T] {
	// ...
}
```

La cola debe ser enlazada, es decir que en lugar de usar un arreglo, usa nodos enlazados, de los cuales se desencola el primero y se encola a continuación del último. Por tanto, se debe implementar solamente con punteros, y no se debe guardar en un campo el tamaño de la cola. El archivo a entregar para la implementación debe ser cola_enlazada.go, que se encuentre en el paquete cola.

Deben entregar también un archivo cola_test.go (que esté dentro del paquete cola_test) que haga las correspondientes pruebas unitarias, análogas a las pedidas para Pila, obviamente considerando que el invariante cambia a su opuesto (FIFO).

## Pruebas

Para correr las pruebas, ejecutar el siguiente comando:

`go test`
