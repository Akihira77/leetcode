package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func hammingWeight(n int) int {
	res := 0

	for n > 0 {
		res += n % 2
		n /= 2
	}

	return res
}

func Test191_1(t *testing.T) {
	n := 11
	expected := 3
	actual := hammingWeight(n)

	assert.Equal(t, expected, actual)
}

func Test191_2(t *testing.T) {
	n := 128
	expected := 1
	actual := hammingWeight(n)

	assert.Equal(t, expected, actual)
}

func Test191_3(t *testing.T) {
	n := 2147483645
	expected := 30
	actual := hammingWeight(n)

	assert.Equal(t, expected, actual)
}
