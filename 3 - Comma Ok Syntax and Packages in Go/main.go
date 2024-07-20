package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//  pkg.go.dev is the official site for go package docs

	reader := bufio.NewReader(os.Stdin)
	//  NewReader is reading from standard input output i.e the keyboard
	fmt.Println("Enter rating for our pizza: ")

	// Comma Ok Syntax || Error Ok Syntax - For error handling
	input, _ := reader.ReadString('\n')
	//  _ stores error. If we use err as variable name then we'll get error as we have to use variables in Go.
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of Rating is %T", input)
}
