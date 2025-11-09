package main

import (
	"fmt"

	"example.com/greetings"

	"rsc.io/quote"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Tony ID")
	fmt.Println(message)
	fmt.Println(quote.Go())
}
