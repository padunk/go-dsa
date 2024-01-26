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
			continue
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

// ChatGPT
func deleteDuplicatesGPT(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			val := current.Next.Val

			// Skip all nodes with the same value
			for current.Next != nil && current.Next.Val == val {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}
