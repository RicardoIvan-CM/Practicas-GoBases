package main

import (
	"errors"
	"fmt"
)

// Ejercicio 5

func SalarioMensual(horas int, valorHora float64) (float64, error) {
	if horas < 80 {
		return 0, errors.New("Error: el trabajador no puede haber trabajado menos de 80 horas mensuales")
	}

	salario := float64(horas) * valorHora
	if salario >= 150000 {
		salario -= salario * 0.10
	}
	return salario, nil
}

func main() {
	salario, err := SalarioMensual(100, 1000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El salario es de:", salario)
	}
}
