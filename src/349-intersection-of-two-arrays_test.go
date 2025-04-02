package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	isection := make(map[int]bool, 0)

	for _, v := range nums1 {
		isection[v] = true
	}

	for _, v := range nums2 {
		if isection[v] {
			res = append(res, v)
			isection[v] = false
		}
	}

	return res
}

func Test349_1(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2}
	actual := intersection(nums1, nums2)

	assert.Equal(t, expected, actual)
}

func Test349_2(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	expected := []int{9, 4}
	actual := intersection(nums1, nums2)

	assert.Equal(t, expected, actual)
}
