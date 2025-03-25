package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func wordPattern(pattern string, s string) bool {
	words := make(map[string]string, 0)
	words1 := make(map[string]string, 0)
	var temp string
	index := 0
	space := 0
	s += " "

	for i := range len(s) {
		if string(s[i]) == " " {
			space++
			//Check word from string s
			word, ok := words[temp]
			if index >= len(pattern) || (ok && string(word) != string(pattern[index])) {
				//If we found the current word inside map
				//then check if the pattern from that word is the same as current pattern
				return false
			} else {
				//If the word is not found
				//then check the current pattern
				word, ok = words1[string(pattern[index])]
				//If current pattern is used then check if it being mapped the same as current word
				if ok && word != temp {
					return false
				}

				//Assign current word with current pattern
				words[temp] = string(pattern[index])
				//Assign current pattern with current word
				words1[string(pattern[index])] = temp
			}
			index++
			temp = ""
		} else {
			temp += string(s[i])
		}
	}

	if space != len(pattern) {
		return false
	}

	return true
}

func Test290_1(t *testing.T) {
	pattern := "abba"
	s := "dog cat cat dog"
	expected := true
	actual := wordPattern(pattern, s)

	assert.Equal(t, expected, actual)
}

func Test290_2(t *testing.T) {
	pattern := "abba"
	s := "dog cat cat fish"
	expected := false
	actual := wordPattern(pattern, s)

	assert.Equal(t, expected, actual)
}

func Test290_3(t *testing.T) {
	pattern := "aaaa"
	s := "dog cat cat dog"
	expected := false
	actual := wordPattern(pattern, s)

	assert.Equal(t, expected, actual)
}

func Test290_4(t *testing.T) {
	pattern := "abc"
	s := "b c a"
	expected := true
	actual := wordPattern(pattern, s)

	assert.Equal(t, expected, actual)
}

func Test290_5(t *testing.T) {
	pattern := "aaa"
	s := "aa aa aa aa"
	expected := false
	actual := wordPattern(pattern, s)

	assert.Equal(t, expected, actual)
}

func Test290_6(t *testing.T) {
	pattern := "jquery"
	s := "jquery"
	expected := false
	actual := wordPattern(pattern, s)

	assert.Equal(t, expected, actual)
}
