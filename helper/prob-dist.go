package helper

import (
	"math/rand"
)

func IntExponentialDistribution(mean int, min int) int {
	randomNumber := int(rand.ExpFloat64() * float64(mean))
	if randomNumber < min {
		return min
	}
	return randomNumber
}
