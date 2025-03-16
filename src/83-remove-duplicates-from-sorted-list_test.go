package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	res := &ListNode{
		Val:  nums[len(nums)-1],
		Next: nil,
	}

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != res.Val {
			temp := &ListNode{
				Val:  nums[i],
				Next: res,
			}

			res = temp
		}
	}

	return res
}

func Test83_1(t *testing.T) {
	head2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	head1 := &ListNode{
		Val:  1,
		Next: head2,
	}
	head0 := &ListNode{
		Val:  1,
		Next: head1,
	}

	expected1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	expected0 := &ListNode{
		Val:  1,
		Next: expected1,
	}

	actual := deleteDuplicates(head0)

	assert.Equal(t, expected0, actual)
}

func Test83_2(t *testing.T) {
	head4 := &ListNode{
		Val:  3,
		Next: nil,
	}
	head3 := &ListNode{
		Val:  3,
		Next: head4,
	}
	head2 := &ListNode{
		Val:  2,
		Next: head3,
	}
	head1 := &ListNode{
		Val:  1,
		Next: head2,
	}
	head0 := &ListNode{
		Val:  1,
		Next: head1,
	}

	expected2 := &ListNode{
		Val:  3,
		Next: nil,
	}
	expected1 := &ListNode{
		Val:  2,
		Next: expected2,
	}
	expected0 := &ListNode{
		Val:  1,
		Next: expected1,
	}

	actual := deleteDuplicates(head0)

	assert.Equal(t, expected0, actual)
}
