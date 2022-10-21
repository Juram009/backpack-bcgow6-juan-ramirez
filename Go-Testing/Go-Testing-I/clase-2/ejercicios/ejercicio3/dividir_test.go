package ejercicio3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	//Arrange
	num := 10
	den := 2
	expected := 5
	//Act
	res, err := Dividir(num, den)
	//Assert
	assert.Nil(t, err, "El denominador no puede ser 0")
	assert.Equal(t, expected, res)
}
