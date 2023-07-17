# Tabla de Hash

El trabajo que deben entregar de forma grupal es el tipo abstracto de datos Diccionario, con una implementación de Tabla de Hash, que se puede implementar mediante un hash abierto o cerrado, a elección. En caso de tratarse de un Hash abierto, pueden simplemente importar la lista por estar en el mismo módulo.

## Primitivas del Diccionario

```
type Diccionario[K comparable, V any] interface {
	Guardar(clave K, dato V)
	Pertenece(clave K) bool
	Obtener(clave K) V
	Borrar(clave K) V
	Cantidad() int
	Iterar(func(clave K, dato V) bool)
	Iterador() IterDiccionario[K, V]
}
```

Tanto **Borrar** como **Obtener** deben entrar en pánico con el mensaje 'La clave no pertenece al diccionario' en caso que la clave no se encuentre en el diccionario.

Además, la primitiva de creación del Hash deberá ser:

`func CrearHash[K comparable, V any]() Diccionario[K, V]`

Nuevamente, el iterador interno (Iterar) debe iterar internamente el hash, aplicando la función pasada por parámetro a la clave y los datos.

## Funcion de hashing… ¿Genérica?

Es de considerar que para implementar el hash será necesario definir una función de hashing internamente. Pueden definir la que más les guste (siempre poniendo referencia o nombre de la misma). Lamentablemente, no podemos trabajar de forma completamente genérica con una función de hashing directamente, por lo que deberemos realizar una transformación. Si bien no es obligatorio pasar la clave a un arreglo de bytes ([]byte), es lo recomendado. Luego, la función de hashing puede siempre trabajar con la versión de arreglo de bytes correspondiente a la clave. El siguiente código (que pueden utilizar, modificar, o lo que gusten) transforma un tipo de dato genérico a un array de bytes:

```
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
```

## Primitivas del iterador

```
type IterDiccionario[K comparable, V any] interface {
	HaySiguiente() bool
	VerActual() (K, V)
	Siguiente()
}
```

El iterador debe permitir recorrer todos los elementos almacenados en el hash, sin importar el orden en el que son devueltos.

Tanto **VerActual** como **Siguiente** deben entrar en pánico con el mensaje 'El iterador termino de iterar' si ya no quedan elementos a iterar (i.e. HaySiguiente() == false).

## Pruebas

Se adjunta, además, un archivo de pruebas que pueden utilizar para verificar que la estructura funciona correctamente.

Para ejecutar las pruebas, incluyendo las pruebas de volumen (benchmarks, que toman los tiempos y consumos de memoria), ejecutar:

`go test -bench=. -benchmem`
  
<br>
<br>
  
# Árbol Binario de Búsqueda

El trabajo que deben entregar de forma grupal es el tipo de dato abstracto Árbol Binario de Búsqueda (ABB), que es la implementación del tipo DiccionarioOrdenado (una extensión del Diccionario de la entrega anterior). Tanto el DiccionarioOrdenado como el ABB deben estar también dentro del paquete diccionario.

```
type DiccionarioOrdenado[K comparable, V any] interface {
	Diccionario[K, V]

	IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)
	IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
}
```

Todas las primitivas anteriores deben funcionar también, con el agregado que tanto el iterador interno (Iterar) como externo (Iterador) deben iterar en el orden que corresponda al ordenamiento del Diccionario. Se agregan las primitivas que permiten iterar por rangos dados. En caso que desde sea nil, se debe iterar desde la primera clave, y en caso de que hasta sea nil se debe iterar hasta la última (por lo cual, si desde == hasta == nil, se debe comportar como el iterador sin rango).

Además, la primitiva de creación del ABB deberá ser:

`func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccinarioOrdenado[K, V]`

La función de comparación, recibe dos claves y devuelve:

+ Un entero menor que 0 si la primera clave es menor que la segunda.
+ Un entero mayor que 0 si la primera clave es mayor que la segunda.
+ 0 si ambas claves son iguales.

Qué implica que una clave sea igual, mayor o menor que otra va a depender del usuario del TDA. Por ejemplo, strings.Compare cumple con esta especificación (si las claves son cadenas).
