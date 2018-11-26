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
	size int
}

// Head returns the first element of list or nil if the list is empty
func (l *List) Head() *Element {
	return l.head
}

// Size returns the number of elements of list.
func (l *List) Size() int {
	return l.size
}

// Clear removes all elements from the list.
func (l *List) Clear() {
	l.head = nil
	l.size = 0
}

// PushFront inserts a new element with value at the front of list
func (l *List) PushFront(value interface{}) *Element {
	e := &Element{value, l.head}

	l.head = e
	l.size++

	return e
}

// PopFront removes an element at the beginning of list and return removed element
func (l *List) PopFront() *Element {
	e := l.Head()

	nextEl := e.Next()
	if nextEl == nil {
		l.Clear()
	} else {
		l.head = nextEl
		l.size--
	}

	return e
}

// Find returns the first element which matches or nil
func (l *List) Find(value interface{}) *Element {
	if l.Size() == 0 {
		return nil
	}

	for e := l.Head(); e != nil; e = e.Next() {
		if e.Value == value {
			return e
		}
	}

	return nil
}

// Remove removes the element at the given index from the list
func (l *List) Remove(index int) error {
	if !l.withinRange(index) {
		return ErrInvalidIndex
	}

	if l.Size() == 1 {
		l.Clear()
		return nil
	}

	var prev *Element
	el := l.Head()
	for i := 0; i != index; i, el = i+1, el.Next() {
		prev = el
	}

	if prev == nil {
		l.head = el.next
	} else {
		prev.next = el.next
	}

	el = nil
	l.size--

	return nil
}

// Get returns the element at index
func (l *List) Get(index int) (*Element, error) {
	if !l.withinRange(index) {
		return nil, ErrInvalidIndex
	}

	el := l.Head()
	for i := 0; i != index; i++ {
		el = el.Next()
	}

	return el, nil
}

// IndexOf returns index of value in the list
func (l *List) IndexOf(value interface{}) int {
	notFound := -1

	if l.Size() == 0 {
		return notFound
	}

	for e, i := l.Head(), 0; e != nil; e, i = e.Next(), i+1 {
		if e.Value == value {
			return i
		}
	}

	return notFound
}

// Contains checks if values are present in the list
func (l *List) Contains(value interface{}) bool {
	contains := false

	if l.Size() == 0 {
		return contains
	}

	for e := l.Head(); e != nil; e = e.Next() {
		if e.Value == value {
			contains = true
			break
		}
	}

	return contains
}

// New is a constructor for a new list
func New() *List {
	return &List{
		head: nil,
		size: 0,
	}
}

func (l *List) withinRange(index int) bool {
	return index >= 0 && index < l.Size()
}
