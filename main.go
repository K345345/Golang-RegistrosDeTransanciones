package main

import (
	"fmt"
)

func main() {
	var gastos []float64
	var ingresos []float64

	fmt.Println("------------------------")
	fmt.Println("Registros de transacciones")
	fmt.Println("------------------------")

	// Generamos un bucle infinito,  para mantener el programa encendido
principalFor:
	for {
		var opcion string

		fmt.Println("|Ingrese gastos(g), Ingrese ingreso(i), Calcular saldo total(st), Salir(s)")
		fmt.Println("|Ingresa una opcion de las anteriores.|")
		fmt.Scanf("%s", &opcion)

		// si se cumple la condicion paramos el programa
		if opcion == "s" {
			fmt.Println("Saliendo del programa...")
			break
		}

		switch opcion {
		//Ingresamos los gastos generados
		case "g":
			var dato float64
			fmt.Println("Ingrese el registro de gastos: ")
			fmt.Scanf("%f", &dato)
			gastos = append(gastos, dato)
			break

		//Ingresar ingresos generados
		case "i":
			var dato float64
			fmt.Println("Ingrese el registro de ingreso: ")
			fmt.Scanf("%f", &dato)
			ingresos = append(ingresos, dato)
			break

		case "st":
			var a string
			fmt.Println("¿Qué registros desea consultar: Gastos(g), Ingreso(i), Ganancias(G)?")
			fmt.Scanf("%s", &a)

			//Consultar gastos o ingresos
			switch a {

			//Mostramos los registros de los gastos
			case "g":
				for i := 0; i < len(gastos); i++ {
					fmt.Printf("%d - $%.2f\n", i, gastos[i])
				}

			//Mostramos los registros de los ingresos
			case "i":
				for i := 0; i < len(ingresos); i++ {
					fmt.Printf("%d - $%.2f\n", i, ingresos[i])
				}

			case "G":

				var totalIngreso, totalGasto float64 = 0, 0

				// iteramos en el slice ingreso, para obtener el valor de cada posicion
				for i := 0; i < len(ingresos); i++ {
					// Sumamos el totalIngreso con el valor iteraro, gurdamos en la variable totalIngreso
					totalIngreso += ingresos[i]
				}

				fmt.Printf("El total de ingreso es de: %f\n", totalIngreso)

				for i := 0; i < len(gastos); i++ {
					totalGasto += gastos[i]
				}

				fmt.Printf("El total de gastos es de: %f\n", totalGasto)

				// Se calcula si hubo ganacias o perdidas
				var Ganancias float64 = totalIngreso - totalGasto

				//Comprobamos si los ingresos superan los gastos
				if totalIngreso > totalGasto {

					fmt.Println("Los ingresos son mayores que los gastos.")
					fmt.Printf("Los ingresos totales son: $%f\n", Ganancias)
				} else if totalIngreso < totalGasto {

					fmt.Println("Los ingresos son menores que los ingresos.")
					fmt.Printf("Las perdidas totales son: $%f\n", Ganancias)
				}
			case "s":
				break principalFor

			default:
				fmt.Println("ERROR: Ingrese un valor valido")
				break
			}

			break

		default:
			fmt.Println("ERROR: Ingrese un valor valido")
			break

		}

	}
}
