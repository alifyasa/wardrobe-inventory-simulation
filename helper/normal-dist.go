package helper

import "math/rand"

func IntNormalDistribution(mean int, std int) int {
	return int(rand.NormFloat64())*std + mean
}
