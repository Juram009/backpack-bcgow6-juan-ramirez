package ejercicio2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	//Arrange
	list := []int{4, 2, 7, 4, 6, 3, 6, 89, 6, 3, 7}
	ordererList := []int{2, 3, 3, 4, 4, 6, 6, 6, 7, 7, 89}
	//Act
	res := burbuja(list)
	//Assert
	assert.Equal(t, ordererList, res)
}
