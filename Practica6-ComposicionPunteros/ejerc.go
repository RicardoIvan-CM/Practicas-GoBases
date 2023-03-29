package main

import "fmt"

//Ejercicio 1

/*
type Alumno struct {
	Nombre,
	Apellido,
	DNI,
	Fecha string
}

func (a *Alumno) Detalle() {
	fmt.Println("Nombre:", a.Nombre)
	fmt.Println("Apellido:", a.Apellido)
	fmt.Println("DNI:", a.DNI)
	fmt.Println("Fecha:", a.Fecha)
}

func main() {
	a1 := Alumno{"Peter", "Parker", "abc123", "01/05/2003"}
	a1.Detalle()
}
*/

//Ejercicio2

type Producto interface {
	Precio() float64
}

type ProductoPequenio struct {
	Costo float64
}

func (p ProductoPequenio) Precio() float64 {
	return p.Costo
}

type ProductoMediano struct {
	Costo float64
}

func (p ProductoMediano) Precio() float64 {
	return p.Costo + p.Costo*0.03
}

type ProductoGrande struct {
	Costo float64
}

func (p ProductoGrande) Precio() float64 {
	return p.Costo + p.Costo*0.06 + 2500
}

func factory(tipo string, precio float64) Producto {
	switch tipo {
	case "Pequenio":
		return ProductoPequenio{precio}
	case "Mediano":
		return ProductoMediano{precio}
	case "Grande":
		return ProductoGrande{precio}
	}
	return nil
}

func main() {
	fmt.Println(factory("Grande", 5000).Precio())
}
