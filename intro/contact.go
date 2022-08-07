package main

import "fmt"

type contact struct {
	name  string
	age   int
	numbers map[string]string
}

type contact1 struct {
	name  string
	age   int
}

func newContact(name string, age int, numbers map[string]string) contact {
	c := contact{
		name:  name,
		age:   age,
		numbers: numbers,
	}

	return c
}

func newContact1(name string, age int) contact {
	c := contact{
		name:  name,
		age:   age,
	}

	return c
}


func (c contact) displayContact() string {
	display := fmt.Sprintf("Contact: %v (%v ans)\n", c.name, c.age)
	display += "-----------------------------\n"

	for key, value := range c.numbers {
		display += fmt.Sprintf("%v : %v\n", key, value)
	}

	return display
}
