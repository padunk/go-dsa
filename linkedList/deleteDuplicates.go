package linkedlist

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head, Val: -100}
	last := dummy
	first := dummy
	second := dummy.Next

	for second != nil {
		if first.Val == second.Val {
			last.Next = second.Next
			first.Next = nil
			second = second.Next
		}
		if first.Val < second.Val {
			if first.Next == nil {
				first = second
				second = second.Next
			} else {
				last = first
				first = first.Next
				second = second.Next
			}
		}

	}

	return dummy.Next
}
