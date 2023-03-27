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
