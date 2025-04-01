package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverseString(s []byte) {
	for i := range (len(s) + 1) / 2 {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func Test344_1(t *testing.T) {
	s := []byte("hello")
	expected := []byte("olleh")
	reverseString(s)

	assert.Equal(t, expected, []byte(s))
}
