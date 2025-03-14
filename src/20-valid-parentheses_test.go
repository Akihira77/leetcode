package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isValid(s string) bool {
	sLen := len(s)
	openBraces := make([]string, 0)

	for i := range sLen {
		currStr := string(s[i])
		if currStr == "(" || currStr == "{" || currStr == "[" {
			openBraces = append(openBraces, currStr)
		} else {
			if len(openBraces) == 0 {
				return false
			}

			lastOpenBrace := openBraces[len(openBraces)-1]
			if lastOpenBrace == "(" && currStr != ")" {
				return false
			} else if lastOpenBrace == "{" && currStr != "}" {
				return false
			} else if lastOpenBrace == "[" && currStr != "]" {
				return false
			}

			openBraces = openBraces[:len(openBraces)-1]
		}
	}

	return len(openBraces) == 0
}

func Test20_1(t *testing.T) {
	s := "()"
	expected := true
	actual := isValid(s)

	assert.Equal(t, expected, actual)
}

func Test20_2(t *testing.T) {
	s := "(){}[]"
	expected := true
	actual := isValid(s)

	assert.Equal(t, expected, actual)
}

func Test20_3(t *testing.T) {
	s := "(]"
	expected := false
	actual := isValid(s)

	assert.Equal(t, expected, actual)
}

func Test20_4(t *testing.T) {
	s := "([])"
	expected := true
	actual := isValid(s)

	assert.Equal(t, expected, actual)
}

func Test20_5(t *testing.T) {
	s := "({[}])"
	expected := false
	actual := isValid(s)

	assert.Equal(t, expected, actual)
}
