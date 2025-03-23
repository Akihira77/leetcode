package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}

	return n == 1
}

func Test263_1(t *testing.T) {
	n := 6
	expected := true
	actual := isUgly(n)

	assert.Equal(t, expected, actual)
}

func Test263_2(t *testing.T) {
	n := 1
	expected := true
	actual := isUgly(n)

	assert.Equal(t, expected, actual)
}

func Test263_3(t *testing.T) {
	n := 14
	expected := false
	actual := isUgly(n)

	assert.Equal(t, expected, actual)
}

func Test263_4(t *testing.T) {
	n := -2147483648
	expected := false
	actual := isUgly(n)

	assert.Equal(t, expected, actual)
}
