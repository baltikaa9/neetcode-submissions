// https://neetcode.io/problems/minimum-stack/

package main

type MinStack struct {
	vals []int
	mins []int
}

func (s *MinStack) Push(val int) {
	s.vals = append(s.vals, val)

	if len(s.mins) == 0 || val <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, val)
	}
}

func (s *MinStack) Pop() {
	if len(s.vals) > 0 {
		if s.vals[len(s.vals)-1] == s.mins[len(s.mins)-1] {
			s.mins = s.mins[:len(s.mins)-1]
		}
		s.vals = s.vals[:len(s.vals)-1]
	}
}

func (s *MinStack) Top() int {
	if len(s.vals) > 0 {
		return s.vals[len(s.vals)-1]
	}
	return 0
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}

func main() {
	s := MinStack{}
	s.Push(3)
	s.Push(2)
	s.Push(1)
	// println(s.Top())
	println(s.GetMin())
	s.Pop()
	// println(s.Top())
	println(s.GetMin())
}
