package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	// No inheritance in Golang, no super or parent

	tushar := User{"Tushar", "tushar.dimri91@gmail.com", true, 16}
	fmt.Println(tushar)
	fmt.Printf("Tushar's details are: %+v\n", tushar)
	fmt.Printf("Name is %v and email is %v", tushar.Name, tushar.Email)
}
