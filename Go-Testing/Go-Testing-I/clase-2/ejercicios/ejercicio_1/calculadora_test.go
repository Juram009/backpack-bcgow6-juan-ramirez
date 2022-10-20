package ejercicio_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResta(t *testing.T) {
	//Arrange
	num1 := 10
	num2 := 5
	expected := 55
	//Act
	res := Sub(num1, num2)
	//Assert
	assert.Equal(t, expected, res, "El numero resultado: %d, es distinto del esperado: %d", expected, res)
	/*
		if res != expected {
			t.Errorf("El numero resultado: %d, es distinto del esperado: %d", expected, res)
		}
	*/
}
