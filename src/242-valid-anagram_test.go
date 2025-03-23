package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[rune]int, 0)

	for i, c := range s {
		freq[c]++
		freq[rune(t[i])]--
	}

	for _, c := range t {
		if freq[c] != 0 {
			return false
		}
	}

	return true
}

func Test242_1(t *testing.T) {
	s := "anagram"
	t1 := "nagaram"
	expected := true
	actual := isAnagram(s, t1)

	assert.Equal(t, expected, actual)
}

func Test242_2(t *testing.T) {
	s := "rat"
	t1 := "car"
	expected := false
	actual := isAnagram(s, t1)

	assert.Equal(t, expected, actual)
}
