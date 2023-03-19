package instruction_processor

import (
	"fmt"
	"strings"

	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
)

/*
Decodifica la instrucción y la ejecuta
	instruccion: instruccion a decodificar
	pcPTR: Program Counter
*/
func DecodeExecute(instruction string, pcPtr *int, stack *data_structures.Stack, storage *data_structures.MapTable){

	fmt.Println("Decoding", instruction)

	instructionList := strings.SplitAfterN(instruction, " ",2)
	
	switch instructionList[0] {
	
	case "LOAD_CONST ":
		fmt.Println("Instruction decoded: LOAD_CONST")
		loadConst(instructionList[1], stack)

	case "LOAD_FAST ":
		fmt.Println("Instruction decoded: LOAD_FAST")
		loadFast(instructionList[1], stack, storage)
		
	case "STORE_FAST ":
		fmt.Println("Instruction decoded: STORE_FAST")
		storeFast(instructionList[1], stack, storage)

	case "LOAD_GLOBAL ":
		fmt.Println("Instruction decoded: LOAD_GLOBAL")
		loadGlobal(instructionList[1], stack)

	case "CALL_FUNCTION ":
		fmt.Println("Instruction decoded: CALL_FUNCTION")
		callFunction(instructionList[1], stack)
	
	case "COMPARE_OP ":
		fmt.Println("Instruction decoded: COMPARE_OP")
		compareOp(instructionList[1], stack)

	case "BINARY_SUBSTRACT":
		fmt.Println("Instruction decoded: BINARY_SUBSTRACt")
		binarySubstract(stack)

	default:
		fmt.Println("Instruction not implemented")
		
	}
}