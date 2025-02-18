package main

import (
	"bunexample/gointernal"
)

func main() {
	Foo()

}

func Foo() {
	// Create a new instance of the server
	server := gointernal.NewServer()

	// Start the server
	server.Run("8080")
}
