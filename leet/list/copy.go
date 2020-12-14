package list

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	newH := &Node{Val: 0}

	p := head
	q := newH

	var t = make(map[*Node]*Node)

	for p != nil {
		q.Next = &Node{Val: p.Val}
		t[p] = q.Next
		q = q.Next
		p = p.Next
	}

	p = head
	q = newH.Next

	for p != nil {
		tmp, _ := t[p.Random]
		q.Random = tmp
		p = p.Next
		q = q.Next
	}
	return newH.Next

}
