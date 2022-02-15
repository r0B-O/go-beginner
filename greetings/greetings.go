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

// Hellos returns a map that associates each of the named people with
// a greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)
	// Loop through the received slice names, calling the Hello function
	// to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map associate the received message with the name
		messages[name] = message
	}
	return messages, nil
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
