package main

import "fmt"

func main() {
	fmt.Println("=== Operators ===")

	a := 20
	b := 7
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	// fmt.Println(a++)
	// fmt.Println(b--)

	// Logical Operators
	fmt.Println(a == 20 && b == 7)
	fmt.Println(a == 19 || b == 8)
	fmt.Println(!(a == 56))

	// Bitwise Operators
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)

	fmt.Println(b >> 2)
	fmt.Println(b << 2)
}
