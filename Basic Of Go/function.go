package main

import "fmt"

func main() {
	printHello()
	sum(20, 30)
	// age, name := addition(200, 400)

	//omit the value
	age, _ := addition(200, 400)

	fmt.Println("my age", age)
	// fmt.Println("my name", name)
}

func printHello() {
	fmt.Println("Hello")
}

func sum(a int, b int) int {
	// fmt.Println(a + b)
	return a + b
}

// Named Return Values (like 2ta value return kore)
func addition(a int, b int) (res int, name string) {
	// fmt.Println(a + b)
	res = a + b
	name = "Go"
	return
}
