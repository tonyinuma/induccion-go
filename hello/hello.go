package main

import (
	"fmt"
	"log"

	"example.com/greetings"

	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Tony", "Luis", "Carlos"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// Get a greeting message and print it.
	fmt.Println(messages)
	fmt.Println(quote.Go())
}
