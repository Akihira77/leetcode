package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = 0

		if digits[i] > 9 {
			carry = 1
			digits[i] %= 10
			continue
		}

	}

	if carry == 1 {
		digits = append(digits, 0)
		copy(digits[1:], digits)
		digits[0] = 1
	}

	return digits
}

func Test66_1(t *testing.T) {
	digits := []int{1, 2, 3}
	expected := []int{1, 2, 4}
	actual := plusOne(digits)

	assert.Equal(t, expected, actual)
}

func Test66_2(t *testing.T) {
	digits := []int{4, 3, 2, 1}
	expected := []int{4, 3, 2, 2}
	actual := plusOne(digits)

	assert.Equal(t, expected, actual)
}

func Test66_3(t *testing.T) {
	digits := []int{9}
	expected := []int{1, 0}
	actual := plusOne(digits)

	assert.Equal(t, expected, actual)
}

func Test66_4(t *testing.T) {
	digits := []int{9, 9}
	expected := []int{1, 0, 0}
	actual := plusOne(digits)

	assert.Equal(t, expected, actual)
}

func Test66_5(t *testing.T) {
	digits := []int{8, 9, 9}
	expected := []int{9, 0, 0}
	actual := plusOne(digits)

	assert.Equal(t, expected, actual)
}
