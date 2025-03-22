package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	arr := make([][]string, 0)

	curr := []string{strconv.Itoa(nums[0])}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			curr = append(curr, strconv.Itoa(nums[i]))
		} else {
			arr = append(arr, curr)
			curr = []string{strconv.Itoa(nums[i])}
		}
	}

	if len(curr) > 0 {
		arr = append(arr, curr)
		curr = []string{}
	}

	for _, v := range arr {
		if len(v) > 1 {
			curr = append(curr, v[0]+"->"+v[len(v)-1])
		} else {
			curr = append(curr, v[0])
		}
	}

	return curr
}

func Test228_1(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7}
	expected := []string{
		"0->2",
		"4->5",
		"7",
	}
	actual := summaryRanges(nums)

	assert.Equal(t, expected, actual)
}

func Test228_2(t *testing.T) {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	expected := []string{
		"0",
		"2->4",
		"6",
		"8->9",
	}
	actual := summaryRanges(nums)

	assert.Equal(t, expected, actual)
}
