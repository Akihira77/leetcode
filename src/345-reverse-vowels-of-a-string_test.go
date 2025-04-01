package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverseVowels(s string) string {
	isVowel := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}

	runes := []byte(s)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !isVowel(runes[left]) {
			left++
		}

		for left < right && !isVowel(runes[right]) {
			right--
		}

		runes[left], runes[right] = runes[right], runes[left]

		left++
		right--
	}

	return string(runes)
}

func Test345_1(t *testing.T) {
	s := "IceCreAm"
	expected := "AceCreIm"
	actual := reverseVowels(s)

	assert.Equal(t, expected, actual)
}
