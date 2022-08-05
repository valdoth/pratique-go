package main

import "fmt"

func main() {
	fmt.Println("Les boucles")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	x = 0

	for {
		if x > 4 {
			break
		}
		fmt.Println(x)
		x++
	}

	x = 0
	for ; x <= 10; x++ {
		if x%2 == 1 {
			continue
		}
		fmt.Println(x)
	}
}
