package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Stats string

const (
	Active   Stats = "active"
	Inactive Stats = "inactive"
)

type (
	User struct {
		Name     string
		LastName string
		Stats    Stats
		Address  struct {
			City    string
			Country string
		}
	}

	UserResponse struct {
		Slug string
		Name string
	}
)

type (
	Preferences struct {
		UserID       string `json:"user_id,omitempty"`
		ReceiveEmail bool   `json:"receive_email"`
	}
)

const (
	WelcomeEmail  string = "a1b2c3d4e5"
	ResetPassword string = "f6g7h8i9j10"
)

func main() {
	user := User{Name: "Leandro", LastName: "Ximenes", Stats: "Other"}
	newUser := User{Name: "Jhon", LastName: "Don", Stats: Active}

	userPreferene := Preferences{ReceiveEmail: true}

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", newUser)

	userResponse := convertUser(user)
	fmt.Printf("%+v\n", userResponse)

	fmt.Printf("Welcome email code is: %s \n", WelcomeEmail)

	fmt.Println("\nchecking JSON formatation:")
	fmt.Printf("%+v\n", userPreferene)

	fmt.Println("\nchecking JSON formatation after Marshal:")
	uJsoned, _ := json.Marshal(userPreferene)
	fmt.Println(string(uJsoned))
}

func convertUser(user User) UserResponse {
	fullName := fmt.Sprintf("%s %s", user.Name, user.LastName)
	slug := strings.ToLower(user.Name) + strings.ToLower(user.LastName)
	userResponse := UserResponse{Slug: slug, Name: fullName}

	return userResponse

}
