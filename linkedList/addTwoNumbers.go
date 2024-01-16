package linkedlist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultCurr := result
	current1 := l1
	current2 := l2
	additional := 0

	for current1 != nil || current2 != nil {
		sum := 0
		if current1 == nil {
			sum = current2.Val + additional
			current2 = current2.Next
		} else if current2 == nil {
			sum = current1.Val + additional
			current1 = current1.Next
		} else {
			sum = current1.Val + current2.Val + additional
			current1 = current1.Next
			current2 = current2.Next
		}

		if sum > 9 {
			additional = sum / 10
			sum %= 10
		} else {
			additional = 0
		}
		resultCurr.Val = sum

		if current1 != nil || current2 != nil {
			resultCurr.Next = &ListNode{}
			resultCurr = resultCurr.Next
		}
	}

	if additional != 0 {
		resultCurr.Next = &ListNode{Val: additional}
	}

	return result
}

// other solution
func addTwoNumbersOther(l1, l2 *ListNode) *ListNode {
	var node = &ListNode{}
	sum, carry := 0, 0
	head := node
	for l1 != nil || l2 != nil {
		sum = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		sum %= 10
		node.Next = &ListNode{Val: sum}
		node = node.Next
	}
	if carry != 0 {
		node.Next = &ListNode{Val: carry}
	}
	return head.Next
}
