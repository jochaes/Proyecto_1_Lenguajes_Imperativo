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
func print(params []any, stack *data_structures.Stack){

	fmt.Println("*****")
	fmt.Println("Funcion Interna de print")
	fmt.Println("Printing stack: ")


	// j := any(params[0]).(int)

	// for i := 0; i < j ; i++ {
	// 	fmt.Print( stack[0] )
	// }

	fmt.Print( *stack )

	fmt.Println("\n*****")

}