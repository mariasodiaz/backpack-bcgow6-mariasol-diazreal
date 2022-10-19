package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	numeros := []int{5, 2, 7, 9, 3}

	expected := []int{2, 3, 5, 7, 9}

	actual := Ordenar(numeros)

	assert.Equal(t, expected, actual)
}
