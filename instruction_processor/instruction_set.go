package instruction_processor

import (
	"fmt"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
	"strconv"
	"reflect"
	"strings"
)

/*
Coloca cualquier constante de cualquier tipo de dato permitido en la pila
constVal: Valor a insertar en la pila
stack: Pila en la que se va a puhs el const
*/
func loadConst[T any]( constVal T, stack *data_structures.Stack){
	fmt.Println("Executing LOAD_CONST")
	
	//Todo lo que le llega es siempre un string
	newConst := any(constVal).(string)
	//Convierte las cosas a string o a entero 
	if strings.Contains(newConst,`"`){
		fmt.Println("Metí un string")
		newConst = strings.ReplaceAll(newConst,`"`,"")
		stack.Push(newConst)
	}else if strings.Contains(newConst,"[") && strings.Contains(newConst,"]"){

		newConst = strings.ReplaceAll(newConst,"[", "")
		newConst = strings.ReplaceAll(newConst,"]", "")
		constSlice := strings.Split(newConst,",")

		//Es un array de strings
		if _,e := strconv.Atoi(constSlice[0]); e != nil {
			stack.Push(constSlice) //Agrega el slice a la pila
		} else{ //Es un slice de enteros
			

			constSliceInt := make([]int, len(constSlice))     //Crea un slice para los enteros
			for i := 0; i < len(constSlice); i++ {						//Convierte los strings a enteros
				constSliceInt[i],_= strconv.Atoi(constSlice[i])
			}

			stack.Push(constSliceInt)
		}

	}else{
		newConstInt,err := strconv.Atoi(newConst)

		if err != nil{
			fmt.Println(err)
			return
		}
		stack.Push(newConstInt)
	}
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
Retorna los operandos para la funciones de Binary... Internamente hace dos pops a la pila y los retorna
*/
func GetOperands(stack *data_structures.Stack, tipo string) (any,any){

	operand2,err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return nil,nil
	}

	operand1, err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return nil,nil
	}

	if tipo == "str"{
		

		if reflect.ValueOf(operand1).Kind() != reflect.String{
			operand1 = strconv.Itoa( any(operand1).(int) )
		}

		if reflect.ValueOf(operand2).Kind() != reflect.String{
			operand2 = strconv.Itoa( any(operand1).(int) )
		}

	}else{

		if reflect.ValueOf(operand1).Kind() != reflect.Int{
			operand1,_ = strconv.Atoi( any(operand1).(string) )
		}

		if reflect.ValueOf(operand2).Kind() != reflect.Int{
			operand2,_ = strconv.Atoi( any(operand2).(string) )
		}
	}

	return operand1, operand2

}

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
	fmt.Println("Executing COMPARE_OP")

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
				stack.Push(1)
			}else{
				stack.Push(0)
			}
		case ">":
			if op1 > op2 {
				stack.Push(1)
			}else{
				stack.Push(0)
			}
		case "==":
			if op1 == op2 {
				stack.Push(1)
			}else{
				stack.Push(0)
			}
	}

}

/*
Realiza una resta de los operandos (OP1 - OP2)
	stack: Pila del programa en dónde se ecuentran los operandos 

	carga el resultado de la resta la pila
*/
func binarySubstract(stack *data_structures.Stack){
	fmt.Println("Executing BINARY_SUBSTRACT")

	op1, op2 := GetOperands( stack, "int")

	if op1 == nil || op2 == nil {
		fmt.Println("ERROR GETTING OPERANDS")
		return 
	}

	result := any(op1).(int) - any(op2).(int)

	stack.Push(result)

}

/*
Realiza una suma de los operandos (OP1 + OP2)
	stack: Pila del programa en dónde se ecuentran los operandos 

	carga el resultado de la suma en la pila
*/
func binaryAdd(stack *data_structures.Stack){
	fmt.Println("Executing BINARY_ADD")

	op1, op2 := GetOperands( stack, "int")

	if op1 == nil || op2 == nil {
		fmt.Println("ERROR GETTING OPERANDS")
		return 
	}

	result := any(op1).(int) + any(op2).(int)

	stack.Push(result)
}

