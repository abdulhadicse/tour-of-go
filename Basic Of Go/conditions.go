package main

import "fmt"

func main() {
	// if 28 == 30 {
	// 	fmt.Println("Hello True")
	// } else if 40 == 40 {
	// 	fmt.Println("It's 40!")
	// } else {
	// 	fmt.Println("False")
	// }

	switch 40 % 8 {
	case 4, 0:
		fmt.Println("It's 4")
	case 5:
		fmt.Println("It's 5")
	default:
		fmt.Println("It's failed")
	}

}
