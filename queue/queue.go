package queue

// Queue holds elements in a slice
type Queue struct {
	elements []interface{}
}

// Enqueue will add the specified elements to the queue
func (q *Queue) Enqueue(values ...interface{}) {
	q.elements = append(q.elements, values...)
}

// Dequeue retrieves the element at the front of the queue and removing it
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}

	element := q.elements[0]
	q.elements = q.elements[1:]

	return element, nil
}

// Peek retrieves the element at the front of the queue without removing it
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}

	return q.elements[0], nil
}

// Size returns size of the queue
func (q *Queue) Size() int {
	return len(q.elements)
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// New is a constructor for a new queue
func New() *Queue {
	return &Queue{
		elements: make([]interface{}, 0),
	}
}