/*
Realiza una multiplicación de los operandos (OP1 * OP2)
	stack: Pila del programa en dónde se ecuentran los operandos 

	carga el resultado de la multiplicación en la pila
*/
func binaryMultiply(stack *data_structures.Stack){
	fmt.Println("Executing BINARY_MULTIPLY")

	op1, op2 := GetOperands( stack, "int")

	if op1 == nil || op2 == nil {
		fmt.Println("ERROR GETTING OPERANDS")
		return 
	}

	result := any(op1).(int) * any(op2).(int)

	stack.Push(result)
}

/*
Realiza una división entera de los operandos (OP1 / OP2)
	stack: Pila del programa en dónde se ecuentran los operandos 

	carga el resultado de la división en la pila
*/
func binaryDivide(stack *data_structures.Stack){
	fmt.Println("Executing BINARY_DIVIDE")

	op1, op2 := GetOperands( stack, "int")

	if op1 == nil || op2 == nil {
		fmt.Println("ERROR GETTING OPERANDS")
		return 
	}

	result := any(op1).(int) / any(op2).(int)

	stack.Push(result)
}

/*
Realiza una AND entre los operandos (OP1 AND OP2)
	stack: Pila del programa en dónde se ecuentran los operandos 

	carga el resultado del AND en la pila
*/
func binaryAnd(stack *data_structures.Stack){
	fmt.Println("Executing BINARY_AND")

	op1, op2 := GetOperands( stack, "int")

	if op1 == nil || op2 == nil {
		fmt.Println("ERROR GETTING OPERANDS")
		return 
	}

	newbool1 := op1 != 0
	newbool2 := op2 != 0

	if newbool1 && newbool2{
		stack.Push(1)
	}else{
		stack.Push(0)
	}
}

/*
Realiza un OR entre los operandos (OP1 OR OP2)
	stack: Pila del programa en dónde se ecuentran los operandos 

	carga el resultado del OR en la pila
*/
func binaryOr(stack *data_structures.Stack){
	fmt.Println("Executing BINARY_OR")

	op1, op2 := GetOperands( stack, "int")

	if op1 == nil || op2 == nil {
		fmt.Println("ERROR GETTING OPERANDS")
		return 
	}

	newbool1 := op1 != 0
	newbool2 := op2 != 0

	if newbool1 || newbool2{
		stack.Push(1)
	}else{
		stack.Push(0)
	}
}


/*
Realiza el cálculo del cociente de la division entre 2 operandos (op1 % op2)
	
	stack: pila del pograma
	retorna el resultado en la pila

*/
func binaryModulo(stack *data_structures.Stack){
	fmt.Println("Executing BINARY_MODULO")

	op1, op2 := GetOperands( stack, "int")

	if op1 == nil || op2 == nil {
		fmt.Println("ERROR GETTING OPERANDS")
		return 
	}

	result := any(op1).(int) % any(op2).(int)

	stack.Push(result)

}

/*
Realiza la operación: array[index] = value 
 En la pila deben haber los siguientes elementos: [index, array, value] <- Tope de la pila 

 Programa de prueba:
 En el almacen en "y" ya debemos tener guardada una lista

 	LOAD_CONST 0     - Indice
	LOAD_GLOBAL y		 - Referencia al array
	LOAD_CONST 90    - Valor que queremos guardar
	STORE_SUBSCR


*/
func storeSubscr(stack *data_structures.Stack, storage *data_structures.MapTable){
	fmt.Println("Executing STORE_SUBSCR")

	value,err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return 
	}

	arrayRef,err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return 
	}

	index,err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return 
	}

	array,err := storage.Get(any(arrayRef).(string))

	if err != nil {
		fmt.Println(err)
		return
	}

	//Si el tipo del value eas entero, entonces asume que el array es un []int

	if reflect.ValueOf(value).Kind() == reflect.Int{
		
		any(array).([]int)[any(index).(int)] = any(value).(int)

	}else{
		any(array).([]string)[any(index).(int)] = any(value).(string)
	}

	storage.Update(any(arrayRef).(string), array)
}


