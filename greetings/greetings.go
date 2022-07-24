package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// init sets initial values for variables used in the function.
func init() {

	x := time.Now().UnixNano()
	rand.Seed(x)
	log.Println(x)
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// tstmessage := fmt.Sprintf(randomFormat(), name)
	tstmessage := fmt.Sprintf(randomFormat(), name)
	return tstmessage, nil
}
