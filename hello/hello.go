package main

import (
	"fmt"
	"github.com/r0B-O/go-beginner/greetings"
)

func main () {
	// Get a  greeting message and print it
	message := greetings.Hello("rob-o")
	fmt.Println(message)
}