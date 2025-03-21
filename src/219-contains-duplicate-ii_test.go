package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	freq := make(map[int]int, 0)

	for i, v := range nums {
		num, ok := freq[v]
		if ok && i-num <= k {
			return true
		}
		freq[v] = i
	}

	return false
}

func Test219_1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	expected := true
	actual := containsNearbyDuplicate(nums, k)

	assert.Equal(t, expected, actual)
}

func Test219_2(t *testing.T) {
	nums := []int{1, 0, 1, 1}
	k := 1
	expected := true
	actual := containsNearbyDuplicate(nums, k)

	assert.Equal(t, expected, actual)
}

func Test219_3(t *testing.T) {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	expected := false
	actual := containsNearbyDuplicate(nums, k)

	assert.Equal(t, expected, actual)
}
