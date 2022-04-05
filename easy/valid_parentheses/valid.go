package validParentheses

type Stack struct {
	stackArray []rune
}

func NewStack() *Stack {
	return &Stack{make([]rune, 0)}
}

func (s *Stack) Push(c rune) {
	s.stackArray = append(s.stackArray, c)
}

func (s *Stack) IsEmpty() bool {
	return len(s.stackArray) == 0
}

func (s *Stack) Pop() rune {
	if s.IsEmpty() {
		return 0
	}

	ret := s.stackArray[s.Size()-1]
	s.stackArray = s.stackArray[:s.Size()-1]

	return ret
}

func (s *Stack) Size() int {
	return len(s.stackArray)
}

func (s *Stack) Peek() rune {
	if s.IsEmpty() {
		return 0
	}
	return s.stackArray[s.Size()-1]
}

func IsValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := NewStack()

	for _, char := range s {
		if _, found := pairs[char]; found {
			stack.Push(char)
		} else if pairs[stack.Peek()] != char {
			return false
		} else {
			stack.Pop()
		}
	}

	return stack.Size() == 0
}
