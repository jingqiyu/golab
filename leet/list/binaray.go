package list

import (
	"math"
)

func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	cur := head
	var prev *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	var ret int
	var i int

	for prev != nil {
		ret += int(math.Pow(2, float64(i))) * prev.Val
		i++

		prev = prev.Next
	}
	return ret
}
