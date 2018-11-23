package stack

import "errors"

var (
	// ErrEmptyStack is returned when an non-applicable stack operation was called
	// due to the stack's empty element state
	ErrEmptyStack = errors.New("stack: stack is empty")
)
