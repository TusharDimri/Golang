package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomoato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3]) // last of range is not inclusive
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[:3])
	// fmt.Println(fruitList)

	// Using Memory Allocation

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 666 //Error - Index out of bounds

	highScore = append(highScore, 555, 666, 777) // Surprisingly, no error
	fmt.Println((highScore))

	sort.Ints(highScore)
	fmt.Println(highScore)

	// How to remove a value from slice based on index:
	var courses = []string{"reactjs", "javascript", "python", "swift", "ruby"}
	fmt.Println(courses)
	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
