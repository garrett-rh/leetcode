package validParentheses

import (
	"strings"
)

type Stack struct {
	stackArray []int
}

func NewStack() *Stack {
	return &Stack{make([]int, 0)}
}

func (s *Stack) Push(c int) {
	s.stackArray = append(s.stackArray, c)
}

func (s *Stack) IsEmpty() bool {
	return len(s.stackArray) == 0
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	ret := s.stackArray[len(s.stackArray)-1]
	s.stackArray = s.stackArray[:len(s.stackArray)-1]

	return ret
}

func MinRemoveToMakeValid(s string) string {
	var builtString strings.Builder
	stack := NewStack()
	matches := make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else if s[i] == ')' {
			if !stack.IsEmpty() {
				matches[i] = true
				matches[stack.Pop()] = true
			}
		} else {
			matches[i] = true
		}
	}

	for i := 0; i < len(s); i++ {
		if matches[i] {
			builtString.WriteByte(s[i])
		}
	}

	if builtString.Len() == 0 {
		return ""
	}

	return builtString.String()
}
