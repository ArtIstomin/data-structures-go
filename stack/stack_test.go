package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	assert := assert.New(t)
	s := New()
	execepted := []interface{}{"test1", "test2"}

	s.Enqueue("test1", "test2")

	assert.Len(s.elements, 2)
	assert.Equal(execepted, s.elements)
	assert.Equal(execepted[0], s.elements[0])
}
func TestDequeue(t *testing.T) {
	assert := assert.New(t)
	s := New()
	testElements := []interface{}{"test1", "test2"}

	s.Enqueue(testElements...)

	firstElement, err := s.Dequeue()
	assert.Nil(err)
	assert.Len(s.elements, 1)
	assert.Equal("test2", firstElement)
	assert.Equal([]interface{}{"test1"}, s.elements)

	secondElement, err := s.Dequeue()
	assert.Nil(err)
	assert.Len(s.elements, 0)
	assert.Equal("test1", secondElement)
	assert.Equal([]interface{}{}, s.elements)

	_, err = s.Dequeue()
	assert.EqualError(err, ErrEmptyStack.Error())
}

func TestPeek(t *testing.T) {
	assert := assert.New(t)
	s := New()
	testElements := []interface{}{"test1", "test2"}

	_, err := s.Peek()
	assert.EqualError(err, ErrEmptyStack.Error())

	s.Enqueue(testElements[0])
	element, err := s.Peek()
	assert.Nil(err)
	assert.Len(s.elements, 1)
	assert.Equal(testElements[0], element)
	assert.Equal(testElements[:1], s.elements)

	s.Enqueue(testElements[1])
	element, err = s.Peek()
	assert.Nil(err)
	assert.Len(s.elements, 2)
	assert.Equal(testElements[1], element)
	assert.Equal(testElements, s.elements)
}

func TestSize(t *testing.T) {
	assert := assert.New(t)
	s := New()

	size := s.Size()
	assert.Zero(size)

	s.Enqueue("test1", "test2")
	size = s.Size()
	assert.Equal(2, size)

	s.Dequeue()
	size = s.Size()
	assert.Equal(1, size)
}

func TestIsEmpty(t *testing.T) {
	assert := assert.New(t)
	s := New()

	isEmpty := s.IsEmpty()
	assert.True(isEmpty)

	s.Enqueue("test")
	isEmpty = s.IsEmpty()
	assert.False(isEmpty)
}
