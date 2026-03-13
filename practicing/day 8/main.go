package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Funções")
	fmt.Println("Funções anonimas")

	func() {
		fmt.Println("Esta é uma funcao anonima")
	}()

	somaTotal := agregra(1, 24, 556, 7676)
	fmt.Printf("Soma é %d\n", somaTotal)

	//lembrando ponteiros
	numero := 10

	var ponteiro *int = &numero
	ponteiro2 := &numero

	fmt.Println(numero, ponteiro, ponteiro2)
	fmt.Println(reflect.TypeOf(numero), reflect.TypeOf(ponteiro), reflect.TypeOf(ponteiro2))

	*ponteiro = 22
	fmt.Println(numero, *ponteiro, *ponteiro2)

}

func agregra(numbers ...int) (total int) {
	for _, number := range numbers {
		total += number
	}

	// apenas o return porque é uma função com retorno nomeado.
	return
}
