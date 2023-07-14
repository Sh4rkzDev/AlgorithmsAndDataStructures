package errors

import "fmt"

func errorString(er string) string {
	return fmt.Sprintf("Error en comando %s\n", er)
}

type ErrorCmd struct {
	Cmd string
}

func (e ErrorCmd) Error() string {
	return errorString(e.Cmd)
}

type ErrorArchivo struct{}

func (e ErrorArchivo) Error() string {
	return errorString("agregar_archivo")
}

type ErrorTablero struct{}

func (e ErrorTablero) Error() string {
	return errorString("ver_tablero")
}

type ErrorInfo struct{}

func (e ErrorInfo) Error() string {
	return errorString("info_vuelo")
}

type ErrorPrior struct{}

func (e ErrorPrior) Error() string {
	return errorString("prioridad_vuelos")
}

type ErrorSig struct{}

func (e ErrorSig) Error() string {
	return errorString("siguiente_vuelo")
}

type ErrorBorrar struct{}

func (e ErrorBorrar) Error() string {
	return errorString("borrar")
}
