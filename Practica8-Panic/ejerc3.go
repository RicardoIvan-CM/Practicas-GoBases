package main

import (
	"errors"
	"fmt"
)

//Ejercicio3

type Cliente struct {
	Legajo,
	Nombre,
	DNI,
	Telefono,
	Domicilio string
}

var (
	ErrDatos = errors.New("Error: uno o más datos son inválidos")
)

func AgregarCliente(arr *[]Cliente, agregado *Cliente) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for _, cliente := range *arr {
		if cliente == *agregado {
			panic("Error: el cliente ya existe")
		}
	}

	dato, err := ValidarCliente(arr, agregado)
	if err != nil {
		panic("Error: El dato " + dato + " es inválido")
	}

	*arr = append(*arr, *agregado)
}

func ValidarCliente(arr *[]Cliente, cliente *Cliente) (string, error) {
	err := ErrDatos
	if cliente.Legajo == "" {
		return "legajo", err
	}
	if cliente.Nombre == "" {
		return "nombre", err
	}
	if cliente.DNI == "" {
		return "dni", err
	}
	if cliente.Domicilio == "" {
		return "domicilio", err
	}
	if cliente.Telefono == "" {
		return "telefono", err
	}
	return "", nil
}

func main() {
	c1 := Cliente{"Lorem ipsum", "Peter Pan", "abc123", "5511223344", "Calle Falsa 123"}
	c2 := Cliente{"Ipsum lorem", "John Doe", "123abc", "1122334455", "Falsa Calle abc"}

	var clientes = &[]Cliente{
		c1,
	}

	AgregarCliente(clientes, &c2)

	fmt.Println(*clientes)
}
