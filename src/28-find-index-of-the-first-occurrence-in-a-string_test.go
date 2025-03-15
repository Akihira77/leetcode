package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func strStr(haystack string, needle string) int {
	for i := range len(haystack) - len(needle) + 1 {
		if string(haystack[i:i+len(needle)]) == needle {
			return i
		}
	}

	return -1
}

func Test28_1(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	expected := 0
	actual := strStr(haystack, needle)

	assert.Equal(t, expected, actual)
}

func Test28_2(t *testing.T) {
	haystack := "leetcode"
	needle := "leeto"
	expected := -1
	actual := strStr(haystack, needle)

	assert.Equal(t, expected, actual)
}

func Test28_3(t *testing.T) {
	haystack := "a"
	needle := "a"
	expected := 0
	actual := strStr(haystack, needle)

	assert.Equal(t, expected, actual)
}

func Test28_4(t *testing.T) {
	haystack := "abc"
	needle := "c"
	expected := 2
	actual := strStr(haystack, needle)

	assert.Equal(t, expected, actual)
}
