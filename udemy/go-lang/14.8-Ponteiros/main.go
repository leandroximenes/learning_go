package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Stats string
}

func UpdateUser(user *User) {
	user.Stats = "updated"
}

func main() {
	fmt.Println("## Pointers ##")
	fmt.Println("")
	fmt.Println("Creating pointers")

	var nome string = "Jhon"
	var nomePtr *string = &nome
	var alias *string = &nome
	var othAlias *string = nomePtr

	fmt.Println("nome		", nome)         // Jhon
	fmt.Println("nomePtr   	", nomePtr) // 0xc000104ef8
	fmt.Println("*nomePtr	", *nomePtr)  // Jhon
	fmt.Println("&nome		", &nome)       // 0xc000104ef8
	fmt.Println("&nomePtr	", &nomePtr)  // 0xc000104ef0

	fmt.Println("*alias		", *alias) // Jhon
	*alias = "The bigger one"
	fmt.Println("&alias		", &alias)      // 0xc000104ef0
	fmt.Println("*alias		", *alias)      // The bigger one
	fmt.Println("nome		", nome)          // The bigger one
	fmt.Println("*othAlias	", *othAlias) // The bigger one

	jhon := User{"Jhon", 42, "new"}
	fmt.Println(jhon)
	UpdateUser(&jhon)
	fmt.Println(jhon)

}
