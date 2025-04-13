package main

import (
	"errors"
	"fmt"
)

func main() {
	var variable string = "firtst string"
	nextVariable := "second variable"

	fmt.Println(variable)
	fmt.Println(nextVariable)

	color, branch := "white", "BMW"
	fmt.Println(color, branch)

	var isCorrectBeRight = false

	if !isCorrectBeRight {
		var err error = errors.New("try to be right")
		fmt.Println(err)
	}
}
