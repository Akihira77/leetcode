package leetcode

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPerfectSquare(num int) bool {
	sqrt := math.Sqrt(float64(num))
	return sqrt*sqrt == float64(num) && sqrt == float64(int(sqrt))
}

func Test367_1(t *testing.T) {
	num := 16
	expected := true
	actual := isPerfectSquare(num)

	assert.Equal(t, expected, actual)
}

func Test367_2(t *testing.T) {
	num := 14
	expected := false
	actual := isPerfectSquare(num)

	assert.Equal(t, expected, actual)
}
