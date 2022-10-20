package clase1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 10
	num2 := 5

	expected := 5

	actual := Restar(num1, num2)

	assert.Equal(t, expected, actual)
}
