package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)

	xStrLen := len(xStr)
	for i := range xStrLen / 2 {
		if string(xStr[i]) != string(xStr[xStrLen-i-1]) {
			return false
		}
	}

	return true
}

func Test9_1(t *testing.T) {
	x := 121
	actual := isPalindrome(x)
	expected := true

	assert.Equal(t, expected, actual)
}

func Test9_2(t *testing.T) {
	x := -121
	actual := isPalindrome(x)
	expected := false

	assert.Equal(t, expected, actual)
}

func Test9_3(t *testing.T) {
	x := 10
	actual := isPalindrome(x)
	expected := false

	assert.Equal(t, expected, actual)
}

func Test9_4(t *testing.T) {
	x := 1001
	actual := isPalindrome(x)
	expected := true

	assert.Equal(t, expected, actual)
}
