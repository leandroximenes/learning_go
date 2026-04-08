package main

import (
	"fmt"
)

type Notifier interface {
	Notify() string
}

type Email struct {
	To      string
	From    string
	Message string
}

func (e Email) Notify() string {
	if e.To == "" {
		panic("To is required")
	}
	return fmt.Sprintf("Sending email to %s", e.To)
}

type SMS struct {
	To      string
	From    string
	Message string
}

func (s SMS) Notify() string {
	return fmt.Sprintf("Sending SMS to %s", s.To)
}

func SendMessage(e Notifier) string {
	return e.Notify()
}

func SendSave(e Notifier) string {
	defer func() {
		if r := recover(); r != "" {
			fmt.Println("Erro ao enviar notificação")
		}
	}()

	return SendMessage(e)
}

func testDeferComRetorno() string {
	defer fmt.Println("Teste na função")
	return "hello World"
}

func main() {
	defer fmt.Println("Finalizado o programa com defer")
	saudacao := testDeferComRetorno()
	fmt.Print(saudacao)
	email := Email{From: "Alessandra", Message: "Enviando um email "}
	sms := SMS{To: "Sara", From: "Sofia", Message: "Enviando um SMS "}

	emailMessage := SendSave(email)
	smsMessage := SendSave(sms)

	fmt.Println(emailMessage)
	fmt.Println(smsMessage)
}
