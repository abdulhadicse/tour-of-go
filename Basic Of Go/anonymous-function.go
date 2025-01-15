package main

import "fmt"

func main() {
	//basic anonymous func
	// func() {
	// 	fmt.Println("Anonymous functions")
	// }()

	//anonymous function with params
	// func(name string) {
	// 	fmt.Println("Hello ", name)
	// }("John")

	// function body assign a variable/name
	// gretting := func(name string) {
	// 	fmt.Println("Hello ", name)
	// }

	// gretting("Tara") // function call by value;

	// function as a parameter
	// sum := func(x int, y int) int {
	// 	return x + y
	// }

	// res := calculate(20, 30, sum)
	// fmt.Println(res)

	// higher order function or function return a function
	// anotherExecute := execute(10)
	// anotherExecute(60)

}

func execute(a int) func(int) {
	return func(x int) {
		fmt.Println(x + a)
	}
}

func calculate(z int, c int, operation func(x int, y int) int) int {
	return operation(z, c)
}
