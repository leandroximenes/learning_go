package airportrobot

import "fmt"

// Write your code here.
type Greeter interface {
	LanguageName() string
	Greet(text string) string
}

type Italian struct {
}

func (i Italian) LanguageName() string {
	return "Italian"
}
func (i Italian) Greet(text string) string {
	return fmt.Sprintf("Ciao %s", text)
}

type Portuguese struct {
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}
func (i Portuguese) Greet(text string) string {
	return fmt.Sprintf("Olá %s", text)
}

func SayHello(greet string, g Greeter) string {
	result := fmt.Sprintf("I can speak %s: %s!", g.LanguageName(), g.Greet(greet))
	return result
}

// Try to solve all the tasks first before running the tests.
// This exercise does not have tests for each individual task.
