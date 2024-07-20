package main

import "fmt"

func main() {
	//  Maps - Also called as Hash Tables

	languages := make(map[string]string)

	languages["Js"] = "Javascript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["Js"])

	// Deleting from a map:
	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	// Loops are interesting in Golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}
}
