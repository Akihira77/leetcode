package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countBits(n int) []int {
	res := make([]int, 0)

	for i := range n + 1 {
		ones := 0

		for i != 0 {
			if i&1 == 1 {
				ones++
			}

			i /= 2
		}

		res = append(res, ones)
	}

	return res
}

func Test338_1(t *testing.T) {
	n := 2
	expected := []int{0, 1, 1}
	actual := countBits(n)

	assert.Equal(t, expected, actual)
}

func Test338_2(t *testing.T) {
	n := 5
	expected := []int{0, 1, 1, 2, 1, 2}
	actual := countBits(n)

	assert.Equal(t, expected, actual)
}
