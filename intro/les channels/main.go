package main

import (
	"fmt"
	"myprogram/functions"
)

func CalculateValue(intChan chan int) {
	randomNumber := functions.GenerateRandomNumber(50)
	intChan  <- randomNumber
}

func main() {
	fmt.Println("Les channels")

	foo := make(chan int)
	defer close(foo)

	go CalculateValue(foo)

	num := <-foo

	fmt.Println(num)

}
