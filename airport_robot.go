package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

type Italian struct {
}

type Portuguese struct {
}

type German struct {
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}

func (i Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func (i German) LanguageName() string {
	return "German"
}

func (i German) Greet(visitorName string) string {
	return fmt.Sprintf("Hallo %s!", visitorName)
}

func SayHello(visitorName string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(visitorName))
}
