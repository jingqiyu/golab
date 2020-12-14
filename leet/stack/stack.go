package stack

type Stack struct {
	Val []int
	top int
	len int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v int) {
	s.top = v
	s.Val = append(s.Val, v)
	s.len++
}

func (s *Stack) Peek() int {
	return s.top
}

func (s *Stack) Pop() (int, bool) {
	if s.len == 0 {
		return 0, false
	}

	tmp := s.top
	s.Val = s.Val[0 : s.len-1]
	s.len--
	if s.len != 0 {
		s.top = s.Val[s.len-1]
	}
	return tmp, true
}
