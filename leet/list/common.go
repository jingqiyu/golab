package list

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func GenList(src ...int) *ListNode {
	if len(src) == 0 {
		return nil
	}

	head := &ListNode{
		Val: 0,
	}
	var p = head
	for _, v := range src {
		p.Next = &ListNode{Val: v, Next: nil}
		p = p.Next
	}
	return head.Next
}

func PrintList(l *ListNode) {
	var bf bytes.Buffer

	for l != nil {
		bf.WriteString(strconv.Itoa(l.Val))
		bf.WriteString("->")
		l = l.Next
	}
	bf.WriteString("nil")

	fmt.Println(bf.String())

}
