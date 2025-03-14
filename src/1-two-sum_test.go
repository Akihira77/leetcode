package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	candidates := make(map[int]int)

	for i, v := range nums {
		candidates[target-v] = i
	}

	res := []int{}
	for i, a := range nums {
		otherI := candidates[a]
		b := nums[otherI]
		if a+b == target && i != otherI {
			res = []int{i, otherI}
			break
		}
	}

	return res
}

func Test1_1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	actual := twoSum(nums, target)
	expected := []int{0, 1}

	assert.Equal(t, expected, actual)
}

func Test1_2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	actual := twoSum(nums, target)
	expected := []int{1, 2}

	assert.Equal(t, expected, actual)
}

func Test1_3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	actual := twoSum(nums, target)
	expected := []int{0, 1}

	assert.Equal(t, expected, actual)
}
