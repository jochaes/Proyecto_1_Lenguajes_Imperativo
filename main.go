package main

import(
	"fmt"
	"os"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/file_management"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/instruction_processor"
)

//Declaracion de Variables
var pc int 						//Program Counter: Guarda el numero de instruccion a ejecutar
var ir string 				//Instruction Register: Guarda la instrucción que se debe decodificar y ejecutar
var program []string 	//Memoria del programa

/**
	Carga el programa en memoria 
	filame: nombre del archivo a cargar, el archivo debe estar en la carpeta ./test_code
**/
func loadProgram(fileName string){
	var filePtr *os.File = file_management.OpenFile("test_code/"+fileName)
	program = file_management.ReadFile(filePtr)
	file_management.CloseFile(filePtr)
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
