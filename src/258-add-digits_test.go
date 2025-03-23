package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func addDigits(num int) int {
	temp := 0

	for {
		temp += num % 10
		if temp >= 10 {
			temp = (temp / 10) + (temp % 10)
		}

		num /= 10

		if num == 0 && temp >= 10 {
			num = temp
		} else if num == 0 && temp <= 10 {
			break
		}
	}

	return temp
}

func Test258_1(t *testing.T) {
	num := 38
	expected := 2
	actual := addDigits(num)

	assert.Equal(t, expected, actual)
}

func Test258_2(t *testing.T) {
	num := 576
	expected := 9
	actual := addDigits(num)

	assert.Equal(t, expected, actual)
}

func Test258_3(t *testing.T) {
	num := 56789
	expected := 8
	actual := addDigits(num)

	assert.Equal(t, expected, actual)
}

func Test258_4(t *testing.T) {
	num := 0
	expected := 0
	actual := addDigits(num)

	assert.Equal(t, expected, actual)
}
