package caesar

import (
	"slices"
)

var Alphabet = []rune("abcdefghijklmnopqrstuvwxyzæøåABCDEFGHIJKLMNOPQRSTUVWXYZÆØÅ")

// positive remainder modulo
func mod(a, b int) int {
	return ((a % b) + b) % b
}

func shift(input string, n int) string {
	runes := []rune(input)
	for i, ch := range runes {
		if index := slices.Index(Alphabet, ch); index >= 0 {
			runes[i] = Alphabet[mod(index+n, len(Alphabet))]
		}
	}
	return string(runes)
}

func Encrypt(input string, key int) string {
	return shift(input, key)
}

func Decrypt(input string, key int) string {
	return shift(input, -key)
}
