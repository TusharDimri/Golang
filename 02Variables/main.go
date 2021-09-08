package main

import "fmt"

// jwtToken := 3000000
// Above line will give error as we are allowed to use := operator(Walrus Operator) only inside a method.

// var jwtToken = 3000000
// var jwtToken int = 3000000
// Above 2 lines are fine

const LoginToken string = "gsvsfibVsvVS"

// Above variable is a Public Varible as its first letter is capital.

func main() {
	var username string = "Tushar Dimri"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T\n", smallVal)

	/*
		var smallVal uint8 = 256
		fmt.Println(smallVal)
		fmt.Printf("Variable is of type %T\n", smallVal)

		Above lines will give error because uint can hold values between 0 and 255 only(Both 0 and 255 included).
	*/

	var smallFloat float32 = 255.455134513753847565
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T\n", smallFloat)

	var largeFloat float64 = 255.455134513753847565
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type %T\n", largeFloat)

	// Dafault values and aliases:-

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type %T\n", anotherVariable)
	// So, 0 is the default value.

	var emptyString string
	fmt.Println(emptyString)
	fmt.Printf("Variable is of type %T\n", emptyString)
	// So, '\0' is the default value.

	// implicit type:-
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type %T\n", website)

	// no var style
	no_of_users := 3000000
	fmt.Println(no_of_users)
	fmt.Printf("Variable is of type %T\n", no_of_users)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T\n", LoginToken)

}
