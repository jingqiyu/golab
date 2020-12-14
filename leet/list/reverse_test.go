package list

import (
	"testing"
)

func Test_reverseList(t *testing.T) {
	list := GenList([]int{1, 2, 3}...)
	PrintList(list)

	list = reverseList(list)
	PrintList(list)
}

func Test_getKthFromEnd(t *testing.T) {
	list := GenList([]int{1, 2, 3, 4, 5}...)
	PrintList(list)
	list = getKthFromEnd(list, 5)
	PrintList(list)

}

func Test_mergeTwoLists(t *testing.T) {
	list := GenList([]int{1, 2, 3, 4, 5}...)
	list2 := GenList([]int{3, 4, 5, 6, 7}...)

	list = mergeTwoLists(list, list2)
	PrintList(list)
}
