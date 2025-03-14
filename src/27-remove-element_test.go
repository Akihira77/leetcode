package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"slices"
)

func removeElement(nums []int, val int) int {
	for i := range len(nums) {
		if nums[i] == val {
			nums[i] = 1000
		}
	}

	res := len(nums)
	i := 0
	for {
		if i >= len(nums) {
			return res
		}

		if nums[i] == 1000 {
			nums = slices.Delete(nums, i, i+1)
			res--
			continue
		}
		i++
	}
}

func Test27_1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	expectedNums := []int{2, 2, 0, 0}
	val := 3
	expected := 2
	actual := removeElement(nums, val)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedNums, nums)
}

func Test27_2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	expectedNums := []int{0, 1, 3, 0, 4, 0, 0, 0}
	val := 2
	expected := 5
	actual := removeElement(nums, val)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedNums, nums)
}
