package linkedlist

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	tempListSlice := []*ListNode{}
	i := 1
	current := head
	// first := &ListNode{}
	var first interface{}
	var last interface{}

	for current != nil {
		if i == left-1 {
			first = current
		}

		if i == right+1 {
			last = current
			break
		}

		if i >= left && i <= right {
			tempListSlice = append(tempListSlice, current)
		}
		i++
		current = current.Next
	}

	for j := len(tempListSlice) - 1; j >= 1; j-- {
		tempListSlice[j].Next = tempListSlice[j-1]
	}

	if v, ok := first.(*ListNode); ok {
		v.Next = tempListSlice[len(tempListSlice)-1]
	}

	if v, ok := last.(*ListNode); ok {
		tempListSlice[0].Next = v
	}

	return head
}

func reverseBetweenGPT(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	// Move to the node just before the reversal starts
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	// Reverse the sublist from left to right
	current := prev.Next
	var next *ListNode

	for i := 0; i < right-left+1; i++ {
		temp := current.Next
		current.Next = next
		next = current
		current = temp
	}

	// Connect the reversed sublist back to the original list
	prev.Next.Next = current
	prev.Next = next

	return dummy.Next
}
