package main

import (
	"fmt"
)

const (
	PASSWORD string = "1234"
)

type LoginError struct {
	Username string
	Reason   string
}

func (err *LoginError) Error() string {
	return fmt.Sprintf("user %s failed to login: %s", err.Username, err.Reason)
}

func ValidateFields(userName string, password string) error {
	var err *LoginError = &LoginError{Username: userName}

	if userName == "" {
		err.Reason = "username is required"
		return err
	}

	if password == "" {
		err.Reason = "password is required"
		return err
	}


	return nil
}

func Login(userName string, password string) error {
	err := ValidateFields(userName, password)

	if err != nil {
		return err
	}

	if password != PASSWORD {
		return &LoginError{
			Username: userName,
			Reason:   "invalid password",
		}
	}

	return nil
}

func main() {
	fmt.Println("Starting review with interfaces and errors handling ")

	err := Login("leandroximenes", "")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("user log in with success")
}
