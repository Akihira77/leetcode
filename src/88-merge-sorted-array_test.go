package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//NOTE: Naive solution because constraint is small
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	nums1 = append(nums1[:m], nums2...)
//
// 	for i := range len(nums1) {
// 		for j := i + 1; j < len(nums1); j++ {
// 			if nums1[i] >= nums1[j] {
// 				nums1[i], nums1[j] = nums1[j], nums1[i]
// 			}
// 		}
// 	}
// }

// NOTE: TimeComplexity O(max(m, n)) solution but MemoryComplexity is O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, 0)
	nums1 = nums1[:m]

	currNum1Idx := 0
	for i := 0; i < n; {
		if currNum1Idx >= m || nums1[currNum1Idx] > nums2[i] {
			temp = append(temp, nums2[i])
			i++
		} else if nums1[currNum1Idx] == nums2[i] {
			temp = append(temp, nums1[currNum1Idx], nums2[i])
			currNum1Idx++
			i++
		} else {
			temp = append(temp, nums1[currNum1Idx])
			currNum1Idx++
		}
	}

	for currNum1Idx < m {
		temp = append(temp, nums1[currNum1Idx])
		currNum1Idx++
	}

	nums1 = append(nums1[:0], temp...)
}

func Test88_1(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	expected := []int{1, 2, 2, 3, 5, 6}
	merge(nums1, m, nums2, n)

	assert.Equal(t, expected, nums1)
}

func Test88_2(t *testing.T) {
	nums1 := []int{1}
	m := 1
	nums2 := []int{}
	n := 0
	expected := []int{1}
	merge(nums1, m, nums2, n)

	assert.Equal(t, expected, nums1)
}

func Test88_3(t *testing.T) {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	expected := []int{1}
	merge(nums1, m, nums2, n)

	assert.Equal(t, expected, nums1)
}

func Test88_4(t *testing.T) {
	nums1 := []int{2, 0}
	m := 1
	nums2 := []int{1}
	n := 1
	expected := []int{1, 2}
	merge(nums1, m, nums2, n)

	assert.Equal(t, expected, nums1)
}
