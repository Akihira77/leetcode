package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestCommonPrefix(strs []string) string {
	idx := 0
	res := ""

Loop:
	for {
		v := strs[0]
		for i, v1 := range strs {
			if len(v) == 0 || len(v1) == 0 {
				break Loop
			}

			if len(v) <= idx || len(v1) <= idx {
				break Loop
			}

			if string(v[idx]) != string(v1[idx]) {
				break Loop
			}

			if i+1 == len(strs) {
				res += string(v[len(res)])
			}
		}

		idx++
	}

	return res
}

func Test14_1(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	actual := longestCommonPrefix(strs)
	expected := "fl"

	assert.Equal(t, expected, actual)
}

func Test14_2(t *testing.T) {
	strs := []string{"dog", "racecar", "car"}
	res := longestCommonPrefix(strs)
	expected := ""

	assert.Equal(t, expected, res)
}

func Test14_3(t *testing.T) {
	strs := []string{"flower", "flower", "flower"}
	res := longestCommonPrefix(strs)
	expected := "flower"

	assert.Equal(t, expected, res)
}
