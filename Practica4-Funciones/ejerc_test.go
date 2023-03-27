package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpuesto(t *testing.T) {

	t.Run("Salario por debajo de 50000", func(t *testing.T) {
		//Arrange
		var salario float64 = 10000
		var resultadoEsperado float64 = 0

		//Act
		resultadoObtenido := Impuesto(salario)

		//Assert
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("Salario por encima de 50000", func(t *testing.T) {
		//Arrange
		var salario float64 = 55000
		var resultadoEsperado float64 = salario * 0.17

		//Act
		resultadoObtenido := Impuesto(salario)

		//Assert
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("Salario por encima de 150000", func(t *testing.T) {
		//Arrange
		var salario float64 = 200000
		var resultadoEsperado float64 = salario * 0.27

		//Act
		resultadoObtenido := Impuesto(salario)

		//Assert
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})
}

func TestPromedio(t *testing.T) {
	t.Run("Promedio Correcto", func(t *testing.T) {
		//Arrange
		var valores = []int{
			5, 7, 8, 7, 6, 9, 10, 10, 7,
		}
		var resultadoEsperado float64 = float64(5+7+8+7+6+9+10+10+7) / float64(9)

		//Act
		resultadoObtenido, err := Promedio(valores...)

		//Assert
		assert.Nil(t, err, "Se esperaba un valor nulo")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("No hay valores a promediar", func(t *testing.T) {
		//Arrange
		var valores = []int{}
		var resultadoEsperado float64 = 0
		var errorEsperado string = "No hay valores a comparar"

		//Act
		resultadoObtenido, err := Promedio(valores...)

		//Assert
		assert.Equal(t, errorEsperado, err.Error(), "El error es incorrecto")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})
}
