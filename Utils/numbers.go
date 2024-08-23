package Utils

import "strconv"

/*
*
EsNumero Funcion que regresa True si es un numero o False si no
*/
func EsNumero(cadena string) bool {
	if _, err := strconv.Atoi(cadena); err == nil {
		return true
	}
	if _, err := strconv.ParseFloat(cadena, 64); err == nil {
		return true
	}
	return false
}
