package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool, 0)

	for _, v := range nums[:len(nums)/2] {
		if visited[v] {
			return true
		}

		visited[v] = true
	}

	for _, v := range nums[len(nums)/2:] {
		if visited[v] {
			return true
		}

		visited[v] = true
	}

	return false
}

func Test217_1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	expected := true
	actual := containsDuplicate(nums)

	assert.Equal(t, expected, actual)
}

func Test217_2(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := false
	actual := containsDuplicate(nums)

	assert.Equal(t, expected, actual)
}

func Test217_3(t *testing.T) {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	expected := true
	actual := containsDuplicate(nums)

	assert.Equal(t, expected, actual)
}

func Test217_4(t *testing.T) {
	nums := []int{3, 3}
	expected := true
	actual := containsDuplicate(nums)

	assert.Equal(t, expected, actual)
}
