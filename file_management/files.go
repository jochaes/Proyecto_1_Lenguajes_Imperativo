/*
Este paquete se encarga de las operaciones sobre archivos (abrir, leer, cerrar)
*/
package file_management

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Abre y devuelve el puntero a un archivo

	fileName: ruta del archivo
*/
func OpenFile(fileName string) *os.File {

	//Abre el Archivo
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return f
}

/*
Lee, guarda y devuelve las lineas en un slice

	filePtr: puntero al archivo abierto
*/
func ReadFile(filePtr *os.File) []string {
	var instructions []string

	scanner := bufio.NewScanner(filePtr)

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return instructions
}

/*
Cierra un archivo previemante abierto

	fileptr: puntero al archivo
*/
func CloseFile(filePtr *os.File) {
	filePtr.Close()

}
