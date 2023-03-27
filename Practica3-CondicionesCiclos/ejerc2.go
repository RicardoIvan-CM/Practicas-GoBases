package practica3

import "fmt"

func main() {
	edad := 24
	empleado := true
	antiguedad := 3
	sueldo := 150000
	if edad <= 22 {
		fmt.Println("Tu edad no permite realizarte un préstamo")
	} else if !empleado {
		fmt.Println("No se permite realizarte un préstamo si no tienes empleo")
	} else if antiguedad < 1 {
		fmt.Println("No tienes la suficiente antigüedad laboral para recibir un préstamo")
	} else {
		if sueldo > 100000 {
			fmt.Println("Felicidades, se te otorgará un prestamo y tu sueldo permite que no se cobre interés")
		} else {
			fmt.Println("Se te otorgará un préstamo pero se cobrará interés")
		}
	}
}
