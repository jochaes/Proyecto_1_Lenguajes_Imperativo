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
func DecodeExecute(instruction string, pcPtr *int, stack *data_structures.Stack, storage *data_structures.MapTable) {

	fmt.Println("Decoding", instruction)

	instructionList := strings.SplitAfterN(instruction, " ", 2)

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
		fmt.Println("Instruction decoded: BINARY_SUBSTRACT")
		binarySubstract(stack)

	case "BINARY_ADD":
		fmt.Println("Instruction decoded: BINARY_ADD")
		binaryAdd(stack)

	case "BINARY_MULTIPLY":
		fmt.Println("Instruction decoded: BINARY_MULTIPLY")
		binaryMultiply(stack)

	case "BINARY_DIVIDE":
		fmt.Println("Instruction decoded: BINARY_DIVIDE")
		binaryDivide(stack)

	case "BINARY_AND":
		fmt.Println("Instruction decoded: BINARY_AND")
		binaryAnd(stack)

	case "BINARY_OR":
		fmt.Println("Instruction decoded: BINARY_OR")
		binaryOr(stack)

	case "BINARY_MODULO":
		fmt.Println("Instruction decoded: BINARY_MODULO")
		binaryModulo(stack)

	case "STORE_SUBSCR":
		fmt.Println("Instruction decoded: STORE_SUBSCR")
		storeSubscr(stack, storage)
	case "BINARY_SUBSCR":
		fmt.Println("Instruction decoded: BINARY_SUBSCR")
		binarySubscr(stack, storage)
	case "JUMP_ABSOLUTE ":
		fmt.Println("Instruction decoded: JUMP_ABSOLUTE")
		jumpAbsolute(instructionList[1], pcPtr)
	case "JUMP_IF_TRUE ":
		fmt.Println("Instruction decoded: JUMP_IF_TRUE")
		jumpTrue(instructionList[1], stack,pcPtr)
	case "JUMP_IF_FALSE ":
		fmt.Println("Instruction decoded: JUMP_IF_FALSE")
		jumpFalse(instructionList[1],stack, pcPtr)

	default:
		fmt.Println("Instruction not implemented")
	}
}

