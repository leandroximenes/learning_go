package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello!!")

	isEmpty, err := isEmpty("teste")
	if err != nil {
		fmt.Println("New Error")
		return
	}

	if isEmpty {
		fmt.Println("empty strig")
	} else {
		fmt.Println("strig is not empty")
	}

	fmt.Println("Program finished")
}

func isEmpty(text string) (bool, error) {
	if text == "" {
		return true, nil
	}

	return false, nil
}
