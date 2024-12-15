package helper

import (
	"math/rand"
)

func IntExponentialDistribution(mean int, min int, max int) int {
	randomNumber := int(rand.ExpFloat64() * float64(mean))
	if randomNumber < min {
		return min
	}
	if randomNumber > max {
		return max
	}
	return randomNumber
}
