package main

import "fmt"

type address struct {
	street string
	number int
}

type person struct {
	address
	name string
	age  int
}

func main() {
	address := address{"street nine", 20}

	var user person
	user.age = 10
	user.name = "leandro"
	user.address = address

	custumer := person{address, "Alessa", 20}

	fmt.Println(user)
	fmt.Println(custumer)

}
