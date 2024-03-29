package data_structures

import (
	"errors"
	"fmt"
)

// Estructura que va a componer la pila
type Node struct {
	Valor any
}

// Pila
type Stack []Node

/*
Agrega elementos a la pila
value: El dato a agregar a la pila
*/
func (stack *Stack) Push(value any) {
	if debugMode {
		fmt.Println("stack(push):", value)
	}
	*stack = append(*stack, Node{value})
}

/*
Elimina el top del stack
*/
func (stack *Stack) Pop() (any, error) {
	if len(*stack) > 0 {
		res, _ := (*stack).Peek()
		if debugMode {
			fmt.Println("stack(pop):", res)
		}
		*stack = (*stack)[:len(*stack)-1]
		return res, nil
	} else {
		return 0, errors.New("stack(pop) error: the stack is empty")
	}
}

/*
Devuelve el elemento al tope del stack
*/
func (stack *Stack) Peek() (any, error) {
	if len(*stack) > 0 {
		if debugMode {
			fmt.Println("stack(Peek):", (*stack))
		} //Imprime todo el stack
		peekVal := (*stack)[len(*stack)-1]
		return peekVal.Valor, nil
	} else {
		return nil, errors.New("stack(peek) error: the stack is empty")
	}

}
