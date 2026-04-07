package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct {
	Name string
}

func (d Dog) MakeSound() string {
	return "au au"
}

func MakeSound(a Animal) string {
	return a.MakeSound()
}

func main() {
	fmt.Println("Revisando interfaces")

	myDog := Dog{Name: "Meg"}
	fmt.Println(myDog.MakeSound())

	fmt.Println(MakeSound(myDog))
}
