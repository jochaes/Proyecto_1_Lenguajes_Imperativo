package instruction_processor

import (
	"fmt"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
	"strconv"
)

/*
Coloca cualquier constante de cualquier tipo de dato permitido en la pila
constVal: Valor a insertar en la pila
stack: Pila en la que se va a puhs el const
*/
func loadConst[T any]( constVal T, stack *data_structures.Stack){
	fmt.Println("Executing LOAD_CONST")
	stack.Push(constVal)
}

/*
Coloca el valor del contenido de la variable en el tope de la pila
	varname: nombre de la variable
	stack: pila a colocar el contenido de la variable
	pc: program counter 
*/
func loadFast(varName string, stack *data_structures.Stack, storage *data_structures.MapTable){
	fmt.Println("Executing LOAD_FAST")
	
	val, e := storage.Get(varName)

	if e!= nil{
		fmt.Println(e)
		return
	}else{
		stack.Push(val)
	}

}

/*
Coloca el valor del contenido de la variable en el tope de la pila
	varname: nombre de la variable
	stack: pila a colocar el contenido de la variable
	pc: program counter 
*/
func storeFast(varName string, stack *data_structures.Stack, storage *data_structures.MapTable){
	fmt.Println("Executing LOAD_FAST")
	
	val, e := stack.Pop()

	if e!= nil{
		fmt.Println(e)
		return
		
	}else{
		//Guarda el NUEVO valor en el almacen
		e := storage.Insert(varName, val)

		if e!= nil{
			//Si ya existe entonces lo actualiza 
			storage.Update(varName, val)
		}
	}

}



/************************************ Load global y Call Function ************************************


Estas 2 funciones están programadas con el fin de que siempre se llame a la misma función: print()
print tiene como parámetro la cantidad de elementos que va a imprimir del stack 

Para ejecutar esta función se utiliza el siguiente bytecode 
LOAD_GLOBAL print    ->Carga la función en el stack    									                        push(print)
LOAD_CONST 1         -> Carga el parametro de print en la pila 					                        push(1)
CALL_FUNCTION 1      -> Indica que va a tomar 1 elemento de la pila, y luego llama a la función pop():1 pop():print
										 -> Ejecuta la función: Imprime 1 elemento de la pila  
*/


/*
Carga en el stack el valor de una variable global o la referencia de una función 
	En este caso se va a utilizar únicamente para cargar la función interna que imprime la pila,
	por lo tanto el programa no tendrá variables globales, ni funciones. 
	name: nombre de la función o variable global 
	stack: pila del programa en dónde se cargará la referencia de la función. 
*/
func loadGlobal[T any]( name T, stack *data_structures.Stack){
	fmt.Println("Executing LOAD_GLOBAL")
	stack.Push(name)
}


/*
Llama a la función con sus parámetros
	En este caso, solo ejecutará la función interna de que imprime el stack, que recibe como parámetro 
	 la cantidad de elementos que va a imprimir. 

	params: parametro o parametros de la función que se va a llamar
	stack: pila del programa
*/
func callFunction[T any]( params T, stack *data_structures.Stack ){
	fmt.Println("Executing CALL_FUNCTION")
	

	//Cómo ya sabemos que solo se va a ejecutar la función interna de printStack entonces el parámetro
	// es la cantidad de elementos del stack que va a imprimir, entonces convierte params a entero
	cantidadParams,err := strconv.Atoi( any(params).(string) )

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	parameters := make([]any, cantidadParams)

	for i := 0; i < cantidadParams; i++ {
		val, e := stack.Pop()

		if e!= nil{
			fmt.Println(e)
			return
		}else{
			//Añade los paramatros de la función en un slice 
			parameters = append(parameters, val)
		}
	}

	//En este punto ya todos los parametros están en el slice, y en el top de la pila está el nombre de la función
	function, e := stack.Pop()
	if e!= nil{
		fmt.Println(e)
		return
	}else{

		if function == "print"{

			print(parameters, stack)

		}
	}
}

//************************************ ************************************ ************************************

/*
Realiza una comparación booleana según el op que reciba 
	Los operandos deben ser números 
	op: <, >, ==
	stack: pila de dónde va a sacar los dos operandos
	
	retorna el resultado en la pila. 
	1 = TRUE
	0 = FALSE
*/

func compareOp[T any]( op T, stack *data_structures.Stack){

	operator := any(op).(string)
	
	operand2,err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return
	}

	operand1, err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return
	}

	op1,_ := strconv.Atoi( any(operand1).(string) )
	op2,_ := strconv.Atoi( any(operand2).(string) )


	switch operator {
		case "<":
			if op1 < op2 {
				stack.Push("TRUE")
			}else{
				stack.Push("FALSE")
			}
		case ">":
			if op1 > op2 {
				stack.Push("TRUE")
			}else{
				stack.Push("FALSE")
			}
		case "==":
			if op1 == op2 {
				stack.Push("TRUE")
			}else{
				stack.Push("FALSE")
			}
	}

}
