package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num1 := 10
	num2 := 0

	expectedError := fmt.Sprintf("el denominador no puede ser 0")

	_, actual := Dividir(num1, num2)

	assert.NotNil(t, actual)
	assert.ErrorContains(t, actual, expectedError)
}
