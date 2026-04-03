package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "au au..."
}

type Cat struct{}

func (c Cat) Speak() string {
	return "meow meow"
}

func makeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	fmt.Println("Learning interfaces")

	dog := Dog{}
	cat := Cat{}

	makeSound(dog)
	makeSound(cat)

}
