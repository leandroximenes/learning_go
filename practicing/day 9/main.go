package main

import "fmt"

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
		"c3": {ID: "c3", Name: "Sara", LastName: "Serpa", Status: InActive},
	}

	users["a2"] = newUser("b2", "Alessandra", "Serpa")

	ActivateUsers(users, "a1", "b2", "c3")

	for _, user := range users {
		fmt.Printf("%+v\n", *user)
	}

}

func newUser(id, name, lastName string) (u *User) {
	user := User{ID: id, Name: name, LastName: lastName, Status: Active}

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
