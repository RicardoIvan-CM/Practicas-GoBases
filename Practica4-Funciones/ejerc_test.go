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
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("No hay valores a promediar", func(t *testing.T) {
		//Arrange
		var valores = []int{}
		var resultadoEsperado float64 = 0
		var errorEsperado string = "No hay valores a promediar"

		//Act
		resultadoObtenido, err := Promedio(valores...)

		//Assert
		assert.EqualError(t, err, errorEsperado, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("Hay un valor negativo", func(t *testing.T) {
		//Arrange
		var valores = []int{
			8, 7, -5, 10,
		}
		var resultadoEsperado float64 = 0
		var errorEsperado string = "No se aceptan valores negativos"

		//Act
		resultadoObtenido, err := Promedio(valores...)

		//Assert
		assert.EqualError(t, err, errorEsperado, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})
}

func TestSalario(t *testing.T) {
	t.Run("C", func(t *testing.T) {
		//Arrange
		var minutosMes int = 300
		var categoria = "C"
		var resultadoEsperado float64 = 5000

		//Act
		resultadoObtenido, err := Salario(minutosMes, categoria)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("B", func(t *testing.T) {
		//Arrange
		var minutosMes int = 300
		var categoria = "B"
		var resultadoEsperado float64 = 9000

		//Act
		resultadoObtenido, err := Salario(minutosMes, categoria)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("A", func(t *testing.T) {
		//Arrange
		var minutosMes int = 300
		var categoria = "A"

		var resultadoEsperado float64 = 22500
		//Act
		resultadoObtenido, err := Salario(minutosMes, categoria)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("Categoría no existente", func(t *testing.T) {
		//Arrange
		var minutosMes int = 300
		var categoria = "Z"
		var errorEsperado = "No se encontró la categoría"
		var resultadoEsperado float64 = 0

		//Act
		resultadoObtenido, err := Salario(minutosMes, categoria)

		//Assert
		assert.EqualError(t, err, errorEsperado, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})
}

func TestEstadisticas(t *testing.T) {
	t.Run("Mínimo", func(t *testing.T) {
		//Arrange
		var operacion string = "minimum"
		var resultadoEsperado float64 = 1

		//Act
		resultadoObtenido, err := Estadisticas(operacion)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("Máximo", func(t *testing.T) {
		//Arrange
		var operacion string = "maximum"
		var resultadoEsperado float64 = 5

		//Act
		resultadoObtenido, err := Estadisticas(operacion)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("Promedio", func(t *testing.T) {
		//Arrange
		var operacion string = "average"
		var resultadoEsperado float64 = 3

		//Act
		resultadoObtenido, err := Estadisticas(operacion)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("Otra operación", func(t *testing.T) {
		//Arrange
		var operacion string = "otra"
		var errorEsperado string = "No se encontró la operación"
		var resultadoEsperado float64 = 0

		//Act
		resultadoObtenido, err := Estadisticas(operacion)

		//Assert
		assert.EqualError(t, err, errorEsperado, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})
}

func TestAnimal(t *testing.T) {
	t.Run("Perro", func(t *testing.T) {
		//Arrange
		var tipo string = "dog"
		var cantidad int = 3
		var resultadoEsperado int = 30000

		//Act
		funcAnimal, _ := Animal(tipo)
		resultadoObtenido, err := funcAnimal(cantidad)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("Gato", func(t *testing.T) {
		//Arrange
		var tipo string = "cat"
		var cantidad int = 3
		var resultadoEsperado int = 15000

		//Act
		funcAnimal, _ := Animal(tipo)
		resultadoObtenido, err := funcAnimal(cantidad)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("hamster", func(t *testing.T) {
		//Arrange
		var tipo string = "hamster"
		var cantidad int = 3
		var resultadoEsperado int = 750

		//Act
		funcAnimal, _ := Animal(tipo)
		resultadoObtenido, err := funcAnimal(cantidad)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("tarantula", func(t *testing.T) {
		//Arrange
		var tipo string = "tarantula"
		var cantidad int = 3
		var resultadoEsperado int = 450

		//Act
		funcAnimal, _ := Animal(tipo)
		resultadoObtenido, err := funcAnimal(cantidad)

		//Assert
		assert.NoError(t, err, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})

	t.Run("Otra operación", func(t *testing.T) {
		//Arrange
		var tipo string = "pez"
		var errorEsperado string = "No se encontró el animal solicitado"

		//Act
		_, err := Animal(tipo)

		//Assert
		assert.EqualError(t, err, errorEsperado, "Los errores no coinciden")
	})

	t.Run("Cantidad negativa", func(t *testing.T) {
		//Arrange
		var tipo string = "cat"
		var cantidad int = -8
		var resultadoEsperado int = 0
		var errorEsperado string = "La cantidad no debe ser menor a 0"

		//Act
		funcAnimal, _ := Animal(tipo)
		resultadoObtenido, err := funcAnimal(cantidad)

		//Assert
		assert.EqualError(t, err, errorEsperado, "Los errores no coinciden")
		assert.Equal(t, resultadoEsperado, resultadoObtenido, "deben ser iguales")
	})
}
