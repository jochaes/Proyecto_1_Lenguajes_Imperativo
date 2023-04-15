package main

import (
	"fmt"
	"os"

	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/config"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/file_management"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/instruction_processor"
	//"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
)

// Declaracion de Variables
var pc int           //Program Counter: Guarda el numero de instruccion a ejecutar
var ir string        //Instruction Register: Guarda la instrucción que se debe decodificar y ejecutar
var program []string //Memoria del programa
var stack data_structures.Stack
var storage (data_structures.MapTable)
var debugMode bool = config.DebugMode

/*
*

	Carga el programa en memoria
	filame: nombre del archivo a cargar, el archivo debe estar en la carpeta ./test_code

*
*/
func loadProgram(filePath string) {
	var filePtr *os.File = file_management.OpenFile(filePath)
	program = file_management.ReadFile(filePtr)
	file_management.CloseFile(filePtr)
}


// Ciclo de ejecución del programa
func main() {

	//Inicializacion de variables
	pc = 0
	ir = ""
	storage = make(data_structures.MapTable)

	//Carga el programa en memoria
	
	//loadProgram("test_programs/Programa_Prueba.txt") // Programa de prueba de la especificación del proyecto
	loadProgram("test_programs/Sumar_Lista.txt") // Suma una lista de enteros
	//loadProgram("test_programs/Lista_Imprimir.txt") // Imprime una lista al revez
	//loadProgram("test_programs/AND_Multiplicar_Var.txt") // Multiplica una variable por 5 si var > 1 && var < 5
	//loadProgram("test_programs/Logic_Or.txt") //Imprime los resultados de la tabla de verdad de la operacion OR true = 1, false = 0
	//loadProgram("test_programs/Dividir.txt") // x = Y / Z

	//Simulación del ciclo de fetch-decodificación-ejecución
	for {

		// En caso de que el PC se salga del rango del programa, se termina la ejecución
		if pc >= len(program) {
			fmt.Println("End of Execution")
			break
		}

		ir = program[pc]                                               //Carga la instrucción al registro de instrucción
		instruction_processor.DecodeExecute(ir, &pc, &stack, &storage) //Decodifica y ejecuta la instrucción
		//En caso de que la instrucción END sea ejecutada, se termina la ejecución
		if pc < 0 {
			break
		} 
		pc += 1 //Incrementa el PC para cargar la siguiente instrucción
		if debugMode {
			fmt.Println()
		}
	}

	fmt.Println("\nStack y Almacen")
	fmt.Println(stack)
	fmt.Println(storage)

}
