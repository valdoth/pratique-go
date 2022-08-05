package main

import "fmt"

var x int

func main() {
	fmt.Println("Les scopes")

	x = 5
	fmt.Println(x)
	f()
	showY()
}

func f() {
	x := 10
	fmt.Println(x)
}
