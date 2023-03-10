package main

import (
	"fmt"
	"os"

	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/file_management"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/instruction_processor"
	//"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
)

//Declaracion de Variables
var pc int 						//Program Counter: Guarda el numero de instruccion a ejecutar
var ir string 				//Instruction Register: Guarda la instrucción que se debe decodificar y ejecutar
var program []string 	//Memoria del programa
var stack data_structures.Stack
var storage (data_structures.MapTable)

/**
	Carga el programa en memoria 
	filame: nombre del archivo a cargar, el archivo debe estar en la carpeta ./test_code
**/
func loadProgram(fileName string){
	var filePtr *os.File = file_management.OpenFile("test_code/"+fileName)
	program = file_management.ReadFile(filePtr)
	file_management.CloseFile(filePtr)
}

func testStack(){

	stack.Push(1)
	stack.Push("hola")

	if e:=stack.Pop();e!= nil{
		fmt.Println(e)
	}
	if e:=stack.Pop();e!= nil{
		fmt.Println(e)
	}
	if e:=stack.Pop();e!= nil{
		fmt.Println(e)
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

	//Carga el programa el memoria
	loadProgram("prueba_1.txt")

	//Simulación del ciclo de fetch-decodificación-ejecución
	for{
		if pc >= len(program) { 
			fmt.Println("End of Execution")
			break	
		}

		ir = program[pc] 																//Carga la instrucción al registro de instrucción
		instruction_processor.DecodeExecute(ir, &pc)		//Decodifica y ejecuta la instrucción
		pc += 1																				  //Incrementa el PC para cargar la siguiente instrucción
	}

}
