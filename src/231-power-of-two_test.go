package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n&1 != 0 {
			return false
		}

		n /= 2
	}

	return true
}

func Test231_1(t *testing.T) {
	n := 1
	expected := true
	actual := isPowerOfTwo(n)

	assert.Equal(t, expected, actual)
}

func Test231_2(t *testing.T) {
	n := 4
	expected := true
	actual := isPowerOfTwo(n)

	assert.Equal(t, expected, actual)
}

func Test231_3(t *testing.T) {
	n := 256
	expected := true
	actual := isPowerOfTwo(n)

	assert.Equal(t, expected, actual)
}

func Test231_4(t *testing.T) {
	n := 768
	expected := false
	actual := isPowerOfTwo(n)

	assert.Equal(t, expected, actual)
}
