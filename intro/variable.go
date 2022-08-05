package main

import "fmt"

func main() {
	var (
		x int
		y string
		z bool
	)
	x = 16
	y = "Tsiaro"
	z = true
	// y := 17
	fmt.Printf("Bonjour à tous, je suis %v, j'ai %vans et j'ai donc plus de 18 ans, c'est %v!\n", y, x, z)
}
