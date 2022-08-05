package main

import "fmt"

func updateA(number int) int {
	number = 5
	return number
}

func updateB(item map[string]int) {
	item["bonbon"] = 4
}

func main() {
	number := 10
	fmt.Println(number)
	number = updateA(number)
	fmt.Println(number)

	groceryList := map[string]int {
		"prince": 6,
		"viande": 10,
	}
	fmt.Println(groceryList)
	updateB(groceryList)
	fmt.Println(groceryList)
}
