package randomizer

import (
	"math/rand"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Ints(n int, maxExclusive int) []int {
	randomInts := []int{}
	for i := 0; i < n; i++ {
		randomInts = append(randomInts, rand.Intn(maxExclusive))
	}
	return randomInts
}

func Strings(n int, length int) []string {
	randomStrings := []string{}
	for i := 0; i < n; i++ {
		randomStrings = append(randomStrings, randomString(length))
	}
	return randomStrings
}

func randomString(length int) string {
	randomStr := make([]byte, length)
	for i := range randomStr {
		randomStr[i] = letters[rand.Intn(len(letters))]
	}
	return string(randomStr)
}
