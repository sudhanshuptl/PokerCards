package main

import "fmt"

type bot interface {
	// if there is any type that has function getGreeting with return type of string is not part of this interface
	getGreeting() string
}

type englishBot struct{}
type hinglishBot struct{}

// Since we are not using varible of type englishBot inside getGreeting function so removed it from reciever
func (englishBot) getGreeting() string {
	// very custom logic
	return "Hello World"
}

// Since we are not using varible of type hinglishBot inside getGreeting function so removed it from reciever
func (hinglishBot) getGreeting() string {
	// very custom logic
	return "namaste"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
