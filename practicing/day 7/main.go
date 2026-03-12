package main

import (
	"fmt"
	"strings"
)

type (
	User struct {
		Name     string
		LastName string
	}
)

func main() {
	fmt.Println("Review about pointers")
	fmt.Println("working with slices")

	users := []User{}

	users = append(users, User{Name: "Leandro", LastName: "Ximenes"})
	users = append(users, User{Name: "Daniel", LastName: "Xavier"})

	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

	fmt.Println("trabalhando com maps")

	usersMap := map[string]*User{
		"a1": {Name: "Alessa", LastName: "Serpa"},
		"b2": {Name: "Sara", LastName: "Serpa"},
	}

	for k, v := range usersMap {
		fmt.Printf("dump in item %s: %+v\n", k, v)
		v.Name = strings.ToUpper(v.Name)
		//or
		usersMap[k].LastName = strings.ToUpper(v.LastName)
		fmt.Printf("dump in maps: %+v\n", *v)
	}

	fmt.Printf("dump in maps: %+v\n", usersMap)
}
