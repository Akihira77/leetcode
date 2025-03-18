package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getRow(rowIndex int) []int {
	var triangle [][]int

	for i := range rowIndex + 1 {
		tempRow := make([]int, 0)
		for _ = range i + 1 {
			tempRow = append(tempRow, 1)
		}

		triangle = append(triangle, tempRow)
	}

	for i := 2; i <= rowIndex; i++ {
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle[rowIndex]

}
func Test119_1(t *testing.T) {
	rowIndex := 3
	expected := []int{1, 3, 3, 1}
	actual := getRow(rowIndex)

	assert.Equal(t, expected, actual)
}

func Test119_2(t *testing.T) {
	rowIndex := 0
	expected := []int{1}
	actual := getRow(rowIndex)

	assert.Equal(t, expected, actual)
}

func Test119_3(t *testing.T) {
	rowIndex := 1
	expected := []int{1, 1}
	actual := getRow(rowIndex)

	assert.Equal(t, expected, actual)
}
