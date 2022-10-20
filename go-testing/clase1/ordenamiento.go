package clase1

import "sort"

func Ordenar(numeros []int) []int {
	sort.Ints(numeros)
	return numeros
}
