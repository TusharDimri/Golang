package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok / error ok syntax:-
	input, _ := reader.ReadString('\n')
	/* comma ok / comma error syntax for error handling. if everything goes ok then we store the value input by the user in the variable input but in case we get an error the error will be stored in the variable _ .
	_ is used to name variables which we don't intend to use.
	*/

	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of the rating is %T", input)
}
