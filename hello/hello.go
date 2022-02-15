package main

import (
	"fmt"
	"log"

	"github.com/r0B-O/go-beginner/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice or array of names
	names := []string{"rob-o", "Sree", "Kiran"}

	// Get a  greeting message and print it
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// in the console
	for _, greeting := range message {
		fmt.Println(greeting)
	}
}
