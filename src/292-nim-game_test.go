package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canWinNim(n int) bool {
	return n%4 != 0
}

func Test292_1(t *testing.T) {
	n := 4
	expected := false
	actual := canWinNim(n)

	assert.Equal(t, expected, actual)
}

func Test292_2(t *testing.T) {
	n := 2
	expected := true
	actual := canWinNim(n)

	assert.Equal(t, expected, actual)
}

func Test292_3(t *testing.T) {
	n := 1
	expected := true
	actual := canWinNim(n)

	assert.Equal(t, expected, actual)
}

func Test292_4(t *testing.T) {
	n := (1 >> 31) - 1
	expected := true
	actual := canWinNim(n)

	assert.Equal(t, expected, actual)
}
