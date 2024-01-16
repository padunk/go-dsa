package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	listMap := make(map[int]*ListNode)
	length := 0

	for current != nil {
		length++
		listMap[length] = current
		current = current.Next
	}

	if length == 1 && n == 1 {
		return nil
	}

	indexToDelete := length - n + 1
	listBefore := listMap[indexToDelete-1]
	listAfter := listMap[indexToDelete+1]
	if listBefore != nil && listAfter != nil {
		listBefore.Next = listAfter
	}
	if listAfter == nil {
		listBefore.Next = nil
	}

	if indexToDelete == 1 {
		return head.Next
	}

	listMap[indexToDelete].Next = nil

	return head
}

func removeNthFromEndGPT(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy

	// Move the second pointer to the (n+1)th node from the beginning
	for i := 1; i <= n+1; i++ {
		second = second.Next
	}

	// Move both pointers until the second pointer reaches the end
	for second != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the nth node
	first.Next = first.Next.Next

	return dummy.Next
}

// other solution
func removeNthFromEndOther(head *ListNode, n int) *ListNode {
	var lagP *ListNode = nil
	startNode := head
	for cur := 1; head != nil; cur++ {
		if lagP != nil {
			lagP = lagP.Next
		} else if lagP == nil && cur > n {
			lagP = startNode
		}
		head = head.Next
	}
	if lagP == nil {
		startNode = startNode.Next
	} else if lagP != nil && lagP.Next != nil {
		lagP.Next = lagP.Next.Next
	}
	return startNode
}
