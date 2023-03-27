package practica3

import "fmt"

func main() {
	cuenta := 0
	palabra := "Hola"
	for cuenta < len(palabra) {
		cuenta++
	}
	fmt.Println(cuenta)
	for _, letra := range palabra {
		fmt.Println(string(letra))
	}
}
