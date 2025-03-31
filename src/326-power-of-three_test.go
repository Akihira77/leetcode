package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 || n%3 != 0 {
		return false
	}

	return isPowerOfThree(n / 3)
}

func Test326_1(t *testing.T) {
	n := 27
	expected := true
	actual := isPowerOfThree(n)

	assert.Equal(t, expected, actual)
}

func Test326_2(t *testing.T) {
	n := 0
	expected := false
	actual := isPowerOfThree(n)

	assert.Equal(t, expected, actual)
}

func Test326_3(t *testing.T) {
	n := -1
	expected := false
	actual := isPowerOfThree(n)

	assert.Equal(t, expected, actual)
}

func Test326_4(t *testing.T) {
	n := 1>>31 - 1
	expected := false
	actual := isPowerOfThree(n)

	assert.Equal(t, expected, actual)
}

func Test326_5(t *testing.T) {
	n := 45
	expected := false
	actual := isPowerOfThree(n)

	assert.Equal(t, expected, actual)
}
