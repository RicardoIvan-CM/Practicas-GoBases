package main

// Ejercicio 2
/*
type MiError struct {
	msg string
}

func (e MiError) Error() string {
	return e.msg
}

func main() {
	var salary int
	var errorEsperado MiError = MiError{"Error: el salario ingresado no alcanza el mínimo imponible"}
	var errorActual MiError
	if salary <= 10000 {
		errorActual = MiError{"Error: el salario ingresado no alcanza el mínimo imponible"}
	}
	if errors.Is(errorActual, errorEsperado) {
		fmt.Println(errorActual.Error())
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}*/
