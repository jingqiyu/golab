package list

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	var prev *ListNode
	var tmp *ListNode
	var cur = head
	for cur != nil {
		tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

//倒数第K
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	p := head
	q := head
	i := k
	for p != nil {
		if i == 0 {
			break
		}
		p = p.Next
		i--
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	p := l3
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else if l2 == nil {
		p.Next = l1
	}
	return l3.Next
}
