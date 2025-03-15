package leetcode

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

   100
     1

   1010
   1011
*/

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	aLen := len(a)
	bLen := len(b)
	res := ""
	carry := 0

	for i := 1; i <= bLen; i++ {
		intA, _ := strconv.Atoi(string(a[aLen-i]))
		intB, _ := strconv.Atoi(string(b[bLen-i]))

		intA = (intA + intB + carry)
		carry = intA / 2
		intA %= 2
		res = strconv.Itoa(intA) + res
	}

	for i := bLen + 1; i <= aLen; i++ {
		intA, _ := strconv.Atoi(string(a[aLen-i]))

		intA = (intA + carry)
		carry = intA / 2
		intA %= 2
		res = strconv.Itoa(intA) + res
	}

	if carry == 1 {
		res = strconv.Itoa(carry) + res
	}

	return res
}

func Test67_1(t *testing.T) {
	a := "11"
	b := "1"
	expected := "100"
	actual := addBinary(a, b)

	assert.Equal(t, expected, actual)
}

func Test67_2(t *testing.T) {
	a := "100"
	b := "1"
	expected := "101"
	actual := addBinary(a, b)

	assert.Equal(t, expected, actual)
}

func Test67_3(t *testing.T) {
	a := "001"
	b := "1111"
	expected := "10000"
	actual := addBinary(a, b)

	assert.Equal(t, expected, actual)
}
