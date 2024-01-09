package helpers

import (
	"math/rand"
)

func RandomNumberGenerator(length, count int) []int {
	if count > length {
		count = length
	}

	allIndexes := make([]int, length)

	for i := 0; i < length; i++ {
		allIndexes[i] = i
	}

	for i := length - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		allIndexes[i], allIndexes[j] = allIndexes[j], allIndexes[i]
	}

	return allIndexes[:count]
}