/*
Carga en el tope de la pila el elemento de un arreglo en el indice indicado 
	stack: pila del programa
	storage: almacen de datos 
*/

func binarySubscr (stack *data_structures.Stack, storage *data_structures.MapTable){
	
	//Saca la referencia del array del Stack 
	arrayRef,err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return 
	}

	//Carga el array del almacen
	array,err := storage.Get(any(arrayRef).(string))
	if err != nil {
		fmt.Println(err)
		return
	}

	//Saca de la pila en indice 
	index,err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return 
	}

	//Saca el valor del array y lo pone al tope de la pila 
	stack.Push( any(array).([]any)[any(index).(int)] )
}


/*
Salta a la linea de código indicada por target
	target: Target es un indice indicativo en el archivo que debe coincidir con el numero de instruccion una vez almacenado en el codigo de memoria. Osea el numero al que va a saltar
	pc: Program counter, carga la proxima instruccion en memoria 
*/
func jumpAbsolute(target string, pc *int){
	
	newTarget,_ := strconv.Atoi(target)
	newTarget -= 1   //Ya que la instrucción suma 1 cuando termina el decode/execute

	//Actualiza el PC para que la siguiente instruccion a ejecutar sea el target 
	*pc = newTarget
}

/*
Si el tope de la pila es True salta al target
	target: Target es un indice indicativo en el archivo que debe coincidir con el numero de instruccion una vez almacenado en el codigo de memoria. Osea el numero al que va a saltar
	pc: Program counter, carga la proxima instruccion en memoria 
*/

func jumpTrue(target string, stack *data_structures.Stack, pc *int){
	newTarget,_ := strconv.Atoi(target)
	newTarget -= 1   //Ya que la instrucción suma 1 cuando termina el decode/execute

	//Saca la referencia del array del Stack 
	value,err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return 
	}

	newValue,_ :=  strconv.Atoi(any(value).(string))

	//1 = true
	//0 = false
	if newValue == 1 {
		*pc = newTarget
	}


}

/*
Si el tope de la pila es False salta al target
	target: Target es un indice indicativo en el archivo que debe coincidir con el numero de instruccion una vez almacenado en el codigo de memoria. Osea el numero al que va a saltar
	pc: Program counter, carga la proxima instruccion en memoria 
*/

func jumpFalse(target string, stack *data_structures.Stack, pc *int){
	newTarget,_ := strconv.Atoi(target)
	newTarget -= 1   //Ya que la instrucción suma 1 cuando termina el decode/execute

	//Saca la referencia del array del Stack 
	value,err := stack.Pop()
	if err != nil{
		fmt.Println(err)
		return 
	}

	newValue,_ :=  strconv.Atoi(any(value).(string))

	//1 = true
	//0 = false
	if newValue == 0 {
		*pc = newTarget
	}
}


/*
Construye una lista con "elements" cantidad de elementos 
	elements: cantidad de elementos en la lista, que están en la pila
	stack: pila del programa, en dónde se encuentran los elementos de la lista 
	
	Deja la lista en la pila
*/
func buildList(elements string, stack *data_structures.Stack){
	//Saca la cantidad de elementos que tiene que sacar de la lista
	listElements,_ := strconv.Atoi(elements)
	
	list := make([]any, listElements)


	//El último elemento está al tope de la pila
	// por eso va al "revez"
	for i := listElements-1; i >= 0; i-- {

		element,err := stack.Pop()
		if err != nil{
			fmt.Println(err)
			return 
		}
		list[i] = element
	}

	stack.Push(list)
}

func end(){

	fmt.Println("Última instrucción del programa")

}