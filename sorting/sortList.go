package sorting

import (
	"cmp"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var lists []*ListNode
	temp := head
	for temp != nil {
		lists = append(lists, temp)
		temp = temp.Next
	}

	// You need to sort the lists slice here based on the Val of ListNode
	slices.SortFunc(lists, func(i, j *ListNode) int {
		return cmp.Compare(i.Val, j.Val)
	})

	for i := 0; i < len(lists)-1; i++ {
		lists[i].Next = lists[i+1]
	}

	// Make sure the last node points to nil
	lists[len(lists)-1].Next = nil

	return lists[0]
}

func sortListGPT(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// split list into two halves
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// sort each half
	half := slow.Next
	slow.Next = nil
	l1 := sortList(head)
	l2 := sortList(half)
	// merge sorted halves
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummy.Next
}
