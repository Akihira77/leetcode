package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func moveZeroes(nums []int) {
	counter := 0
	for i := range nums {
		if nums[i] != 0 {
			counter++
		}
	}

	for i := 0; i < len(nums); {
		if len(nums[:i]) == counter {
			break
		}

		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		} else {
			i++
		}
	}
}

func Test283_1(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}
	moveZeroes(nums)

	assert.Equal(t, expected, nums)
}

func Test283_2(t *testing.T) {
	nums := []int{0}
	expected := []int{0}
	moveZeroes(nums)

	assert.Equal(t, expected, nums)
}

func Test283_3(t *testing.T) {
	nums := []int{0, 0, 1}
	expected := []int{1, 0, 0}
	moveZeroes(nums)

	assert.Equal(t, expected, nums)
}
