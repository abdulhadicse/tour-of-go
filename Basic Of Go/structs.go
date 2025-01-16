package main

import "fmt"

// declare
type Person struct {
	Name string
	Age  int
}

func (p *Person) Print() {
	fmt.Printf("Hello %s, Your are now %d\n", p.Name, p.Age)
	p.Age = 299
}

func main() {
	fmt.Println("Hello Structs")

	// assign value
	p := Person{"Abul", 20}
	p.Print()

	// accessing value
	// fmt.Println(p.Name)
	// fmt.Println(p.Age)

	// update value
	// p.Age = 100

	fmt.Println(p)
}
