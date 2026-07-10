package greetings

import (
	"errors"
	"fmt"
)

// In Go, a function that starts with a capital letter (an "exported name") can be called by a function not in
// the same package.
func Hello(name string) (string, error) {
	
	if name == "" {
		return "", errors.New("empty name");
	}
	
	// Return a greeting with name parameter
	// ":=" is shorthand for declaring and initializing
	// "%v" is the universal format verb which would print any value using it's default representation
	//		This is typically safer for more versatile, mixed type objects. "%s" would achieve the same results below
	//		as that is the specific formatter for a string ("represent this value as a string").
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}