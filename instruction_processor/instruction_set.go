package instruction_processor

import (
	"fmt"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
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
		
	}else{
		//Guarda el NUEVO valor en el almacen
		e := storage.Insert(varName, val)

		if e!= nil{
			//Si ya existe entonces lo actualiza 
			storage.Update(varName, val)
		}
	}

}