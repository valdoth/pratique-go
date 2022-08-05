package main

import "fmt"

func main() {
	fmt.Println("Les maps")

	supermarketPrice := map[string]int {
		"prince": 8,
		"eau": 2,
		"meat": 6,
		"viande": 5,
	}
	fmt.Println(supermarketPrice)
	fmt.Println(supermarketPrice["prince"])

	for key, value := range supermarketPrice {
		fmt.Println(key, value)
	}

	daysInAYear := 0

	daysInAMonth := map[string] int {
		"Janvier": 31, "Février": 28, "Mars": 31,
		"Avril": 30, "Mai": 31, "Juin": 30,
		"Juillet": 31, "Aout": 31, "Septembre": 30,
		"Octobre": 31, "Novembre": 30, "Decembre": 31,
	}

	for _, value := range daysInAMonth {
		daysInAYear += value
	}

	fmt.Println("Nombre de jours dans une annéé:", daysInAYear, "jours!")
}
