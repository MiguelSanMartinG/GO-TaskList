package Utils

import "strconv"

func EsNumero(cadena string) bool {
	if _, err := strconv.Atoi(cadena); err == nil {
		return true
	}
	if _, err := strconv.ParseFloat(cadena, 64); err == nil {
		return true
	}
	return false
}
