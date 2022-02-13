package greetings

import "fmt"

// Function Hello returns a greeting for the named person
func Hello(name string) string {
	// Return a message that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	
	return message
}