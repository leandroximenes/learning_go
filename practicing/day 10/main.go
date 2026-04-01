package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Day 10. Let's go!")

	name := "João"
	age := 36

	message, _, err := evaluatePerson(name, age)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(message)

	// New topic
	fmt.Println("slices & maps")

	people := []map[string]interface{}{
		{"name": "Ana", "age": 10},
		{"name": "Bia", "age": 15},
		{"name": "Lia", "age": 22},
		{"name": "", "age": 30},
	}

	counts := make(map[string]int)
	for _, person := range people {
		name, ok := person["name"].(string)
		if !ok {
			fmt.Println("Invalid name type")
			continue
		}

		age, ok := person["age"].(int)
		if !ok {
			fmt.Println("Invalid name type")
			continue
		}

		if _, category, err := evaluatePerson(name, age); err != nil {
			fmt.Println(err)
			continue
		} else {
			counts[category]++
		}
	}

	fmt.Println("Resumo:")
	fmt.Printf("child: %d \n", counts["child"])
	fmt.Printf("teenager: %d \n", counts["teenager"])
	fmt.Printf("adult: %d \n", counts["adult"])

}

func evaluatePerson(name string, age int) (string, string, error) {
	if name == "" {
		return "", "", errors.New("Name attribute is required")
	}

	if age < 0 {
		return "", "", errors.New("Invalid age")
	}

	if age < 13 {
		return fmt.Sprintf("%s is a child", name), "child", nil
	} else if age <= 17 {
		return fmt.Sprintf("%s is a teenager", name), "teenager", nil
	}

	return fmt.Sprintf("%s is an adult", name), "adult", nil
}
