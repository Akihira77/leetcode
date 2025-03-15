package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLastWord(s string) int {
	end := -1
	start := -1
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			if end == -1 {
				end = i
			}
		}

		if string(s[i]) == " " && end != -1 {
			start = i
			break
		}
	}

	return end - start
}

func Test58_1(t *testing.T) {
	s := "Hello World"
	expected := 5
	actual := lengthOfLastWord(s)

	assert.Equal(t, expected, actual)
}

func Test58_2(t *testing.T) {
	s := "   fly me   to   the moon  "
	expected := 4
	actual := lengthOfLastWord(s)

	assert.Equal(t, expected, actual)
}

func Test58_3(t *testing.T) {
	s := "luffy is still joyboy"
	expected := 6
	actual := lengthOfLastWord(s)

	assert.Equal(t, expected, actual)
}
