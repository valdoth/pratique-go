package main

import (
	"errors"
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("Bonjour à tous, je m'appelle %v.\n", name)
}

func calculatePerimRect(lo, la int) int {
	return 2 * (lo + la)
}

func sayBye(name string) (string, error) {
	if name == "" {
		return "", errors.New("Erreur: Vous avez oublié de spécifier un nom!")
	}
	byeMessage := fmt.Sprintf("%v s'en va! Bonne soirée!", name)
	return byeMessage, nil
}

func main() {
	fmt.Println("Les fonctions")

	sayHello("Tsiaro")

	r1 := calculatePerimRect(5, 2)
	fmt.Printf("Le périmétre de mon réctangle est: %v.\n", r1)

	fmt.Println(sayBye("Tsiaro"))
	fmt.Println(sayBye(""))
}
