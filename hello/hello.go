package main

import (
	"fmt"
	"log"
	"labkit.in/greetings"
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)	

	names := []string {"vidhya","arun","aathi","ayyanar","naji"}

	// Request a greeting message.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
