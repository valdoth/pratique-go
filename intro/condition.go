package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Les conditions en go")
	rand.Seed(time.Now().UnixNano())
	if x := rand.Int(); x%2 == 0 {
		fmt.Println(x, "est un nombre paire!")
	} else {
		fmt.Println(x, "est un nombre impaire!")
	}

	y := rand.Int() % 2

	if y == 0 {
		fmt.Println(y, "Paire")
	} else {
		fmt.Println(y, "Impaire")
	}

	age := 19

	if age > 18 {
		fmt.Println("Je suis majeur!")
	} else if age == 18 {
		fmt.Println("Je viens tout juste d'être majeur!")
	} else {
		fmt.Println("Je suis mineur!")
	}
}
