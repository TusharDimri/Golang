package main

import "fmt"

// jwtToken := 3000 // Not allowed here, its allowed only inside a method
// var jwtToken int = 30000 // Allowed

const LoginToken string = "wrnvwrivbqve" // Public Variable

// if first first letter is capital then the variable is public

func main() {
	var username string = "Tushar Dimri"
	fmt.Println("Value: ", username)
	fmt.Printf("Type: %T", username)

	var isLoggedIn bool = false
	fmt.Println("Value: ", isLoggedIn)
	fmt.Printf("Type: %T", isLoggedIn)

	var small uint8 = 12
	// var small uint8 = 256 Out of bounds
	// var small uint = 12
	fmt.Println("Value: ", small)
	fmt.Printf("Type: %T", small)

	var smallFloat float32 = 32.32
	fmt.Println("Value: ", smallFloat)
	fmt.Printf("Type: %T", smallFloat)

	// Dafault Values and Aliases
	var anotherVariable int
	fmt.Println("Value: ", anotherVariable) // Default Value is printed
	fmt.Printf("Type: %T", anotherVariable)

	// Implicit Type
	var website = "oakshade-ai.in" // Handled by Lexer
	// website = 3 // Error
	fmt.Println(website)

	// No var style
	numberoOfUsers := 30000.0
	fmt.Println(numberoOfUsers)

	fmt.Println("Value: ", LoginToken)
	fmt.Printf("Type: %T", LoginToken)
}
