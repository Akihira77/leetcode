package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func generate(numRows int) [][]int {
	var triangle [][]int

	for i := range numRows {
		tempRow := make([]int, 0)
		for _ = range i + 1 {
			tempRow = append(tempRow, 1)
		}

		triangle = append(triangle, tempRow)
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func Test118_1(t *testing.T) {
	numRows := 5
	expected := [][]int{
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
	}
	actual := generate(numRows)

	assert.Equal(t, expected, actual)
}

func Test118_2(t *testing.T) {
	numRows := 1
	expected := [][]int{
		[]int{1},
	}
	actual := generate(numRows)

	assert.Equal(t, expected, actual)
}
