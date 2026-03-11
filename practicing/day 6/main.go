package main

import "fmt"

type User struct {
	Name string
}

func main() {
	fmt.Println("Pointers")

	var variable1 int = 10
	var variable2 int = variable1
	var pointer *int

	variable1++

	pointer = &variable2

	*pointer = 52

	fmt.Println(variable1, variable2)
	fmt.Println(*pointer)

	var exerc2 int = 2

	changeValue(&exerc2)

	fmt.Println(exerc2)

	user := User{}
	changeName(&user)
	fmt.Println(user.Name)
	fmt.Printf("%+v\n", user)

	//1. Store memory address
	//2. the value of the pointer p
	//3. 20
}

func changeValue(n *int) {
	*n = 50
}

func changeName(u *User) {
	(*u).Name = "leandro"
}
