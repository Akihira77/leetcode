package leetcode

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(lista *ListNode, listb *ListNode) *ListNode {
	temp := make([]int, 0)
	var scanB bool
	var currNode *ListNode

	if lista != nil {
		scanB = false
		currNode = lista
	} else if listb != nil {
		scanB = true
		currNode = listb
	} else {
		return nil
	}

	for {
		if currNode == nil && scanB {
			break
		}

		temp = append(temp, currNode.Val)
		if currNode.Next == nil {
			if listb == nil {
				break
			}

			if currNode.Next == nil && scanB {
				break
			}

			currNode = listb
			temp = append(temp, currNode.Val)
			scanB = true
		}
		fmt.Println(temp)

		currNode = currNode.Next
	}

	tempLen := len(temp)
	fmt.Println(temp)
	if tempLen == 0 {
		return &ListNode{}
	}

	slices.Sort(temp)

	nextNode := &ListNode{
		Val:  temp[tempLen-1],
		Next: nil,
	}

	for i := tempLen - 2; i >= 0; i-- {
		currNode := &ListNode{
			Val:  temp[i],
			Next: nextNode,
		}

		nextNode = currNode
	}

	return nextNode
}

func Test21_1(t *testing.T) {
	lista_3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	lista_2 := &ListNode{
		Val:  2,
		Next: lista_3,
	}
	lista_1 := &ListNode{
		Val:  1,
		Next: lista_2,
	}

	listb_3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	listb_2 := &ListNode{
		Val:  3,
		Next: listb_3,
	}
	listb_1 := &ListNode{
		Val:  1,
		Next: listb_2,
	}

	listc_6 := &ListNode{
		Val:  4,
		Next: nil,
	}
	listc_5 := &ListNode{
		Val:  4,
		Next: listc_6,
	}
	listc_4 := &ListNode{
		Val:  3,
		Next: listc_5,
	}
	listc_3 := &ListNode{
		Val:  2,
		Next: listc_4,
	}
	listc_2 := &ListNode{
		Val:  1,
		Next: listc_3,
	}
	listc_1 := &ListNode{
		Val:  1,
		Next: listc_2,
	}

	res := mergeTwoLists(lista_1, listb_1)
	assert.Equal(t, res, listc_1)
}

func Test21_2(t *testing.T) {
	var lista_1 *ListNode
	var listb_1 *ListNode
	var listc_1 *ListNode

	res := mergeTwoLists(lista_1, listb_1)
	assert.Equal(t, res, listc_1)
}

func Test21_3(t *testing.T) {
	var lista_1 *ListNode
	listb_1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	listc_1 := &ListNode{
		Val:  0,
		Next: nil,
	}

	res := mergeTwoLists(lista_1, listb_1)
	assert.Equal(t, res, listc_1)
}

func Test21_4(t *testing.T) {
	var lista_1 *ListNode
	var listb_1 *ListNode
	var listc_1 *ListNode

	res := mergeTwoLists(lista_1, listb_1)
	assert.Equal(t, res, listc_1)
}

func Test21_5(t *testing.T) {
	lista_1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	var listb_1 *ListNode
	listc_1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	res := mergeTwoLists(lista_1, listb_1)
	assert.Equal(t, res, listc_1)
}
