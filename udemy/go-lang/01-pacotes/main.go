package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	auxiliar.Write()

	error := checkmail.ValidateFormat("leandroximenes.com")
	fmt.Println(error)
}
