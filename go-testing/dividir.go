package main

import (
	"fmt"
)

func Dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("el denominador no puede ser 0")
	}
	return a / b, nil
}
