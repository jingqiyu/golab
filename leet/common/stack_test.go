package common

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Print()

	fmt.Println("stack peek = ", stack.Peek())

	pop, ok := stack.Pop()

	fmt.Println("stack pop=", pop, "ok=", ok)

	stack.Print()

}
