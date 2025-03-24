package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
**/
var bad int

func isBadVersion(version int) bool {
	return version >= bad
}

func firstBadVersion(n int) int {
	l := 1
	r := n

	for l <= r {
		n = l + (r-l)/2

		if isBadVersion(n) {
			if !isBadVersion(n - 1) {
				break
			}
			r = n - 1
		} else {
			l = n + 1
		}
	}

	return n
}

func Test278_1(t *testing.T) {
	n := 5
	bad = 4
	expected := 4
	actual := firstBadVersion(n)

	assert.Equal(t, expected, actual)
}

func Test278_2(t *testing.T) {
	n := (1 << 31) - 1
	bad = 1000_005
	expected := 1000_005
	actual := firstBadVersion(n)

	assert.Equal(t, expected, actual)
}

func Test278_3(t *testing.T) {
	n := 3
	bad = 1
	expected := 1
	actual := firstBadVersion(n)

	assert.Equal(t, expected, actual)
}

func Test278_4(t *testing.T) {
	n := 2126753390
	bad = 1702766719
	expected := 1702766719
	actual := firstBadVersion(n)

	assert.Equal(t, expected, actual)
}

func Test278_5(t *testing.T) {
	n := 3
	bad = 2
	expected := 2
	actual := firstBadVersion(n)

	assert.Equal(t, expected, actual)
}

func Test278_6(t *testing.T) {
	n := 5
	bad = 5
	expected := 5
	actual := firstBadVersion(n)

	assert.Equal(t, expected, actual)
}
