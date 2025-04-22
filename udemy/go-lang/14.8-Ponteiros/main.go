package main

import "fmt"

func main() {
	fmt.Println("## Pointers ##")
	fmt.Println("")
	fmt.Println("Creating pointers")

	var nome string = "Jhon"
	var nomePtr *string = &nome
	var alias *string = &nome
	var othAlias *string = nomePtr
	
	println("nome		", nome)		// Jhon
	println("nomePtr   	", nomePtr)		// 0xc000104ef8
	println("*nomePtr	", *nomePtr)	// Jhon
	println("&nome		", &nome)		// 0xc000104ef8
	println("&nomePtr	", &nomePtr)	// 0xc000104ef0
	
	println("*alias		", *alias)	// Jhon
	*alias = "The bigger one"
	println("&alias		", &alias)	// 0xc000104ef0
	println("*alias		", *alias)	// The bigger one
	println("nome		", nome)	// The bigger one
	println("*othAlias	", *othAlias)	// The bigger one



}
