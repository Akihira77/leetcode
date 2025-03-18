package leetcode

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(s string) bool {
	s1 := ""
	s = strings.ToLower(s)
	for _, r := range s {
		if (rune(0) <= r && r <= rune(9)) || (rune('0') <= r && r <= rune('9')) || (rune('a') <= r && r <= rune('z')) {
			s1 += string(r)
		}
	}

	for i := range (len(s1) + 1) / 2 {
		if s1[i] != s1[len(s1)-i-1] {
			return false
		}
	}

	return true
}

func Test125_1(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	expected := true
	actual := isPalindrome(s)

	assert.Equal(t, expected, actual)
}

func Test125_2(t *testing.T) {
	s := "race a car"
	expected := false
	actual := isPalindrome(s)

	assert.Equal(t, expected, actual)
}

func Test125_3(t *testing.T) {
	s := ""
	expected := true
	actual := isPalindrome(s)

	assert.Equal(t, expected, actual)
}

func Test125_4(t *testing.T) {
	s := "0P"
	expected := false
	actual := isPalindrome(s)

	assert.Equal(t, expected, actual)
}
