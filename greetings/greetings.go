package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Function Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return a message with this error
	if name == "" {
		return "", errors.New("name is empty")
	}
	// Return a message that embeds the name in a message
	// Create a message using a random greeting and format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets intial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of the set of greeting messages. The returned
// message is selected at random
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi. %v, Welcome!",
		"Great to see you, %v",
		"Hail %v, Well met!",
	}

	// Return a randomly selected message format by specifying a random
	// index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
