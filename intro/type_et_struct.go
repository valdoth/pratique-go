package main

import "fmt"

type Example struct {
	a string
	b int
	c bool
}

func main() {
	fmt.Println("type et struct")

	myExample := Example{
		a: "Tsiaro",
		b: 30,
		c: true,
	}

	fooExample := Example{"Jane", 28, false}

	fmt.Println(myExample)
	fmt.Println(fooExample)

	myContact := newContact("Tsiaro", 30)
	fmt.Println(myContact)
}
