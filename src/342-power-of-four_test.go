package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 || n%4 != 0 {
		return false
	}

	return isPowerOfFour(n / 4)
}

func Test342_1(t *testing.T) {
	n := 16
	expected := true
	actual := isPowerOfFour(n)

	assert.Equal(t, expected, actual)
}

func Test342_2(t *testing.T) {
	n := 5
	expected := false
	actual := isPowerOfFour(n)

	assert.Equal(t, expected, actual)
}

func Test342_3(t *testing.T) {
	n := 1
	expected := true
	actual := isPowerOfFour(n)

	assert.Equal(t, expected, actual)
}

func Test342_4(t *testing.T) {
	n := 1>>31 - 1
	expected := false
	actual := isPowerOfFour(n)

	assert.Equal(t, expected, actual)
}

func Test342_5(t *testing.T) {
	n := 6
	expected := false
	actual := isPowerOfFour(n)

	assert.Equal(t, expected, actual)
}

func Test342_6(t *testing.T) {
	n := 4194304
	expected := true
	actual := isPowerOfFour(n)

	assert.Equal(t, expected, actual)
}
