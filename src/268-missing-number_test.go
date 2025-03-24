package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func missingNumber(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return (len(nums) * (len(nums) + 1) / 2) - sum
}

func Test268_1(t *testing.T) {
	nums := []int{3, 0, 1}
	expected := 2
	actual := missingNumber(nums)

	assert.Equal(t, expected, actual)
}

func Test268_2(t *testing.T) {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expected := 8
	actual := missingNumber(nums)

	assert.Equal(t, expected, actual)
}

func Test268_3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := 0
	actual := missingNumber(nums)

	assert.Equal(t, expected, actual)
}
