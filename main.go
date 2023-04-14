package main

import (
	"fmt"
	"os"

	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/file_management"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/instruction_processor"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/config"
	//"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
)

//Declaracion de Variables
var pc int 						//Program Counter: Guarda el numero de instruccion a ejecutar
var ir string 				//Instruction Register: Guarda la instrucción que se debe decodificar y ejecutar
var program []string 	//Memoria del programa
var stack data_structures.Stack
var storage (data_structures.MapTable)
var debugMode bool = config.DebugMode

/**
	Carga el programa en memoria 
	filame: nombre del archivo a cargar, el archivo debe estar en la carpeta ./test_code
**/
func loadProgram(filePath string){
	var filePtr *os.File = file_management.OpenFile(filePath)
	program = file_management.ReadFile(filePtr)
	file_management.CloseFile(filePtr)
}

func testStack(){

	stack.Push(1)
	stack.Push("hola")

	if val,e:=stack.Pop();e!= nil{
		fmt.Println(e)
	}else{
		fmt.Println(val)
	}
	if val,e:=stack.Pop();e!= nil{
		fmt.Println(e)
	}else{
		fmt.Println(val)
	}
	if val,e:=stack.Pop();e!= nil{
		fmt.Println(e)
	}else{
		fmt.Println(val)
	}
	fmt.Println(stack)
}

func testMapTable(){
	storage = make(data_structures.MapTable)

	storage.Insert("hola", 1)
	storage.Insert("adios", 2)
	storage.Insert("queso", "queso")

	fmt.Println(storage)
	storage.Delete("adios")
	fmt.Println(storage.Get("adios"))
	fmt.Println(storage)

}


//Ciclo de ejecución del programa 
func main(){

	
	//Inicializacion de variables
	pc = 0
	ir = ""
	storage = make(data_structures.MapTable)


	//Carga el programa el memoria
	loadProgram("test_programs/Programa_Prueba.txt")

	//Simulación del ciclo de fetch-decodificación-ejecución
	for{
		
		if pc >= len(program) { 
			fmt.Println("End of Execution")
			break	
		}

		ir = program[pc] 																//Carga la instrucción al registro de instrucción
		instruction_processor.DecodeExecute(ir, &pc, &stack, &storage)		//Decodifica y ejecuta la instrucción
		pc += 1																				  //Incrementa el PC para cargar la siguiente instrucción
		if debugMode{ fmt.Println()}
	}

	
	fmt.Println("\nStack y Almacen")
	fmt.Println(stack)
	fmt.Println(storage)

}



