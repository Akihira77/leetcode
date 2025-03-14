package leetcode

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(nums []int) int {
	for i := range len(nums) - 1 {
		if nums[i] == nums[i+1] {
			nums[i] = 1000
		}
	}

	res := 0
	for _, v := range nums {
		if v != 1000 {
			res++
		}
	}

	slices.Sort(nums)
	return res
}

func Test26_1(t *testing.T) {
	nums := []int{1, 1, 2}
	expected := 2
	actual := removeDuplicates(nums)

	assert.Equal(t, expected, actual)
}

func Test26_2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := 5
	actual := removeDuplicates(nums)

	assert.Equal(t, expected, actual)
}
