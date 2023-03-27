package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(impuesto(44280))
}

func impuesto(salario float64) float64 {
	var impuesto float64
	if salario > 50000 {
		impuesto += salario * 0.17
		if salario > 150000 {
			impuesto += salario * 0.10
		}
	}
	return impuesto
}

func promedio(valores ...int) float64 {
	var suma float64 = 0
	var cuenta int = 0
	for _, valor := range valores {
		suma += float64(valor)
		cuenta++
	}
	return suma / float64(cuenta)
}

func salario(minutosMes int, categoria string) float64 {
	var resultado float64
	horas := float64(minutosMes) / 60
	switch categoria {
	case "C":
		resultado = horas * 1000
	case "B":
		salarioMes := horas * 1500
		resultado = salarioMes + salarioMes*0.20
	case "A":
		salarioMes := horas * 3000
		resultado = salarioMes + salarioMes*0.50
	default:
		errors.New("No se encontró la categoría")
	}
	return resultado
}

func minimo(valores ...int) float64 {
	var min int
	for _, valor := range valores {
		if min < valor {
			min = valor
		}
	}
	return float64(min)
}

func maximo(valores ...int) float64 {
	var max int
	for _, valor := range valores {
		if max > valor {
			max = valor
		}
	}
	return float64(max)
}

/*
func estadisticas(operacion string) (operacionFunc func(...int) float64,err string){
	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	switch operacion {
	case minimum:
		operacionFunc = minimo
	case average:
		operacionFunc = promedio
	case maximum:
		operacionFunc = maximo
	default:
		return 0
	}

	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)


}*/
