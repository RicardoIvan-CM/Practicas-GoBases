package main

import (
	"errors"
	"fmt"
)

func main() {
	alimentoFunc, err := animal("cat")
	if err != nil {
		fmt.Println(err)
	} else {
		alimento, err := alimentoFunc(10)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("La cantidad de alimento es", alimento)
		}
	}
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

func promedio(valores ...int) (float64, error) {
	if len(valores) < 1 {
		return 0, errors.New("No hay valores a comparar")
	}
	var suma float64 = 0
	var cuenta int = 0
	for _, valor := range valores {
		if valor < 0 {
			return 0, errors.New("No se aceptan valores negativos")
		}
		suma += float64(valor)
		cuenta++
	}
	return suma / float64(cuenta), nil
}

func salario(minutosMes int, categoria string) (resultado float64, err error) {
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
		err = errors.New("No se encontró la categoría")
	}
	return
}

func minimo(valores ...int) (float64, error) {
	if len(valores) < 1 {
		return 0, errors.New("No hay valores a comparar")
	}
	var min int = valores[0]
	for _, valor := range valores {
		if valor < 0 {
			return 0, errors.New("No se aceptan valores negativos")
		}
		if valor < min {
			min = valor
		}
	}
	return float64(min), nil
}

func maximo(valores ...int) (float64, error) {
	if len(valores) < 1 {
		return 0, errors.New("No hay valores a comparar")
	}
	var max int
	for _, valor := range valores {
		if valor < 0 {
			return 0, errors.New("No se aceptan valores negativos")
		}
		if valor > max {
			max = valor
		}
	}
	return float64(max), nil
}

func estadisticas(operacion string) (resultado float64, err error) {
	var operacionFunc func(...int) (float64, error)

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
		err = errors.New("No se encontró la operación")
	}

	if err == nil {
		resultado, err = operacionFunc(2, 3, 3, 4, 1, 2, 4, 5)
	}
	return resultado, err
}

func alimentoPerro(cantidad int) (int, error) {
	if cantidad < 0 {
		return 0, errors.New("La cantidad debe ser menor a 0")
	}
	return 10000 * cantidad, nil
}

func alimentoGato(cantidad int) (int, error) {
	if cantidad < 0 {
		return 0, errors.New("La cantidad no debe ser menor a 0")
	}
	return 5000 * cantidad, nil
}

func alimentoHamster(cantidad int) (int, error) {
	if cantidad < 0 {
		return 0, errors.New("La cantidad no debe ser menor a 0")
	}
	return 250 * cantidad, nil
}

func alimentoTarantula(cantidad int) (int, error) {
	if cantidad < 0 {
		return 0, errors.New("La cantidad no debe ser menor a 0")
	}
	return 150 * cantidad, nil
}

func animal(tipo string) (alimentoFunc func(int) (int, error), err error) {
	const (
		perro     = "dog"
		gato      = "cat"
		hamster   = "hamster"
		tarantula = "tarantula"
	)

	switch tipo {
	case perro:
		alimentoFunc = alimentoPerro
	case gato:
		alimentoFunc = alimentoGato
	case hamster:
		alimentoFunc = alimentoHamster
	case tarantula:
		alimentoFunc = alimentoTarantula
	default:
		return nil, errors.New("No se encontró el animal solicitado")
	}

	return alimentoFunc, nil
}
