package main

import "fmt"

func update(number int) {
	number = 5
}

func updateWithPointer(number *int, value int) {
	*number = value
}

func main() {
	number := 10
	fmt.Println(number)
	fmt.Println("Adresse mémoire de number:", &number)
	update(number)
	fmt.Println(number)

	myPointer := &number	// pointer
	fmt.Println(myPointer)
	fmt.Println("La valeur de l'adresse mémoire", myPointer, "est", *myPointer)
	updateWithPointer(myPointer, 67)
	fmt.Println("La valeur de l'adresse mémoire", myPointer, "à changé en", *myPointer)

}
