package queue

import "errors"

var (
	// ErrEmptyQueue is returned when an non-applicable queue operation was called
	// due to the queue's empty element state
	ErrEmptyQueue = errors.New("queue: queue is empty")
)
