package main

import "fmt"

func main() {
	fmt.Println("Lier une fonction à notre type")

	myContact := newContact("Alex", 30, map[string]string{"Fixe": "10.14.25.36.35", "Portable": "12.12.12.12"})
	// fmt.Println(myContact)
	fmt.Println(myContact.displayContact())

	myNewContact := newContact("Justine", 27, map[string]string{"Fixe": "18.52.98.14.25", "Portable": "13.14.15.16"})
	// fmt.Println(myContact)
	fmt.Println(myNewContact.displayContact())
}
