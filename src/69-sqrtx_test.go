package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mySqrt(x int) int {
	if x == 2 || x == 3 {
		return 1
	}

	if x == 5 || x == 7 {
		return 2
	}

	for xi := range x/2 + 1 {
		if xi*xi > x {
			x = xi - 1
			break
		}

		if xi*xi == x {
			x = xi
			break
		}
	}

	return x
}

func Test69_1(t *testing.T) {
	x := 4
	expected := 2
	actual := mySqrt(x)

	assert.Equal(t, expected, actual)
}

func Test69_2(t *testing.T) {
	x := 8
	expected := 2
	actual := mySqrt(x)

	assert.Equal(t, expected, actual)
}

func Test69_3(t *testing.T) {
	x := 25
	expected := 5
	actual := mySqrt(x)

	assert.Equal(t, expected, actual)
}

func Test69_4(t *testing.T) {
	x := 120
	expected := 10
	actual := mySqrt(x)

	assert.Equal(t, expected, actual)
}

func Test69_5(t *testing.T) {
	x := 1<<31 - 1 // 2147483647
	expected := 46340
	fmt.Println(x, " & ", expected)
	actual := mySqrt(x)

	assert.Equal(t, expected, actual)
}

func Test69_6(t *testing.T) {
	x := 0
	expected := 0
	actual := mySqrt(x)

	assert.Equal(t, expected, actual)
}

func Test69_7(t *testing.T) {
	x := 2
	expected := 1
	actual := mySqrt(x)

	assert.Equal(t, expected, actual)
}
