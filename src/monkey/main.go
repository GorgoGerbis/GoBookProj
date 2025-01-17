// main.go

package main // Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).

import (
	// Standard library packages:
	"fmt" // Import fmt package: Contains functions for formatting text, including printing to the console.
	"os"
	"os/user"
	"monkey/repl"
)

// A main function executes by default when you run 'go run .' the main package.
func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("HEllo %s! This is the Monkey Programming Language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
