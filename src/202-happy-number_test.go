package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isHappy(n int) bool {
	freq := make(map[int]bool, 0)

	splitDigitAndSum := func(num int) int {
		sum := 0
		var currNum int

		for num > 0 {
			currNum = num % 10
			sum += (currNum) * (currNum)
			num /= 10
		}

		return sum
	}

	var sum int
	for {
		sum = splitDigitAndSum(n)
		if sum == 1 {
			return true
		}

		if freq[n] {
			return false
		}
		freq[n] = true

		n = sum
	}
}

func Test202_1(t *testing.T) {
	n := 19
	expected := true
	actual := isHappy(n)

	assert.Equal(t, expected, actual)
}

func Test202_2(t *testing.T) {
	n := 2
	expected := false
	actual := isHappy(n)

	assert.Equal(t, expected, actual)
}
