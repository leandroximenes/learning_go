package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string
	LastName string
	Idt      uint
	Address  Address
}

type Address struct {
	Street  string `json:"street,omitempty"` //rename and omit if empty
	Number  uint   `json:"number,omitempty"`
	ZipCode uint   `json:"zip_code,omitempty,string"` //rename, omit if empty and parse t string
	Country string `json:"-"`                         //skip
}

func main() {
	fmt.Println("Initiate day 3. Structs")
	fmt.Println("Structs are type")

	u := User{
		Name:    "leandro",
		Address: Address{Street: "hidden streat"},
	}

	add := Address{ZipCode: 2, Country: "EUA"}

	uJsoned, _ := json.Marshal(u)
	addresJson, _ := json.Marshal(add)

	fmt.Println(u)
	fmt.Println(string(addresJson))
	fmt.Println(string(uJsoned))
}
