package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func convertToTitle(columnNumber int) string {
	res := ""
	startChar := int('A') - 1

	for columnNumber > 0 {
		if columnNumber%26 == 0 {
			res = string(rune(startChar+26)) + res
		} else {
			res = string(rune(startChar+columnNumber%26)) + res
		}

		columnNumber = (columnNumber - 1) / 26
	}

	return res
}

func Test168_1(t *testing.T) {
	columnNumber := 1
	expected := "A"
	actual := convertToTitle(columnNumber)

	assert.Equal(t, expected, actual)
}

func Test168_2(t *testing.T) {
	columnNumber := 28
	expected := "AB"
	actual := convertToTitle(columnNumber)

	assert.Equal(t, expected, actual)
}

func Test168_3(t *testing.T) {
	columnNumber := 701
	expected := "ZY"
	actual := convertToTitle(columnNumber)

	assert.Equal(t, expected, actual)
}

func Test168_4(t *testing.T) {
	columnNumber := 52
	expected := "AZ"
	actual := convertToTitle(columnNumber)

	assert.Equal(t, expected, actual)
}
