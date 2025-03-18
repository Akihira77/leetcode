package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}

	return nums[0]
}

func Test136_1(t *testing.T) {
	nums := []int{2, 2, 1}
	expected := 1
	actual := singleNumber(nums)

	assert.Equal(t, expected, actual)
}

func Test136_2(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	expected := 4
	actual := singleNumber(nums)

	assert.Equal(t, expected, actual)
}

func Test136_3(t *testing.T) {
	nums := []int{1}
	expected := 1
	actual := singleNumber(nums)

	assert.Equal(t, expected, actual)
}
