package instruction_processor

import (
	"fmt"
	"strings"
)

/**
	Decodifica la instrucción y la ejecuta 
	instruccion: instruccion a decodificar
	pcPTR: Program Counter
**/
func DecodeExecute(instruction string, pcPtr *int){

	fmt.Println("Decoding", instruction)

	instructionList := strings.Split(instruction, " ")
	
	switch instructionList[0] {
	case "LOAD_FAST":
		fmt.Println("Instruction decoded: LOAD_FAST")
		
	case "STORE_FAST":
		fmt.Println("Instruction decoded: STORE_FAST")

	default:
		fmt.Println("Instruction not implemented")
		
	}
}