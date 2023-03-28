package main

import (
	"fmt"
)

/*
//Ejercicio 1

var Products = []Product{
	Product{1, "Laptop", 15000, "Computadora portátil", "Electrónicos"},
	Product{2, "Naranja", 20, "Fruta cítrica", "Frutas"},
}

func main() {
	p1 := Product{3, "Chocolate", 43, "Dulce de cacao", "Golosinas"}
	p1.GetAll()
	res, err := getById(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	p1.Save()
	p1.GetAll()
	res, err = getById(3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p Product) GetAll() {
	fmt.Println(Products)
}

func getById(id int) (Product, error) {
	for _, product := range Products {
		if id == product.ID {
			return product, nil
		}
	}
	return Product{}, errors.New("No se encontró al producto")
}
*/

//Ejercicio 2

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Person
	EmployeeID int
	Position   string
}

func (e *Employee) PrintEmployee() {
	fmt.Println(*e)
}

func main() {
	p1 := Person{1, "Peter Pan", "05/19/2002"}
	e1 := Employee{p1, 2, "Manager"}
	e1.PrintEmployee()

	fmt.Println(e1.Person.ID)
}
