package greetings

import (
	"errors"
	"fmt"
)

// Function Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return a message with this error
	if name == "" {
		return "", errors.New("Name is empty")
	}
	// Return a message that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	
	return message, nil
}