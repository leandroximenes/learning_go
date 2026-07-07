package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Checking slice copy")

	matrizSrc := [][]int{{1, 2}, {3, 4}}
	matrizCopy := slices.Clone(matrizSrc)

	matrizCopy[0][0] = 99

	fmt.Println("slice src:", matrizSrc)
	fmt.Println("slice copy:", matrizCopy)

	meuNome := "teste"
	var nome *string
	nome = &meuNome

	*nome = "Valor do ponteiro alterado"
	fmt.Printf("O nome é um endereço de memoria: %+v \n", nome)
	fmt.Printf("Para imprimir o valor de um ponteiro: %+v \n", *nome)

	fmt.Println("Testando funções")
	items := []int{1, 2, 3, 4, 5}

	items = simpleMap(items, func(val int) int {
		return val * val
	})
	fmt.Printf("Resultado da funçao map: %+v \n", items)

	items = simpleFilter(items, func(val int) bool {
		return val%2 == 0
	})
	fmt.Printf("Resultado da funçao filter: %+v \n", items)

	result := simpleReduce(items, func(acc, val int) int {
		return acc + val
	})
	fmt.Printf("Resultado da funçao reduce: %+v \n", result)
}

func simpleMap(items []int, f func(val int) int) []int {
	var result []int

	for _, item := range items {
		result = append(result, f(item))
	}

	return result
}

func simpleFilter(items []int, f func(val int) bool) []int {
	var result []int

	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}

	return result
}

func simpleReduce(items []int, f func(acc int, val int) int) int {
	var result int

	for _, item := range items {
		result = f(result, item)
	}

	return result
}
