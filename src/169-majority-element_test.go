package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func majorityElement(nums []int) int {
	freq := make(map[int]int, 0)

	freqMust := len(nums) / 2

	for i := range len(nums)/2 + 1 {
		freq[nums[i]]++
		if freq[nums[i]] > freqMust {
			return nums[i]
		}

		freq[nums[len(nums)-i-1]]++
		if freq[nums[len(nums)-i-1]] > freqMust {
			return nums[len(nums)-i-1]
		}
	}

	return nums[0]
}

func Test169_1(t *testing.T) {
	nums := []int{3, 2, 3}
	expected := 3
	actual := majorityElement(nums)

	assert.Equal(t, expected, actual)
}

func Test169_2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	expected := 2
	actual := majorityElement(nums)

	assert.Equal(t, expected, actual)
}
