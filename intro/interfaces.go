package main

import "fmt"

type Animal interface {
	Noise() string
	NumberOfLegs() int
}

type Cat struct {
	Name string
	Breed string
}

type Spider struct {
	Name string
	Breed string
	Venomous bool
}

func main() {
	fmt.Println("Les interfaces")

	cat := Cat {
		Name: "Kitty",
		Breed: "Siamois",
	}
	PrintAnimalInfo(&cat)


	spider := Spider {
		Name: "Spiddy",
		Breed: "Veuve Noire",
		Venomous: true,
	}
	PrintAnimalInfo(&spider)
}


func PrintAnimalInfo(a Animal) {
	fmt.Println("Le bruit de cet animal est", a.Noise(), "et il possède", a.NumberOfLegs(), "pattes!")
}

func (c *Cat) Noise() string {
	return "Miaou"
}

func (c *Cat) NumberOfLegs() int {
	return 4
}

func (s *Spider) Noise() string {
	return "Hiss"
}

func (s *Spider) NumberOfLegs() int {
	return 8
}
