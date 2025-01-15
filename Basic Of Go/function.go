package main

import "fmt"

func main() {
	printHello()
	sum(20, 30)
	// age, name := addition(200, 400)

	//omit the value
	// age, _ := addition(200, 400)
	// fmt.Println("my age", age)
	// fmt.Println("my name", name)

	//Pass by values
	// value := 30
	// passByvalue(value)
	// fmt.Println(value)

	// Pass by references
	// value := 34
	// passByreference(&value)
	// fmt.Println(value)
}

func passByreference(y *int) {
	*y = 20
}

func passByvalue(x int) {
	x = 20
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
