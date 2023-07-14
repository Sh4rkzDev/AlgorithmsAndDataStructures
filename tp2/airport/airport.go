package airport

type Airport interface {

	//Recibe por parametro ruta al archivo conteniendo informacion de los vuelos. En caso que ocurra algun error, lo devolvera, en caso contrario devolvera nil
	AgregarArchivo(string) error

	//Recibe la cantidad a mostrar de vuelos, el modo a mostrar(ascendente o descendente), y el rango de fechas
	//con el formato YYYY-MM-DDTHH:MM:SS. Devuelve un arreglo de strings con el formato "FECHA - ID", y un error en caso que haya ocurrido alguno, sino nil
	VerTablero(int, string, string, string) ([]string, error)

	//Recibe el codigo del vuelo a mostrar su informacion. Devuelve toda la informacion del vuelo como string y error en caso de haberlo, sino nil
	InfoVuelo(string) (string, error)

	//Recibe la cantidad de vuelos prioritarios a mostrar. Devuelve todos los vuelos ordenados por prioridad
	PrioridadVuelos(int) []string

	//Recibe el origen y destino de llegada, y la fecha con el formato YYYY-MM-DDTHH:MM:SS. En caso de no haber ningun vuelo con los datos ingrresados, devolvera un string vacio
	SiguienteVuelo(string, string, string) string

	//Borra los vuelos contenidos entre el rango de fechas ingresado del sistema. Devolvera un arreglo de strings con toda la informacion de los vuelos, y un error en caso de haberlo
	Borrar(string, string) ([]string, error)
}
