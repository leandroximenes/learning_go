package main

import (
	"errors"
	"fmt"
	"modulo/auxiliar"
	"strings"

	"github.com/badoux/checkmail"
)

const DOMINIO = "gmail.com"

func main() {
	fmt.Println("Meu primeiro progama em go lang")
	auxiliar.Escrever()

	var email string
	var err error

	email, err = buildEmail("leandroj.r.ximenes")
	if err != nil {
		fmt.Println("Erro ao criar email")
		return
	}
	
	if err := checkEmail(email); err != nil {
		fmt.Println("Aconteceu um erro")
		return
	}

	fmt.Printf("tudo ok com o email %s \n", email)
	fmt.Println("fim do programa")
}

func checkEmail(email string) error {
	if err := checkmail.ValidateHost(email); err != nil {
		fmt.Println("Host invalido")
		return err
	}

	if err := checkmail.ValidateMX(email); err != nil {
		fmt.Println("MX invalido")
		return err
	}

	if err := checkmail.ValidateFormat(email); err != nil {
		fmt.Println("Formatod invalido")
		return err
	}

	return nil
}

func buildEmail(email string) (string, error) {
	if strings.Contains(email, "@") {
		return "", errors.New("formato invalido")
	}

	email = email + "@" + DOMINIO
	return email, nil
}
