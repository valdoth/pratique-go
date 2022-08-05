package main

import "fmt"

func main() {
	fmt.Println("Découverte des opérateur")
	var (
		x int
		y int
	)

	x = 15
	y = 15

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x / y)
	fmt.Println(x * y)
	fmt.Println(x % y)
	fmt.Println("--------------")
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(y < x)
	fmt.Println(y <= x)
	fmt.Println(x < y)
	fmt.Println(x <= y)
	fmt.Println("---------------")
	fmt.Println(x ==y && x != y)
	fmt.Println(x ==y || x != y)
}

