package stack

// Stack holds elements in a slice
type Stack struct {
	elements []interface{}
}

// Enqueue will add the specified elements to the stack
func (s *Stack) Enqueue(values ...interface{}) {
	s.elements = append(s.elements, values...)
}

// Dequeue retrieves the element at the end of the stack and removing it
func (s *Stack) Dequeue() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}

	element := s.elements[s.Size()-1]
	s.elements = s.elements[:s.Size()-1]

	return element, nil
}

// Peek retrieves the element at the front of the stack without removing it
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}

	return s.elements[s.Size()-1], nil
}

// Size returns size of the stack
func (s *Stack) Size() int {
	return len(s.elements)
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// New is a constructor for a new stack
func New() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}
