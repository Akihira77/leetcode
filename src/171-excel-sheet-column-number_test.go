package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func titleToNumber(columnTitle string) int {
	res := 0

	pow := func(num, power int) int {
		for power > 1 {
			num *= 26
			power--
		}

		return num
	}

	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += pow(int(columnTitle[i]-'A')+1, len(columnTitle)-i)
	}

	return res
}

func Test171_1(t *testing.T) {
	columnTitle := "A"
	expected := 1
	actual := titleToNumber(columnTitle)

	assert.Equal(t, expected, actual)
}

func Test171_2(t *testing.T) {
	columnTitle := "AB"
	expected := 28
	actual := titleToNumber(columnTitle)

	assert.Equal(t, expected, actual)
}

func Test171_3(t *testing.T) {
	columnTitle := "ZY"
	expected := 701
	actual := titleToNumber(columnTitle)

	assert.Equal(t, expected, actual)
}

func Test171_4(t *testing.T) {
	columnTitle := "AAA"
	expected := 703
	actual := titleToNumber(columnTitle)

	assert.Equal(t, expected, actual)
}
