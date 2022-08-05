package main

import "fmt"

func main() {
	fmt.Println("Les types en go")
	var (
		b bool
		s string
		i int
		u uint
		u8 uint8
		i8 int8
		i16 int16
		u16 uint16
		f float32
	)

	b = true
	s = "Tsiaro"
	i = -15
	u = 15
	u8 = 254
	i8 = 127
	i16 = -21500
	u16 = 40000
	f = 3.14

	fmt.Println(b, s, i, u, u8, i8, i16, u16, f)
}
