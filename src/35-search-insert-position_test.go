package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	res := 0

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			res = mid
			break
		} else if nums[mid] > target {
			if mid == 0 {
				res = 0
				break
			}
			if mid > 0 && nums[mid-1] < target {
				res = mid
				break
			}

			r = mid - 1
		} else {
			if mid == len(nums)-1 {
				res = len(nums)
				break
			} else if mid+1 < len(nums) && nums[mid+1] >= target {
				res = mid + 1
				break
			}

			l = mid + 1
		}
	}

	return res
}

func Test35_1(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	expected := 2
	actual := searchInsert(nums, target)

	assert.Equal(t, expected, actual)
}

func Test35_2(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 2
	expected := 1
	actual := searchInsert(nums, target)

	assert.Equal(t, expected, actual)
}

func Test35_3(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7
	expected := 4
	actual := searchInsert(nums, target)

	assert.Equal(t, expected, actual)
}

func Test35_4(t *testing.T) {
	nums := []int{1, 3}
	target := 2
	expected := 1
	actual := searchInsert(nums, target)

	assert.Equal(t, expected, actual)
}

func Test35_5(t *testing.T) {
	nums := []int{1, 3}
	target := 1
	expected := 0
	actual := searchInsert(nums, target)

	assert.Equal(t, expected, actual)
}
