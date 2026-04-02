package main

import (
	"errors"
	"fmt"
)

type Category string

func (c Category) String() string {
	switch c {
	case Child:
		return "Child"
	case Teenager:
		return "Teenager"
	case Adult:
		return "Adult"
	default:
		return "Unknown"
	}
}

func (c Category) Article() string {
	switch c {
	case Adult:
		return "an"
	default:
		return "a"
	}
}

type Person struct {
	Name string
	Age  int
}

const (
	Child    Category = "child"
	Teenager Category = "teenager"
	Adult    Category = "adult"
)

func main() {
	fmt.Println("Day 10. Let's go!")

	personExample := Person{Name: "João", Age: 5}
	message, _, err := evaluatePerson(personExample)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(message)

	// New topic
	fmt.Println("slices & maps")

	people := []Person{
		{Name: "Ana", Age: 10},
		{Name: "Bia", Age: 15},
		{Name: "Lia", Age: 22},
		{Name: "", Age: 30},
	}

	counts := make(map[Category]int)
	for _, person := range people {
		_, category, err := evaluatePerson(person)
		if err != nil {
			fmt.Println(err)
			continue
		}

		counts[category]++
	}

	printSummary(counts)

}

func evaluatePerson(person Person) (string, Category, error) {
	if person.Name == "" {
		return "", "", fmt.Errorf("Name attribute is required")
	}

	category, err := getCategory(person.Age)
	if err != nil {
		return "", "", err
	}
	message := fmt.Sprintf("%s is %s %s", person.Name, category.Article(), category)
	return message, category, nil
}

func getCategory(age int) (Category, error) {
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

func printSummary(counts map[Category]int) {
	categories := []Category{Child, Teenager, Adult}

	fmt.Println("Resumo:")

	for _, category := range categories {
		fmt.Printf("%s: %d\n", category, counts[category])
	}

}
