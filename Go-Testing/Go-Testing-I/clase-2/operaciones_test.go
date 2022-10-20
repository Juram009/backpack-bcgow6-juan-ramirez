package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	//Arrange
	num1 := 10
	num2 := 5
	expected := 15
	//Act
	res, err := Add(num1, num2)
	//Assert
	assert.Equal(t, expected, res, "El numero resultado: %d, es distinto del esperado: %d", expected, res)
	assert.Nil(t, err)
	/*
		if res != expected {
			t.Errorf("El numero resultado: %d, es distinto del esperado: %d", expected, res)
		}
	*/
}
