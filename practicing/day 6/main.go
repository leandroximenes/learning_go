package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name     string
	LastName string
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

	user := getUser("1")
	if user == nil {
		fmt.Println("No user found")
		return
	}

	changeName(user)
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
	(*u).Name = strings.ToUpper(u.Name)
	// or
	u.LastName = strings.ToUpper(u.LastName)
}

func getUser(id string) (user *User) {
	if id != "1" {
		return nil
	}

	u := User{Name: "leandro", LastName: "ximenes"}
	return &u
}
