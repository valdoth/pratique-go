package main

import "fmt"

func main() {
	fmt.Println("Les arrays")

	var list [3]int
	list[0] = 10
	list[1] = 20
	list[2] = 30
	fmt.Println(list)
	fmt.Println(list[0])
	fmt.Println(list[1])
	fmt.Println(list[2])


	newList := [...]int{40, 50}
	fmt.Println(newList)

	bigList := [...]int{10, 20, 30, 40, 50, 60, 420, 777777, 50085}
	fmt.Println(bigList)
	for pos, value := range bigList {
		fmt.Println("La valuer à l'indice", pos, "est égale à", value)
	}
}
