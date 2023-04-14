package instruction_processor

import (
	"fmt"

	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
)

/*
Función que se encarga de imprimir una cantidad de elementos del stack

	params: parámetros de la función (En este caso es sólo la cantidad de elementos que va a imprimir)
	stack: la pila del programa
*/
func print(params []any, stack *data_structures.Stack) {

	fmt.Println("\n*****")
	fmt.Println("Funcion Interna de print")
	fmt.Print("Print:\n")

	j := len(params)

	for i := 0; i < j; i++ {
		fmt.Println(params[i])
	}

	//fmt.Print( *stack )
	fmt.Println("*****")

}
