package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type NumArray struct {
	Prefix []int
}

func Constructor(nums []int) NumArray {
	var na NumArray
	na.Prefix = make([]int, len(nums))
	na.Prefix[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		na.Prefix[i] = na.Prefix[i-1] + nums[i]
	}

	return na
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.Prefix[right]
	}

	return this.Prefix[right] - this.Prefix[left-1]
}

func Test303_1(t *testing.T) {
	nums := Constructor([]int{-2, 0, 3, -5, 2, -1})
	actual := make([]int, 0)
	actual = append(actual, nums.SumRange(0, 2))
	actual = append(actual, nums.SumRange(2, 5))
	actual = append(actual, nums.SumRange(0, 5))

	expected := []int{1, -1, -3}

	assert.Equal(t, expected, actual)
}
