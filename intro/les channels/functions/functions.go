package functions

import "math/rand"

func GenerateRandomNumber(n int) int  {
	randomNumber := rand.Intn(n)
	return randomNumber
}
