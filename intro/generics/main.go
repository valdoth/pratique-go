package main

import "fmt"

func main() {
	fmt.Println("Les generics")

	fmt.Println(min[float32](0.5, 0.7))
	fmt.Println(min[int32](2, 1))

	fmt.Println(min1(0.5, 0.7))

	type f float64
	var foo f = 23.9
	fmt.Println(min2(foo, 0.7))


	fmt.Println(min3("moi", "avec"))
}

func min[T int32 | float32](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}

func min1[T float64](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}


func min2[T ~float64](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}

type Number interface {
	int8 | int16 | int64 | int | uint | uint8 | uint16 | uint32 | uint64 | float32 | ~float64 | string
}

func min3[T Number](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}
