package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of predefined Logger, including the log entry prefix and a flag to disable 
	// printing, the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	
	names := []string{
		"Gladys",
		"Malik",
		"Makayla",
	}
	

	// Get a greeting message and print it
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}