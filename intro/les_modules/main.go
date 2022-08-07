package main

import (
	"fmt"
	"github.com/valdoth/myprogram/functions"
)

func main() {
	fmt.Println("Les modules")

	var fooVar functions.Foo
	fooVar.TypeInt = 18
	fooVar.TypeString = "Tsiaro"
	fooVar.TypeBool = true

	fmt.Println(fooVar.TypeInt)
	fmt.Println(fooVar.TypeString)
	fmt.Println(fooVar.TypeBool)
}
