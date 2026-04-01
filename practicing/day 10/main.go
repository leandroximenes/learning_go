package main

import (
	"errors"
	"fmt"
)

const (
	Child    = "child"
	Teenager = "teenager"
	Adult    = "adult"
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
			fmt.Println("Invalid age type")
			continue
		}

		_, category, err := evaluatePerson(name, age)
		if err != nil {
			fmt.Println(err)
			continue
		}

		counts[category]++
	}

	printSummary(counts)

}

func evaluatePerson(name string, age int) (string, string, error) {
	if name == "" {
		return "", "", errors.New("Name attribute is required")
	}

	if age < 0 {
		return "", "", fmt.Errorf("Invalid age: %d", age)
	}

	category, err := getCategory(age)
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("%s is an %s", name, category), category, nil
}

func getCategory(age int) (string, error) {
	if age < 0 {
		return "", errors.New("invalid age")
	}
	if age < 13 {
		return Child, nil
	}
	if age <= 17 {
		return Teenager, nil
	}
	return Adult, nil
}

func printSummary(counts map[string]int) {
	categories := []string{Child, Teenager, Adult}

	fmt.Println("Resumo: ")

	for _, category := range categories {
		fmt.Printf("%s: %d\n", category, counts[category])
	}

}
