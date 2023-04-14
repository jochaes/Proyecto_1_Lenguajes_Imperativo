package instruction_processor

import (
	"fmt"
	"strings"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/config"
	"github.com/jochaes/Proyecto_1_Lenguajes_Imperativo/data_structures"
	
)

var debugMode bool = config.DebugMode
/*
Decodifica la instrucción y la ejecuta

	instruccion: instruccion a decodificar
	pcPTR: Program Counter
*/
func DecodeExecute(instruction string, pcPtr *int, stack *data_structures.Stack, storage *data_structures.MapTable) {

	if debugMode{ fmt.Println("Decoding", instruction)}

	instructionList := strings.SplitAfterN(instruction, " ", 2)

	switch instructionList[0] {

	case "LOAD_CONST ":
		if debugMode{fmt.Println("Instruction decoded: LOAD_CONST")}
		loadConst(instructionList[1], stack)

	case "LOAD_FAST ":
		if debugMode{fmt.Println("Instruction decoded: LOAD_FAST")}
		loadFast(instructionList[1], stack, storage)

	case "STORE_FAST ":
		if debugMode{fmt.Println("Instruction decoded: STORE_FAST")}
		storeFast(instructionList[1], stack, storage)

	case "LOAD_GLOBAL ":
		if debugMode{fmt.Println("Instruction decoded: LOAD_GLOBAL")}
		loadGlobal(instructionList[1], stack)

	case "CALL_FUNCTION ":
		if debugMode{fmt.Println("Instruction decoded: CALL_FUNCTION")}
		callFunction(instructionList[1], stack)

	case "COMPARE_OP ":
		if debugMode{fmt.Println("Instruction decoded: COMPARE_OP")}
		compareOp(instructionList[1], stack)

	case "BINARY_SUBSTRACT":
		if debugMode{fmt.Println("Instruction decoded: BINARY_SUBSTRACT")}
		binarySubstract(stack)

	case "BINARY_ADD":
		if debugMode{fmt.Println("Instruction decoded: BINARY_ADD")}
		binaryAdd(stack)

	case "BINARY_MULTIPLY":
		if debugMode{fmt.Println("Instruction decoded: BINARY_MULTIPLY")}
		binaryMultiply(stack)

	case "BINARY_DIVIDE":
		if debugMode{fmt.Println("Instruction decoded: BINARY_DIVIDE")}
		binaryDivide(stack)

	case "BINARY_AND":
		if debugMode{fmt.Println("Instruction decoded: BINARY_AND")}
		binaryAnd(stack)

	case "BINARY_OR":
		if debugMode{fmt.Println("Instruction decoded: BINARY_OR")}
		binaryOr(stack)

	case "BINARY_MODULO":
		if debugMode{fmt.Println("Instruction decoded: BINARY_MODULO")}
		binaryModulo(stack)

	case "STORE_SUBSCR":
		if debugMode{fmt.Println("Instruction decoded: STORE_SUBSCR")}
		storeSubscr(stack, storage)

	case "BINARY_SUBSCR":
		if debugMode{fmt.Println("Instruction decoded: BINARY_SUBSCR")}
		binarySubscr(stack, storage)

	case "JUMP_ABSOLUTE ":
		if debugMode{fmt.Println("Instruction decoded: JUMP_ABSOLUTE")}
		jumpAbsolute(instructionList[1], pcPtr)

	case "JUMP_IF_TRUE ":
		if debugMode{fmt.Println("Instruction decoded: JUMP_IF_TRUE")}
		jumpTrue(instructionList[1], stack,pcPtr)

	case "JUMP_IF_FALSE ":
		if debugMode{fmt.Println("Instruction decoded: JUMP_IF_FALSE")}
		jumpFalse(instructionList[1],stack, pcPtr)

	case "BUILD_LIST ":
		if debugMode{fmt.Println("Instruction decoded: BUILD_LIST")}
		buildList(instructionList[1],stack)

	case "END":
		if debugMode{fmt.Println("Instruction decoded: END")}
		end()
		
	default:
		if debugMode{fmt.Println("Instruction not implemented")}
	}
}

