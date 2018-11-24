package list

// Element is an element of a singly linked list
type Element struct {
	Value interface{}
	next  *Element
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	return e.next
}

// List represents a singly linked list
type List struct {
	head *Element
	size uint
}

// Head returns the first element of list or nil if the list is empty
func (l *List) Head() *Element {
	return l.head
}

// Size returns the number of elements of list.
func (l *List) Size() uint {
	return l.size
}

// PushFront inserts a new element with value at the front of list
func (l *List) PushFront(value interface{}) *Element {
	el := &Element{value, l.head}

	l.head = el
	l.size++

	return el
}

// PopFront removes an element at the beginning of list and return removed element
func (l *List) PopFront() *Element {
	el := l.Head()

	nextEl := el.Next()
	if nextEl == nil {
		l.Clear()
	} else {
		l.head = nextEl
		l.size--
	}

	return el
}

// Clear removes all elements from the list.
func (l *List) Clear() {
	l.head = nil
	l.size = 0
}

// New is a constructor for a new list
func New() *List {
	return &List{
		head: nil,
		size: 0,
	}
}
