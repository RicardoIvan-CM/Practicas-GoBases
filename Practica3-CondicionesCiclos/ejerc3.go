package practica3

import "fmt"

func main() {
	meses := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiempre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}
	fmt.Println(meses)

	mes := 14
	if meses[mes] == "" {
		fmt.Println("No existe ese mes")
	} else {
		fmt.Print(mes)
		fmt.Println(",", meses[mes])
	}

	/*
		switch mes {
		case 1:
			fmt.Println("Enero")
		case 2:
			fmt.Println("Febrero")
		case 3:
			fmt.Println("Marzo")
		case 4:
			fmt.Println("Abril")
		case 5:
			fmt.Println("Mayo")
		case 6:
			fmt.Println("Junio")
		case 7:
			fmt.Println("Julio")
		case 8:
			fmt.Println("Agosto")
		case 9:
			fmt.Println("Septiembre")
		case 10:
			fmt.Println("Octubre")
		case 11:
			fmt.Println("Noviembre")
		case 12:
			fmt.Println("Diciembre")
		default:
			fmt.Println("No existe el mes")
		}*/
}
