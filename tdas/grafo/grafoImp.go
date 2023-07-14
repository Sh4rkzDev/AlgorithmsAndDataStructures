package grafo

import (
	"fmt"
	"reflect"
	dic "tdas/diccionario"
)

type grafo[V comparable, W interface{}] struct {
	vert dic.Diccionario[V, dic.Diccionario[V, W]]
	dir  bool
}

func noPertenece[V comparable](v V) string {
	return fmt.Sprintf("El vertice %+v no pertenece al grafo", v)
}

func verificarInts(w interface{}) {
	rv := reflect.TypeOf(w).Kind()
	if rv == reflect.Int {
		return
	}
	if rv == reflect.Struct {
		structInst := reflect.ValueOf(w)
		for i := 0; i < structInst.NumField(); i++ {
			if structInst.Field(i).Kind() != reflect.Int {
				panic("La estructura debe contener solo campos de tipo Int")
			}
		}
	}
	if rv == reflect.Array || rv == reflect.Slice {
		arr := reflect.ValueOf(w)
		if arr.Index(0).Kind() != reflect.Int {
			panic("El arreglo debe ser de tipo Int")
		}
	}
}

func CrearGrafo[V comparable, W any](dir bool) Grafo[V, W] {
	aux := dic.CrearHash[V, dic.Diccionario[V, W]]()
	return &grafo[V, W]{aux, dir}
}

func (g *grafo[V, W]) Cantidad() int {
	return g.vert.Cantidad()
}

func (g *grafo[V, W]) AgregarVertice(v V) {
	if g.Pertenece(v) {
		return
	}
	ady := dic.CrearHash[V, W]()
	g.vert.Guardar(v, ady)
}

func (g *grafo[V, W]) BorrarVertice(v V) {
	if !g.Pertenece(v) {
		panic(noPertenece(v))
	}
	g.vert.Borrar(v)
	g.vert.Iterar(func(vertice V, ady dic.Diccionario[V, W]) bool {
		if ady.Pertenece(v) {
			ady.Borrar(v)
		}
		return true
	})
}

func (g *grafo[V, W]) AgregarArista(origen, dest V, p W) {
	if !g.vert.Pertenece(origen) {
		panic(noPertenece(origen))
	}
	if !g.vert.Pertenece(dest) {
		panic(noPertenece(dest))
	}
	verificarInts(p)
	g.vert.Obtener(origen).Guardar(dest, p)
	if !g.dir {
		g.vert.Obtener(dest).Guardar(origen, p)
	}
}

func (g *grafo[V, W]) BorrarArista(origen, dest V) {
	if !g.vert.Pertenece(origen) {
		panic(noPertenece(origen))
	}
	if !g.vert.Pertenece(dest) {
		panic(noPertenece(dest))
	}
	if !g.vert.Obtener(origen).Pertenece(dest) {
		panic("La arista uniendo a los vertices indicados no existe")
	}
	g.vert.Obtener(origen).Borrar(dest)
	if !g.dir {
		g.vert.Obtener(dest).Borrar(origen)
	}
}

func (g *grafo[V, W]) EstanUnidos(origen, dest V) bool {
	if !g.vert.Pertenece(origen) {
		panic(noPertenece(origen))
	}
	if !g.vert.Pertenece(dest) {
		panic(noPertenece(dest))
	}
	return g.vert.Obtener(origen).Pertenece(dest)
}

func (g *grafo[V, W]) Peso(origen, dest V) W {
	if !g.vert.Pertenece(origen) {
		panic(noPertenece(origen))
	}
	if !g.vert.Pertenece(dest) {
		panic(noPertenece(dest))
	}
	if !g.vert.Obtener(origen).Pertenece(dest) {
		panic("La arista uniendo a los vertices indicados no existe")
	}
	return g.vert.Obtener(origen).Obtener(dest)
}

func (g *grafo[V, W]) Pertenece(v V) bool {
	return g.vert.Pertenece(v)
}

func (g *grafo[V, W]) Vertices() []V {
	res := make([]V, g.Cantidad())
	i := 0
	g.vert.Iterar(func(vertice V, ady dic.Diccionario[V, W]) bool {
		res[i] = vertice
		i++
		return true
	})
	return res
}

func (g *grafo[V, W]) Adyacentes(v V) []V {
	if !g.Pertenece(v) {
		panic(noPertenece[V](v))
	}
	res := make([]V, g.vert.Obtener(v).Cantidad())
	i := 0
	g.vert.Obtener(v).Iterar(func(ady V, p W) bool {
		res[i] = ady
		i++
		return true
	})
	return res
}

func (g *grafo[V, W]) ObtenerVertice() V {
	if g.Cantidad() == 0 {
		panic("El grafo se encuentra vacio")
	}
	var res V
	g.vert.Iterar(func(clave V, dato dic.Diccionario[V, W]) bool {
		res = clave
		return false
	})
	return res
}

func (g *grafo[V, W]) Iterar(visit func(V) bool) {
	g.vert.Iterar(func(vertice V, ady dic.Diccionario[V, W]) bool {
		return visit(vertice)
	})
}
