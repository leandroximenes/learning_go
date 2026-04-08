package main

import "fmt"

type Notifier interface {
	Notify() string
}

type Email struct {
	To      string
	From    string
	Message string
}

func (e Email) Notify() string {
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

func main() {
	email := Email{To: "Leandro", From: "Alessandra", Message: "Enviando um email "}
	sms := SMS{To: "Sara", From: "Sofia", Message: "Enviando um SMS "}

	emailMessage := SendMessage(email)
	smsMessage := SendMessage(sms)

	fmt.Println(emailMessage)
	fmt.Println(smsMessage)
}
