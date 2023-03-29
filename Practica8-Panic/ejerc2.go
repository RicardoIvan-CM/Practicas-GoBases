package main

/*
import (
	"fmt"
	"io"
	"os"
)

//Ejercicio2

func main() {
	file, err := os.Open("./Practica8-Panic/customers.txt")

	defer func() {
		file.Close()
		fmt.Println("ejecución finalizada")
	}()

	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	text, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(text))
}
*/
