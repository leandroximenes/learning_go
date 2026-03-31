package main

import (
	"fmt"
	"sort"
)

type Status string

const (
	Active   Status = "active"
	InActive Status = "inactive"
)

type User struct {
	ID       string
	Name     string
	LastName string
	Status   Status
}

func main() {
	fmt.Println("Vamos começar a revisão")

	users := map[string]*User{
		"a1": newUser("a1", "Leandro", "Ximenes"),
		"b2": newUser("b2", "Alessandra", "Serpa"),
		"c3": &User{ID: "c3", Name: "Sara", LastName: "Serpa", Status: InActive},
	}

	ActivateUsers(users, "a1", "b2", "c3")
	userList := ListActiveUsers(users)

	for _, user := range users {
		fmt.Printf("%+v\n", *user)
	}

	sort.Slice(userList, func(i, j int) bool {
		return userList[i].Name < userList[j].Name
	})
}

func newUser(id, name, lastName string) (u *User) {
	user := User{
		ID:       id,
		Name:     name,
		LastName: lastName,
		Status:   Active,
	}

	return &user
}

func ActivateUsers(users map[string]*User, ids ...string) {
	for _, id := range ids {
		user, isUser := users[id]

		if isUser {
			user.Status = Active
		}
	}
}

func ListActiveUsers(users map[string]*User) []User {
	var activeUsers []User
	for _, user := range users {
		if user.Status == Active {
			activeUsers = append(activeUsers, *user)
		}
	}

	return activeUsers
}
