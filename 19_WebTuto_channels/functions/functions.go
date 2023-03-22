package functions

import "math/rand"

func GenerateRandomNumber(n int) int {
	randomNumer := rand.Intn(n)
	return randomNumer
}
