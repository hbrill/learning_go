package tutorial2

import "fmt"

func Hello(name string) string {
	// var message string is how to declare a variable. := is a shortcut to declare and assign
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
