package main

import (
	"day5/extras"
	"fmt"
)

const (
	GLOBAL_ID string = "a1b2c3d5e"
)

type (
	Address struct {
		City string `json:"city"`
	}

	User struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		LastName string `json:"last_name"`
		Address  Address
	}

	AdminUser struct {
		User
		Password string `json:"password"`
	}
)

func main() {
	fmt.Println("Reviweing things")
	fmt.Println("#1 Packages")
	extras.SayHello()

	fmt.Println("#const and structs")

	user := User{ID: GLOBAL_ID, Address: Address{City: "Brasilia"}}
	user.LastName = "Ximenes"

	admin := AdminUser{}
	admin.Password = "a1"

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", admin)

	fmt.Println("#slice")

	slice := []int{10, 20, 30, 40}
	slice = append(slice, 50)

	fmt.Printf("%+v\n", slice)

}
