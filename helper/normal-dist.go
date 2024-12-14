package helper

import (
	"math"
	"math/rand"
)

func NonNegativeIntNormalDistribution(mean int, std int) int {
	result := int(rand.NormFloat64())*std + mean
	return int(math.Max(float64(result), 0))
}
