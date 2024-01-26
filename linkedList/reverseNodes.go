package linkedlist

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// Calculate the length of the linked list
	length := 0
	for head != nil {
		length++
		head = head.Next
	}

	// Reverse k nodes at a time
	for length >= k {
		current := prev.Next
		next := current.Next

		for i := 1; i < k; i++ {
			current.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
			next = current.Next
		}

		prev = current
		length -= k
	}

	return dummy.Next
}
