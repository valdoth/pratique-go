package main

import "fmt"

func main() {
	x, y := 50, 50
	fmt.Printf("Salut à tous!\n")
	fmt.Printf("Salut à tous!\n")
	fmt.Printf("Mon nombre (default): %v\n", x)
	fmt.Printf("Mon nombre (base 2): %b\n", x)
	fmt.Printf("Mon nombre (base 8): %O\n", x)
	fmt.Printf("Mon nombre (base 10): %v\n", x)
	fmt.Printf("Mon nombre (base 16): %X\n", x)

	fmt.Printf("La valeur de X est égale à la valeur de y -> %t\n", x == y)

	fmt.Printf("L'écriture scientifique de 128.165559 -> %E\n", 128.165559)
}
