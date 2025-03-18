package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit(prices []int) int {
	minPrice := prices[0]
	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > res {
			res = prices[i] - minPrice
		}
	}

	return res
}

func Test121_1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5
	actual := maxProfit(prices)

	assert.Equal(t, expected, actual)
}

func Test121_2(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	expected := 0
	actual := maxProfit(prices)

	assert.Equal(t, expected, actual)
}

func Test121_3(t *testing.T) {
	prices := []int{1, 2, 3, 4, 5, 6}
	expected := 5
	actual := maxProfit(prices)

	assert.Equal(t, expected, actual)
}
