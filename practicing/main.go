package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Meu primeiro progama em go lang")
	auxiliar.Escrever()

	email := "leandroj.r.ximenes"
	if err := checkmail.ValidateHost(email); err != nil {
		fmt.Println("Aconteceu um erro")
		return
	}

	fmt.Printf("tudo ok com o email %s \n", email)
	fmt.Println("fim do programa")
}
