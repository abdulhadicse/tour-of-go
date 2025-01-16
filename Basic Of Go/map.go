package main

import "fmt"

func main() {
	fmt.Println("Hello Map")

	//creating map and initialization

	// var myMap map[string]int = map[string]int{
	// 	"color": 203,
	// 	"age":   20,
	// }

	myMap := map[string]int{
		"age":  20,
		"name": 30,
	}

	// accessing elements
	// myMap["age"]

	//modifying value
	// myMap["age"] = 40

	//adding new elements
	// myMap["class"] = 8

	//iterate elements
	// for k, v := range myMap {
	// 	fmt.Println(k, v)
	// }

	//deleting elements
	// delete(myMap, "age")

	fmt.Println(myMap)
}
