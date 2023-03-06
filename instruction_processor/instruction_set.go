package instruction_processor

import "fmt"


/**
	Coloca el valor del contenido de la variable en el tope de la pila
	varname: nombre de la variable
	stack: pila a colocar el contenido de la variable
	pc: program counter 
**/

func loadFast(varName string, stack *[]int, pc *int){
	fmt.Println("Executing LOAD_FAST")
	fmt.Println("varname:",varName )	
}