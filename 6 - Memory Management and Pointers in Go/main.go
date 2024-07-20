package main

import "fmt"

func main() {
	// Memory Allocation and Deallocqtion happens autmatically in Go
	/*
		Memory Management:
			new():
			--> Allocate memory but not initialized.
			--> You will get a memory address.
			--> Zeroed Storage.
			make():
			--> Allocate memory and initialized.
			--> You will get a memory address.
			--> Non-Zeroed Storage.

			Garbage Collection happens automatically
			Out Of Scope or nil.
			https://pkg.go.dev/runtime
	*/

	var ptr *int

	fmt.Println("value of ptr is: ", ptr) // nil

	mynumber := 23
	var pointer = &mynumber

	fmt.Println("value of ptr is: ", pointer)
	fmt.Println("value inside the pointer is: ", *pointer)

	*pointer = *pointer * 2
	fmt.Println("New value is: ", mynumber)
}
