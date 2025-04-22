package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Stats string
}

// this is a function
func UpdateUser(user *User) {
	user.Stats = "updated"
}

// this is a method
func (u *User) SetAge(age int) error {
	u.Age = age
	return nil
}

// ineffective assignment to field User.Age (SA4005)go-staticcheck
func (u User) IncreaseAge() { //It will not work because it's not a pointer and there isn't a return
	u.Age++
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
	paul := User{"Paul", 30, "new"}
	fmt.Println(jhon)

	UpdateUser(&jhon)
	fmt.Println(jhon)

	jhon.SetAge(24)
	fmt.Println(jhon)

	fmt.Println(paul)
	err := paul.SetAge(18)
	if err != nil {
		fmt.Println("Erro to set age")
	}
	fmt.Println(paul)
	paul.IncreaseAge()
	fmt.Println(paul)
}
